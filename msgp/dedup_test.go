package msgp

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDedup(t *testing.T) {

	var buf bytes.Buffer
	wr := NewWriter(&buf)
	origK := 37
	wr.DedupWriteExt(origK)
	wr.Flush()

	//	fmt.Printf("\n buf='%#v'\n", buf)

	rd := NewReader(&buf)

	k, err := rd.DedupReadExt()
	if err != nil {
		panic(err)
	}
	if k != origK {
		panic(fmt.Sprintf("expected %v, got %v", origK, k))
	}
}
