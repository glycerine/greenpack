package testdata

//go:generate greenpack

type HasAny struct {
	// any was borken.
	Inside any `zid:"0"`
}
