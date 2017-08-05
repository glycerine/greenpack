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
	const maxFields0zjyh6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zjyh6 uint32
	totalEncodedFields0zjyh6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zjyh6 := totalEncodedFields0zjyh6
	missingFieldsLeft0zjyh6 := maxFields0zjyh6 - totalEncodedFields0zjyh6

	var nextMiss0zjyh6 int32 = -1
	var found0zjyh6 [maxFields0zjyh6]bool
	var curField0zjyh6 string

doneWithStruct0zjyh6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zjyh6 > 0 || missingFieldsLeft0zjyh6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zjyh6, missingFieldsLeft0zjyh6, msgp.ShowFound(found0zjyh6[:]), decodeMsgFieldOrder0zjyh6)
		if encodedFieldsLeft0zjyh6 > 0 {
			encodedFieldsLeft0zjyh6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zjyh6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zjyh6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zjyh6 = 0
			}
			for nextMiss0zjyh6 < maxFields0zjyh6 && (found0zjyh6[nextMiss0zjyh6] || decodeMsgFieldSkip0zjyh6[nextMiss0zjyh6]) {
				nextMiss0zjyh6++
			}
			if nextMiss0zjyh6 == maxFields0zjyh6 {
				// filled all the empty fields!
				break doneWithStruct0zjyh6
			}
			missingFieldsLeft0zjyh6--
			curField0zjyh6 = decodeMsgFieldOrder0zjyh6[nextMiss0zjyh6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zjyh6)
		switch curField0zjyh6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zjyh6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zjyh6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zjyh6[3] = true
			var zcde7 uint32
			zcde7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zcde7 > 0 {
				z.Snot = make(map[string]bool, zcde7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zcde7 > 0 {
				zcde7--
				var zogc1 string
				var zuuy2 bool
				zogc1, err = dc.ReadString()
				if err != nil {
					return
				}
				zuuy2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zogc1] = zuuy2
			}
		case "Notyet_zid03_map":
			found0zjyh6[4] = true
			var ztxt8 uint32
			ztxt8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && ztxt8 > 0 {
				z.Notyet = make(map[string]bool, ztxt8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for ztxt8 > 0 {
				ztxt8--
				var zgxp3 string
				var zmbn4 bool
				zgxp3, err = dc.ReadString()
				if err != nil {
					return
				}
				zmbn4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zgxp3] = zmbn4
			}
		case "Setm_zid04_slc":
			found0zjyh6[5] = true
			var zsfm9 uint32
			zsfm9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zsfm9) {
				z.Setm = (z.Setm)[:zsfm9]
			} else {
				z.Setm = make([]*inn, zsfm9)
			}
			for zngp5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zngp5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zngp5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zngp5] == nil {
						z.Setm[zngp5] = new(inn)
					}
					err = z.Setm[zngp5].DecodeMsg(dc)
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
	if nextMiss0zjyh6 != -1 {
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
var decodeMsgFieldOrder0zjyh6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zjyh6 = []bool{false, true, false, false, false, false}

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
	var empty_zuvz10 [6]bool
	fieldsInUse_zmlr11 := z.fieldsNotEmpty(empty_zuvz10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zmlr11)
	if err != nil {
		return err
	}

	if !empty_zuvz10[0] {
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

	if !empty_zuvz10[2] {
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

	if !empty_zuvz10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zogc1, zuuy2 := range z.Snot {
			err = en.WriteString(zogc1)
			if err != nil {
				return
			}
			err = en.WriteBool(zuuy2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zuvz10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zgxp3, zmbn4 := range z.Notyet {
			err = en.WriteString(zgxp3)
			if err != nil {
				return
			}
			err = en.WriteBool(zmbn4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zuvz10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zngp5 := range z.Setm {
			if z.Setm[zngp5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zngp5].EncodeMsg(en)
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
		for zogc1, zuuy2 := range z.Snot {
			o = msgp.AppendString(o, zogc1)
			o = msgp.AppendBool(o, zuuy2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zgxp3, zmbn4 := range z.Notyet {
			o = msgp.AppendString(o, zgxp3)
			o = msgp.AppendBool(o, zmbn4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zngp5 := range z.Setm {
			if z.Setm[zngp5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zngp5].MarshalMsg(o)
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
	const maxFields1zyar12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zyar12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zyar12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zyar12 := totalEncodedFields1zyar12
	missingFieldsLeft1zyar12 := maxFields1zyar12 - totalEncodedFields1zyar12

	var nextMiss1zyar12 int32 = -1
	var found1zyar12 [maxFields1zyar12]bool
	var curField1zyar12 string

doneWithStruct1zyar12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zyar12 > 0 || missingFieldsLeft1zyar12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zyar12, missingFieldsLeft1zyar12, msgp.ShowFound(found1zyar12[:]), unmarshalMsgFieldOrder1zyar12)
		if encodedFieldsLeft1zyar12 > 0 {
			encodedFieldsLeft1zyar12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zyar12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zyar12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zyar12 = 0
			}
			for nextMiss1zyar12 < maxFields1zyar12 && (found1zyar12[nextMiss1zyar12] || unmarshalMsgFieldSkip1zyar12[nextMiss1zyar12]) {
				nextMiss1zyar12++
			}
			if nextMiss1zyar12 == maxFields1zyar12 {
				// filled all the empty fields!
				break doneWithStruct1zyar12
			}
			missingFieldsLeft1zyar12--
			curField1zyar12 = unmarshalMsgFieldOrder1zyar12[nextMiss1zyar12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zyar12)
		switch curField1zyar12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zyar12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zyar12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zyar12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var ziae13 uint32
				ziae13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && ziae13 > 0 {
					z.Snot = make(map[string]bool, ziae13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for ziae13 > 0 {
					var zogc1 string
					var zuuy2 bool
					ziae13--
					zogc1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zuuy2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zogc1] = zuuy2
				}
			}
		case "Notyet_zid03_map":
			found1zyar12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zvmm14 uint32
				zvmm14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zvmm14 > 0 {
					z.Notyet = make(map[string]bool, zvmm14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zvmm14 > 0 {
					var zgxp3 string
					var zmbn4 bool
					zvmm14--
					zgxp3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zmbn4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zgxp3] = zmbn4
				}
			}
		case "Setm_zid04_slc":
			found1zyar12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zxvi15 uint32
				zxvi15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zxvi15) {
					z.Setm = (z.Setm)[:zxvi15]
				} else {
					z.Setm = make([]*inn, zxvi15)
				}
				for zngp5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zngp5] != nil {
							z.Setm[zngp5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zngp5] {
								z.Setm[zngp5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zngp5] == nil {
								z.Setm[zngp5] = new(inn)
							}
							bts, err = z.Setm[zngp5].UnmarshalMsg(bts)
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
	if nextMiss1zyar12 != -1 {
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
var unmarshalMsgFieldOrder1zyar12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zyar12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zogc1, zuuy2 := range z.Snot {
			_ = zuuy2
			_ = zogc1
			s += msgp.StringPrefixSize + len(zogc1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zgxp3, zmbn4 := range z.Notyet {
			_ = zmbn4
			_ = zgxp3
			s += msgp.StringPrefixSize + len(zgxp3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zngp5 := range z.Setm {
		if z.Setm[zngp5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zngp5].Msgsize()
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
	const maxFields2zxva18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zxva18 uint32
	totalEncodedFields2zxva18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zxva18 := totalEncodedFields2zxva18
	missingFieldsLeft2zxva18 := maxFields2zxva18 - totalEncodedFields2zxva18

	var nextMiss2zxva18 int32 = -1
	var found2zxva18 [maxFields2zxva18]bool
	var curField2zxva18 string

doneWithStruct2zxva18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zxva18 > 0 || missingFieldsLeft2zxva18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zxva18, missingFieldsLeft2zxva18, msgp.ShowFound(found2zxva18[:]), decodeMsgFieldOrder2zxva18)
		if encodedFieldsLeft2zxva18 > 0 {
			encodedFieldsLeft2zxva18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zxva18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zxva18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zxva18 = 0
			}
			for nextMiss2zxva18 < maxFields2zxva18 && (found2zxva18[nextMiss2zxva18] || decodeMsgFieldSkip2zxva18[nextMiss2zxva18]) {
				nextMiss2zxva18++
			}
			if nextMiss2zxva18 == maxFields2zxva18 {
				// filled all the empty fields!
				break doneWithStruct2zxva18
			}
			missingFieldsLeft2zxva18--
			curField2zxva18 = decodeMsgFieldOrder2zxva18[nextMiss2zxva18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zxva18)
		switch curField2zxva18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zxva18[0] = true
			var zdyc19 uint32
			zdyc19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zdyc19 > 0 {
				z.j = make(map[string]bool, zdyc19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zdyc19 > 0 {
				zdyc19--
				var zaoa16 string
				var zwgz17 bool
				zaoa16, err = dc.ReadString()
				if err != nil {
					return
				}
				zwgz17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zaoa16] = zwgz17
			}
		case "e_zid01_i64":
			found2zxva18[1] = true
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
	if nextMiss2zxva18 != -1 {
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
var decodeMsgFieldOrder2zxva18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zxva18 = []bool{false, false}

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
	var empty_zrpw20 [2]bool
	fieldsInUse_zppi21 := z.fieldsNotEmpty(empty_zrpw20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zppi21)
	if err != nil {
		return err
	}

	if !empty_zrpw20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zaoa16, zwgz17 := range z.j {
			err = en.WriteString(zaoa16)
			if err != nil {
				return
			}
			err = en.WriteBool(zwgz17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zrpw20[1] {
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
		for zaoa16, zwgz17 := range z.j {
			o = msgp.AppendString(o, zaoa16)
			o = msgp.AppendBool(o, zwgz17)
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
	const maxFields3zxgm22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zxgm22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zxgm22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zxgm22 := totalEncodedFields3zxgm22
	missingFieldsLeft3zxgm22 := maxFields3zxgm22 - totalEncodedFields3zxgm22

	var nextMiss3zxgm22 int32 = -1
	var found3zxgm22 [maxFields3zxgm22]bool
	var curField3zxgm22 string

doneWithStruct3zxgm22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zxgm22 > 0 || missingFieldsLeft3zxgm22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zxgm22, missingFieldsLeft3zxgm22, msgp.ShowFound(found3zxgm22[:]), unmarshalMsgFieldOrder3zxgm22)
		if encodedFieldsLeft3zxgm22 > 0 {
			encodedFieldsLeft3zxgm22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zxgm22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zxgm22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zxgm22 = 0
			}
			for nextMiss3zxgm22 < maxFields3zxgm22 && (found3zxgm22[nextMiss3zxgm22] || unmarshalMsgFieldSkip3zxgm22[nextMiss3zxgm22]) {
				nextMiss3zxgm22++
			}
			if nextMiss3zxgm22 == maxFields3zxgm22 {
				// filled all the empty fields!
				break doneWithStruct3zxgm22
			}
			missingFieldsLeft3zxgm22--
			curField3zxgm22 = unmarshalMsgFieldOrder3zxgm22[nextMiss3zxgm22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zxgm22)
		switch curField3zxgm22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zxgm22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zmbi23 uint32
				zmbi23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zmbi23 > 0 {
					z.j = make(map[string]bool, zmbi23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zmbi23 > 0 {
					var zaoa16 string
					var zwgz17 bool
					zmbi23--
					zaoa16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zwgz17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zaoa16] = zwgz17
				}
			}
		case "e_zid01_i64":
			found3zxgm22[1] = true
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
	if nextMiss3zxgm22 != -1 {
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
var unmarshalMsgFieldOrder3zxgm22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zxgm22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zaoa16, zwgz17 := range z.j {
			_ = zwgz17
			_ = zaoa16
			s += msgp.StringPrefixSize + len(zaoa16) + msgp.BoolSize
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
	const maxFields4zdxx26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zdxx26 uint32
	totalEncodedFields4zdxx26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zdxx26 := totalEncodedFields4zdxx26
	missingFieldsLeft4zdxx26 := maxFields4zdxx26 - totalEncodedFields4zdxx26

	var nextMiss4zdxx26 int32 = -1
	var found4zdxx26 [maxFields4zdxx26]bool
	var curField4zdxx26 string

doneWithStruct4zdxx26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zdxx26 > 0 || missingFieldsLeft4zdxx26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zdxx26, missingFieldsLeft4zdxx26, msgp.ShowFound(found4zdxx26[:]), decodeMsgFieldOrder4zdxx26)
		if encodedFieldsLeft4zdxx26 > 0 {
			encodedFieldsLeft4zdxx26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zdxx26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zdxx26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zdxx26 = 0
			}
			for nextMiss4zdxx26 < maxFields4zdxx26 && (found4zdxx26[nextMiss4zdxx26] || decodeMsgFieldSkip4zdxx26[nextMiss4zdxx26]) {
				nextMiss4zdxx26++
			}
			if nextMiss4zdxx26 == maxFields4zdxx26 {
				// filled all the empty fields!
				break doneWithStruct4zdxx26
			}
			missingFieldsLeft4zdxx26--
			curField4zdxx26 = decodeMsgFieldOrder4zdxx26[nextMiss4zdxx26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zdxx26)
		switch curField4zdxx26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zdxx26[0] = true
			var zrtd27 uint32
			zrtd27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zrtd27 > 0 {
				z.m = make(map[string]*Tr, zrtd27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zrtd27 > 0 {
				zrtd27--
				var zniu24 string
				var ztlu25 *Tr
				zniu24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if ztlu25 != nil {
						dc.PushAlwaysNil()
						err = ztlu25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if ztlu25 == nil {
						ztlu25 = new(Tr)
					}
					err = ztlu25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zniu24] = ztlu25
			}
		case "s_zid01_str":
			found4zdxx26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zdxx26[2] = true
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
	if nextMiss4zdxx26 != -1 {
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
var decodeMsgFieldOrder4zdxx26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zdxx26 = []bool{false, false, false}

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
	var empty_zkrp28 [3]bool
	fieldsInUse_zdkm29 := z.fieldsNotEmpty(empty_zkrp28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zdkm29)
	if err != nil {
		return err
	}

	if !empty_zkrp28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zniu24, ztlu25 := range z.m {
			err = en.WriteString(zniu24)
			if err != nil {
				return
			}
			if ztlu25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = ztlu25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zkrp28[1] {
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

	if !empty_zkrp28[2] {
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
		for zniu24, ztlu25 := range z.m {
			o = msgp.AppendString(o, zniu24)
			if ztlu25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = ztlu25.MarshalMsg(o)
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
	const maxFields5zyjh30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zyjh30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zyjh30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zyjh30 := totalEncodedFields5zyjh30
	missingFieldsLeft5zyjh30 := maxFields5zyjh30 - totalEncodedFields5zyjh30

	var nextMiss5zyjh30 int32 = -1
	var found5zyjh30 [maxFields5zyjh30]bool
	var curField5zyjh30 string

doneWithStruct5zyjh30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zyjh30 > 0 || missingFieldsLeft5zyjh30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zyjh30, missingFieldsLeft5zyjh30, msgp.ShowFound(found5zyjh30[:]), unmarshalMsgFieldOrder5zyjh30)
		if encodedFieldsLeft5zyjh30 > 0 {
			encodedFieldsLeft5zyjh30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zyjh30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zyjh30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zyjh30 = 0
			}
			for nextMiss5zyjh30 < maxFields5zyjh30 && (found5zyjh30[nextMiss5zyjh30] || unmarshalMsgFieldSkip5zyjh30[nextMiss5zyjh30]) {
				nextMiss5zyjh30++
			}
			if nextMiss5zyjh30 == maxFields5zyjh30 {
				// filled all the empty fields!
				break doneWithStruct5zyjh30
			}
			missingFieldsLeft5zyjh30--
			curField5zyjh30 = unmarshalMsgFieldOrder5zyjh30[nextMiss5zyjh30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zyjh30)
		switch curField5zyjh30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zyjh30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zebf31 uint32
				zebf31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zebf31 > 0 {
					z.m = make(map[string]*Tr, zebf31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zebf31 > 0 {
					var zniu24 string
					var ztlu25 *Tr
					zebf31--
					zniu24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if ztlu25 != nil {
							ztlu25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != ztlu25 {
								ztlu25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if ztlu25 == nil {
								ztlu25 = new(Tr)
							}
							bts, err = ztlu25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zniu24] = ztlu25
				}
			}
		case "s_zid01_str":
			found5zyjh30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zyjh30[2] = true
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
	if nextMiss5zyjh30 != -1 {
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
var unmarshalMsgFieldOrder5zyjh30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zyjh30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zniu24, ztlu25 := range z.m {
			_ = ztlu25
			_ = zniu24
			s += msgp.StringPrefixSize + len(zniu24)
			if ztlu25 == nil {
				s += msgp.NilSize
			} else {
				s += ztlu25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
