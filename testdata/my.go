package testdata

import (
	"time"
)

//go:generate greenpack
//go:generate greenpack -o my_msgp_gen.go -method-prefix=MSGP -io=false -tests=false

const greenSchemaId64 = 0x6eb25cc0f9a3e

//func main() {}

type S struct {
	Placeholder int
}

type S2 struct {
	A struct{}         `zid:"0" msg:"alpha"`
	B string           `msg:"beta" zid:"1"`
	R map[string]uint8 `zid:"2" msg:"ralph"`
	P uint16           `zid:"3" msg:",deprecated"`
	Q uint32           `zid:"4"`
	T int64            `zid:"5" msg:",showzero"`

	// test const array and *ast.SelectorExpr as array size
	Arr [6]float64 `zid:"6" msg:"arr"`

	MyTree *Tree `zid:"7"`
}

type Tree struct {
	Chld []Tree `zid:"0"`
	Str  string `zid:"1"`
	Par  *S2    `zid:"2"`
}

type Big struct {
	Slice     []S2        `zid:"0"  msg:",showzero"`
	Transform map[int]*S2 `zid:"1"  msg:",showzero"`
	Myptr     *S2         `zid:"2"  msg:",showzero"`
	Myarray   [3]string   `zid:"3"  msg:",showzero"`
	MySlice   []string    `zid:"4"  msg:",showzero"`
}

type A struct {
	Name   string    `zid:"0" msg:"name"`
	Bday   time.Time `zid:"1"`
	Phone  string    `zid:"2" msg:"phone,omitempty"`
	Sibs   int       `zid:"3" msg:",omitempty"`
	GPA    float64   `zid:"4"`
	Friend bool      `zid:"5"`
}

type Sys struct {
	F interface{} `zid:"0"`
}
