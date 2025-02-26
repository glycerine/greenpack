package gen

import (
	"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/msgp"
)

func encode(w io.Writer, cfg *cfg.GreenConfig) *encodeGen {
	return &encodeGen{
		p:   printer{w: w, cfg: cfg},
		cfg: cfg,
	}
}

type encodeGen struct {
	passes
	p    printer
	fuse []byte
	cfg  *cfg.GreenConfig
}

func (e *encodeGen) MethodPrefix() string {
	return e.cfg.MethodPrefix
}

func (e *encodeGen) Method() Method { return Encode }

func (e *encodeGen) Apply(dirs []string) error {
	return nil
}

func (e *encodeGen) writeAndCheck(typ string, argfmt string, arg interface{}) {
	e.p.printf("\nerr = en.Write%s(%s)", typ, fmt.Sprintf(argfmt, arg))
	e.p.print(errcheck)
}

func (e *encodeGen) fuseHook() {
	if len(e.fuse) > 0 {
		e.appendraw(e.fuse)
		e.fuse = e.fuse[:0]
	}
}

func (e *encodeGen) Fuse(b []byte) {
	if len(e.fuse) > 0 {
		e.fuse = append(e.fuse, b...)
	} else {
		e.fuse = b
	}
}

func (e *encodeGen) Execute(p Elem) error {
	if !e.p.ok() {
		return e.p.err
	}
	p = e.applyall(p)
	if p == nil {
		return nil
	}
	if !IsPrintable(p) {
		return nil
	}

	e.p.comment(fmt.Sprintf("%sEncodeMsg implements msgp.Encodable", e.cfg.MethodPrefix))

	varname := p.Varname()
	rcvr := imutMethodReceiver(p)
	//rcvr := methodReceiver(p)

	//e.p.printf("\nfunc (%s %s) %sEncodeMsg(en *msgp.Writer) (err error) {", p.Varname(), imutMethodReceiver(p), e.cfg.MethodPrefix)
	e.p.printf("\nfunc (%s %s) %sEncodeMsg(en *msgp.Writer) (err error) {", varname, rcvr, e.cfg.MethodPrefix)

	hasPtr := false
	if hasPointersOrInterfaces(p) {
		hasPtr = true
		//e.p.dedupWriteTop(false)
	}
	e.p.preSaveHook()
	next(e, p, nil)
	if hasPtr {
		e.p.dedupWriteCleanup(false)
	}
	e.p.nakedReturn()
	return e.p.err
}

func (e *encodeGen) gStruct(s *Struct, x *extra) {
	if !e.p.ok() {
		return
	}

	if e.cfg.AllTuple || s.AsTuple {
		e.tuple(s)
	} else {
		e.structmap(s)
	}
	return
}

func (e *encodeGen) tuple(s *Struct) {
	nfields := len(s.Fields) - s.SkipCount
	data := msgp.AppendArrayHeader(nil, uint32(nfields))
	e.p.printf("\n// array header, size %d", nfields)
	e.Fuse(data)
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if !e.p.ok() {
			return
		}
		next(e, s.Fields[i].FieldElem, nil)
	}
}

func (e *encodeGen) appendraw(bts []byte) {
	e.p.print("\nerr = en.Append(")
	for i, b := range bts {
		if i != 0 {
			e.p.print(", ")
		}
		e.p.printf("0x%x", b)
	}
	e.p.print(")\nif err != nil { return err }")
}

func (e *encodeGen) structmap(s *Struct) {
	nfields := len(s.Fields) - s.SkipCount
	var data []byte
	empty := "empty_" + gensym()
	inUse := "fieldsInUse_" + gensym()

	allOmitEmpty := !e.cfg.SerzEmpty
	skipclue := e.cfg.SkipZidClue || e.cfg.Msgpack2

	if allOmitEmpty || s.hasOmitEmptyTags {
		e.p.printf("\n\n// honor the omitempty tags\n")
		e.p.printf("var %s [%d]bool\n", empty, len(s.Fields))
		e.p.printf("%s := %s.fieldsNotEmpty(%s[:])\n",
			inUse, s.vname, empty)
		e.p.printf("\n// map header\n")

		if !e.cfg.NoEmbeddedStructNames {
			// +1 for the -1:structName type-identifier.
			e.p.printf("	err = en.WriteMapHeader(%s+1)\n", inUse)
		} else {
			e.p.printf("	err = en.WriteMapHeader(%s)\n", inUse)
		}

		e.p.printf("	if err != nil {\n")
		e.p.printf("		return err\n}\n")
	} else {
		data = msgp.AppendMapHeader(nil, uint32(nfields))
		e.p.printf("\n// map header, size %d", nfields)
		e.Fuse(data)
		e.fuseHook()
	}

	if !e.cfg.NoEmbeddedStructNames {
		// record the struct name under integer key -1
		recv := s.TypeName() // imutMethodReceiver(s)?
		e.p.printf("\n// runtime struct type identification for '%s'\n", recv)
		hexname := ""
		for i := range recv {
			hexname += fmt.Sprintf("0x%x,", recv[i])
		}
		e.p.printf("err = en.Append(0xa1, 0x40)\n") // "@" sign key, value has type name
		e.p.printf("	if err != nil {\n")
		e.p.printf("		return err\n}\n")

		e.p.printf("err = en.WriteStringFromBytes([]byte{%s})\n", hexname)
		e.p.printf("	if err != nil {\n")
		e.p.printf("		return err\n}\n")
	}

	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if s.Fields[i].NeedsReflection {
			e.p.printf("\n// '%s' generic => reflection\n", s.Fields[i].FieldName)
			continue
		}
		if !e.p.ok() {
			return
		}

		if allOmitEmpty || (s.hasOmitEmptyTags && s.Fields[i].OmitEmpty) {
			e.p.printf("\n if !%s[%d] {", empty, i)
		}

		fld := s.Fields[i].FieldTagZidClue
		if skipclue {
			fld = s.Fields[i].FieldTag
		}
		data = msgp.AppendString(nil, fld)
		e.p.printf("\n// write %q", fld)
		e.Fuse(data)
		e.fuseHook()
		next(e, s.Fields[i].FieldElem, &extra{pointerOrIface: s.Fields[i].IsPointer || s.Fields[i].IsIface})

		if allOmitEmpty || (s.hasOmitEmptyTags && s.Fields[i].OmitEmpty) {
			e.p.printf("\n }\n")
		}
	}
}

func (e *encodeGen) gMap(m *Map, x *extra) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := m.Varname()
	e.writeAndCheck(mapHeader, lenAsUint32, vname)

	e.p.printf("\nfor %s, %s := range %s {", m.Keyidx, m.Validx, vname)
	e.writeAndCheck(m.KeyTyp, literalFmt, m.Keyidx)
	next(e, m.Value, nil)
	e.p.closeblock()
}

func (e *encodeGen) gPtr(s *Ptr, x *extra) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.p.print("\n // gPtr.encodeGen():\n")
	e.p.printf("\nif %s == nil { err = en.WriteNil(); if err != nil { return; } } else {", s.Varname())

	doDedupHere := !e.cfg.NoDedup
	if doDedupHere {
		switch s.Value.(type) {
		case *BaseElem:
			// not inlined; don't dedup here as
			// the interface/pointer/struct will handle
			// dedup itself.
			doDedupHere = false
		default:
			e.p.printf("\n // record the pointer for deduplication  \n")
			e.p.printf(` var dup bool
 dup, err = en.DedupWriteIsDup(%s)
			if err != nil {
				return
			}
			if !dup {
			`, s.Varname())
		}
	}
	next(e, s.Value, nil)
	if doDedupHere {
		e.p.closeblock()
	}
	e.p.closeblock()
}

func (e *encodeGen) gSlice(s *Slice, x *extra) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.writeAndCheck(arrayHeader, lenAsUint32, s.Varname())
	e.p.rangeBlock(s.Index, s.Varname(), e, s.Els)
}

func (e *encodeGen) gArray(a *Array, x *extra) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	// shortcut for [const]byte
	if be, ok := a.Els.(*BaseElem); ok && (be.Value == Byte || be.Value == Uint8) {
		e.p.printf("\nerr = en.WriteBytes(%s[:])", a.Varname())
		e.p.print(errcheck)
		return
	}

	e.writeAndCheck(arrayHeader, literalFmt, a.SizeResolved)
	e.p.rangeBlock(a.Index, a.Varname(), e, a.Els)
}

func (e *encodeGen) gBase(b *BaseElem, x *extra) {
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := b.Varname()
	if b.Convert {
		vname = tobaseConvert(b)
	}

	if b.Value == IDENT { // unknown identity
		if !e.cfg.NoDedup {
			e.p.printf("\n // encodeGen.gBase IDENT \n")
			e.p.printf(`
		// record the interface for deduplication
		var dup bool
		dup, err = en.DedupWriteIsDup(%s)
		if err != nil {
			return
		}
		if !dup {`, vname)
		}
		e.p.printf("\nerr = %s.%sEncodeMsg(en)", vname, e.cfg.MethodPrefix)
		e.p.print(errcheck)
		if !e.cfg.NoDedup {
			e.p.closeblock()
		}
	} else { // typical case
		e.writeAndCheck(b.BaseName(), literalFmt, vname)
	}
}
