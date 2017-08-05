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
	const maxFields0zdgq6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zdgq6 uint32
	totalEncodedFields0zdgq6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zdgq6 := totalEncodedFields0zdgq6
	missingFieldsLeft0zdgq6 := maxFields0zdgq6 - totalEncodedFields0zdgq6

	var nextMiss0zdgq6 int32 = -1
	var found0zdgq6 [maxFields0zdgq6]bool
	var curField0zdgq6 string

doneWithStruct0zdgq6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zdgq6 > 0 || missingFieldsLeft0zdgq6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zdgq6, missingFieldsLeft0zdgq6, msgp.ShowFound(found0zdgq6[:]), decodeMsgFieldOrder0zdgq6)
		if encodedFieldsLeft0zdgq6 > 0 {
			encodedFieldsLeft0zdgq6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zdgq6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zdgq6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zdgq6 = 0
			}
			for nextMiss0zdgq6 < maxFields0zdgq6 && (found0zdgq6[nextMiss0zdgq6] || decodeMsgFieldSkip0zdgq6[nextMiss0zdgq6]) {
				nextMiss0zdgq6++
			}
			if nextMiss0zdgq6 == maxFields0zdgq6 {
				// filled all the empty fields!
				break doneWithStruct0zdgq6
			}
			missingFieldsLeft0zdgq6--
			curField0zdgq6 = decodeMsgFieldOrder0zdgq6[nextMiss0zdgq6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zdgq6)
		switch curField0zdgq6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zdgq6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zdgq6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zdgq6[3] = true
			var zjkl7 uint32
			zjkl7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zjkl7 > 0 {
				z.Snot = make(map[string]bool, zjkl7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zjkl7 > 0 {
				zjkl7--
				var zpfn1 string
				var znfu2 bool
				zpfn1, err = dc.ReadString()
				if err != nil {
					return
				}
				znfu2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zpfn1] = znfu2
			}
		case "Notyet_zid03_map":
			found0zdgq6[4] = true
			var zcnf8 uint32
			zcnf8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zcnf8 > 0 {
				z.Notyet = make(map[string]bool, zcnf8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zcnf8 > 0 {
				zcnf8--
				var zawy3 string
				var znou4 bool
				zawy3, err = dc.ReadString()
				if err != nil {
					return
				}
				znou4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zawy3] = znou4
			}
		case "Setm_zid04_slc":
			found0zdgq6[5] = true
			var zsjn9 uint32
			zsjn9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zsjn9) {
				z.Setm = (z.Setm)[:zsjn9]
			} else {
				z.Setm = make([]*inn, zsjn9)
			}
			for znct5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[znct5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[znct5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[znct5] == nil {
						z.Setm[znct5] = new(inn)
					}
					err = z.Setm[znct5].DecodeMsg(dc)
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
	if nextMiss0zdgq6 != -1 {
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
var decodeMsgFieldOrder0zdgq6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zdgq6 = []bool{false, true, false, false, false, false}

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
	var empty_zvzb10 [6]bool
	fieldsInUse_zwtg11 := z.fieldsNotEmpty(empty_zvzb10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zwtg11)
	if err != nil {
		return err
	}

	if !empty_zvzb10[0] {
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

	if !empty_zvzb10[2] {
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

	if !empty_zvzb10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zpfn1, znfu2 := range z.Snot {
			err = en.WriteString(zpfn1)
			if err != nil {
				return
			}
			err = en.WriteBool(znfu2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvzb10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zawy3, znou4 := range z.Notyet {
			err = en.WriteString(zawy3)
			if err != nil {
				return
			}
			err = en.WriteBool(znou4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvzb10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for znct5 := range z.Setm {
			if z.Setm[znct5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[znct5].EncodeMsg(en)
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
		for zpfn1, znfu2 := range z.Snot {
			o = msgp.AppendString(o, zpfn1)
			o = msgp.AppendBool(o, znfu2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zawy3, znou4 := range z.Notyet {
			o = msgp.AppendString(o, zawy3)
			o = msgp.AppendBool(o, znou4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for znct5 := range z.Setm {
			if z.Setm[znct5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[znct5].MarshalMsg(o)
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
	const maxFields1zjbv12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zjbv12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zjbv12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zjbv12 := totalEncodedFields1zjbv12
	missingFieldsLeft1zjbv12 := maxFields1zjbv12 - totalEncodedFields1zjbv12

	var nextMiss1zjbv12 int32 = -1
	var found1zjbv12 [maxFields1zjbv12]bool
	var curField1zjbv12 string

doneWithStruct1zjbv12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zjbv12 > 0 || missingFieldsLeft1zjbv12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zjbv12, missingFieldsLeft1zjbv12, msgp.ShowFound(found1zjbv12[:]), unmarshalMsgFieldOrder1zjbv12)
		if encodedFieldsLeft1zjbv12 > 0 {
			encodedFieldsLeft1zjbv12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zjbv12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zjbv12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zjbv12 = 0
			}
			for nextMiss1zjbv12 < maxFields1zjbv12 && (found1zjbv12[nextMiss1zjbv12] || unmarshalMsgFieldSkip1zjbv12[nextMiss1zjbv12]) {
				nextMiss1zjbv12++
			}
			if nextMiss1zjbv12 == maxFields1zjbv12 {
				// filled all the empty fields!
				break doneWithStruct1zjbv12
			}
			missingFieldsLeft1zjbv12--
			curField1zjbv12 = unmarshalMsgFieldOrder1zjbv12[nextMiss1zjbv12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zjbv12)
		switch curField1zjbv12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zjbv12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zjbv12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zjbv12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zfyn13 uint32
				zfyn13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zfyn13 > 0 {
					z.Snot = make(map[string]bool, zfyn13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zfyn13 > 0 {
					var zpfn1 string
					var znfu2 bool
					zfyn13--
					zpfn1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					znfu2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zpfn1] = znfu2
				}
			}
		case "Notyet_zid03_map":
			found1zjbv12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zvaf14 uint32
				zvaf14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zvaf14 > 0 {
					z.Notyet = make(map[string]bool, zvaf14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zvaf14 > 0 {
					var zawy3 string
					var znou4 bool
					zvaf14--
					zawy3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					znou4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zawy3] = znou4
				}
			}
		case "Setm_zid04_slc":
			found1zjbv12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zppp15 uint32
				zppp15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zppp15) {
					z.Setm = (z.Setm)[:zppp15]
				} else {
					z.Setm = make([]*inn, zppp15)
				}
				for znct5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[znct5] != nil {
							z.Setm[znct5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[znct5] {
								z.Setm[znct5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[znct5] == nil {
								z.Setm[znct5] = new(inn)
							}
							bts, err = z.Setm[znct5].UnmarshalMsg(bts)
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
	if nextMiss1zjbv12 != -1 {
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
var unmarshalMsgFieldOrder1zjbv12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zjbv12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zpfn1, znfu2 := range z.Snot {
			_ = znfu2
			_ = zpfn1
			s += msgp.StringPrefixSize + len(zpfn1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zawy3, znou4 := range z.Notyet {
			_ = znou4
			_ = zawy3
			s += msgp.StringPrefixSize + len(zawy3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for znct5 := range z.Setm {
		if z.Setm[znct5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[znct5].Msgsize()
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
	const maxFields2zchw18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zchw18 uint32
	totalEncodedFields2zchw18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zchw18 := totalEncodedFields2zchw18
	missingFieldsLeft2zchw18 := maxFields2zchw18 - totalEncodedFields2zchw18

	var nextMiss2zchw18 int32 = -1
	var found2zchw18 [maxFields2zchw18]bool
	var curField2zchw18 string

doneWithStruct2zchw18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zchw18 > 0 || missingFieldsLeft2zchw18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zchw18, missingFieldsLeft2zchw18, msgp.ShowFound(found2zchw18[:]), decodeMsgFieldOrder2zchw18)
		if encodedFieldsLeft2zchw18 > 0 {
			encodedFieldsLeft2zchw18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zchw18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zchw18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zchw18 = 0
			}
			for nextMiss2zchw18 < maxFields2zchw18 && (found2zchw18[nextMiss2zchw18] || decodeMsgFieldSkip2zchw18[nextMiss2zchw18]) {
				nextMiss2zchw18++
			}
			if nextMiss2zchw18 == maxFields2zchw18 {
				// filled all the empty fields!
				break doneWithStruct2zchw18
			}
			missingFieldsLeft2zchw18--
			curField2zchw18 = decodeMsgFieldOrder2zchw18[nextMiss2zchw18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zchw18)
		switch curField2zchw18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zchw18[0] = true
			var zrcg19 uint32
			zrcg19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zrcg19 > 0 {
				z.j = make(map[string]bool, zrcg19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zrcg19 > 0 {
				zrcg19--
				var znvm16 string
				var zhxw17 bool
				znvm16, err = dc.ReadString()
				if err != nil {
					return
				}
				zhxw17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[znvm16] = zhxw17
			}
		case "e_zid01_i64":
			found2zchw18[1] = true
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
	if nextMiss2zchw18 != -1 {
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
var decodeMsgFieldOrder2zchw18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zchw18 = []bool{false, false}

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
	var empty_zfpr20 [2]bool
	fieldsInUse_znnr21 := z.fieldsNotEmpty(empty_zfpr20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_znnr21)
	if err != nil {
		return err
	}

	if !empty_zfpr20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for znvm16, zhxw17 := range z.j {
			err = en.WriteString(znvm16)
			if err != nil {
				return
			}
			err = en.WriteBool(zhxw17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zfpr20[1] {
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
		for znvm16, zhxw17 := range z.j {
			o = msgp.AppendString(o, znvm16)
			o = msgp.AppendBool(o, zhxw17)
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
	const maxFields3zvrh22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zvrh22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zvrh22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zvrh22 := totalEncodedFields3zvrh22
	missingFieldsLeft3zvrh22 := maxFields3zvrh22 - totalEncodedFields3zvrh22

	var nextMiss3zvrh22 int32 = -1
	var found3zvrh22 [maxFields3zvrh22]bool
	var curField3zvrh22 string

doneWithStruct3zvrh22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zvrh22 > 0 || missingFieldsLeft3zvrh22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvrh22, missingFieldsLeft3zvrh22, msgp.ShowFound(found3zvrh22[:]), unmarshalMsgFieldOrder3zvrh22)
		if encodedFieldsLeft3zvrh22 > 0 {
			encodedFieldsLeft3zvrh22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zvrh22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zvrh22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zvrh22 = 0
			}
			for nextMiss3zvrh22 < maxFields3zvrh22 && (found3zvrh22[nextMiss3zvrh22] || unmarshalMsgFieldSkip3zvrh22[nextMiss3zvrh22]) {
				nextMiss3zvrh22++
			}
			if nextMiss3zvrh22 == maxFields3zvrh22 {
				// filled all the empty fields!
				break doneWithStruct3zvrh22
			}
			missingFieldsLeft3zvrh22--
			curField3zvrh22 = unmarshalMsgFieldOrder3zvrh22[nextMiss3zvrh22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zvrh22)
		switch curField3zvrh22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zvrh22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zxzo23 uint32
				zxzo23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zxzo23 > 0 {
					z.j = make(map[string]bool, zxzo23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zxzo23 > 0 {
					var znvm16 string
					var zhxw17 bool
					zxzo23--
					znvm16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zhxw17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[znvm16] = zhxw17
				}
			}
		case "e_zid01_i64":
			found3zvrh22[1] = true
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
	if nextMiss3zvrh22 != -1 {
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
var unmarshalMsgFieldOrder3zvrh22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zvrh22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for znvm16, zhxw17 := range z.j {
			_ = zhxw17
			_ = znvm16
			s += msgp.StringPrefixSize + len(znvm16) + msgp.BoolSize
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
	const maxFields4zjws26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zjws26 uint32
	totalEncodedFields4zjws26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zjws26 := totalEncodedFields4zjws26
	missingFieldsLeft4zjws26 := maxFields4zjws26 - totalEncodedFields4zjws26

	var nextMiss4zjws26 int32 = -1
	var found4zjws26 [maxFields4zjws26]bool
	var curField4zjws26 string

doneWithStruct4zjws26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zjws26 > 0 || missingFieldsLeft4zjws26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zjws26, missingFieldsLeft4zjws26, msgp.ShowFound(found4zjws26[:]), decodeMsgFieldOrder4zjws26)
		if encodedFieldsLeft4zjws26 > 0 {
			encodedFieldsLeft4zjws26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zjws26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zjws26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zjws26 = 0
			}
			for nextMiss4zjws26 < maxFields4zjws26 && (found4zjws26[nextMiss4zjws26] || decodeMsgFieldSkip4zjws26[nextMiss4zjws26]) {
				nextMiss4zjws26++
			}
			if nextMiss4zjws26 == maxFields4zjws26 {
				// filled all the empty fields!
				break doneWithStruct4zjws26
			}
			missingFieldsLeft4zjws26--
			curField4zjws26 = decodeMsgFieldOrder4zjws26[nextMiss4zjws26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zjws26)
		switch curField4zjws26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zjws26[0] = true
			var zpxn27 uint32
			zpxn27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zpxn27 > 0 {
				z.m = make(map[string]*Tr, zpxn27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zpxn27 > 0 {
				zpxn27--
				var zywh24 string
				var zvbs25 *Tr
				zywh24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zvbs25 != nil {
						dc.PushAlwaysNil()
						err = zvbs25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zvbs25 == nil {
						zvbs25 = new(Tr)
					}
					err = zvbs25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zywh24] = zvbs25
			}
		case "s_zid01_str":
			found4zjws26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zjws26[2] = true
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
	if nextMiss4zjws26 != -1 {
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
var decodeMsgFieldOrder4zjws26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zjws26 = []bool{false, false, false}

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
	var empty_zzqt28 [3]bool
	fieldsInUse_zjnc29 := z.fieldsNotEmpty(empty_zzqt28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjnc29)
	if err != nil {
		return err
	}

	if !empty_zzqt28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zywh24, zvbs25 := range z.m {
			err = en.WriteString(zywh24)
			if err != nil {
				return
			}
			if zvbs25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zvbs25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zzqt28[1] {
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

	if !empty_zzqt28[2] {
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
		for zywh24, zvbs25 := range z.m {
			o = msgp.AppendString(o, zywh24)
			if zvbs25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zvbs25.MarshalMsg(o)
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
	const maxFields5zysi30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zysi30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zysi30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zysi30 := totalEncodedFields5zysi30
	missingFieldsLeft5zysi30 := maxFields5zysi30 - totalEncodedFields5zysi30

	var nextMiss5zysi30 int32 = -1
	var found5zysi30 [maxFields5zysi30]bool
	var curField5zysi30 string

doneWithStruct5zysi30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zysi30 > 0 || missingFieldsLeft5zysi30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zysi30, missingFieldsLeft5zysi30, msgp.ShowFound(found5zysi30[:]), unmarshalMsgFieldOrder5zysi30)
		if encodedFieldsLeft5zysi30 > 0 {
			encodedFieldsLeft5zysi30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zysi30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zysi30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zysi30 = 0
			}
			for nextMiss5zysi30 < maxFields5zysi30 && (found5zysi30[nextMiss5zysi30] || unmarshalMsgFieldSkip5zysi30[nextMiss5zysi30]) {
				nextMiss5zysi30++
			}
			if nextMiss5zysi30 == maxFields5zysi30 {
				// filled all the empty fields!
				break doneWithStruct5zysi30
			}
			missingFieldsLeft5zysi30--
			curField5zysi30 = unmarshalMsgFieldOrder5zysi30[nextMiss5zysi30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zysi30)
		switch curField5zysi30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zysi30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zuuf31 uint32
				zuuf31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zuuf31 > 0 {
					z.m = make(map[string]*Tr, zuuf31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zuuf31 > 0 {
					var zywh24 string
					var zvbs25 *Tr
					zuuf31--
					zywh24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zvbs25 != nil {
							zvbs25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zvbs25 {
								zvbs25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zvbs25 == nil {
								zvbs25 = new(Tr)
							}
							bts, err = zvbs25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zywh24] = zvbs25
				}
			}
		case "s_zid01_str":
			found5zysi30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zysi30[2] = true
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
	if nextMiss5zysi30 != -1 {
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
var unmarshalMsgFieldOrder5zysi30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zysi30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zywh24, zvbs25 := range z.m {
			_ = zvbs25
			_ = zywh24
			s += msgp.StringPrefixSize + len(zywh24)
			if zvbs25 == nil {
				s += msgp.NilSize
			} else {
				s += zvbs25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
