package parse

import (
	"fmt"
	"sort"

	"github.com/glycerine/greenpack/gen"
)

// This file defines when and how we
// propagate type information from
// one type declaration to another.
// After the processing pass, every
// non-primitive type is marshalled/unmarshalled/etc.
// through a function call. Here, we propagate
// the type information into the caller's type
// tree *if* the child type is simple enough.
//
// For example, types like
//
//    type A [4]int
//
// will get pushed into parent methods,
// whereas types like
//
//    type B [3]map[string]struct{A, B [4]string}
//
// will not.

// this is an approximate measure
// of the number of children in a node
const maxComplex = 5

// begin recursive search for identities with the
// given name and replace them with be
func (f *FileSet) findShim(id string, be *gen.BaseElem) {
	for name, el := range f.Identities {
		pushstate(name)
		switch el := el.(type) {
		case *gen.Struct:
			for i := range el.Fields {
				if !el.Fields[i].Skip {
					f.nextShim(&el.Fields[i].FieldElem, id, be)
				}
			}
		case *gen.Array:
			f.nextShim(&el.Els, id, be)
		case *gen.Slice:
			f.nextShim(&el.Els, id, be)
		case *gen.Map:
			f.nextShim(&el.Value, id, be)
		case *gen.Ptr:
			f.nextShim(&el.Value, id, be)
		}
		popstate()
	}
	// we'll need this at the top level as well
	f.Identities[id] = be
}

func (f *FileSet) nextShim(ref *gen.Elem, id string, be *gen.BaseElem) {
	if (*ref).TypeName() == id {
		vn := (*ref).Varname()
		*ref = be.Copy()
		(*ref).SetVarname(vn)
	} else {
		switch el := (*ref).(type) {
		case *gen.Struct:
			for i := range el.Fields {
				if !el.Fields[i].Skip {
					f.nextShim(&el.Fields[i].FieldElem, id, be)
				}
			}
		case *gen.Array:
			f.nextShim(&el.Els, id, be)
		case *gen.Slice:
			f.nextShim(&el.Els, id, be)
		case *gen.Map:
			f.nextShim(&el.Value, id, be)
		case *gen.Ptr:
			f.nextShim(&el.Value, id, be)
		}
	}
}

type ComplexityThenNameSorter struct {
	Name    string
	Plexity int
}
type PlexSorterSlice []*ComplexityThenNameSorter

func (p PlexSorterSlice) Len() int { return len(p) }
func (p PlexSorterSlice) Less(i, j int) bool {
	switch {
	case p[i].Plexity > p[j].Plexity:
		return true
	case p[i].Plexity < p[j].Plexity:
		return false
	}
	return p[i].Name < p[j].Name
}
func (p PlexSorterSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// propInline identifies and inlines candidates
func (f *FileSet) propInline() {

	// We sort struct names to get deterministic output.
	// Otherwise the order of one struct being
	// inlined into another can change, which increases
	// the target's complexity and changes the subsequent inlining choice.
	// We really like deterministic output to avoid
	// arbitrary changes impacting that show up in
	// version control history when not expected.
	//
	// But we also want to promote inlining for speed.
	// So we also sort by un-inlined complexity, biggest first.
	// Break ties by name.
	//
	var plexNameSort PlexSorterSlice
	for k, node := range f.Identities {
		plexNameSort = append(plexNameSort,
			&ComplexityThenNameSorter{Name: k, Plexity: node.Complexity()})
	}
	sort.Sort(plexNameSort)

	for _, pns := range plexNameSort {
		name := pns.Name
		el := f.Identities[name]
		pushstate(name)
		switch el := el.(type) {
		case *gen.Struct:
			for i := range el.Fields {
				if !el.Fields[i].Skip {
					f.nextInline(&el.Fields[i].FieldElem, name)
				}
			}
		case *gen.Array:
			f.nextInline(&el.Els, name)
		case *gen.Slice:
			f.nextInline(&el.Els, name)
		case *gen.Map:
			f.nextInline(&el.Value, name)
		case *gen.Ptr:
			f.nextInline(&el.Value, name)
		}
		popstate()
	}
}

const fatalloop = `detected infinite recursion in inlining loop!
Please file a bug at github.com/glycerine/greenpack/issues!
Thanks!
`

func (f *FileSet) nextInline(ref *gen.Elem, root string) {
	switch el := (*ref).(type) {
	case *gen.BaseElem:
		// ensure that we're not inlining
		// a type into itself
		typ := el.TypeName()
		if el.Value == gen.IDENT && typ != root {
			if node, ok := f.Identities[typ]; ok {
				plexity := node.Complexity()
				if plexity < maxComplex {
					infof("inlining %s\n", typ)
					// This should never happen; it will cause
					// infinite recursion.
					if node == *ref {
						panic(fatalloop)
					}

					*ref = node.Copy()
					f.nextInline(ref, node.TypeName())
				}
			} else if !ok && !el.Resolved() {
				// this is the point at which we're sure that
				// we've got a type that isn't a primitive,
				// a library builtin, or a processed type
				warnf("unresolved identifier: %s\n", typ)
			}
		}
	case *gen.Struct:
		for i := range el.Fields {
			if !el.Fields[i].Skip {
				f.nextInline(&el.Fields[i].FieldElem, root)
			}
		}
	case *gen.Array:
		f.nextInline(&el.Els, root)
	case *gen.Slice:
		f.nextInline(&el.Els, root)
	case *gen.Map:
		f.nextInline(&el.Value, root)
	case *gen.Ptr:
		f.nextInline(&el.Value, root)
	default:
		panic(fmt.Sprintf("bad elem type %T/val=%#v", el, el))
	}
}
