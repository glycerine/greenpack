package testdata

//go:generate greenpack  -omit-clue

type OmitClueTestStruct struct {
	S string `zid:"0"`
	N int64  `zid:"1"`
}
