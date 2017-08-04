package msgp

import (
	"testing"
)

// embedding of version & type info at the
// end of the field name, as in "myField_zid01_i64"
// For back-compatibility with existing
// msgpack2 implementations, this is still
// a legit variable name.
func TestClue(t *testing.T) {

	// convert from (name, type, zid) -> fieldName
	name := "my_Field"
	clue := "i64"
	var zid int64 = 9
	fieldName := Clue2Field(name, clue, zid)
	if fieldName != "my_Field_zid09_i64" {
		t.Fatalf("unexpected fieldName")
	}
	f := Clue2Field(name, clue, 199)
	if f != "my_Field_zid199_i64" {
		t.Fatalf("unexpected fieldName")
	}

	// and the inverse: parse fieldName -> (name, type, zid)
	name2, clue2, zid2, err := Field2Clue(fieldName)
	if err != nil {
		t.Fatal(err)
	}
	if name2 != name {
		t.Fatal("unexpected name2")
	}
	if clue2 != clue {
		t.Fatal("unexpected clue2")
	}
	if zid2 != zid {
		t.Fatal("unexpected zid2")
	}
}
