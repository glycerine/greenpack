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
	const maxFields0zujb1 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zujb1 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zujb1, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zujb1 := totalEncodedFields0zujb1
	missingFieldsLeft0zujb1 := maxFields0zujb1 - totalEncodedFields0zujb1

	var nextMiss0zujb1 int32 = -1
	var found0zujb1 [maxFields0zujb1]bool
	var curField0zujb1 string

doneWithStruct0zujb1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zujb1 > 0 || missingFieldsLeft0zujb1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zujb1, missingFieldsLeft0zujb1, msgp.ShowFound(found0zujb1[:]), unmarshalMsgFieldOrder0zujb1)
		if encodedFieldsLeft0zujb1 > 0 {
			encodedFieldsLeft0zujb1--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zujb1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zujb1 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zujb1 = 0
			}
			for nextMiss0zujb1 < maxFields0zujb1 && (found0zujb1[nextMiss0zujb1] || unmarshalMsgFieldSkip0zujb1[nextMiss0zujb1]) {
				nextMiss0zujb1++
			}
			if nextMiss0zujb1 == maxFields0zujb1 {
				// filled all the empty fields!
				break doneWithStruct0zujb1
			}
			missingFieldsLeft0zujb1--
			curField0zujb1 = unmarshalMsgFieldOrder0zujb1[nextMiss0zujb1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zujb1)
		switch curField0zujb1 {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0zujb1[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0zujb1[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0zujb1[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0zujb1[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0zujb1[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0zujb1[5] = true
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
	if nextMiss0zujb1 != -1 {
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
var unmarshalMsgFieldOrder0zujb1 = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0zujb1 = []bool{false, false, false, false, false, false}

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
		for zexu2 := range z.Slice {
			o, err = z.Slice[zexu2].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Transform_zid01_map"
		o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
		for zqxv3, ztjt4 := range z.Transform {
			o = msgp.AppendInt(o, zqxv3)
			if ztjt4 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = ztjt4.MSGPMarshalMsg(o)
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
		for zyls5 := range z.Myarray {
			o = msgp.AppendString(o, z.Myarray[zyls5])
		}
	}

	if !empty[4] {
		// string "MySlice_zid04_slc"
		o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
		for zlra6 := range z.MySlice {
			o = msgp.AppendString(o, z.MySlice[zlra6])
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
	const maxFields1zdpc7 = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zdpc7 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zdpc7, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zdpc7 := totalEncodedFields1zdpc7
	missingFieldsLeft1zdpc7 := maxFields1zdpc7 - totalEncodedFields1zdpc7

	var nextMiss1zdpc7 int32 = -1
	var found1zdpc7 [maxFields1zdpc7]bool
	var curField1zdpc7 string

doneWithStruct1zdpc7:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zdpc7 > 0 || missingFieldsLeft1zdpc7 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zdpc7, missingFieldsLeft1zdpc7, msgp.ShowFound(found1zdpc7[:]), unmarshalMsgFieldOrder1zdpc7)
		if encodedFieldsLeft1zdpc7 > 0 {
			encodedFieldsLeft1zdpc7--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zdpc7 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zdpc7 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zdpc7 = 0
			}
			for nextMiss1zdpc7 < maxFields1zdpc7 && (found1zdpc7[nextMiss1zdpc7] || unmarshalMsgFieldSkip1zdpc7[nextMiss1zdpc7]) {
				nextMiss1zdpc7++
			}
			if nextMiss1zdpc7 == maxFields1zdpc7 {
				// filled all the empty fields!
				break doneWithStruct1zdpc7
			}
			missingFieldsLeft1zdpc7--
			curField1zdpc7 = unmarshalMsgFieldOrder1zdpc7[nextMiss1zdpc7]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zdpc7)
		switch curField1zdpc7 {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found1zdpc7[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var znbg8 uint32
				znbg8, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(znbg8) {
					z.Slice = (z.Slice)[:znbg8]
				} else {
					z.Slice = make([]S2, znbg8)
				}
				for zexu2 := range z.Slice {
					bts, err = z.Slice[zexu2].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found1zdpc7[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var znbu9 uint32
				znbu9, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && znbu9 > 0 {
					z.Transform = make(map[int]*S2, znbu9)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for znbu9 > 0 {
					var zqxv3 int
					var ztjt4 *S2
					znbu9--
					zqxv3, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if ztjt4 != nil {
							ztjt4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != ztjt4 {
								ztjt4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if ztjt4 == nil {
								ztjt4 = new(S2)
							}
							bts, err = ztjt4.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zqxv3] = ztjt4
				}
			}
		case "Myptr_zid02_ptr":
			found1zdpc7[2] = true
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
			found1zdpc7[3] = true
			var zivc10 uint32
			zivc10, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zivc10 != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zivc10}
				return
			}
			for zyls5 := range z.Myarray {
				z.Myarray[zyls5], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found1zdpc7[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zkzj11 uint32
				zkzj11, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zkzj11) {
					z.MySlice = (z.MySlice)[:zkzj11]
				} else {
					z.MySlice = make([]string, zkzj11)
				}
				for zlra6 := range z.MySlice {
					z.MySlice[zlra6], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zdpc7 != -1 {
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
var unmarshalMsgFieldOrder1zdpc7 = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip1zdpc7 = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zexu2 := range z.Slice {
		s += z.Slice[zexu2].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zqxv3, ztjt4 := range z.Transform {
			_ = ztjt4
			_ = zqxv3
			s += msgp.IntSize
			if ztjt4 == nil {
				s += msgp.NilSize
			} else {
				s += ztjt4.MSGPMsgsize()
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
	for zyls5 := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zyls5])
	}
	s += 18 + msgp.ArrayHeaderSize
	for zlra6 := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zlra6])
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
		for zwyq12, zzbm13 := range z.R {
			o = msgp.AppendString(o, zwyq12)
			o = msgp.AppendUint8(o, zzbm13)
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
		for zwhh14 := range z.Arr {
			o = msgp.AppendFloat64(o, z.Arr[zwhh14])
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
	const maxFields2zgjk15 = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zgjk15 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zgjk15, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zgjk15 := totalEncodedFields2zgjk15
	missingFieldsLeft2zgjk15 := maxFields2zgjk15 - totalEncodedFields2zgjk15

	var nextMiss2zgjk15 int32 = -1
	var found2zgjk15 [maxFields2zgjk15]bool
	var curField2zgjk15 string

doneWithStruct2zgjk15:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zgjk15 > 0 || missingFieldsLeft2zgjk15 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zgjk15, missingFieldsLeft2zgjk15, msgp.ShowFound(found2zgjk15[:]), unmarshalMsgFieldOrder2zgjk15)
		if encodedFieldsLeft2zgjk15 > 0 {
			encodedFieldsLeft2zgjk15--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zgjk15 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zgjk15 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zgjk15 = 0
			}
			for nextMiss2zgjk15 < maxFields2zgjk15 && (found2zgjk15[nextMiss2zgjk15] || unmarshalMsgFieldSkip2zgjk15[nextMiss2zgjk15]) {
				nextMiss2zgjk15++
			}
			if nextMiss2zgjk15 == maxFields2zgjk15 {
				// filled all the empty fields!
				break doneWithStruct2zgjk15
			}
			missingFieldsLeft2zgjk15--
			curField2zgjk15 = unmarshalMsgFieldOrder2zgjk15[nextMiss2zgjk15]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zgjk15)
		switch curField2zgjk15 {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found2zgjk15[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found2zgjk15[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zblj16 uint32
				zblj16, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zblj16 > 0 {
					z.R = make(map[string]uint8, zblj16)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zblj16 > 0 {
					var zwyq12 string
					var zzbm13 uint8
					zblj16--
					zwyq12, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zzbm13, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zwyq12] = zzbm13
				}
			}
		case "P_zid03_u16":
			found2zgjk15[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found2zgjk15[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found2zgjk15[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found2zgjk15[6] = true
			var zbaf17 uint32
			zbaf17, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zbaf17 != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zbaf17}
				return
			}
			for zwhh14 := range z.Arr {
				z.Arr[zwhh14], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found2zgjk15[7] = true
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
	if nextMiss2zgjk15 != -1 {
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
var unmarshalMsgFieldOrder2zgjk15 = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip2zgjk15 = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zwyq12, zzbm13 := range z.R {
			_ = zzbm13
			_ = zwyq12
			s += msgp.StringPrefixSize + len(zwyq12) + msgp.Uint8Size
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
	const maxFields3zfny18 = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zfny18 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zfny18, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zfny18 := totalEncodedFields3zfny18
	missingFieldsLeft3zfny18 := maxFields3zfny18 - totalEncodedFields3zfny18

	var nextMiss3zfny18 int32 = -1
	var found3zfny18 [maxFields3zfny18]bool
	var curField3zfny18 string

doneWithStruct3zfny18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zfny18 > 0 || missingFieldsLeft3zfny18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zfny18, missingFieldsLeft3zfny18, msgp.ShowFound(found3zfny18[:]), unmarshalMsgFieldOrder3zfny18)
		if encodedFieldsLeft3zfny18 > 0 {
			encodedFieldsLeft3zfny18--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zfny18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zfny18 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zfny18 = 0
			}
			for nextMiss3zfny18 < maxFields3zfny18 && (found3zfny18[nextMiss3zfny18] || unmarshalMsgFieldSkip3zfny18[nextMiss3zfny18]) {
				nextMiss3zfny18++
			}
			if nextMiss3zfny18 == maxFields3zfny18 {
				// filled all the empty fields!
				break doneWithStruct3zfny18
			}
			missingFieldsLeft3zfny18--
			curField3zfny18 = unmarshalMsgFieldOrder3zfny18[nextMiss3zfny18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zfny18)
		switch curField3zfny18 {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found3zfny18[0] = true
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
	if nextMiss3zfny18 != -1 {
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
var unmarshalMsgFieldOrder3zfny18 = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip3zfny18 = []bool{false}

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
		for zabl19 := range z.Chld {
			o, err = z.Chld[zabl19].MSGPMarshalMsg(o)
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
	const maxFields4zlkm20 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zlkm20 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zlkm20, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zlkm20 := totalEncodedFields4zlkm20
	missingFieldsLeft4zlkm20 := maxFields4zlkm20 - totalEncodedFields4zlkm20

	var nextMiss4zlkm20 int32 = -1
	var found4zlkm20 [maxFields4zlkm20]bool
	var curField4zlkm20 string

doneWithStruct4zlkm20:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zlkm20 > 0 || missingFieldsLeft4zlkm20 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zlkm20, missingFieldsLeft4zlkm20, msgp.ShowFound(found4zlkm20[:]), unmarshalMsgFieldOrder4zlkm20)
		if encodedFieldsLeft4zlkm20 > 0 {
			encodedFieldsLeft4zlkm20--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zlkm20 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zlkm20 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zlkm20 = 0
			}
			for nextMiss4zlkm20 < maxFields4zlkm20 && (found4zlkm20[nextMiss4zlkm20] || unmarshalMsgFieldSkip4zlkm20[nextMiss4zlkm20]) {
				nextMiss4zlkm20++
			}
			if nextMiss4zlkm20 == maxFields4zlkm20 {
				// filled all the empty fields!
				break doneWithStruct4zlkm20
			}
			missingFieldsLeft4zlkm20--
			curField4zlkm20 = unmarshalMsgFieldOrder4zlkm20[nextMiss4zlkm20]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zlkm20)
		switch curField4zlkm20 {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found4zlkm20[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zbby21 uint32
				zbby21, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zbby21) {
					z.Chld = (z.Chld)[:zbby21]
				} else {
					z.Chld = make([]Tree, zbby21)
				}
				for zabl19 := range z.Chld {
					bts, err = z.Chld[zabl19].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found4zlkm20[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found4zlkm20[2] = true
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
	if nextMiss4zlkm20 != -1 {
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
var unmarshalMsgFieldOrder4zlkm20 = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip4zlkm20 = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zabl19 := range z.Chld {
		s += z.Chld[zabl19].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
