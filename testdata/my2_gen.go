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
	const maxFields0zfyt6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zfyt6 uint32
	totalEncodedFields0zfyt6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zfyt6 := totalEncodedFields0zfyt6
	missingFieldsLeft0zfyt6 := maxFields0zfyt6 - totalEncodedFields0zfyt6

	var nextMiss0zfyt6 int32 = -1
	var found0zfyt6 [maxFields0zfyt6]bool
	var curField0zfyt6 string

doneWithStruct0zfyt6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zfyt6 > 0 || missingFieldsLeft0zfyt6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zfyt6, missingFieldsLeft0zfyt6, msgp.ShowFound(found0zfyt6[:]), decodeMsgFieldOrder0zfyt6)
		if encodedFieldsLeft0zfyt6 > 0 {
			encodedFieldsLeft0zfyt6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zfyt6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zfyt6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zfyt6 = 0
			}
			for nextMiss0zfyt6 < maxFields0zfyt6 && (found0zfyt6[nextMiss0zfyt6] || decodeMsgFieldSkip0zfyt6[nextMiss0zfyt6]) {
				nextMiss0zfyt6++
			}
			if nextMiss0zfyt6 == maxFields0zfyt6 {
				// filled all the empty fields!
				break doneWithStruct0zfyt6
			}
			missingFieldsLeft0zfyt6--
			curField0zfyt6 = decodeMsgFieldOrder0zfyt6[nextMiss0zfyt6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zfyt6)
		switch curField0zfyt6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zfyt6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zfyt6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zfyt6[3] = true
			var zjmt7 uint32
			zjmt7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zjmt7 > 0 {
				z.Snot = make(map[string]bool, zjmt7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zjmt7 > 0 {
				zjmt7--
				var zjmq1 string
				var zpyg2 bool
				zjmq1, err = dc.ReadString()
				if err != nil {
					return
				}
				zpyg2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zjmq1] = zpyg2
			}
		case "Notyet_zid03_map":
			found0zfyt6[4] = true
			var zezl8 uint32
			zezl8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zezl8 > 0 {
				z.Notyet = make(map[string]bool, zezl8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zezl8 > 0 {
				zezl8--
				var zvps3 string
				var zldt4 bool
				zvps3, err = dc.ReadString()
				if err != nil {
					return
				}
				zldt4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zvps3] = zldt4
			}
		case "Setm_zid04_slc":
			found0zfyt6[5] = true
			var zdsg9 uint32
			zdsg9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zdsg9) {
				z.Setm = (z.Setm)[:zdsg9]
			} else {
				z.Setm = make([]*inn, zdsg9)
			}
			for zfzj5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zfzj5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zfzj5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zfzj5] == nil {
						z.Setm[zfzj5] = new(inn)
					}
					err = z.Setm[zfzj5].DecodeMsg(dc)
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
	if nextMiss0zfyt6 != -1 {
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
var decodeMsgFieldOrder0zfyt6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zfyt6 = []bool{false, true, false, false, false, false}

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
	var empty_zlac10 [6]bool
	fieldsInUse_zgoz11 := z.fieldsNotEmpty(empty_zlac10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgoz11)
	if err != nil {
		return err
	}

	if !empty_zlac10[0] {
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

	if !empty_zlac10[2] {
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

	if !empty_zlac10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zjmq1, zpyg2 := range z.Snot {
			err = en.WriteString(zjmq1)
			if err != nil {
				return
			}
			err = en.WriteBool(zpyg2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zlac10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zvps3, zldt4 := range z.Notyet {
			err = en.WriteString(zvps3)
			if err != nil {
				return
			}
			err = en.WriteBool(zldt4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zlac10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zfzj5 := range z.Setm {
			if z.Setm[zfzj5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zfzj5].EncodeMsg(en)
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
		for zjmq1, zpyg2 := range z.Snot {
			o = msgp.AppendString(o, zjmq1)
			o = msgp.AppendBool(o, zpyg2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zvps3, zldt4 := range z.Notyet {
			o = msgp.AppendString(o, zvps3)
			o = msgp.AppendBool(o, zldt4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zfzj5 := range z.Setm {
			if z.Setm[zfzj5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zfzj5].MarshalMsg(o)
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
	const maxFields1znod12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1znod12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1znod12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1znod12 := totalEncodedFields1znod12
	missingFieldsLeft1znod12 := maxFields1znod12 - totalEncodedFields1znod12

	var nextMiss1znod12 int32 = -1
	var found1znod12 [maxFields1znod12]bool
	var curField1znod12 string

doneWithStruct1znod12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1znod12 > 0 || missingFieldsLeft1znod12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1znod12, missingFieldsLeft1znod12, msgp.ShowFound(found1znod12[:]), unmarshalMsgFieldOrder1znod12)
		if encodedFieldsLeft1znod12 > 0 {
			encodedFieldsLeft1znod12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1znod12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1znod12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1znod12 = 0
			}
			for nextMiss1znod12 < maxFields1znod12 && (found1znod12[nextMiss1znod12] || unmarshalMsgFieldSkip1znod12[nextMiss1znod12]) {
				nextMiss1znod12++
			}
			if nextMiss1znod12 == maxFields1znod12 {
				// filled all the empty fields!
				break doneWithStruct1znod12
			}
			missingFieldsLeft1znod12--
			curField1znod12 = unmarshalMsgFieldOrder1znod12[nextMiss1znod12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1znod12)
		switch curField1znod12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1znod12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1znod12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1znod12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zhfd13 uint32
				zhfd13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zhfd13 > 0 {
					z.Snot = make(map[string]bool, zhfd13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zhfd13 > 0 {
					var zjmq1 string
					var zpyg2 bool
					zhfd13--
					zjmq1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zpyg2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zjmq1] = zpyg2
				}
			}
		case "Notyet_zid03_map":
			found1znod12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var ziad14 uint32
				ziad14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && ziad14 > 0 {
					z.Notyet = make(map[string]bool, ziad14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for ziad14 > 0 {
					var zvps3 string
					var zldt4 bool
					ziad14--
					zvps3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zldt4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zvps3] = zldt4
				}
			}
		case "Setm_zid04_slc":
			found1znod12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zwpl15 uint32
				zwpl15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zwpl15) {
					z.Setm = (z.Setm)[:zwpl15]
				} else {
					z.Setm = make([]*inn, zwpl15)
				}
				for zfzj5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zfzj5] != nil {
							z.Setm[zfzj5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zfzj5] {
								z.Setm[zfzj5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zfzj5] == nil {
								z.Setm[zfzj5] = new(inn)
							}
							bts, err = z.Setm[zfzj5].UnmarshalMsg(bts)
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
	if nextMiss1znod12 != -1 {
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
var unmarshalMsgFieldOrder1znod12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1znod12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zjmq1, zpyg2 := range z.Snot {
			_ = zpyg2
			_ = zjmq1
			s += msgp.StringPrefixSize + len(zjmq1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zvps3, zldt4 := range z.Notyet {
			_ = zldt4
			_ = zvps3
			s += msgp.StringPrefixSize + len(zvps3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zfzj5 := range z.Setm {
		if z.Setm[zfzj5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zfzj5].Msgsize()
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
	const maxFields2zyqq18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zyqq18 uint32
	totalEncodedFields2zyqq18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zyqq18 := totalEncodedFields2zyqq18
	missingFieldsLeft2zyqq18 := maxFields2zyqq18 - totalEncodedFields2zyqq18

	var nextMiss2zyqq18 int32 = -1
	var found2zyqq18 [maxFields2zyqq18]bool
	var curField2zyqq18 string

doneWithStruct2zyqq18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zyqq18 > 0 || missingFieldsLeft2zyqq18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zyqq18, missingFieldsLeft2zyqq18, msgp.ShowFound(found2zyqq18[:]), decodeMsgFieldOrder2zyqq18)
		if encodedFieldsLeft2zyqq18 > 0 {
			encodedFieldsLeft2zyqq18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zyqq18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zyqq18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zyqq18 = 0
			}
			for nextMiss2zyqq18 < maxFields2zyqq18 && (found2zyqq18[nextMiss2zyqq18] || decodeMsgFieldSkip2zyqq18[nextMiss2zyqq18]) {
				nextMiss2zyqq18++
			}
			if nextMiss2zyqq18 == maxFields2zyqq18 {
				// filled all the empty fields!
				break doneWithStruct2zyqq18
			}
			missingFieldsLeft2zyqq18--
			curField2zyqq18 = decodeMsgFieldOrder2zyqq18[nextMiss2zyqq18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zyqq18)
		switch curField2zyqq18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zyqq18[0] = true
			var zqag19 uint32
			zqag19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zqag19 > 0 {
				z.j = make(map[string]bool, zqag19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zqag19 > 0 {
				zqag19--
				var zhkc16 string
				var zfqb17 bool
				zhkc16, err = dc.ReadString()
				if err != nil {
					return
				}
				zfqb17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zhkc16] = zfqb17
			}
		case "e_zid01_i64":
			found2zyqq18[1] = true
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
	if nextMiss2zyqq18 != -1 {
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
var decodeMsgFieldOrder2zyqq18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zyqq18 = []bool{false, false}

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
	var empty_zter20 [2]bool
	fieldsInUse_zqlw21 := z.fieldsNotEmpty(empty_zter20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqlw21)
	if err != nil {
		return err
	}

	if !empty_zter20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zhkc16, zfqb17 := range z.j {
			err = en.WriteString(zhkc16)
			if err != nil {
				return
			}
			err = en.WriteBool(zfqb17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zter20[1] {
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
		for zhkc16, zfqb17 := range z.j {
			o = msgp.AppendString(o, zhkc16)
			o = msgp.AppendBool(o, zfqb17)
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
	const maxFields3zpjo22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zpjo22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zpjo22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zpjo22 := totalEncodedFields3zpjo22
	missingFieldsLeft3zpjo22 := maxFields3zpjo22 - totalEncodedFields3zpjo22

	var nextMiss3zpjo22 int32 = -1
	var found3zpjo22 [maxFields3zpjo22]bool
	var curField3zpjo22 string

doneWithStruct3zpjo22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zpjo22 > 0 || missingFieldsLeft3zpjo22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zpjo22, missingFieldsLeft3zpjo22, msgp.ShowFound(found3zpjo22[:]), unmarshalMsgFieldOrder3zpjo22)
		if encodedFieldsLeft3zpjo22 > 0 {
			encodedFieldsLeft3zpjo22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zpjo22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zpjo22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zpjo22 = 0
			}
			for nextMiss3zpjo22 < maxFields3zpjo22 && (found3zpjo22[nextMiss3zpjo22] || unmarshalMsgFieldSkip3zpjo22[nextMiss3zpjo22]) {
				nextMiss3zpjo22++
			}
			if nextMiss3zpjo22 == maxFields3zpjo22 {
				// filled all the empty fields!
				break doneWithStruct3zpjo22
			}
			missingFieldsLeft3zpjo22--
			curField3zpjo22 = unmarshalMsgFieldOrder3zpjo22[nextMiss3zpjo22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zpjo22)
		switch curField3zpjo22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zpjo22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zils23 uint32
				zils23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zils23 > 0 {
					z.j = make(map[string]bool, zils23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zils23 > 0 {
					var zhkc16 string
					var zfqb17 bool
					zils23--
					zhkc16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zfqb17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zhkc16] = zfqb17
				}
			}
		case "e_zid01_i64":
			found3zpjo22[1] = true
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
	if nextMiss3zpjo22 != -1 {
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
var unmarshalMsgFieldOrder3zpjo22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zpjo22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zhkc16, zfqb17 := range z.j {
			_ = zfqb17
			_ = zhkc16
			s += msgp.StringPrefixSize + len(zhkc16) + msgp.BoolSize
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
	const maxFields4zbxz26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zbxz26 uint32
	totalEncodedFields4zbxz26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zbxz26 := totalEncodedFields4zbxz26
	missingFieldsLeft4zbxz26 := maxFields4zbxz26 - totalEncodedFields4zbxz26

	var nextMiss4zbxz26 int32 = -1
	var found4zbxz26 [maxFields4zbxz26]bool
	var curField4zbxz26 string

doneWithStruct4zbxz26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zbxz26 > 0 || missingFieldsLeft4zbxz26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zbxz26, missingFieldsLeft4zbxz26, msgp.ShowFound(found4zbxz26[:]), decodeMsgFieldOrder4zbxz26)
		if encodedFieldsLeft4zbxz26 > 0 {
			encodedFieldsLeft4zbxz26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zbxz26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zbxz26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zbxz26 = 0
			}
			for nextMiss4zbxz26 < maxFields4zbxz26 && (found4zbxz26[nextMiss4zbxz26] || decodeMsgFieldSkip4zbxz26[nextMiss4zbxz26]) {
				nextMiss4zbxz26++
			}
			if nextMiss4zbxz26 == maxFields4zbxz26 {
				// filled all the empty fields!
				break doneWithStruct4zbxz26
			}
			missingFieldsLeft4zbxz26--
			curField4zbxz26 = decodeMsgFieldOrder4zbxz26[nextMiss4zbxz26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zbxz26)
		switch curField4zbxz26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zbxz26[0] = true
			var zwpa27 uint32
			zwpa27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zwpa27 > 0 {
				z.m = make(map[string]*Tr, zwpa27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zwpa27 > 0 {
				zwpa27--
				var zodx24 string
				var zvru25 *Tr
				zodx24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zvru25 != nil {
						dc.PushAlwaysNil()
						err = zvru25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zvru25 == nil {
						zvru25 = new(Tr)
					}
					err = zvru25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zodx24] = zvru25
			}
		case "s_zid01_str":
			found4zbxz26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zbxz26[2] = true
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
	if nextMiss4zbxz26 != -1 {
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
var decodeMsgFieldOrder4zbxz26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zbxz26 = []bool{false, false, false}

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
	var empty_zowz28 [3]bool
	fieldsInUse_zixc29 := z.fieldsNotEmpty(empty_zowz28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zixc29)
	if err != nil {
		return err
	}

	if !empty_zowz28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zodx24, zvru25 := range z.m {
			err = en.WriteString(zodx24)
			if err != nil {
				return
			}
			if zvru25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zvru25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zowz28[1] {
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

	if !empty_zowz28[2] {
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
		for zodx24, zvru25 := range z.m {
			o = msgp.AppendString(o, zodx24)
			if zvru25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zvru25.MarshalMsg(o)
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
	const maxFields5zddn30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zddn30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zddn30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zddn30 := totalEncodedFields5zddn30
	missingFieldsLeft5zddn30 := maxFields5zddn30 - totalEncodedFields5zddn30

	var nextMiss5zddn30 int32 = -1
	var found5zddn30 [maxFields5zddn30]bool
	var curField5zddn30 string

doneWithStruct5zddn30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zddn30 > 0 || missingFieldsLeft5zddn30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zddn30, missingFieldsLeft5zddn30, msgp.ShowFound(found5zddn30[:]), unmarshalMsgFieldOrder5zddn30)
		if encodedFieldsLeft5zddn30 > 0 {
			encodedFieldsLeft5zddn30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zddn30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zddn30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zddn30 = 0
			}
			for nextMiss5zddn30 < maxFields5zddn30 && (found5zddn30[nextMiss5zddn30] || unmarshalMsgFieldSkip5zddn30[nextMiss5zddn30]) {
				nextMiss5zddn30++
			}
			if nextMiss5zddn30 == maxFields5zddn30 {
				// filled all the empty fields!
				break doneWithStruct5zddn30
			}
			missingFieldsLeft5zddn30--
			curField5zddn30 = unmarshalMsgFieldOrder5zddn30[nextMiss5zddn30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zddn30)
		switch curField5zddn30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zddn30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zzfd31 uint32
				zzfd31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zzfd31 > 0 {
					z.m = make(map[string]*Tr, zzfd31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zzfd31 > 0 {
					var zodx24 string
					var zvru25 *Tr
					zzfd31--
					zodx24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zvru25 != nil {
							zvru25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zvru25 {
								zvru25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zvru25 == nil {
								zvru25 = new(Tr)
							}
							bts, err = zvru25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zodx24] = zvru25
				}
			}
		case "s_zid01_str":
			found5zddn30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zddn30[2] = true
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
	if nextMiss5zddn30 != -1 {
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
var unmarshalMsgFieldOrder5zddn30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zddn30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zodx24, zvru25 := range z.m {
			_ = zvru25
			_ = zodx24
			s += msgp.StringPrefixSize + len(zodx24)
			if zvru25 == nil {
				s += msgp.NilSize
			} else {
				s += zvru25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
