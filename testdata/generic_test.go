package testdata

import (
	"bytes"
	"testing"

	"github.com/glycerine/greenpack/msgp"
)

func Test_with_fields_MarshalUnmarshalMatrix(t *testing.T) {

	// The Nrow and Ncol fields should not be
	// dropped. The generic Dat cannot be transported
	// of course--we don't know its type. We could
	// do reflection at runtime?

	v := Matrix[int]{
		Ncol: 8,
		Nrow: 7,
	}
	bts, err := v.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	var v2 Matrix[int]
	left, err := v2.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if v2.Nrow != v.Nrow {
		t.Fatalf("Nrow not serialized!")
	}
	if v2.Ncol != v.Ncol {
		t.Fatalf("Ncol not serialized!")
	}

	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func Test_with_fields_EncodeDecodeMatrix(t *testing.T) {
	v := Matrix[int]{
		Ncol: 8,
		Nrow: 7,
	}
	var buf bytes.Buffer
	msgp.Encode(&buf, &v)

	m := v.Msgsize()
	if buf.Len() > m {
		t.Logf("WARNING: Msgsize() for %v is inaccurate", v)
	}

	vn := Matrix[int]{}
	err := msgp.Decode(&buf, &vn)
	if err != nil {
		t.Error(err)
	}

	if vn.Nrow != v.Nrow {
		t.Fatalf("Nrow not serialized!")
	}
	if vn.Ncol != v.Ncol {
		t.Fatalf("Ncol not serialized!")
	}

	buf.Reset()
	msgp.Encode(&buf, &v)
	err = msgp.NewReader(&buf).Skip()
	if err != nil {
		t.Error(err)
	}
}
