package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	cv "github.com/glycerine/goconvey/convey"
	"github.com/glycerine/greenpack/cfg"
	"github.com/glycerine/greenpack/gen"
	"testing"
)

func Test001DuplicateOrMissingGapZidFieldsNotAllowed(t *testing.T) {

	cv.Convey("duplicate zid fields and sequences of zid fields with gaps are detected and forbidden", t, func() {

		fmt.Printf("\n  duplicate zid not allowed\n")
		s := "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"0\"`\n" +
			"}\n"
		cv.So(testCode(s, nil), cv.ShouldNotBeNil)

		fmt.Printf("\n  zid must start at 0\n")
		s = "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"1\"`\n" +
			"   Wilma  string `zid:\"2\"`\n" +
			"}\n"
		cv.So(testCode(s, nil), cv.ShouldNotBeNil)

		fmt.Printf("\n  negative zid mean skipped fields\n")
		s = "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"-1\"`\n" +
			"}\n"
		cv.So(testCode(s, nil), cv.ShouldBeNil)

		fmt.Printf("\n  0, 1 should be fine and not error\n")
		s = "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"1\"`\n" +
			"}\n"
		cv.So(testCode(s, nil), cv.ShouldBeNil)

		fmt.Printf("\n  empty struct{} values should be allowed without error\n")
		s = "package p; type Flint struct {}"
		cv.So(testCode(s, nil), cv.ShouldBeNil)

		fmt.Printf("\n  struct{} with zid should be allowed and count in the zid sequence\n")
		s = "package hoo; type S2 struct {A struct{} `zid:\"0\"`;B string   `zid:\"1\"`;}"
		cv.So(testCode(s, nil), cv.ShouldBeNil)

		fmt.Printf("\n  can't have one zid on 2 fields\n")
		s = "package hoo; type S2 struct {A string `zid:\"0\"`; " +
			"B, C string   `zid:\"1\"`;}" // multiple fields, one zid.
		cv.So(testCode(s, nil), cv.ShouldNotBeNil)

		fmt.Printf("\n  can't have one zid on 3 fields\n")
		s = "package hoo; type S2 struct {A string `zid:\"0\"`; " +
			"B, C, D string   `zid:\"1\"`;}" // multiple fields, one zid.
		cv.So(testCode(s, nil), cv.ShouldNotBeNil)

		fmt.Printf("\n  both tag name and Go name available in schema\n")
		s = "package h; type S struct {Alfonzo string `zid:\"0\" msg:\"alpha\"`}"
		var o [500]byte
		cv.So(testCode(s, o[:]), cv.ShouldBeNil)
		so := string(o[:])
		//fmt.Printf("so=%v\n", so)
		cv.So(so, cv.ShouldContainSubstring, `"FieldGoName__str": "Alfonzo"`)
		cv.So(so, cv.ShouldContainSubstring, `"FieldTagName__str": "alpha"`)
	})
}

func Test002EmptyStructsNotMarshalled(t *testing.T) {

	cv.Convey("skipped fields (e.g. struct{} empty values) shouldn't be marshalled", t, func() {

		s := "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"0\"`\n" +
			"   Wilma  string `zid:\"1\"`\n" +
			"   Skiperoo  func()   \n" +
			"   Skiperoo2  struct{}   \n" +
			"}\n"
		cv.So(testCode(s, nil), cv.ShouldBeNil)
	})
}

// testCode is a helper for checking for parsing errors
func testCode(code string, out []byte) error {
	// struct{} fields should still count in their zid
	gofile, err := ioutil.TempFile(".", "tmp-test-001")
	panicOn(err)
	ofile := gofile.Name() + ".out"

	fmt.Fprintf(gofile, code)
	gofile.Close()

	fmt.Printf("\n in file '%s', checking:\n%v\n", gofile.Name(),
		code)

	cfg := cfg.GreenConfig{
		Out:        ofile,
		GoFile:     gofile.Name(),
		Encode:     true,
		Marshal:    true,
		Tests:      true,
		Unexported: false,
	}
	fileSet, err := File(&cfg)
	//fmt.Printf("fileSet.Identities = '%#v'\n", fileSet.Identities)
	//for k, v := range fileSet.Identities {
	//	fmt.Printf("k = '%v'; v= '%#v'\n", k, v)
	//}
	if len(out) > 0 {
		if len(fileSet.Identities) > 0 {
			err := fileSet.SaveMsgpackFile(gofile.Name(), ofile)
			if err != nil {
				panic(err)
			}
		}
		o, err := ioutil.ReadFile(ofile + ".json")
		if err != nil {
			panic(err)
		}
		copy(out, o)
	}
	os.Remove(gofile.Name())
	os.Remove(ofile)
	os.Remove(ofile + ".json")
	return err
}

// write out fields in zid order: re-order the
// field sequence to match the zid-based order.
// This allows for consistent and reproducible
// serialization.

func Test003OrderFieldsByZid(t *testing.T) {

	cv.Convey("for reproducible serialization, order fields by zid", t, func() {
		code := "\npackage fred\n\n" +
			"type Flint struct {\n" +
			"   Barney string `zid:\"1\"`\n" +
			"   Wilma  string `zid:\"0\"`\n" +
			"}\n"

		gofile, err := ioutil.TempFile(".", "tmp-test-001")
		panicOn(err)
		ofile := gofile.Name() + ".out"

		defer func() {
			os.Remove(gofile.Name())
			os.Remove(ofile)
		}()

		fmt.Fprintf(gofile, code)
		gofile.Close()

		fmt.Printf("\n in file '%s', checking:\n%v\n", gofile.Name(),
			code)

		cfg := cfg.GreenConfig{
			Out:        ofile,
			GoFile:     gofile.Name(),
			Encode:     true,
			Marshal:    true,
			Tests:      true,
			Unexported: false,
		}
		fs, err := File(&cfg)
		if len(fs.Identities) > 0 {
			rct := fs.Identities["Flint"].(*gen.Struct)
			//fmt.Printf("\n debug: spec[flint]='%#v'\n", rct)
			//			lastZid := -1
			seenNeg := false
			for i, fld := range rct.Fields {
				curZid := fld.ZebraId
				if seenNeg && curZid != -1 {
					panic(fmt.Sprintf("all zid of -1 must come at the end; we saw %v at pos %v after a -1", curZid, i))
				}
				fmt.Printf("field %v has zid %v\n", i, curZid)
				// allow zid of -1 (not specified) to be at
				// the end only
				if curZid >= 0 {
					// must match i, so be in increasing order!
					if curZid != int64(i) {
						panic(fmt.Sprintf("expected curZid to be %v, but got %v", i, curZid))
					}
				} else {
					seenNeg = true
				}

			}
		}
	})
}

func Test004OrderFieldsByZid(t *testing.T) {

	cv.Convey("zidSetSlice.Sort() order should sort c(-1,1,0,-1) -> c(0,1,-1,-1)", t, func() {

		z := []zid{zid{zid: -1, origPos: 0}, zid{zid: 1, origPos: 1}, zid{zid: 0, origPos: 2}, zid{zid: -1, origPos: 3}}

		sort.Sort(zidSetSlice(z))
		cv.So(z[0].zid, cv.ShouldEqual, 0)
		cv.So(z[1].zid, cv.ShouldEqual, 1)
		cv.So(z[2].zid, cv.ShouldEqual, -1)
		cv.So(z[3].zid, cv.ShouldEqual, -1)
	})
}

func Test005GenericsNotMarshalled(t *testing.T) {

	cv.Convey("uninstantiated generic structs and fields (e.g. struct Gen[T any]{ my T }; ) shouldn't be marshalled", t, func() {

		s := `
package hasgenerics

type Addable interface {
	~complex128 | ~float64
}
type Matrix[T Addable] struct {
	Nrow int
  	Dat        []T
}

type HasMatrix struct {
    M128 Matrix[complex128]
    M64 Matrix[float64]
}
`
		cv.So(testCode(s, nil), cv.ShouldBeNil)
	})
}
