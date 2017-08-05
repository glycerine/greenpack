package testdata

//go:generate greenpack -tuple-by-default

type TupleByDefaultTestStruct struct {
	S string `zid:"0"`
	N int64  `zid:"1"`
}
