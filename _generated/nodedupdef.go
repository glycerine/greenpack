package _generated

import (
	"fmt"
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

//

type Inner2 struct {
	Bubbles int
}

func (in *Inner2) Shout() {}

type Shouter2 interface {
	Shout()
	msgp.Serz
}
type Middle2 struct {
	Children []Shouter2 `msg:",iface"`
}

func (nd *Middle2) NewValueAsInterface(zid int64, typename string) interface{} {
	if typename == "Inner2" {
		return &Inner2{}
	}
	panic(fmt.Sprintf("Middle2.NewValueAsInterface doesn't know how to procude '%s'", typename))
}

func (s *Middle2) Yowza() {}

type Imid2 interface {
	Yowza()
	msgp.Serz
}

type Outer2 struct {
	Slc []Imid2 `msg:"slc,iface" json:"slc" zid:"0"`
}

func (nd *Outer2) NewValueAsInterface(zid int64, typename string) interface{} {
	if typename == "Middle2" {
		return &Middle2{}
	}
	if typename == "Inner2" {
		return &Inner2{}
	}
	panic("unknown typename")
}
