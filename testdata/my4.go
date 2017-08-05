package testdata

//go:generate greenpack -alltuple

type TupleByDefaultTestStruct struct {
	S string `zid:"0"`
	N int64  `zid:"1"`
}
