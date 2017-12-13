package _generated

import (
	"bytes"
	"fmt"
	"github.com/glycerine/greenpack/msgp"
	"reflect"
	"testing"
	"time"
)

// benchmark encoding a small, "fast" type.
// the point here is to see how much garbage
// is generated intrinsically by the encoding/
// decoding process as opposed to the nature
// of the struct.
func BenchmarkFastEncode(b *testing.B) {
	v := &TestFast{
		Lat:  40.12398,
		Long: -41.9082,
		Alt:  201.08290,
		Data: []byte("whaaaaargharbl"),
	}
	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	en := msgp.NewWriter(msgp.Nowhere)
	b.SetBytes(int64(buf.Len()))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.EncodeMsg(en)
	}
	en.Flush()
}

// benchmark decoding a small, "fast" type.
// the point here is to see how much garbage
// is generated intrinsically by the encoding/
// decoding process as opposed to the nature
// of the struct.
func BenchmarkFastDecode(b *testing.B) {
	v := &TestFast{
		Lat:  40.12398,
		Long: -41.9082,
		Alt:  201.08290,
		Data: []byte("whaaaaargharbl"),
	}

	var buf bytes.Buffer
	msgp.Encode(&buf, v)
	dc := msgp.NewReader(msgp.NewEndlessReader(buf.Bytes(), b))
	b.SetBytes(int64(buf.Len()))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.DecodeMsg(dc)
	}
}

// This covers the following cases:
//  - Recursive types
//  - Non-builtin identifiers (and recursive types)
//  - time.Time
//  - map[string]string
//  - anonymous structs
//
func Test1EncodeDecode(t *testing.T) {
	f := 32.00
	tt := &TestType{
		F: &f,
		Els: map[string]string{
			"thing_one": "one",
			"thing_two": "two",
		},
		Obj: struct {
			ValueA string `msg:"value_a"`
			ValueB []byte `msg:"value_b"`
		}{
			ValueA: "here's the first inner value",
			ValueB: []byte("here's the second inner value"),
		},
		Child:    nil,
		Time:     time.Now().Round(0),
		Appended: msgp.Raw([]byte{0xc0}), // 'nil'
	}

	var buf bytes.Buffer

	err := msgp.Encode(&buf, tt)
	if err != nil {
		t.Fatal(err)
	}

	tnew := new(TestType)

	err = msgp.Decode(&buf, tnew)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(tt, tnew) {
		t.Logf("in: %#v", tt)
		t.Logf("out: %#v", tnew)
		t.Fatal("objects not equal")
	}

	tanother := new(TestType)

	buf.Reset()
	msgp.Encode(&buf, tt)

	var left []byte
	left, err = tanother.UnmarshalMsg(buf.Bytes())
	if err != nil {
		t.Error(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left", len(left))
	}

	if !reflect.DeepEqual(tt, tanother) {
		t.Logf("in: %#v", tt)
		t.Logf("out: %#v", tanother)
		t.Fatal("objects not equal")
	}
}

func TestIssue168(t *testing.T) {
	buf := bytes.Buffer{}
	test := TestObj{}

	msgp.Encode(&buf, &TestObj{ID1: "1", ID2: "2"})
	msgp.Decode(&buf, &test)

	if test.ID1 != "1" || test.ID2 != "2" {
		t.Fatalf("got back %+v", test)
	}
}

func Test11111HonorDefaultOmitEmpty(t *testing.T) {
	// test that an empty struct is minimally
	// encoding, as if omitempty is applied
	// everywhere possible.
	//

	tt := &SimpleTestType{}

	var buf bytes.Buffer

	err := msgp.Encode(&buf, tt)
	if err != nil {
		t.Fatal(err)
	}

	if len(buf.Bytes()) != 18 {
		t.Fatalf("should have encoding of 18 bytes since omitempty is on by default, had %v", len(buf.Bytes()))
	}

	tnew := new(SimpleTestType)

	err = msgp.Decode(&buf, tnew)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(tt, tnew) {
		t.Logf("in: %#v", tt)
		t.Logf("out: %#v", tnew)
		t.Fatal("objects not equal")
	}
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

// Dedup Test
func Test444DedupOfSamePointerWorks(t *testing.T) {

	ptr := &Target{ID: 3}
	iface := &Greeter{Style: 7}
	nd := &NeedsDedup{
		MyPtr0:   ptr,
		MyPtr1:   ptr,
		MyIface0: iface,
		MyIface1: iface,
		Slice:    []Hello{iface, iface},
	}

	var buf bytes.Buffer
	wr := msgp.NewWriter(&buf)
	//wr.ResetDedup()
	panicOn(nd.EncodeMsg(wr))
	wr.Flush()

	//fmt.Printf("\nAFTER EncodeMsg WRITE: PointerCount=%v. buf='%#v'\n buf as string='%s'\n", wr.PointerCount(), buf.Bytes(), string(buf.Bytes()))

	rd := msgp.NewReader(&buf)
	//rd.ResetDedup()
	var nd2 NeedsDedup
	panicOn(nd2.DecodeMsg(rd))
	if nd2.MyPtr0 != nd2.MyPtr1 {
		panic(fmt.Sprintf("expected pointers to be the same"))
	}
	if nd2.MyIface0.(*Greeter) != nd2.MyIface1.(*Greeter) {
		panic(fmt.Sprintf("expected pointers behind interfaces to be the same"))
	}
	// check slices of interfaces for dedup too.
	if nd2.Slice[0].(*Greeter) != nd2.Slice[1].(*Greeter) {
		panic(fmt.Sprintf("expected pointers behind interfaces to be the same"))
	}
	if nd2.Slice[0].(*Greeter) != nd2.MyIface0.(*Greeter) {
		panic(fmt.Sprintf("expected pointers behind interfaces to be the same"))
	}
}

// Dedup Test2
func Test445DedupOfSamePointerWorks(t *testing.T) {

	// dedup within interface slice alone?

	//ptr := &Target{ID: 3}
	iface := &Greeter{Style: 7}
	nd := &NeedsDedup{
		//MyPtr0: ptr,
		//MyPtr1: ptr,
		//		MyIface0: iface,
		//		MyIface1: iface,
		Slice: []Hello{iface, iface},
	}

	var buf bytes.Buffer
	wr := msgp.NewWriter(&buf)
	//wr.ResetDedup()
	panicOn(nd.EncodeMsg(wr))
	wr.Flush()

	//fmt.Printf("\nAFTER EncodeMsg WRITE: PointerCount=%v. buf='%#v'\n buf as string='%s'\n", wr.PointerCount(), buf.Bytes(), string(buf.Bytes()))

	rd := msgp.NewReader(&buf)
	//rd.ResetDedup()
	var nd2 NeedsDedup
	panicOn(nd2.DecodeMsg(rd))
	// check slices of interfaces for dedup too.
	if nd2.Slice[0].(*Greeter) != nd2.Slice[1].(*Greeter) {
		panic(fmt.Sprintf("expected pointers behind interfaces to be the same"))
	}
}
