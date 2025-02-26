package testdata

//go:generate greenpack

type HasEmptyInterface struct {
	EmptyIface interface{} `zid:"0"`
}
