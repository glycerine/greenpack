package gen

// ugorji_test: tests of reflection based
// msgpack serz, for generics.

import (
	"bytes"
	//"fmt"
	"math"
	"reflect"
	"testing"

	codec "github.com/glycerine/go-1/codec"
	//goon "github.com/shurcooL/go-goon"
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

	var h codec.Handle = &codec.MsgpackHandle{}
	by := make([]byte, 0, 256<<10)
	buf := bytes.NewBuffer(by)

	// will auto-expand the backing slice if too small,
	// so always read buf.Bytes() not by.
	enc := codec.NewEncoder(buf, h)

	v := mkTestStruct()
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
	vv("len by = %v", len(buf.Bytes()))

	v2 := &Struct2{}

	dec := codec.NewDecoder(buf, h)
	err = dec.Decode(v2)
	panicOn(err)

	if !reflect.DeepEqual(v, v2) {
		t.Fatalf("expected '%#v', got '%#v'", v, v2) // green
	}
}
