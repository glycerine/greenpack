package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// MSGPfieldsNotEmpty supports omitempty tags
func (z *A) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[0] = (len(z.Name) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.Bday.IsZero()) // time.Time, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Sibs == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.GPA == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (!z.Friend) // bool, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *A) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "name_zid00_str"
		o = append(o, 0xae, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Name)
	}

	if !empty[1] {
		// string "Bday_zid01_tim"
		o = append(o, 0xae, 0x42, 0x64, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
		o = msgp.AppendTime(o, z.Bday)
	}

	if !empty[2] {
		// string "phone_zid02_str"
		o = append(o, 0xaf, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// string "Sibs_zid03_int"
		o = append(o, 0xae, 0x53, 0x69, 0x62, 0x73, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x69, 0x6e, 0x74)
		o = msgp.AppendInt(o, z.Sibs)
	}

	if !empty[4] {
		// string "GPA_zid04_f64"
		o = append(o, 0xad, 0x47, 0x50, 0x41, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x66, 0x36, 0x34)
		o = msgp.AppendFloat64(o, z.GPA)
	}

	if !empty[5] {
		// string "Friend_zid05_boo"
		o = append(o, 0xb0, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x62, 0x6f, 0x6f)
		o = msgp.AppendBool(o, z.Friend)
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *A) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *A) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields0zxcc1 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zxcc1 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zxcc1, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zxcc1 := totalEncodedFields0zxcc1
	missingFieldsLeft0zxcc1 := maxFields0zxcc1 - totalEncodedFields0zxcc1

	var nextMiss0zxcc1 int32 = -1
	var found0zxcc1 [maxFields0zxcc1]bool
	var curField0zxcc1 string

doneWithStruct0zxcc1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zxcc1 > 0 || missingFieldsLeft0zxcc1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zxcc1, missingFieldsLeft0zxcc1, msgp.ShowFound(found0zxcc1[:]), unmarshalMsgFieldOrder0zxcc1)
		if encodedFieldsLeft0zxcc1 > 0 {
			encodedFieldsLeft0zxcc1--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zxcc1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zxcc1 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zxcc1 = 0
			}
			for nextMiss0zxcc1 < maxFields0zxcc1 && (found0zxcc1[nextMiss0zxcc1] || unmarshalMsgFieldSkip0zxcc1[nextMiss0zxcc1]) {
				nextMiss0zxcc1++
			}
			if nextMiss0zxcc1 == maxFields0zxcc1 {
				// filled all the empty fields!
				break doneWithStruct0zxcc1
			}
			missingFieldsLeft0zxcc1--
			curField0zxcc1 = unmarshalMsgFieldOrder0zxcc1[nextMiss0zxcc1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zxcc1)
		switch curField0zxcc1 {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0zxcc1[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0zxcc1[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0zxcc1[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0zxcc1[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0zxcc1[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0zxcc1[5] = true
			z.Friend, bts, err = nbs.ReadBoolBytes(bts)

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
	if nextMiss0zxcc1 != -1 {
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

// fields of A
var unmarshalMsgFieldOrder0zxcc1 = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0zxcc1 = []bool{false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.Name) + 15 + msgp.TimeSize + 16 + msgp.StringPrefixSize + len(z.Phone) + 15 + msgp.IntSize + 14 + msgp.Float64Size + 17 + msgp.BoolSize
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Big) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.Slice) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Transform) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Myptr == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.Myarray) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.MySlice) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Big) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [5]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "Slice_zid00_slc"
		o = append(o, 0xaf, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
		for zxon2 := range z.Slice {
			o, err = z.Slice[zxon2].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Transform_zid01_map"
		o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
		for zlkq3, zryq4 := range z.Transform {
			o = msgp.AppendInt(o, zlkq3)
			if zryq4 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zryq4.MSGPMarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[2] {
		// string "Myptr_zid02_ptr"
		o = append(o, 0xaf, 0x4d, 0x79, 0x70, 0x74, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x70, 0x74, 0x72)
		if z.Myptr == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Myptr.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[3] {
		// string "Myarray_zid03_ary"
		o = append(o, 0xb1, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x61, 0x72, 0x79)
		o = msgp.AppendArrayHeader(o, 3)
		for zdwc5 := range z.Myarray {
			o = msgp.AppendString(o, z.Myarray[zdwc5])
		}
	}

	if !empty[4] {
		// string "MySlice_zid04_slc"
		o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
		for zshc6 := range z.MySlice {
			o = msgp.AppendString(o, z.MySlice[zshc6])
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Big) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Big) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zdtk7 = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zdtk7 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zdtk7, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zdtk7 := totalEncodedFields1zdtk7
	missingFieldsLeft1zdtk7 := maxFields1zdtk7 - totalEncodedFields1zdtk7

	var nextMiss1zdtk7 int32 = -1
	var found1zdtk7 [maxFields1zdtk7]bool
	var curField1zdtk7 string

doneWithStruct1zdtk7:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zdtk7 > 0 || missingFieldsLeft1zdtk7 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zdtk7, missingFieldsLeft1zdtk7, msgp.ShowFound(found1zdtk7[:]), unmarshalMsgFieldOrder1zdtk7)
		if encodedFieldsLeft1zdtk7 > 0 {
			encodedFieldsLeft1zdtk7--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zdtk7 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zdtk7 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zdtk7 = 0
			}
			for nextMiss1zdtk7 < maxFields1zdtk7 && (found1zdtk7[nextMiss1zdtk7] || unmarshalMsgFieldSkip1zdtk7[nextMiss1zdtk7]) {
				nextMiss1zdtk7++
			}
			if nextMiss1zdtk7 == maxFields1zdtk7 {
				// filled all the empty fields!
				break doneWithStruct1zdtk7
			}
			missingFieldsLeft1zdtk7--
			curField1zdtk7 = unmarshalMsgFieldOrder1zdtk7[nextMiss1zdtk7]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zdtk7)
		switch curField1zdtk7 {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found1zdtk7[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zbgl8 uint32
				zbgl8, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zbgl8) {
					z.Slice = (z.Slice)[:zbgl8]
				} else {
					z.Slice = make([]S2, zbgl8)
				}
				for zxon2 := range z.Slice {
					bts, err = z.Slice[zxon2].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found1zdtk7[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zgeb9 uint32
				zgeb9, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zgeb9 > 0 {
					z.Transform = make(map[int]*S2, zgeb9)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zgeb9 > 0 {
					var zlkq3 int
					var zryq4 *S2
					zgeb9--
					zlkq3, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zryq4 != nil {
							zryq4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zryq4 {
								zryq4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zryq4 == nil {
								zryq4 = new(S2)
							}
							bts, err = zryq4.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zlkq3] = zryq4
				}
			}
		case "Myptr_zid02_ptr":
			found1zdtk7[2] = true
			if nbs.AlwaysNil {
				if z.Myptr != nil {
					z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Myptr {
						z.Myptr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Myptr == nil {
						z.Myptr = new(S2)
					}
					bts, err = z.Myptr.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Myarray_zid03_ary":
			found1zdtk7[3] = true
			var zoqq10 uint32
			zoqq10, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zoqq10 != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zoqq10}
				return
			}
			for zdwc5 := range z.Myarray {
				z.Myarray[zdwc5], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found1zdtk7[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zqqw11 uint32
				zqqw11, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zqqw11) {
					z.MySlice = (z.MySlice)[:zqqw11]
				} else {
					z.MySlice = make([]string, zqqw11)
				}
				for zshc6 := range z.MySlice {
					z.MySlice[zshc6], bts, err = nbs.ReadStringBytes(bts)

					if err != nil {
						return
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
	if nextMiss1zdtk7 != -1 {
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

// fields of Big
var unmarshalMsgFieldOrder1zdtk7 = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip1zdtk7 = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zxon2 := range z.Slice {
		s += z.Slice[zxon2].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zlkq3, zryq4 := range z.Transform {
			_ = zryq4
			_ = zlkq3
			s += msgp.IntSize
			if zryq4 == nil {
				s += msgp.NilSize
			} else {
				s += zryq4.MSGPMsgsize()
			}
		}
	}
	s += 16
	if z.Myptr == nil {
		s += msgp.NilSize
	} else {
		s += z.Myptr.MSGPMsgsize()
	}
	s += 18 + msgp.ArrayHeaderSize
	for zdwc5 := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zdwc5])
	}
	s += 18 + msgp.ArrayHeaderSize
	for zshc6 := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zshc6])
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *S2) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 7
	}
	var fieldsInUse uint32 = 7
	isempty[1] = (len(z.B) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.R) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.P == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.Q == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.T == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}
	isempty[6] = (len(z.Arr) == 0) // string, omitempty
	if isempty[6] {
		fieldsInUse--
	}
	isempty[7] = (z.MyTree == nil) // pointer, omitempty
	if isempty[7] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *S2) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [8]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[1] {
		// string "beta_zid01_str"
		o = append(o, 0xae, 0x62, 0x65, 0x74, 0x61, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.B)
	}

	if !empty[2] {
		// string "ralph_zid02_map"
		o = append(o, 0xaf, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.R)))
		for zobg12, zgzj13 := range z.R {
			o = msgp.AppendString(o, zobg12)
			o = msgp.AppendUint8(o, zgzj13)
		}
	}

	if !empty[3] {
		// string "P_zid03_u16"
		o = append(o, 0xab, 0x50, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x75, 0x31, 0x36)
		o = msgp.AppendUint16(o, z.P)
	}

	if !empty[4] {
		// string "Q_zid04_u32"
		o = append(o, 0xab, 0x51, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x75, 0x33, 0x32)
		o = msgp.AppendUint32(o, z.Q)
	}

	if !empty[5] {
		// string "T_zid05_i64"
		o = append(o, 0xab, 0x54, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x69, 0x36, 0x34)
		o = msgp.AppendInt64(o, z.T)
	}

	if !empty[6] {
		// string "arr_zid06_ary"
		o = append(o, 0xad, 0x61, 0x72, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x36, 0x5f, 0x61, 0x72, 0x79)
		o = msgp.AppendArrayHeader(o, 6)
		for zmyr14 := range z.Arr {
			o = msgp.AppendFloat64(o, z.Arr[zmyr14])
		}
	}

	if !empty[7] {
		// string "MyTree_zid07_ptr"
		o = append(o, 0xb0, 0x4d, 0x79, 0x54, 0x72, 0x65, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x37, 0x5f, 0x70, 0x74, 0x72)
		if z.MyTree == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.MyTree.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *S2) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *S2) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields2zhgz15 = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zhgz15 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zhgz15, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zhgz15 := totalEncodedFields2zhgz15
	missingFieldsLeft2zhgz15 := maxFields2zhgz15 - totalEncodedFields2zhgz15

	var nextMiss2zhgz15 int32 = -1
	var found2zhgz15 [maxFields2zhgz15]bool
	var curField2zhgz15 string

doneWithStruct2zhgz15:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zhgz15 > 0 || missingFieldsLeft2zhgz15 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zhgz15, missingFieldsLeft2zhgz15, msgp.ShowFound(found2zhgz15[:]), unmarshalMsgFieldOrder2zhgz15)
		if encodedFieldsLeft2zhgz15 > 0 {
			encodedFieldsLeft2zhgz15--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zhgz15 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zhgz15 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zhgz15 = 0
			}
			for nextMiss2zhgz15 < maxFields2zhgz15 && (found2zhgz15[nextMiss2zhgz15] || unmarshalMsgFieldSkip2zhgz15[nextMiss2zhgz15]) {
				nextMiss2zhgz15++
			}
			if nextMiss2zhgz15 == maxFields2zhgz15 {
				// filled all the empty fields!
				break doneWithStruct2zhgz15
			}
			missingFieldsLeft2zhgz15--
			curField2zhgz15 = unmarshalMsgFieldOrder2zhgz15[nextMiss2zhgz15]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zhgz15)
		switch curField2zhgz15 {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found2zhgz15[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found2zhgz15[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zfeq16 uint32
				zfeq16, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zfeq16 > 0 {
					z.R = make(map[string]uint8, zfeq16)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zfeq16 > 0 {
					var zobg12 string
					var zgzj13 uint8
					zfeq16--
					zobg12, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zgzj13, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zobg12] = zgzj13
				}
			}
		case "P_zid03_u16":
			found2zhgz15[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found2zhgz15[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found2zhgz15[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found2zhgz15[6] = true
			var zpfz17 uint32
			zpfz17, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zpfz17 != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zpfz17}
				return
			}
			for zmyr14 := range z.Arr {
				z.Arr[zmyr14], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found2zhgz15[7] = true
			if nbs.AlwaysNil {
				if z.MyTree != nil {
					z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.MyTree {
						z.MyTree.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.MyTree == nil {
						z.MyTree = new(Tree)
					}
					bts, err = z.MyTree.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
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
	if nextMiss2zhgz15 != -1 {
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

// fields of S2
var unmarshalMsgFieldOrder2zhgz15 = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip2zhgz15 = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zobg12, zgzj13 := range z.R {
			_ = zgzj13
			_ = zobg12
			s += msgp.StringPrefixSize + len(zobg12) + msgp.Uint8Size
		}
	}
	s += 12 + msgp.Uint16Size + 12 + msgp.Uint32Size + 12 + msgp.Int64Size + 14 + msgp.ArrayHeaderSize + (6 * (msgp.Float64Size)) + 17
	if z.MyTree == nil {
		s += msgp.NilSize
	} else {
		s += z.MyTree.MSGPMsgsize()
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z Sys) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 1
	}
	var fieldsInUse uint32 = 1
	isempty[0] = false
	if isempty[0] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z Sys) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [1]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "F_zid00_ifc"
		o = append(o, 0xab, 0x46, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x69, 0x66, 0x63)
		o, err = msgp.AppendIntf(o, z.F)
		if err != nil {
			return
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Sys) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Sys) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zgxf18 = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zgxf18 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zgxf18, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zgxf18 := totalEncodedFields3zgxf18
	missingFieldsLeft3zgxf18 := maxFields3zgxf18 - totalEncodedFields3zgxf18

	var nextMiss3zgxf18 int32 = -1
	var found3zgxf18 [maxFields3zgxf18]bool
	var curField3zgxf18 string

doneWithStruct3zgxf18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zgxf18 > 0 || missingFieldsLeft3zgxf18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zgxf18, missingFieldsLeft3zgxf18, msgp.ShowFound(found3zgxf18[:]), unmarshalMsgFieldOrder3zgxf18)
		if encodedFieldsLeft3zgxf18 > 0 {
			encodedFieldsLeft3zgxf18--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zgxf18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zgxf18 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zgxf18 = 0
			}
			for nextMiss3zgxf18 < maxFields3zgxf18 && (found3zgxf18[nextMiss3zgxf18] || unmarshalMsgFieldSkip3zgxf18[nextMiss3zgxf18]) {
				nextMiss3zgxf18++
			}
			if nextMiss3zgxf18 == maxFields3zgxf18 {
				// filled all the empty fields!
				break doneWithStruct3zgxf18
			}
			missingFieldsLeft3zgxf18--
			curField3zgxf18 = unmarshalMsgFieldOrder3zgxf18[nextMiss3zgxf18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zgxf18)
		switch curField3zgxf18 {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found3zgxf18[0] = true
			z.F, bts, err = nbs.ReadIntfBytes(bts)

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
	if nextMiss3zgxf18 != -1 {
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

// fields of Sys
var unmarshalMsgFieldOrder3zgxf18 = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip3zgxf18 = []bool{false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) MSGPMsgsize() (s int) {
	s = 1 + 12 + msgp.GuessSize(z.F)
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Tree) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 3
	}
	var fieldsInUse uint32 = 3
	isempty[0] = (len(z.Chld) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Str) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.Par == nil) // pointer, omitempty
	if isempty[2] {
		fieldsInUse--
	}

	return fieldsInUse
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Tree) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())

	// honor the omitempty tags
	var empty [3]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse)

	if !empty[0] {
		// string "Chld_zid00_slc"
		o = append(o, 0xae, 0x43, 0x68, 0x6c, 0x64, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
		for zcat19 := range z.Chld {
			o, err = z.Chld[zcat19].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Str_zid01_str"
		o = append(o, 0xad, 0x53, 0x74, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
		o = msgp.AppendString(o, z.Str)
	}

	if !empty[2] {
		// string "Par_zid02_ptr"
		o = append(o, 0xad, 0x50, 0x61, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x70, 0x74, 0x72)
		if z.Par == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Par.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	return
}

// MSGPUnmarshalMsg implements msgp.Unmarshaler
func (z *Tree) MSGPUnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.MSGPUnmarshalMsgWithCfg(bts, nil)
}
func (z *Tree) MSGPUnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields4zzgd20 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zzgd20 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zzgd20, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zzgd20 := totalEncodedFields4zzgd20
	missingFieldsLeft4zzgd20 := maxFields4zzgd20 - totalEncodedFields4zzgd20

	var nextMiss4zzgd20 int32 = -1
	var found4zzgd20 [maxFields4zzgd20]bool
	var curField4zzgd20 string

doneWithStruct4zzgd20:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zzgd20 > 0 || missingFieldsLeft4zzgd20 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zzgd20, missingFieldsLeft4zzgd20, msgp.ShowFound(found4zzgd20[:]), unmarshalMsgFieldOrder4zzgd20)
		if encodedFieldsLeft4zzgd20 > 0 {
			encodedFieldsLeft4zzgd20--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zzgd20 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zzgd20 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zzgd20 = 0
			}
			for nextMiss4zzgd20 < maxFields4zzgd20 && (found4zzgd20[nextMiss4zzgd20] || unmarshalMsgFieldSkip4zzgd20[nextMiss4zzgd20]) {
				nextMiss4zzgd20++
			}
			if nextMiss4zzgd20 == maxFields4zzgd20 {
				// filled all the empty fields!
				break doneWithStruct4zzgd20
			}
			missingFieldsLeft4zzgd20--
			curField4zzgd20 = unmarshalMsgFieldOrder4zzgd20[nextMiss4zzgd20]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zzgd20)
		switch curField4zzgd20 {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found4zzgd20[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zbrn21 uint32
				zbrn21, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zbrn21) {
					z.Chld = (z.Chld)[:zbrn21]
				} else {
					z.Chld = make([]Tree, zbrn21)
				}
				for zcat19 := range z.Chld {
					bts, err = z.Chld[zcat19].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found4zzgd20[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found4zzgd20[2] = true
			if nbs.AlwaysNil {
				if z.Par != nil {
					z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
				}
			} else {
				// not nbs.AlwaysNil
				if msgp.IsNil(bts) {
					bts = bts[1:]
					if nil != z.Par {
						z.Par.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
					}
				} else {
					// not nbs.AlwaysNil and not IsNil(bts): have something to read

					if z.Par == nil {
						z.Par = new(S2)
					}
					bts, err = z.Par.MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
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
	if nextMiss4zzgd20 != -1 {
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

// fields of Tree
var unmarshalMsgFieldOrder4zzgd20 = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip4zzgd20 = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zcat19 := range z.Chld {
		s += z.Chld[zcat19].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
