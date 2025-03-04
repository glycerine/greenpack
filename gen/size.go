package gen

import (
	"fmt"
	"io"
	"strconv"

	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/msgp"
)

type sizeState uint8

const (
	// need to write "s = ..."
	assign sizeState = iota

	// need to write "s += ..."
	add

	// can just append "+ ..."
	expr
)

func sizes(w io.Writer, cfg *cfg.GreenConfig) *sizeGen {
	return &sizeGen{
		p:     printer{w: w, cfg: cfg},
		state: assign,
		cfg:   cfg,
	}
}

type sizeGen struct {
	passes
	p     printer
	state sizeState
	cfg   *cfg.GreenConfig
}

func (s *sizeGen) MethodPrefix() string {
	return s.cfg.MethodPrefix
}

func (s *sizeGen) Method() Method { return Size }

func (s *sizeGen) Apply(dirs []string) error {
	return nil
}

func builtinSize(typ string) string {
	return "msgp." + typ + "Size"
}

// this lets us chain together addition
// operations where possible
func (s *sizeGen) addConstant(sz string) {
	if !s.p.ok() {
		return
	}

	switch s.state {
	case assign:
		s.p.print("\ns = " + sz)
		s.state = expr
		return
	case add:
		s.p.print("\ns += " + sz)
		s.state = expr
		return
	case expr:
		s.p.print(" + " + sz)
		return
	}

	panic("unknown size state")
}

func (s *sizeGen) Execute(p Elem) error {
	if !s.p.ok() {
		return s.p.err
	}
	p = s.applyall(p)
	if p == nil {
		return nil
	}
	if !IsPrintable(p) {
		return nil
	}

	s.p.comment(fmt.Sprintf("%sMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message", s.cfg.MethodPrefix))

	s.p.printf("\nfunc (%s %s) %sMsgsize() (s int) {", p.Varname(), imutMethodReceiver(p), s.cfg.MethodPrefix)
	s.state = assign
	next(s, p, nil)
	s.p.nakedReturn()
	return s.p.err
}

func (s *sizeGen) gStruct(st *Struct, x *extra) {
	if !s.p.ok() {
		return
	}
	skipclue := s.cfg.SkipZidClue || s.cfg.Msgpack2

	nfields := uint32(len(st.Fields) - st.SkipCount)

	if s.cfg.AllTuple || st.AsTuple {
		data := msgp.AppendArrayHeader(nil, nfields)
		s.addConstant(strconv.Itoa(len(data)))
		for i := range st.Fields {
			if st.Fields[i].Skip {
				continue
			}

			if !s.p.ok() {
				return
			}
			next(s, st.Fields[i].FieldElem, nil)
		}
	} else {
		data := msgp.AppendMapHeader(nil, nfields)
		s.addConstant(strconv.Itoa(len(data)))
		for i := range st.Fields {
			if st.Fields[i].Skip {
				continue
			}
			data = data[:0]

			var appendme string
			if skipclue {
				appendme = st.Fields[i].FieldTag
			} else {
				appendme = st.Fields[i].FieldTagZidClue
			}
			data = msgp.AppendString(data, appendme)

			s.addConstant(strconv.Itoa(len(data)))
			next(s, st.Fields[i].FieldElem, nil)
		}
	}
}

func (s *sizeGen) gPtr(p *Ptr, x *extra) {
	s.state = add // inner must use add
	s.p.printf("\nif %s == nil {\ns += msgp.NilSize\n} else {", p.Varname())
	next(s, p.Value, nil)
	s.state = add // closing block; reset to add
	s.p.closeblock()
}

func (s *sizeGen) gSlice(sl *Slice, x *extra) {
	if !s.p.ok() {
		return
	}

	s.addConstant(builtinSize(arrayHeader))

	// if the slice's element is a fixed size
	// (e.g. float64, [32]int, etc.), then
	// print the length times the element size directly
	if str, ok := fixedsizeExpr(sl.Els); ok {
		s.addConstant(fmt.Sprintf("(%s * (%s))", lenExpr(sl), str))
		return
	}

	// add inside the range block, and immediately after
	s.state = add
	s.p.rangeBlock(sl.Index, sl.Varname(), s, sl.Els)
	s.state = add
}

func (s *sizeGen) gArray(a *Array, x *extra) {
	if !s.p.ok() {
		return
	}
	//vv("gArray called; a.Els = '%#v'", a.Els)

	s.addConstant(builtinSize(arrayHeader))

	// if the array's children are a fixed
	// size, we can compile an expression
	// that always represents the array's wire size
	if str, ok := fixedsizeExpr(a); ok {
		s.addConstant(str)
		return
	}

	s.state = add
	s.p.rangeBlock(a.Index, a.Varname(), s, a.Els)
	s.state = add
}

func (s *sizeGen) gMap(m *Map, x *extra) {
	s.addConstant(builtinSize(mapHeader))
	vn := m.Varname()
	s.p.printf("\nif %s != nil {", vn)
	s.p.printf("\nfor %s, %s := range %s {", m.Keyidx, m.Validx, vn)
	s.p.printf("\n_ = %s", m.Validx) // we may not use the value
	s.p.printf("\n_ = %s", m.Keyidx) // we may not use the value
	switch m.KeyTyp {
	case "String":
		s.p.printf("\ns += msgp.StringPrefixSize + len(%s)", m.Keyidx)
	default:
		s.p.printf("\ns += msgp.%sSize", m.KeyTyp)
	}
	s.state = expr
	next(s, m.Value, nil)
	s.p.closeblock()
	s.p.closeblock()
	s.state = add
}

func (s *sizeGen) gBase(b *BaseElem, x *extra) {
	if !s.p.ok() {
		return
	}
	s.addConstant(basesizeExpr(b, s))
}

// returns "len(slice)"
func lenExpr(sl *Slice) string {
	return "len(" + sl.Varname() + ")"
}

// is a given primitive always the same (max)
// size on the wire?
func fixedSize(p Primitive) bool {
	switch p {
	case Intf, Ext, IDENT, Bytes, String:
		return false
	default:
		return true
	}
}

// strip reference from string
func stripRef(s string) string {
	if s[0] == '&' {
		return s[1:]
	}
	return s
}

// return a fixed-size expression, if possible.
// only possible for *BaseElem and *Array.
// returns (expr, ok)
func fixedsizeExpr(e Elem) (string, bool) {
	switch e := e.(type) {
	case *Array:
		if str, ok := fixedsizeExpr(e.Els); ok {
			return fmt.Sprintf("(%s * (%s))", e.SizeResolved, str), true
		}
	case *BaseElem:
		if fixedSize(e.Value) {
			return builtinSize(e.BaseName()), true
		}
	case *Struct:
		var str string
		for _, f := range e.Fields {
			if f.Skip {
				continue
			}

			if fs, ok := fixedsizeExpr(f.FieldElem); ok {
				if str == "" {
					str = fs
				} else {
					str += "+" + fs
				}
			} else {
				return "", false
			}
		}
		var hdrlen int
		mhdr := msgp.AppendMapHeader(nil, uint32(len(e.Fields)-e.SkipCount))
		hdrlen += len(mhdr)
		var strbody []byte
		for _, f := range e.Fields {
			if !f.Skip {

				strbody = msgp.AppendString(strbody[:0], f.FieldTagZidClue)
				hdrlen += len(strbody)
			}
		}
		return fmt.Sprintf("%d + %s", hdrlen, str), true
	}
	return "", false
}

// print size expression of a variable name
func basesizeExpr(b *BaseElem, s *sizeGen) string {
	vname := b.Varname()
	if b.Convert {
		vname = tobaseConvert(b)
	}
	if b.IsIface {
		return "msgp.GuessSize(" + vname + ")"
	}

	switch b.Value {
	case Ext:
		return "msgp.ExtensionPrefixSize + " + stripRef(vname) + ".Len()"
	case Intf:
		return "msgp.GuessSize(" + vname + ")"
	case IDENT:
		return vname + fmt.Sprintf(".%sMsgsize()", s.cfg.MethodPrefix)
	case Bytes:
		return "msgp.BytesPrefixSize + len(" + vname + ")"
	case String:
		return "msgp.StringPrefixSize + len(" + vname + ")"
	default:
		return builtinSize(b.BaseName())
	}
}
