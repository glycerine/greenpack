package parse

import (
	"bytes"
	"fmt"
	cv "github.com/glycerine/goconvey/convey"
	"testing"
)

func Test010IsIfaceScanOfFile(t *testing.T) {

	cv.Convey("scanning a file for interfaces using our heuristic should find them", t, func() {
		s := "type Hello interface {"
		b := bytes.NewBuffer([]byte(s))
		m, err := FindInterfaces(b, "testfile")
		fmt.Printf("\n m = %#v\n", m)
		cv.So(err, cv.ShouldBeNil)
		cv.So(m["Hello"], cv.ShouldEqual, "testfile")
	})
}
