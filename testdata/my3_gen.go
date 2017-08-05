package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *OmitClueTestStruct) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zgensym_cc028b3bcf246c08_1 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zgensym_cc028b3bcf246c08_1 uint32
	totalEncodedFields0zgensym_cc028b3bcf246c08_1, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zgensym_cc028b3bcf246c08_1 := totalEncodedFields0zgensym_cc028b3bcf246c08_1
	missingFieldsLeft0zgensym_cc028b3bcf246c08_1 := maxFields0zgensym_cc028b3bcf246c08_1 - totalEncodedFields0zgensym_cc028b3bcf246c08_1

	var nextMiss0zgensym_cc028b3bcf246c08_1 int32 = -1
	var found0zgensym_cc028b3bcf246c08_1 [maxFields0zgensym_cc028b3bcf246c08_1]bool
	var curField0zgensym_cc028b3bcf246c08_1 string

doneWithStruct0zgensym_cc028b3bcf246c08_1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgensym_cc028b3bcf246c08_1 > 0 || missingFieldsLeft0zgensym_cc028b3bcf246c08_1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgensym_cc028b3bcf246c08_1, missingFieldsLeft0zgensym_cc028b3bcf246c08_1, msgp.ShowFound(found0zgensym_cc028b3bcf246c08_1[:]), decodeMsgFieldOrder0zgensym_cc028b3bcf246c08_1)
		if encodedFieldsLeft0zgensym_cc028b3bcf246c08_1 > 0 {
			encodedFieldsLeft0zgensym_cc028b3bcf246c08_1--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zgensym_cc028b3bcf246c08_1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgensym_cc028b3bcf246c08_1 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zgensym_cc028b3bcf246c08_1 = 0
			}
			for nextMiss0zgensym_cc028b3bcf246c08_1 < maxFields0zgensym_cc028b3bcf246c08_1 && (found0zgensym_cc028b3bcf246c08_1[nextMiss0zgensym_cc028b3bcf246c08_1] || decodeMsgFieldSkip0zgensym_cc028b3bcf246c08_1[nextMiss0zgensym_cc028b3bcf246c08_1]) {
				nextMiss0zgensym_cc028b3bcf246c08_1++
			}
			if nextMiss0zgensym_cc028b3bcf246c08_1 == maxFields0zgensym_cc028b3bcf246c08_1 {
				// filled all the empty fields!
				break doneWithStruct0zgensym_cc028b3bcf246c08_1
			}
			missingFieldsLeft0zgensym_cc028b3bcf246c08_1--
			curField0zgensym_cc028b3bcf246c08_1 = decodeMsgFieldOrder0zgensym_cc028b3bcf246c08_1[nextMiss0zgensym_cc028b3bcf246c08_1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgensym_cc028b3bcf246c08_1)
		switch curField0zgensym_cc028b3bcf246c08_1 {
		// -- templateDecodeMsg ends here --

		case "S":
			found0zgensym_cc028b3bcf246c08_1[0] = true
			z.S, err = dc.ReadString()
			if err != nil {
				return
			}
		case "N":
			found0zgensym_cc028b3bcf246c08_1[1] = true
			z.N, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zgensym_cc028b3bcf246c08_1 != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of OmitClueTestStruct
var decodeMsgFieldOrder0zgensym_cc028b3bcf246c08_1 = []string{"S", "N"}

var decodeMsgFieldSkip0zgensym_cc028b3bcf246c08_1 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z OmitClueTestStruct) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.S) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.N == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z OmitClueTestStruct) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zgensym_cc028b3bcf246c08_2 [2]bool
	fieldsInUse_zgensym_cc028b3bcf246c08_3 := z.fieldsNotEmpty(empty_zgensym_cc028b3bcf246c08_2[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgensym_cc028b3bcf246c08_3)
	if err != nil {
		return err
	}

	if !empty_zgensym_cc028b3bcf246c08_2[0] {
		// write "S"
		err = en.Append(0xa1, 0x53)
		if err != nil {
			return err
		}
		err = en.WriteString(z.S)
		if err != nil {
			return
		}
	}

	if !empty_zgensym_cc028b3bcf246c08_2[1] {
		// write "N"
		err = en.Append(0xa1, 0x4e)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.N)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z OmitClueTestStruct) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "S"
		o = append(o, 0xa1, 0x53)
		o = msgp.AppendString(o, z.S)
	}

	if !empty[1] {
		// string "N"
		o = append(o, 0xa1, 0x4e)
		o = msgp.AppendInt64(o, z.N)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OmitClueTestStruct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *OmitClueTestStruct) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zgensym_cc028b3bcf246c08_5 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zgensym_cc028b3bcf246c08_5 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zgensym_cc028b3bcf246c08_5, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zgensym_cc028b3bcf246c08_5 := totalEncodedFields4zgensym_cc028b3bcf246c08_5
	missingFieldsLeft4zgensym_cc028b3bcf246c08_5 := maxFields4zgensym_cc028b3bcf246c08_5 - totalEncodedFields4zgensym_cc028b3bcf246c08_5

	var nextMiss4zgensym_cc028b3bcf246c08_5 int32 = -1
	var found4zgensym_cc028b3bcf246c08_5 [maxFields4zgensym_cc028b3bcf246c08_5]bool
	var curField4zgensym_cc028b3bcf246c08_5 string

doneWithStruct4zgensym_cc028b3bcf246c08_5:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zgensym_cc028b3bcf246c08_5 > 0 || missingFieldsLeft4zgensym_cc028b3bcf246c08_5 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zgensym_cc028b3bcf246c08_5, missingFieldsLeft4zgensym_cc028b3bcf246c08_5, msgp.ShowFound(found4zgensym_cc028b3bcf246c08_5[:]), unmarshalMsgFieldOrder4zgensym_cc028b3bcf246c08_5)
		if encodedFieldsLeft4zgensym_cc028b3bcf246c08_5 > 0 {
			encodedFieldsLeft4zgensym_cc028b3bcf246c08_5--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zgensym_cc028b3bcf246c08_5 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zgensym_cc028b3bcf246c08_5 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zgensym_cc028b3bcf246c08_5 = 0
			}
			for nextMiss4zgensym_cc028b3bcf246c08_5 < maxFields4zgensym_cc028b3bcf246c08_5 && (found4zgensym_cc028b3bcf246c08_5[nextMiss4zgensym_cc028b3bcf246c08_5] || unmarshalMsgFieldSkip4zgensym_cc028b3bcf246c08_5[nextMiss4zgensym_cc028b3bcf246c08_5]) {
				nextMiss4zgensym_cc028b3bcf246c08_5++
			}
			if nextMiss4zgensym_cc028b3bcf246c08_5 == maxFields4zgensym_cc028b3bcf246c08_5 {
				// filled all the empty fields!
				break doneWithStruct4zgensym_cc028b3bcf246c08_5
			}
			missingFieldsLeft4zgensym_cc028b3bcf246c08_5--
			curField4zgensym_cc028b3bcf246c08_5 = unmarshalMsgFieldOrder4zgensym_cc028b3bcf246c08_5[nextMiss4zgensym_cc028b3bcf246c08_5]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zgensym_cc028b3bcf246c08_5)
		switch curField4zgensym_cc028b3bcf246c08_5 {
		// -- templateUnmarshalMsg ends here --

		case "S":
			found4zgensym_cc028b3bcf246c08_5[0] = true
			z.S, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "N":
			found4zgensym_cc028b3bcf246c08_5[1] = true
			z.N, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss4zgensym_cc028b3bcf246c08_5 != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of OmitClueTestStruct
var unmarshalMsgFieldOrder4zgensym_cc028b3bcf246c08_5 = []string{"S", "N"}

var unmarshalMsgFieldSkip4zgensym_cc028b3bcf246c08_5 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z OmitClueTestStruct) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.S) + 12 + msgp.Int64Size
	return
}
