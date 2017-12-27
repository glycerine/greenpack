package _generated

import (
	"bytes"
	"fmt"
	"github.com/glycerine/greenpack/msgp"
	"testing"
)

var _ = fmt.Printf

func TestIssue6(t *testing.T) {
	// exercises pointer dedup within a slice of pointers,
	// where there is only one pointer; and condition
	// was that we are *NOT* inlining, so complexity
	// score push up by having at least 4 elements
	// in the struct.

	buf := bytes.NewBuffer(nil)
	err := msgp.Encode(buf, &Commit{Events: []*Event{&Event{}}})
	panicOn(err)

	//	fmt.Printf("we see bytes: '%#v'\n", buf.Bytes())
	//	f, err := os.Create("dump.out")
	//	panicOn(err)
	//	f.Write(buf.Bytes())
	//	f.Close()

	var c Commit
	err = msgp.Decode(buf, &c)
	panicOn(err)
}
