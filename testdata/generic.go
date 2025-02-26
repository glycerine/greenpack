package testdata

import (
	"fmt"
)

//go:generate greenpack

// should be ignored, but parsed smoothly without complaint.
type Addable interface {
	~complex128 | ~complex64 | ~float64 | ~float32 |
		~byte | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 | ~int
}

// Dat field should be ignored
// (because has generic type parameter);
// the rest of the struct should be serialized.
type Matrix[T Addable] struct {
	Dat  []T
	Nrow int
	Ncol int
}

func useGenericMatrix() {
	m := &Matrix[float64]{}
	fmt.Printf("m = '%#v'", m)
}

type StructThatHasFieldsWithGenerics struct {
	M128 Matrix[complex128]
	M64  Matrix[float64]
}

type CompareToIgnoredFields struct {
	IgnoreMe int `msg:"-"`
}
