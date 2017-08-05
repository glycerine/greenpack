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
	const maxFields0zlun6 = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zlun6 uint32
	totalEncodedFields0zlun6, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zlun6 := totalEncodedFields0zlun6
	missingFieldsLeft0zlun6 := maxFields0zlun6 - totalEncodedFields0zlun6

	var nextMiss0zlun6 int32 = -1
	var found0zlun6 [maxFields0zlun6]bool
	var curField0zlun6 string

doneWithStruct0zlun6:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zlun6 > 0 || missingFieldsLeft0zlun6 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zlun6, missingFieldsLeft0zlun6, msgp.ShowFound(found0zlun6[:]), decodeMsgFieldOrder0zlun6)
		if encodedFieldsLeft0zlun6 > 0 {
			encodedFieldsLeft0zlun6--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zlun6 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zlun6 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zlun6 = 0
			}
			for nextMiss0zlun6 < maxFields0zlun6 && (found0zlun6[nextMiss0zlun6] || decodeMsgFieldSkip0zlun6[nextMiss0zlun6]) {
				nextMiss0zlun6++
			}
			if nextMiss0zlun6 == maxFields0zlun6 {
				// filled all the empty fields!
				break doneWithStruct0zlun6
			}
			missingFieldsLeft0zlun6--
			curField0zlun6 = decodeMsgFieldOrder0zlun6[nextMiss0zlun6]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zlun6)
		switch curField0zlun6 {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zlun6[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zlun6[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zlun6[3] = true
			var zfzf7 uint32
			zfzf7, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zfzf7 > 0 {
				z.Snot = make(map[string]bool, zfzf7)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zfzf7 > 0 {
				zfzf7--
				var zqbu1 string
				var zzzf2 bool
				zqbu1, err = dc.ReadString()
				if err != nil {
					return
				}
				zzzf2, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zqbu1] = zzzf2
			}
		case "Notyet_zid03_map":
			found0zlun6[4] = true
			var zgqk8 uint32
			zgqk8, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zgqk8 > 0 {
				z.Notyet = make(map[string]bool, zgqk8)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zgqk8 > 0 {
				zgqk8--
				var zbff3 string
				var zvpy4 bool
				zbff3, err = dc.ReadString()
				if err != nil {
					return
				}
				zvpy4, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zbff3] = zvpy4
			}
		case "Setm_zid04_slc":
			found0zlun6[5] = true
			var zlzx9 uint32
			zlzx9, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zlzx9) {
				z.Setm = (z.Setm)[:zlzx9]
			} else {
				z.Setm = make([]*inn, zlzx9)
			}
			for zkkt5 := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zkkt5] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zkkt5].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zkkt5] == nil {
						z.Setm[zkkt5] = new(inn)
					}
					err = z.Setm[zkkt5].DecodeMsg(dc)
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
	if nextMiss0zlun6 != -1 {
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
var decodeMsgFieldOrder0zlun6 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zlun6 = []bool{false, true, false, false, false, false}

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
	var empty_zvpg10 [6]bool
	fieldsInUse_zwdy11 := z.fieldsNotEmpty(empty_zvpg10[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zwdy11)
	if err != nil {
		return err
	}

	if !empty_zvpg10[0] {
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

	if !empty_zvpg10[2] {
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

	if !empty_zvpg10[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zqbu1, zzzf2 := range z.Snot {
			err = en.WriteString(zqbu1)
			if err != nil {
				return
			}
			err = en.WriteBool(zzzf2)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvpg10[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zbff3, zvpy4 := range z.Notyet {
			err = en.WriteString(zbff3)
			if err != nil {
				return
			}
			err = en.WriteBool(zvpy4)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvpg10[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zkkt5 := range z.Setm {
			if z.Setm[zkkt5] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zkkt5].EncodeMsg(en)
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
		for zqbu1, zzzf2 := range z.Snot {
			o = msgp.AppendString(o, zqbu1)
			o = msgp.AppendBool(o, zzzf2)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zbff3, zvpy4 := range z.Notyet {
			o = msgp.AppendString(o, zbff3)
			o = msgp.AppendBool(o, zvpy4)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zkkt5 := range z.Setm {
			if z.Setm[zkkt5] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zkkt5].MarshalMsg(o)
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
	const maxFields1zxee12 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zxee12 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zxee12, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zxee12 := totalEncodedFields1zxee12
	missingFieldsLeft1zxee12 := maxFields1zxee12 - totalEncodedFields1zxee12

	var nextMiss1zxee12 int32 = -1
	var found1zxee12 [maxFields1zxee12]bool
	var curField1zxee12 string

doneWithStruct1zxee12:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zxee12 > 0 || missingFieldsLeft1zxee12 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zxee12, missingFieldsLeft1zxee12, msgp.ShowFound(found1zxee12[:]), unmarshalMsgFieldOrder1zxee12)
		if encodedFieldsLeft1zxee12 > 0 {
			encodedFieldsLeft1zxee12--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zxee12 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zxee12 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zxee12 = 0
			}
			for nextMiss1zxee12 < maxFields1zxee12 && (found1zxee12[nextMiss1zxee12] || unmarshalMsgFieldSkip1zxee12[nextMiss1zxee12]) {
				nextMiss1zxee12++
			}
			if nextMiss1zxee12 == maxFields1zxee12 {
				// filled all the empty fields!
				break doneWithStruct1zxee12
			}
			missingFieldsLeft1zxee12--
			curField1zxee12 = unmarshalMsgFieldOrder1zxee12[nextMiss1zxee12]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zxee12)
		switch curField1zxee12 {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zxee12[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zxee12[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zxee12[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zmpp13 uint32
				zmpp13, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zmpp13 > 0 {
					z.Snot = make(map[string]bool, zmpp13)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zmpp13 > 0 {
					var zqbu1 string
					var zzzf2 bool
					zmpp13--
					zqbu1, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zzzf2, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zqbu1] = zzzf2
				}
			}
		case "Notyet_zid03_map":
			found1zxee12[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zfhk14 uint32
				zfhk14, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zfhk14 > 0 {
					z.Notyet = make(map[string]bool, zfhk14)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zfhk14 > 0 {
					var zbff3 string
					var zvpy4 bool
					zfhk14--
					zbff3, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zvpy4, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zbff3] = zvpy4
				}
			}
		case "Setm_zid04_slc":
			found1zxee12[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var ztsq15 uint32
				ztsq15, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(ztsq15) {
					z.Setm = (z.Setm)[:ztsq15]
				} else {
					z.Setm = make([]*inn, ztsq15)
				}
				for zkkt5 := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zkkt5] != nil {
							z.Setm[zkkt5].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zkkt5] {
								z.Setm[zkkt5].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zkkt5] == nil {
								z.Setm[zkkt5] = new(inn)
							}
							bts, err = z.Setm[zkkt5].UnmarshalMsg(bts)
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
	if nextMiss1zxee12 != -1 {
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
var unmarshalMsgFieldOrder1zxee12 = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zxee12 = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zqbu1, zzzf2 := range z.Snot {
			_ = zzzf2
			_ = zqbu1
			s += msgp.StringPrefixSize + len(zqbu1) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zbff3, zvpy4 := range z.Notyet {
			_ = zvpy4
			_ = zbff3
			s += msgp.StringPrefixSize + len(zbff3) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zkkt5 := range z.Setm {
		if z.Setm[zkkt5] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zkkt5].Msgsize()
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
	const maxFields2zthi18 = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zthi18 uint32
	totalEncodedFields2zthi18, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zthi18 := totalEncodedFields2zthi18
	missingFieldsLeft2zthi18 := maxFields2zthi18 - totalEncodedFields2zthi18

	var nextMiss2zthi18 int32 = -1
	var found2zthi18 [maxFields2zthi18]bool
	var curField2zthi18 string

doneWithStruct2zthi18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zthi18 > 0 || missingFieldsLeft2zthi18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zthi18, missingFieldsLeft2zthi18, msgp.ShowFound(found2zthi18[:]), decodeMsgFieldOrder2zthi18)
		if encodedFieldsLeft2zthi18 > 0 {
			encodedFieldsLeft2zthi18--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zthi18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zthi18 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zthi18 = 0
			}
			for nextMiss2zthi18 < maxFields2zthi18 && (found2zthi18[nextMiss2zthi18] || decodeMsgFieldSkip2zthi18[nextMiss2zthi18]) {
				nextMiss2zthi18++
			}
			if nextMiss2zthi18 == maxFields2zthi18 {
				// filled all the empty fields!
				break doneWithStruct2zthi18
			}
			missingFieldsLeft2zthi18--
			curField2zthi18 = decodeMsgFieldOrder2zthi18[nextMiss2zthi18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zthi18)
		switch curField2zthi18 {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zthi18[0] = true
			var zwtc19 uint32
			zwtc19, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zwtc19 > 0 {
				z.j = make(map[string]bool, zwtc19)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zwtc19 > 0 {
				zwtc19--
				var zcxf16 string
				var zvss17 bool
				zcxf16, err = dc.ReadString()
				if err != nil {
					return
				}
				zvss17, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zcxf16] = zvss17
			}
		case "e_zid01_i64":
			found2zthi18[1] = true
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
	if nextMiss2zthi18 != -1 {
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
var decodeMsgFieldOrder2zthi18 = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zthi18 = []bool{false, false}

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
	var empty_zvzj20 [2]bool
	fieldsInUse_zaqp21 := z.fieldsNotEmpty(empty_zvzj20[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zaqp21)
	if err != nil {
		return err
	}

	if !empty_zvzj20[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zcxf16, zvss17 := range z.j {
			err = en.WriteString(zcxf16)
			if err != nil {
				return
			}
			err = en.WriteBool(zvss17)
			if err != nil {
				return
			}
		}
	}

	if !empty_zvzj20[1] {
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
		for zcxf16, zvss17 := range z.j {
			o = msgp.AppendString(o, zcxf16)
			o = msgp.AppendBool(o, zvss17)
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
	const maxFields3zacq22 = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zacq22 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zacq22, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zacq22 := totalEncodedFields3zacq22
	missingFieldsLeft3zacq22 := maxFields3zacq22 - totalEncodedFields3zacq22

	var nextMiss3zacq22 int32 = -1
	var found3zacq22 [maxFields3zacq22]bool
	var curField3zacq22 string

doneWithStruct3zacq22:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zacq22 > 0 || missingFieldsLeft3zacq22 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zacq22, missingFieldsLeft3zacq22, msgp.ShowFound(found3zacq22[:]), unmarshalMsgFieldOrder3zacq22)
		if encodedFieldsLeft3zacq22 > 0 {
			encodedFieldsLeft3zacq22--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zacq22 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zacq22 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zacq22 = 0
			}
			for nextMiss3zacq22 < maxFields3zacq22 && (found3zacq22[nextMiss3zacq22] || unmarshalMsgFieldSkip3zacq22[nextMiss3zacq22]) {
				nextMiss3zacq22++
			}
			if nextMiss3zacq22 == maxFields3zacq22 {
				// filled all the empty fields!
				break doneWithStruct3zacq22
			}
			missingFieldsLeft3zacq22--
			curField3zacq22 = unmarshalMsgFieldOrder3zacq22[nextMiss3zacq22]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zacq22)
		switch curField3zacq22 {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zacq22[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zpnr23 uint32
				zpnr23, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zpnr23 > 0 {
					z.j = make(map[string]bool, zpnr23)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zpnr23 > 0 {
					var zcxf16 string
					var zvss17 bool
					zpnr23--
					zcxf16, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zvss17, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zcxf16] = zvss17
				}
			}
		case "e_zid01_i64":
			found3zacq22[1] = true
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
	if nextMiss3zacq22 != -1 {
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
var unmarshalMsgFieldOrder3zacq22 = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zacq22 = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zcxf16, zvss17 := range z.j {
			_ = zvss17
			_ = zcxf16
			s += msgp.StringPrefixSize + len(zcxf16) + msgp.BoolSize
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
	const maxFields4ztjk26 = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4ztjk26 uint32
	totalEncodedFields4ztjk26, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4ztjk26 := totalEncodedFields4ztjk26
	missingFieldsLeft4ztjk26 := maxFields4ztjk26 - totalEncodedFields4ztjk26

	var nextMiss4ztjk26 int32 = -1
	var found4ztjk26 [maxFields4ztjk26]bool
	var curField4ztjk26 string

doneWithStruct4ztjk26:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4ztjk26 > 0 || missingFieldsLeft4ztjk26 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4ztjk26, missingFieldsLeft4ztjk26, msgp.ShowFound(found4ztjk26[:]), decodeMsgFieldOrder4ztjk26)
		if encodedFieldsLeft4ztjk26 > 0 {
			encodedFieldsLeft4ztjk26--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4ztjk26 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4ztjk26 < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4ztjk26 = 0
			}
			for nextMiss4ztjk26 < maxFields4ztjk26 && (found4ztjk26[nextMiss4ztjk26] || decodeMsgFieldSkip4ztjk26[nextMiss4ztjk26]) {
				nextMiss4ztjk26++
			}
			if nextMiss4ztjk26 == maxFields4ztjk26 {
				// filled all the empty fields!
				break doneWithStruct4ztjk26
			}
			missingFieldsLeft4ztjk26--
			curField4ztjk26 = decodeMsgFieldOrder4ztjk26[nextMiss4ztjk26]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4ztjk26)
		switch curField4ztjk26 {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4ztjk26[0] = true
			var zaen27 uint32
			zaen27, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zaen27 > 0 {
				z.m = make(map[string]*Tr, zaen27)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zaen27 > 0 {
				zaen27--
				var zomi24 string
				var zwsk25 *Tr
				zomi24, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zwsk25 != nil {
						dc.PushAlwaysNil()
						err = zwsk25.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zwsk25 == nil {
						zwsk25 = new(Tr)
					}
					err = zwsk25.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zomi24] = zwsk25
			}
		case "s_zid01_str":
			found4ztjk26[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4ztjk26[2] = true
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
	if nextMiss4ztjk26 != -1 {
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
var decodeMsgFieldOrder4ztjk26 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4ztjk26 = []bool{false, false, false}

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
	var empty_zdue28 [3]bool
	fieldsInUse_zsab29 := z.fieldsNotEmpty(empty_zdue28[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zsab29)
	if err != nil {
		return err
	}

	if !empty_zdue28[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zomi24, zwsk25 := range z.m {
			err = en.WriteString(zomi24)
			if err != nil {
				return
			}
			if zwsk25 == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zwsk25.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zdue28[1] {
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

	if !empty_zdue28[2] {
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
		for zomi24, zwsk25 := range z.m {
			o = msgp.AppendString(o, zomi24)
			if zwsk25 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zwsk25.MarshalMsg(o)
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
	const maxFields5zkiq30 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zkiq30 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zkiq30, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zkiq30 := totalEncodedFields5zkiq30
	missingFieldsLeft5zkiq30 := maxFields5zkiq30 - totalEncodedFields5zkiq30

	var nextMiss5zkiq30 int32 = -1
	var found5zkiq30 [maxFields5zkiq30]bool
	var curField5zkiq30 string

doneWithStruct5zkiq30:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zkiq30 > 0 || missingFieldsLeft5zkiq30 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zkiq30, missingFieldsLeft5zkiq30, msgp.ShowFound(found5zkiq30[:]), unmarshalMsgFieldOrder5zkiq30)
		if encodedFieldsLeft5zkiq30 > 0 {
			encodedFieldsLeft5zkiq30--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zkiq30 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zkiq30 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zkiq30 = 0
			}
			for nextMiss5zkiq30 < maxFields5zkiq30 && (found5zkiq30[nextMiss5zkiq30] || unmarshalMsgFieldSkip5zkiq30[nextMiss5zkiq30]) {
				nextMiss5zkiq30++
			}
			if nextMiss5zkiq30 == maxFields5zkiq30 {
				// filled all the empty fields!
				break doneWithStruct5zkiq30
			}
			missingFieldsLeft5zkiq30--
			curField5zkiq30 = unmarshalMsgFieldOrder5zkiq30[nextMiss5zkiq30]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zkiq30)
		switch curField5zkiq30 {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zkiq30[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zsmk31 uint32
				zsmk31, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zsmk31 > 0 {
					z.m = make(map[string]*Tr, zsmk31)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zsmk31 > 0 {
					var zomi24 string
					var zwsk25 *Tr
					zsmk31--
					zomi24, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zwsk25 != nil {
							zwsk25.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zwsk25 {
								zwsk25.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zwsk25 == nil {
								zwsk25 = new(Tr)
							}
							bts, err = zwsk25.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zomi24] = zwsk25
				}
			}
		case "s_zid01_str":
			found5zkiq30[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zkiq30[2] = true
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
	if nextMiss5zkiq30 != -1 {
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
var unmarshalMsgFieldOrder5zkiq30 = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zkiq30 = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zomi24, zwsk25 := range z.m {
			_ = zwsk25
			_ = zomi24
			s += msgp.StringPrefixSize + len(zomi24)
			if zwsk25 == nil {
				s += msgp.NilSize
			} else {
				s += zwsk25.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
