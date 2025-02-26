package gen

import (
	"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/msgp"
)

func marshal(w io.Writer, cfg *cfg.GreenConfig) *marshalGen {
	return &marshalGen{
		p:   printer{w: w, cfg: cfg},
		cfg: cfg,
	}
}

type marshalGen struct {
	passes
	p    printer
	fuse []byte
	cfg  *cfg.GreenConfig
}

func (m *marshalGen) MethodPrefix() string {
	return m.cfg.MethodPrefix
}

func (m *marshalGen) Method() Method { return Marshal }

func (m *marshalGen) Apply(dirs []string) error {
	return nil
}

func (m *marshalGen) Execute(p Elem) error {
	if !m.p.ok() {
		return m.p.err
	}
	p = m.applyall(p)
	if p == nil {
		return nil
	}
	if !IsPrintable(p) {
		return nil
	}

	m.p.comment(fmt.Sprintf("%sMarshalMsg implements msgp.Marshaler", m.cfg.MethodPrefix))

	// save the vname before
	// calling methodReceiver so
	// that z.Msgsize() is printed correctly
	c := p.Varname()
	rcvr := imutMethodReceiver(p)
	//rcvr := methodReceiver(p)

	m.p.printf("\nfunc (%s %s) %sMarshalMsg(b []byte) (o []byte, err error) {", c, rcvr, m.cfg.MethodPrefix)
	hasPtr := false
	if hasPointersOrInterfaces(p) {
		hasPtr = true
		m.p.dedupWriteTop(true)
	}
	m.p.preSaveHook()
	m.p.printf("\no = msgp.Require(b, %s.%sMsgsize())", c, m.cfg.MethodPrefix)
	next(m, p, nil)
	if hasPtr {
		m.p.dedupWriteCleanup(true)
	}
	m.p.nakedReturn()
	return m.p.err
}

func (m *marshalGen) rawAppend(typ string, argfmt string, arg interface{}) {
	m.p.printf("\no = msgp.Append%s(o, %s)", typ, fmt.Sprintf(argfmt, arg))
}

func (m *marshalGen) fuseHook() {
	if len(m.fuse) > 0 {
		m.rawbytes(m.fuse)
		m.fuse = m.fuse[:0]
	}
}

func (m *marshalGen) Fuse(b []byte) {
	if len(m.fuse) == 0 {
		m.fuse = b
	} else {
		m.fuse = append(m.fuse, b...)
	}
}

func (m *marshalGen) gStruct(s *Struct, x *extra) {
	if !m.p.ok() {
		return
	}

	if m.cfg.AllTuple || s.AsTuple {
		m.tuple(s)
	} else {
		m.mapstruct(s)
	}
	return
}

func (m *marshalGen) tuple(s *Struct) {
	data := make([]byte, 0, 5)
	data = msgp.AppendArrayHeader(data, uint32(len(s.Fields)-s.SkipCount))
	m.p.printf("\n// array header, size %d", len(s.Fields)-s.SkipCount)
	m.Fuse(data)
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if !m.p.ok() {
			return
		}
		next(m, s.Fields[i].FieldElem, nil)
	}
}

func (m *marshalGen) mapstruct(s *Struct) {
	data := make([]byte, 0, 64)

	nfields := len(s.Fields) - s.SkipCount

	allOmitEmpty := !m.cfg.SerzEmpty
	skipclue := m.cfg.SkipZidClue || m.cfg.Msgpack2

	if allOmitEmpty || s.hasOmitEmptyTags {
		m.p.printf("\n\n// honor the omitempty tags\n")
		m.p.printf("var empty [%d]bool\n", len(s.Fields))
		m.p.printf("fieldsInUse := %s.fieldsNotEmpty(empty[:])\n", s.vname)
		m.p.printf("	o = msgp.AppendMapHeader(o, fieldsInUse)\n")
	} else {
		data = msgp.AppendMapHeader(data, uint32(nfields))
		m.p.printf("\n// map header, size %d", nfields)
		m.Fuse(data)
		m.fuseHook()
	}

	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if s.Fields[i].NeedsReflection {
			m.p.printf("\n// '%s' generic => reflection\n", s.Fields[i].FieldName)
			m.Fuse(data)
			m.fuseHook()
			continue
		}

		if !m.p.ok() {
			return
		}
		if allOmitEmpty || (s.hasOmitEmptyTags && s.Fields[i].OmitEmpty) {
			m.p.printf("\n if !empty[%d] {", i)
		}

		fld := s.Fields[i].FieldTagZidClue
		if skipclue {
			fld = s.Fields[i].FieldTag
		}

		switch s.KeyTyp {
		case "Int64":
			data = msgp.AppendInt64(nil, s.Fields[i].ZebraId)
		default:
			data = msgp.AppendString(nil, fld)
		}

		m.p.printf("\n// string %q", fld)
		m.Fuse(data)
		m.fuseHook()
		next(m, s.Fields[i].FieldElem, &extra{pointerOrIface: s.Fields[i].IsPointer || s.Fields[i].IsIface})

		if allOmitEmpty || (s.hasOmitEmptyTags && s.Fields[i].OmitEmpty) {
			m.p.printf("\n }\n")
		}
	}
}

// append raw data
func (m *marshalGen) rawbytes(bts []byte) {
	m.p.print("\no = append(o, ")
	for _, b := range bts {
		m.p.printf("0x%x,", b)
	}
	m.p.print(")")
}

func (m *marshalGen) gMap(s *Map, x *extra) {
	if !m.p.ok() {
		return
	}
	m.fuseHook()
	vname := s.Varname()
	m.rawAppend(mapHeader, lenAsUint32, vname)
	m.p.printf("\nfor %s, %s := range %s {", s.Keyidx, s.Validx, vname)
	m.rawAppend(s.KeyTyp, literalFmt, s.Keyidx)
	next(m, s.Value, nil)
	m.p.closeblock()
}

func (m *marshalGen) gSlice(s *Slice, x *extra) {
	if !m.p.ok() {
		return
	}
	m.fuseHook()
	vname := s.Varname()
	m.rawAppend(arrayHeader, lenAsUint32, vname)
	m.p.rangeBlock(s.Index, vname, m, s.Els)
}

func (m *marshalGen) gArray(a *Array, x *extra) {
	if !m.p.ok() {
		return
	}
	m.fuseHook()
	if be, ok := a.Els.(*BaseElem); ok && be.Value == Byte {
		m.rawAppend("Bytes", "%s[:]", a.Varname())
		return
	}

	m.rawAppend(arrayHeader, literalFmt, a.SizeResolved)
	m.p.rangeBlock(a.Index, a.Varname(), m, a.Els)
}

func (m *marshalGen) gPtr(p *Ptr, x *extra) {
	if !m.p.ok() {
		return
	}
	m.fuseHook()
	m.p.printf("\n // marshalGen.gPtr() \n")
	m.p.printf("\nif %s == nil {\no = msgp.AppendNil(o)\n} else { ", p.Varname())
	m.p.printf("\n // hmm.. no en, no place to check en.DedupWriteIsDup(z)\n")
	next(m, p.Value, nil)
	m.p.closeblock()
}

func (m *marshalGen) gBase(b *BaseElem, x *extra) {
	if !m.p.ok() {
		return
	}
	m.fuseHook()
	vname := b.Varname()

	if b.Convert {
		vname = tobaseConvert(b)
	}

	var echeck bool
	switch b.Value {
	case IDENT:
		echeck = true
		if b.isIface {
			echeck = true
			m.p.printf("\no, err = msgp.AppendIntf(o, %s) // is.iface", vname)
		} else {
			m.p.printf("\no, err = %s.%sMarshalMsg(o) // not is.iface", vname, m.cfg.MethodPrefix)
		}
	case Intf, Ext:
		echeck = true
		m.p.printf("\no, err = msgp.Append%s(o, %s)", b.BaseName(), vname)
	default:
		m.rawAppend(b.BaseName(), literalFmt, vname)
	}

	if echeck {
		m.p.print(errcheck)
	}
}
