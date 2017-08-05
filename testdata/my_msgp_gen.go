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
	const maxFields0zhfg1 = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zhfg1 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zhfg1, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zhfg1 := totalEncodedFields0zhfg1
	missingFieldsLeft0zhfg1 := maxFields0zhfg1 - totalEncodedFields0zhfg1

	var nextMiss0zhfg1 int32 = -1
	var found0zhfg1 [maxFields0zhfg1]bool
	var curField0zhfg1 string

doneWithStruct0zhfg1:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zhfg1 > 0 || missingFieldsLeft0zhfg1 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zhfg1, missingFieldsLeft0zhfg1, msgp.ShowFound(found0zhfg1[:]), unmarshalMsgFieldOrder0zhfg1)
		if encodedFieldsLeft0zhfg1 > 0 {
			encodedFieldsLeft0zhfg1--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zhfg1 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zhfg1 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zhfg1 = 0
			}
			for nextMiss0zhfg1 < maxFields0zhfg1 && (found0zhfg1[nextMiss0zhfg1] || unmarshalMsgFieldSkip0zhfg1[nextMiss0zhfg1]) {
				nextMiss0zhfg1++
			}
			if nextMiss0zhfg1 == maxFields0zhfg1 {
				// filled all the empty fields!
				break doneWithStruct0zhfg1
			}
			missingFieldsLeft0zhfg1--
			curField0zhfg1 = unmarshalMsgFieldOrder0zhfg1[nextMiss0zhfg1]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zhfg1)
		switch curField0zhfg1 {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0zhfg1[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0zhfg1[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0zhfg1[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0zhfg1[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0zhfg1[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0zhfg1[5] = true
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
	if nextMiss0zhfg1 != -1 {
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
var unmarshalMsgFieldOrder0zhfg1 = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0zhfg1 = []bool{false, false, false, false, false, false}

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
		for zznr2 := range z.Slice {
			o, err = z.Slice[zznr2].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Transform_zid01_map"
		o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
		for ztmx3, zamq4 := range z.Transform {
			o = msgp.AppendInt(o, ztmx3)
			if zamq4 == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zamq4.MSGPMarshalMsg(o)
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
		for zbjo5 := range z.Myarray {
			o = msgp.AppendString(o, z.Myarray[zbjo5])
		}
	}

	if !empty[4] {
		// string "MySlice_zid04_slc"
		o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
		for ztdw6 := range z.MySlice {
			o = msgp.AppendString(o, z.MySlice[ztdw6])
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
	const maxFields1zuzv7 = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zuzv7 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zuzv7, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zuzv7 := totalEncodedFields1zuzv7
	missingFieldsLeft1zuzv7 := maxFields1zuzv7 - totalEncodedFields1zuzv7

	var nextMiss1zuzv7 int32 = -1
	var found1zuzv7 [maxFields1zuzv7]bool
	var curField1zuzv7 string

doneWithStruct1zuzv7:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zuzv7 > 0 || missingFieldsLeft1zuzv7 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zuzv7, missingFieldsLeft1zuzv7, msgp.ShowFound(found1zuzv7[:]), unmarshalMsgFieldOrder1zuzv7)
		if encodedFieldsLeft1zuzv7 > 0 {
			encodedFieldsLeft1zuzv7--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zuzv7 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zuzv7 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zuzv7 = 0
			}
			for nextMiss1zuzv7 < maxFields1zuzv7 && (found1zuzv7[nextMiss1zuzv7] || unmarshalMsgFieldSkip1zuzv7[nextMiss1zuzv7]) {
				nextMiss1zuzv7++
			}
			if nextMiss1zuzv7 == maxFields1zuzv7 {
				// filled all the empty fields!
				break doneWithStruct1zuzv7
			}
			missingFieldsLeft1zuzv7--
			curField1zuzv7 = unmarshalMsgFieldOrder1zuzv7[nextMiss1zuzv7]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zuzv7)
		switch curField1zuzv7 {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found1zuzv7[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zjuy8 uint32
				zjuy8, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zjuy8) {
					z.Slice = (z.Slice)[:zjuy8]
				} else {
					z.Slice = make([]S2, zjuy8)
				}
				for zznr2 := range z.Slice {
					bts, err = z.Slice[zznr2].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found1zuzv7[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var ztkb9 uint32
				ztkb9, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && ztkb9 > 0 {
					z.Transform = make(map[int]*S2, ztkb9)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for ztkb9 > 0 {
					var ztmx3 int
					var zamq4 *S2
					ztkb9--
					ztmx3, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zamq4 != nil {
							zamq4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zamq4 {
								zamq4.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zamq4 == nil {
								zamq4 = new(S2)
							}
							bts, err = zamq4.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[ztmx3] = zamq4
				}
			}
		case "Myptr_zid02_ptr":
			found1zuzv7[2] = true
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
			found1zuzv7[3] = true
			var zsnv10 uint32
			zsnv10, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zsnv10 != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zsnv10}
				return
			}
			for zbjo5 := range z.Myarray {
				z.Myarray[zbjo5], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found1zuzv7[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zgah11 uint32
				zgah11, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zgah11) {
					z.MySlice = (z.MySlice)[:zgah11]
				} else {
					z.MySlice = make([]string, zgah11)
				}
				for ztdw6 := range z.MySlice {
					z.MySlice[ztdw6], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zuzv7 != -1 {
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
var unmarshalMsgFieldOrder1zuzv7 = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip1zuzv7 = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zznr2 := range z.Slice {
		s += z.Slice[zznr2].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for ztmx3, zamq4 := range z.Transform {
			_ = zamq4
			_ = ztmx3
			s += msgp.IntSize
			if zamq4 == nil {
				s += msgp.NilSize
			} else {
				s += zamq4.MSGPMsgsize()
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
	for zbjo5 := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zbjo5])
	}
	s += 18 + msgp.ArrayHeaderSize
	for ztdw6 := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[ztdw6])
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
		for zmkf12, zcsk13 := range z.R {
			o = msgp.AppendString(o, zmkf12)
			o = msgp.AppendUint8(o, zcsk13)
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
		for zhis14 := range z.Arr {
			o = msgp.AppendFloat64(o, z.Arr[zhis14])
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
	const maxFields2znwf15 = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2znwf15 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2znwf15, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2znwf15 := totalEncodedFields2znwf15
	missingFieldsLeft2znwf15 := maxFields2znwf15 - totalEncodedFields2znwf15

	var nextMiss2znwf15 int32 = -1
	var found2znwf15 [maxFields2znwf15]bool
	var curField2znwf15 string

doneWithStruct2znwf15:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2znwf15 > 0 || missingFieldsLeft2znwf15 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2znwf15, missingFieldsLeft2znwf15, msgp.ShowFound(found2znwf15[:]), unmarshalMsgFieldOrder2znwf15)
		if encodedFieldsLeft2znwf15 > 0 {
			encodedFieldsLeft2znwf15--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2znwf15 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2znwf15 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2znwf15 = 0
			}
			for nextMiss2znwf15 < maxFields2znwf15 && (found2znwf15[nextMiss2znwf15] || unmarshalMsgFieldSkip2znwf15[nextMiss2znwf15]) {
				nextMiss2znwf15++
			}
			if nextMiss2znwf15 == maxFields2znwf15 {
				// filled all the empty fields!
				break doneWithStruct2znwf15
			}
			missingFieldsLeft2znwf15--
			curField2znwf15 = unmarshalMsgFieldOrder2znwf15[nextMiss2znwf15]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2znwf15)
		switch curField2znwf15 {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found2znwf15[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found2znwf15[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var znky16 uint32
				znky16, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && znky16 > 0 {
					z.R = make(map[string]uint8, znky16)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for znky16 > 0 {
					var zmkf12 string
					var zcsk13 uint8
					znky16--
					zmkf12, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zcsk13, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zmkf12] = zcsk13
				}
			}
		case "P_zid03_u16":
			found2znwf15[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found2znwf15[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found2znwf15[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found2znwf15[6] = true
			var zflh17 uint32
			zflh17, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zflh17 != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zflh17}
				return
			}
			for zhis14 := range z.Arr {
				z.Arr[zhis14], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found2znwf15[7] = true
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
	if nextMiss2znwf15 != -1 {
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
var unmarshalMsgFieldOrder2znwf15 = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip2znwf15 = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zmkf12, zcsk13 := range z.R {
			_ = zcsk13
			_ = zmkf12
			s += msgp.StringPrefixSize + len(zmkf12) + msgp.Uint8Size
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
	const maxFields3zhei18 = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zhei18 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zhei18, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zhei18 := totalEncodedFields3zhei18
	missingFieldsLeft3zhei18 := maxFields3zhei18 - totalEncodedFields3zhei18

	var nextMiss3zhei18 int32 = -1
	var found3zhei18 [maxFields3zhei18]bool
	var curField3zhei18 string

doneWithStruct3zhei18:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zhei18 > 0 || missingFieldsLeft3zhei18 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zhei18, missingFieldsLeft3zhei18, msgp.ShowFound(found3zhei18[:]), unmarshalMsgFieldOrder3zhei18)
		if encodedFieldsLeft3zhei18 > 0 {
			encodedFieldsLeft3zhei18--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zhei18 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zhei18 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zhei18 = 0
			}
			for nextMiss3zhei18 < maxFields3zhei18 && (found3zhei18[nextMiss3zhei18] || unmarshalMsgFieldSkip3zhei18[nextMiss3zhei18]) {
				nextMiss3zhei18++
			}
			if nextMiss3zhei18 == maxFields3zhei18 {
				// filled all the empty fields!
				break doneWithStruct3zhei18
			}
			missingFieldsLeft3zhei18--
			curField3zhei18 = unmarshalMsgFieldOrder3zhei18[nextMiss3zhei18]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zhei18)
		switch curField3zhei18 {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found3zhei18[0] = true
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
	if nextMiss3zhei18 != -1 {
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
var unmarshalMsgFieldOrder3zhei18 = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip3zhei18 = []bool{false}

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
		for zgrl19 := range z.Chld {
			o, err = z.Chld[zgrl19].MSGPMarshalMsg(o)
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
	const maxFields4zhro20 = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zhro20 uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zhro20, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zhro20 := totalEncodedFields4zhro20
	missingFieldsLeft4zhro20 := maxFields4zhro20 - totalEncodedFields4zhro20

	var nextMiss4zhro20 int32 = -1
	var found4zhro20 [maxFields4zhro20]bool
	var curField4zhro20 string

doneWithStruct4zhro20:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zhro20 > 0 || missingFieldsLeft4zhro20 > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zhro20, missingFieldsLeft4zhro20, msgp.ShowFound(found4zhro20[:]), unmarshalMsgFieldOrder4zhro20)
		if encodedFieldsLeft4zhro20 > 0 {
			encodedFieldsLeft4zhro20--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zhro20 = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zhro20 < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zhro20 = 0
			}
			for nextMiss4zhro20 < maxFields4zhro20 && (found4zhro20[nextMiss4zhro20] || unmarshalMsgFieldSkip4zhro20[nextMiss4zhro20]) {
				nextMiss4zhro20++
			}
			if nextMiss4zhro20 == maxFields4zhro20 {
				// filled all the empty fields!
				break doneWithStruct4zhro20
			}
			missingFieldsLeft4zhro20--
			curField4zhro20 = unmarshalMsgFieldOrder4zhro20[nextMiss4zhro20]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zhro20)
		switch curField4zhro20 {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found4zhro20[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zjko21 uint32
				zjko21, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zjko21) {
					z.Chld = (z.Chld)[:zjko21]
				} else {
					z.Chld = make([]Tree, zjko21)
				}
				for zgrl19 := range z.Chld {
					bts, err = z.Chld[zgrl19].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found4zhro20[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found4zhro20[2] = true
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
	if nextMiss4zhro20 != -1 {
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
var unmarshalMsgFieldOrder4zhro20 = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip4zhro20 = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zgrl19 := range z.Chld {
		s += z.Chld[zgrl19].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
