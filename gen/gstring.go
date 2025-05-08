package gen

import (
	//"fmt"
	"io"
	"strings"

	"github.com/glycerine/greenpack/cfg"
)

func gstring(w io.Writer, cfg *cfg.GreenConfig) *gstringGen {
	return &gstringGen{
		p:   printer{w: w, cfg: cfg},
		cfg: cfg,
	}
}

type gstringGen struct {
	passes
	p    printer
	fuse []byte
	cfg  *cfg.GreenConfig
}

func (e *gstringGen) MethodPrefix() string {
	return e.cfg.MethodPrefix
}

func (e *gstringGen) Method() Method { return Gstring }

func (e *gstringGen) Apply(dirs []string) error {
	return nil
}

func (e *gstringGen) writeAndCheck(typ string, argfmt string, arg interface{}) {
	//e.p.printf("\nerr = en.Write%s(%s)", typ, fmt.Sprintf(argfmt, arg))
	//e.p.printf("%s %s", arg, typ)
}

func (e *gstringGen) fuseHook() {
	if len(e.fuse) > 0 {
		e.appendraw(e.fuse)
		e.fuse = e.fuse[:0]
	}
}

func (e *gstringGen) Fuse(b []byte) {
	if len(e.fuse) > 0 {
		e.fuse = append(e.fuse, b...)
	} else {
		e.fuse = b
	}
}

func (e *gstringGen) Execute(p Elem) error {
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

	// only gen Gstring for structs
	_, isStruct := p.(*Struct)
	if !isStruct {
		return nil
	}

	e.p.printf(`func (%s %s) %sGstring() (r string) {
`, p.Varname(), imutMethodReceiver(p), e.cfg.MethodPrefix)

	e.p.printf(`
		r = "&%v{\n"
	`, p.TypeName())

	next(e, p, nil)
	e.p.printf(` r += "}\n"`)
	e.p.nakedReturn()
	return e.p.err
}

func (e *gstringGen) gStruct(s *Struct, x *extra) {
	if !e.p.ok() {
		return
	}

	e.structmap(s)

	return
}

func (e *gstringGen) appendraw(bts []byte) {
	return
	e.p.print("\nerr = en.Append(")
	for i, b := range bts {
		if i != 0 {
			e.p.print(", ")
		}
		e.p.printf("0x%x", b)
	}
	e.p.print(")\nif err != nil { return err }")
}

func (e *gstringGen) structmap(s *Struct) {

	// align ":" in output like R.
	longest := 0
	for i := range s.Fields {
		x := len(s.Fields[i].FieldName)
		if x > longest {
			longest = x
		}
	}
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if !e.p.ok() {
			return
		}

		f := s.Fields[i].FieldName
		spaces := strings.Repeat(" ", longest-len(f))
		if s.Fields[i].FieldElem.TypeName() == "string" {
			e.p.printf(`r += fmt.Sprintf("%v%v: \"%%v\",\n", %v.%v)
`, spaces, f, s.Varname(), f)
		} else {
			e.p.printf(`r += fmt.Sprintf("%v%v: %%v,\n", %v.%v)
`, spaces, f, s.Varname(), f)
		}
		// type-switch dispatch to the correct method
		//next(e, s.Fields[i].FieldElem, &extra{pointerOrIface: s.Fields[i].IsPointer || s.Fields[i].IsIface})

	}
	//e.p.printf(``)

}

func (e *gstringGen) gMap(m *Map, x *extra) {
	return
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

func (e *gstringGen) gPtr(s *Ptr, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.p.print("\n // gPtr.gstringGen():\n")
	e.p.printf("\nif %s == nil { err = en.WriteNil(); if err != nil { return; } } else {", s.Varname())

	next(e, s.Value, nil)
	e.p.closeblock()
}

func (e *gstringGen) gSlice(s *Slice, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	e.writeAndCheck(arrayHeader, lenAsUint32, s.Varname())
	e.p.rangeBlock(s.Index, s.Varname(), e, s.Els)
}

func (e *gstringGen) gArray(a *Array, x *extra) {
	return
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

func (e *gstringGen) gBase(b *BaseElem, x *extra) {
	return
	if !e.p.ok() {
		return
	}
	e.fuseHook()
	vname := b.Varname()
	//if b.Convert {
	//	vname = tobaseConvert(b)
	//}

	if b.Value == IDENT { // unknown identity
		e.p.printf("\n // gstringGen.gBase unknown IDENT '%v': ignore for now.\n", vname)
		e.p.printf("\n // %s.%sGstring()", vname, e.cfg.MethodPrefix)
	} else { // typical case
		e.writeAndCheck(b.BaseName(), literalFmt, vname)
	}
}
