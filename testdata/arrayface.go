package testdata

//go:generate greenpack

type node any

// sizer code was crashing on node interface being nil.
type Helper struct {
	Children [4]node
}
