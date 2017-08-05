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
	const maxFields0zton = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zton uint32
	totalEncodedFields0zton, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zton := totalEncodedFields0zton
	missingFieldsLeft0zton := maxFields0zton - totalEncodedFields0zton

	var nextMiss0zton int32 = -1
	var found0zton [maxFields0zton]bool
	var curField0zton string

doneWithStruct0zton:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zton > 0 || missingFieldsLeft0zton > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zton, missingFieldsLeft0zton, msgp.ShowFound(found0zton[:]), decodeMsgFieldOrder0zton)
		if encodedFieldsLeft0zton > 0 {
			encodedFieldsLeft0zton--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zton = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zton < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zton = 0
			}
			for nextMiss0zton < maxFields0zton && (found0zton[nextMiss0zton] || decodeMsgFieldSkip0zton[nextMiss0zton]) {
				nextMiss0zton++
			}
			if nextMiss0zton == maxFields0zton {
				// filled all the empty fields!
				break doneWithStruct0zton
			}
			missingFieldsLeft0zton--
			curField0zton = decodeMsgFieldOrder0zton[nextMiss0zton]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zton)
		switch curField0zton {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zton[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zton[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zton[3] = true
			var zasp uint32
			zasp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zasp > 0 {
				z.Snot = make(map[string]bool, zasp)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zasp > 0 {
				zasp--
				var zrwc string
				var zwba bool
				zrwc, err = dc.ReadString()
				if err != nil {
					return
				}
				zwba, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zrwc] = zwba
			}
		case "Notyet_zid03_map":
			found0zton[4] = true
			var zlys uint32
			zlys, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zlys > 0 {
				z.Notyet = make(map[string]bool, zlys)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zlys > 0 {
				zlys--
				var zeyu string
				var zmlj bool
				zeyu, err = dc.ReadString()
				if err != nil {
					return
				}
				zmlj, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zeyu] = zmlj
			}
		case "Setm_zid04_slc":
			found0zton[5] = true
			var zkjm uint32
			zkjm, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zkjm) {
				z.Setm = (z.Setm)[:zkjm]
			} else {
				z.Setm = make([]*inn, zkjm)
			}
			for zuqc := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zuqc] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zuqc].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zuqc] == nil {
						z.Setm[zuqc] = new(inn)
					}
					err = z.Setm[zuqc].DecodeMsg(dc)
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
	if nextMiss0zton != -1 {
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
var decodeMsgFieldOrder0zton = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zton = []bool{false, true, false, false, false, false}

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
	for zrwc, zwba := range z.Snot {
		err = en.WriteString(zrwc)
		if err != nil {
			return
		}
		err = en.WriteBool(zwba)
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
	for zeyu, zmlj := range z.Notyet {
		err = en.WriteString(zeyu)
		if err != nil {
			return
		}
		err = en.WriteBool(zmlj)
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
	for zuqc := range z.Setm {
		if z.Setm[zuqc] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Setm[zuqc].EncodeMsg(en)
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
	for zrwc, zwba := range z.Snot {
		o = msgp.AppendString(o, zrwc)
		o = msgp.AppendBool(o, zwba)
	}
	// string "Notyet_zid03_map"
	o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
	for zeyu, zmlj := range z.Notyet {
		o = msgp.AppendString(o, zeyu)
		o = msgp.AppendBool(o, zmlj)
	}
	// string "Setm_zid04_slc"
	o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
	for zuqc := range z.Setm {
		if z.Setm[zuqc] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Setm[zuqc].MarshalMsg(o)
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
	const maxFields1zdpv = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zdpv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zdpv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zdpv := totalEncodedFields1zdpv
	missingFieldsLeft1zdpv := maxFields1zdpv - totalEncodedFields1zdpv

	var nextMiss1zdpv int32 = -1
	var found1zdpv [maxFields1zdpv]bool
	var curField1zdpv string

doneWithStruct1zdpv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zdpv > 0 || missingFieldsLeft1zdpv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zdpv, missingFieldsLeft1zdpv, msgp.ShowFound(found1zdpv[:]), unmarshalMsgFieldOrder1zdpv)
		if encodedFieldsLeft1zdpv > 0 {
			encodedFieldsLeft1zdpv--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zdpv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zdpv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zdpv = 0
			}
			for nextMiss1zdpv < maxFields1zdpv && (found1zdpv[nextMiss1zdpv] || unmarshalMsgFieldSkip1zdpv[nextMiss1zdpv]) {
				nextMiss1zdpv++
			}
			if nextMiss1zdpv == maxFields1zdpv {
				// filled all the empty fields!
				break doneWithStruct1zdpv
			}
			missingFieldsLeft1zdpv--
			curField1zdpv = unmarshalMsgFieldOrder1zdpv[nextMiss1zdpv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zdpv)
		switch curField1zdpv {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zdpv[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zdpv[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zdpv[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zbrf uint32
				zbrf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zbrf > 0 {
					z.Snot = make(map[string]bool, zbrf)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zbrf > 0 {
					var zrwc string
					var zwba bool
					zbrf--
					zrwc, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zwba, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zrwc] = zwba
				}
			}
		case "Notyet_zid03_map":
			found1zdpv[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var ztnq uint32
				ztnq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && ztnq > 0 {
					z.Notyet = make(map[string]bool, ztnq)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for ztnq > 0 {
					var zeyu string
					var zmlj bool
					ztnq--
					zeyu, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zmlj, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zeyu] = zmlj
				}
			}
		case "Setm_zid04_slc":
			found1zdpv[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var znol uint32
				znol, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(znol) {
					z.Setm = (z.Setm)[:znol]
				} else {
					z.Setm = make([]*inn, znol)
				}
				for zuqc := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zuqc] != nil {
							z.Setm[zuqc].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zuqc] {
								z.Setm[zuqc].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zuqc] == nil {
								z.Setm[zuqc] = new(inn)
							}
							bts, err = z.Setm[zuqc].UnmarshalMsg(bts)
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
	if nextMiss1zdpv != -1 {
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
var unmarshalMsgFieldOrder1zdpv = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zdpv = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zrwc, zwba := range z.Snot {
			_ = zwba
			_ = zrwc
			s += msgp.StringPrefixSize + len(zrwc) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zeyu, zmlj := range z.Notyet {
			_ = zmlj
			_ = zeyu
			s += msgp.StringPrefixSize + len(zeyu) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zuqc := range z.Setm {
		if z.Setm[zuqc] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zuqc].Msgsize()
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
	const maxFields2zqoe = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zqoe uint32
	totalEncodedFields2zqoe, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zqoe := totalEncodedFields2zqoe
	missingFieldsLeft2zqoe := maxFields2zqoe - totalEncodedFields2zqoe

	var nextMiss2zqoe int32 = -1
	var found2zqoe [maxFields2zqoe]bool
	var curField2zqoe string

doneWithStruct2zqoe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zqoe > 0 || missingFieldsLeft2zqoe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zqoe, missingFieldsLeft2zqoe, msgp.ShowFound(found2zqoe[:]), decodeMsgFieldOrder2zqoe)
		if encodedFieldsLeft2zqoe > 0 {
			encodedFieldsLeft2zqoe--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zqoe = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zqoe < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zqoe = 0
			}
			for nextMiss2zqoe < maxFields2zqoe && (found2zqoe[nextMiss2zqoe] || decodeMsgFieldSkip2zqoe[nextMiss2zqoe]) {
				nextMiss2zqoe++
			}
			if nextMiss2zqoe == maxFields2zqoe {
				// filled all the empty fields!
				break doneWithStruct2zqoe
			}
			missingFieldsLeft2zqoe--
			curField2zqoe = decodeMsgFieldOrder2zqoe[nextMiss2zqoe]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zqoe)
		switch curField2zqoe {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zqoe[0] = true
			var zcuo uint32
			zcuo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zcuo > 0 {
				z.j = make(map[string]bool, zcuo)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zcuo > 0 {
				zcuo--
				var zpoj string
				var zzzl bool
				zpoj, err = dc.ReadString()
				if err != nil {
					return
				}
				zzzl, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zpoj] = zzzl
			}
		case "e_zid01_i64":
			found2zqoe[1] = true
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
	if nextMiss2zqoe != -1 {
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
var decodeMsgFieldOrder2zqoe = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zqoe = []bool{false, false}

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
	for zpoj, zzzl := range z.j {
		err = en.WriteString(zpoj)
		if err != nil {
			return
		}
		err = en.WriteBool(zzzl)
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
	for zpoj, zzzl := range z.j {
		o = msgp.AppendString(o, zpoj)
		o = msgp.AppendBool(o, zzzl)
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
	const maxFields3zsgu = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zsgu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zsgu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zsgu := totalEncodedFields3zsgu
	missingFieldsLeft3zsgu := maxFields3zsgu - totalEncodedFields3zsgu

	var nextMiss3zsgu int32 = -1
	var found3zsgu [maxFields3zsgu]bool
	var curField3zsgu string

doneWithStruct3zsgu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zsgu > 0 || missingFieldsLeft3zsgu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zsgu, missingFieldsLeft3zsgu, msgp.ShowFound(found3zsgu[:]), unmarshalMsgFieldOrder3zsgu)
		if encodedFieldsLeft3zsgu > 0 {
			encodedFieldsLeft3zsgu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zsgu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zsgu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zsgu = 0
			}
			for nextMiss3zsgu < maxFields3zsgu && (found3zsgu[nextMiss3zsgu] || unmarshalMsgFieldSkip3zsgu[nextMiss3zsgu]) {
				nextMiss3zsgu++
			}
			if nextMiss3zsgu == maxFields3zsgu {
				// filled all the empty fields!
				break doneWithStruct3zsgu
			}
			missingFieldsLeft3zsgu--
			curField3zsgu = unmarshalMsgFieldOrder3zsgu[nextMiss3zsgu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zsgu)
		switch curField3zsgu {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zsgu[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zxhj uint32
				zxhj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zxhj > 0 {
					z.j = make(map[string]bool, zxhj)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zxhj > 0 {
					var zpoj string
					var zzzl bool
					zxhj--
					zpoj, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zzzl, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zpoj] = zzzl
				}
			}
		case "e_zid01_i64":
			found3zsgu[1] = true
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
	if nextMiss3zsgu != -1 {
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
var unmarshalMsgFieldOrder3zsgu = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zsgu = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zpoj, zzzl := range z.j {
			_ = zzzl
			_ = zpoj
			s += msgp.StringPrefixSize + len(zpoj) + msgp.BoolSize
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
	const maxFields4zwqz = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zwqz uint32
	totalEncodedFields4zwqz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zwqz := totalEncodedFields4zwqz
	missingFieldsLeft4zwqz := maxFields4zwqz - totalEncodedFields4zwqz

	var nextMiss4zwqz int32 = -1
	var found4zwqz [maxFields4zwqz]bool
	var curField4zwqz string

doneWithStruct4zwqz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zwqz > 0 || missingFieldsLeft4zwqz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zwqz, missingFieldsLeft4zwqz, msgp.ShowFound(found4zwqz[:]), decodeMsgFieldOrder4zwqz)
		if encodedFieldsLeft4zwqz > 0 {
			encodedFieldsLeft4zwqz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zwqz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zwqz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zwqz = 0
			}
			for nextMiss4zwqz < maxFields4zwqz && (found4zwqz[nextMiss4zwqz] || decodeMsgFieldSkip4zwqz[nextMiss4zwqz]) {
				nextMiss4zwqz++
			}
			if nextMiss4zwqz == maxFields4zwqz {
				// filled all the empty fields!
				break doneWithStruct4zwqz
			}
			missingFieldsLeft4zwqz--
			curField4zwqz = decodeMsgFieldOrder4zwqz[nextMiss4zwqz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zwqz)
		switch curField4zwqz {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zwqz[0] = true
			var zqgn uint32
			zqgn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zqgn > 0 {
				z.m = make(map[string]*Tr, zqgn)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zqgn > 0 {
				zqgn--
				var zzck string
				var zqnf *Tr
				zzck, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zqnf != nil {
						dc.PushAlwaysNil()
						err = zqnf.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zqnf == nil {
						zqnf = new(Tr)
					}
					err = zqnf.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zzck] = zqnf
			}
		case "s_zid01_str":
			found4zwqz[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zwqz[2] = true
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
	if nextMiss4zwqz != -1 {
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
var decodeMsgFieldOrder4zwqz = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zwqz = []bool{false, false, false}

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
	for zzck, zqnf := range z.m {
		err = en.WriteString(zzck)
		if err != nil {
			return
		}
		if zqnf == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zqnf.EncodeMsg(en)
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
	for zzck, zqnf := range z.m {
		o = msgp.AppendString(o, zzck)
		if zqnf == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zqnf.MarshalMsg(o)
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
	const maxFields5zdeo = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zdeo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zdeo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zdeo := totalEncodedFields5zdeo
	missingFieldsLeft5zdeo := maxFields5zdeo - totalEncodedFields5zdeo

	var nextMiss5zdeo int32 = -1
	var found5zdeo [maxFields5zdeo]bool
	var curField5zdeo string

doneWithStruct5zdeo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zdeo > 0 || missingFieldsLeft5zdeo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zdeo, missingFieldsLeft5zdeo, msgp.ShowFound(found5zdeo[:]), unmarshalMsgFieldOrder5zdeo)
		if encodedFieldsLeft5zdeo > 0 {
			encodedFieldsLeft5zdeo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zdeo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zdeo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zdeo = 0
			}
			for nextMiss5zdeo < maxFields5zdeo && (found5zdeo[nextMiss5zdeo] || unmarshalMsgFieldSkip5zdeo[nextMiss5zdeo]) {
				nextMiss5zdeo++
			}
			if nextMiss5zdeo == maxFields5zdeo {
				// filled all the empty fields!
				break doneWithStruct5zdeo
			}
			missingFieldsLeft5zdeo--
			curField5zdeo = unmarshalMsgFieldOrder5zdeo[nextMiss5zdeo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zdeo)
		switch curField5zdeo {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zdeo[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zuzt uint32
				zuzt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zuzt > 0 {
					z.m = make(map[string]*Tr, zuzt)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zuzt > 0 {
					var zzck string
					var zqnf *Tr
					zuzt--
					zzck, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zqnf != nil {
							zqnf.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zqnf {
								zqnf.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zqnf == nil {
								zqnf = new(Tr)
							}
							bts, err = zqnf.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zzck] = zqnf
				}
			}
		case "s_zid01_str":
			found5zdeo[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zdeo[2] = true
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
	if nextMiss5zdeo != -1 {
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
var unmarshalMsgFieldOrder5zdeo = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zdeo = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zzck, zqnf := range z.m {
			_ = zqnf
			_ = zzck
			s += msgp.StringPrefixSize + len(zzck)
			if zqnf == nil {
				s += msgp.NilSize
			} else {
				s += zqnf.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}

// FileMy2 holds Greenpack schema from file 'my2.go'
type FileMy2 struct{}

// GreenSchemaInMsgpack2Format provides the Greenpack Schema in msgpack2 format, length 2195 bytes
func (FileMy2) GreenSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaf, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x6d, 0x79,
		0x32, 0x2e, 0x67, 0x6f, 0xb2, 0x53, 0x6f, 0x75, 0x72, 0x63,
		0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa8, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
		0x74, 0x61, 0xb2, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0x00, 0xac, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73,
		0x5f, 0x5f, 0x6d, 0x61, 0x70, 0x83, 0xa3, 0x69, 0x6e, 0x6e,
		0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e,
		0x6e, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f,
		0x73, 0x6c, 0x63, 0x92, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f,
		0x5f, 0x69, 0x36, 0x34, 0x00, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa1, 0x6a, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa1, 0x6a, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
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
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x87, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x01, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x65, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x65, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x5f, 0x5f, 0x11, 0xb2, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x11, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa1,
		0x75, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x75,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73,
		0x6c, 0x63, 0x93, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0x00, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa1, 0x6d, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa1, 0x6d, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xae, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0x5d, 0x2a, 0x54, 0x72, 0xaf, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x5f, 0x5f, 0x18, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x84, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x18, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa3, 0x4d, 0x61, 0x70, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x1c, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xab, 0x44,
		0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72,
		0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x16, 0xa8,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa2, 0x54,
		0x72, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36,
		0x34, 0x01, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1,
		0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1,
		0x73, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
		0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f,
		0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72,
		0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa8, 0x5a, 0x69, 0x64,
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0x02, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa1, 0x6e, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa1, 0x6e, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xaf,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
		0x65, 0x5f, 0x5f, 0x11, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f,
		0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f,
		0x5f, 0x11, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa2, 0x54, 0x72,
		0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa2, 0x54, 0x72,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73,
		0x6c, 0x63, 0x96, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0x00, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa1, 0x55, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa1, 0x55, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x5f, 0x5f, 0x02, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x85, 0xa8, 0x5a,
		0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa2, 0x65, 0x74, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x2d, 0xaf, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x5f, 0x5f, 0x1c, 0xa9, 0x53, 0x6b, 0x69, 0x70,
		0x5f, 0x5f, 0x62, 0x6f, 0x6f, 0xc3, 0x87, 0xa8, 0x5a, 0x69,
		0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x01, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa2, 0x4e, 0x74, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa2, 0x4e, 0x74, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74,
		0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x5f, 0x5f, 0x0d, 0xb2, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f,
		0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64,
		0x5f, 0x5f, 0x0d, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0x86, 0xa8, 0x5a, 0x69,
		0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x02, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x6e,
		0x6f, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xaf, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x5f, 0x5f, 0x18, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x84, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x18, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa3, 0x4d, 0x61, 0x70, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x12, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa8, 0x5a, 0x69, 0x64,
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0x03, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa6, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x4e,
		0x6f, 0x74, 0x79, 0x65, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xaf,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x18, 0xb2, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x84, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x18, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xab, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0xaa, 0x52, 0x61, 0x6e, 0x67, 0x65,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x12, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa8,
		0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x04, 0xb0,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x65, 0x74,
		0x6d, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4,
		0x53, 0x65, 0x74, 0x6d, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa6, 0x5b, 0x5d, 0x2a, 0x69, 0x6e, 0x6e, 0xaf,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x1a, 0xb2, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e,
		0x64, 0x5f, 0x5f, 0x1a, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xab,
		0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x1c,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7,
		0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xab, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x16, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e,
		0x6e, 0xac, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f,
		0x5f, 0x73, 0x6c, 0x63, 0x91, 0xbd, 0x22, 0x67, 0x69, 0x74,
		0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6c,
		0x79, 0x63, 0x65, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x72, 0x62,
		0x74, 0x72, 0x65, 0x65, 0x22,
	}
}

// GreenSchemaInJsonCompact provides the Greenpack Schema in compact JSON format, length 2712 bytes
func (FileMy2) GreenSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath__str":"my2.go","SourcePackage__str":"testdata","GreenSchemaId__i64":0,"Structs__map":{"inn":{"StructName__str":"inn","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"j","FieldTagName__str":"j","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":1,"FieldGoName__str":"e","FieldTagName__str":"e","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}}]},"u":{"StructName__str":"u","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"m","FieldTagName__str":"m","FieldTypeStr__str":"map[string]*Tr","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"Tr"}}}},{"Zid__i64":1,"FieldGoName__str":"s","FieldTagName__str":"s","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":2,"FieldGoName__str":"n","FieldTagName__str":"n","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}}]},"Tr":{"StructName__str":"Tr","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"U","FieldTagName__str":"U","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"et","FieldTagName__str":"-","FieldCategory__":28,"Skip__boo":true},{"Zid__i64":1,"FieldGoName__str":"Nt","FieldTagName__str":"Nt","FieldTypeStr__str":"int","FieldCategory__":23,"FieldPrimitive__":13,"FieldFullType__ptr":{"Kind__":13,"Str__str":"int"}},{"Zid__i64":2,"FieldGoName__str":"Snot","FieldTagName__str":"Snot","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":3,"FieldGoName__str":"Notyet","FieldTagName__str":"Notyet","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":4,"FieldGoName__str":"Setm","FieldTagName__str":"Setm","FieldTypeStr__str":"[]*inn","FieldCategory__":26,"FieldFullType__ptr":{"Kind__":26,"Str__str":"Slice","Domain__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"inn"}}}}]}},"Imports__slc":["\"github.com/glycerine/rbtree\""]}`)
}

// GreenSchemaInJsonPretty provides the Greenpack Schema in pretty JSON format, length 6839 bytes
func (FileMy2) GreenSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath__str": "my2.go",
    "SourcePackage__str": "testdata",
    "GreenSchemaId__i64": 0,
    "Structs__map": {
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
        },
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
        }
    },
    "Imports__slc": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
