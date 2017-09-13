package _generated

import (
	"fmt"
	"time"
)

//go:generate greenpack -fast-strings -o issue4_gen.go

type RowIssue4 struct {
	K []interface{} `msg:"K"`
	V []interface{} `msg:"V"`
	T int64         `msg:"T"`
}

// EncodeRow encode one row of data to a blob
func encodeRow(row RowIssue4) ([]byte, error) {
	return row.MarshalMsg(nil)
}

func newRowIssue4(timestamp int64, objects ...interface{}) *RowIssue4 {
	r := &RowIssue4{
		T: timestamp,
	}
	if len(objects) > 0 {
		r.AppendKey(objects[0]).AppendValue(objects[1:]...)
	}
	return r
}

func (row *RowIssue4) AppendKey(objects ...interface{}) *RowIssue4 {
	row.K = append(row.K, objects...)
	return row
}

func (row *RowIssue4) AppendValue(objects ...interface{}) *RowIssue4 {
	row.V = append(row.V, objects...)
	return row
}

func Now() (ts int64) {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// decodeRowIssue4 decodes one row of data from a blob
func decodeRowIssue4(encodedBytes []byte) (*RowIssue4, error) {
	row := &RowIssue4{}
	_, err := row.UnmarshalMsg(encodedBytes)
	if err != nil {
		err = fmt.Errorf("decode row error %v: %s\n", err, string(encodedBytes))
	}
	return row, err
}

func issue4EncodeDecode() {

	originalKey := uint32(2349234)

	originalData := []interface{}{
		originalKey,
		"jkjkj",
		123,
		"zx,mcv",
	}

	originalTs := Now()

	encodedRow, _ := encodeRow(*newRowIssue4(originalTs, originalData...))

	//fmt.Printf("\nencoded: %#v\n", encodedRow)

	row, _ := decodeRowIssue4(encodedRow)

	if originalTs != row.T {
		panic(fmt.Sprintf("Failed to decode ts %d => %d", originalTs, row.T))
	}

	// check that we got an uint32 back
	if key, ok := row.K[0].(uint32); ok {
		if key != originalKey {
			panic(fmt.Sprintf("Failed to decode key %T/val=%v => %T/val=%v", originalKey, originalKey, key, key))
		}
	} else {
		// row.K[0] is uint64
		panic(fmt.Sprintf("Failed to get matching key type %T/val=%v => %T/val=%v", originalKey, originalKey, row.K[0], row.K[0]))
	}

	if len(row.V) != len(originalData)-1 {
		panic(fmt.Sprintf("Failed to decode to the same length %d => %d", len(originalData), len(row.V)-1))
	}

}
