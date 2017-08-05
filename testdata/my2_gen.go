package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Tr) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields0zbcx6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zbcx6 uint32
	totalEncodedFields0zbcx6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zbcx6 := totalEncodedFields0zbcx6
	missingFieldsLeft0zbcx6 := maxFields0zbcx6 - totalEncodedFields0zbcx6

	var nextMiss0zbcx6 int32 = -1
	var found0zbcx6 [maxFields0zbcx6]bool
	var curField0zbcx6 string

doneWithStruct0zbcx6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zbcx6 > 0 || missingFieldsLeft0zbcx6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zbcx6, missingFieldsLeft0zbcx6, msgp.ShowFound(found0zbcx6[:]), decodeMsgFieldOrder0zbcx6)
		if encodedFieldsLeft0zbcx6 > 0 {
			encodedFieldsLeft0zbcx6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zbcx6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zbcx6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zbcx6 = 0
			}
			for nextMiss0zbcx6 < maxFields0zbcx6 && (found0zbcx6[nextMiss0zbcx6] || decodeMsgFieldSkip0zbcx6[nextMiss0zbcx6]) {
				nextMiss0zbcx6++
			}
			if nextMiss0zbcx6 == maxFields0zbcx6 {
				// filled all the empty fields!
				break doneWithStruct0zbcx6
			}
			missingFieldsLeft0zbcx6--
			curField0zbcx6 = decodeMsgFieldOrder0zbcx6[nextMiss0zbcx6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zbcx6)
		switch curField0zbcx6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zbcx6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zbcx6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zbcx6[3] = true
			var zudo7 uint32
			zudo7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zudo7 > 0 {
				z.Snot = make(map[string]bool, zudo7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zudo7 > 0 {
				zudo7--
				var zcdp1 string
				var ztiy2 bool
				zcdp1, err = dc.ReadString()
				if err != nil {
					return
				}
				ztiy2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zcdp1] = ztiy2
			}
		case "Notyet_zid03_map":
			found0zbcx6[4] = true
			var zxno8 uint32
			zxno8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zxno8 > 0 {
				z.Notyet = make(map[string]bool, zxno8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zxno8 > 0 {
				zxno8--
				var zuah3 string
				var zjgd4 bool
				zuah3, err = dc.ReadString()
				if err != nil {
					return
				}
				zjgd4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zuah3] = zjgd4
			}
		case "Setm_zid04_slc":
			found0zbcx6[5] = true
			var zlnl9 uint32
			zlnl9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zlnl9) {
				z.Setm = (z.Setm)[:zlnl9]
			} else {
				z.Setm = make([]*inn, zlnl9)
			}
			for zcep5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zcep5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zcep5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zcep5] == nil {
						z.Setm[zcep5] = new(inn)
					}
					err = z.Setm[zcep5].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zbcx6 != -1 {
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

// fields of Tr
var decodeMsgFieldOrder0zbcx6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zbcx6 = []bool{false, true, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Tr) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.U) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[2] = (z.Nt == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.Snot) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.Notyet) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (len(z.Setm) == 0) // string, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Tr) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zvip10 [6]bool
	fieldsInUse_zcls11 := z.fieldsNotEmpty(empty_zvip10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zcls11)
	if err != nil {
		return err
	}

	if !empty_zvip10[0] {
		// write "U_zid00_str"
		err = en.Append(0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.U)
		if err != nil {
			return
		}
	}

	if !empty_zvip10[2] {
		// write "Nt_zid01_int"
		err = en.Append(0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Nt)
		if err != nil {
			return
		}
	}

	if !empty_zvip10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zcdp1, ztiy2 := range z.Snot {
			err = en.WriteString(zcdp1)
			if err != nil {
				return
			}
			err = en.WriteBool(ztiy2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvip10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zuah3, zjgd4 := range z.Notyet {
			err = en.WriteString(zuah3)
			if err != nil {
				return
			}
			err = en.WriteBool(zjgd4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvip10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zcep5 := range z.Setm {
			if z.Setm[zcep5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zcep5].EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Tr) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "U_zid00_str"
		o = append(o, 0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.U)
	}

	if !empty[2] {
		// string "Nt_zid01_int"
		o = append(o, 0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
		o = msgp.AppendInt(o, z.Nt)
	}

	if !empty[3] {
		// string "Snot_zid02_map"
		o = append(o, 0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
		for zcdp1, ztiy2 := range z.Snot {
			o = msgp.AppendString(o, zcdp1)
			o = msgp.AppendBool(o, ztiy2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zuah3, zjgd4 := range z.Notyet {
			o = msgp.AppendString(o, zuah3)
			o = msgp.AppendBool(o, zjgd4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zcep5 := range z.Setm {
			if z.Setm[zcep5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zcep5].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Tr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *Tr) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zgar12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zgar12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zgar12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zgar12 := totalEncodedFields1zgar12
	missingFieldsLeft1zgar12 := maxFields1zgar12 - totalEncodedFields1zgar12

	var nextMiss1zgar12 int32 = -1
	var found1zgar12 [maxFields1zgar12]bool
	var curField1zgar12 string

doneWithStruct1zgar12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgar12 > 0 || missingFieldsLeft1zgar12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgar12, missingFieldsLeft1zgar12, msgp.ShowFound(found1zgar12[:]), unmarshalMsgFieldOrder1zgar12)
		if encodedFieldsLeft1zgar12 > 0 {
			encodedFieldsLeft1zgar12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zgar12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zgar12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zgar12 = 0
			}
			for nextMiss1zgar12 < maxFields1zgar12 && (found1zgar12[nextMiss1zgar12] || unmarshalMsgFieldSkip1zgar12[nextMiss1zgar12]) {
				nextMiss1zgar12++
			}
			if nextMiss1zgar12 == maxFields1zgar12 {
				// filled all the empty fields!
				break doneWithStruct1zgar12
			}
			missingFieldsLeft1zgar12--
			curField1zgar12 = unmarshalMsgFieldOrder1zgar12[nextMiss1zgar12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgar12)
		switch curField1zgar12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zgar12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zgar12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zgar12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zeel13 uint32
				zeel13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zeel13 > 0 {
					z.Snot = make(map[string]bool, zeel13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zeel13 > 0 {
					var zcdp1 string
					var ztiy2 bool
					zeel13--
					zcdp1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					ztiy2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zcdp1] = ztiy2
				}
			}
		case "Notyet_zid03_map":
			found1zgar12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zwdn14 uint32
				zwdn14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zwdn14 > 0 {
					z.Notyet = make(map[string]bool, zwdn14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zwdn14 > 0 {
					var zuah3 string
					var zjgd4 bool
					zwdn14--
					zuah3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zjgd4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zuah3] = zjgd4
				}
			}
		case "Setm_zid04_slc":
			found1zgar12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zkmn15 uint32
				zkmn15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zkmn15) {
					z.Setm = (z.Setm)[:zkmn15]
				} else {
					z.Setm = make([]*inn, zkmn15)
				}
				for zcep5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zcep5] != nil {
							z.Setm[zcep5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zcep5] {
								z.Setm[zcep5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zcep5] == nil {
								z.Setm[zcep5] = new(inn)
							}
							bts, err = z.Setm[zcep5].UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
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
	if nextMiss1zgar12 != -1 {
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

// fields of Tr
var unmarshalMsgFieldOrder1zgar12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zgar12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zcdp1, ztiy2 := range z.Snot {
			_ = ztiy2
			_ = zcdp1
			s += msgp.StringPrefixSize + len(zcdp1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zuah3, zjgd4 := range z.Notyet {
			_ = zjgd4
			_ = zuah3
			s += msgp.StringPrefixSize + len(zuah3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zcep5 := range z.Setm {
		if z.Setm[zcep5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zcep5].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *inn) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields2zhjp18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zhjp18 uint32
	totalEncodedFields2zhjp18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zhjp18 := totalEncodedFields2zhjp18
	missingFieldsLeft2zhjp18 := maxFields2zhjp18 - totalEncodedFields2zhjp18

	var nextMiss2zhjp18 int32 = -1
	var found2zhjp18 [maxFields2zhjp18]bool
	var curField2zhjp18 string

doneWithStruct2zhjp18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zhjp18 > 0 || missingFieldsLeft2zhjp18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zhjp18, missingFieldsLeft2zhjp18, msgp.ShowFound(found2zhjp18[:]), decodeMsgFieldOrder2zhjp18)
		if encodedFieldsLeft2zhjp18 > 0 {
			encodedFieldsLeft2zhjp18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zhjp18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zhjp18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zhjp18 = 0
			}
			for nextMiss2zhjp18 < maxFields2zhjp18 && (found2zhjp18[nextMiss2zhjp18] || decodeMsgFieldSkip2zhjp18[nextMiss2zhjp18]) {
				nextMiss2zhjp18++
			}
			if nextMiss2zhjp18 == maxFields2zhjp18 {
				// filled all the empty fields!
				break doneWithStruct2zhjp18
			}
			missingFieldsLeft2zhjp18--
			curField2zhjp18 = decodeMsgFieldOrder2zhjp18[nextMiss2zhjp18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zhjp18)
		switch curField2zhjp18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zhjp18[0] = true
			var zwaj19 uint32
			zwaj19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zwaj19 > 0 {
				z.j = make(map[string]bool, zwaj19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zwaj19 > 0 {
				zwaj19--
				var zaeb16 string
				var ztld17 bool
				zaeb16, err = dc.ReadString()
				if err != nil {
					return
				}
				ztld17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zaeb16] = ztld17
			}
		case "e_zid01_i64":
			found2zhjp18[1] = true
			z.e, err = dc.ReadInt64()
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
	if nextMiss2zhjp18 != -1 {
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

// fields of inn
var decodeMsgFieldOrder2zhjp18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zhjp18 = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *inn) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.j) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.e == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *inn) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zmto20 [2]bool
	fieldsInUse_zsln21 := z.fieldsNotEmpty(empty_zmto20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zsln21)
	if err != nil {
		return err
	}

	if !empty_zmto20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zaeb16, ztld17 := range z.j {
			err = en.WriteString(zaeb16)
			if err != nil {
				return
			}
			err = en.WriteBool(ztld17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zmto20[1] {
		// write "e_zid01_i64"
		err = en.Append(0xab, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x36, 0x34)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.e)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *inn) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "j_zid00_map"
		o = append(o, 0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.j)))
		for zaeb16, ztld17 := range z.j {
			o = msgp.AppendString(o, zaeb16)
			o = msgp.AppendBool(o, ztld17)
		}
	}

	if !empty[1] {
		// string "e_zid01_i64"
		o = append(o, 0xab, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x36, 0x34)
		o = msgp.AppendInt64(o, z.e)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *inn) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *inn) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zsix22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zsix22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zsix22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zsix22 := totalEncodedFields3zsix22
	missingFieldsLeft3zsix22 := maxFields3zsix22 - totalEncodedFields3zsix22

	var nextMiss3zsix22 int32 = -1
	var found3zsix22 [maxFields3zsix22]bool
	var curField3zsix22 string

doneWithStruct3zsix22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zsix22 > 0 || missingFieldsLeft3zsix22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zsix22, missingFieldsLeft3zsix22, msgp.ShowFound(found3zsix22[:]), unmarshalMsgFieldOrder3zsix22)
		if encodedFieldsLeft3zsix22 > 0 {
			encodedFieldsLeft3zsix22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zsix22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zsix22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zsix22 = 0
			}
			for nextMiss3zsix22 < maxFields3zsix22 && (found3zsix22[nextMiss3zsix22] || unmarshalMsgFieldSkip3zsix22[nextMiss3zsix22]) {
				nextMiss3zsix22++
			}
			if nextMiss3zsix22 == maxFields3zsix22 {
				// filled all the empty fields!
				break doneWithStruct3zsix22
			}
			missingFieldsLeft3zsix22--
			curField3zsix22 = unmarshalMsgFieldOrder3zsix22[nextMiss3zsix22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zsix22)
		switch curField3zsix22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zsix22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zhjs23 uint32
				zhjs23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zhjs23 > 0 {
					z.j = make(map[string]bool, zhjs23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zhjs23 > 0 {
					var zaeb16 string
					var ztld17 bool
					zhjs23--
					zaeb16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					ztld17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zaeb16] = ztld17
				}
			}
		case "e_zid01_i64":
			found3zsix22[1] = true
			z.e, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss3zsix22 != -1 {
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

// fields of inn
var unmarshalMsgFieldOrder3zsix22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zsix22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zaeb16, ztld17 := range z.j {
			_ = ztld17
			_ = zaeb16
			s += msgp.StringPrefixSize + len(zaeb16) + msgp.BoolSize
		}
	}
	s += 12 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *u) DecodeMsg(dc *msgp.Reader) (err error) {
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
	const maxFields4zbzl26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zbzl26 uint32
	totalEncodedFields4zbzl26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zbzl26 := totalEncodedFields4zbzl26
	missingFieldsLeft4zbzl26 := maxFields4zbzl26 - totalEncodedFields4zbzl26

	var nextMiss4zbzl26 int32 = -1
	var found4zbzl26 [maxFields4zbzl26]bool
	var curField4zbzl26 string

doneWithStruct4zbzl26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zbzl26 > 0 || missingFieldsLeft4zbzl26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zbzl26, missingFieldsLeft4zbzl26, msgp.ShowFound(found4zbzl26[:]), decodeMsgFieldOrder4zbzl26)
		if encodedFieldsLeft4zbzl26 > 0 {
			encodedFieldsLeft4zbzl26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zbzl26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zbzl26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zbzl26 = 0
			}
			for nextMiss4zbzl26 < maxFields4zbzl26 && (found4zbzl26[nextMiss4zbzl26] || decodeMsgFieldSkip4zbzl26[nextMiss4zbzl26]) {
				nextMiss4zbzl26++
			}
			if nextMiss4zbzl26 == maxFields4zbzl26 {
				// filled all the empty fields!
				break doneWithStruct4zbzl26
			}
			missingFieldsLeft4zbzl26--
			curField4zbzl26 = decodeMsgFieldOrder4zbzl26[nextMiss4zbzl26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zbzl26)
		switch curField4zbzl26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zbzl26[0] = true
			var zers27 uint32
			zers27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zers27 > 0 {
				z.m = make(map[string]*Tr, zers27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zers27 > 0 {
				zers27--
				var zran24 string
				var zjhk25 *Tr
				zran24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zjhk25 != nil {
						dc.PushAlwaysNil()
						err = zjhk25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zjhk25 == nil {
						zjhk25 = new(Tr)
					}
					err = zjhk25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zran24] = zjhk25
			}
		case "s_zid01_str":
			found4zbzl26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zbzl26[2] = true
			z.n, err = dc.ReadInt64()
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
	if nextMiss4zbzl26 != -1 {
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

// fields of u
var decodeMsgFieldOrder4zbzl26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zbzl26 = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *u) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (len(z.m) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.s) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.n == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *u) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zzmw28 [3]bool
	fieldsInUse_zfig29 := z.fieldsNotEmpty(empty_zzmw28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zfig29)
	if err != nil {
		return err
	}

	if !empty_zzmw28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zran24, zjhk25 := range z.m {
			err = en.WriteString(zran24)
			if err != nil {
				return
			}
			if zjhk25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zjhk25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zzmw28[1] {
		// write "s_zid01_str"
		err = en.Append(0xab, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		if err != nil {
			return err
		}
		err = en.WriteString(z.s)
		if err != nil {
			return
		}
	}

	if !empty_zzmw28[2] {
		// write "n_zid02_i64"
		err = en.Append(0xab, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x36, 0x34)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.n)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *u) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "m_zid00_map"
		o = append(o, 0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.m)))
		for zran24, zjhk25 := range z.m {
			o = msgp.AppendString(o, zran24)
			if zjhk25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zjhk25.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[1] {
		// string "s_zid01_str"
		o = append(o, 0xab, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.s)
	}

	if !empty[2] {
		// string "n_zid02_i64"
		o = append(o, 0xab, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x36, 0x34)
		o = msgp.AppendInt64(o, z.n)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *u) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *u) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields5zdiy30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zdiy30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zdiy30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zdiy30 := totalEncodedFields5zdiy30
	missingFieldsLeft5zdiy30 := maxFields5zdiy30 - totalEncodedFields5zdiy30

	var nextMiss5zdiy30 int32 = -1
	var found5zdiy30 [maxFields5zdiy30]bool
	var curField5zdiy30 string

doneWithStruct5zdiy30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zdiy30 > 0 || missingFieldsLeft5zdiy30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zdiy30, missingFieldsLeft5zdiy30, msgp.ShowFound(found5zdiy30[:]), unmarshalMsgFieldOrder5zdiy30)
		if encodedFieldsLeft5zdiy30 > 0 {
			encodedFieldsLeft5zdiy30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zdiy30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zdiy30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zdiy30 = 0
			}
			for nextMiss5zdiy30 < maxFields5zdiy30 && (found5zdiy30[nextMiss5zdiy30] || unmarshalMsgFieldSkip5zdiy30[nextMiss5zdiy30]) {
				nextMiss5zdiy30++
			}
			if nextMiss5zdiy30 == maxFields5zdiy30 {
				// filled all the empty fields!
				break doneWithStruct5zdiy30
			}
			missingFieldsLeft5zdiy30--
			curField5zdiy30 = unmarshalMsgFieldOrder5zdiy30[nextMiss5zdiy30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zdiy30)
		switch curField5zdiy30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zdiy30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zsgz31 uint32
				zsgz31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zsgz31 > 0 {
					z.m = make(map[string]*Tr, zsgz31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zsgz31 > 0 {
					var zran24 string
					var zjhk25 *Tr
					zsgz31--
					zran24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zjhk25 != nil {
							zjhk25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zjhk25 {
								zjhk25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zjhk25 == nil {
								zjhk25 = new(Tr)
							}
							bts, err = zjhk25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zran24] = zjhk25
				}
			}
		case "s_zid01_str":
			found5zdiy30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zdiy30[2] = true
			z.n, bts, err = nbs.ReadInt64Bytes(bts)

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
	if nextMiss5zdiy30 != -1 {
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

// fields of u
var unmarshalMsgFieldOrder5zdiy30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zdiy30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zran24, zjhk25 := range z.m {
			_ = zjhk25
			_ = zran24
			s += msgp.StringPrefixSize + len(zran24)
			if zjhk25 == nil {
				s += msgp.NilSize
			} else {
				s += zjhk25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
