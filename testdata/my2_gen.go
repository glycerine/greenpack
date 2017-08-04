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
	const maxFields0zdwr = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zdwr uint32
	totalEncodedFields0zdwr, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zdwr := totalEncodedFields0zdwr
	missingFieldsLeft0zdwr := maxFields0zdwr - totalEncodedFields0zdwr

	var nextMiss0zdwr int32 = -1
	var found0zdwr [maxFields0zdwr]bool
	var curField0zdwr string

doneWithStruct0zdwr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zdwr > 0 || missingFieldsLeft0zdwr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zdwr, missingFieldsLeft0zdwr, msgp.ShowFound(found0zdwr[:]), decodeMsgFieldOrder0zdwr)
		if encodedFieldsLeft0zdwr > 0 {
			encodedFieldsLeft0zdwr--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zdwr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zdwr < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zdwr = 0
			}
			for nextMiss0zdwr < maxFields0zdwr && (found0zdwr[nextMiss0zdwr] || decodeMsgFieldSkip0zdwr[nextMiss0zdwr]) {
				nextMiss0zdwr++
			}
			if nextMiss0zdwr == maxFields0zdwr {
				// filled all the empty fields!
				break doneWithStruct0zdwr
			}
			missingFieldsLeft0zdwr--
			curField0zdwr = decodeMsgFieldOrder0zdwr[nextMiss0zdwr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zdwr)
		switch curField0zdwr {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zdwr[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zdwr[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zdwr[3] = true
			var zmru uint32
			zmru, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zmru > 0 {
				z.Snot = make(map[string]bool, zmru)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zmru > 0 {
				zmru--
				var zfjk string
				var zend bool
				zfjk, err = dc.ReadString()
				if err != nil {
					return
				}
				zend, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zfjk] = zend
			}
		case "Notyet_zid03_map":
			found0zdwr[4] = true
			var zirq uint32
			zirq, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zirq > 0 {
				z.Notyet = make(map[string]bool, zirq)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zirq > 0 {
				zirq--
				var zcaa string
				var zrqo bool
				zcaa, err = dc.ReadString()
				if err != nil {
					return
				}
				zrqo, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zcaa] = zrqo
			}
		case "Setm_zid04_slc":
			found0zdwr[5] = true
			var znoi uint32
			znoi, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(znoi) {
				z.Setm = (z.Setm)[:znoi]
			} else {
				z.Setm = make([]*inn, znoi)
			}
			for zssu := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zssu] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zssu].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zssu] == nil {
						z.Setm[zssu] = new(inn)
					}
					err = z.Setm[zssu].DecodeMsg(dc)
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
	if nextMiss0zdwr != -1 {
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
var decodeMsgFieldOrder0zdwr = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zdwr = []bool{false, true, false, false, false, false}

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
	for zfjk, zend := range z.Snot {
		err = en.WriteString(zfjk)
		if err != nil {
			return
		}
		err = en.WriteBool(zend)
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
	for zcaa, zrqo := range z.Notyet {
		err = en.WriteString(zcaa)
		if err != nil {
			return
		}
		err = en.WriteBool(zrqo)
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
	for zssu := range z.Setm {
		if z.Setm[zssu] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Setm[zssu].EncodeMsg(en)
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
	for zfjk, zend := range z.Snot {
		o = msgp.AppendString(o, zfjk)
		o = msgp.AppendBool(o, zend)
	}
	// string "Notyet_zid03_map"
	o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
	for zcaa, zrqo := range z.Notyet {
		o = msgp.AppendString(o, zcaa)
		o = msgp.AppendBool(o, zrqo)
	}
	// string "Setm_zid04_slc"
	o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
	for zssu := range z.Setm {
		if z.Setm[zssu] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Setm[zssu].MarshalMsg(o)
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
	const maxFields1zokd = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zokd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zokd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zokd := totalEncodedFields1zokd
	missingFieldsLeft1zokd := maxFields1zokd - totalEncodedFields1zokd

	var nextMiss1zokd int32 = -1
	var found1zokd [maxFields1zokd]bool
	var curField1zokd string

doneWithStruct1zokd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zokd > 0 || missingFieldsLeft1zokd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zokd, missingFieldsLeft1zokd, msgp.ShowFound(found1zokd[:]), unmarshalMsgFieldOrder1zokd)
		if encodedFieldsLeft1zokd > 0 {
			encodedFieldsLeft1zokd--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zokd = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zokd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zokd = 0
			}
			for nextMiss1zokd < maxFields1zokd && (found1zokd[nextMiss1zokd] || unmarshalMsgFieldSkip1zokd[nextMiss1zokd]) {
				nextMiss1zokd++
			}
			if nextMiss1zokd == maxFields1zokd {
				// filled all the empty fields!
				break doneWithStruct1zokd
			}
			missingFieldsLeft1zokd--
			curField1zokd = unmarshalMsgFieldOrder1zokd[nextMiss1zokd]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zokd)
		switch curField1zokd {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zokd[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zokd[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zokd[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zvnj uint32
				zvnj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zvnj > 0 {
					z.Snot = make(map[string]bool, zvnj)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zvnj > 0 {
					var zfjk string
					var zend bool
					zvnj--
					zfjk, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zend, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zfjk] = zend
				}
			}
		case "Notyet_zid03_map":
			found1zokd[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var ztqw uint32
				ztqw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && ztqw > 0 {
					z.Notyet = make(map[string]bool, ztqw)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for ztqw > 0 {
					var zcaa string
					var zrqo bool
					ztqw--
					zcaa, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zrqo, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zcaa] = zrqo
				}
			}
		case "Setm_zid04_slc":
			found1zokd[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zchy uint32
				zchy, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zchy) {
					z.Setm = (z.Setm)[:zchy]
				} else {
					z.Setm = make([]*inn, zchy)
				}
				for zssu := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zssu] != nil {
							z.Setm[zssu].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zssu] {
								z.Setm[zssu].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zssu] == nil {
								z.Setm[zssu] = new(inn)
							}
							bts, err = z.Setm[zssu].UnmarshalMsg(bts)
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
	if nextMiss1zokd != -1 {
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
var unmarshalMsgFieldOrder1zokd = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zokd = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zfjk, zend := range z.Snot {
			_ = zend
			_ = zfjk
			s += msgp.StringPrefixSize + len(zfjk) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zcaa, zrqo := range z.Notyet {
			_ = zrqo
			_ = zcaa
			s += msgp.StringPrefixSize + len(zcaa) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zssu := range z.Setm {
		if z.Setm[zssu] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zssu].Msgsize()
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
	const maxFields2ztwn = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2ztwn uint32
	totalEncodedFields2ztwn, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2ztwn := totalEncodedFields2ztwn
	missingFieldsLeft2ztwn := maxFields2ztwn - totalEncodedFields2ztwn

	var nextMiss2ztwn int32 = -1
	var found2ztwn [maxFields2ztwn]bool
	var curField2ztwn string

doneWithStruct2ztwn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2ztwn > 0 || missingFieldsLeft2ztwn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2ztwn, missingFieldsLeft2ztwn, msgp.ShowFound(found2ztwn[:]), decodeMsgFieldOrder2ztwn)
		if encodedFieldsLeft2ztwn > 0 {
			encodedFieldsLeft2ztwn--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2ztwn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2ztwn < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2ztwn = 0
			}
			for nextMiss2ztwn < maxFields2ztwn && (found2ztwn[nextMiss2ztwn] || decodeMsgFieldSkip2ztwn[nextMiss2ztwn]) {
				nextMiss2ztwn++
			}
			if nextMiss2ztwn == maxFields2ztwn {
				// filled all the empty fields!
				break doneWithStruct2ztwn
			}
			missingFieldsLeft2ztwn--
			curField2ztwn = decodeMsgFieldOrder2ztwn[nextMiss2ztwn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2ztwn)
		switch curField2ztwn {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2ztwn[0] = true
			var zkny uint32
			zkny, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zkny > 0 {
				z.j = make(map[string]bool, zkny)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zkny > 0 {
				zkny--
				var zjtf string
				var zmjy bool
				zjtf, err = dc.ReadString()
				if err != nil {
					return
				}
				zmjy, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zjtf] = zmjy
			}
		case "e_zid01_i64":
			found2ztwn[1] = true
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
	if nextMiss2ztwn != -1 {
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
var decodeMsgFieldOrder2ztwn = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2ztwn = []bool{false, false}

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
	for zjtf, zmjy := range z.j {
		err = en.WriteString(zjtf)
		if err != nil {
			return
		}
		err = en.WriteBool(zmjy)
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
	for zjtf, zmjy := range z.j {
		o = msgp.AppendString(o, zjtf)
		o = msgp.AppendBool(o, zmjy)
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
	const maxFields3zjgv = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zjgv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zjgv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zjgv := totalEncodedFields3zjgv
	missingFieldsLeft3zjgv := maxFields3zjgv - totalEncodedFields3zjgv

	var nextMiss3zjgv int32 = -1
	var found3zjgv [maxFields3zjgv]bool
	var curField3zjgv string

doneWithStruct3zjgv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zjgv > 0 || missingFieldsLeft3zjgv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zjgv, missingFieldsLeft3zjgv, msgp.ShowFound(found3zjgv[:]), unmarshalMsgFieldOrder3zjgv)
		if encodedFieldsLeft3zjgv > 0 {
			encodedFieldsLeft3zjgv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zjgv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zjgv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zjgv = 0
			}
			for nextMiss3zjgv < maxFields3zjgv && (found3zjgv[nextMiss3zjgv] || unmarshalMsgFieldSkip3zjgv[nextMiss3zjgv]) {
				nextMiss3zjgv++
			}
			if nextMiss3zjgv == maxFields3zjgv {
				// filled all the empty fields!
				break doneWithStruct3zjgv
			}
			missingFieldsLeft3zjgv--
			curField3zjgv = unmarshalMsgFieldOrder3zjgv[nextMiss3zjgv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zjgv)
		switch curField3zjgv {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zjgv[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var znye uint32
				znye, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && znye > 0 {
					z.j = make(map[string]bool, znye)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for znye > 0 {
					var zjtf string
					var zmjy bool
					znye--
					zjtf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zmjy, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zjtf] = zmjy
				}
			}
		case "e_zid01_i64":
			found3zjgv[1] = true
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
	if nextMiss3zjgv != -1 {
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
var unmarshalMsgFieldOrder3zjgv = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zjgv = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zjtf, zmjy := range z.j {
			_ = zmjy
			_ = zjtf
			s += msgp.StringPrefixSize + len(zjtf) + msgp.BoolSize
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
	const maxFields4zabv = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zabv uint32
	totalEncodedFields4zabv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zabv := totalEncodedFields4zabv
	missingFieldsLeft4zabv := maxFields4zabv - totalEncodedFields4zabv

	var nextMiss4zabv int32 = -1
	var found4zabv [maxFields4zabv]bool
	var curField4zabv string

doneWithStruct4zabv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zabv > 0 || missingFieldsLeft4zabv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zabv, missingFieldsLeft4zabv, msgp.ShowFound(found4zabv[:]), decodeMsgFieldOrder4zabv)
		if encodedFieldsLeft4zabv > 0 {
			encodedFieldsLeft4zabv--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zabv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zabv < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zabv = 0
			}
			for nextMiss4zabv < maxFields4zabv && (found4zabv[nextMiss4zabv] || decodeMsgFieldSkip4zabv[nextMiss4zabv]) {
				nextMiss4zabv++
			}
			if nextMiss4zabv == maxFields4zabv {
				// filled all the empty fields!
				break doneWithStruct4zabv
			}
			missingFieldsLeft4zabv--
			curField4zabv = decodeMsgFieldOrder4zabv[nextMiss4zabv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zabv)
		switch curField4zabv {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zabv[0] = true
			var zkey uint32
			zkey, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zkey > 0 {
				z.m = make(map[string]*Tr, zkey)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zkey > 0 {
				zkey--
				var zfpw string
				var zcqs *Tr
				zfpw, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zcqs != nil {
						dc.PushAlwaysNil()
						err = zcqs.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zcqs == nil {
						zcqs = new(Tr)
					}
					err = zcqs.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zfpw] = zcqs
			}
		case "s_zid01_str":
			found4zabv[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zabv[2] = true
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
	if nextMiss4zabv != -1 {
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
var decodeMsgFieldOrder4zabv = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zabv = []bool{false, false, false}

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
	for zfpw, zcqs := range z.m {
		err = en.WriteString(zfpw)
		if err != nil {
			return
		}
		if zcqs == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zcqs.EncodeMsg(en)
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
	for zfpw, zcqs := range z.m {
		o = msgp.AppendString(o, zfpw)
		if zcqs == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zcqs.MarshalMsg(o)
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
	const maxFields5zwei = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zwei uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zwei, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zwei := totalEncodedFields5zwei
	missingFieldsLeft5zwei := maxFields5zwei - totalEncodedFields5zwei

	var nextMiss5zwei int32 = -1
	var found5zwei [maxFields5zwei]bool
	var curField5zwei string

doneWithStruct5zwei:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zwei > 0 || missingFieldsLeft5zwei > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zwei, missingFieldsLeft5zwei, msgp.ShowFound(found5zwei[:]), unmarshalMsgFieldOrder5zwei)
		if encodedFieldsLeft5zwei > 0 {
			encodedFieldsLeft5zwei--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zwei = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zwei < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zwei = 0
			}
			for nextMiss5zwei < maxFields5zwei && (found5zwei[nextMiss5zwei] || unmarshalMsgFieldSkip5zwei[nextMiss5zwei]) {
				nextMiss5zwei++
			}
			if nextMiss5zwei == maxFields5zwei {
				// filled all the empty fields!
				break doneWithStruct5zwei
			}
			missingFieldsLeft5zwei--
			curField5zwei = unmarshalMsgFieldOrder5zwei[nextMiss5zwei]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zwei)
		switch curField5zwei {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zwei[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zagk uint32
				zagk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zagk > 0 {
					z.m = make(map[string]*Tr, zagk)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zagk > 0 {
					var zfpw string
					var zcqs *Tr
					zagk--
					zfpw, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zcqs != nil {
							zcqs.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zcqs {
								zcqs.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zcqs == nil {
								zcqs = new(Tr)
							}
							bts, err = zcqs.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zfpw] = zcqs
				}
			}
		case "s_zid01_str":
			found5zwei[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zwei[2] = true
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
	if nextMiss5zwei != -1 {
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
var unmarshalMsgFieldOrder5zwei = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zwei = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zfpw, zcqs := range z.m {
			_ = zcqs
			_ = zfpw
			s += msgp.StringPrefixSize + len(zfpw)
			if zcqs == nil {
				s += msgp.NilSize
			} else {
				s += zcqs.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}

// FileMy2 holds Greenpack schema from file 'my2.go'
type FileMy2 struct{}

// ZebraSchemaInMsgpack2Format provides the Greenpack Schema in msgpack2 format, length 2195 bytes
func (FileMy2) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaf, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x6d, 0x79,
		0x32, 0x2e, 0x67, 0x6f, 0xb2, 0x53, 0x6f, 0x75, 0x72, 0x63,
		0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa8, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
		0x74, 0x61, 0xb2, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0x00, 0xac, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73,
		0x5f, 0x5f, 0x6d, 0x61, 0x70, 0x83, 0xa2, 0x54, 0x72, 0x82,
		0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa2, 0x54, 0x72, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c,
		0x63, 0x96, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69,
		0x36, 0x34, 0x00, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa1, 0x55, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa1, 0x55, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f,
		0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x85, 0xa8, 0x5a, 0x69,
		0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa2, 0x65, 0x74, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x2d, 0xaf, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x5f, 0x5f, 0x1c, 0xa9, 0x53, 0x6b, 0x69, 0x70, 0x5f,
		0x5f, 0x62, 0x6f, 0x6f, 0xc3, 0x87, 0xa8, 0x5a, 0x69, 0x64,
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0x01, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa2, 0x4e, 0x74, 0xb1, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa2, 0x4e, 0x74, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xaf,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
		0x65, 0x5f, 0x5f, 0x0d, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f,
		0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f,
		0x5f, 0x0d, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa3, 0x69, 0x6e, 0x74, 0x86, 0xa8, 0x5a, 0x69, 0x64,
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0x02, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x6e, 0x6f,
		0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xaf,
		0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xaf, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
		0x5f, 0x18, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x84, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x18,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3,
		0x4d, 0x61, 0x70, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x12,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4,
		0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f,
		0x5f, 0x69, 0x36, 0x34, 0x03, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa6, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x4e, 0x6f,
		0x74, 0x79, 0x65, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x18, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x84, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x18, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xab, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x12, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x04, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x65, 0x74, 0x6d,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53,
		0x65, 0x74, 0x6d, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x5b, 0x5d, 0x2a, 0x69, 0x6e, 0x6e, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x1a, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x1a, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xab, 0x44,
		0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72,
		0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x1c, 0xa8,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7, 0x50,
		0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xab, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x16, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x6e,
		0xa3, 0x69, 0x6e, 0x6e, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa3, 0x69, 0x6e, 0x6e, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x92, 0x86, 0xa8,
		0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x00, 0xb0,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6a, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6a, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70,
		0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f,
		0x6f, 0x6c, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x18, 0xb2,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x84, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x18, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaa, 0x52, 0x61,
		0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x12, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f,
		0x6c, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0x01, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1,
		0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1,
		0x65, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xaf, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
		0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x11,
		0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x11, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0xa1, 0x75, 0x82, 0xaf, 0x53, 0x74, 0x72,
		0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa1, 0x75, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x93, 0x86, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x00, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6d, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6d, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xae, 0x6d, 0x61, 0x70, 0x5b,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a, 0x54, 0x72,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x18, 0xb2, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x84, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x18, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xab, 0x44,
		0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72,
		0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xaa, 0x52, 0x61, 0x6e, 0x67,
		0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x1c, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74,
		0x65, 0x72, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x16, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa2, 0x54, 0x72, 0x87, 0xa8, 0x5a, 0x69, 0x64,
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0x01, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa1, 0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa1, 0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x5f, 0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87,
		0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x02,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6e, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6e, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x11, 0xb2, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b,
		0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x11, 0xa8, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0xac, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f,
		0x5f, 0x73, 0x6c, 0x63, 0x91, 0xbd, 0x22, 0x67, 0x69, 0x74,
		0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6c,
		0x79, 0x63, 0x65, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x62,
		0x74, 0x72, 0x65, 0x65, 0x22,
	}
}

// ZebraSchemaInJsonCompact provides the Greenpack Schema in compact JSON format, length 2712 bytes
func (FileMy2) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath__str":"my2.go","SourcePackage__str":"testdata","ZebraSchemaId__i64":0,"Structs__map":{"Tr":{"StructName__str":"Tr","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"U","FieldTagName__str":"U","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"et","FieldTagName__str":"-","FieldCategory__":28,"Skip__boo":true},{"Zid__i64":1,"FieldGoName__str":"Nt","FieldTagName__str":"Nt","FieldTypeStr__str":"int","FieldCategory__":23,"FieldPrimitive__":13,"FieldFullType__ptr":{"Kind__":13,"Str__str":"int"}},{"Zid__i64":2,"FieldGoName__str":"Snot","FieldTagName__str":"Snot","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":3,"FieldGoName__str":"Notyet","FieldTagName__str":"Notyet","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":4,"FieldGoName__str":"Setm","FieldTagName__str":"Setm","FieldTypeStr__str":"[]*inn","FieldCategory__":26,"FieldFullType__ptr":{"Kind__":26,"Str__str":"Slice","Domain__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"inn"}}}}]},"inn":{"StructName__str":"inn","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"j","FieldTagName__str":"j","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":1,"FieldGoName__str":"e","FieldTagName__str":"e","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}}]},"u":{"StructName__str":"u","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"m","FieldTagName__str":"m","FieldTypeStr__str":"map[string]*Tr","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"Tr"}}}},{"Zid__i64":1,"FieldGoName__str":"s","FieldTagName__str":"s","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":2,"FieldGoName__str":"n","FieldTagName__str":"n","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}}]}},"Imports__slc":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the Greenpack Schema in pretty JSON format, length 6839 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath__str": "my2.go",
    "SourcePackage__str": "testdata",
    "ZebraSchemaId__i64": 0,
    "Structs__map": {
        "Tr": {
            "StructName__str": "Tr",
            "Fields__slc": [
                {
                    "Zid__i64": 0,
                    "FieldGoName__str": "U",
                    "FieldTagName__str": "U",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    }
                },
                {
                    "Zid__i64": -1,
                    "FieldGoName__str": "et",
                    "FieldTagName__str": "-",
                    "FieldCategory__": 28,
                    "Skip__boo": true
                },
                {
                    "Zid__i64": 1,
                    "FieldGoName__str": "Nt",
                    "FieldTagName__str": "Nt",
                    "FieldTypeStr__str": "int",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 13,
                    "FieldFullType__ptr": {
                        "Kind__": 13,
                        "Str__str": "int"
                    }
                },
                {
                    "Zid__i64": 2,
                    "FieldGoName__str": "Snot",
                    "FieldTagName__str": "Snot",
                    "FieldTypeStr__str": "map[string]bool",
                    "FieldCategory__": 24,
                    "FieldFullType__ptr": {
                        "Kind__": 24,
                        "Str__str": "Map",
                        "Domain__ptr": {
                            "Kind__": 2,
                            "Str__str": "string"
                        },
                        "Range__ptr": {
                            "Kind__": 18,
                            "Str__str": "bool"
                        }
                    }
                },
                {
                    "Zid__i64": 3,
                    "FieldGoName__str": "Notyet",
                    "FieldTagName__str": "Notyet",
                    "FieldTypeStr__str": "map[string]bool",
                    "FieldCategory__": 24,
                    "FieldFullType__ptr": {
                        "Kind__": 24,
                        "Str__str": "Map",
                        "Domain__ptr": {
                            "Kind__": 2,
                            "Str__str": "string"
                        },
                        "Range__ptr": {
                            "Kind__": 18,
                            "Str__str": "bool"
                        }
                    }
                },
                {
                    "Zid__i64": 4,
                    "FieldGoName__str": "Setm",
                    "FieldTagName__str": "Setm",
                    "FieldTypeStr__str": "[]*inn",
                    "FieldCategory__": 26,
                    "FieldFullType__ptr": {
                        "Kind__": 26,
                        "Str__str": "Slice",
                        "Domain__ptr": {
                            "Kind__": 28,
                            "Str__str": "Pointer",
                            "Domain__ptr": {
                                "Kind__": 22,
                                "Str__str": "inn"
                            }
                        }
                    }
                }
            ]
        },
        "inn": {
            "StructName__str": "inn",
            "Fields__slc": [
                {
                    "Zid__i64": 0,
                    "FieldGoName__str": "j",
                    "FieldTagName__str": "j",
                    "FieldTypeStr__str": "map[string]bool",
                    "FieldCategory__": 24,
                    "FieldFullType__ptr": {
                        "Kind__": 24,
                        "Str__str": "Map",
                        "Domain__ptr": {
                            "Kind__": 2,
                            "Str__str": "string"
                        },
                        "Range__ptr": {
                            "Kind__": 18,
                            "Str__str": "bool"
                        }
                    }
                },
                {
                    "Zid__i64": 1,
                    "FieldGoName__str": "e",
                    "FieldTagName__str": "e",
                    "FieldTypeStr__str": "int64",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 17,
                    "FieldFullType__ptr": {
                        "Kind__": 17,
                        "Str__str": "int64"
                    }
                }
            ]
        },
        "u": {
            "StructName__str": "u",
            "Fields__slc": [
                {
                    "Zid__i64": 0,
                    "FieldGoName__str": "m",
                    "FieldTagName__str": "m",
                    "FieldTypeStr__str": "map[string]*Tr",
                    "FieldCategory__": 24,
                    "FieldFullType__ptr": {
                        "Kind__": 24,
                        "Str__str": "Map",
                        "Domain__ptr": {
                            "Kind__": 2,
                            "Str__str": "string"
                        },
                        "Range__ptr": {
                            "Kind__": 28,
                            "Str__str": "Pointer",
                            "Domain__ptr": {
                                "Kind__": 22,
                                "Str__str": "Tr"
                            }
                        }
                    }
                },
                {
                    "Zid__i64": 1,
                    "FieldGoName__str": "s",
                    "FieldTagName__str": "s",
                    "FieldTypeStr__str": "string",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 2,
                    "FieldFullType__ptr": {
                        "Kind__": 2,
                        "Str__str": "string"
                    }
                },
                {
                    "Zid__i64": 2,
                    "FieldGoName__str": "n",
                    "FieldTagName__str": "n",
                    "FieldTypeStr__str": "int64",
                    "FieldCategory__": 23,
                    "FieldPrimitive__": 17,
                    "FieldFullType__ptr": {
                        "Kind__": 17,
                        "Str__str": "int64"
                    }
                }
            ]
        }
    },
    "Imports__slc": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
