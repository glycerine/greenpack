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
	const maxFields0zgpl6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zgpl6 uint32
	totalEncodedFields0zgpl6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zgpl6 := totalEncodedFields0zgpl6
	missingFieldsLeft0zgpl6 := maxFields0zgpl6 - totalEncodedFields0zgpl6

	var nextMiss0zgpl6 int32 = -1
	var found0zgpl6 [maxFields0zgpl6]bool
	var curField0zgpl6 string

doneWithStruct0zgpl6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgpl6 > 0 || missingFieldsLeft0zgpl6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgpl6, missingFieldsLeft0zgpl6, msgp.ShowFound(found0zgpl6[:]), decodeMsgFieldOrder0zgpl6)
		if encodedFieldsLeft0zgpl6 > 0 {
			encodedFieldsLeft0zgpl6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zgpl6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgpl6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zgpl6 = 0
			}
			for nextMiss0zgpl6 < maxFields0zgpl6 && (found0zgpl6[nextMiss0zgpl6] || decodeMsgFieldSkip0zgpl6[nextMiss0zgpl6]) {
				nextMiss0zgpl6++
			}
			if nextMiss0zgpl6 == maxFields0zgpl6 {
				// filled all the empty fields!
				break doneWithStruct0zgpl6
			}
			missingFieldsLeft0zgpl6--
			curField0zgpl6 = decodeMsgFieldOrder0zgpl6[nextMiss0zgpl6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgpl6)
		switch curField0zgpl6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zgpl6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zgpl6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zgpl6[3] = true
			var zwkz7 uint32
			zwkz7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zwkz7 > 0 {
				z.Snot = make(map[string]bool, zwkz7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zwkz7 > 0 {
				zwkz7--
				var zips1 string
				var zjim2 bool
				zips1, err = dc.ReadString()
				if err != nil {
					return
				}
				zjim2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zips1] = zjim2
			}
		case "Notyet_zid03_map":
			found0zgpl6[4] = true
			var zndi8 uint32
			zndi8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zndi8 > 0 {
				z.Notyet = make(map[string]bool, zndi8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zndi8 > 0 {
				zndi8--
				var zmjs3 string
				var zait4 bool
				zmjs3, err = dc.ReadString()
				if err != nil {
					return
				}
				zait4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zmjs3] = zait4
			}
		case "Setm_zid04_slc":
			found0zgpl6[5] = true
			var zton9 uint32
			zton9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zton9) {
				z.Setm = (z.Setm)[:zton9]
			} else {
				z.Setm = make([]*inn, zton9)
			}
			for zhnw5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zhnw5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zhnw5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zhnw5] == nil {
						z.Setm[zhnw5] = new(inn)
					}
					err = z.Setm[zhnw5].DecodeMsg(dc)
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
	if nextMiss0zgpl6 != -1 {
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
var decodeMsgFieldOrder0zgpl6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zgpl6 = []bool{false, true, false, false, false, false}

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
	var empty_zhvx10 [6]bool
	fieldsInUse_zihl11 := z.fieldsNotEmpty(empty_zhvx10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zihl11)
	if err != nil {
		return err
	}

	if !empty_zhvx10[0] {
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

	if !empty_zhvx10[2] {
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

	if !empty_zhvx10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zips1, zjim2 := range z.Snot {
			err = en.WriteString(zips1)
			if err != nil {
				return
			}
			err = en.WriteBool(zjim2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zhvx10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zmjs3, zait4 := range z.Notyet {
			err = en.WriteString(zmjs3)
			if err != nil {
				return
			}
			err = en.WriteBool(zait4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zhvx10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zhnw5 := range z.Setm {
			if z.Setm[zhnw5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zhnw5].EncodeMsg(en)
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
		for zips1, zjim2 := range z.Snot {
			o = msgp.AppendString(o, zips1)
			o = msgp.AppendBool(o, zjim2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zmjs3, zait4 := range z.Notyet {
			o = msgp.AppendString(o, zmjs3)
			o = msgp.AppendBool(o, zait4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zhnw5 := range z.Setm {
			if z.Setm[zhnw5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zhnw5].MarshalMsg(o)
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
	const maxFields1zwax12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zwax12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zwax12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zwax12 := totalEncodedFields1zwax12
	missingFieldsLeft1zwax12 := maxFields1zwax12 - totalEncodedFields1zwax12

	var nextMiss1zwax12 int32 = -1
	var found1zwax12 [maxFields1zwax12]bool
	var curField1zwax12 string

doneWithStruct1zwax12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zwax12 > 0 || missingFieldsLeft1zwax12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zwax12, missingFieldsLeft1zwax12, msgp.ShowFound(found1zwax12[:]), unmarshalMsgFieldOrder1zwax12)
		if encodedFieldsLeft1zwax12 > 0 {
			encodedFieldsLeft1zwax12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zwax12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zwax12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zwax12 = 0
			}
			for nextMiss1zwax12 < maxFields1zwax12 && (found1zwax12[nextMiss1zwax12] || unmarshalMsgFieldSkip1zwax12[nextMiss1zwax12]) {
				nextMiss1zwax12++
			}
			if nextMiss1zwax12 == maxFields1zwax12 {
				// filled all the empty fields!
				break doneWithStruct1zwax12
			}
			missingFieldsLeft1zwax12--
			curField1zwax12 = unmarshalMsgFieldOrder1zwax12[nextMiss1zwax12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zwax12)
		switch curField1zwax12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zwax12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zwax12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zwax12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zaxz13 uint32
				zaxz13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zaxz13 > 0 {
					z.Snot = make(map[string]bool, zaxz13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zaxz13 > 0 {
					var zips1 string
					var zjim2 bool
					zaxz13--
					zips1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zjim2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zips1] = zjim2
				}
			}
		case "Notyet_zid03_map":
			found1zwax12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zvbg14 uint32
				zvbg14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zvbg14 > 0 {
					z.Notyet = make(map[string]bool, zvbg14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zvbg14 > 0 {
					var zmjs3 string
					var zait4 bool
					zvbg14--
					zmjs3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zait4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zmjs3] = zait4
				}
			}
		case "Setm_zid04_slc":
			found1zwax12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zqkv15 uint32
				zqkv15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zqkv15) {
					z.Setm = (z.Setm)[:zqkv15]
				} else {
					z.Setm = make([]*inn, zqkv15)
				}
				for zhnw5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zhnw5] != nil {
							z.Setm[zhnw5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zhnw5] {
								z.Setm[zhnw5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zhnw5] == nil {
								z.Setm[zhnw5] = new(inn)
							}
							bts, err = z.Setm[zhnw5].UnmarshalMsg(bts)
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
	if nextMiss1zwax12 != -1 {
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
var unmarshalMsgFieldOrder1zwax12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zwax12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zips1, zjim2 := range z.Snot {
			_ = zjim2
			_ = zips1
			s += msgp.StringPrefixSize + len(zips1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zmjs3, zait4 := range z.Notyet {
			_ = zait4
			_ = zmjs3
			s += msgp.StringPrefixSize + len(zmjs3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zhnw5 := range z.Setm {
		if z.Setm[zhnw5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zhnw5].Msgsize()
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
	const maxFields2zctd18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zctd18 uint32
	totalEncodedFields2zctd18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zctd18 := totalEncodedFields2zctd18
	missingFieldsLeft2zctd18 := maxFields2zctd18 - totalEncodedFields2zctd18

	var nextMiss2zctd18 int32 = -1
	var found2zctd18 [maxFields2zctd18]bool
	var curField2zctd18 string

doneWithStruct2zctd18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zctd18 > 0 || missingFieldsLeft2zctd18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zctd18, missingFieldsLeft2zctd18, msgp.ShowFound(found2zctd18[:]), decodeMsgFieldOrder2zctd18)
		if encodedFieldsLeft2zctd18 > 0 {
			encodedFieldsLeft2zctd18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zctd18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zctd18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zctd18 = 0
			}
			for nextMiss2zctd18 < maxFields2zctd18 && (found2zctd18[nextMiss2zctd18] || decodeMsgFieldSkip2zctd18[nextMiss2zctd18]) {
				nextMiss2zctd18++
			}
			if nextMiss2zctd18 == maxFields2zctd18 {
				// filled all the empty fields!
				break doneWithStruct2zctd18
			}
			missingFieldsLeft2zctd18--
			curField2zctd18 = decodeMsgFieldOrder2zctd18[nextMiss2zctd18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zctd18)
		switch curField2zctd18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zctd18[0] = true
			var znwg19 uint32
			znwg19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && znwg19 > 0 {
				z.j = make(map[string]bool, znwg19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for znwg19 > 0 {
				znwg19--
				var zjjt16 string
				var zthc17 bool
				zjjt16, err = dc.ReadString()
				if err != nil {
					return
				}
				zthc17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zjjt16] = zthc17
			}
		case "e_zid01_i64":
			found2zctd18[1] = true
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
	if nextMiss2zctd18 != -1 {
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
var decodeMsgFieldOrder2zctd18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zctd18 = []bool{false, false}

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
	var empty_zyss20 [2]bool
	fieldsInUse_zpde21 := z.fieldsNotEmpty(empty_zyss20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpde21)
	if err != nil {
		return err
	}

	if !empty_zyss20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zjjt16, zthc17 := range z.j {
			err = en.WriteString(zjjt16)
			if err != nil {
				return
			}
			err = en.WriteBool(zthc17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zyss20[1] {
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
		for zjjt16, zthc17 := range z.j {
			o = msgp.AppendString(o, zjjt16)
			o = msgp.AppendBool(o, zthc17)
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
	const maxFields3zmzh22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zmzh22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zmzh22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zmzh22 := totalEncodedFields3zmzh22
	missingFieldsLeft3zmzh22 := maxFields3zmzh22 - totalEncodedFields3zmzh22

	var nextMiss3zmzh22 int32 = -1
	var found3zmzh22 [maxFields3zmzh22]bool
	var curField3zmzh22 string

doneWithStruct3zmzh22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zmzh22 > 0 || missingFieldsLeft3zmzh22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zmzh22, missingFieldsLeft3zmzh22, msgp.ShowFound(found3zmzh22[:]), unmarshalMsgFieldOrder3zmzh22)
		if encodedFieldsLeft3zmzh22 > 0 {
			encodedFieldsLeft3zmzh22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zmzh22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zmzh22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zmzh22 = 0
			}
			for nextMiss3zmzh22 < maxFields3zmzh22 && (found3zmzh22[nextMiss3zmzh22] || unmarshalMsgFieldSkip3zmzh22[nextMiss3zmzh22]) {
				nextMiss3zmzh22++
			}
			if nextMiss3zmzh22 == maxFields3zmzh22 {
				// filled all the empty fields!
				break doneWithStruct3zmzh22
			}
			missingFieldsLeft3zmzh22--
			curField3zmzh22 = unmarshalMsgFieldOrder3zmzh22[nextMiss3zmzh22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zmzh22)
		switch curField3zmzh22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zmzh22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zmmp23 uint32
				zmmp23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zmmp23 > 0 {
					z.j = make(map[string]bool, zmmp23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zmmp23 > 0 {
					var zjjt16 string
					var zthc17 bool
					zmmp23--
					zjjt16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zthc17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zjjt16] = zthc17
				}
			}
		case "e_zid01_i64":
			found3zmzh22[1] = true
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
	if nextMiss3zmzh22 != -1 {
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
var unmarshalMsgFieldOrder3zmzh22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zmzh22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zjjt16, zthc17 := range z.j {
			_ = zthc17
			_ = zjjt16
			s += msgp.StringPrefixSize + len(zjjt16) + msgp.BoolSize
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
	const maxFields4zwzy26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zwzy26 uint32
	totalEncodedFields4zwzy26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zwzy26 := totalEncodedFields4zwzy26
	missingFieldsLeft4zwzy26 := maxFields4zwzy26 - totalEncodedFields4zwzy26

	var nextMiss4zwzy26 int32 = -1
	var found4zwzy26 [maxFields4zwzy26]bool
	var curField4zwzy26 string

doneWithStruct4zwzy26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zwzy26 > 0 || missingFieldsLeft4zwzy26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zwzy26, missingFieldsLeft4zwzy26, msgp.ShowFound(found4zwzy26[:]), decodeMsgFieldOrder4zwzy26)
		if encodedFieldsLeft4zwzy26 > 0 {
			encodedFieldsLeft4zwzy26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zwzy26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zwzy26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zwzy26 = 0
			}
			for nextMiss4zwzy26 < maxFields4zwzy26 && (found4zwzy26[nextMiss4zwzy26] || decodeMsgFieldSkip4zwzy26[nextMiss4zwzy26]) {
				nextMiss4zwzy26++
			}
			if nextMiss4zwzy26 == maxFields4zwzy26 {
				// filled all the empty fields!
				break doneWithStruct4zwzy26
			}
			missingFieldsLeft4zwzy26--
			curField4zwzy26 = decodeMsgFieldOrder4zwzy26[nextMiss4zwzy26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zwzy26)
		switch curField4zwzy26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zwzy26[0] = true
			var zupg27 uint32
			zupg27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zupg27 > 0 {
				z.m = make(map[string]*Tr, zupg27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zupg27 > 0 {
				zupg27--
				var zkne24 string
				var zcst25 *Tr
				zkne24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zcst25 != nil {
						dc.PushAlwaysNil()
						err = zcst25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zcst25 == nil {
						zcst25 = new(Tr)
					}
					err = zcst25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zkne24] = zcst25
			}
		case "s_zid01_str":
			found4zwzy26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zwzy26[2] = true
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
	if nextMiss4zwzy26 != -1 {
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
var decodeMsgFieldOrder4zwzy26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zwzy26 = []bool{false, false, false}

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
	var empty_zlve28 [3]bool
	fieldsInUse_zuzg29 := z.fieldsNotEmpty(empty_zlve28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zuzg29)
	if err != nil {
		return err
	}

	if !empty_zlve28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zkne24, zcst25 := range z.m {
			err = en.WriteString(zkne24)
			if err != nil {
				return
			}
			if zcst25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zcst25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zlve28[1] {
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

	if !empty_zlve28[2] {
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
		for zkne24, zcst25 := range z.m {
			o = msgp.AppendString(o, zkne24)
			if zcst25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zcst25.MarshalMsg(o)
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
	const maxFields5zjyx30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zjyx30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zjyx30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zjyx30 := totalEncodedFields5zjyx30
	missingFieldsLeft5zjyx30 := maxFields5zjyx30 - totalEncodedFields5zjyx30

	var nextMiss5zjyx30 int32 = -1
	var found5zjyx30 [maxFields5zjyx30]bool
	var curField5zjyx30 string

doneWithStruct5zjyx30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zjyx30 > 0 || missingFieldsLeft5zjyx30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zjyx30, missingFieldsLeft5zjyx30, msgp.ShowFound(found5zjyx30[:]), unmarshalMsgFieldOrder5zjyx30)
		if encodedFieldsLeft5zjyx30 > 0 {
			encodedFieldsLeft5zjyx30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zjyx30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zjyx30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zjyx30 = 0
			}
			for nextMiss5zjyx30 < maxFields5zjyx30 && (found5zjyx30[nextMiss5zjyx30] || unmarshalMsgFieldSkip5zjyx30[nextMiss5zjyx30]) {
				nextMiss5zjyx30++
			}
			if nextMiss5zjyx30 == maxFields5zjyx30 {
				// filled all the empty fields!
				break doneWithStruct5zjyx30
			}
			missingFieldsLeft5zjyx30--
			curField5zjyx30 = unmarshalMsgFieldOrder5zjyx30[nextMiss5zjyx30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zjyx30)
		switch curField5zjyx30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zjyx30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var znud31 uint32
				znud31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && znud31 > 0 {
					z.m = make(map[string]*Tr, znud31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for znud31 > 0 {
					var zkne24 string
					var zcst25 *Tr
					znud31--
					zkne24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zcst25 != nil {
							zcst25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zcst25 {
								zcst25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zcst25 == nil {
								zcst25 = new(Tr)
							}
							bts, err = zcst25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zkne24] = zcst25
				}
			}
		case "s_zid01_str":
			found5zjyx30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zjyx30[2] = true
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
	if nextMiss5zjyx30 != -1 {
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
var unmarshalMsgFieldOrder5zjyx30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zjyx30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zkne24, zcst25 := range z.m {
			_ = zcst25
			_ = zkne24
			s += msgp.StringPrefixSize + len(zkne24)
			if zcst25 == nil {
				s += msgp.NilSize
			} else {
				s += zcst25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
