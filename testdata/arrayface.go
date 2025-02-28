package testdata

//go:generate greenpack

type Node any

// sizer code was crashing on node interface being nil.
type Helper struct {
	Children [4]Node
}

// check slice too
type SliceOfIface struct {
	Slice []Node
}
