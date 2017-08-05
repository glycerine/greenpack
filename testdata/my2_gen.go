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
	const maxFields0zqgi = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zqgi uint32
	totalEncodedFields0zqgi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zqgi := totalEncodedFields0zqgi
	missingFieldsLeft0zqgi := maxFields0zqgi - totalEncodedFields0zqgi

	var nextMiss0zqgi int32 = -1
	var found0zqgi [maxFields0zqgi]bool
	var curField0zqgi string

doneWithStruct0zqgi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zqgi > 0 || missingFieldsLeft0zqgi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zqgi, missingFieldsLeft0zqgi, msgp.ShowFound(found0zqgi[:]), decodeMsgFieldOrder0zqgi)
		if encodedFieldsLeft0zqgi > 0 {
			encodedFieldsLeft0zqgi--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zqgi = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zqgi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zqgi = 0
			}
			for nextMiss0zqgi < maxFields0zqgi && (found0zqgi[nextMiss0zqgi] || decodeMsgFieldSkip0zqgi[nextMiss0zqgi]) {
				nextMiss0zqgi++
			}
			if nextMiss0zqgi == maxFields0zqgi {
				// filled all the empty fields!
				break doneWithStruct0zqgi
			}
			missingFieldsLeft0zqgi--
			curField0zqgi = decodeMsgFieldOrder0zqgi[nextMiss0zqgi]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zqgi)
		switch curField0zqgi {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zqgi[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zqgi[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zqgi[3] = true
			var zkzc uint32
			zkzc, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zkzc > 0 {
				z.Snot = make(map[string]bool, zkzc)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zkzc > 0 {
				zkzc--
				var ztlq string
				var zllt bool
				ztlq, err = dc.ReadString()
				if err != nil {
					return
				}
				zllt, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[ztlq] = zllt
			}
		case "Notyet_zid03_map":
			found0zqgi[4] = true
			var zvrj uint32
			zvrj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zvrj > 0 {
				z.Notyet = make(map[string]bool, zvrj)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zvrj > 0 {
				zvrj--
				var zpkg string
				var zmvf bool
				zpkg, err = dc.ReadString()
				if err != nil {
					return
				}
				zmvf, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zpkg] = zmvf
			}
		case "Setm_zid04_slc":
			found0zqgi[5] = true
			var ztcs uint32
			ztcs, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(ztcs) {
				z.Setm = (z.Setm)[:ztcs]
			} else {
				z.Setm = make([]*inn, ztcs)
			}
			for zhxv := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zhxv] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zhxv].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zhxv] == nil {
						z.Setm[zhxv] = new(inn)
					}
					err = z.Setm[zhxv].DecodeMsg(dc)
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
	if nextMiss0zqgi != -1 {
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
var decodeMsgFieldOrder0zqgi = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zqgi = []bool{false, true, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Tr) fieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// EncodeMsg implements msgp.Encodable
func (z *Tr) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 5
	// write "U_zid00_str"
	err = en.Append(0x85, 0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
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
	for ztlq, zllt := range z.Snot {
		err = en.WriteString(ztlq)
		if err != nil {
			return
		}
		err = en.WriteBool(zllt)
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
	for zpkg, zmvf := range z.Notyet {
		err = en.WriteString(zpkg)
		if err != nil {
			return
		}
		err = en.WriteBool(zmvf)
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
	for zhxv := range z.Setm {
		if z.Setm[zhxv] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Setm[zhxv].EncodeMsg(en)
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
	// map header, size 5
	// string "U_zid00_str"
	o = append(o, 0x85, 0xab, 0x55, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.U)
	// string "Nt_zid01_int"
	o = append(o, 0xac, 0x4e, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x69, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.Nt)
	// string "Snot_zid02_map"
	o = append(o, 0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
	for ztlq, zllt := range z.Snot {
		o = msgp.AppendString(o, ztlq)
		o = msgp.AppendBool(o, zllt)
	}
	// string "Notyet_zid03_map"
	o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
	for zpkg, zmvf := range z.Notyet {
		o = msgp.AppendString(o, zpkg)
		o = msgp.AppendBool(o, zmvf)
	}
	// string "Setm_zid04_slc"
	o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
	for zhxv := range z.Setm {
		if z.Setm[zhxv] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Setm[zhxv].MarshalMsg(o)
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
	const maxFields1ziwf = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1ziwf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1ziwf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1ziwf := totalEncodedFields1ziwf
	missingFieldsLeft1ziwf := maxFields1ziwf - totalEncodedFields1ziwf

	var nextMiss1ziwf int32 = -1
	var found1ziwf [maxFields1ziwf]bool
	var curField1ziwf string

doneWithStruct1ziwf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1ziwf > 0 || missingFieldsLeft1ziwf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1ziwf, missingFieldsLeft1ziwf, msgp.ShowFound(found1ziwf[:]), unmarshalMsgFieldOrder1ziwf)
		if encodedFieldsLeft1ziwf > 0 {
			encodedFieldsLeft1ziwf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1ziwf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1ziwf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1ziwf = 0
			}
			for nextMiss1ziwf < maxFields1ziwf && (found1ziwf[nextMiss1ziwf] || unmarshalMsgFieldSkip1ziwf[nextMiss1ziwf]) {
				nextMiss1ziwf++
			}
			if nextMiss1ziwf == maxFields1ziwf {
				// filled all the empty fields!
				break doneWithStruct1ziwf
			}
			missingFieldsLeft1ziwf--
			curField1ziwf = unmarshalMsgFieldOrder1ziwf[nextMiss1ziwf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1ziwf)
		switch curField1ziwf {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1ziwf[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1ziwf[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1ziwf[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zfpf uint32
				zfpf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zfpf > 0 {
					z.Snot = make(map[string]bool, zfpf)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zfpf > 0 {
					var ztlq string
					var zllt bool
					zfpf--
					ztlq, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zllt, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[ztlq] = zllt
				}
			}
		case "Notyet_zid03_map":
			found1ziwf[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zllx uint32
				zllx, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zllx > 0 {
					z.Notyet = make(map[string]bool, zllx)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zllx > 0 {
					var zpkg string
					var zmvf bool
					zllx--
					zpkg, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zmvf, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zpkg] = zmvf
				}
			}
		case "Setm_zid04_slc":
			found1ziwf[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zzix uint32
				zzix, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zzix) {
					z.Setm = (z.Setm)[:zzix]
				} else {
					z.Setm = make([]*inn, zzix)
				}
				for zhxv := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zhxv] != nil {
							z.Setm[zhxv].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zhxv] {
								z.Setm[zhxv].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zhxv] == nil {
								z.Setm[zhxv] = new(inn)
							}
							bts, err = z.Setm[zhxv].UnmarshalMsg(bts)
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
	if nextMiss1ziwf != -1 {
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
var unmarshalMsgFieldOrder1ziwf = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1ziwf = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for ztlq, zllt := range z.Snot {
			_ = zllt
			_ = ztlq
			s += msgp.StringPrefixSize + len(ztlq) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zpkg, zmvf := range z.Notyet {
			_ = zmvf
			_ = zpkg
			s += msgp.StringPrefixSize + len(zpkg) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zhxv := range z.Setm {
		if z.Setm[zhxv] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zhxv].Msgsize()
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
	const maxFields2zwzb = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zwzb uint32
	totalEncodedFields2zwzb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zwzb := totalEncodedFields2zwzb
	missingFieldsLeft2zwzb := maxFields2zwzb - totalEncodedFields2zwzb

	var nextMiss2zwzb int32 = -1
	var found2zwzb [maxFields2zwzb]bool
	var curField2zwzb string

doneWithStruct2zwzb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zwzb > 0 || missingFieldsLeft2zwzb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zwzb, missingFieldsLeft2zwzb, msgp.ShowFound(found2zwzb[:]), decodeMsgFieldOrder2zwzb)
		if encodedFieldsLeft2zwzb > 0 {
			encodedFieldsLeft2zwzb--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zwzb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zwzb < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zwzb = 0
			}
			for nextMiss2zwzb < maxFields2zwzb && (found2zwzb[nextMiss2zwzb] || decodeMsgFieldSkip2zwzb[nextMiss2zwzb]) {
				nextMiss2zwzb++
			}
			if nextMiss2zwzb == maxFields2zwzb {
				// filled all the empty fields!
				break doneWithStruct2zwzb
			}
			missingFieldsLeft2zwzb--
			curField2zwzb = decodeMsgFieldOrder2zwzb[nextMiss2zwzb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zwzb)
		switch curField2zwzb {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zwzb[0] = true
			var zjyq uint32
			zjyq, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zjyq > 0 {
				z.j = make(map[string]bool, zjyq)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zjyq > 0 {
				zjyq--
				var zaar string
				var zppn bool
				zaar, err = dc.ReadString()
				if err != nil {
					return
				}
				zppn, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zaar] = zppn
			}
		case "e_zid01_i64":
			found2zwzb[1] = true
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
	if nextMiss2zwzb != -1 {
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
var decodeMsgFieldOrder2zwzb = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zwzb = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *inn) fieldsNotEmpty(isempty []bool) uint32 {
	return 2
}

// EncodeMsg implements msgp.Encodable
func (z *inn) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 2
	// write "j_zid00_map"
	err = en.Append(0x82, 0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.j)))
	if err != nil {
		return
	}
	for zaar, zppn := range z.j {
		err = en.WriteString(zaar)
		if err != nil {
			return
		}
		err = en.WriteBool(zppn)
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
	// map header, size 2
	// string "j_zid00_map"
	o = append(o, 0x82, 0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.j)))
	for zaar, zppn := range z.j {
		o = msgp.AppendString(o, zaar)
		o = msgp.AppendBool(o, zppn)
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
	const maxFields3zvfp = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zvfp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zvfp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zvfp := totalEncodedFields3zvfp
	missingFieldsLeft3zvfp := maxFields3zvfp - totalEncodedFields3zvfp

	var nextMiss3zvfp int32 = -1
	var found3zvfp [maxFields3zvfp]bool
	var curField3zvfp string

doneWithStruct3zvfp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zvfp > 0 || missingFieldsLeft3zvfp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvfp, missingFieldsLeft3zvfp, msgp.ShowFound(found3zvfp[:]), unmarshalMsgFieldOrder3zvfp)
		if encodedFieldsLeft3zvfp > 0 {
			encodedFieldsLeft3zvfp--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zvfp = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zvfp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zvfp = 0
			}
			for nextMiss3zvfp < maxFields3zvfp && (found3zvfp[nextMiss3zvfp] || unmarshalMsgFieldSkip3zvfp[nextMiss3zvfp]) {
				nextMiss3zvfp++
			}
			if nextMiss3zvfp == maxFields3zvfp {
				// filled all the empty fields!
				break doneWithStruct3zvfp
			}
			missingFieldsLeft3zvfp--
			curField3zvfp = unmarshalMsgFieldOrder3zvfp[nextMiss3zvfp]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zvfp)
		switch curField3zvfp {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zvfp[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zapp uint32
				zapp, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zapp > 0 {
					z.j = make(map[string]bool, zapp)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zapp > 0 {
					var zaar string
					var zppn bool
					zapp--
					zaar, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zppn, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zaar] = zppn
				}
			}
		case "e_zid01_i64":
			found3zvfp[1] = true
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
	if nextMiss3zvfp != -1 {
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
var unmarshalMsgFieldOrder3zvfp = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zvfp = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zaar, zppn := range z.j {
			_ = zppn
			_ = zaar
			s += msgp.StringPrefixSize + len(zaar) + msgp.BoolSize
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
	const maxFields4zrnv = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zrnv uint32
	totalEncodedFields4zrnv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zrnv := totalEncodedFields4zrnv
	missingFieldsLeft4zrnv := maxFields4zrnv - totalEncodedFields4zrnv

	var nextMiss4zrnv int32 = -1
	var found4zrnv [maxFields4zrnv]bool
	var curField4zrnv string

doneWithStruct4zrnv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zrnv > 0 || missingFieldsLeft4zrnv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zrnv, missingFieldsLeft4zrnv, msgp.ShowFound(found4zrnv[:]), decodeMsgFieldOrder4zrnv)
		if encodedFieldsLeft4zrnv > 0 {
			encodedFieldsLeft4zrnv--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zrnv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zrnv < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zrnv = 0
			}
			for nextMiss4zrnv < maxFields4zrnv && (found4zrnv[nextMiss4zrnv] || decodeMsgFieldSkip4zrnv[nextMiss4zrnv]) {
				nextMiss4zrnv++
			}
			if nextMiss4zrnv == maxFields4zrnv {
				// filled all the empty fields!
				break doneWithStruct4zrnv
			}
			missingFieldsLeft4zrnv--
			curField4zrnv = decodeMsgFieldOrder4zrnv[nextMiss4zrnv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zrnv)
		switch curField4zrnv {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zrnv[0] = true
			var zbrx uint32
			zbrx, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zbrx > 0 {
				z.m = make(map[string]*Tr, zbrx)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zbrx > 0 {
				zbrx--
				var zdzq string
				var zmrx *Tr
				zdzq, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zmrx != nil {
						dc.PushAlwaysNil()
						err = zmrx.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zmrx == nil {
						zmrx = new(Tr)
					}
					err = zmrx.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zdzq] = zmrx
			}
		case "s_zid01_str":
			found4zrnv[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zrnv[2] = true
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
	if nextMiss4zrnv != -1 {
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
var decodeMsgFieldOrder4zrnv = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zrnv = []bool{false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *u) fieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// EncodeMsg implements msgp.Encodable
func (z *u) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// map header, size 3
	// write "m_zid00_map"
	err = en.Append(0x83, 0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.m)))
	if err != nil {
		return
	}
	for zdzq, zmrx := range z.m {
		err = en.WriteString(zdzq)
		if err != nil {
			return
		}
		if zmrx == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zmrx.EncodeMsg(en)
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
	// map header, size 3
	// string "m_zid00_map"
	o = append(o, 0x83, 0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.m)))
	for zdzq, zmrx := range z.m {
		o = msgp.AppendString(o, zdzq)
		if zmrx == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zmrx.MarshalMsg(o)
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
	const maxFields5zvqh = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zvqh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zvqh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zvqh := totalEncodedFields5zvqh
	missingFieldsLeft5zvqh := maxFields5zvqh - totalEncodedFields5zvqh

	var nextMiss5zvqh int32 = -1
	var found5zvqh [maxFields5zvqh]bool
	var curField5zvqh string

doneWithStruct5zvqh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zvqh > 0 || missingFieldsLeft5zvqh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zvqh, missingFieldsLeft5zvqh, msgp.ShowFound(found5zvqh[:]), unmarshalMsgFieldOrder5zvqh)
		if encodedFieldsLeft5zvqh > 0 {
			encodedFieldsLeft5zvqh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zvqh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zvqh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zvqh = 0
			}
			for nextMiss5zvqh < maxFields5zvqh && (found5zvqh[nextMiss5zvqh] || unmarshalMsgFieldSkip5zvqh[nextMiss5zvqh]) {
				nextMiss5zvqh++
			}
			if nextMiss5zvqh == maxFields5zvqh {
				// filled all the empty fields!
				break doneWithStruct5zvqh
			}
			missingFieldsLeft5zvqh--
			curField5zvqh = unmarshalMsgFieldOrder5zvqh[nextMiss5zvqh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zvqh)
		switch curField5zvqh {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zvqh[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zcsg uint32
				zcsg, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zcsg > 0 {
					z.m = make(map[string]*Tr, zcsg)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zcsg > 0 {
					var zdzq string
					var zmrx *Tr
					zcsg--
					zdzq, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zmrx != nil {
							zmrx.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zmrx {
								zmrx.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zmrx == nil {
								zmrx = new(Tr)
							}
							bts, err = zmrx.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zdzq] = zmrx
				}
			}
		case "s_zid01_str":
			found5zvqh[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zvqh[2] = true
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
	if nextMiss5zvqh != -1 {
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
var unmarshalMsgFieldOrder5zvqh = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zvqh = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zdzq, zmrx := range z.m {
			_ = zmrx
			_ = zdzq
			s += msgp.StringPrefixSize + len(zdzq)
			if zmrx == nil {
				s += msgp.NilSize
			} else {
				s += zmrx.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
