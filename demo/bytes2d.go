package main

import (
	"fmt"
	//"github.com/glycerine/greenpack/msgp"
)

//go:generate greenpack

type SliceOfSlice struct {
	Matrix [][]byte `zid:"0"`
}

func main() {
	sos := &SliceOfSlice{
		Matrix: make([][]byte, 2),
	}
	for i := range sos.Matrix {
		sos.Matrix[i] = make([]byte, 3)
		for j := range sos.Matrix[i] {
			sos.Matrix[i][j] = byte(i*3 + j)
		}
	}
	fmt.Printf("before serialization: sos = '%#v'\n\n", sos)

	by, err := sos.MarshalMsg(nil)
	panicOn(err)

	fmt.Printf("in byte form: '%#v'\n\n", by)

	var sos2 SliceOfSlice

	_, err = sos2.UnmarshalMsg(by)
	panicOn(err)

	fmt.Printf("after deserialization, sos2 = '%#v'\n\n", sos2)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

/* run log
$ go generate
======== Greenpack Code Generator  =======
>>> Input: "bytes2d.go"
>>> Wrote and formatted "bytes2d_gen_test.go"
>>> Wrote and formatted "bytes2d_gen.go"

$ go run .
before serialization: sos = '&main.SliceOfSlice{Matrix:[][]uint8{[]uint8{0x0, 0x1, 0x2}, []uint8{0x3, 0x4, 0x5}}}'

in byte form: '[]byte{0x81, 0xb0, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63, 0x92, 0xc4, 0x3, 0x0, 0x1, 0x2, 0xc4, 0x3, 0x3, 0x4, 0x5}'

after deserialization, sos2 = 'main.SliceOfSlice{Matrix:[][]uint8{[]uint8{0x0, 0x1, 0x2}, []uint8{0x3, 0x4, 0x5}}}'
*/
