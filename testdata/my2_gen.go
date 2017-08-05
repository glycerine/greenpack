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
	const maxFields0znuv6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0znuv6 uint32
	totalEncodedFields0znuv6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0znuv6 := totalEncodedFields0znuv6
	missingFieldsLeft0znuv6 := maxFields0znuv6 - totalEncodedFields0znuv6

	var nextMiss0znuv6 int32 = -1
	var found0znuv6 [maxFields0znuv6]bool
	var curField0znuv6 string

doneWithStruct0znuv6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0znuv6 > 0 || missingFieldsLeft0znuv6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0znuv6, missingFieldsLeft0znuv6, msgp.ShowFound(found0znuv6[:]), decodeMsgFieldOrder0znuv6)
		if encodedFieldsLeft0znuv6 > 0 {
			encodedFieldsLeft0znuv6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0znuv6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0znuv6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0znuv6 = 0
			}
			for nextMiss0znuv6 < maxFields0znuv6 && (found0znuv6[nextMiss0znuv6] || decodeMsgFieldSkip0znuv6[nextMiss0znuv6]) {
				nextMiss0znuv6++
			}
			if nextMiss0znuv6 == maxFields0znuv6 {
				// filled all the empty fields!
				break doneWithStruct0znuv6
			}
			missingFieldsLeft0znuv6--
			curField0znuv6 = decodeMsgFieldOrder0znuv6[nextMiss0znuv6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0znuv6)
		switch curField0znuv6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0znuv6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0znuv6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0znuv6[3] = true
			var zbwa7 uint32
			zbwa7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zbwa7 > 0 {
				z.Snot = make(map[string]bool, zbwa7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zbwa7 > 0 {
				zbwa7--
				var zwxo1 string
				var zotw2 bool
				zwxo1, err = dc.ReadString()
				if err != nil {
					return
				}
				zotw2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zwxo1] = zotw2
			}
		case "Notyet_zid03_map":
			found0znuv6[4] = true
			var zujq8 uint32
			zujq8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zujq8 > 0 {
				z.Notyet = make(map[string]bool, zujq8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zujq8 > 0 {
				zujq8--
				var zevc3 string
				var zhsg4 bool
				zevc3, err = dc.ReadString()
				if err != nil {
					return
				}
				zhsg4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zevc3] = zhsg4
			}
		case "Setm_zid04_slc":
			found0znuv6[5] = true
			var zyej9 uint32
			zyej9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zyej9) {
				z.Setm = (z.Setm)[:zyej9]
			} else {
				z.Setm = make([]*inn, zyej9)
			}
			for zjvu5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zjvu5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zjvu5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zjvu5] == nil {
						z.Setm[zjvu5] = new(inn)
					}
					err = z.Setm[zjvu5].DecodeMsg(dc)
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
	if nextMiss0znuv6 != -1 {
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
var decodeMsgFieldOrder0znuv6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0znuv6 = []bool{false, true, false, false, false, false}

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
	var empty_zyde10 [6]bool
	fieldsInUse_zdai11 := z.fieldsNotEmpty(empty_zyde10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zdai11)
	if err != nil {
		return err
	}

	if !empty_zyde10[0] {
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

	if !empty_zyde10[2] {
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

	if !empty_zyde10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zwxo1, zotw2 := range z.Snot {
			err = en.WriteString(zwxo1)
			if err != nil {
				return
			}
			err = en.WriteBool(zotw2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zyde10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zevc3, zhsg4 := range z.Notyet {
			err = en.WriteString(zevc3)
			if err != nil {
				return
			}
			err = en.WriteBool(zhsg4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zyde10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zjvu5 := range z.Setm {
			if z.Setm[zjvu5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zjvu5].EncodeMsg(en)
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
		for zwxo1, zotw2 := range z.Snot {
			o = msgp.AppendString(o, zwxo1)
			o = msgp.AppendBool(o, zotw2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zevc3, zhsg4 := range z.Notyet {
			o = msgp.AppendString(o, zevc3)
			o = msgp.AppendBool(o, zhsg4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zjvu5 := range z.Setm {
			if z.Setm[zjvu5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zjvu5].MarshalMsg(o)
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
	const maxFields1zemm12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zemm12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zemm12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zemm12 := totalEncodedFields1zemm12
	missingFieldsLeft1zemm12 := maxFields1zemm12 - totalEncodedFields1zemm12

	var nextMiss1zemm12 int32 = -1
	var found1zemm12 [maxFields1zemm12]bool
	var curField1zemm12 string

doneWithStruct1zemm12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zemm12 > 0 || missingFieldsLeft1zemm12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zemm12, missingFieldsLeft1zemm12, msgp.ShowFound(found1zemm12[:]), unmarshalMsgFieldOrder1zemm12)
		if encodedFieldsLeft1zemm12 > 0 {
			encodedFieldsLeft1zemm12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zemm12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zemm12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zemm12 = 0
			}
			for nextMiss1zemm12 < maxFields1zemm12 && (found1zemm12[nextMiss1zemm12] || unmarshalMsgFieldSkip1zemm12[nextMiss1zemm12]) {
				nextMiss1zemm12++
			}
			if nextMiss1zemm12 == maxFields1zemm12 {
				// filled all the empty fields!
				break doneWithStruct1zemm12
			}
			missingFieldsLeft1zemm12--
			curField1zemm12 = unmarshalMsgFieldOrder1zemm12[nextMiss1zemm12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zemm12)
		switch curField1zemm12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zemm12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zemm12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zemm12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zwtk13 uint32
				zwtk13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zwtk13 > 0 {
					z.Snot = make(map[string]bool, zwtk13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zwtk13 > 0 {
					var zwxo1 string
					var zotw2 bool
					zwtk13--
					zwxo1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zotw2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zwxo1] = zotw2
				}
			}
		case "Notyet_zid03_map":
			found1zemm12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zftv14 uint32
				zftv14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zftv14 > 0 {
					z.Notyet = make(map[string]bool, zftv14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zftv14 > 0 {
					var zevc3 string
					var zhsg4 bool
					zftv14--
					zevc3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zhsg4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zevc3] = zhsg4
				}
			}
		case "Setm_zid04_slc":
			found1zemm12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zcpb15 uint32
				zcpb15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zcpb15) {
					z.Setm = (z.Setm)[:zcpb15]
				} else {
					z.Setm = make([]*inn, zcpb15)
				}
				for zjvu5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zjvu5] != nil {
							z.Setm[zjvu5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zjvu5] {
								z.Setm[zjvu5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zjvu5] == nil {
								z.Setm[zjvu5] = new(inn)
							}
							bts, err = z.Setm[zjvu5].UnmarshalMsg(bts)
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
	if nextMiss1zemm12 != -1 {
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
var unmarshalMsgFieldOrder1zemm12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zemm12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zwxo1, zotw2 := range z.Snot {
			_ = zotw2
			_ = zwxo1
			s += msgp.StringPrefixSize + len(zwxo1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zevc3, zhsg4 := range z.Notyet {
			_ = zhsg4
			_ = zevc3
			s += msgp.StringPrefixSize + len(zevc3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zjvu5 := range z.Setm {
		if z.Setm[zjvu5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zjvu5].Msgsize()
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
	const maxFields2zldk18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zldk18 uint32
	totalEncodedFields2zldk18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zldk18 := totalEncodedFields2zldk18
	missingFieldsLeft2zldk18 := maxFields2zldk18 - totalEncodedFields2zldk18

	var nextMiss2zldk18 int32 = -1
	var found2zldk18 [maxFields2zldk18]bool
	var curField2zldk18 string

doneWithStruct2zldk18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zldk18 > 0 || missingFieldsLeft2zldk18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zldk18, missingFieldsLeft2zldk18, msgp.ShowFound(found2zldk18[:]), decodeMsgFieldOrder2zldk18)
		if encodedFieldsLeft2zldk18 > 0 {
			encodedFieldsLeft2zldk18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zldk18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zldk18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zldk18 = 0
			}
			for nextMiss2zldk18 < maxFields2zldk18 && (found2zldk18[nextMiss2zldk18] || decodeMsgFieldSkip2zldk18[nextMiss2zldk18]) {
				nextMiss2zldk18++
			}
			if nextMiss2zldk18 == maxFields2zldk18 {
				// filled all the empty fields!
				break doneWithStruct2zldk18
			}
			missingFieldsLeft2zldk18--
			curField2zldk18 = decodeMsgFieldOrder2zldk18[nextMiss2zldk18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zldk18)
		switch curField2zldk18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zldk18[0] = true
			var zakr19 uint32
			zakr19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zakr19 > 0 {
				z.j = make(map[string]bool, zakr19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zakr19 > 0 {
				zakr19--
				var zlbf16 string
				var zlnt17 bool
				zlbf16, err = dc.ReadString()
				if err != nil {
					return
				}
				zlnt17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zlbf16] = zlnt17
			}
		case "e_zid01_i64":
			found2zldk18[1] = true
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
	if nextMiss2zldk18 != -1 {
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
var decodeMsgFieldOrder2zldk18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zldk18 = []bool{false, false}

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
	var empty_zbpi20 [2]bool
	fieldsInUse_zsvt21 := z.fieldsNotEmpty(empty_zbpi20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zsvt21)
	if err != nil {
		return err
	}

	if !empty_zbpi20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zlbf16, zlnt17 := range z.j {
			err = en.WriteString(zlbf16)
			if err != nil {
				return
			}
			err = en.WriteBool(zlnt17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zbpi20[1] {
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
		for zlbf16, zlnt17 := range z.j {
			o = msgp.AppendString(o, zlbf16)
			o = msgp.AppendBool(o, zlnt17)
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
	const maxFields3zpkw22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zpkw22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zpkw22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zpkw22 := totalEncodedFields3zpkw22
	missingFieldsLeft3zpkw22 := maxFields3zpkw22 - totalEncodedFields3zpkw22

	var nextMiss3zpkw22 int32 = -1
	var found3zpkw22 [maxFields3zpkw22]bool
	var curField3zpkw22 string

doneWithStruct3zpkw22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zpkw22 > 0 || missingFieldsLeft3zpkw22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zpkw22, missingFieldsLeft3zpkw22, msgp.ShowFound(found3zpkw22[:]), unmarshalMsgFieldOrder3zpkw22)
		if encodedFieldsLeft3zpkw22 > 0 {
			encodedFieldsLeft3zpkw22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zpkw22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zpkw22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zpkw22 = 0
			}
			for nextMiss3zpkw22 < maxFields3zpkw22 && (found3zpkw22[nextMiss3zpkw22] || unmarshalMsgFieldSkip3zpkw22[nextMiss3zpkw22]) {
				nextMiss3zpkw22++
			}
			if nextMiss3zpkw22 == maxFields3zpkw22 {
				// filled all the empty fields!
				break doneWithStruct3zpkw22
			}
			missingFieldsLeft3zpkw22--
			curField3zpkw22 = unmarshalMsgFieldOrder3zpkw22[nextMiss3zpkw22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zpkw22)
		switch curField3zpkw22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zpkw22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zpfc23 uint32
				zpfc23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zpfc23 > 0 {
					z.j = make(map[string]bool, zpfc23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zpfc23 > 0 {
					var zlbf16 string
					var zlnt17 bool
					zpfc23--
					zlbf16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zlnt17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zlbf16] = zlnt17
				}
			}
		case "e_zid01_i64":
			found3zpkw22[1] = true
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
	if nextMiss3zpkw22 != -1 {
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
var unmarshalMsgFieldOrder3zpkw22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zpkw22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zlbf16, zlnt17 := range z.j {
			_ = zlnt17
			_ = zlbf16
			s += msgp.StringPrefixSize + len(zlbf16) + msgp.BoolSize
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
	const maxFields4zuur26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zuur26 uint32
	totalEncodedFields4zuur26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zuur26 := totalEncodedFields4zuur26
	missingFieldsLeft4zuur26 := maxFields4zuur26 - totalEncodedFields4zuur26

	var nextMiss4zuur26 int32 = -1
	var found4zuur26 [maxFields4zuur26]bool
	var curField4zuur26 string

doneWithStruct4zuur26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zuur26 > 0 || missingFieldsLeft4zuur26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zuur26, missingFieldsLeft4zuur26, msgp.ShowFound(found4zuur26[:]), decodeMsgFieldOrder4zuur26)
		if encodedFieldsLeft4zuur26 > 0 {
			encodedFieldsLeft4zuur26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zuur26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zuur26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zuur26 = 0
			}
			for nextMiss4zuur26 < maxFields4zuur26 && (found4zuur26[nextMiss4zuur26] || decodeMsgFieldSkip4zuur26[nextMiss4zuur26]) {
				nextMiss4zuur26++
			}
			if nextMiss4zuur26 == maxFields4zuur26 {
				// filled all the empty fields!
				break doneWithStruct4zuur26
			}
			missingFieldsLeft4zuur26--
			curField4zuur26 = decodeMsgFieldOrder4zuur26[nextMiss4zuur26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zuur26)
		switch curField4zuur26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zuur26[0] = true
			var zokq27 uint32
			zokq27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zokq27 > 0 {
				z.m = make(map[string]*Tr, zokq27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zokq27 > 0 {
				zokq27--
				var zgdh24 string
				var zygs25 *Tr
				zgdh24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zygs25 != nil {
						dc.PushAlwaysNil()
						err = zygs25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zygs25 == nil {
						zygs25 = new(Tr)
					}
					err = zygs25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zgdh24] = zygs25
			}
		case "s_zid01_str":
			found4zuur26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zuur26[2] = true
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
	if nextMiss4zuur26 != -1 {
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
var decodeMsgFieldOrder4zuur26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zuur26 = []bool{false, false, false}

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
	var empty_zwbe28 [3]bool
	fieldsInUse_zizv29 := z.fieldsNotEmpty(empty_zwbe28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zizv29)
	if err != nil {
		return err
	}

	if !empty_zwbe28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zgdh24, zygs25 := range z.m {
			err = en.WriteString(zgdh24)
			if err != nil {
				return
			}
			if zygs25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zygs25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zwbe28[1] {
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

	if !empty_zwbe28[2] {
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
		for zgdh24, zygs25 := range z.m {
			o = msgp.AppendString(o, zgdh24)
			if zygs25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zygs25.MarshalMsg(o)
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
	const maxFields5zhao30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zhao30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zhao30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zhao30 := totalEncodedFields5zhao30
	missingFieldsLeft5zhao30 := maxFields5zhao30 - totalEncodedFields5zhao30

	var nextMiss5zhao30 int32 = -1
	var found5zhao30 [maxFields5zhao30]bool
	var curField5zhao30 string

doneWithStruct5zhao30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zhao30 > 0 || missingFieldsLeft5zhao30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zhao30, missingFieldsLeft5zhao30, msgp.ShowFound(found5zhao30[:]), unmarshalMsgFieldOrder5zhao30)
		if encodedFieldsLeft5zhao30 > 0 {
			encodedFieldsLeft5zhao30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zhao30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zhao30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zhao30 = 0
			}
			for nextMiss5zhao30 < maxFields5zhao30 && (found5zhao30[nextMiss5zhao30] || unmarshalMsgFieldSkip5zhao30[nextMiss5zhao30]) {
				nextMiss5zhao30++
			}
			if nextMiss5zhao30 == maxFields5zhao30 {
				// filled all the empty fields!
				break doneWithStruct5zhao30
			}
			missingFieldsLeft5zhao30--
			curField5zhao30 = unmarshalMsgFieldOrder5zhao30[nextMiss5zhao30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zhao30)
		switch curField5zhao30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zhao30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zgkt31 uint32
				zgkt31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zgkt31 > 0 {
					z.m = make(map[string]*Tr, zgkt31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zgkt31 > 0 {
					var zgdh24 string
					var zygs25 *Tr
					zgkt31--
					zgdh24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zygs25 != nil {
							zygs25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zygs25 {
								zygs25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zygs25 == nil {
								zygs25 = new(Tr)
							}
							bts, err = zygs25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zgdh24] = zygs25
				}
			}
		case "s_zid01_str":
			found5zhao30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zhao30[2] = true
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
	if nextMiss5zhao30 != -1 {
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
var unmarshalMsgFieldOrder5zhao30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zhao30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zgdh24, zygs25 := range z.m {
			_ = zygs25
			_ = zgdh24
			s += msgp.StringPrefixSize + len(zgdh24)
			if zygs25 == nil {
				s += msgp.NilSize
			} else {
				s += zygs25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
