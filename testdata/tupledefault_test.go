package testdata

import (
	"bytes"
	"testing"

	cv "github.com/glycerine/goconvey/convey"
	"github.com/glycerine/greenpack/msgp"
)

func Test012AllTuple(t *testing.T) {

	cv.Convey("greenpack -alltuple always writes tuples without fieldnames", t, func() {

		v := TupleByDefaultTestStruct{
			S: "hello",
			N: -19,
		}
		bts, err := v.MarshalMsg(nil)
		if err != nil {
			t.Fatal(err)
		}

		// verify Encode works the same as Marshal
		var b4 bytes.Buffer
		err = msgp.Encode(&b4, &v)
		if err != nil {
			t.Fatal(err)
		}
		cv.So(string(b4.Bytes()), cv.ShouldResemble, string(bts))

		found0 := bytes.Contains(bts, []byte("S"))
		found1 := bytes.Contains(bts, []byte("N"))
		//fmt.Printf("\n bts = '%x'/'%s'\n", bts, string(bts))

		// tuples don't write field names.
		cv.So(found0, cv.ShouldBeFalse)
		cv.So(found1, cv.ShouldBeFalse)

		// verify unmarshal works
		var v2 TupleByDefaultTestStruct
		_, err = v2.UnmarshalMsg(bts)
		if err != nil {
			t.Fatal(err)
		}
		cv.So(v2.N, cv.ShouldEqual, v.N)
		cv.So(v2.S, cv.ShouldEqual, v.S)

		// verify decode works
		var v3 TupleByDefaultTestStruct
		err = msgp.Decode(bytes.NewBuffer(bts), &v3)
		if err != nil {
			t.Fatal(err)
		}

		cv.So(v3.N, cv.ShouldEqual, v.N)
		cv.So(v3.S, cv.ShouldEqual, v.S)

	})
}
