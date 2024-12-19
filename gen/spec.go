package gen

import (
	"fmt"
	"io"

	"github.com/glycerine/greenpack/cfg"
)

const (
	errcheck = "\nif err != nil { return }"
	//errcheck    = "\nif err != nil { panic(err) }"
	lenAsUint32 = "uint32(len(%s))"
	literalFmt  = "%s"
	intFmt      = "%d"
	quotedFmt   = `"%s"`
	mapHeader   = "MapHeader"
	arrayHeader = "ArrayHeader"
	mapKey      = "MapKeyPtr"
	stringTyp   = "String"
	int64Typ    = "Int64"
	u32         = "uint32"
)

// Method is a bitfield representing something that the
// generator knows how to print.
type Method uint16

// are the bits in 'f' set in 'm'?
func (m Method) isset(f Method) bool { return (m&f == f) }

// String implements fmt.Stringer
func (m Method) String() string {
	switch m {
	case 0, invalidmeth:
		return "<invalid method>"
	case Decode:
		return "decode"
	case Encode:
		return "encode"
	case Marshal:
		return "marshal"
	case Unmarshal:
		return "unmarshal"
	case Size:
		return "size"
	case Test:
		return "test"
	case FieldsEmpty:
		return "fieldsempty"
	case StoreToSQL:
		return "storeToSQL"
	case GetFromSQL:
		return "getFromSQL"
	default:
		// return e.g. "decode+encode+test"
		modes := [...]Method{Decode, Encode, Marshal, Unmarshal, Size, Test, StoreToSQL}
		any := false
		nm := ""
		for _, mm := range modes {
			if m.isset(mm) {
				if any {
					nm += "+" + mm.String()
				} else {
					nm += mm.String()
					any = true
				}
			}
		}
		return nm

	}
}

func strtoMeth(s string) Method {
	switch s {
	case "encode":
		return Encode
	case "decode":
		return Decode
	case "marshal":
		return Marshal
	case "unmarshal":
		return Unmarshal
	case "size":
		return Size
	case "test":
		return Test
	case "fieldsempty":
		return FieldsEmpty
	case "storeToSQL":
		return StoreToSQL
	case "getFromSQL":
		return GetFromSQL
	default:
		return 0
	}
}

const (
	Decode      Method = 1 << iota // msgp.Decodable
	Encode                         // msgp.Encodable
	Marshal                        // msgp.Marshaler
	Unmarshal                      // msgp.Unmarshaler
	Size                           // msgp.Sizer
	Test                           // generate tests
	FieldsEmpty                    // support omitempty tag
	StoreToSQL                     // from struct to SQL database.
	GetFromSQL                     // from SQL database to struct
	invalidmeth                    // this isn't a method

	encodetest  = Encode | Decode | Test | FieldsEmpty     // tests for Encodable and Decodable
	marshaltest = Marshal | Unmarshal | Test | FieldsEmpty // tests for Marshaler and Unmarshaler
)

type Printer struct {
	gens []generator
	cfg  *cfg.GreenConfig
}

func NewPrinter(m Method, out io.Writer, tests io.Writer, cfg *cfg.GreenConfig) *Printer {
	if m.isset(Test) && tests == nil {
		panic("cannot print tests with 'nil' tests argument!")
	}
	gens := make([]generator, 0, 8)
	if m.isset(Decode) {
		gens = append(gens, decode(out, cfg))
	}
	// must run FieldsEmpty before Encode/Marshal, so as
	// to set Struct.hasOmitEmptyTags
	if m.isset(FieldsEmpty) {
		gens = append(gens, fieldsempty(out, cfg))
	}
	if m.isset(Encode) {
		gens = append(gens, encode(out, cfg))
	}
	if m.isset(Marshal) {
		gens = append(gens, marshal(out, cfg))
	}
	if m.isset(Unmarshal) {
		gens = append(gens, unmarshal(out, cfg))
	}
	if m.isset(Size) {
		gens = append(gens, sizes(out, cfg))
	}
	if m.isset(StoreToSQL) {
		gens = append(gens, storeToSQL(out, cfg))
	}
	if m.isset(GetFromSQL) {
		gens = append(gens, getFromSQL(out, cfg))
	}
	if m.isset(marshaltest) {
		gens = append(gens, mtest(tests, cfg))
	}
	if m.isset(encodetest) {
		gens = append(gens, etest(tests, cfg))
	}
	if len(gens) == 0 {
		panic("NewPrinter called with invalid method flags")
	}
	return &Printer{gens: gens, cfg: cfg}
}

// TransformPass is a pass that transforms individual
// elements. (Note that if the returned is different from
// the argument, it should not point to the same objects.)
type TransformPass func(Elem) Elem

// IgnoreTypename is a pass that just ignores
// types of a given name.
func IgnoreTypename(name string) TransformPass {
	return func(e Elem) Elem {
		if e.TypeName() == name {
			return nil
		}
		return e
	}
}

// ApplyDirective applies a directive to a named pass
// and all of its dependents.
func (p *Printer) ApplyDirective(pass Method, t TransformPass) {
	for _, g := range p.gens {
		if g.Method().isset(pass) {
			g.Add(t)
		}
	}
}

// Print prints an Elem.
func (p *Printer) Print(e Elem) error {
	for _, g := range p.gens {
		//		fmt.Printf("\n\n 99999 about to begin executing the gen g='%#v'\n   ---------> on element e='%#v'.\n\n", g, e)

		err := g.Execute(e)
		if err != nil {
			return err
		}
	}
	return nil
}

// generator is the interface through
// which code is generated.
type generator interface {
	Method() Method
	Add(p TransformPass)
	Execute(Elem) error // execute writes the method for the provided object.
}

type passes []TransformPass

func (p *passes) Add(t TransformPass) {
	*p = append(*p, t)
}

func (p *passes) applyall(e Elem) Elem {
	for _, t := range *p {
		e = t(e)
		if e == nil {
			return nil
		}
	}
	return e
}

type traversal interface {
	gMap(*Map, *extra)
	gSlice(*Slice, *extra)
	gArray(*Array, *extra)
	gPtr(*Ptr, *extra)
	gBase(*BaseElem, *extra)
	gStruct(*Struct, *extra)
}

type extra struct {
	pointerOrIface bool
}

// type-switch dispatch to the correct
// method given the type of 'e'
func next(t traversal, e Elem, x *extra) {
	//	if DEBUG {
	//		fmt.Printf("\n top of next(), e='%#v'.\n =========>>  traversal='%#v'\n", e, t)
	//	}

	switch e := e.(type) {
	case *Map:
		t.gMap(e, x)
	case *Struct:
		t.gStruct(e, x)
	case *Slice:
		t.gSlice(e, x)
	case *Array:
		t.gArray(e, x)
	case *Ptr:
		t.gPtr(e, x)
	case *BaseElem:
		t.gBase(e, x)
	default:
		panic("bad element type")
	}
}

// possibly-immutable method receiver
func imutMethodReceiver(p Elem) string {
	switch e := p.(type) {
	case *Struct:
		// TODO(HACK): actually do real math here.
		if len(e.Fields)-e.SkipCount <= 3 {
			for i := range e.Fields {
				if e.Fields[i].Skip {
					continue
				}
				if be, ok := e.Fields[i].FieldElem.(*BaseElem); !ok || (be.Value == IDENT || be.Value == Bytes) {
					goto nope
				}
			}
			return p.TypeName()
		}
	nope:
		return "*" + p.TypeName()

	// gets dereferenced automatically
	case *Array:
		return "*" + p.TypeName()

	// everything else can be
	// by-value.
	default:
		return p.TypeName()
	}
}

// if necessary, wraps a type
// so that its method receiver
// is of the write type.
func methodReceiver(p Elem) string {
	switch p.(type) {

	// structs and arrays are
	// dereferenced automatically,
	// so no need to alter varname
	case *Struct, *Array:
		return "*" + p.TypeName()
	// set variable name to
	// *varname
	default:
		p.SetVarname("(*" + p.Varname() + ")")
		return "*" + p.TypeName()
	}
}

func unsetReceiver(p Elem) {
	switch p.(type) {
	case *Struct, *Array:
	default:
		p.SetVarname("z")
	}
}

// shared utility for generators
type printer struct {
	cfg *cfg.GreenConfig
	w   io.Writer
	err error
}

// writes "var {{name}} {{typ}};"
func (p *printer) declare(name string, typ string) {
	p.printf("\nvar %s %s", name, typ)
}

// does:
//
//	if m == nil && size > 0 {
//	    m = make(type, size)
//	} else if len(m) > 0 {
//
//	    for key, _ := range m { delete(m, key) }
//	}
func (p *printer) resizeMap(size string, m *Map) {
	vn := m.Varname()
	if !p.ok() {
		return
	}
	p.printf("\nif %s == nil && %s > 0 {", vn, size)
	p.printf("\n%s = make(%s, %s)", vn, m.TypeName(), size)
	p.printf("\n} else if len(%s) > 0 {", vn)
	p.clearMap(vn)
	p.closeblock()
}

// assign key to value based on varnames
func (p *printer) mapAssign(m *Map) {
	if !p.ok() {
		return
	}
	p.printf("\n%s[%s] = %s", m.Varname(), m.Keyidx, m.Validx)
}

// clear map keys
func (p *printer) clearMap(name string) {
	p.printf("\nfor key, _ := range %[1]s { delete(%[1]s, key) }", name)
}

func (p *printer) resizeSlice(size string, s *Slice) {
	p.printf("\nif cap(%[1]s) >= int(%[2]s) { %[1]s = (%[1]s)[:%[2]s] } else { %[1]s = make(%[3]s, %[2]s) }", s.Varname(), size, s.TypeName())
}

func (p *printer) arrayCheck(want string, got string, additionalGuard string) {
	p.printf("\nif %[3]s %[1]s != %[2]s { err = msgp.ArrayError{Wanted: %[2]s, Got: %[1]s}; return }", got, want, additionalGuard)
}

func (p *printer) closeblock() { p.print("\n}") }

// does:
//
//	for idx := range iter {
//	    {{generate inner}}
//	}
func (p *printer) rangeBlock(idx string, iter string, t traversal, inner Elem) {
	p.printf("\n for %s := range %s {", idx, iter)
	next(t, inner, nil)
	p.closeblock()
}

func (p *printer) nakedReturn() {
	if p.ok() {
		p.print("\nreturn\n}\n")
	}
}

func (p *printer) postLoadHook() {
	if p.ok() {
		p.print("\nif p, ok := interface{}(z).(msgp.PostLoad); ok { p.PostLoadHook() }\n")
	}
}

func (p *printer) preSaveHook() {
	if p.ok() {
		p.print("\nif p, ok := interface{}(z).(msgp.PreSave); ok { p.PreSaveHook() }\n")
	}
}

func (p *printer) dedupReadTop(isUnmarshal bool) {
	if p.cfg.NoDedup {
		return
	}
	if p.ok() {
		// TODO: do we need to figure out how to turn off inlining if a struct has pointer values, so that we
		//  can simply return early if we need to once we've pointed to a dup instead of inflating anew.
	}
}

func (p *printer) dedupReadCleanup(isUnmarshal bool) {
	if p.cfg.NoDedup {
		return
	}
	if p.ok() {
	}
}

func (p *printer) dedupWriteCleanup(isMarshal bool) {
	if p.cfg.NoDedup {
		return
	}
	if p.ok() {
	}
}

func (p *printer) dedupWriteTop(isMarshal bool) {
	if p.cfg.NoDedup {
		return
	}
	if p.ok() {
		if !isMarshal {
			p.print(`	dup, err := en.DedupWriteIsDup(z)
	if dup || err != nil {
		return err
	}
`)
		}
	}
}

func (p *printer) comment(s string) {
	p.print("\n// " + s)
}

func (p *printer) printf(format string, args ...interface{}) {
	if p.err == nil {
		_, p.err = fmt.Fprintf(p.w, format, args...)
	}
}

func (p *printer) print(format string) {
	if p.err == nil {
		_, p.err = io.WriteString(p.w, format)
	}
}

func (p *printer) initPtr(pt *Ptr, isDecode bool) {
	if pt.Needsinit() {
		vname := pt.Varname()
		p.printf("\nif %s == nil { %s = new(%s); }\n", vname, vname, pt.Value.TypeName())
		if isDecode {
			p.printf("dc.DedupIndexEachPtr(%s)\n", vname)
		}
	}
}

func (p *printer) ok() bool { return p.err == nil }

func tobaseConvert(b *BaseElem) string {
	return b.ToBase() + "(" + b.Varname() + ")"
}

func (p *printer) decodeRangeBlock(idx string, parent Elem, t traversal, inner Elem) {
	iter := parent.Varname()
	myzid := parent.GetZid()

	if inner.IsInterface() {
		inner.SetIsInInterfaceSlice()
		target, concreteName := gensym(), gensym()
		cfac := gensym()
		p.printf(`
		// NB: we have a slice of interfaces, so we need to
		//  fill target with the concrete implementation`)

		p.printf("\n for %s := range %s {\n", idx, iter)

		var dedupIndexLine0 string
		if !p.cfg.NoDedup {
			dedupIndexLine0 = fmt.Sprintf(`dc.DedupIndexEachPtr(target_%s)`, target)
			// deduplicate inteface values in the slice, place 1 of 3.
			p.printf(`if kptr, dup := dc.DedupReadIsDup("%s[%s]","%s"); dup {
					%s[%s] = kptr.(%s)
					continue
             }`, iter, idx, inner.TypeName(), iter, idx, inner.TypeName())
		}
		p.printf(`
		concreteName_%s := dc.NextStructName()
        `, concreteName)

		p.printf("target_%s :=  %s[%s]\n", target, iter, idx)
		p.printf(`if concreteName_%s != "" {
				if cfac_%s, cfac_%s_OK := interface{}(z).(msgp.ConcreteFactory); cfac_%s_OK {
					target_%s = cfac_%s.NewValueAsInterface(%v, concreteName_%s).(%s)
				}
                %s
  			    err = target_%s.DecodeMsg(dc)
			    if err != nil {
				    return
			    }
		`, concreteName, cfac, cfac, cfac, target, cfac, myzid, concreteName, inner.TypeName(), dedupIndexLine0, target)
		p.printf(`
                %s[%s] = target_%s
                continue
              }
        `, iter, idx, target) // dedup, place 2 of 3

		if !p.cfg.NoDedup {
			p.printf("\n dc.DedupIndexEachPtr(%s[%s])\n", iter, idx) // dedup, 3 of 3.
		}
		p.printf("\nerr = %s[%s].DecodeMsg(dc) // from decodeRangeBlock in spec.go:511. IsInInterfaceSlice: %v", iter, idx, inner.IsInInterfaceSlice())

		next(t, inner, nil)
	} else {
		p.printf("\n for %s := range %s {", idx, iter)
		next(t, inner, nil)
	}
	p.closeblock()
}

func (p *printer) unmarshalRangeBlock(idx string, parent Elem, t traversal, inner Elem) {
	iter := parent.Varname()
	myzid := parent.GetZid()

	if inner.IsInterface() {
		inner.SetIsInInterfaceSlice()

		target, concreteName := gensym(), gensym()
		cfac := gensym()
		p.printf(`
		// NB: we have a slice of interfaces, so we need to
		//  fill target with the concrete implementation`)

		p.printf("\n for %s := range %s {\n", idx, iter)

		p.printf(`
        var concreteName_%s string
		concreteName_%s, bts = nbs.NextStructName(bts)
        `, concreteName, concreteName)

		p.printf("target_%s :=  %s[%s]\n", target, iter, idx)
		p.printf(`if concreteName_%s != "" {
				if cfac_%s, cfac_%s_OK := interface{}(z).(msgp.ConcreteFactory); cfac_%s_OK {
					target_%s = cfac_%s.NewValueAsInterface(%v, concreteName_%s).(%s)
				}
  			    bts, err = target_%s.UnmarshalMsg(bts)
			    if err != nil {
				    return
			    }
		`, concreteName, cfac, cfac, cfac, target, cfac, myzid, concreteName, inner.TypeName(), target)
		p.printf(`
                %s[%s] = target_%s
                continue
              }
        `, iter, idx, target)

		p.printf("\nbts, err = %s[%s].UnmarshalMsg(bts) // from unmarshalRangeBlock in spec.go:486. IsInInterfaceSlice: %v", iter, idx, inner.IsInInterfaceSlice())

		next(t, inner, nil)
	} else {
		p.printf("\n for %s := range %s {", idx, iter)
		next(t, inner, nil)
	}
	p.closeblock()
}

func hasPointersOrInterfaces(e Elem) bool {
	switch e := e.(type) {
	case *Map:
		// it might... Elem could be checked. TODO, detect for sure instead of guessing:
		return true
	case *Struct:
		for _, f := range e.Fields {
			if f.IsIface || f.IsPointer {
				return true
			}
		}
	}
	return false
}
