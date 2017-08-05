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
	const maxFields0zbec = 6

	// -- templateDecodeMsg starts here--
	var totalEncodedFields0zbec uint32
	totalEncodedFields0zbec, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zbec := totalEncodedFields0zbec
	missingFieldsLeft0zbec := maxFields0zbec - totalEncodedFields0zbec

	var nextMiss0zbec int32 = -1
	var found0zbec [maxFields0zbec]bool
	var curField0zbec string

doneWithStruct0zbec:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zbec > 0 || missingFieldsLeft0zbec > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zbec, missingFieldsLeft0zbec, msgp.ShowFound(found0zbec[:]), decodeMsgFieldOrder0zbec)
		if encodedFieldsLeft0zbec > 0 {
			encodedFieldsLeft0zbec--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField0zbec = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zbec < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zbec = 0
			}
			for nextMiss0zbec < maxFields0zbec && (found0zbec[nextMiss0zbec] || decodeMsgFieldSkip0zbec[nextMiss0zbec]) {
				nextMiss0zbec++
			}
			if nextMiss0zbec == maxFields0zbec {
				// filled all the empty fields!
				break doneWithStruct0zbec
			}
			missingFieldsLeft0zbec--
			curField0zbec = decodeMsgFieldOrder0zbec[nextMiss0zbec]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zbec)
		switch curField0zbec {
		// -- templateDecodeMsg ends here --

		case "U_zid00_str":
			found0zbec[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found0zbec[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found0zbec[3] = true
			var zgus uint32
			zgus, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zgus > 0 {
				z.Snot = make(map[string]bool, zgus)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zgus > 0 {
				zgus--
				var zshl string
				var zoxw bool
				zshl, err = dc.ReadString()
				if err != nil {
					return
				}
				zoxw, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zshl] = zoxw
			}
		case "Notyet_zid03_map":
			found0zbec[4] = true
			var zupf uint32
			zupf, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zupf > 0 {
				z.Notyet = make(map[string]bool, zupf)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zupf > 0 {
				zupf--
				var zdjo string
				var zvbh bool
				zdjo, err = dc.ReadString()
				if err != nil {
					return
				}
				zvbh, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zdjo] = zvbh
			}
		case "Setm_zid04_slc":
			found0zbec[5] = true
			var zfpp uint32
			zfpp, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zfpp) {
				z.Setm = (z.Setm)[:zfpp]
			} else {
				z.Setm = make([]*inn, zfpp)
			}
			for znnc := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[znnc] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[znnc].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[znnc] == nil {
						z.Setm[znnc] = new(inn)
					}
					err = z.Setm[znnc].DecodeMsg(dc)
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
	if nextMiss0zbec != -1 {
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
var decodeMsgFieldOrder0zbec = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var decodeMsgFieldSkip0zbec = []bool{false, true, false, false, false, false}

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
	var empty_zqar [6]bool
	fieldsInUse_zkju := z.fieldsNotEmpty(empty_zqar[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zkju)
	if err != nil {
		return err
	}

	if !empty_zqar[0] {
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

	if !empty_zqar[2] {
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

	if !empty_zqar[3] {
		// write "Snot_zid02_map"
		err = en.Append(0xae, 0x53, 0x6e, 0x6f, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zshl, zoxw := range z.Snot {
			err = en.WriteString(zshl)
			if err != nil {
				return
			}
			err = en.WriteBool(zoxw)
			if err != nil {
				return
			}
		}
	}

	if !empty_zqar[4] {
		// write "Notyet_zid03_map"
		err = en.Append(0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zdjo, zvbh := range z.Notyet {
			err = en.WriteString(zdjo)
			if err != nil {
				return
			}
			err = en.WriteBool(zvbh)
			if err != nil {
				return
			}
		}
	}

	if !empty_zqar[5] {
		// write "Setm_zid04_slc"
		err = en.Append(0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for znnc := range z.Setm {
			if z.Setm[znnc] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[znnc].EncodeMsg(en)
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
		for zshl, zoxw := range z.Snot {
			o = msgp.AppendString(o, zshl)
			o = msgp.AppendBool(o, zoxw)
		}
	}

	if !empty[4] {
		// string "Notyet_zid03_map"
		o = append(o, 0xb0, 0x4e, 0x6f, 0x74, 0x79, 0x65, 0x74, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zdjo, zvbh := range z.Notyet {
			o = msgp.AppendString(o, zdjo)
			o = msgp.AppendBool(o, zvbh)
		}
	}

	if !empty[5] {
		// string "Setm_zid04_slc"
		o = append(o, 0xae, 0x53, 0x65, 0x74, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for znnc := range z.Setm {
			if z.Setm[znnc] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[znnc].MarshalMsg(o)
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
	const maxFields1zflu = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zflu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zflu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zflu := totalEncodedFields1zflu
	missingFieldsLeft1zflu := maxFields1zflu - totalEncodedFields1zflu

	var nextMiss1zflu int32 = -1
	var found1zflu [maxFields1zflu]bool
	var curField1zflu string

doneWithStruct1zflu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zflu > 0 || missingFieldsLeft1zflu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zflu, missingFieldsLeft1zflu, msgp.ShowFound(found1zflu[:]), unmarshalMsgFieldOrder1zflu)
		if encodedFieldsLeft1zflu > 0 {
			encodedFieldsLeft1zflu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zflu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zflu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zflu = 0
			}
			for nextMiss1zflu < maxFields1zflu && (found1zflu[nextMiss1zflu] || unmarshalMsgFieldSkip1zflu[nextMiss1zflu]) {
				nextMiss1zflu++
			}
			if nextMiss1zflu == maxFields1zflu {
				// filled all the empty fields!
				break doneWithStruct1zflu
			}
			missingFieldsLeft1zflu--
			curField1zflu = unmarshalMsgFieldOrder1zflu[nextMiss1zflu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zflu)
		switch curField1zflu {
		// -- templateUnmarshalMsg ends here --

		case "U_zid00_str":
			found1zflu[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Nt_zid01_int":
			found1zflu[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "Snot_zid02_map":
			found1zflu[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zhou uint32
				zhou, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zhou > 0 {
					z.Snot = make(map[string]bool, zhou)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zhou > 0 {
					var zshl string
					var zoxw bool
					zhou--
					zshl, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zoxw, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zshl] = zoxw
				}
			}
		case "Notyet_zid03_map":
			found1zflu[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var znom uint32
				znom, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && znom > 0 {
					z.Notyet = make(map[string]bool, znom)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for znom > 0 {
					var zdjo string
					var zvbh bool
					znom--
					zdjo, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zvbh, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zdjo] = zvbh
				}
			}
		case "Setm_zid04_slc":
			found1zflu[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zbna uint32
				zbna, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zbna) {
					z.Setm = (z.Setm)[:zbna]
				} else {
					z.Setm = make([]*inn, zbna)
				}
				for znnc := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[znnc] != nil {
							z.Setm[znnc].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[znnc] {
								z.Setm[znnc].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[znnc] == nil {
								z.Setm[znnc] = new(inn)
							}
							bts, err = z.Setm[znnc].UnmarshalMsg(bts)
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
	if nextMiss1zflu != -1 {
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
var unmarshalMsgFieldOrder1zflu = []string{"U_zid00_str", "", "Nt_zid01_int", "Snot_zid02_map", "Notyet_zid03_map", "Setm_zid04_slc"}

var unmarshalMsgFieldSkip1zflu = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.U) + 13 + msgp.IntSize + 15 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zshl, zoxw := range z.Snot {
			_ = zoxw
			_ = zshl
			s += msgp.StringPrefixSize + len(zshl) + msgp.BoolSize
		}
	}
	s += 17 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zdjo, zvbh := range z.Notyet {
			_ = zvbh
			_ = zdjo
			s += msgp.StringPrefixSize + len(zdjo) + msgp.BoolSize
		}
	}
	s += 15 + msgp.ArrayHeaderSize
	for znnc := range z.Setm {
		if z.Setm[znnc] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[znnc].Msgsize()
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
	const maxFields2zskg = 2

	// -- templateDecodeMsg starts here--
	var totalEncodedFields2zskg uint32
	totalEncodedFields2zskg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zskg := totalEncodedFields2zskg
	missingFieldsLeft2zskg := maxFields2zskg - totalEncodedFields2zskg

	var nextMiss2zskg int32 = -1
	var found2zskg [maxFields2zskg]bool
	var curField2zskg string

doneWithStruct2zskg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zskg > 0 || missingFieldsLeft2zskg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zskg, missingFieldsLeft2zskg, msgp.ShowFound(found2zskg[:]), decodeMsgFieldOrder2zskg)
		if encodedFieldsLeft2zskg > 0 {
			encodedFieldsLeft2zskg--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField2zskg = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zskg < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zskg = 0
			}
			for nextMiss2zskg < maxFields2zskg && (found2zskg[nextMiss2zskg] || decodeMsgFieldSkip2zskg[nextMiss2zskg]) {
				nextMiss2zskg++
			}
			if nextMiss2zskg == maxFields2zskg {
				// filled all the empty fields!
				break doneWithStruct2zskg
			}
			missingFieldsLeft2zskg--
			curField2zskg = decodeMsgFieldOrder2zskg[nextMiss2zskg]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zskg)
		switch curField2zskg {
		// -- templateDecodeMsg ends here --

		case "j_zid00_map":
			found2zskg[0] = true
			var zdep uint32
			zdep, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zdep > 0 {
				z.j = make(map[string]bool, zdep)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zdep > 0 {
				zdep--
				var zsgh string
				var zypd bool
				zsgh, err = dc.ReadString()
				if err != nil {
					return
				}
				zypd, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zsgh] = zypd
			}
		case "e_zid01_i64":
			found2zskg[1] = true
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
	if nextMiss2zskg != -1 {
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
var decodeMsgFieldOrder2zskg = []string{"j_zid00_map", "e_zid01_i64"}

var decodeMsgFieldSkip2zskg = []bool{false, false}

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
	var empty_zian [2]bool
	fieldsInUse_znmz := z.fieldsNotEmpty(empty_zian[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_znmz)
	if err != nil {
		return err
	}

	if !empty_zian[0] {
		// write "j_zid00_map"
		err = en.Append(0xab, 0x6a, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zsgh, zypd := range z.j {
			err = en.WriteString(zsgh)
			if err != nil {
				return
			}
			err = en.WriteBool(zypd)
			if err != nil {
				return
			}
		}
	}

	if !empty_zian[1] {
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
		for zsgh, zypd := range z.j {
			o = msgp.AppendString(o, zsgh)
			o = msgp.AppendBool(o, zypd)
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
	const maxFields3zgaj = 2

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zgaj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zgaj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zgaj := totalEncodedFields3zgaj
	missingFieldsLeft3zgaj := maxFields3zgaj - totalEncodedFields3zgaj

	var nextMiss3zgaj int32 = -1
	var found3zgaj [maxFields3zgaj]bool
	var curField3zgaj string

doneWithStruct3zgaj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zgaj > 0 || missingFieldsLeft3zgaj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zgaj, missingFieldsLeft3zgaj, msgp.ShowFound(found3zgaj[:]), unmarshalMsgFieldOrder3zgaj)
		if encodedFieldsLeft3zgaj > 0 {
			encodedFieldsLeft3zgaj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zgaj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zgaj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zgaj = 0
			}
			for nextMiss3zgaj < maxFields3zgaj && (found3zgaj[nextMiss3zgaj] || unmarshalMsgFieldSkip3zgaj[nextMiss3zgaj]) {
				nextMiss3zgaj++
			}
			if nextMiss3zgaj == maxFields3zgaj {
				// filled all the empty fields!
				break doneWithStruct3zgaj
			}
			missingFieldsLeft3zgaj--
			curField3zgaj = unmarshalMsgFieldOrder3zgaj[nextMiss3zgaj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zgaj)
		switch curField3zgaj {
		// -- templateUnmarshalMsg ends here --

		case "j_zid00_map":
			found3zgaj[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var zxmc uint32
				zxmc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && zxmc > 0 {
					z.j = make(map[string]bool, zxmc)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for zxmc > 0 {
					var zsgh string
					var zypd bool
					zxmc--
					zsgh, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zypd, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zsgh] = zypd
				}
			}
		case "e_zid01_i64":
			found3zgaj[1] = true
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
	if nextMiss3zgaj != -1 {
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
var unmarshalMsgFieldOrder3zgaj = []string{"j_zid00_map", "e_zid01_i64"}

var unmarshalMsgFieldSkip3zgaj = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.j != nil {
		for zsgh, zypd := range z.j {
			_ = zypd
			_ = zsgh
			s += msgp.StringPrefixSize + len(zsgh) + msgp.BoolSize
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
	const maxFields4zhik = 3

	// -- templateDecodeMsg starts here--
	var totalEncodedFields4zhik uint32
	totalEncodedFields4zhik, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zhik := totalEncodedFields4zhik
	missingFieldsLeft4zhik := maxFields4zhik - totalEncodedFields4zhik

	var nextMiss4zhik int32 = -1
	var found4zhik [maxFields4zhik]bool
	var curField4zhik string

doneWithStruct4zhik:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zhik > 0 || missingFieldsLeft4zhik > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zhik, missingFieldsLeft4zhik, msgp.ShowFound(found4zhik[:]), decodeMsgFieldOrder4zhik)
		if encodedFieldsLeft4zhik > 0 {
			encodedFieldsLeft4zhik--
			field, err = dc.ReadMapKeyPtr()
			if err != nil {
				return
			}
			curField4zhik = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zhik < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zhik = 0
			}
			for nextMiss4zhik < maxFields4zhik && (found4zhik[nextMiss4zhik] || decodeMsgFieldSkip4zhik[nextMiss4zhik]) {
				nextMiss4zhik++
			}
			if nextMiss4zhik == maxFields4zhik {
				// filled all the empty fields!
				break doneWithStruct4zhik
			}
			missingFieldsLeft4zhik--
			curField4zhik = decodeMsgFieldOrder4zhik[nextMiss4zhik]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zhik)
		switch curField4zhik {
		// -- templateDecodeMsg ends here --

		case "m_zid00_map":
			found4zhik[0] = true
			var zter uint32
			zter, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zter > 0 {
				z.m = make(map[string]*Tr, zter)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zter > 0 {
				zter--
				var ztrs string
				var zidt *Tr
				ztrs, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zidt != nil {
						dc.PushAlwaysNil()
						err = zidt.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zidt == nil {
						zidt = new(Tr)
					}
					err = zidt.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[ztrs] = zidt
			}
		case "s_zid01_str":
			found4zhik[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case "n_zid02_i64":
			found4zhik[2] = true
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
	if nextMiss4zhik != -1 {
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
var decodeMsgFieldOrder4zhik = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var decodeMsgFieldSkip4zhik = []bool{false, false, false}

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
	var empty_zfay [3]bool
	fieldsInUse_zgjg := z.fieldsNotEmpty(empty_zfay[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgjg)
	if err != nil {
		return err
	}

	if !empty_zfay[0] {
		// write "m_zid00_map"
		err = en.Append(0xab, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for ztrs, zidt := range z.m {
			err = en.WriteString(ztrs)
			if err != nil {
				return
			}
			if zidt == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zidt.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zfay[1] {
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

	if !empty_zfay[2] {
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
		for ztrs, zidt := range z.m {
			o = msgp.AppendString(o, ztrs)
			if zidt == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zidt.MarshalMsg(o)
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
	const maxFields5zgke = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields5zgke uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zgke, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zgke := totalEncodedFields5zgke
	missingFieldsLeft5zgke := maxFields5zgke - totalEncodedFields5zgke

	var nextMiss5zgke int32 = -1
	var found5zgke [maxFields5zgke]bool
	var curField5zgke string

doneWithStruct5zgke:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zgke > 0 || missingFieldsLeft5zgke > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zgke, missingFieldsLeft5zgke, msgp.ShowFound(found5zgke[:]), unmarshalMsgFieldOrder5zgke)
		if encodedFieldsLeft5zgke > 0 {
			encodedFieldsLeft5zgke--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField5zgke = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss5zgke < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zgke = 0
			}
			for nextMiss5zgke < maxFields5zgke && (found5zgke[nextMiss5zgke] || unmarshalMsgFieldSkip5zgke[nextMiss5zgke]) {
				nextMiss5zgke++
			}
			if nextMiss5zgke == maxFields5zgke {
				// filled all the empty fields!
				break doneWithStruct5zgke
			}
			missingFieldsLeft5zgke--
			curField5zgke = unmarshalMsgFieldOrder5zgke[nextMiss5zgke]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zgke)
		switch curField5zgke {
		// -- templateUnmarshalMsg ends here --

		case "m_zid00_map":
			found5zgke[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zwhd uint32
				zwhd, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zwhd > 0 {
					z.m = make(map[string]*Tr, zwhd)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zwhd > 0 {
					var ztrs string
					var zidt *Tr
					zwhd--
					ztrs, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zidt != nil {
							zidt.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zidt {
								zidt.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zidt == nil {
								zidt = new(Tr)
							}
							bts, err = zidt.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[ztrs] = zidt
				}
			}
		case "s_zid01_str":
			found5zgke[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "n_zid02_i64":
			found5zgke[2] = true
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
	if nextMiss5zgke != -1 {
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
var unmarshalMsgFieldOrder5zgke = []string{"m_zid00_map", "s_zid01_str", "n_zid02_i64"}

var unmarshalMsgFieldSkip5zgke = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.m != nil {
		for ztrs, zidt := range z.m {
			_ = zidt
			_ = ztrs
			s += msgp.StringPrefixSize + len(ztrs)
			if zidt == nil {
				s += msgp.NilSize
			} else {
				s += zidt.Msgsize()
			}
		}
	}
	s += 12 + msgp.StringPrefixSize + len(z.s) + 12 + msgp.Int64Size
	return
}
