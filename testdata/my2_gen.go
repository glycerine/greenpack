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
	const maxFields0zhke = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zhke uint32
	totalEncodedFields0zhke, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zhke := totalEncodedFields0zhke
	missingFieldsLeft0zhke := maxFields0zhke - totalEncodedFields0zhke

	var nextMiss0zhke int32 = -1
	var found0zhke [maxFields0zhke]bool
	var curField0zhke string

doneWithStruct0zhke:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zhke > 0 || missingFieldsLeft0zhke > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zhke, missingFieldsLeft0zhke, msgp.ShowFound(found0zhke[:]), decodeMsgFieldOrder0zhke)
		if encodedFieldsLeft0zhke > 0 {
			encodedFieldsLeft0zhke--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zhke = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zhke < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zhke = 0
			}
			for nextMiss0zhke < maxFields0zhke && (found0zhke[nextMiss0zhke] || decodeMsgFieldSkip0zhke[nextMiss0zhke]) {
				nextMiss0zhke++
			}
			if nextMiss0zhke == maxFields0zhke {
				// filled all the empty fields!
				break doneWithStruct0zhke
			}
			missingFieldsLeft0zhke--
			curField0zhke = decodeMsgFieldOrder0zhke[nextMiss0zhke]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zhke)
		switch curField0zhke {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zhke[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zhke[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zhke[3] = true
			var zpbl uint32
			zpbl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zpbl > 0 {
				z.Snot = make(map[string]bool, zpbl)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zpbl > 0 {
				zpbl--
				var zcfk string
				var zdec bool
				zcfk, err = dc.ReadString()
				if err != nil {
					return
				}
				zdec, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zcfk] = zdec
			}
		case "Notyet_zid03_map":
			found0zhke[4] = true
			var zewd uint32
			zewd, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zewd > 0 {
				z.Notyet = make(map[string]bool, zewd)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zewd > 0 {
				zewd--
				var zcfe string
				var zoks bool
				zcfe, err = dc.ReadString()
				if err != nil {
					return
				}
				zoks, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zcfe] = zoks
			}
		case "Setm_zid04_slc":
			found0zhke[5] = true
			var zcri uint32
			zcri, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zcri) {
				z.Setm = (z.Setm)[:zcri]
			} else {
				z.Setm = make([]*inn, zcri)
			}
			for zrqw := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zrqw] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zrqw].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zrqw] == nil {
						z.Setm[zrqw] = new(inn)
					}
					err = z.Setm[zrqw].DecodeMsg(dc)
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
	if nextMiss0zhke != -1 {
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
var decodeMsgFieldOrder0zhke = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zhke = []bool{false, true, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Tr) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Tr) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zsdr [6]bool
	fieldsInUse_zpfp := z.fieldsNotEmpty(empty_zsdr[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpfp)
	if err != nil {
		return err
	}

	// write "U_zid00_str"
	err = en.Append(0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.U)
	if err != nil {
		return
	}
	// write "Nt_zid01_int"
	err = en.Append(0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Nt)
	if err != nil {
		return
	}
	// write "Snot_zid02_map"
	err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Snot)))
	if err != nil {
		return
	}
	for zcfk, zdec := range z.Snot {
		err = en.WriteString(zcfk)
		if err != nil {
			return
		}
		err = en.WriteBool(zdec)
		if err != nil {
			return
		}
	}
	// write "Notyet_zid03_map"
	err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Notyet)))
	if err != nil {
		return
	}
	for zcfe, zoks := range z.Notyet {
		err = en.WriteString(zcfe)
		if err != nil {
			return
		}
		err = en.WriteBool(zoks)
		if err != nil {
			return
		}
	}
	// write "Setm_zid04_slc"
	err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Setm)))
	if err != nil {
		return
	}
	for zrqw := range z.Setm {
		if z.Setm[zrqw] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Setm[zrqw].EncodeMsg(en)
			if err != nil {
				return
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

	// string "U_zid00_str"
	o = append(o, 0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.U)
	// string "Nt_zid01_int"
	o = append(o, 0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.Nt)
	// string "Snot_zid02_map"
	o = append(o, 0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
	for zcfk, zdec := range z.Snot {
		o = msgp.AppendString(o, zcfk)
		o = msgp.AppendBool(o, zdec)
	}
	// string "Notyet_zid03_map"
	o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
	for zcfe, zoks := range z.Notyet {
		o = msgp.AppendString(o, zcfe)
		o = msgp.AppendBool(o, zoks)
	}
	// string "Setm_zid04_slc"
	o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
	for zrqw := range z.Setm {
		if z.Setm[zrqw] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Setm[zrqw].MarshalMsg(o)
			if err != nil {
				return
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
	const maxFields1ztkr = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1ztkr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1ztkr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1ztkr := totalEncodedFields1ztkr
	missingFieldsLeft1ztkr := maxFields1ztkr - totalEncodedFields1ztkr

	var nextMiss1ztkr int32 = -1
	var found1ztkr [maxFields1ztkr]bool
	var curField1ztkr string

doneWithStruct1ztkr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1ztkr > 0 || missingFieldsLeft1ztkr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1ztkr, missingFieldsLeft1ztkr, msgp.ShowFound(found1ztkr[:]), unmarshalMsgFieldOrder1ztkr)
		if encodedFieldsLeft1ztkr > 0 {
			encodedFieldsLeft1ztkr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1ztkr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1ztkr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1ztkr = 0
			}
			for nextMiss1ztkr < maxFields1ztkr && (found1ztkr[nextMiss1ztkr] || unmarshalMsgFieldSkip1ztkr[nextMiss1ztkr]) {
				nextMiss1ztkr++
			}
			if nextMiss1ztkr == maxFields1ztkr {
				// filled all the empty fields!
				break doneWithStruct1ztkr
			}
			missingFieldsLeft1ztkr--
			curField1ztkr = unmarshalMsgFieldOrder1ztkr[nextMiss1ztkr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1ztkr)
		switch curField1ztkr {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1ztkr[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1ztkr[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1ztkr[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zwmc uint32
				zwmc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zwmc > 0 {
					z.Snot = make(map[string]bool, zwmc)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zwmc > 0 {
					var zcfk string
					var zdec bool
					zwmc--
					zcfk, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zdec, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zcfk] = zdec
				}
			}
		case "Notyet_zid03_map":
			found1ztkr[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zzdg uint32
				zzdg, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zzdg > 0 {
					z.Notyet = make(map[string]bool, zzdg)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zzdg > 0 {
					var zcfe string
					var zoks bool
					zzdg--
					zcfe, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zoks, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zcfe] = zoks
				}
			}
		case "Setm_zid04_slc":
			found1ztkr[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zibu uint32
				zibu, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zibu) {
					z.Setm = (z.Setm)[:zibu]
				} else {
					z.Setm = make([]*inn, zibu)
				}
				for zrqw := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zrqw] != nil {
							z.Setm[zrqw].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zrqw] {
								z.Setm[zrqw].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zrqw] == nil {
								z.Setm[zrqw] = new(inn)
							}
							bts, err = z.Setm[zrqw].UnmarshalMsg(bts)
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
	if nextMiss1ztkr != -1 {
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
var unmarshalMsgFieldOrder1ztkr = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1ztkr = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zcfk, zdec := range z.Snot {
			_ = zdec
			_ = zcfk
			s += msgp.StringPrefixSize + len(zcfk) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zcfe, zoks := range z.Notyet {
			_ = zoks
			_ = zcfe
			s += msgp.StringPrefixSize + len(zcfe) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zrqw := range z.Setm {
		if z.Setm[zrqw] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zrqw].Msgsize()
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
	const maxFields2zeab = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zeab uint32
	totalEncodedFields2zeab, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zeab := totalEncodedFields2zeab
	missingFieldsLeft2zeab := maxFields2zeab - totalEncodedFields2zeab

	var nextMiss2zeab int32 = -1
	var found2zeab [maxFields2zeab]bool
	var curField2zeab string

doneWithStruct2zeab:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zeab > 0 || missingFieldsLeft2zeab > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zeab, missingFieldsLeft2zeab, msgp.ShowFound(found2zeab[:]), decodeMsgFieldOrder2zeab)
		if encodedFieldsLeft2zeab > 0 {
			encodedFieldsLeft2zeab--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zeab = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zeab < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zeab = 0
			}
			for nextMiss2zeab < maxFields2zeab && (found2zeab[nextMiss2zeab] || decodeMsgFieldSkip2zeab[nextMiss2zeab]) {
				nextMiss2zeab++
			}
			if nextMiss2zeab == maxFields2zeab {
				// filled all the empty fields!
				break doneWithStruct2zeab
			}
			missingFieldsLeft2zeab--
			curField2zeab = decodeMsgFieldOrder2zeab[nextMiss2zeab]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zeab)
		switch curField2zeab {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zeab[0] = true
			var zhri uint32
			zhri, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zhri > 0 {
				z.j = make(map[string]bool, zhri)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zhri > 0 {
				zhri--
				var zdcg string
				var zewi bool
				zdcg, err = dc.ReadString()
				if err != nil {
					return
				}
				zewi, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zdcg] = zewi
			}
		case "e_zid01_i64":
			found2zeab[1] = true
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
	if nextMiss2zeab != -1 {
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
var decodeMsgFieldOrder2zeab = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zeab = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *inn) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *inn) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zmul [2]bool
	fieldsInUse_ziub := z.fieldsNotEmpty(empty_zmul[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_ziub)
	if err != nil {
		return err
	}

	// write "j_zid00_map"
	err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.j)))
	if err != nil {
		return
	}
	for zdcg, zewi := range z.j {
		err = en.WriteString(zdcg)
		if err != nil {
			return
		}
		err = en.WriteBool(zewi)
		if err != nil {
			return
		}
	}
	// write "e_zid01_i64"
	err = en.Append(0xab, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x36, 0x34)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.e)
	if err != nil {
		return
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

	// string "j_zid00_map"
	o = append(o, 0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.j)))
	for zdcg, zewi := range z.j {
		o = msgp.AppendString(o, zdcg)
		o = msgp.AppendBool(o, zewi)
	}
	// string "e_zid01_i64"
	o = append(o, 0xab, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.e)
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
	const maxFields3zphn = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zphn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zphn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zphn := totalEncodedFields3zphn
	missingFieldsLeft3zphn := maxFields3zphn - totalEncodedFields3zphn

	var nextMiss3zphn int32 = -1
	var found3zphn [maxFields3zphn]bool
	var curField3zphn string

doneWithStruct3zphn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zphn > 0 || missingFieldsLeft3zphn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zphn, missingFieldsLeft3zphn, msgp.ShowFound(found3zphn[:]), unmarshalMsgFieldOrder3zphn)
		if encodedFieldsLeft3zphn > 0 {
			encodedFieldsLeft3zphn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zphn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zphn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zphn = 0
			}
			for nextMiss3zphn < maxFields3zphn && (found3zphn[nextMiss3zphn] || unmarshalMsgFieldSkip3zphn[nextMiss3zphn]) {
				nextMiss3zphn++
			}
			if nextMiss3zphn == maxFields3zphn {
				// filled all the empty fields!
				break doneWithStruct3zphn
			}
			missingFieldsLeft3zphn--
			curField3zphn = unmarshalMsgFieldOrder3zphn[nextMiss3zphn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zphn)
		switch curField3zphn {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zphn[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zirc uint32
				zirc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zirc > 0 {
					z.j = make(map[string]bool, zirc)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zirc > 0 {
					var zdcg string
					var zewi bool
					zirc--
					zdcg, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zewi, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zdcg] = zewi
				}
			}
		case "e_zid01_i64":
			found3zphn[1] = true
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
	if nextMiss3zphn != -1 {
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
var unmarshalMsgFieldOrder3zphn = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zphn = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zdcg, zewi := range z.j {
			_ = zewi
			_ = zdcg
			s += msgp.StringPrefixSize + len(zdcg) + msgp.BoolSize
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
	const maxFields4zvww = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zvww uint32
	totalEncodedFields4zvww, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zvww := totalEncodedFields4zvww
	missingFieldsLeft4zvww := maxFields4zvww - totalEncodedFields4zvww

	var nextMiss4zvww int32 = -1
	var found4zvww [maxFields4zvww]bool
	var curField4zvww string

doneWithStruct4zvww:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zvww > 0 || missingFieldsLeft4zvww > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zvww, missingFieldsLeft4zvww, msgp.ShowFound(found4zvww[:]), decodeMsgFieldOrder4zvww)
		if encodedFieldsLeft4zvww > 0 {
			encodedFieldsLeft4zvww--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zvww = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zvww < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zvww = 0
			}
			for nextMiss4zvww < maxFields4zvww && (found4zvww[nextMiss4zvww] || decodeMsgFieldSkip4zvww[nextMiss4zvww]) {
				nextMiss4zvww++
			}
			if nextMiss4zvww == maxFields4zvww {
				// filled all the empty fields!
				break doneWithStruct4zvww
			}
			missingFieldsLeft4zvww--
			curField4zvww = decodeMsgFieldOrder4zvww[nextMiss4zvww]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zvww)
		switch curField4zvww {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zvww[0] = true
			var zqpb uint32
			zqpb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zqpb > 0 {
				z.m = make(map[string]*Tr, zqpb)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zqpb > 0 {
				zqpb--
				var znrd string
				var zbqo *Tr
				znrd, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zbqo != nil {
						dc.PushAlwaysNil()
						err = zbqo.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zbqo == nil {
						zbqo = new(Tr)
					}
					err = zbqo.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[znrd] = zbqo
			}
		case "s_zid01_str":
			found4zvww[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zvww[2] = true
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
	if nextMiss4zvww != -1 {
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
var decodeMsgFieldOrder4zvww = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zvww = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *u) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *u) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zhnl [3]bool
	fieldsInUse_znjm := z.fieldsNotEmpty(empty_zhnl[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_znjm)
	if err != nil {
		return err
	}

	// write "m_zid00_map"
	err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.m)))
	if err != nil {
		return
	}
	for znrd, zbqo := range z.m {
		err = en.WriteString(znrd)
		if err != nil {
			return
		}
		if zbqo == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zbqo.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "s_zid01_str"
	err = en.Append(0xab, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.s)
	if err != nil {
		return
	}
	// write "n_zid02_i64"
	err = en.Append(0xab, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x36, 0x34)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.n)
	if err != nil {
		return
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

	// string "m_zid00_map"
	o = append(o, 0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.m)))
	for znrd, zbqo := range z.m {
		o = msgp.AppendString(o, znrd)
		if zbqo == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zbqo.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "s_zid01_str"
	o = append(o, 0xab, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.s)
	// string "n_zid02_i64"
	o = append(o, 0xab, 0x6e, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x69, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.n)
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
	const maxFields5zada = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zada uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zada, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zada := totalEncodedFields5zada
	missingFieldsLeft5zada := maxFields5zada - totalEncodedFields5zada

	var nextMiss5zada int32 = -1
	var found5zada [maxFields5zada]bool
	var curField5zada string

doneWithStruct5zada:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zada > 0 || missingFieldsLeft5zada > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zada, missingFieldsLeft5zada, msgp.ShowFound(found5zada[:]), unmarshalMsgFieldOrder5zada)
		if encodedFieldsLeft5zada > 0 {
			encodedFieldsLeft5zada--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zada = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zada < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zada = 0
			}
			for nextMiss5zada < maxFields5zada && (found5zada[nextMiss5zada] || unmarshalMsgFieldSkip5zada[nextMiss5zada]) {
				nextMiss5zada++
			}
			if nextMiss5zada == maxFields5zada {
				// filled all the empty fields!
				break doneWithStruct5zada
			}
			missingFieldsLeft5zada--
			curField5zada = unmarshalMsgFieldOrder5zada[nextMiss5zada]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zada)
		switch curField5zada {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zada[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zswu uint32
				zswu, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zswu > 0 {
					z.m = make(map[string]*Tr, zswu)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zswu > 0 {
					var znrd string
					var zbqo *Tr
					zswu--
					znrd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zbqo != nil {
							zbqo.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zbqo {
								zbqo.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zbqo == nil {
								zbqo = new(Tr)
							}
							bts, err = zbqo.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[znrd] = zbqo
				}
			}
		case "s_zid01_str":
			found5zada[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zada[2] = true
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
	if nextMiss5zada != -1 {
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
var unmarshalMsgFieldOrder5zada = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zada = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for znrd, zbqo := range z.m {
			_ = zbqo
			_ = znrd
			s += msgp.StringPrefixSize + len(znrd)
			if zbqo == nil {
				s += msgp.NilSize
			} else {
				s += zbqo.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
