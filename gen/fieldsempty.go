package gen

import (
	"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
)

func fieldsempty(w io.Writer, cfg *cfg.GreenConfig) *fieldsEmpty {
	return &fieldsEmpty{
		p:   printer{w: w},
		cfg: cfg,
	}
}

type fieldsEmpty struct {
	passes
	p     printer
	recvr string
	cfg   *cfg.GreenConfig
}

func (e *fieldsEmpty) MethodPrefix() string {
	return e.cfg.MethodPrefix
}

func (e *fieldsEmpty) Method() Method { return FieldsEmpty }

func (e *fieldsEmpty) Execute(p Elem) error {
	//	if DEBUG {
	//		fmt.Printf("Execute()-ing elem = p='%#v'\n", p)
	//	}
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

	e.recvr = fmt.Sprintf("%s %s", p.Varname(), imutMethodReceiver(p))

	next(e, p, nil)
	return e.p.err
}

func (e *fieldsEmpty) gStruct(s *Struct, x *extra) {
	if e.cfg.AllTuple {
		return
	}
	//	fmt.Printf("\n\n 77777 gStruct starting for fieldsempty s='%#v'\n\n ******* and e.recvr = '%#v'\n\n", s, e.recvr)
	e.p.printf("// %sfieldsNotEmpty supports omitempty tags\n", e.cfg.MethodPrefix)

	e.p.printf("func (%s) %sfieldsNotEmpty(isempty []bool) uint32 {",
		e.recvr, e.cfg.MethodPrefix)

	allFieldsEmpty := !e.cfg.SerzEmpty

	nfields := len(s.Fields) - s.SkipCount
	numOE := 0
	for i := range s.Fields {
		if s.Fields[i].OmitEmpty {
			numOE++
		}
	}

	// default is SerzEmpty false, so by
	// default we will treat all as being omitempty.
	// This is safe since we will zero any re-used
	// struct's fields when reading.
	//
	if e.cfg.SerzEmpty {
		// even under SerzEmpty, we still respect the specific ,omitempty tag:
		if numOE == 0 {
			// no fields tagged with omitempty, just return the full field count.
			e.p.printf("\nreturn %d }\n", nfields)
			return
		}
	}
	// remember this to avoid recomputing it in other passes.
	s.hasOmitEmptyTags = true

	om := emptyOmitter(&e.p, s.vname)

	e.p.printf("if len(isempty) == 0 { return %d }\n", nfields)
	e.p.printf("var fieldsInUse uint32 = %d\n", nfields)
	for i := range s.Fields {
		if s.Fields[i].Skip {
			continue
		}
		if allFieldsEmpty || s.Fields[i].OmitEmpty {
			e.p.printf("isempty[%d] = ", i)
			next(om, s.Fields[i].FieldElem, &extra{pointerOrIface: s.Fields[i].IsPointer || s.Fields[i].IsIface})

			e.p.printf("if isempty[%d] { fieldsInUse-- ; }\n", i)
			//or: e.p.printf("if isempty[%d] { fieldsInUse-- ; fmt.Printf(\"\\n %s is not in use!\\n \")}\n", i, s.Fields[i].FieldTagZidClue)
		}
	}
	//e.p.printf("\n fmt.Printf(\"\\n\\n fieldsInUse=%%v\", fieldsInUse) \n\n")
	e.p.printf("\n return fieldsInUse \n}\n")
}

func (e *fieldsEmpty) gPtr(p *Ptr, x *extra) {
	//next(e, p.Value)
}

func (e *fieldsEmpty) gSlice(sl *Slice, x *extra) {
	//next(e, sl.Els)
}

func (e *fieldsEmpty) gArray(a *Array, x *extra) {
	//fmt.Printf("\n\n 66666 gArray starting for fieldsempty a='%#v'\n\n", a)
	//next(e, a.Els)
}

func (e *fieldsEmpty) gMap(m *Map, x *extra) {
	//	fmt.Printf("\n\n 55555 gMap starting for fieldsempty m='%#v'\n\n", m)
	//next(e, m.Value)
}

func (e *fieldsEmpty) gBase(b *BaseElem, x *extra) {}
