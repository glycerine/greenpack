package _generated

import (
	"testing"
)

//go:generate greenpack

type Row struct {
	K []interface{} `msg:"K"`
	V []interface{} `msg:"V"`
	T int64         `msg:"T"`
}

// https://github.com/glycerine/greenpack/issues/1
func TestMarshalMapOfConcreteType(t *testing.T) {

	m := make(map[string][]float64)
	m["def"] = []float64{1.2, 3.4}

	row := Row{}
	row.K = []interface{}{"abc"}
	row.V = []interface{}{m}
	_, err := row.MarshalMsg(nil)
	if err != nil {
		panic(err) // panic: msgp: type "map[string][]float64" not supported
	}
}

func TestMarshalMapOfInterface(t *testing.T) {

	m := make(map[string]interface{})
	m["def"] = []float64{1.2, 3.4}

	row := Row{}
	row.K = []interface{}{"abc"}
	row.V = []interface{}{m}
	_, err := row.MarshalMsg(nil)
	if err != nil {
		panic(err) // panic: msgp: type "map[string][]float64" not supported
	}
}
