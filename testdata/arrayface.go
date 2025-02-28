package testdata

import (
	"fmt"

	"github.com/glycerine/greenpack/msgp"
)

//go:generate greenpack

type Node interface {
	msgp.Serz
}

// sizer code was crashing on node interface being nil.
type Helper struct {
	Children [4]Node
}

func (s *Helper) NewValueAsInterface(zid int64, typename string) interface{} {
	fmt.Printf("Helper.NewValueAsInterface zid=%v, typename='%v'\n", zid, typename)
	return nil
}

// check slice too
type SliceOfIface struct {
	Slice []Node
}

func (s *SliceOfIface) NewValueAsInterface(zid int64, typename string) interface{} {
	fmt.Printf("SliceOfIface) NewValueAsInterface zid=%v, typename='%v'\n", zid, typename)
	return nil
}
