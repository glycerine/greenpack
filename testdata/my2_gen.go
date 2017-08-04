package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
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
	const maxFields0zopw = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zopw uint32
	totalEncodedFields0zopw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zopw := totalEncodedFields0zopw
	missingFieldsLeft0zopw := maxFields0zopw - totalEncodedFields0zopw

	var nextMiss0zopw int = -1
	var found0zopw [maxFields0zopw]bool
	var curField0zopw int

doneWithStruct0zopw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zopw > 0 || missingFieldsLeft0zopw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zopw, missingFieldsLeft0zopw, msgp.ShowFound(found0zopw[:]), decodeMsgFieldOrder0zopw)
		if encodedFieldsLeft0zopw > 0 {
			encodedFieldsLeft0zopw--
			curField0zopw, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zopw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zopw = 0
			}
			for nextMiss0zopw < maxFields0zopw && (found0zopw[nextMiss0zopw] || decodeMsgFieldSkip0zopw[nextMiss0zopw]) {
				nextMiss0zopw++
			}
			if nextMiss0zopw == maxFields0zopw {
				// filled all the empty fields!
				break doneWithStruct0zopw
			}
			missingFieldsLeft0zopw--
			curField0zopw = nextMiss0zopw
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zopw)
		switch curField0zopw {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found0zopw[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found0zopw[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found0zopw[3] = true
			var zeia uint32
			zeia, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zeia > 0 {
				z.Snot = make(map[string]bool, zeia)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zeia > 0 {
				zeia--
				var ziqo string
				var zgap bool
				ziqo, err = dc.ReadString()
				if err != nil {
					return
				}
				zgap, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[ziqo] = zgap
			}
		case 3:
			// zid 3 for "Notyet"
			found0zopw[4] = true
			var zvil uint32
			zvil, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zvil > 0 {
				z.Notyet = make(map[string]bool, zvil)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zvil > 0 {
				zvil--
				var zyee string
				var zfad bool
				zyee, err = dc.ReadString()
				if err != nil {
					return
				}
				zfad, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zyee] = zfad
			}
		case 4:
			// zid 4 for "Setm"
			found0zopw[5] = true
			var zres uint32
			zres, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zres) {
				z.Setm = (z.Setm)[:zres]
			} else {
				z.Setm = make([]*inn, zres)
			}
			for zjsw := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zjsw] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zjsw].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zjsw] == nil {
						z.Setm[zjsw] = new(inn)
					}
					err = z.Setm[zjsw].DecodeMsg(dc)
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
	if nextMiss0zopw != -1 {
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
var decodeMsgFieldOrder0zopw = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var decodeMsgFieldSkip0zopw = []bool{false, true, false, false, false, false}

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
	var empty_zimb [6]bool
	fieldsInUse_zlty := z.fieldsNotEmpty(empty_zimb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlty)
	if err != nil {
		return err
	}

	if !empty_zimb[0] {
		// zid 0 for "U"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteString(z.U)
		if err != nil {
			return
		}
	}

	if !empty_zimb[2] {
		// zid 1 for "Nt"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Nt)
		if err != nil {
			return
		}
	}

	if !empty_zimb[3] {
		// zid 2 for "Snot"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for ziqo, zgap := range z.Snot {
			err = en.WriteString(ziqo)
			if err != nil {
				return
			}
			err = en.WriteBool(zgap)
			if err != nil {
				return
			}
		}
	}

	if !empty_zimb[4] {
		// zid 3 for "Notyet"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zyee, zfad := range z.Notyet {
			err = en.WriteString(zyee)
			if err != nil {
				return
			}
			err = en.WriteBool(zfad)
			if err != nil {
				return
			}
		}
	}

	if !empty_zimb[5] {
		// zid 4 for "Setm"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zjsw := range z.Setm {
			if z.Setm[zjsw] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zjsw].EncodeMsg(en)
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
		// zid 0 for "U"
		o = append(o, 0x0)
		o = msgp.AppendString(o, z.U)
	}

	if !empty[2] {
		// zid 1 for "Nt"
		o = append(o, 0x1)
		o = msgp.AppendInt(o, z.Nt)
	}

	if !empty[3] {
		// zid 2 for "Snot"
		o = append(o, 0x2)
		o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
		for ziqo, zgap := range z.Snot {
			o = msgp.AppendString(o, ziqo)
			o = msgp.AppendBool(o, zgap)
		}
	}

	if !empty[4] {
		// zid 3 for "Notyet"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zyee, zfad := range z.Notyet {
			o = msgp.AppendString(o, zyee)
			o = msgp.AppendBool(o, zfad)
		}
	}

	if !empty[5] {
		// zid 4 for "Setm"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zjsw := range z.Setm {
			if z.Setm[zjsw] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zjsw].MarshalMsg(o)
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
	const maxFields1zzxb = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zzxb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zzxb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zzxb := totalEncodedFields1zzxb
	missingFieldsLeft1zzxb := maxFields1zzxb - totalEncodedFields1zzxb

	var nextMiss1zzxb int = -1
	var found1zzxb [maxFields1zzxb]bool
	var curField1zzxb int

doneWithStruct1zzxb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zzxb > 0 || missingFieldsLeft1zzxb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zzxb, missingFieldsLeft1zzxb, msgp.ShowFound(found1zzxb[:]), unmarshalMsgFieldOrder1zzxb)
		if encodedFieldsLeft1zzxb > 0 {
			encodedFieldsLeft1zzxb--
			curField1zzxb, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zzxb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zzxb = 0
			}
			for nextMiss1zzxb < maxFields1zzxb && (found1zzxb[nextMiss1zzxb] || unmarshalMsgFieldSkip1zzxb[nextMiss1zzxb]) {
				nextMiss1zzxb++
			}
			if nextMiss1zzxb == maxFields1zzxb {
				// filled all the empty fields!
				break doneWithStruct1zzxb
			}
			missingFieldsLeft1zzxb--
			curField1zzxb = nextMiss1zzxb
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zzxb)
		switch curField1zzxb {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found1zzxb[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found1zzxb[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found1zzxb[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zuql uint32
				zuql, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zuql > 0 {
					z.Snot = make(map[string]bool, zuql)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zuql > 0 {
					var ziqo string
					var zgap bool
					zuql--
					ziqo, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgap, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[ziqo] = zgap
				}
			}
		case 3:
			// zid 3 for "Notyet"
			found1zzxb[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zsqi uint32
				zsqi, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zsqi > 0 {
					z.Notyet = make(map[string]bool, zsqi)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zsqi > 0 {
					var zyee string
					var zfad bool
					zsqi--
					zyee, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zfad, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zyee] = zfad
				}
			}
		case 4:
			// zid 4 for "Setm"
			found1zzxb[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zirz uint32
				zirz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zirz) {
					z.Setm = (z.Setm)[:zirz]
				} else {
					z.Setm = make([]*inn, zirz)
				}
				for zjsw := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zjsw] != nil {
							z.Setm[zjsw].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zjsw] {
								z.Setm[zjsw].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zjsw] == nil {
								z.Setm[zjsw] = new(inn)
							}
							bts, err = z.Setm[zjsw].UnmarshalMsg(bts)
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
	if nextMiss1zzxb != -1 {
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
var unmarshalMsgFieldOrder1zzxb = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var unmarshalMsgFieldSkip1zzxb = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.U) + 3 + msgp.IntSize + 5 + msgp.MapHeaderSize
	if z.Snot != nil {
		for ziqo, zgap := range z.Snot {
			_ = zgap
			_ = ziqo
			s += msgp.StringPrefixSize + len(ziqo) + msgp.BoolSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zyee, zfad := range z.Notyet {
			_ = zfad
			_ = zyee
			s += msgp.StringPrefixSize + len(zyee) + msgp.BoolSize
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zjsw := range z.Setm {
		if z.Setm[zjsw] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zjsw].Msgsize()
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
	const maxFields2ztzz = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2ztzz uint32
	totalEncodedFields2ztzz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2ztzz := totalEncodedFields2ztzz
	missingFieldsLeft2ztzz := maxFields2ztzz - totalEncodedFields2ztzz

	var nextMiss2ztzz int = -1
	var found2ztzz [maxFields2ztzz]bool
	var curField2ztzz int

doneWithStruct2ztzz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2ztzz > 0 || missingFieldsLeft2ztzz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2ztzz, missingFieldsLeft2ztzz, msgp.ShowFound(found2ztzz[:]), decodeMsgFieldOrder2ztzz)
		if encodedFieldsLeft2ztzz > 0 {
			encodedFieldsLeft2ztzz--
			curField2ztzz, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2ztzz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2ztzz = 0
			}
			for nextMiss2ztzz < maxFields2ztzz && (found2ztzz[nextMiss2ztzz] || decodeMsgFieldSkip2ztzz[nextMiss2ztzz]) {
				nextMiss2ztzz++
			}
			if nextMiss2ztzz == maxFields2ztzz {
				// filled all the empty fields!
				break doneWithStruct2ztzz
			}
			missingFieldsLeft2ztzz--
			curField2ztzz = nextMiss2ztzz
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2ztzz)
		switch curField2ztzz {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found2ztzz[0] = true
			var zjlv uint32
			zjlv, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zjlv > 0 {
				z.j = make(map[string]bool, zjlv)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zjlv > 0 {
				zjlv--
				var zqeo string
				var zamo bool
				zqeo, err = dc.ReadString()
				if err != nil {
					return
				}
				zamo, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zqeo] = zamo
			}
		case 1:
			// zid 1 for "e"
			found2ztzz[1] = true
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
	if nextMiss2ztzz != -1 {
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
var decodeMsgFieldOrder2ztzz = []string{"j", "e"}

var decodeMsgFieldSkip2ztzz = []bool{false, false}

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
	var empty_zwca [2]bool
	fieldsInUse_zfct := z.fieldsNotEmpty(empty_zwca[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zfct)
	if err != nil {
		return err
	}

	if !empty_zwca[0] {
		// zid 0 for "j"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zqeo, zamo := range z.j {
			err = en.WriteString(zqeo)
			if err != nil {
				return
			}
			err = en.WriteBool(zamo)
			if err != nil {
				return
			}
		}
	}

	if !empty_zwca[1] {
		// zid 1 for "e"
		err = en.Append(0x1)
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
		// zid 0 for "j"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.j)))
		for zqeo, zamo := range z.j {
			o = msgp.AppendString(o, zqeo)
			o = msgp.AppendBool(o, zamo)
		}
	}

	if !empty[1] {
		// zid 1 for "e"
		o = append(o, 0x1)
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
	const maxFields3zrel = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zrel uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zrel, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zrel := totalEncodedFields3zrel
	missingFieldsLeft3zrel := maxFields3zrel - totalEncodedFields3zrel

	var nextMiss3zrel int = -1
	var found3zrel [maxFields3zrel]bool
	var curField3zrel int

doneWithStruct3zrel:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zrel > 0 || missingFieldsLeft3zrel > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zrel, missingFieldsLeft3zrel, msgp.ShowFound(found3zrel[:]), unmarshalMsgFieldOrder3zrel)
		if encodedFieldsLeft3zrel > 0 {
			encodedFieldsLeft3zrel--
			curField3zrel, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zrel < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zrel = 0
			}
			for nextMiss3zrel < maxFields3zrel && (found3zrel[nextMiss3zrel] || unmarshalMsgFieldSkip3zrel[nextMiss3zrel]) {
				nextMiss3zrel++
			}
			if nextMiss3zrel == maxFields3zrel {
				// filled all the empty fields!
				break doneWithStruct3zrel
			}
			missingFieldsLeft3zrel--
			curField3zrel = nextMiss3zrel
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zrel)
		switch curField3zrel {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found3zrel[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zlad uint32
				zlad, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zlad > 0 {
					z.j = make(map[string]bool, zlad)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zlad > 0 {
					var zqeo string
					var zamo bool
					zlad--
					zqeo, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zamo, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zqeo] = zamo
				}
			}
		case 1:
			// zid 1 for "e"
			found3zrel[1] = true
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
	if nextMiss3zrel != -1 {
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
var unmarshalMsgFieldOrder3zrel = []string{"j", "e"}

var unmarshalMsgFieldSkip3zrel = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.j != nil {
		for zqeo, zamo := range z.j {
			_ = zamo
			_ = zqeo
			s += msgp.StringPrefixSize + len(zqeo) + msgp.BoolSize
		}
	}
	s += 2 + msgp.Int64Size
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
	const maxFields4zfsu = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zfsu uint32
	totalEncodedFields4zfsu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zfsu := totalEncodedFields4zfsu
	missingFieldsLeft4zfsu := maxFields4zfsu - totalEncodedFields4zfsu

	var nextMiss4zfsu int = -1
	var found4zfsu [maxFields4zfsu]bool
	var curField4zfsu int

doneWithStruct4zfsu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zfsu > 0 || missingFieldsLeft4zfsu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zfsu, missingFieldsLeft4zfsu, msgp.ShowFound(found4zfsu[:]), decodeMsgFieldOrder4zfsu)
		if encodedFieldsLeft4zfsu > 0 {
			encodedFieldsLeft4zfsu--
			curField4zfsu, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zfsu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zfsu = 0
			}
			for nextMiss4zfsu < maxFields4zfsu && (found4zfsu[nextMiss4zfsu] || decodeMsgFieldSkip4zfsu[nextMiss4zfsu]) {
				nextMiss4zfsu++
			}
			if nextMiss4zfsu == maxFields4zfsu {
				// filled all the empty fields!
				break doneWithStruct4zfsu
			}
			missingFieldsLeft4zfsu--
			curField4zfsu = nextMiss4zfsu
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zfsu)
		switch curField4zfsu {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found4zfsu[0] = true
			var zimy uint32
			zimy, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zimy > 0 {
				z.m = make(map[string]*Tr, zimy)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zimy > 0 {
				zimy--
				var zytd string
				var zhnq *Tr
				zytd, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zhnq != nil {
						dc.PushAlwaysNil()
						err = zhnq.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zhnq == nil {
						zhnq = new(Tr)
					}
					err = zhnq.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zytd] = zhnq
			}
		case 1:
			// zid 1 for "s"
			found4zfsu[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found4zfsu[2] = true
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
	if nextMiss4zfsu != -1 {
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
var decodeMsgFieldOrder4zfsu = []string{"m", "s", "n"}

var decodeMsgFieldSkip4zfsu = []bool{false, false, false}

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
	var empty_ztqt [3]bool
	fieldsInUse_zujo := z.fieldsNotEmpty(empty_ztqt[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zujo)
	if err != nil {
		return err
	}

	if !empty_ztqt[0] {
		// zid 0 for "m"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zytd, zhnq := range z.m {
			err = en.WriteString(zytd)
			if err != nil {
				return
			}
			if zhnq == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zhnq.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_ztqt[1] {
		// zid 1 for "s"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.s)
		if err != nil {
			return
		}
	}

	if !empty_ztqt[2] {
		// zid 2 for "n"
		err = en.Append(0x2)
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
		// zid 0 for "m"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.m)))
		for zytd, zhnq := range z.m {
			o = msgp.AppendString(o, zytd)
			if zhnq == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zhnq.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[1] {
		// zid 1 for "s"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.s)
	}

	if !empty[2] {
		// zid 2 for "n"
		o = append(o, 0x2)
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
	const maxFields5zpqy = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zpqy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zpqy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zpqy := totalEncodedFields5zpqy
	missingFieldsLeft5zpqy := maxFields5zpqy - totalEncodedFields5zpqy

	var nextMiss5zpqy int = -1
	var found5zpqy [maxFields5zpqy]bool
	var curField5zpqy int

doneWithStruct5zpqy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zpqy > 0 || missingFieldsLeft5zpqy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zpqy, missingFieldsLeft5zpqy, msgp.ShowFound(found5zpqy[:]), unmarshalMsgFieldOrder5zpqy)
		if encodedFieldsLeft5zpqy > 0 {
			encodedFieldsLeft5zpqy--
			curField5zpqy, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zpqy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zpqy = 0
			}
			for nextMiss5zpqy < maxFields5zpqy && (found5zpqy[nextMiss5zpqy] || unmarshalMsgFieldSkip5zpqy[nextMiss5zpqy]) {
				nextMiss5zpqy++
			}
			if nextMiss5zpqy == maxFields5zpqy {
				// filled all the empty fields!
				break doneWithStruct5zpqy
			}
			missingFieldsLeft5zpqy--
			curField5zpqy = nextMiss5zpqy
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zpqy)
		switch curField5zpqy {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found5zpqy[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var znvl uint32
				znvl, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && znvl > 0 {
					z.m = make(map[string]*Tr, znvl)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for znvl > 0 {
					var zytd string
					var zhnq *Tr
					znvl--
					zytd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zhnq != nil {
							zhnq.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zhnq {
								zhnq.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zhnq == nil {
								zhnq = new(Tr)
							}
							bts, err = zhnq.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zytd] = zhnq
				}
			}
		case 1:
			// zid 1 for "s"
			found5zpqy[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found5zpqy[2] = true
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
	if nextMiss5zpqy != -1 {
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
var unmarshalMsgFieldOrder5zpqy = []string{"m", "s", "n"}

var unmarshalMsgFieldSkip5zpqy = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.m != nil {
		for zytd, zhnq := range z.m {
			_ = zhnq
			_ = zytd
			s += msgp.StringPrefixSize + len(zytd)
			if zhnq == nil {
				s += msgp.NilSize
			} else {
				s += zhnq.Msgsize()
			}
		}
	}
	s += 2 + msgp.StringPrefixSize + len(z.s) + 2 + msgp.Int64Size
	return
}

// FileMy2 holds ZebraPack schema from file 'my2.go'
type FileMy2 struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 1636 bytes
func (FileMy2) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xa6, 0x6d, 0x79, 0x32, 0x2e, 0x67, 0x6f, 0xad,
		0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
		0x61, 0x67, 0x65, 0xa8, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
		0x74, 0x61, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x00, 0xa7, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x73, 0x83, 0xa1, 0x75, 0x82, 0xaa,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x75, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x93,
		0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6d,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xae, 0x6d,
		0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d,
		0x2a, 0x54, 0x72, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7,
		0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x16, 0xa3, 0x53, 0x74, 0x72, 0xa2, 0x54, 0x72, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x73, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa1, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69,
		0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6e, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74,
		0x36, 0x34, 0xa2, 0x54, 0x72, 0x82, 0xaa, 0x53, 0x74, 0x72,
		0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x54, 0x72,
		0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x55, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa1, 0x55, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x85, 0xa3, 0x5a, 0x69,
		0x64, 0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x65, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x2d, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xa4, 0x53, 0x6b,
		0x69, 0x70, 0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa2, 0x4e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x4e, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0d, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0d, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0x86, 0xa3, 0x5a,
		0x69, 0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e,
		0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3,
		0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa3,
		0x5a, 0x69, 0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74,
		0x79, 0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74,
		0x79, 0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70,
		0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f,
		0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53,
		0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02,
		0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62,
		0x6f, 0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa4, 0x53, 0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
		0x53, 0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x5b, 0x5d,
		0x2a, 0x69, 0x6e, 0x6e, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a,
		0xa3, 0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x83, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50,
		0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16,
		0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x6e, 0xa3, 0x69,
		0x6e, 0x6e, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
		0x4e, 0x61, 0x6d, 0x65, 0xa3, 0x69, 0x6e, 0x6e, 0xa6, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x73, 0x92, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6a, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x6a, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72,
		0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69,
		0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53,
		0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5,
		0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f,
		0x6c, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x65, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17,
		0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74,
		0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa7, 0x49, 0x6d,
		0x70, 0x6f, 0x72, 0x74, 0x73, 0x91, 0xbd, 0x22, 0x67, 0x69,
		0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
		0x6c, 0x79, 0x63, 0x65, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x72,
		0x62, 0x74, 0x72, 0x65, 0x65, 0x22,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 2153 bytes
func (FileMy2) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"my2.go","SourcePackage":"testdata","ZebraSchemaId":0,"Structs":{"u":{"StructName":"u","Fields":[{"Zid":0,"FieldGoName":"m","FieldTagName":"m","FieldTypeStr":"map[string]*Tr","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Tr"}}}},{"Zid":1,"FieldGoName":"s","FieldTagName":"s","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"n","FieldTagName":"n","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Tr":{"StructName":"Tr","Fields":[{"Zid":0,"FieldGoName":"U","FieldTagName":"U","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"et","FieldTagName":"-","FieldCategory":28,"Skip":true},{"Zid":1,"FieldGoName":"Nt","FieldTagName":"Nt","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}},{"Zid":2,"FieldGoName":"Snot","FieldTagName":"Snot","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":3,"FieldGoName":"Notyet","FieldTagName":"Notyet","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":4,"FieldGoName":"Setm","FieldTagName":"Setm","FieldTypeStr":"[]*inn","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"inn"}}}}]},"inn":{"StructName":"inn","Fields":[{"Zid":0,"FieldGoName":"j","FieldTagName":"j","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":1,"FieldGoName":"e","FieldTagName":"e","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]}},"Imports":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 6280 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "my2.go",
    "SourcePackage": "testdata",
    "ZebraSchemaId": 0,
    "Structs": {
        "u": {
            "StructName": "u",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "m",
                    "FieldTagName": "m",
                    "FieldTypeStr": "map[string]*Tr",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "Tr"
                            }
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "s",
                    "FieldTagName": "s",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "n",
                    "FieldTagName": "n",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        },
        "Tr": {
            "StructName": "Tr",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "U",
                    "FieldTagName": "U",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "et",
                    "FieldTagName": "-",
                    "FieldCategory": 28,
                    "Skip": true
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Nt",
                    "FieldTagName": "Nt",
                    "FieldTypeStr": "int",
                    "FieldCategory": 23,
                    "FieldPrimitive": 13,
                    "FieldFullType": {
                        "Kind": 13,
                        "Str": "int"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Snot",
                    "FieldTagName": "Snot",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 3,
                    "FieldGoName": "Notyet",
                    "FieldTagName": "Notyet",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 4,
                    "FieldGoName": "Setm",
                    "FieldTagName": "Setm",
                    "FieldTypeStr": "[]*inn",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "inn"
                            }
                        }
                    }
                }
            ]
        },
        "inn": {
            "StructName": "inn",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "j",
                    "FieldTagName": "j",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "e",
                    "FieldTagName": "e",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        }
    },
    "Imports": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
