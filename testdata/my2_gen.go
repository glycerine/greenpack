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
	const maxFields0zwog = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zwog uint32
	totalEncodedFields0zwog, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zwog := totalEncodedFields0zwog
	missingFieldsLeft0zwog := maxFields0zwog - totalEncodedFields0zwog

	var nextMiss0zwog int32 = -1
	var found0zwog [maxFields0zwog]bool
	var curField0zwog string

doneWithStruct0zwog:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zwog > 0 || missingFieldsLeft0zwog > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zwog, missingFieldsLeft0zwog, msgp.ShowFound(found0zwog[:]), decodeMsgFieldOrder0zwog)
		if encodedFieldsLeft0zwog > 0 {
			encodedFieldsLeft0zwog--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zwog = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zwog < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zwog = 0
			}
			for nextMiss0zwog < maxFields0zwog && (found0zwog[nextMiss0zwog] || decodeMsgFieldSkip0zwog[nextMiss0zwog]) {
				nextMiss0zwog++
			}
			if nextMiss0zwog == maxFields0zwog {
				// filled all the empty fields!
				break doneWithStruct0zwog
			}
			missingFieldsLeft0zwog--
			curField0zwog = decodeMsgFieldOrder0zwog[nextMiss0zwog]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zwog)
		switch curField0zwog {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zwog[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zwog[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zwog[3] = true
			var zkyy uint32
			zkyy, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zkyy > 0 {
				z.Snot = make(map[string]bool, zkyy)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zkyy > 0 {
				zkyy--
				var zcbd string
				var zimi bool
				zcbd, err = dc.ReadString()
				if err != nil {
					return
				}
				zimi, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zcbd] = zimi
			}
		case "Notyet_zid03_map":
			found0zwog[4] = true
			var zgkp uint32
			zgkp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zgkp > 0 {
				z.Notyet = make(map[string]bool, zgkp)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zgkp > 0 {
				zgkp--
				var ztmf string
				var zgyh bool
				ztmf, err = dc.ReadString()
				if err != nil {
					return
				}
				zgyh, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[ztmf] = zgyh
			}
		case "Setm_zid04_slc":
			found0zwog[5] = true
			var zrsm uint32
			zrsm, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zrsm) {
				z.Setm = (z.Setm)[:zrsm]
			} else {
				z.Setm = make([]*inn, zrsm)
			}
			for zege := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zege] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zege].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zege] == nil {
						z.Setm[zege] = new(inn)
					}
					err = z.Setm[zege].DecodeMsg(dc)
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
	if nextMiss0zwog != -1 {
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
var decodeMsgFieldOrder0zwog = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zwog = []bool{false, true, false, false, false, false}

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
	var empty_zhdz [6]bool
	fieldsInUse_zqjs := z.fieldsNotEmpty(empty_zhdz[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqjs)
	if err != nil {
		return err
	}

	if !empty_zhdz[0] {
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

	if !empty_zhdz[2] {
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

	if !empty_zhdz[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zcbd, zimi := range z.Snot {
			err = en.WriteString(zcbd)
			if err != nil {
				return
			}
			err = en.WriteBool(zimi)
			if err != nil {
				return
			}
		}
	}

	if !empty_zhdz[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for ztmf, zgyh := range z.Notyet {
			err = en.WriteString(ztmf)
			if err != nil {
				return
			}
			err = en.WriteBool(zgyh)
			if err != nil {
				return
			}
		}
	}

	if !empty_zhdz[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zege := range z.Setm {
			if z.Setm[zege] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zege].EncodeMsg(en)
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
		for zcbd, zimi := range z.Snot {
			o = msgp.AppendString(o, zcbd)
			o = msgp.AppendBool(o, zimi)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for ztmf, zgyh := range z.Notyet {
			o = msgp.AppendString(o, ztmf)
			o = msgp.AppendBool(o, zgyh)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zege := range z.Setm {
			if z.Setm[zege] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zege].MarshalMsg(o)
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
	const maxFields1zxys = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zxys uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zxys, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zxys := totalEncodedFields1zxys
	missingFieldsLeft1zxys := maxFields1zxys - totalEncodedFields1zxys

	var nextMiss1zxys int32 = -1
	var found1zxys [maxFields1zxys]bool
	var curField1zxys string

doneWithStruct1zxys:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zxys > 0 || missingFieldsLeft1zxys > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zxys, missingFieldsLeft1zxys, msgp.ShowFound(found1zxys[:]), unmarshalMsgFieldOrder1zxys)
		if encodedFieldsLeft1zxys > 0 {
			encodedFieldsLeft1zxys--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zxys = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zxys < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zxys = 0
			}
			for nextMiss1zxys < maxFields1zxys && (found1zxys[nextMiss1zxys] || unmarshalMsgFieldSkip1zxys[nextMiss1zxys]) {
				nextMiss1zxys++
			}
			if nextMiss1zxys == maxFields1zxys {
				// filled all the empty fields!
				break doneWithStruct1zxys
			}
			missingFieldsLeft1zxys--
			curField1zxys = unmarshalMsgFieldOrder1zxys[nextMiss1zxys]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zxys)
		switch curField1zxys {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zxys[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zxys[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zxys[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zxnu uint32
				zxnu, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zxnu > 0 {
					z.Snot = make(map[string]bool, zxnu)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zxnu > 0 {
					var zcbd string
					var zimi bool
					zxnu--
					zcbd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zimi, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zcbd] = zimi
				}
			}
		case "Notyet_zid03_map":
			found1zxys[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zysd uint32
				zysd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zysd > 0 {
					z.Notyet = make(map[string]bool, zysd)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zysd > 0 {
					var ztmf string
					var zgyh bool
					zysd--
					ztmf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgyh, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[ztmf] = zgyh
				}
			}
		case "Setm_zid04_slc":
			found1zxys[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zpvu uint32
				zpvu, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zpvu) {
					z.Setm = (z.Setm)[:zpvu]
				} else {
					z.Setm = make([]*inn, zpvu)
				}
				for zege := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zege] != nil {
							z.Setm[zege].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zege] {
								z.Setm[zege].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zege] == nil {
								z.Setm[zege] = new(inn)
							}
							bts, err = z.Setm[zege].UnmarshalMsg(bts)
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
	if nextMiss1zxys != -1 {
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
var unmarshalMsgFieldOrder1zxys = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zxys = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zcbd, zimi := range z.Snot {
			_ = zimi
			_ = zcbd
			s += msgp.StringPrefixSize + len(zcbd) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for ztmf, zgyh := range z.Notyet {
			_ = zgyh
			_ = ztmf
			s += msgp.StringPrefixSize + len(ztmf) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for zege := range z.Setm {
		if z.Setm[zege] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zege].Msgsize()
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
	const maxFields2znuw = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2znuw uint32
	totalEncodedFields2znuw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2znuw := totalEncodedFields2znuw
	missingFieldsLeft2znuw := maxFields2znuw - totalEncodedFields2znuw

	var nextMiss2znuw int32 = -1
	var found2znuw [maxFields2znuw]bool
	var curField2znuw string

doneWithStruct2znuw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2znuw > 0 || missingFieldsLeft2znuw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2znuw, missingFieldsLeft2znuw, msgp.ShowFound(found2znuw[:]), decodeMsgFieldOrder2znuw)
		if encodedFieldsLeft2znuw > 0 {
			encodedFieldsLeft2znuw--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2znuw = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2znuw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2znuw = 0
			}
			for nextMiss2znuw < maxFields2znuw && (found2znuw[nextMiss2znuw] || decodeMsgFieldSkip2znuw[nextMiss2znuw]) {
				nextMiss2znuw++
			}
			if nextMiss2znuw == maxFields2znuw {
				// filled all the empty fields!
				break doneWithStruct2znuw
			}
			missingFieldsLeft2znuw--
			curField2znuw = decodeMsgFieldOrder2znuw[nextMiss2znuw]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2znuw)
		switch curField2znuw {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2znuw[0] = true
			var zegg uint32
			zegg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zegg > 0 {
				z.j = make(map[string]bool, zegg)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zegg > 0 {
				zegg--
				var zhth string
				var zuhx bool
				zhth, err = dc.ReadString()
				if err != nil {
					return
				}
				zuhx, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zhth] = zuhx
			}
		case "e_zid01_i64":
			found2znuw[1] = true
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
	if nextMiss2znuw != -1 {
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
var decodeMsgFieldOrder2znuw = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2znuw = []bool{false, false}

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
	var empty_zuzu [2]bool
	fieldsInUse_zeaj := z.fieldsNotEmpty(empty_zuzu[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zeaj)
	if err != nil {
		return err
	}

	if !empty_zuzu[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zhth, zuhx := range z.j {
			err = en.WriteString(zhth)
			if err != nil {
				return
			}
			err = en.WriteBool(zuhx)
			if err != nil {
				return
			}
		}
	}

	if !empty_zuzu[1] {
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
		for zhth, zuhx := range z.j {
			o = msgp.AppendString(o, zhth)
			o = msgp.AppendBool(o, zuhx)
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
	const maxFields3zyag = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zyag uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zyag, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zyag := totalEncodedFields3zyag
	missingFieldsLeft3zyag := maxFields3zyag - totalEncodedFields3zyag

	var nextMiss3zyag int32 = -1
	var found3zyag [maxFields3zyag]bool
	var curField3zyag string

doneWithStruct3zyag:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zyag > 0 || missingFieldsLeft3zyag > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zyag, missingFieldsLeft3zyag, msgp.ShowFound(found3zyag[:]), unmarshalMsgFieldOrder3zyag)
		if encodedFieldsLeft3zyag > 0 {
			encodedFieldsLeft3zyag--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zyag = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zyag < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zyag = 0
			}
			for nextMiss3zyag < maxFields3zyag && (found3zyag[nextMiss3zyag] || unmarshalMsgFieldSkip3zyag[nextMiss3zyag]) {
				nextMiss3zyag++
			}
			if nextMiss3zyag == maxFields3zyag {
				// filled all the empty fields!
				break doneWithStruct3zyag
			}
			missingFieldsLeft3zyag--
			curField3zyag = unmarshalMsgFieldOrder3zyag[nextMiss3zyag]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zyag)
		switch curField3zyag {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zyag[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zobo uint32
				zobo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zobo > 0 {
					z.j = make(map[string]bool, zobo)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zobo > 0 {
					var zhth string
					var zuhx bool
					zobo--
					zhth, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zuhx, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zhth] = zuhx
				}
			}
		case "e_zid01_i64":
			found3zyag[1] = true
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
	if nextMiss3zyag != -1 {
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
var unmarshalMsgFieldOrder3zyag = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zyag = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zhth, zuhx := range z.j {
			_ = zuhx
			_ = zhth
			s += msgp.StringPrefixSize + len(zhth) + msgp.BoolSize
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
	const maxFields4zegv = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zegv uint32
	totalEncodedFields4zegv, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zegv := totalEncodedFields4zegv
	missingFieldsLeft4zegv := maxFields4zegv - totalEncodedFields4zegv

	var nextMiss4zegv int32 = -1
	var found4zegv [maxFields4zegv]bool
	var curField4zegv string

doneWithStruct4zegv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zegv > 0 || missingFieldsLeft4zegv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zegv, missingFieldsLeft4zegv, msgp.ShowFound(found4zegv[:]), decodeMsgFieldOrder4zegv)
		if encodedFieldsLeft4zegv > 0 {
			encodedFieldsLeft4zegv--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zegv = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zegv < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zegv = 0
			}
			for nextMiss4zegv < maxFields4zegv && (found4zegv[nextMiss4zegv] || decodeMsgFieldSkip4zegv[nextMiss4zegv]) {
				nextMiss4zegv++
			}
			if nextMiss4zegv == maxFields4zegv {
				// filled all the empty fields!
				break doneWithStruct4zegv
			}
			missingFieldsLeft4zegv--
			curField4zegv = decodeMsgFieldOrder4zegv[nextMiss4zegv]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zegv)
		switch curField4zegv {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zegv[0] = true
			var zxca uint32
			zxca, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zxca > 0 {
				z.m = make(map[string]*Tr, zxca)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zxca > 0 {
				zxca--
				var zdea string
				var zpxl *Tr
				zdea, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zpxl != nil {
						dc.PushAlwaysNil()
						err = zpxl.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zpxl == nil {
						zpxl = new(Tr)
					}
					err = zpxl.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zdea] = zpxl
			}
		case "s_zid01_str":
			found4zegv[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zegv[2] = true
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
	if nextMiss4zegv != -1 {
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
var decodeMsgFieldOrder4zegv = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zegv = []bool{false, false, false}

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
	var empty_zprh [3]bool
	fieldsInUse_zjih := z.fieldsNotEmpty(empty_zprh[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjih)
	if err != nil {
		return err
	}

	if !empty_zprh[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zdea, zpxl := range z.m {
			err = en.WriteString(zdea)
			if err != nil {
				return
			}
			if zpxl == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zpxl.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zprh[1] {
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

	if !empty_zprh[2] {
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
		for zdea, zpxl := range z.m {
			o = msgp.AppendString(o, zdea)
			if zpxl == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zpxl.MarshalMsg(o)
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
	const maxFields5zpao = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zpao uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zpao, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zpao := totalEncodedFields5zpao
	missingFieldsLeft5zpao := maxFields5zpao - totalEncodedFields5zpao

	var nextMiss5zpao int32 = -1
	var found5zpao [maxFields5zpao]bool
	var curField5zpao string

doneWithStruct5zpao:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zpao > 0 || missingFieldsLeft5zpao > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zpao, missingFieldsLeft5zpao, msgp.ShowFound(found5zpao[:]), unmarshalMsgFieldOrder5zpao)
		if encodedFieldsLeft5zpao > 0 {
			encodedFieldsLeft5zpao--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zpao = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zpao < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zpao = 0
			}
			for nextMiss5zpao < maxFields5zpao && (found5zpao[nextMiss5zpao] || unmarshalMsgFieldSkip5zpao[nextMiss5zpao]) {
				nextMiss5zpao++
			}
			if nextMiss5zpao == maxFields5zpao {
				// filled all the empty fields!
				break doneWithStruct5zpao
			}
			missingFieldsLeft5zpao--
			curField5zpao = unmarshalMsgFieldOrder5zpao[nextMiss5zpao]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zpao)
		switch curField5zpao {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zpao[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zhqk uint32
				zhqk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zhqk > 0 {
					z.m = make(map[string]*Tr, zhqk)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zhqk > 0 {
					var zdea string
					var zpxl *Tr
					zhqk--
					zdea, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zpxl != nil {
							zpxl.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zpxl {
								zpxl.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zpxl == nil {
								zpxl = new(Tr)
							}
							bts, err = zpxl.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zdea] = zpxl
				}
			}
		case "s_zid01_str":
			found5zpao[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zpao[2] = true
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
	if nextMiss5zpao != -1 {
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
var unmarshalMsgFieldOrder5zpao = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zpao = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for zdea, zpxl := range z.m {
			_ = zpxl
			_ = zdea
			s += msgp.StringPrefixSize + len(zdea)
			if zpxl == nil {
				s += msgp.NilSize
			} else {
				s += zpxl.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
