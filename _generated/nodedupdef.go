package _generated

import (
	"github.com/glycerine/greenpack/msgp"
)

//go:generate greenpack -no-dedup -o nodedup_generated.go

type Target2 struct {
	ID int
}

type Greeter2 struct {
	Style int
}

func (t *Greeter2) Hi() {}

type Hello2 interface {
	Hi()
	msgp.Serz
}

type NoDedup struct {
	MyPtr0 *Target2
	MyPtr1 *Target2

	MyIface0 Hello2 `msg:",iface"`
	MyIface1 Hello2 `msg:",iface"`

	Slice    []Hello2 `msg:",iface"`
	SlicePtr []*Greeter2
}

func (nd *NoDedup) NewValueAsInterface(zid int64, typename string) interface{} {
	//fmt.Printf("\n DEBUG, in NoDedup.NewValueAsInterface(): typename = '%s'\n", typename)
	if typename == "Greeter2" {
		return &Greeter2{}
	}
	panic("unknown typename")
}
