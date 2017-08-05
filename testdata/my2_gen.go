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
	const maxFields0zugz = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zugz uint32
	totalEncodedFields0zugz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zugz := totalEncodedFields0zugz
	missingFieldsLeft0zugz := maxFields0zugz - totalEncodedFields0zugz

	var nextMiss0zugz int32 = -1
	var found0zugz [maxFields0zugz]bool
	var curField0zugz string

doneWithStruct0zugz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zugz > 0 || missingFieldsLeft0zugz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zugz, missingFieldsLeft0zugz, msgp.ShowFound(found0zugz[:]), decodeMsgFieldOrder0zugz)
		if encodedFieldsLeft0zugz > 0 {
			encodedFieldsLeft0zugz--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zugz = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zugz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zugz = 0
			}
			for nextMiss0zugz < maxFields0zugz && (found0zugz[nextMiss0zugz] || decodeMsgFieldSkip0zugz[nextMiss0zugz]) {
				nextMiss0zugz++
			}
			if nextMiss0zugz == maxFields0zugz {
				// filled all the empty fields!
				break doneWithStruct0zugz
			}
			missingFieldsLeft0zugz--
			curField0zugz = decodeMsgFieldOrder0zugz[nextMiss0zugz]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zugz)
		switch curField0zugz {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zugz[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zugz[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zugz[3] = true
			var zgja uint32
			zgja, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zgja > 0 {
				z.Snot = make(map[string]bool, zgja)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zgja > 0 {
				zgja--
				var zgol string
				var zbya bool
				zgol, err = dc.ReadString()
				if err != nil {
					return
				}
				zbya, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zgol] = zbya
			}
		case "Notyet_zid03_map":
			found0zugz[4] = true
			var zwkn uint32
			zwkn, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zwkn > 0 {
				z.Notyet = make(map[string]bool, zwkn)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zwkn > 0 {
				zwkn--
				var zzbs string
				var zulo bool
				zzbs, err = dc.ReadString()
				if err != nil {
					return
				}
				zulo, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zzbs] = zulo
			}
		case "Setm_zid04_slc":
			found0zugz[5] = true
			var zanu uint32
			zanu, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zanu) {
				z.Setm = (z.Setm)[:zanu]
			} else {
				z.Setm = make([]*inn, zanu)
			}
			for zqou := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zqou] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zqou].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zqou] == nil {
						z.Setm[zqou] = new(inn)
					}
					err = z.Setm[zqou].DecodeMsg(dc)
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
	if nextMiss0zugz != -1 {
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
var decodeMsgFieldOrder0zugz = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zugz = []bool{false, true, false, false, false, false}

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
	var empty_zmsx [6]bool
	fieldsInUse_zuni := z.fieldsNotEmpty(empty_zmsx[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zuni)
	if err != nil {
		return err
	}

	if !empty_zmsx[0] {
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

	if !empty_zmsx[2] {
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

	if !empty_zmsx[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zgol, zbya := range z.Snot {
			err = en.WriteString(zgol)
			if err != nil {
				return
			}
			err = en.WriteBool(zbya)
			if err != nil {
				return
			}
		}
	}

	if !empty_zmsx[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zzbs, zulo := range z.Notyet {
			err = en.WriteString(zzbs)
			if err != nil {
				return
			}
			err = en.WriteBool(zulo)
			if err != nil {
				return
			}
		}
	}

	if !empty_zmsx[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zqou := range z.Setm {
			if z.Setm[zqou] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zqou].EncodeMsg(en)
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
		for zgol, zbya := range z.Snot {
			o = msgp.AppendString(o, zgol)
			o = msgp.AppendBool(o, zbya)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zzbs, zulo := range z.Notyet {
			o = msgp.AppendString(o, zzbs)
			o = msgp.AppendBool(o, zulo)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zqou := range z.Setm {
			if z.Setm[zqou] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zqou].MarshalMsg(o)
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
	const maxFields1zukr = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zukr uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zukr, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zukr := totalEncodedFields1zukr
	missingFieldsLeft1zukr := maxFields1zukr - totalEncodedFields1zukr

	var nextMiss1zukr int32 = -1
	var found1zukr [maxFields1zukr]bool
	var curField1zukr string

doneWithStruct1zukr:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zukr > 0 || missingFieldsLeft1zukr > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zukr, missingFieldsLeft1zukr, msgp.ShowFound(found1zukr[:]), unmarshalMsgFieldOrder1zukr)
		if encodedFieldsLeft1zukr > 0 {
			encodedFieldsLeft1zukr--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zukr = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zukr < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zukr = 0
			}
			for nextMiss1zukr < maxFields1zukr && (found1zukr[nextMiss1zukr] || unmarshalMsgFieldSkip1zukr[nextMiss1zukr]) {
				nextMiss1zukr++
			}
			if nextMiss1zukr == maxFields1zukr {
				// filled all the empty fields!
				break doneWithStruct1zukr
			}
			missingFieldsLeft1zukr--
			curField1zukr = unmarshalMsgFieldOrder1zukr[nextMiss1zukr]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zukr)
		switch curField1zukr {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zukr[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zukr[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zukr[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zjyv uint32
				zjyv, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zjyv > 0 {
					z.Snot = make(map[string]bool, zjyv)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zjyv > 0 {
					var zgol string
					var zbya bool
					zjyv--
					zgol, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zbya, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zgol] = zbya
				}
			}
		case "Notyet_zid03_map":
			found1zukr[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zcyc uint32
				zcyc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zcyc > 0 {
					z.Notyet = make(map[string]bool, zcyc)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zcyc > 0 {
					var zzbs string
					var zulo bool
					zcyc--
					zzbs, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zulo, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zzbs] = zulo
				}
			}
		case "Setm_zid04_slc":
			found1zukr[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zyuv uint32
				zyuv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zyuv) {
					z.Setm = (z.Setm)[:zyuv]
				} else {
					z.Setm = make([]*inn, zyuv)
				}
				for zqou := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zqou] != nil {
							z.Setm[zqou].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zqou] {
								z.Setm[zqou].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zqou] == nil {
								z.Setm[zqou] = new(inn)
							}
							bts, err = z.Setm[zqou].UnmarshalMsg(bts)
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
	if nextMiss1zukr != -1 {
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
var unmarshalMsgFieldOrder1zukr = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zukr = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zgol, zbya := range z.Snot {
			_ = zbya
			_ = zgol
			s += msgp.StringPrefixSize + len(zgol) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zzbs, zulo := range z.Notyet {
			_ = zulo
			_ = zzbs
			s += msgp.StringPrefixSize + len(zzbs) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zqou := range z.Setm {
		if z.Setm[zqou] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zqou].Msgsize()
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
	const maxFields2zgin = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zgin uint32
	totalEncodedFields2zgin, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zgin := totalEncodedFields2zgin
	missingFieldsLeft2zgin := maxFields2zgin - totalEncodedFields2zgin

	var nextMiss2zgin int32 = -1
	var found2zgin [maxFields2zgin]bool
	var curField2zgin string

doneWithStruct2zgin:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zgin > 0 || missingFieldsLeft2zgin > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zgin, missingFieldsLeft2zgin, msgp.ShowFound(found2zgin[:]), decodeMsgFieldOrder2zgin)
		if encodedFieldsLeft2zgin > 0 {
			encodedFieldsLeft2zgin--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zgin = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zgin < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zgin = 0
			}
			for nextMiss2zgin < maxFields2zgin && (found2zgin[nextMiss2zgin] || decodeMsgFieldSkip2zgin[nextMiss2zgin]) {
				nextMiss2zgin++
			}
			if nextMiss2zgin == maxFields2zgin {
				// filled all the empty fields!
				break doneWithStruct2zgin
			}
			missingFieldsLeft2zgin--
			curField2zgin = decodeMsgFieldOrder2zgin[nextMiss2zgin]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zgin)
		switch curField2zgin {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zgin[0] = true
			var zvmt uint32
			zvmt, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zvmt > 0 {
				z.j = make(map[string]bool, zvmt)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zvmt > 0 {
				zvmt--
				var zgvs string
				var zohy bool
				zgvs, err = dc.ReadString()
				if err != nil {
					return
				}
				zohy, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zgvs] = zohy
			}
		case "e_zid01_i64":
			found2zgin[1] = true
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
	if nextMiss2zgin != -1 {
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
var decodeMsgFieldOrder2zgin = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zgin = []bool{false, false}

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
	var empty_zklp [2]bool
	fieldsInUse_zqfg := z.fieldsNotEmpty(empty_zklp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqfg)
	if err != nil {
		return err
	}

	if !empty_zklp[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zgvs, zohy := range z.j {
			err = en.WriteString(zgvs)
			if err != nil {
				return
			}
			err = en.WriteBool(zohy)
			if err != nil {
				return
			}
		}
	}

	if !empty_zklp[1] {
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
		for zgvs, zohy := range z.j {
			o = msgp.AppendString(o, zgvs)
			o = msgp.AppendBool(o, zohy)
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
	const maxFields3zeln = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zeln uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zeln, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zeln := totalEncodedFields3zeln
	missingFieldsLeft3zeln := maxFields3zeln - totalEncodedFields3zeln

	var nextMiss3zeln int32 = -1
	var found3zeln [maxFields3zeln]bool
	var curField3zeln string

doneWithStruct3zeln:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zeln > 0 || missingFieldsLeft3zeln > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zeln, missingFieldsLeft3zeln, msgp.ShowFound(found3zeln[:]), unmarshalMsgFieldOrder3zeln)
		if encodedFieldsLeft3zeln > 0 {
			encodedFieldsLeft3zeln--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zeln = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zeln < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zeln = 0
			}
			for nextMiss3zeln < maxFields3zeln && (found3zeln[nextMiss3zeln] || unmarshalMsgFieldSkip3zeln[nextMiss3zeln]) {
				nextMiss3zeln++
			}
			if nextMiss3zeln == maxFields3zeln {
				// filled all the empty fields!
				break doneWithStruct3zeln
			}
			missingFieldsLeft3zeln--
			curField3zeln = unmarshalMsgFieldOrder3zeln[nextMiss3zeln]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zeln)
		switch curField3zeln {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zeln[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zqmi uint32
				zqmi, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zqmi > 0 {
					z.j = make(map[string]bool, zqmi)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zqmi > 0 {
					var zgvs string
					var zohy bool
					zqmi--
					zgvs, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zohy, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zgvs] = zohy
				}
			}
		case "e_zid01_i64":
			found3zeln[1] = true
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
	if nextMiss3zeln != -1 {
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
var unmarshalMsgFieldOrder3zeln = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zeln = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zgvs, zohy := range z.j {
			_ = zohy
			_ = zgvs
			s += msgp.StringPrefixSize + len(zgvs) + msgp.BoolSize
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
	const maxFields4zbos = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zbos uint32
	totalEncodedFields4zbos, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zbos := totalEncodedFields4zbos
	missingFieldsLeft4zbos := maxFields4zbos - totalEncodedFields4zbos

	var nextMiss4zbos int32 = -1
	var found4zbos [maxFields4zbos]bool
	var curField4zbos string

doneWithStruct4zbos:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zbos > 0 || missingFieldsLeft4zbos > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zbos, missingFieldsLeft4zbos, msgp.ShowFound(found4zbos[:]), decodeMsgFieldOrder4zbos)
		if encodedFieldsLeft4zbos > 0 {
			encodedFieldsLeft4zbos--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zbos = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zbos < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zbos = 0
			}
			for nextMiss4zbos < maxFields4zbos && (found4zbos[nextMiss4zbos] || decodeMsgFieldSkip4zbos[nextMiss4zbos]) {
				nextMiss4zbos++
			}
			if nextMiss4zbos == maxFields4zbos {
				// filled all the empty fields!
				break doneWithStruct4zbos
			}
			missingFieldsLeft4zbos--
			curField4zbos = decodeMsgFieldOrder4zbos[nextMiss4zbos]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zbos)
		switch curField4zbos {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zbos[0] = true
			var zshq uint32
			zshq, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zshq > 0 {
				z.m = make(map[string]*Tr, zshq)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zshq > 0 {
				zshq--
				var zdzv string
				var zxel *Tr
				zdzv, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zxel != nil {
						dc.PushAlwaysNil()
						err = zxel.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zxel == nil {
						zxel = new(Tr)
					}
					err = zxel.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zdzv] = zxel
			}
		case "s_zid01_str":
			found4zbos[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zbos[2] = true
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
	if nextMiss4zbos != -1 {
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
var decodeMsgFieldOrder4zbos = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zbos = []bool{false, false, false}

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
	var empty_zwqq [3]bool
	fieldsInUse_zlqn := z.fieldsNotEmpty(empty_zwqq[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zlqn)
	if err != nil {
		return err
	}

	if !empty_zwqq[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zdzv, zxel := range z.m {
			err = en.WriteString(zdzv)
			if err != nil {
				return
			}
			if zxel == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zxel.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zwqq[1] {
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

	if !empty_zwqq[2] {
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
		for zdzv, zxel := range z.m {
			o = msgp.AppendString(o, zdzv)
			if zxel == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zxel.MarshalMsg(o)
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
	const maxFields5zgyb = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zgyb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zgyb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zgyb := totalEncodedFields5zgyb
	missingFieldsLeft5zgyb := maxFields5zgyb - totalEncodedFields5zgyb

	var nextMiss5zgyb int32 = -1
	var found5zgyb [maxFields5zgyb]bool
	var curField5zgyb string

doneWithStruct5zgyb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zgyb > 0 || missingFieldsLeft5zgyb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zgyb, missingFieldsLeft5zgyb, msgp.ShowFound(found5zgyb[:]), unmarshalMsgFieldOrder5zgyb)
		if encodedFieldsLeft5zgyb > 0 {
			encodedFieldsLeft5zgyb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zgyb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zgyb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zgyb = 0
			}
			for nextMiss5zgyb < maxFields5zgyb && (found5zgyb[nextMiss5zgyb] || unmarshalMsgFieldSkip5zgyb[nextMiss5zgyb]) {
				nextMiss5zgyb++
			}
			if nextMiss5zgyb == maxFields5zgyb {
				// filled all the empty fields!
				break doneWithStruct5zgyb
			}
			missingFieldsLeft5zgyb--
			curField5zgyb = unmarshalMsgFieldOrder5zgyb[nextMiss5zgyb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zgyb)
		switch curField5zgyb {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zgyb[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zbuc uint32
				zbuc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zbuc > 0 {
					z.m = make(map[string]*Tr, zbuc)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zbuc > 0 {
					var zdzv string
					var zxel *Tr
					zbuc--
					zdzv, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zxel != nil {
							zxel.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zxel {
								zxel.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zxel == nil {
								zxel = new(Tr)
							}
							bts, err = zxel.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zdzv] = zxel
				}
			}
		case "s_zid01_str":
			found5zgyb[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zgyb[2] = true
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
	if nextMiss5zgyb != -1 {
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
var unmarshalMsgFieldOrder5zgyb = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zgyb = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zdzv, zxel := range z.m {
			_ = zxel
			_ = zdzv
			s += msgp.StringPrefixSize + len(zdzv)
			if zxel == nil {
				s += msgp.NilSize
			} else {
				s += zxel.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
