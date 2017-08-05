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
	const maxFields0zwmc = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zwmc uint32
	totalEncodedFields0zwmc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zwmc := totalEncodedFields0zwmc
	missingFieldsLeft0zwmc := maxFields0zwmc - totalEncodedFields0zwmc

	var nextMiss0zwmc int32 = -1
	var found0zwmc [maxFields0zwmc]bool
	var curField0zwmc string

doneWithStruct0zwmc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zwmc > 0 || missingFieldsLeft0zwmc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zwmc, missingFieldsLeft0zwmc, msgp.ShowFound(found0zwmc[:]), decodeMsgFieldOrder0zwmc)
		if encodedFieldsLeft0zwmc > 0 {
			encodedFieldsLeft0zwmc--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zwmc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zwmc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zwmc = 0
			}
			for nextMiss0zwmc < maxFields0zwmc && (found0zwmc[nextMiss0zwmc] || decodeMsgFieldSkip0zwmc[nextMiss0zwmc]) {
				nextMiss0zwmc++
			}
			if nextMiss0zwmc == maxFields0zwmc {
				// filled all the empty fields!
				break doneWithStruct0zwmc
			}
			missingFieldsLeft0zwmc--
			curField0zwmc = decodeMsgFieldOrder0zwmc[nextMiss0zwmc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zwmc)
		switch curField0zwmc {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zwmc[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zwmc[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zwmc[3] = true
			var zpmd uint32
			zpmd, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zpmd > 0 {
				z.Snot = make(map[string]bool, zpmd)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zpmd > 0 {
				zpmd--
				var zdbc string
				var zvfu bool
				zdbc, err = dc.ReadString()
				if err != nil {
					return
				}
				zvfu, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zdbc] = zvfu
			}
		case "Notyet_zid03_map":
			found0zwmc[4] = true
			var zhsy uint32
			zhsy, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zhsy > 0 {
				z.Notyet = make(map[string]bool, zhsy)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zhsy > 0 {
				zhsy--
				var zmwt string
				var zred bool
				zmwt, err = dc.ReadString()
				if err != nil {
					return
				}
				zred, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zmwt] = zred
			}
		case "Setm_zid04_slc":
			found0zwmc[5] = true
			var zumf uint32
			zumf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zumf) {
				z.Setm = (z.Setm)[:zumf]
			} else {
				z.Setm = make([]*inn, zumf)
			}
			for zcmj := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zcmj] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zcmj].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zcmj] == nil {
						z.Setm[zcmj] = new(inn)
					}
					err = z.Setm[zcmj].DecodeMsg(dc)
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
	if nextMiss0zwmc != -1 {
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
var decodeMsgFieldOrder0zwmc = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zwmc = []bool{false, true, false, false, false, false}

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
	var empty_zaht [6]bool
	fieldsInUse_zxrm := z.fieldsNotEmpty(empty_zaht[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxrm)
	if err != nil {
		return err
	}

	if !empty_zaht[0] {
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

	if !empty_zaht[2] {
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

	if !empty_zaht[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zdbc, zvfu := range z.Snot {
			err = en.WriteString(zdbc)
			if err != nil {
				return
			}
			err = en.WriteBool(zvfu)
			if err != nil {
				return
			}
		}
	}

	if !empty_zaht[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zmwt, zred := range z.Notyet {
			err = en.WriteString(zmwt)
			if err != nil {
				return
			}
			err = en.WriteBool(zred)
			if err != nil {
				return
			}
		}
	}

	if !empty_zaht[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zcmj := range z.Setm {
			if z.Setm[zcmj] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zcmj].EncodeMsg(en)
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
		for zdbc, zvfu := range z.Snot {
			o = msgp.AppendString(o, zdbc)
			o = msgp.AppendBool(o, zvfu)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zmwt, zred := range z.Notyet {
			o = msgp.AppendString(o, zmwt)
			o = msgp.AppendBool(o, zred)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zcmj := range z.Setm {
			if z.Setm[zcmj] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zcmj].MarshalMsg(o)
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
	const maxFields1zreg = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zreg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zreg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zreg := totalEncodedFields1zreg
	missingFieldsLeft1zreg := maxFields1zreg - totalEncodedFields1zreg

	var nextMiss1zreg int32 = -1
	var found1zreg [maxFields1zreg]bool
	var curField1zreg string

doneWithStruct1zreg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zreg > 0 || missingFieldsLeft1zreg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zreg, missingFieldsLeft1zreg, msgp.ShowFound(found1zreg[:]), unmarshalMsgFieldOrder1zreg)
		if encodedFieldsLeft1zreg > 0 {
			encodedFieldsLeft1zreg--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zreg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zreg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zreg = 0
			}
			for nextMiss1zreg < maxFields1zreg && (found1zreg[nextMiss1zreg] || unmarshalMsgFieldSkip1zreg[nextMiss1zreg]) {
				nextMiss1zreg++
			}
			if nextMiss1zreg == maxFields1zreg {
				// filled all the empty fields!
				break doneWithStruct1zreg
			}
			missingFieldsLeft1zreg--
			curField1zreg = unmarshalMsgFieldOrder1zreg[nextMiss1zreg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zreg)
		switch curField1zreg {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zreg[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zreg[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zreg[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var znpe uint32
				znpe, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && znpe > 0 {
					z.Snot = make(map[string]bool, znpe)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for znpe > 0 {
					var zdbc string
					var zvfu bool
					znpe--
					zdbc, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zvfu, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zdbc] = zvfu
				}
			}
		case "Notyet_zid03_map":
			found1zreg[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zulm uint32
				zulm, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zulm > 0 {
					z.Notyet = make(map[string]bool, zulm)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zulm > 0 {
					var zmwt string
					var zred bool
					zulm--
					zmwt, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zred, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zmwt] = zred
				}
			}
		case "Setm_zid04_slc":
			found1zreg[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zydz uint32
				zydz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zydz) {
					z.Setm = (z.Setm)[:zydz]
				} else {
					z.Setm = make([]*inn, zydz)
				}
				for zcmj := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zcmj] != nil {
							z.Setm[zcmj].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zcmj] {
								z.Setm[zcmj].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zcmj] == nil {
								z.Setm[zcmj] = new(inn)
							}
							bts, err = z.Setm[zcmj].UnmarshalMsg(bts)
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
	if nextMiss1zreg != -1 {
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
var unmarshalMsgFieldOrder1zreg = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zreg = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zdbc, zvfu := range z.Snot {
			_ = zvfu
			_ = zdbc
			s += msgp.StringPrefixSize + len(zdbc) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zmwt, zred := range z.Notyet {
			_ = zred
			_ = zmwt
			s += msgp.StringPrefixSize + len(zmwt) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zcmj := range z.Setm {
		if z.Setm[zcmj] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zcmj].Msgsize()
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
	const maxFields2zylh = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zylh uint32
	totalEncodedFields2zylh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zylh := totalEncodedFields2zylh
	missingFieldsLeft2zylh := maxFields2zylh - totalEncodedFields2zylh

	var nextMiss2zylh int32 = -1
	var found2zylh [maxFields2zylh]bool
	var curField2zylh string

doneWithStruct2zylh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zylh > 0 || missingFieldsLeft2zylh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zylh, missingFieldsLeft2zylh, msgp.ShowFound(found2zylh[:]), decodeMsgFieldOrder2zylh)
		if encodedFieldsLeft2zylh > 0 {
			encodedFieldsLeft2zylh--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zylh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zylh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zylh = 0
			}
			for nextMiss2zylh < maxFields2zylh && (found2zylh[nextMiss2zylh] || decodeMsgFieldSkip2zylh[nextMiss2zylh]) {
				nextMiss2zylh++
			}
			if nextMiss2zylh == maxFields2zylh {
				// filled all the empty fields!
				break doneWithStruct2zylh
			}
			missingFieldsLeft2zylh--
			curField2zylh = decodeMsgFieldOrder2zylh[nextMiss2zylh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zylh)
		switch curField2zylh {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zylh[0] = true
			var zubu uint32
			zubu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zubu > 0 {
				z.j = make(map[string]bool, zubu)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zubu > 0 {
				zubu--
				var zhlf string
				var zdsm bool
				zhlf, err = dc.ReadString()
				if err != nil {
					return
				}
				zdsm, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zhlf] = zdsm
			}
		case "e_zid01_i64":
			found2zylh[1] = true
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
	if nextMiss2zylh != -1 {
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
var decodeMsgFieldOrder2zylh = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zylh = []bool{false, false}

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
	var empty_zkzp [2]bool
	fieldsInUse_zelx := z.fieldsNotEmpty(empty_zkzp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zelx)
	if err != nil {
		return err
	}

	if !empty_zkzp[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zhlf, zdsm := range z.j {
			err = en.WriteString(zhlf)
			if err != nil {
				return
			}
			err = en.WriteBool(zdsm)
			if err != nil {
				return
			}
		}
	}

	if !empty_zkzp[1] {
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
		for zhlf, zdsm := range z.j {
			o = msgp.AppendString(o, zhlf)
			o = msgp.AppendBool(o, zdsm)
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
	const maxFields3zsxi = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zsxi uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zsxi, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zsxi := totalEncodedFields3zsxi
	missingFieldsLeft3zsxi := maxFields3zsxi - totalEncodedFields3zsxi

	var nextMiss3zsxi int32 = -1
	var found3zsxi [maxFields3zsxi]bool
	var curField3zsxi string

doneWithStruct3zsxi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zsxi > 0 || missingFieldsLeft3zsxi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zsxi, missingFieldsLeft3zsxi, msgp.ShowFound(found3zsxi[:]), unmarshalMsgFieldOrder3zsxi)
		if encodedFieldsLeft3zsxi > 0 {
			encodedFieldsLeft3zsxi--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zsxi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zsxi < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zsxi = 0
			}
			for nextMiss3zsxi < maxFields3zsxi && (found3zsxi[nextMiss3zsxi] || unmarshalMsgFieldSkip3zsxi[nextMiss3zsxi]) {
				nextMiss3zsxi++
			}
			if nextMiss3zsxi == maxFields3zsxi {
				// filled all the empty fields!
				break doneWithStruct3zsxi
			}
			missingFieldsLeft3zsxi--
			curField3zsxi = unmarshalMsgFieldOrder3zsxi[nextMiss3zsxi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zsxi)
		switch curField3zsxi {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zsxi[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zvsk uint32
				zvsk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zvsk > 0 {
					z.j = make(map[string]bool, zvsk)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zvsk > 0 {
					var zhlf string
					var zdsm bool
					zvsk--
					zhlf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zdsm, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zhlf] = zdsm
				}
			}
		case "e_zid01_i64":
			found3zsxi[1] = true
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
	if nextMiss3zsxi != -1 {
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
var unmarshalMsgFieldOrder3zsxi = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zsxi = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zhlf, zdsm := range z.j {
			_ = zdsm
			_ = zhlf
			s += msgp.StringPrefixSize + len(zhlf) + msgp.BoolSize
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
	const maxFields4zedj = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zedj uint32
	totalEncodedFields4zedj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zedj := totalEncodedFields4zedj
	missingFieldsLeft4zedj := maxFields4zedj - totalEncodedFields4zedj

	var nextMiss4zedj int32 = -1
	var found4zedj [maxFields4zedj]bool
	var curField4zedj string

doneWithStruct4zedj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zedj > 0 || missingFieldsLeft4zedj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zedj, missingFieldsLeft4zedj, msgp.ShowFound(found4zedj[:]), decodeMsgFieldOrder4zedj)
		if encodedFieldsLeft4zedj > 0 {
			encodedFieldsLeft4zedj--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zedj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zedj < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zedj = 0
			}
			for nextMiss4zedj < maxFields4zedj && (found4zedj[nextMiss4zedj] || decodeMsgFieldSkip4zedj[nextMiss4zedj]) {
				nextMiss4zedj++
			}
			if nextMiss4zedj == maxFields4zedj {
				// filled all the empty fields!
				break doneWithStruct4zedj
			}
			missingFieldsLeft4zedj--
			curField4zedj = decodeMsgFieldOrder4zedj[nextMiss4zedj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zedj)
		switch curField4zedj {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zedj[0] = true
			var zicn uint32
			zicn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zicn > 0 {
				z.m = make(map[string]*Tr, zicn)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zicn > 0 {
				zicn--
				var zegs string
				var zzys *Tr
				zegs, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zzys != nil {
						dc.PushAlwaysNil()
						err = zzys.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zzys == nil {
						zzys = new(Tr)
					}
					err = zzys.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zegs] = zzys
			}
		case "s_zid01_str":
			found4zedj[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zedj[2] = true
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
	if nextMiss4zedj != -1 {
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
var decodeMsgFieldOrder4zedj = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zedj = []bool{false, false, false}

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
	var empty_zysm [3]bool
	fieldsInUse_zrew := z.fieldsNotEmpty(empty_zysm[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zrew)
	if err != nil {
		return err
	}

	if !empty_zysm[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zegs, zzys := range z.m {
			err = en.WriteString(zegs)
			if err != nil {
				return
			}
			if zzys == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zzys.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zysm[1] {
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

	if !empty_zysm[2] {
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
		for zegs, zzys := range z.m {
			o = msgp.AppendString(o, zegs)
			if zzys == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zzys.MarshalMsg(o)
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
	const maxFields5zgmz = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zgmz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zgmz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zgmz := totalEncodedFields5zgmz
	missingFieldsLeft5zgmz := maxFields5zgmz - totalEncodedFields5zgmz

	var nextMiss5zgmz int32 = -1
	var found5zgmz [maxFields5zgmz]bool
	var curField5zgmz string

doneWithStruct5zgmz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zgmz > 0 || missingFieldsLeft5zgmz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zgmz, missingFieldsLeft5zgmz, msgp.ShowFound(found5zgmz[:]), unmarshalMsgFieldOrder5zgmz)
		if encodedFieldsLeft5zgmz > 0 {
			encodedFieldsLeft5zgmz--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zgmz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zgmz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zgmz = 0
			}
			for nextMiss5zgmz < maxFields5zgmz && (found5zgmz[nextMiss5zgmz] || unmarshalMsgFieldSkip5zgmz[nextMiss5zgmz]) {
				nextMiss5zgmz++
			}
			if nextMiss5zgmz == maxFields5zgmz {
				// filled all the empty fields!
				break doneWithStruct5zgmz
			}
			missingFieldsLeft5zgmz--
			curField5zgmz = unmarshalMsgFieldOrder5zgmz[nextMiss5zgmz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zgmz)
		switch curField5zgmz {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zgmz[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zeaa uint32
				zeaa, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zeaa > 0 {
					z.m = make(map[string]*Tr, zeaa)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zeaa > 0 {
					var zegs string
					var zzys *Tr
					zeaa--
					zegs, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zzys != nil {
							zzys.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zzys {
								zzys.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zzys == nil {
								zzys = new(Tr)
							}
							bts, err = zzys.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zegs] = zzys
				}
			}
		case "s_zid01_str":
			found5zgmz[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zgmz[2] = true
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
	if nextMiss5zgmz != -1 {
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
var unmarshalMsgFieldOrder5zgmz = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zgmz = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zegs, zzys := range z.m {
			_ = zzys
			_ = zegs
			s += msgp.StringPrefixSize + len(zegs)
			if zzys == nil {
				s += msgp.NilSize
			} else {
				s += zzys.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
