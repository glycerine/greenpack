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
	const maxFields0ztkk = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0ztkk uint32
	totalEncodedFields0ztkk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0ztkk := totalEncodedFields0ztkk
	missingFieldsLeft0ztkk := maxFields0ztkk - totalEncodedFields0ztkk

	var nextMiss0ztkk int32 = -1
	var found0ztkk [maxFields0ztkk]bool
	var curField0ztkk string

doneWithStruct0ztkk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ztkk > 0 || missingFieldsLeft0ztkk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ztkk, missingFieldsLeft0ztkk, msgp.ShowFound(found0ztkk[:]), decodeMsgFieldOrder0ztkk)
		if encodedFieldsLeft0ztkk > 0 {
			encodedFieldsLeft0ztkk--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0ztkk = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ztkk < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0ztkk = 0
			}
			for nextMiss0ztkk < maxFields0ztkk && (found0ztkk[nextMiss0ztkk] || decodeMsgFieldSkip0ztkk[nextMiss0ztkk]) {
				nextMiss0ztkk++
			}
			if nextMiss0ztkk == maxFields0ztkk {
				// filled all the empty fields!
				break doneWithStruct0ztkk
			}
			missingFieldsLeft0ztkk--
			curField0ztkk = decodeMsgFieldOrder0ztkk[nextMiss0ztkk]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ztkk)
		switch curField0ztkk {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0ztkk[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0ztkk[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0ztkk[3] = true
			var zsub uint32
			zsub, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zsub > 0 {
				z.Snot = make(map[string]bool, zsub)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zsub > 0 {
				zsub--
				var ztxx string
				var zomj bool
				ztxx, err = dc.ReadString()
				if err != nil {
					return
				}
				zomj, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[ztxx] = zomj
			}
		case "Notyet_zid03_map":
			found0ztkk[4] = true
			var zpuy uint32
			zpuy, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zpuy > 0 {
				z.Notyet = make(map[string]bool, zpuy)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zpuy > 0 {
				zpuy--
				var zetm string
				var zysc bool
				zetm, err = dc.ReadString()
				if err != nil {
					return
				}
				zysc, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zetm] = zysc
			}
		case "Setm_zid04_slc":
			found0ztkk[5] = true
			var znbr uint32
			znbr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(znbr) {
				z.Setm = (z.Setm)[:znbr]
			} else {
				z.Setm = make([]*inn, znbr)
			}
			for zooh := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zooh] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zooh].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zooh] == nil {
						z.Setm[zooh] = new(inn)
					}
					err = z.Setm[zooh].DecodeMsg(dc)
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
	if nextMiss0ztkk != -1 {
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
var decodeMsgFieldOrder0ztkk = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0ztkk = []bool{false, true, false, false, false, false}

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
	for ztxx, zomj := range z.Snot {
		err = en.WriteString(ztxx)
		if err != nil {
			return
		}
		err = en.WriteBool(zomj)
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
	for zetm, zysc := range z.Notyet {
		err = en.WriteString(zetm)
		if err != nil {
			return
		}
		err = en.WriteBool(zysc)
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
	for zooh := range z.Setm {
		if z.Setm[zooh] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Setm[zooh].EncodeMsg(en)
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
	for ztxx, zomj := range z.Snot {
		o = msgp.AppendString(o, ztxx)
		o = msgp.AppendBool(o, zomj)
	}
	// string "Notyet_zid03_map"
	o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
	for zetm, zysc := range z.Notyet {
		o = msgp.AppendString(o, zetm)
		o = msgp.AppendBool(o, zysc)
	}
	// string "Setm_zid04_slc"
	o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
	for zooh := range z.Setm {
		if z.Setm[zooh] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Setm[zooh].MarshalMsg(o)
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
	const maxFields1zbvr = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zbvr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zbvr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zbvr := totalEncodedFields1zbvr
	missingFieldsLeft1zbvr := maxFields1zbvr - totalEncodedFields1zbvr

	var nextMiss1zbvr int32 = -1
	var found1zbvr [maxFields1zbvr]bool
	var curField1zbvr string

doneWithStruct1zbvr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zbvr > 0 || missingFieldsLeft1zbvr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zbvr, missingFieldsLeft1zbvr, msgp.ShowFound(found1zbvr[:]), unmarshalMsgFieldOrder1zbvr)
		if encodedFieldsLeft1zbvr > 0 {
			encodedFieldsLeft1zbvr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zbvr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zbvr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zbvr = 0
			}
			for nextMiss1zbvr < maxFields1zbvr && (found1zbvr[nextMiss1zbvr] || unmarshalMsgFieldSkip1zbvr[nextMiss1zbvr]) {
				nextMiss1zbvr++
			}
			if nextMiss1zbvr == maxFields1zbvr {
				// filled all the empty fields!
				break doneWithStruct1zbvr
			}
			missingFieldsLeft1zbvr--
			curField1zbvr = unmarshalMsgFieldOrder1zbvr[nextMiss1zbvr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zbvr)
		switch curField1zbvr {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zbvr[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zbvr[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zbvr[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zubk uint32
				zubk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zubk > 0 {
					z.Snot = make(map[string]bool, zubk)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zubk > 0 {
					var ztxx string
					var zomj bool
					zubk--
					ztxx, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zomj, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[ztxx] = zomj
				}
			}
		case "Notyet_zid03_map":
			found1zbvr[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zcmf uint32
				zcmf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zcmf > 0 {
					z.Notyet = make(map[string]bool, zcmf)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zcmf > 0 {
					var zetm string
					var zysc bool
					zcmf--
					zetm, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zysc, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zetm] = zysc
				}
			}
		case "Setm_zid04_slc":
			found1zbvr[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zhvy uint32
				zhvy, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zhvy) {
					z.Setm = (z.Setm)[:zhvy]
				} else {
					z.Setm = make([]*inn, zhvy)
				}
				for zooh := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zooh] != nil {
							z.Setm[zooh].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zooh] {
								z.Setm[zooh].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zooh] == nil {
								z.Setm[zooh] = new(inn)
							}
							bts, err = z.Setm[zooh].UnmarshalMsg(bts)
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
	if nextMiss1zbvr != -1 {
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
var unmarshalMsgFieldOrder1zbvr = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zbvr = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for ztxx, zomj := range z.Snot {
			_ = zomj
			_ = ztxx
			s += msgp.StringPrefixSize + len(ztxx) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zetm, zysc := range z.Notyet {
			_ = zysc
			_ = zetm
			s += msgp.StringPrefixSize + len(zetm) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zooh := range z.Setm {
		if z.Setm[zooh] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zooh].Msgsize()
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
	const maxFields2zxir = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zxir uint32
	totalEncodedFields2zxir, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zxir := totalEncodedFields2zxir
	missingFieldsLeft2zxir := maxFields2zxir - totalEncodedFields2zxir

	var nextMiss2zxir int32 = -1
	var found2zxir [maxFields2zxir]bool
	var curField2zxir string

doneWithStruct2zxir:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zxir > 0 || missingFieldsLeft2zxir > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zxir, missingFieldsLeft2zxir, msgp.ShowFound(found2zxir[:]), decodeMsgFieldOrder2zxir)
		if encodedFieldsLeft2zxir > 0 {
			encodedFieldsLeft2zxir--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zxir = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zxir < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zxir = 0
			}
			for nextMiss2zxir < maxFields2zxir && (found2zxir[nextMiss2zxir] || decodeMsgFieldSkip2zxir[nextMiss2zxir]) {
				nextMiss2zxir++
			}
			if nextMiss2zxir == maxFields2zxir {
				// filled all the empty fields!
				break doneWithStruct2zxir
			}
			missingFieldsLeft2zxir--
			curField2zxir = decodeMsgFieldOrder2zxir[nextMiss2zxir]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zxir)
		switch curField2zxir {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zxir[0] = true
			var zivu uint32
			zivu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zivu > 0 {
				z.j = make(map[string]bool, zivu)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zivu > 0 {
				zivu--
				var znie string
				var zvxn bool
				znie, err = dc.ReadString()
				if err != nil {
					return
				}
				zvxn, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[znie] = zvxn
			}
		case "e_zid01_i64":
			found2zxir[1] = true
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
	if nextMiss2zxir != -1 {
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
var decodeMsgFieldOrder2zxir = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zxir = []bool{false, false}

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
	for znie, zvxn := range z.j {
		err = en.WriteString(znie)
		if err != nil {
			return
		}
		err = en.WriteBool(zvxn)
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
	for znie, zvxn := range z.j {
		o = msgp.AppendString(o, znie)
		o = msgp.AppendBool(o, zvxn)
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
	const maxFields3ztic = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3ztic uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3ztic, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3ztic := totalEncodedFields3ztic
	missingFieldsLeft3ztic := maxFields3ztic - totalEncodedFields3ztic

	var nextMiss3ztic int32 = -1
	var found3ztic [maxFields3ztic]bool
	var curField3ztic string

doneWithStruct3ztic:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3ztic > 0 || missingFieldsLeft3ztic > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3ztic, missingFieldsLeft3ztic, msgp.ShowFound(found3ztic[:]), unmarshalMsgFieldOrder3ztic)
		if encodedFieldsLeft3ztic > 0 {
			encodedFieldsLeft3ztic--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3ztic = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3ztic < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3ztic = 0
			}
			for nextMiss3ztic < maxFields3ztic && (found3ztic[nextMiss3ztic] || unmarshalMsgFieldSkip3ztic[nextMiss3ztic]) {
				nextMiss3ztic++
			}
			if nextMiss3ztic == maxFields3ztic {
				// filled all the empty fields!
				break doneWithStruct3ztic
			}
			missingFieldsLeft3ztic--
			curField3ztic = unmarshalMsgFieldOrder3ztic[nextMiss3ztic]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3ztic)
		switch curField3ztic {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3ztic[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zlwd uint32
				zlwd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zlwd > 0 {
					z.j = make(map[string]bool, zlwd)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zlwd > 0 {
					var znie string
					var zvxn bool
					zlwd--
					znie, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zvxn, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[znie] = zvxn
				}
			}
		case "e_zid01_i64":
			found3ztic[1] = true
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
	if nextMiss3ztic != -1 {
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
var unmarshalMsgFieldOrder3ztic = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3ztic = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for znie, zvxn := range z.j {
			_ = zvxn
			_ = znie
			s += msgp.StringPrefixSize + len(znie) + msgp.BoolSize
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
	const maxFields4zkes = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zkes uint32
	totalEncodedFields4zkes, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zkes := totalEncodedFields4zkes
	missingFieldsLeft4zkes := maxFields4zkes - totalEncodedFields4zkes

	var nextMiss4zkes int32 = -1
	var found4zkes [maxFields4zkes]bool
	var curField4zkes string

doneWithStruct4zkes:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zkes > 0 || missingFieldsLeft4zkes > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zkes, missingFieldsLeft4zkes, msgp.ShowFound(found4zkes[:]), decodeMsgFieldOrder4zkes)
		if encodedFieldsLeft4zkes > 0 {
			encodedFieldsLeft4zkes--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zkes = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zkes < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zkes = 0
			}
			for nextMiss4zkes < maxFields4zkes && (found4zkes[nextMiss4zkes] || decodeMsgFieldSkip4zkes[nextMiss4zkes]) {
				nextMiss4zkes++
			}
			if nextMiss4zkes == maxFields4zkes {
				// filled all the empty fields!
				break doneWithStruct4zkes
			}
			missingFieldsLeft4zkes--
			curField4zkes = decodeMsgFieldOrder4zkes[nextMiss4zkes]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zkes)
		switch curField4zkes {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zkes[0] = true
			var zyeg uint32
			zyeg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zyeg > 0 {
				z.m = make(map[string]*Tr, zyeg)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zyeg > 0 {
				zyeg--
				var zwok string
				var zctu *Tr
				zwok, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zctu != nil {
						dc.PushAlwaysNil()
						err = zctu.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zctu == nil {
						zctu = new(Tr)
					}
					err = zctu.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zwok] = zctu
			}
		case "s_zid01_str":
			found4zkes[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zkes[2] = true
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
	if nextMiss4zkes != -1 {
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
var decodeMsgFieldOrder4zkes = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zkes = []bool{false, false, false}

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
	for zwok, zctu := range z.m {
		err = en.WriteString(zwok)
		if err != nil {
			return
		}
		if zctu == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = zctu.EncodeMsg(en)
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
	for zwok, zctu := range z.m {
		o = msgp.AppendString(o, zwok)
		if zctu == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zctu.MarshalMsg(o)
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
	const maxFields5zqbl = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zqbl uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zqbl, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zqbl := totalEncodedFields5zqbl
	missingFieldsLeft5zqbl := maxFields5zqbl - totalEncodedFields5zqbl

	var nextMiss5zqbl int32 = -1
	var found5zqbl [maxFields5zqbl]bool
	var curField5zqbl string

doneWithStruct5zqbl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zqbl > 0 || missingFieldsLeft5zqbl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zqbl, missingFieldsLeft5zqbl, msgp.ShowFound(found5zqbl[:]), unmarshalMsgFieldOrder5zqbl)
		if encodedFieldsLeft5zqbl > 0 {
			encodedFieldsLeft5zqbl--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zqbl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zqbl < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zqbl = 0
			}
			for nextMiss5zqbl < maxFields5zqbl && (found5zqbl[nextMiss5zqbl] || unmarshalMsgFieldSkip5zqbl[nextMiss5zqbl]) {
				nextMiss5zqbl++
			}
			if nextMiss5zqbl == maxFields5zqbl {
				// filled all the empty fields!
				break doneWithStruct5zqbl
			}
			missingFieldsLeft5zqbl--
			curField5zqbl = unmarshalMsgFieldOrder5zqbl[nextMiss5zqbl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zqbl)
		switch curField5zqbl {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zqbl[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zdaw uint32
				zdaw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zdaw > 0 {
					z.m = make(map[string]*Tr, zdaw)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zdaw > 0 {
					var zwok string
					var zctu *Tr
					zdaw--
					zwok, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zctu != nil {
							zctu.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zctu {
								zctu.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zctu == nil {
								zctu = new(Tr)
							}
							bts, err = zctu.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zwok] = zctu
				}
			}
		case "s_zid01_str":
			found5zqbl[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zqbl[2] = true
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
	if nextMiss5zqbl != -1 {
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
var unmarshalMsgFieldOrder5zqbl = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zqbl = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zwok, zctu := range z.m {
			_ = zctu
			_ = zwok
			s += msgp.StringPrefixSize + len(zwok)
			if zctu == nil {
				s += msgp.NilSize
			} else {
				s += zctu.Msgsize()
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
		0x5f, 0x5f, 0x6d, 0x61, 0x70, 0x83, 0xa1, 0x75, 0x82, 0xaf,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x75, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x93,
		0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34,
		0x00, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6d,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6d,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xae, 0x6d,
		0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d,
		0x2a, 0x54, 0x72, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x18,
		0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x84,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x18, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f,
		0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f,
		0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaa, 0x52,
		0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x1c, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7, 0x50, 0x6f,
		0x69, 0x6e, 0x74, 0x65, 0x72, 0xab, 0x44, 0x6f, 0x6d, 0x61,
		0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b,
		0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x16, 0xa8, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa2, 0x54, 0x72, 0x87, 0xa8,
		0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x01, 0xb0,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x73, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x73, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x17,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x02, 0xb2, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b,
		0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69,
		0x36, 0x34, 0x02, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa1, 0x6e, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa1, 0x6e, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xaf, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
		0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f,
		0x11, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72,
		0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x11, 0xa8,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa5, 0x69,
		0x6e, 0x74, 0x36, 0x34, 0xa2, 0x54, 0x72, 0x82, 0xaf, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa2, 0x54, 0x72, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x73, 0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x96,
		0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34,
		0x00, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x55,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x55,
		0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xaf, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
		0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x5f, 0x02,
		0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x02, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0x85, 0xa8, 0x5a, 0x69, 0x64, 0x5f,
		0x5f, 0x69, 0x36, 0x34, 0xff, 0xb0, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa2, 0x65, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa1, 0x2d, 0xaf, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
		0x5f, 0x1c, 0xa9, 0x53, 0x6b, 0x69, 0x70, 0x5f, 0x5f, 0x62,
		0x6f, 0x6f, 0xc3, 0x87, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0x01, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa2, 0x4e, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73,
		0x74, 0x72, 0xa2, 0x4e, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xaf, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x5f, 0x5f, 0x17, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f,
		0x5f, 0x0d, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74,
		0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x0d,
		0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3,
		0x69, 0x6e, 0x74, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f,
		0x69, 0x36, 0x34, 0x02, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xb1, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xaf, 0x6d, 0x61,
		0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62,
		0x6f, 0x6f, 0x6c, 0xaf, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x5f, 0x18,
		0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x84,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x18, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x5f,
		0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f,
		0x5f, 0x02, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xaa, 0x52,
		0x61, 0x6e, 0x67, 0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82,
		0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x12, 0xa8, 0x53,
		0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f,
		0x6f, 0x6c, 0x86, 0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69,
		0x36, 0x34, 0x03, 0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa6, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0xb1, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6, 0x4e, 0x6f, 0x74, 0x79,
		0x65, 0x74, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
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
		0x5f, 0x5f, 0x69, 0x36, 0x34, 0x04, 0xb0, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
		0x73, 0x74, 0x72, 0xa4, 0x53, 0x65, 0x74, 0x6d, 0xb1, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa4, 0x53, 0x65, 0x74,
		0x6d, 0xb1, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa6,
		0x5b, 0x5d, 0x2a, 0x69, 0x6e, 0x6e, 0xaf, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x5f, 0x5f, 0x1a, 0xb2, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x5f, 0x70,
		0x74, 0x72, 0x83, 0xa6, 0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f,
		0x1a, 0xa8, 0x53, 0x74, 0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72,
		0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xab, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x83, 0xa6,
		0x4b, 0x69, 0x6e, 0x64, 0x5f, 0x5f, 0x1c, 0xa8, 0x53, 0x74,
		0x72, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69,
		0x6e, 0x74, 0x65, 0x72, 0xab, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x16, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x6e, 0xa3, 0x69,
		0x6e, 0x6e, 0x82, 0xaf, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
		0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa3,
		0x69, 0x6e, 0x6e, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
		0x5f, 0x5f, 0x73, 0x6c, 0x63, 0x92, 0x86, 0xa8, 0x5a, 0x69,
		0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x00, 0xb0, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6a, 0xb1, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa1, 0x6a, 0xb1, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c,
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
		0x65, 0x5f, 0x5f, 0x70, 0x74, 0x72, 0x82, 0xa6, 0x4b, 0x69,
		0x6e, 0x64, 0x5f, 0x5f, 0x12, 0xa8, 0x53, 0x74, 0x72, 0x5f,
		0x5f, 0x73, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x87,
		0xa8, 0x5a, 0x69, 0x64, 0x5f, 0x5f, 0x69, 0x36, 0x34, 0x01,
		0xb0, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x65, 0xb1,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0x5f, 0x5f, 0x73, 0x74, 0x72, 0xa1, 0x65, 0xb1,
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
	return []byte(`{"SourcePath__str":"my2.go","SourcePackage__str":"testdata","ZebraSchemaId__i64":0,"Structs__map":{"u":{"StructName__str":"u","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"m","FieldTagName__str":"m","FieldTypeStr__str":"map[string]*Tr","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"Tr"}}}},{"Zid__i64":1,"FieldGoName__str":"s","FieldTagName__str":"s","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":2,"FieldGoName__str":"n","FieldTagName__str":"n","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}}]},"Tr":{"StructName__str":"Tr","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"U","FieldTagName__str":"U","FieldTypeStr__str":"string","FieldCategory__":23,"FieldPrimitive__":2,"FieldFullType__ptr":{"Kind__":2,"Str__str":"string"}},{"Zid__i64":-1,"FieldGoName__str":"et","FieldTagName__str":"-","FieldCategory__":28,"Skip__boo":true},{"Zid__i64":1,"FieldGoName__str":"Nt","FieldTagName__str":"Nt","FieldTypeStr__str":"int","FieldCategory__":23,"FieldPrimitive__":13,"FieldFullType__ptr":{"Kind__":13,"Str__str":"int"}},{"Zid__i64":2,"FieldGoName__str":"Snot","FieldTagName__str":"Snot","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":3,"FieldGoName__str":"Notyet","FieldTagName__str":"Notyet","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":4,"FieldGoName__str":"Setm","FieldTagName__str":"Setm","FieldTypeStr__str":"[]*inn","FieldCategory__":26,"FieldFullType__ptr":{"Kind__":26,"Str__str":"Slice","Domain__ptr":{"Kind__":28,"Str__str":"Pointer","Domain__ptr":{"Kind__":22,"Str__str":"inn"}}}}]},"inn":{"StructName__str":"inn","Fields__slc":[{"Zid__i64":0,"FieldGoName__str":"j","FieldTagName__str":"j","FieldTypeStr__str":"map[string]bool","FieldCategory__":24,"FieldFullType__ptr":{"Kind__":24,"Str__str":"Map","Domain__ptr":{"Kind__":2,"Str__str":"string"},"Range__ptr":{"Kind__":18,"Str__str":"bool"}}},{"Zid__i64":1,"FieldGoName__str":"e","FieldTagName__str":"e","FieldTypeStr__str":"int64","FieldCategory__":23,"FieldPrimitive__":17,"FieldFullType__ptr":{"Kind__":17,"Str__str":"int64"}}]}},"Imports__slc":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the Greenpack Schema in pretty JSON format, length 6839 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath__str": "my2.go",
    "SourcePackage__str": "testdata",
    "ZebraSchemaId__i64": 0,
    "Structs__map": {
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
        }
    },
    "Imports__slc": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
