package _generated

import (
	"fmt"
	"os"
	"time"

	"github.com/glycerine/greenpack/msgp"
)

//go:generate greenpack -o generated.go -no-dedup=false

// All of the struct
// definitions in this
// file are fed to the code
// generator when `make test` is
// called, followed by an
// invocation of `go test -v` in this
// directory. A simple way of testing
// a struct definition is
// by adding it to this file.

type Block [32]byte

// tests edge-cases with
// compiling size compilation.
type X struct {
	Values    [32]byte    // should compile to 32*msgp.ByteSize; encoded as Bin
	More      Block       // should be identical to the above
	Others    [][32]int32 // should compile to len(x.Others)*32*msgp.Int32Size
	Matrix    [][]int32   // should not optimize
	ManyFixed []Fixed
}

// test fixed-size struct
// size compilation
type Fixed struct {
	A float64
	B bool
}

type PreviouslyAnon struct {
	ValueA string `msg:"value_a"`
	ValueB []byte `msg:"value_b"`
}

type TestType struct {
	F        *float64          `msg:"float"`
	Els      map[string]string `msg:"elements"`
	Obj      PreviouslyAnon    `msg:"object"`
	Child    *TestType         `msg:"child"`
	Time     time.Time         `msg:"time"`
	Any      interface{}       `msg:"any"`
	Appended msgp.Raw          `msg:"appended"`
	Num      msgp.Number       `msg:"num"`
	Slice1   []string
	Slice2   []string
	SlicePtr *[]string
	Dur      time.Duration
}

type SimpleTestType struct {
	F   *float64          `msg:"float"`
	Els map[string]string `msg:"elements"`
	Obj *PreviouslyAnon   `msg:"object"`
}

//msgp:tuple Object
type Object struct {
	ObjectNo string   `msg:"objno"`
	Slice1   []string `msg:"slice1"`
	Slice2   []string `msg:"slice2"`
	MapMap   map[string]map[string]string
}

//msgp:tuple TestBench

type TestBench struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

//msgp:tuple TestFast

type TestFast struct {
	Lat, Long, Alt float64 // test inline decl
	Data           []byte
}

// Test nested aliases
type FastAlias TestFast
type AliasContainer struct {
	Fast FastAlias
}

// Test dependency resolution
type IntA int
type IntB IntA
type IntC IntB

type TestHidden struct {
	A   string
	B   []float64
	Bad func(string) bool // This results in a warning: field "Bad" unsupported
}

type Embedded struct {
	*Embedded   // test embedded field
	Children    []Embedded
	PtrChildren []*Embedded
	Other       string
}

const eight = 8

type Things struct {
	Cmplx complex64 `msg:"complex"` // test slices
	Vals  []int32   `msg:"values"`
	//Arr   [msgp.ExtensionPrefixSize]float64 `msg:"arr"`            // test const array and *ast.SelectorExpr as array size
	Arr  [6]float64         `msg:"arr"`            // test const array and *ast.SelectorExpr as array size
	Arr2 [4]float64         `msg:"arr2"`           // test basic lit array
	Ext  *msgp.RawExtension `msg:"ext,extension"`  // test extension
	Oext msgp.RawExtension  `msg:"oext,extension"` // test extension reference
}

//msgp:shim SpecialID as:[]byte using:toBytes/fromBytes

type SpecialID string
type TestObj struct{ ID1, ID2 SpecialID }

func toBytes(id SpecialID) []byte   { return []byte(string(id)) }
func fromBytes(id []byte) SpecialID { return SpecialID(string(id)) }

type MyEnum byte

const (
	A MyEnum = iota
	B
	C
	D
	invalid
)

// test shim directive (below)

//msgp:shim MyEnum as:string using:(MyEnum).String/myenumStr

//msgp:shim *os.File as:string using:filetostr/filefromstr

func filetostr(f *os.File) string {
	return f.Name()
}

func filefromstr(s string) *os.File {
	f, _ := os.Open(s)
	return f
}

func (m MyEnum) String() string {
	switch m {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	default:
		return "<invalid>"
	}
}

func myenumStr(s string) MyEnum {
	switch s {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	default:
		return invalid
	}
}

// test pass-specific directive
// :encode ignore Insane

// anonymous structs can't be serialized with greenpack,
// because the fieldsempty/everything omitempty by default
// logic requires named types when we define the helper
// method fieldsNotEmpty().
type Insane [3]map[string]InsaneInner
type InsaneInner struct{ A, B CustomInt }

type Custom struct {
	Bts   CustomBytes          `msg:"bts"`
	Mp    map[string]*Embedded `msg:"mp"`
	Enums []MyEnum             `msg:"enums"` // test explicit enum shim
	Some  FileHandle           `msg:file_handle`
}

type Files []*os.File

type FileHandle struct {
	Relevent Files  `msg:"files"`
	Name     string `msg:"name"`
}

type CustomInt int
type CustomBytes []byte

// Test omitempty tag
type TestOmitEmpty struct {

	// scalars
	Name     string    `msg:",omitempty"`
	BirthDay time.Time `msg:",omitempty"`
	Phone    string    `msg:",omitempty"`
	Siblings int       `msg:",omitempty"`
	Spouse   bool      `msg:",omitempty"`
	Money    float64   `msg:",omitempty"`

	// slices
	SliceName     []string    `msg:",omitempty"`
	SliceBirthDay []time.Time `msg:",omitempty"`
	SlicePhone    []string    `msg:",omitempty"`
	SliceSiblings []int       `msg:",omitempty"`
	SliceSpouse   []bool      `msg:",omitempty"`
	SliceMoney    []float64   `msg:",omitempty"`

	// arrays
	ArrayName     [3]string    `msg:",omitempty"`
	ArrayBirthDay [3]time.Time `msg:",omitempty"`
	ArrayPhone    [3]string    `msg:",omitempty"`
	ArraySiblings [3]int       `msg:",omitempty"`
	ArraySpouse   [3]bool      `msg:",omitempty"`
	ArrayMoney    [3]float64   `msg:",omitempty"`

	// maps
	MapStringString map[string]string      `msg:",omitempty"`
	MapStringIface  map[string]interface{} `msg:",omitempty"`

	// pointers
	PtrName     *string    `msg:",omitempty"`
	PtrBirthDay *time.Time `msg:",omitempty"`
	PtrPhone    *string    `msg:",omitempty"`
	PtrSiblings *int       `msg:",omitempty"`
	PtrSpouse   *bool      `msg:",omitempty"`
	PtrMoney    *float64   `msg:",omitempty"`

	Inside1    OmitEmptyInside1 `msg:",omitempty"`
	Greetings  string           `msg:",omitempty"`
	Bullwinkle *Rocky           `msg:",omitempty"`
}

type TopNester struct {
	TopId      int
	Greetings  string `msg:",omitempty"`
	Bullwinkle *Rocky `msg:",omitempty"`

	MyIntArray  [3]int               `msg:",omitempty"`
	MyByteArray [3]byte              `msg:",omitempty"`
	MyMap       map[string]string    `msg:",omitempty"`
	MyArrayMap  [3]map[string]string `msg:",omitempty"`

	// the time.Time extension
	TopTime time.Time  `msg:",omitempty"`
	PtrTime *time.Time `msg:",omitempty"`
}

type Rocky struct {
	Bugs  *Bunny `msg:",omitempty"`
	Road  string `msg:",omitempty"`
	Moose *Moose `msg:",omitempty"`
}

type Bunny struct {
	Carrots []int             `msg:",omitempty"`
	Sayings map[string]string `msg:",omitempty"`
	BunnyId int               `msg:",omitempty"`
}

type Moose struct {
	Trees   []int             `msg:",omitempty"`
	Sayings map[string]string `msg:",omitempty"`
	Id      int
}

type OmitEmptyInside1 struct {
	CountOfMonteCrisco int
	Name               string           `msg:"name,omitempty"`
	Inside2            OmitEmptyInside2 `msg:",omitempty"`
}

type OmitEmptyInside2 struct {
	NameSuey string `msg:",omitempty"`
}

type OmitSimple struct {
	CountDrocula int
	Inside1      OmitEmptyInside1 `msg:",omitempty"`
}

type Wrapper struct {
	Tree *Tree
}

type Tree struct {
	Children []Tree
	Element  int
	Parent   *Wrapper
}

type Target struct {
	ID int
}

type Greeter struct {
	Style int
}

func (t *Greeter) Hi() {}

type Hello interface {
	Hi()
	msgp.Serz
}

type NeedsDedup struct {
	MyPtr0 *Target
	MyPtr1 *Target

	MyIface0 Hello `msg:",iface"`
	MyIface1 Hello `msg:",iface"`

	Slice    []Hello `msg:",iface"`
	SlicePtr []*Greeter
}

func (nd *NeedsDedup) NewValueAsInterface(zid int64, typename string) interface{} {
	//fmt.Printf("\n DEBUG, in NeedsDedup.NewValueAsInterface(): typename = '%s'\n", typename)
	if typename == "Greeter" {
		return &Greeter{}
	}
	panic("unknown typename")
}

// for test 500

type Inner struct {
	Bubbles int
}

func (in *Inner) Shout() {}

type Shouter interface {
	Shout()
	msgp.Serz
}
type Middle struct {
	Children []Shouter `msg:",iface"`
}

func (nd *Middle) NewValueAsInterface(zid int64, typename string) interface{} {
	if typename == "Inner" {
		return &Inner{}
	}
	panic(fmt.Sprintf("Middle.NewValueAsInterface doesn't know how to procude '%s'", typename))
}

func (s *Middle) Yowza() {}

type Imid interface {
	Yowza()
	msgp.Serz
}

type Outer struct {
	Slc []Imid `msg:"slc,iface" json:"slc" zid:"0"`
}

func (nd *Outer) NewValueAsInterface(zid int64, typename string) interface{} {
	if typename == "Middle" {
		return &Middle{}
	}
	if typename == "Inner" {
		return &Inner{}
	}
	panic("unknown typename")
}

// struct-valued fields having problems?
type ZError struct {
	What int `zid:"0"`
}

type ZResponse struct {
	ProblemFieldIfStructValue ZError `zid:"0"`
}
