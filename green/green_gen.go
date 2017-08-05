package green

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Field) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields0zrev = 11

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zrev uint32
	totalEncodedFields0zrev, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zrev := totalEncodedFields0zrev
	missingFieldsLeft0zrev := maxFields0zrev - totalEncodedFields0zrev

	var nextMiss0zrev int32 = -1
	var found0zrev [maxFields0zrev]bool
	var curField0zrev string

doneWithStruct0zrev:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zrev > 0 || missingFieldsLeft0zrev > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zrev, missingFieldsLeft0zrev, msgp.ShowFound(found0zrev[:]), decodeMsgFieldOrder0zrev)
		if encodedFieldsLeft0zrev > 0 {
			encodedFieldsLeft0zrev--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zrev = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zrev < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zrev = 0
			}
			for nextMiss0zrev < maxFields0zrev && (found0zrev[nextMiss0zrev] || decodeMsgFieldSkip0zrev[nextMiss0zrev]) {
				nextMiss0zrev++
			}
			if nextMiss0zrev == maxFields0zrev {
				// filled all the empty fields!
				break doneWithStruct0zrev
			}
			missingFieldsLeft0zrev--
			curField0zrev = decodeMsgFieldOrder0zrev[nextMiss0zrev]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zrev)
		switch curField0zrev {
		// -- templateDecodeMsg ends here --

		case "Zid__i64":
			found0zrev[0] = true
			z.Zid, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "FieldGoName__str":
			found0zrev[1] = true
			z.FieldGoName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTagName__str":
			found0zrev[2] = true
			z.FieldTagName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldTypeStr__str":
			found0zrev[3] = true
			z.FieldTypeStr, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FieldCategory__":
			found0zrev[4] = true
			{
				var zcbl uint64
				zcbl, err = dc.ReadUint64()
				z.FieldCategory = Zkind(zcbl)
			}
			if err != nil {
				return
			}
		case "FieldPrimitive__":
			found0zrev[5] = true
			{
				var zmfa uint64
				zmfa, err = dc.ReadUint64()
				z.FieldPrimitive = Zkind(zmfa)
			}
			if err != nil {
				return
			}
		case "FieldFullType__ptr":
			found0zrev[6] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.FieldFullType != nil {
					dc.PushAlwaysNil()
					err = z.FieldFullType.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.FieldFullType == nil {
					z.FieldFullType = new(Ztype)
				}
				err = z.FieldFullType.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "OmitEmpty__boo":
			found0zrev[7] = true
			z.OmitEmpty, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Skip__boo":
			found0zrev[8] = true
			z.Skip, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Deprecated__boo":
			found0zrev[9] = true
			z.Deprecated, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "ShowZero__boo":
			found0zrev[10] = true
			z.ShowZero, err = dc.ReadBool()
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
	if nextMiss0zrev != -1 {
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

// fields of Field
var decodeMsgFieldOrder0zrev = []string{"Zid__i64", "FieldGoName__str", "FieldTagName__str", "FieldTypeStr__str", "FieldCategory__", "FieldPrimitive__", "FieldFullType__ptr", "OmitEmpty__boo", "Skip__boo", "Deprecated__boo", "ShowZero__boo"}

var decodeMsgFieldSkip0zrev = []bool{false, false, false, false, false, false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Field) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 11
	}
	var fieldsInUse uint32 = 11
	isempty[2] = (len(z.FieldTagName) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.FieldTypeStr) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.FieldCategory == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.FieldPrimitive == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}
	isempty[6] = (z.FieldFullType == nil) // pointer, omitempty
	if isempty[6] {
		fieldsInUse--
	}
	isempty[7] = (!z.OmitEmpty) // bool, omitempty
	if isempty[7] {
		fieldsInUse--
	}
	isempty[8] = (!z.Skip) // bool, omitempty
	if isempty[8] {
		fieldsInUse--
	}
	isempty[9] = (!z.Deprecated) // bool, omitempty
	if isempty[9] {
		fieldsInUse--
	}
	isempty[10] = (!z.ShowZero) // bool, omitempty
	if isempty[10] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Field) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zkyn [11]bool
	fieldsInUse_zgyq := z.fieldsNotEmpty(empty_zkyn[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgyq)
	if err != nil {
		return err
	}

	// write "Zid__i64"
	err = en.Append(0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Zid)
	if err != nil {
		return
	}
	// write "FieldGoName__str"
	err = en.Append(0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.FieldGoName)
	if err != nil {
		return
	}
	if !empty_zkyn[2] {
		// write "FieldTagName__str"
		err = en.Append(0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.FieldTagName)
		if err != nil {
			return
		}
	}

	if !empty_zkyn[3] {
		// write "FieldTypeStr__str"
		err = en.Append(0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.FieldTypeStr)
		if err != nil {
			return
		}
	}

	if !empty_zkyn[4] {
		// write "FieldCategory__"
		err = en.Append(0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.FieldCategory))
		if err != nil {
			return
		}
	}

	if !empty_zkyn[5] {
		// write "FieldPrimitive__"
		err = en.Append(0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f)
		if err != nil {
			return err
		}
		err = en.WriteUint64(uint64(z.FieldPrimitive))
		if err != nil {
			return
		}
	}

	if !empty_zkyn[6] {
		// write "FieldFullType__ptr"
		err = en.Append(0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72)
		if err != nil {
			return err
		}
		if z.FieldFullType == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.FieldFullType.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zkyn[7] {
		// write "OmitEmpty__boo"
		err = en.Append(0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.OmitEmpty)
		if err != nil {
			return
		}
	}

	if !empty_zkyn[8] {
		// write "Skip__boo"
		err = en.Append(0xa9, 0x53, 0x6b, 0x69, 0x70, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Skip)
		if err != nil {
			return
		}
	}

	if !empty_zkyn[9] {
		// write "Deprecated__boo"
		err = en.Append(0xaf, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.Deprecated)
		if err != nil {
			return
		}
	}

	if !empty_zkyn[10] {
		// write "ShowZero__boo"
		err = en.Append(0xad, 0x53, 0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.ShowZero)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Field) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [11]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Zid__i64"
	o = append(o, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.Zid)
	// string "FieldGoName__str"
	o = append(o, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.FieldGoName)
	if !empty[2] {
		// string "FieldTagName__str"
		o = append(o, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.FieldTagName)
	}

	if !empty[3] {
		// string "FieldTypeStr__str"
		o = append(o, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.FieldTypeStr)
	}

	if !empty[4] {
		// string "FieldCategory__"
		o = append(o, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f)
		o = msgp.AppendUint64(o, uint64(z.FieldCategory))
	}

	if !empty[5] {
		// string "FieldPrimitive__"
		o = append(o, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f)
		o = msgp.AppendUint64(o, uint64(z.FieldPrimitive))
	}

	if !empty[6] {
		// string "FieldFullType__ptr"
		o = append(o, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72)
		if z.FieldFullType == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.FieldFullType.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[7] {
		// string "OmitEmpty__boo"
		o = append(o, 0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.OmitEmpty)
	}

	if !empty[8] {
		// string "Skip__boo"
		o = append(o, 0xa9, 0x53, 0x6b, 0x69, 0x70, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.Skip)
	}

	if !empty[9] {
		// string "Deprecated__boo"
		o = append(o, 0xaf, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.Deprecated)
	}

	if !empty[10] {
		// string "ShowZero__boo"
		o = append(o, 0xad, 0x53, 0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f, 0x5f, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.ShowZero)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Field) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Field) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zggx = 11

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zggx uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zggx, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zggx := totalEncodedFields1zggx
	missingFieldsLeft1zggx := maxFields1zggx - totalEncodedFields1zggx

	var nextMiss1zggx int32 = -1
	var found1zggx [maxFields1zggx]bool
	var curField1zggx string

doneWithStruct1zggx:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zggx > 0 || missingFieldsLeft1zggx > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zggx, missingFieldsLeft1zggx, msgp.ShowFound(found1zggx[:]), unmarshalMsgFieldOrder1zggx)
		if encodedFieldsLeft1zggx > 0 {
			encodedFieldsLeft1zggx--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zggx = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zggx < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zggx = 0
			}
			for nextMiss1zggx < maxFields1zggx && (found1zggx[nextMiss1zggx] || unmarshalMsgFieldSkip1zggx[nextMiss1zggx]) {
				nextMiss1zggx++
			}
			if nextMiss1zggx == maxFields1zggx {
				// filled all the empty fields!
				break doneWithStruct1zggx
			}
			missingFieldsLeft1zggx--
			curField1zggx = unmarshalMsgFieldOrder1zggx[nextMiss1zggx]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zggx)
		switch curField1zggx {
		// -- templateUnmarshalMsg ends here --

		case "Zid__i64":
			found1zggx[0] = true
			z.Zid, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "FieldGoName__str":
			found1zggx[1] = true
			z.FieldGoName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "FieldTagName__str":
			found1zggx[2] = true
			z.FieldTagName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "FieldTypeStr__str":
			found1zggx[3] = true
			z.FieldTypeStr, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "FieldCategory__":
			found1zggx[4] = true
			{
				var zrla uint64
				zrla, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					return
				}
				z.FieldCategory = Zkind(zrla)
			}
		case "FieldPrimitive__":
			found1zggx[5] = true
			{
				var zgxy uint64
				zgxy, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					return
				}
				z.FieldPrimitive = Zkind(zgxy)
			}
		case "FieldFullType__ptr":
			found1zggx[6] = true
			if nbs.AlwaysNil {
				if z.FieldFullType != nil {
					z.FieldFullType.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.FieldFullType {
						z.FieldFullType.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.FieldFullType == nil {
						z.FieldFullType = new(Ztype)
					}
					bts, err = z.FieldFullType.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "OmitEmpty__boo":
			found1zggx[7] = true
			z.OmitEmpty, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "Skip__boo":
			found1zggx[8] = true
			z.Skip, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "Deprecated__boo":
			found1zggx[9] = true
			z.Deprecated, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case "ShowZero__boo":
			found1zggx[10] = true
			z.ShowZero, bts, err = nbs.ReadBoolBytes(bts)

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
	if nextMiss1zggx != -1 {
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

// fields of Field
var unmarshalMsgFieldOrder1zggx = []string{"Zid__i64", "FieldGoName__str", "FieldTagName__str", "FieldTypeStr__str", "FieldCategory__", "FieldPrimitive__", "FieldFullType__ptr", "OmitEmpty__boo", "Skip__boo", "Deprecated__boo", "ShowZero__boo"}

var unmarshalMsgFieldSkip1zggx = []bool{false, false, false, false, false, false, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Field) Msgsize() (s int) {
	s = 1 + 9 + msgp.Int64Size + 17 + msgp.StringPrefixSize + len(z.FieldGoName) + 18 + msgp.StringPrefixSize + len(z.FieldTagName) + 18 + msgp.StringPrefixSize + len(z.FieldTypeStr) + 16 + msgp.Uint64Size + 17 + msgp.Uint64Size + 19
	if z.FieldFullType == nil {
		s += msgp.NilSize
	} else {
		s += z.FieldFullType.Msgsize()
	}
	s += 15 + msgp.BoolSize + 10 + msgp.BoolSize + 16 + msgp.BoolSize + 14 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Schema) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zktt = 5

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zktt uint32
	totalEncodedFields2zktt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zktt := totalEncodedFields2zktt
	missingFieldsLeft2zktt := maxFields2zktt - totalEncodedFields2zktt

	var nextMiss2zktt int32 = -1
	var found2zktt [maxFields2zktt]bool
	var curField2zktt string

doneWithStruct2zktt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zktt > 0 || missingFieldsLeft2zktt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zktt, missingFieldsLeft2zktt, msgp.ShowFound(found2zktt[:]), decodeMsgFieldOrder2zktt)
		if encodedFieldsLeft2zktt > 0 {
			encodedFieldsLeft2zktt--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zktt = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zktt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zktt = 0
			}
			for nextMiss2zktt < maxFields2zktt && (found2zktt[nextMiss2zktt] || decodeMsgFieldSkip2zktt[nextMiss2zktt]) {
				nextMiss2zktt++
			}
			if nextMiss2zktt == maxFields2zktt {
				// filled all the empty fields!
				break doneWithStruct2zktt
			}
			missingFieldsLeft2zktt--
			curField2zktt = decodeMsgFieldOrder2zktt[nextMiss2zktt]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zktt)
		switch curField2zktt {
		// -- templateDecodeMsg ends here --

		case "SourcePath__str":
			found2zktt[0] = true
			z.SourcePath, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SourcePackage__str":
			found2zktt[1] = true
			z.SourcePackage, err = dc.ReadString()
			if err != nil {
				return
			}
		case "GreenSchemaId__i64":
			found2zktt[2] = true
			z.GreenSchemaId, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Structs__map":
			found2zktt[3] = true
			var zttz uint32
			zttz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Structs == nil && zttz > 0 {
				z.Structs = make(map[string]*Struct, zttz)
			} else if len(z.Structs) > 0 {
				for key, _ := range z.Structs {
					delete(z.Structs, key)
				}
			}
			for zttz > 0 {
				zttz--
				var znok string
				var znzn *Struct
				znok, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					znzn = nil
				} else {
					if znzn == nil {
						znzn = new(Struct)
					}
					const maxFields3zerw = 2

					// -- templateDecodeMsg starts here--
					var totalEncodedFields3zerw uint32
					totalEncodedFields3zerw, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					encodedFieldsLeft3zerw := totalEncodedFields3zerw
					missingFieldsLeft3zerw := maxFields3zerw - totalEncodedFields3zerw

					var nextMiss3zerw int32 = -1
					var found3zerw [maxFields3zerw]bool
					var curField3zerw string

				doneWithStruct3zerw:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft3zerw > 0 || missingFieldsLeft3zerw > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zerw, missingFieldsLeft3zerw, msgp.ShowFound(found3zerw[:]), decodeMsgFieldOrder3zerw)
						if encodedFieldsLeft3zerw > 0 {
							encodedFieldsLeft3zerw--
							field, err = dc.ReadMapKeyPtr()
							if err != nil {
								return
							}
							curField3zerw = msgp.UnsafeString(field)
						} else {
							//missing fields need handling
							if nextMiss3zerw < 0 {
								// tell the reader to only give us Nils
								// until further notice.
								dc.PushAlwaysNil()
								nextMiss3zerw = 0
							}
							for nextMiss3zerw < maxFields3zerw && (found3zerw[nextMiss3zerw] || decodeMsgFieldSkip3zerw[nextMiss3zerw]) {
								nextMiss3zerw++
							}
							if nextMiss3zerw == maxFields3zerw {
								// filled all the empty fields!
								break doneWithStruct3zerw
							}
							missingFieldsLeft3zerw--
							curField3zerw = decodeMsgFieldOrder3zerw[nextMiss3zerw]
						}
						//fmt.Printf("switching on curField: '%v'\n", curField3zerw)
						switch curField3zerw {
						// -- templateDecodeMsg ends here --

						case "StructName__str":
							found3zerw[0] = true
							znzn.StructName, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Fields__slc":
							found3zerw[1] = true
							var zrtt uint32
							zrtt, err = dc.ReadArrayHeader()
							if err != nil {
								return
							}
							if cap(znzn.Fields) >= int(zrtt) {
								znzn.Fields = (znzn.Fields)[:zrtt]
							} else {
								znzn.Fields = make([]Field, zrtt)
							}
							for zqhr := range znzn.Fields {
								err = znzn.Fields[zqhr].DecodeMsg(dc)
								if err != nil {
									return
								}
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
					if nextMiss3zerw != -1 {
						dc.PopAlwaysNil()
					}

				}
				z.Structs[znok] = znzn
			}
		case "Imports__slc":
			found2zktt[4] = true
			var zwgb uint32
			zwgb, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Imports) >= int(zwgb) {
				z.Imports = (z.Imports)[:zwgb]
			} else {
				z.Imports = make([]string, zwgb)
			}
			for zlfy := range z.Imports {
				z.Imports[zlfy], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss2zktt != -1 {
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

// fields of Schema
var decodeMsgFieldOrder2zktt = []string{"SourcePath__str", "SourcePackage__str", "GreenSchemaId__i64", "Structs__map", "Imports__slc"}

var decodeMsgFieldSkip2zktt = []bool{false, false, false, false, false}

// fields of Struct
var decodeMsgFieldOrder3zerw = []string{"StructName__str", "Fields__slc"}

var decodeMsgFieldSkip3zerw = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Schema) fieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// EncodeMsg implements msgp.Encodable
func (z *Schema) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 5
	// write "SourcePath__str"
	err = en.Append(0x85, 0xaf, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePath)
	if err != nil {
		return
	}
	// write "SourcePackage__str"
	err = en.Append(0xb2, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SourcePackage)
	if err != nil {
		return
	}
	// write "GreenSchemaId__i64"
	err = en.Append(0xb2, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.GreenSchemaId)
	if err != nil {
		return
	}
	// write "Structs__map"
	err = en.Append(0xac, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Structs)))
	if err != nil {
		return
	}
	for znok, znzn := range z.Structs {
		err = en.WriteString(znok)
		if err != nil {
			return
		}
		if znzn == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 2
			// write "StructName__str"
			err = en.Append(0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
			if err != nil {
				return err
			}
			err = en.WriteString(znzn.StructName)
			if err != nil {
				return
			}
			// write "Fields__slc"
			err = en.Append(0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63)
			if err != nil {
				return err
			}
			err = en.WriteArrayHeader(uint32(len(znzn.Fields)))
			if err != nil {
				return
			}
			for zqhr := range znzn.Fields {
				err = znzn.Fields[zqhr].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}
	// write "Imports__slc"
	err = en.Append(0xac, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Imports)))
	if err != nil {
		return
	}
	for zlfy := range z.Imports {
		err = en.WriteString(z.Imports[zlfy])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Schema) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "SourcePath__str"
	o = append(o, 0x85, 0xaf, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.SourcePath)
	// string "SourcePackage__str"
	o = append(o, 0xb2, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.SourcePackage)
	// string "GreenSchemaId__i64"
	o = append(o, 0xb2, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.GreenSchemaId)
	// string "Structs__map"
	o = append(o, 0xac, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Structs)))
	for znok, znzn := range z.Structs {
		o = msgp.AppendString(o, znok)
		if znzn == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "StructName__str"
			o = append(o, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
			o = msgp.AppendString(o, znzn.StructName)
			// string "Fields__slc"
			o = append(o, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63)
			o = msgp.AppendArrayHeader(o, uint32(len(znzn.Fields)))
			for zqhr := range znzn.Fields {
				o, err = znzn.Fields[zqhr].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}
	// string "Imports__slc"
	o = append(o, 0xac, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Imports)))
	for zlfy := range z.Imports {
		o = msgp.AppendString(o, z.Imports[zlfy])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Schema) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Schema) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zdco = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zdco uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zdco, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zdco := totalEncodedFields4zdco
	missingFieldsLeft4zdco := maxFields4zdco - totalEncodedFields4zdco

	var nextMiss4zdco int32 = -1
	var found4zdco [maxFields4zdco]bool
	var curField4zdco string

doneWithStruct4zdco:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zdco > 0 || missingFieldsLeft4zdco > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zdco, missingFieldsLeft4zdco, msgp.ShowFound(found4zdco[:]), unmarshalMsgFieldOrder4zdco)
		if encodedFieldsLeft4zdco > 0 {
			encodedFieldsLeft4zdco--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zdco = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zdco < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zdco = 0
			}
			for nextMiss4zdco < maxFields4zdco && (found4zdco[nextMiss4zdco] || unmarshalMsgFieldSkip4zdco[nextMiss4zdco]) {
				nextMiss4zdco++
			}
			if nextMiss4zdco == maxFields4zdco {
				// filled all the empty fields!
				break doneWithStruct4zdco
			}
			missingFieldsLeft4zdco--
			curField4zdco = unmarshalMsgFieldOrder4zdco[nextMiss4zdco]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zdco)
		switch curField4zdco {
		// -- templateUnmarshalMsg ends here --

		case "SourcePath__str":
			found4zdco[0] = true
			z.SourcePath, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "SourcePackage__str":
			found4zdco[1] = true
			z.SourcePackage, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "GreenSchemaId__i64":
			found4zdco[2] = true
			z.GreenSchemaId, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "Structs__map":
			found4zdco[3] = true
			if nbs.AlwaysNil {
				if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}

			} else {

				var zxeh uint32
				zxeh, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Structs == nil && zxeh > 0 {
					z.Structs = make(map[string]*Struct, zxeh)
				} else if len(z.Structs) > 0 {
					for key, _ := range z.Structs {
						delete(z.Structs, key)
					}
				}
				for zxeh > 0 {
					var znok string
					var znzn *Struct
					zxeh--
					znok, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					// default gPtr logic.
					if nbs.PeekNil(bts) && znzn == nil {
						// consume the nil
						bts, err = nbs.ReadNilBytes(bts)
						if err != nil {
							return
						}
					} else {
						// read as-if the wire has bytes, letting nbs take care of nils.

						if znzn == nil {
							znzn = new(Struct)
						}
						const maxFields5zssm = 2

						// -- templateUnmarshalMsg starts here--
						var totalEncodedFields5zssm uint32
						if !nbs.AlwaysNil {
							totalEncodedFields5zssm, bts, err = nbs.ReadMapHeaderBytes(bts)
							if err != nil {
								return
							}
						}
						encodedFieldsLeft5zssm := totalEncodedFields5zssm
						missingFieldsLeft5zssm := maxFields5zssm - totalEncodedFields5zssm

						var nextMiss5zssm int32 = -1
						var found5zssm [maxFields5zssm]bool
						var curField5zssm string

					doneWithStruct5zssm:
						// First fill all the encoded fields, then
						// treat the remaining, missing fields, as Nil.
						for encodedFieldsLeft5zssm > 0 || missingFieldsLeft5zssm > 0 {
							//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zssm, missingFieldsLeft5zssm, msgp.ShowFound(found5zssm[:]), unmarshalMsgFieldOrder5zssm)
							if encodedFieldsLeft5zssm > 0 {
								encodedFieldsLeft5zssm--
								field, bts, err = nbs.ReadMapKeyZC(bts)
								if err != nil {
									return
								}
								curField5zssm = msgp.UnsafeString(field)
							} else {
								//missing fields need handling
								if nextMiss5zssm < 0 {
									// set bts to contain just mnil (0xc0)
									bts = nbs.PushAlwaysNil(bts)
									nextMiss5zssm = 0
								}
								for nextMiss5zssm < maxFields5zssm && (found5zssm[nextMiss5zssm] || unmarshalMsgFieldSkip5zssm[nextMiss5zssm]) {
									nextMiss5zssm++
								}
								if nextMiss5zssm == maxFields5zssm {
									// filled all the empty fields!
									break doneWithStruct5zssm
								}
								missingFieldsLeft5zssm--
								curField5zssm = unmarshalMsgFieldOrder5zssm[nextMiss5zssm]
							}
							//fmt.Printf("switching on curField: '%v'\n", curField5zssm)
							switch curField5zssm {
							// -- templateUnmarshalMsg ends here --

							case "StructName__str":
								found5zssm[0] = true
								znzn.StructName, bts, err = nbs.ReadStringBytes(bts)

								if err != nil {
									return
								}
							case "Fields__slc":
								found5zssm[1] = true
								if nbs.AlwaysNil {
									(znzn.Fields) = (znzn.Fields)[:0]
								} else {

									var znqd uint32
									znqd, bts, err = nbs.ReadArrayHeaderBytes(bts)
									if err != nil {
										return
									}
									if cap(znzn.Fields) >= int(znqd) {
										znzn.Fields = (znzn.Fields)[:znqd]
									} else {
										znzn.Fields = make([]Field, znqd)
									}
									for zqhr := range znzn.Fields {
										bts, err = znzn.Fields[zqhr].UnmarshalMsg(bts)
										if err != nil {
											return
										}
										if err != nil {
											return
										}
									}
								}
							default:
								bts, err = msgp.Skip(bts)
								if err != nil {
									return
								}
							}
						}
						if nextMiss5zssm != -1 {
							bts = nbs.PopAlwaysNil()
						}

					}
					z.Structs[znok] = znzn
				}
			}
		case "Imports__slc":
			found4zdco[4] = true
			if nbs.AlwaysNil {
				(z.Imports) = (z.Imports)[:0]
			} else {

				var zwlu uint32
				zwlu, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Imports) >= int(zwlu) {
					z.Imports = (z.Imports)[:zwlu]
				} else {
					z.Imports = make([]string, zwlu)
				}
				for zlfy := range z.Imports {
					z.Imports[zlfy], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss4zdco != -1 {
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

// fields of Schema
var unmarshalMsgFieldOrder4zdco = []string{"SourcePath__str", "SourcePackage__str", "GreenSchemaId__i64", "Structs__map", "Imports__slc"}

var unmarshalMsgFieldSkip4zdco = []bool{false, false, false, false, false}

// fields of Struct
var unmarshalMsgFieldOrder5zssm = []string{"StructName__str", "Fields__slc"}

var unmarshalMsgFieldSkip5zssm = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Schema) Msgsize() (s int) {
	s = 1 + 16 + msgp.StringPrefixSize + len(z.SourcePath) + 19 + msgp.StringPrefixSize + len(z.SourcePackage) + 19 + msgp.Int64Size + 13 + msgp.MapHeaderSize
	if z.Structs != nil {
		for znok, znzn := range z.Structs {
			_ = znzn
			_ = znok
			s += msgp.StringPrefixSize + len(znok)
			if znzn == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 16 + msgp.StringPrefixSize + len(znzn.StructName) + 12 + msgp.ArrayHeaderSize
				for zqhr := range znzn.Fields {
					s += znzn.Fields[zqhr].Msgsize()
				}
			}
		}
	}
	s += 13 + msgp.ArrayHeaderSize
	for zlfy := range z.Imports {
		s += msgp.StringPrefixSize + len(z.Imports[zlfy])
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Struct) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields6zpqd = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields6zpqd uint32
	totalEncodedFields6zpqd, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zpqd := totalEncodedFields6zpqd
	missingFieldsLeft6zpqd := maxFields6zpqd - totalEncodedFields6zpqd

	var nextMiss6zpqd int32 = -1
	var found6zpqd [maxFields6zpqd]bool
	var curField6zpqd string

doneWithStruct6zpqd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zpqd > 0 || missingFieldsLeft6zpqd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zpqd, missingFieldsLeft6zpqd, msgp.ShowFound(found6zpqd[:]), decodeMsgFieldOrder6zpqd)
		if encodedFieldsLeft6zpqd > 0 {
			encodedFieldsLeft6zpqd--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField6zpqd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss6zpqd < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zpqd = 0
			}
			for nextMiss6zpqd < maxFields6zpqd && (found6zpqd[nextMiss6zpqd] || decodeMsgFieldSkip6zpqd[nextMiss6zpqd]) {
				nextMiss6zpqd++
			}
			if nextMiss6zpqd == maxFields6zpqd {
				// filled all the empty fields!
				break doneWithStruct6zpqd
			}
			missingFieldsLeft6zpqd--
			curField6zpqd = decodeMsgFieldOrder6zpqd[nextMiss6zpqd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zpqd)
		switch curField6zpqd {
		// -- templateDecodeMsg ends here --

		case "StructName__str":
			found6zpqd[0] = true
			z.StructName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Fields__slc":
			found6zpqd[1] = true
			var zeon uint32
			zeon, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Fields) >= int(zeon) {
				z.Fields = (z.Fields)[:zeon]
			} else {
				z.Fields = make([]Field, zeon)
			}
			for zjlq := range z.Fields {
				err = z.Fields[zjlq].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zpqd != -1 {
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

// fields of Struct
var decodeMsgFieldOrder6zpqd = []string{"StructName__str", "Fields__slc"}

var decodeMsgFieldSkip6zpqd = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Struct) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *Struct) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 2
	// write "StructName__str"
	err = en.Append(0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StructName)
	if err != nil {
		return
	}
	// write "Fields__slc"
	err = en.Append(0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Fields)))
	if err != nil {
		return
	}
	for zjlq := range z.Fields {
		err = z.Fields[zjlq].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Struct) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "StructName__str"
	o = append(o, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.StructName)
	// string "Fields__slc"
	o = append(o, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Fields)))
	for zjlq := range z.Fields {
		o, err = z.Fields[zjlq].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Struct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Struct) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields7zjhb = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields7zjhb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields7zjhb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft7zjhb := totalEncodedFields7zjhb
	missingFieldsLeft7zjhb := maxFields7zjhb - totalEncodedFields7zjhb

	var nextMiss7zjhb int32 = -1
	var found7zjhb [maxFields7zjhb]bool
	var curField7zjhb string

doneWithStruct7zjhb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft7zjhb > 0 || missingFieldsLeft7zjhb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zjhb, missingFieldsLeft7zjhb, msgp.ShowFound(found7zjhb[:]), unmarshalMsgFieldOrder7zjhb)
		if encodedFieldsLeft7zjhb > 0 {
			encodedFieldsLeft7zjhb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField7zjhb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss7zjhb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss7zjhb = 0
			}
			for nextMiss7zjhb < maxFields7zjhb && (found7zjhb[nextMiss7zjhb] || unmarshalMsgFieldSkip7zjhb[nextMiss7zjhb]) {
				nextMiss7zjhb++
			}
			if nextMiss7zjhb == maxFields7zjhb {
				// filled all the empty fields!
				break doneWithStruct7zjhb
			}
			missingFieldsLeft7zjhb--
			curField7zjhb = unmarshalMsgFieldOrder7zjhb[nextMiss7zjhb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField7zjhb)
		switch curField7zjhb {
		// -- templateUnmarshalMsg ends here --

		case "StructName__str":
			found7zjhb[0] = true
			z.StructName, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Fields__slc":
			found7zjhb[1] = true
			if nbs.AlwaysNil {
				(z.Fields) = (z.Fields)[:0]
			} else {

				var zstn uint32
				zstn, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Fields) >= int(zstn) {
					z.Fields = (z.Fields)[:zstn]
				} else {
					z.Fields = make([]Field, zstn)
				}
				for zjlq := range z.Fields {
					bts, err = z.Fields[zjlq].UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss7zjhb != -1 {
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

// fields of Struct
var unmarshalMsgFieldOrder7zjhb = []string{"StructName__str", "Fields__slc"}

var unmarshalMsgFieldSkip7zjhb = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 16 + msgp.StringPrefixSize + len(z.StructName) + 12 + msgp.ArrayHeaderSize
	for zjlq := range z.Fields {
		s += z.Fields[zjlq].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Zkind) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	{
		var ztwi uint64
		ztwi, err = dc.ReadUint64()
		(*z) = Zkind(ztwi)
	}
	if err != nil {
		return
	}
	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z Zkind) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	err = en.WriteUint64(uint64(z))
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Zkind) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Zkind) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Zkind) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	{
		var zqxu uint64
		zqxu, bts, err = nbs.ReadUint64Bytes(bts)

		if err != nil {
			return
		}
		(*z) = Zkind(zqxu)
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

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Zkind) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Ztype) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields8zbqj = 4

	// -- templateDecodeMsg starts here--
	var totalEncodedFields8zbqj uint32
	totalEncodedFields8zbqj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft8zbqj := totalEncodedFields8zbqj
	missingFieldsLeft8zbqj := maxFields8zbqj - totalEncodedFields8zbqj

	var nextMiss8zbqj int32 = -1
	var found8zbqj [maxFields8zbqj]bool
	var curField8zbqj string

doneWithStruct8zbqj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft8zbqj > 0 || missingFieldsLeft8zbqj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zbqj, missingFieldsLeft8zbqj, msgp.ShowFound(found8zbqj[:]), decodeMsgFieldOrder8zbqj)
		if encodedFieldsLeft8zbqj > 0 {
			encodedFieldsLeft8zbqj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField8zbqj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss8zbqj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss8zbqj = 0
			}
			for nextMiss8zbqj < maxFields8zbqj && (found8zbqj[nextMiss8zbqj] || decodeMsgFieldSkip8zbqj[nextMiss8zbqj]) {
				nextMiss8zbqj++
			}
			if nextMiss8zbqj == maxFields8zbqj {
				// filled all the empty fields!
				break doneWithStruct8zbqj
			}
			missingFieldsLeft8zbqj--
			curField8zbqj = decodeMsgFieldOrder8zbqj[nextMiss8zbqj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField8zbqj)
		switch curField8zbqj {
		// -- templateDecodeMsg ends here --

		case "Kind__":
			found8zbqj[0] = true
			{
				var ztyx uint64
				ztyx, err = dc.ReadUint64()
				z.Kind = Zkind(ztyx)
			}
			if err != nil {
				return
			}
		case "Str__str":
			found8zbqj[1] = true
			z.Str, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Domain__ptr":
			found8zbqj[2] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.Domain != nil {
					dc.PushAlwaysNil()
					err = z.Domain.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.Domain == nil {
					z.Domain = new(Ztype)
				}
				err = z.Domain.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "Range__ptr":
			found8zbqj[3] = true
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}

				if z.Range != nil {
					dc.PushAlwaysNil()
					err = z.Range.DecodeMsg(dc)
					if err != nil {
						return
					}
					dc.PopAlwaysNil()
				}
			} else {
				// not Nil, we have something to read

				if z.Range == nil {
					z.Range = new(Ztype)
				}
				err = z.Range.DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss8zbqj != -1 {
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

// fields of Ztype
var decodeMsgFieldOrder8zbqj = []string{"Kind__", "Str__str", "Domain__ptr", "Range__ptr"}

var decodeMsgFieldSkip8zbqj = []bool{false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Ztype) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 4
	}
	var fieldsInUse uint32 = 4
	isempty[1] = (len(z.Str) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Domain == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Range == nil) // pointer, omitempty
	if isempty[3] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Ztype) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zvky [4]bool
	fieldsInUse_zpki := z.fieldsNotEmpty(empty_zvky[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpki)
	if err != nil {
		return err
	}

	// write "Kind__"
	err = en.Append(0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f)
	if err != nil {
		return err
	}
	err = en.WriteUint64(uint64(z.Kind))
	if err != nil {
		return
	}
	if !empty_zvky[1] {
		// write "Str__str"
		err = en.Append(0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Str)
		if err != nil {
			return
		}
	}

	if !empty_zvky[2] {
		// write "Domain__ptr"
		err = en.Append(0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72)
		if err != nil {
			return err
		}
		if z.Domain == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Domain.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvky[3] {
		// write "Range__ptr"
		err = en.Append(0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72)
		if err != nil {
			return err
		}
		if z.Range == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Range.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Ztype) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [4]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	// string "Kind__"
	o = append(o, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f)
	o = msgp.AppendUint64(o, uint64(z.Kind))
	if !empty[1] {
		// string "Str__str"
		o = append(o, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Str)
	}

	if !empty[2] {
		// string "Domain__ptr"
		o = append(o, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72)
		if z.Domain == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Domain.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[3] {
		// string "Range__ptr"
		o = append(o, 0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72)
		if z.Range == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Range.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Ztype) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Ztype) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields9zagm = 4

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields9zagm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zagm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft9zagm := totalEncodedFields9zagm
	missingFieldsLeft9zagm := maxFields9zagm - totalEncodedFields9zagm

	var nextMiss9zagm int32 = -1
	var found9zagm [maxFields9zagm]bool
	var curField9zagm string

doneWithStruct9zagm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zagm > 0 || missingFieldsLeft9zagm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zagm, missingFieldsLeft9zagm, msgp.ShowFound(found9zagm[:]), unmarshalMsgFieldOrder9zagm)
		if encodedFieldsLeft9zagm > 0 {
			encodedFieldsLeft9zagm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField9zagm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss9zagm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zagm = 0
			}
			for nextMiss9zagm < maxFields9zagm && (found9zagm[nextMiss9zagm] || unmarshalMsgFieldSkip9zagm[nextMiss9zagm]) {
				nextMiss9zagm++
			}
			if nextMiss9zagm == maxFields9zagm {
				// filled all the empty fields!
				break doneWithStruct9zagm
			}
			missingFieldsLeft9zagm--
			curField9zagm = unmarshalMsgFieldOrder9zagm[nextMiss9zagm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zagm)
		switch curField9zagm {
		// -- templateUnmarshalMsg ends here --

		case "Kind__":
			found9zagm[0] = true
			{
				var zidi uint64
				zidi, bts, err = nbs.ReadUint64Bytes(bts)

				if err != nil {
					return
				}
				z.Kind = Zkind(zidi)
			}
		case "Str__str":
			found9zagm[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Domain__ptr":
			found9zagm[2] = true
			if nbs.AlwaysNil {
				if z.Domain != nil {
					z.Domain.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Domain {
						z.Domain.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Domain == nil {
						z.Domain = new(Ztype)
					}
					bts, err = z.Domain.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Range__ptr":
			found9zagm[3] = true
			if nbs.AlwaysNil {
				if z.Range != nil {
					z.Range.UnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Range {
						z.Range.UnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Range == nil {
						z.Range = new(Ztype)
					}
					bts, err = z.Range.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss9zagm != -1 {
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

// fields of Ztype
var unmarshalMsgFieldOrder9zagm = []string{"Kind__", "Str__str", "Domain__ptr", "Range__ptr"}

var unmarshalMsgFieldSkip9zagm = []bool{false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Ztype) Msgsize() (s int) {
	s = 1 + 7 + msgp.Uint64Size + 9 + msgp.StringPrefixSize + len(z.Str) + 12
	if z.Domain == nil {
		s += msgp.NilSize
	} else {
		s += z.Domain.Msgsize()
	}
	s += 11
	if z.Range == nil {
		s += msgp.NilSize
	} else {
		s += z.Range.Msgsize()
	}
	return
}

// FileGreen holds Greenpack schema from file 'green.go'
type FileGreen struct{}

// ZebraSchemaInMsgpack2Format provides the Greenpack Schema in msgpack2 format, length 4375 bytes
func (FileGreen) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaf, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa8, 0x67, 0x72,
		0x65, 0x65, 0x6e, 0x2e, 0x67, 0x6f, 0xb2, 0x53, 0x6f, 0x75,
		0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x67, 0x72, 0x65, 0x65,
		0x6e, 0xb2, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63, 0x68,
		0x65, 0x6d, 0x61, 0x49, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34,
		0xd3, 0x00, 0x01, 0xa5, 0xa9, 0x4b, 0xd4, 0x96, 0x24, 0xac,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x5f, 0x5f, 0x6d,
		0x61, 0x70, 0x84, 0xa6, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
		0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x95, 0x87, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xaa, 0x53, 0x6f, 0x75, 0x72,
		0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0xb1, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63,
		0x65, 0x50, 0x61, 0x74, 0x68, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x5f, 0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87,
		0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad, 0x53, 0x6f,
		0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
		0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad,
		0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
		0x61, 0x67, 0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x5f, 0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad, 0x47, 0x72, 0x65, 0x65,
		0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad, 0x47, 0x72,
		0x65, 0x65, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49,
		0x64, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xaf, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
		0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x11,
		0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x11, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xb2, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x18, 0xb2,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x84, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x18, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaa, 0x52, 0x61,
		0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x1c, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69,
		0x6e, 0x74, 0x65, 0x72, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x19, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63,
		0x74, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7,
		0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7, 0x49, 0x6d, 0x70, 0x6f,
		0x72, 0x74, 0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa8, 0x5b, 0x5d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x1a, 0xb2, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x1a, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65,
		0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa6, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63,
		0x92, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xaa,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xaa, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f,
		0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x02, 0xb2,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa7, 0x5b, 0x5d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0xaf,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x1a, 0xb2, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x1a, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xab,
		0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x16,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0xa5, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x9b, 0x87, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x5a, 0x69, 0x64, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x5a, 0x69,
		0x64, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xaf, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
		0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x11,
		0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x11, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x5f, 0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f,
		0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x02, 0xb2,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d,
		0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f, 0xc3, 0x88,
		0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
		0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f,
		0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72,
		0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xae, 0x4f, 0x6d, 0x69, 0x74,
		0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f,
		0xc3, 0x88, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa5, 0x5a, 0x6b, 0x69, 0x6e, 0x64, 0xaf,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
		0x65, 0x5f, 0x5f, 0x0b, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f,
		0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f,
		0x5f, 0x0b, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xae, 0x4f,
		0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f,
		0x62, 0x6f, 0x6f, 0xc3, 0x88, 0xa8, 0x5a, 0x69, 0x64, 0x5f,
		0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x5a,
		0x6b, 0x69, 0x6e, 0x64, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f,
		0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x0b, 0xb2,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x0b, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d,
		0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f, 0xc3, 0x87,
		0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x1c, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x1c, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
		0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x16, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xae, 0x4f, 0x6d, 0x69,
		0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f,
		0x6f, 0xc3, 0x88, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69,
		0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa9, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa9, 0x4f,
		0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f,
		0x6c, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x5f, 0x5f, 0x12, 0xb2, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x12, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xae, 0x4f,
		0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f,
		0x62, 0x6f, 0x6f, 0xc3, 0x88, 0xa8, 0x5a, 0x69, 0x64, 0x5f,
		0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa4, 0x53, 0x6b, 0x69, 0x70, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x6b, 0x69, 0x70,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x62,
		0x6f, 0x6f, 0x6c, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x12, 0xb2, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b,
		0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x12, 0xa8, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c,
		0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
		0x5f, 0x5f, 0x62, 0x6f, 0x6f, 0xc3, 0x88, 0xa8, 0x5a, 0x69,
		0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65,
		0x63, 0x61, 0x74, 0x65, 0x64, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xaa, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
		0x61, 0x74, 0x65, 0x64, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xaf, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f,
		0x5f, 0x12, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x12,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4,
		0x62, 0x6f, 0x6f, 0x6c, 0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45,
		0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f, 0xc3,
		0x88, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34,
		0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa8, 0x53,
		0x68, 0x6f, 0x77, 0x5a, 0x65, 0x72, 0x6f, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa8, 0x53, 0x68, 0x6f, 0x77,
		0x5a, 0x65, 0x72, 0x6f, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xaf, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f,
		0x5f, 0x12, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x12,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4,
		0x62, 0x6f, 0x6f, 0x6c, 0xae, 0x4f, 0x6d, 0x69, 0x74, 0x45,
		0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f, 0xc3,
		0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0x82, 0xaf, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c,
		0x63, 0x94, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69,
		0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x5a, 0x6b, 0x69,
		0x6e, 0x64, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
		0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x0b, 0xb2, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x0b, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0x88, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3,
		0x53, 0x74, 0x72, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa3, 0x53, 0x74, 0x72, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x5f, 0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xae,
		0x4f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f,
		0x5f, 0x62, 0x6f, 0x6f, 0xc3, 0x87, 0xa8, 0x5a, 0x69, 0x64,
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x44,
		0x6f, 0x6d, 0x61, 0x69, 0x6e, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x1c, 0xb2, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x1c, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74,
		0x65, 0x72, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x16, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa5, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xae, 0x4f,
		0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f,
		0x62, 0x6f, 0x6f, 0xc3, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f,
		0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x52, 0x61, 0x6e,
		0x67, 0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x2a, 0x5a, 0x74, 0x79, 0x70, 0x65, 0xaf, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x5f, 0x5f, 0x1c, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f,
		0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f,
		0x5f, 0x1c, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xab,
		0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x16,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5,
		0x5a, 0x74, 0x79, 0x70, 0x65, 0xae, 0x4f, 0x6d, 0x69, 0x74,
		0x45, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x5f, 0x62, 0x6f, 0x6f,
		0xc3, 0xac, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f,
		0x5f, 0x73, 0x6c, 0x63, 0x90,
	}
}

// ZebraSchemaInJsonCompact provides the Greenpack Schema in compact JSON format, length 5351 bytes
func (FileGreen) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath__str":"green.go","SourcePackage__str":"green","ZebraSchemaId__i64":463621516989988,"Structs__map":{"Schema":{"StructName__str":"Schema","Fields__slc":[{"Zid__i64":-1,"FieldGoName__str":"SourcePath","FieldTagName__str":"SourcePath","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"SourcePackage","FieldTagName__str":"SourcePackage","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"GreenSchemaId","FieldTagName__str":"GreenSchemaId","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}},{"Zid__i64":-1,"FieldGoName__str":"Structs","FieldTagName__str":"Structs","FieldTypeStr__str":"map[string]*Struct","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":25,"Str__str":"Struct"}}}},{"Zid__i64":-1,"FieldGoName__str":"Imports","FieldTagName__str":"Imports","FieldTypeStr__str":"[]string","FieldCategory__":26,"FieldFullType__ptr":{"Kind__":26,"Str__str":"Slice","Domain__ptr":{"Kind__":2,"Str__str":"string"}}}]},"Struct":{"StructName__str":"Struct","Fields__slc":[{"Zid__i64":-1,"FieldGoName__str":"StructName","FieldTagName__str":"StructName","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"Fields","FieldTagName__str":"Fields","FieldTypeStr__str":"[]Field","FieldCategory__":26,"FieldFullType__ptr":{"Kind__":26,"Str__str":"Slice","Domain__ptr":{"Kind__":22,"Str__str":"Field"}}}]},"Field":{"StructName__str":"Field","Fields__slc":[{"Zid__i64":-1,"FieldGoName__str":"Zid","FieldTagName__str":"Zid","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}},{"Zid__i64":-1,"FieldGoName__str":"FieldGoName","FieldTagName__str":"FieldGoName","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"FieldTagName","FieldTagName__str":"FieldTagName","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"FieldTypeStr","FieldTagName__str":"FieldTypeStr","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"FieldCategory","FieldTagName__str":"FieldCategory","FieldTypeStr__str":"Zkind","FieldCategory__":23,"FieldPrimitive__":11,"FieldFullType__ptr":{"Kind__":11,"Str__str":"uint64"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"FieldPrimitive","FieldTagName__str":"FieldPrimitive","FieldTypeStr__str":"Zkind","FieldCategory__":23,"FieldPrimitive__":11,"FieldFullType__ptr":{"Kind__":11,"Str__str":"uint64"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"FieldFullType","FieldTagName__str":"FieldFullType","FieldTypeStr__str":"*Ztype","FieldCategory__":28,"FieldFullType__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"Ztype"}},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"OmitEmpty","FieldTagName__str":"OmitEmpty","FieldTypeStr__str":"bool","FieldCategory__":23,"FieldPrimitive__":18,"FieldFullType__ptr":{"Kind__":18,"Str__str":"bool"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"Skip","FieldTagName__str":"Skip","FieldTypeStr__str":"bool","FieldCategory__":23,"FieldPrimitive__":18,"FieldFullType__ptr":{"Kind__":18,"Str__str":"bool"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"Deprecated","FieldTagName__str":"Deprecated","FieldTypeStr__str":"bool","FieldCategory__":23,"FieldPrimitive__":18,"FieldFullType__ptr":{"Kind__":18,"Str__str":"bool"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"ShowZero","FieldTagName__str":"ShowZero","FieldTypeStr__str":"bool","FieldCategory__":23,"FieldPrimitive__":18,"FieldFullType__ptr":{"Kind__":18,"Str__str":"bool"},"OmitEmpty__boo":true}]},"Ztype":{"StructName__str":"Ztype","Fields__slc":[{"Zid__i64":-1,"FieldGoName__str":"Kind","FieldTagName__str":"Kind","FieldTypeStr__str":"Zkind","FieldCategory__":23,"FieldPrimitive__":11,"FieldFullType__ptr":{"Kind__":11,"Str__str":"uint64"}},{"Zid__i64":-1,"FieldGoName__str":"Str","FieldTagName__str":"Str","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"Domain","FieldTagName__str":"Domain","FieldTypeStr__str":"*Ztype","FieldCategory__":28,"FieldFullType__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"Ztype"}},"OmitEmpty__boo":true},{"Zid__i64":-1,"FieldGoName__str":"Range","FieldTagName__str":"Range","FieldTypeStr__str":"*Ztype","FieldCategory__":28,"FieldFullType__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"Ztype"}},"OmitEmpty__boo":true}]}},"Imports__slc":[]}`)
}

// ZebraSchemaInJsonPretty provides the Greenpack Schema in pretty JSON format, length 12405 bytes
func (FileGreen) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath__str": "green.go",
    "SourcePackage__str": "green",
    "ZebraSchemaId__i64": 463621516989988,
    "Structs__map": {
        "Schema": {
            "StructName__str": "Schema",
            "Fields__slc": [
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "SourcePath",
                    "FieldTagName__str": "SourcePath",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "SourcePackage",
                    "FieldTagName__str": "SourcePackage",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "GreenSchemaId",
                    "FieldTagName__str": "GreenSchemaId",
                    "FieldTypeStr__str": "int64",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 17,
                    "FieldFullType__ptr": {
                        "Kind__": 17,
                        "Str__str": "int64"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Structs",
                    "FieldTagName__str": "Structs",
                    "FieldTypeStr__str": "map[string]*Struct",
                    "FieldCategory__": 24,
                    "FieldFullType__ptr": {
                        "Kind__": 24,
                        "Str__str": "Map",
                        "Domain__ptr": {
                            "Kind__": 2,
                            "Str__str": "string"
                        },
                        "Range__ptr": {
                            "Kind__": 28,
                            "Str__str": "Pointer",
                            "Domain__ptr": {
                                "Kind__": 25,
                                "Str__str": "Struct"
                            }
                        }
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Imports",
                    "FieldTagName__str": "Imports",
                    "FieldTypeStr__str": "[]string",
                    "FieldCategory__": 26,
                    "FieldFullType__ptr": {
                        "Kind__": 26,
                        "Str__str": "Slice",
                        "Domain__ptr": {
                            "Kind__": 2,
                            "Str__str": "string"
                        }
                    }
                }
            ]
        },
        "Struct": {
            "StructName__str": "Struct",
            "Fields__slc": [
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "StructName",
                    "FieldTagName__str": "StructName",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Fields",
                    "FieldTagName__str": "Fields",
                    "FieldTypeStr__str": "[]Field",
                    "FieldCategory__": 26,
                    "FieldFullType__ptr": {
                        "Kind__": 26,
                        "Str__str": "Slice",
                        "Domain__ptr": {
                            "Kind__": 22,
                            "Str__str": "Field"
                        }
                    }
                }
            ]
        },
        "Field": {
            "StructName__str": "Field",
            "Fields__slc": [
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Zid",
                    "FieldTagName__str": "Zid",
                    "FieldTypeStr__str": "int64",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 17,
                    "FieldFullType__ptr": {
                        "Kind__": 17,
                        "Str__str": "int64"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "FieldGoName",
                    "FieldTagName__str": "FieldGoName",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "FieldTagName",
                    "FieldTagName__str": "FieldTagName",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "FieldTypeStr",
                    "FieldTagName__str": "FieldTypeStr",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "FieldCategory",
                    "FieldTagName__str": "FieldCategory",
                    "FieldTypeStr__str": "Zkind",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 11,
                    "FieldFullType__ptr": {
                        "Kind__": 11,
                        "Str__str": "uint64"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "FieldPrimitive",
                    "FieldTagName__str": "FieldPrimitive",
                    "FieldTypeStr__str": "Zkind",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 11,
                    "FieldFullType__ptr": {
                        "Kind__": 11,
                        "Str__str": "uint64"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "FieldFullType",
                    "FieldTagName__str": "FieldFullType",
                    "FieldTypeStr__str": "*Ztype",
                    "FieldCategory__": 28,
                    "FieldFullType__ptr": {
                        "Kind__": 28,
                        "Str__str": "Pointer",
                        "Domain__ptr": {
                            "Kind__": 22,
                            "Str__str": "Ztype"
                        }
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "OmitEmpty",
                    "FieldTagName__str": "OmitEmpty",
                    "FieldTypeStr__str": "bool",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 18,
                    "FieldFullType__ptr": {
                        "Kind__": 18,
                        "Str__str": "bool"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Skip",
                    "FieldTagName__str": "Skip",
                    "FieldTypeStr__str": "bool",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 18,
                    "FieldFullType__ptr": {
                        "Kind__": 18,
                        "Str__str": "bool"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Deprecated",
                    "FieldTagName__str": "Deprecated",
                    "FieldTypeStr__str": "bool",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 18,
                    "FieldFullType__ptr": {
                        "Kind__": 18,
                        "Str__str": "bool"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "ShowZero",
                    "FieldTagName__str": "ShowZero",
                    "FieldTypeStr__str": "bool",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 18,
                    "FieldFullType__ptr": {
                        "Kind__": 18,
                        "Str__str": "bool"
                    },
                    "OmitEmpty__boo": true
                }
            ]
        },
        "Ztype": {
            "StructName__str": "Ztype",
            "Fields__slc": [
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Kind",
                    "FieldTagName__str": "Kind",
                    "FieldTypeStr__str": "Zkind",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 11,
                    "FieldFullType__ptr": {
                        "Kind__": 11,
                        "Str__str": "uint64"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Str",
                    "FieldTagName__str": "Str",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Domain",
                    "FieldTagName__str": "Domain",
                    "FieldTypeStr__str": "*Ztype",
                    "FieldCategory__": 28,
                    "FieldFullType__ptr": {
                        "Kind__": 28,
                        "Str__str": "Pointer",
                        "Domain__ptr": {
                            "Kind__": 22,
                            "Str__str": "Ztype"
                        }
                    },
                    "OmitEmpty__boo": true
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "Range",
                    "FieldTagName__str": "Range",
                    "FieldTypeStr__str": "*Ztype",
                    "FieldCategory__": 28,
                    "FieldFullType__ptr": {
                        "Kind__": 28,
                        "Str__str": "Pointer",
                        "Domain__ptr": {
                            "Kind__": 22,
                            "Str__str": "Ztype"
                        }
                    },
                    "OmitEmpty__boo": true
                }
            ]
        }
    },
    "Imports__slc": []
}`)
}
