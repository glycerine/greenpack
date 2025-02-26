package gen

/*
// flect_test: tests of reflection based
// msgpack serz, for generics.

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"testing"

	shamaton2 "github.com/shamaton/msgpack/v2"
	goon "github.com/shurcooL/go-goon"
)

type Struct1 struct {
	String string
	By     []byte

	I   int
	I64 int64
	I32 int32
	I16 int16
	I8  int8

	U   uint
	U64 uint64
	U32 uint32
	U16 uint16
	U8  uint8

	Imin   int
	I64min int64
	I32min int32
	I16min int16
	I8min  int8
}

type Struct2 struct {
	S1 *Struct1
	S2 *Struct2 // should not blow up on recursion

	By2 []byte
}

func mkTestStruct() *Struct2 {
	s1 := &Struct1{
		String: "msgpack",
		I:      math.MaxInt,
		Imin:   math.MinInt,
		I8:     math.MaxInt8,
		I8min:  math.MinInt8,
		I16:    math.MaxInt16,
		I16min: math.MinInt16,
		I32:    math.MaxInt32,
		I32min: math.MinInt32,
		I64:    math.MaxInt64,
		I64min: math.MinInt64,
		U:      math.MaxUint,
		U8:     math.MaxUint8,
		U16:    math.MaxUint16,
		U32:    math.MaxUint32,
		U64:    math.MaxUint64,
		By:     []byte("bytes!"),
	}
	s2 := &Struct2{
		S1: s1,
		S2: &Struct2{
			By2: []byte("DEEPer By2"),
		},
		By2: []byte("I'm a By2"),
	}
	return s2
}

func TestReflectionSerz(t *testing.T) {

	v := mkTestStruct()
	by, err := shamaton2.Marshal(v)
	if err != nil {
		panic(err)
	}
	v2 := &Struct2{}
	if err = shamaton2.Unmarshal(by, v2); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(v, v2) {
		t.Fatalf("expected '%#v', got '%#v'", v, v2) // green
	}

	// streaming
	buf := bytes.NewBuffer(nil)

	if err := shamaton2.MarshalWrite(buf, v2); err != nil {
		panic(err)
	}
	v3 := &Struct2{}
	if err := shamaton2.UnmarshalRead(buf, v3); err != nil {
		panic(err)
	}
	fmt.Printf("v2.S1.By = '%v'\n", string(v2.S1.By))
	// v2.S1.By = 'bytes!'

	fmt.Printf("versus v3.S1.By = '%v'\n", string(v3.S1.By))
	// v3.S1.By = 'I'm a '

	if !reflect.DeepEqual(v2, v3) {
		goon.Dump(v2)
		fmt.Printf("versus got back v3:\n")
		goon.Dump(v3)
		t.Fatalf("expected '%#v', got '%#v'", v2, v3) //red
	}
}
*/
/*
$ diff -b v v3
5,10c5,10
(*gen.Struct2)(&gen.Struct2{
	S1: (*gen.Struct1)(&gen.Struct1{
		String: (string)("msgpack"),
		By: ([]uint8)([]uint8{
< 			(uint8)(98),
< 			(uint8)(121),
< 			(uint8)(116),
< 			(uint8)(101),
< 			(uint8)(115),
< 			(uint8)(33),
---
(*gen.Struct2)(&gen.Struct2{
	S1: (*gen.Struct1)(&gen.Struct1{
		String: (string)("msgpack"),
		By: ([]uint8)([]uint8{
> 			(uint8)(73),
> 			(uint8)(39),
> 			(uint8)(109),
> 			(uint8)(32),
> 			(uint8)(97),
> 			(uint8)(32),
32,37c32,34
< 			(uint8)(68),
< 			(uint8)(69),
< 			(uint8)(69),
< 			(uint8)(80),
< 			(uint8)(101),
< 			(uint8)(114),
---
> 			(uint8)(73),
> 			(uint8)(39),
> 			(uint8)(109),
38a36,37
> 			(uint8)(97),
> 			(uint8)(32),
40a40
> 			(uint8)(50),
*/
