package parse

import (
	"bytes"
	"fmt"
	cv "github.com/glycerine/goconvey/convey"
	"testing"
)

var _ = fmt.Printf

func Test010IsIfaceScanOfFile(t *testing.T) {

	cv.Convey("scanning a file for interfaces using our heuristic should find them", t, func() {

		chk := func(s string, expect string) {
			b := bytes.NewBuffer([]byte(s))
			m, err := FindInterfaces(b, "testfile")
			//fmt.Printf("\n m = %#v\n", m)
			cv.So(err, cv.ShouldBeNil)
			cv.So(m[expect], cv.ShouldEqual, "testfile")
		}
		chk("type Hello interface {", "Hello")
		chk("    type    Hello    interface    {", "Hello")
		chk("\ttype\tHello\tinterface\t{", "Hello")

	})
}
