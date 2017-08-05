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
	isempty[2] = (len(z.Phone) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.Sibs == 0) // number, omitempty
	if isempty[3] {
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

	// string "name_zid00_str"
	o = append(o, 0xae, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.Name)
	// string "Bday_zid01_tim"
	o = append(o, 0xae, 0x42, 0x64, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x74, 0x69, 0x6d)
	o = msgp.AppendTime(o, z.Bday)
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

	// string "GPA_zid04_f64"
	o = append(o, 0xad, 0x47, 0x50, 0x41, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x66, 0x36, 0x34)
	o = msgp.AppendFloat64(o, z.GPA)
	// string "Friend_zid05_boo"
	o = append(o, 0xb0, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x62, 0x6f, 0x6f)
	o = msgp.AppendBool(o, z.Friend)
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
	const maxFields0ztkm = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0ztkm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0ztkm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0ztkm := totalEncodedFields0ztkm
	missingFieldsLeft0ztkm := maxFields0ztkm - totalEncodedFields0ztkm

	var nextMiss0ztkm int32 = -1
	var found0ztkm [maxFields0ztkm]bool
	var curField0ztkm string

doneWithStruct0ztkm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0ztkm > 0 || missingFieldsLeft0ztkm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0ztkm, missingFieldsLeft0ztkm, msgp.ShowFound(found0ztkm[:]), unmarshalMsgFieldOrder0ztkm)
		if encodedFieldsLeft0ztkm > 0 {
			encodedFieldsLeft0ztkm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0ztkm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0ztkm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0ztkm = 0
			}
			for nextMiss0ztkm < maxFields0ztkm && (found0ztkm[nextMiss0ztkm] || unmarshalMsgFieldSkip0ztkm[nextMiss0ztkm]) {
				nextMiss0ztkm++
			}
			if nextMiss0ztkm == maxFields0ztkm {
				// filled all the empty fields!
				break doneWithStruct0ztkm
			}
			missingFieldsLeft0ztkm--
			curField0ztkm = unmarshalMsgFieldOrder0ztkm[nextMiss0ztkm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0ztkm)
		switch curField0ztkm {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0ztkm[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0ztkm[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0ztkm[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0ztkm[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0ztkm[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0ztkm[5] = true
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
	if nextMiss0ztkm != -1 {
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
var unmarshalMsgFieldOrder0ztkm = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0ztkm = []bool{false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.Name) + 15 + msgp.TimeSize + 16 + msgp.StringPrefixSize + len(z.Phone) + 15 + msgp.IntSize + 14 + msgp.Float64Size + 17 + msgp.BoolSize
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Big) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 5
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Big) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 5
	// string "Slice_zid00_slc"
	o = append(o, 0x85, 0xaf, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
	for zmks := range z.Slice {
		o, err = z.Slice[zmks].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform_zid01_map"
	o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for ziux, ztmg := range z.Transform {
		o = msgp.AppendInt(o, ziux)
		if ztmg == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = ztmg.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
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
	// string "Myarray_zid03_ary"
	o = append(o, 0xb1, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x61, 0x72, 0x79)
	o = msgp.AppendArrayHeader(o, 3)
	for zqyw := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zqyw])
	}
	// string "MySlice_zid04_slc"
	o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zrsw := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zrsw])
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
	const maxFields1zmvq = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zmvq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zmvq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zmvq := totalEncodedFields1zmvq
	missingFieldsLeft1zmvq := maxFields1zmvq - totalEncodedFields1zmvq

	var nextMiss1zmvq int32 = -1
	var found1zmvq [maxFields1zmvq]bool
	var curField1zmvq string

doneWithStruct1zmvq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zmvq > 0 || missingFieldsLeft1zmvq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zmvq, missingFieldsLeft1zmvq, msgp.ShowFound(found1zmvq[:]), unmarshalMsgFieldOrder1zmvq)
		if encodedFieldsLeft1zmvq > 0 {
			encodedFieldsLeft1zmvq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zmvq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zmvq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zmvq = 0
			}
			for nextMiss1zmvq < maxFields1zmvq && (found1zmvq[nextMiss1zmvq] || unmarshalMsgFieldSkip1zmvq[nextMiss1zmvq]) {
				nextMiss1zmvq++
			}
			if nextMiss1zmvq == maxFields1zmvq {
				// filled all the empty fields!
				break doneWithStruct1zmvq
			}
			missingFieldsLeft1zmvq--
			curField1zmvq = unmarshalMsgFieldOrder1zmvq[nextMiss1zmvq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zmvq)
		switch curField1zmvq {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found1zmvq[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zgwn uint32
				zgwn, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zgwn) {
					z.Slice = (z.Slice)[:zgwn]
				} else {
					z.Slice = make([]S2, zgwn)
				}
				for zmks := range z.Slice {
					bts, err = z.Slice[zmks].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found1zmvq[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zpdt uint32
				zpdt, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zpdt > 0 {
					z.Transform = make(map[int]*S2, zpdt)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zpdt > 0 {
					var ziux int
					var ztmg *S2
					zpdt--
					ziux, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if ztmg != nil {
							ztmg.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != ztmg {
								ztmg.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if ztmg == nil {
								ztmg = new(S2)
							}
							bts, err = ztmg.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[ziux] = ztmg
				}
			}
		case "Myptr_zid02_ptr":
			found1zmvq[2] = true
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
			found1zmvq[3] = true
			var ztrl uint32
			ztrl, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && ztrl != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: ztrl}
				return
			}
			for zqyw := range z.Myarray {
				z.Myarray[zqyw], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found1zmvq[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zifm uint32
				zifm, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zifm) {
					z.MySlice = (z.MySlice)[:zifm]
				} else {
					z.MySlice = make([]string, zifm)
				}
				for zrsw := range z.MySlice {
					z.MySlice[zrsw], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zmvq != -1 {
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
var unmarshalMsgFieldOrder1zmvq = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip1zmvq = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zmks := range z.Slice {
		s += z.Slice[zmks].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for ziux, ztmg := range z.Transform {
			_ = ztmg
			_ = ziux
			s += msgp.IntSize
			if ztmg == nil {
				s += msgp.NilSize
			} else {
				s += ztmg.MSGPMsgsize()
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
	for zqyw := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zqyw])
	}
	s += 18 + msgp.ArrayHeaderSize
	for zrsw := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zrsw])
	}
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *S2) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 7
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *S2) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 7
	// string "beta_zid01_str"
	o = append(o, 0x87, 0xae, 0x62, 0x65, 0x74, 0x61, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.B)
	// string "ralph_zid02_map"
	o = append(o, 0xaf, 0x72, 0x61, 0x6c, 0x70, 0x68, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.R)))
	for zhbo, zzmn := range z.R {
		o = msgp.AppendString(o, zhbo)
		o = msgp.AppendUint8(o, zzmn)
	}
	// string "P_zid03_u16"
	o = append(o, 0xab, 0x50, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x33, 0x5f, 0x75, 0x31, 0x36)
	o = msgp.AppendUint16(o, z.P)
	// string "Q_zid04_u32"
	o = append(o, 0xab, 0x51, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x75, 0x33, 0x32)
	o = msgp.AppendUint32(o, z.Q)
	// string "T_zid05_i64"
	o = append(o, 0xab, 0x54, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x35, 0x5f, 0x69, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.T)
	// string "arr_zid06_ary"
	o = append(o, 0xad, 0x61, 0x72, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x36, 0x5f, 0x61, 0x72, 0x79)
	o = msgp.AppendArrayHeader(o, 6)
	for zvdo := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zvdo])
	}
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
	const maxFields2zlby = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zlby uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zlby, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zlby := totalEncodedFields2zlby
	missingFieldsLeft2zlby := maxFields2zlby - totalEncodedFields2zlby

	var nextMiss2zlby int32 = -1
	var found2zlby [maxFields2zlby]bool
	var curField2zlby string

doneWithStruct2zlby:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zlby > 0 || missingFieldsLeft2zlby > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zlby, missingFieldsLeft2zlby, msgp.ShowFound(found2zlby[:]), unmarshalMsgFieldOrder2zlby)
		if encodedFieldsLeft2zlby > 0 {
			encodedFieldsLeft2zlby--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zlby = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zlby < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zlby = 0
			}
			for nextMiss2zlby < maxFields2zlby && (found2zlby[nextMiss2zlby] || unmarshalMsgFieldSkip2zlby[nextMiss2zlby]) {
				nextMiss2zlby++
			}
			if nextMiss2zlby == maxFields2zlby {
				// filled all the empty fields!
				break doneWithStruct2zlby
			}
			missingFieldsLeft2zlby--
			curField2zlby = unmarshalMsgFieldOrder2zlby[nextMiss2zlby]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zlby)
		switch curField2zlby {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found2zlby[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found2zlby[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zbvb uint32
				zbvb, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zbvb > 0 {
					z.R = make(map[string]uint8, zbvb)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zbvb > 0 {
					var zhbo string
					var zzmn uint8
					zbvb--
					zhbo, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zzmn, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zhbo] = zzmn
				}
			}
		case "P_zid03_u16":
			found2zlby[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found2zlby[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found2zlby[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found2zlby[6] = true
			var zpss uint32
			zpss, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zpss != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zpss}
				return
			}
			for zvdo := range z.Arr {
				z.Arr[zvdo], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found2zlby[7] = true
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
	if nextMiss2zlby != -1 {
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
var unmarshalMsgFieldOrder2zlby = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip2zlby = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zhbo, zzmn := range z.R {
			_ = zzmn
			_ = zhbo
			s += msgp.StringPrefixSize + len(zhbo) + msgp.Uint8Size
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
	return 1
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z Sys) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 1
	// string "F_zid00_ifc"
	o = append(o, 0x81, 0xab, 0x46, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x69, 0x66, 0x63)
	o, err = msgp.AppendIntf(o, z.F)
	if err != nil {
		return
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
	const maxFields3zvvl = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zvvl uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zvvl, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zvvl := totalEncodedFields3zvvl
	missingFieldsLeft3zvvl := maxFields3zvvl - totalEncodedFields3zvvl

	var nextMiss3zvvl int32 = -1
	var found3zvvl [maxFields3zvvl]bool
	var curField3zvvl string

doneWithStruct3zvvl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zvvl > 0 || missingFieldsLeft3zvvl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvvl, missingFieldsLeft3zvvl, msgp.ShowFound(found3zvvl[:]), unmarshalMsgFieldOrder3zvvl)
		if encodedFieldsLeft3zvvl > 0 {
			encodedFieldsLeft3zvvl--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zvvl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zvvl < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zvvl = 0
			}
			for nextMiss3zvvl < maxFields3zvvl && (found3zvvl[nextMiss3zvvl] || unmarshalMsgFieldSkip3zvvl[nextMiss3zvvl]) {
				nextMiss3zvvl++
			}
			if nextMiss3zvvl == maxFields3zvvl {
				// filled all the empty fields!
				break doneWithStruct3zvvl
			}
			missingFieldsLeft3zvvl--
			curField3zvvl = unmarshalMsgFieldOrder3zvvl[nextMiss3zvvl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zvvl)
		switch curField3zvvl {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found3zvvl[0] = true
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
	if nextMiss3zvvl != -1 {
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
var unmarshalMsgFieldOrder3zvvl = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip3zvvl = []bool{false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) MSGPMsgsize() (s int) {
	s = 1 + 12 + msgp.GuessSize(z.F)
	return
}

// MSGPfieldsNotEmpty supports omitempty tags
func (z *Tree) MSGPfieldsNotEmpty(isempty []bool) uint32 {
	return 3
}

// MSGPMarshalMsg implements msgp.Marshaler
func (z *Tree) MSGPMarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.MSGPMsgsize())
	// map header, size 3
	// string "Chld_zid00_slc"
	o = append(o, 0x83, 0xae, 0x43, 0x68, 0x6c, 0x64, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x30, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
	for zdpj := range z.Chld {
		o, err = z.Chld[zdpj].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Str_zid01_str"
	o = append(o, 0xad, 0x53, 0x74, 0x72, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
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
	const maxFields4zvjb = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zvjb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zvjb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zvjb := totalEncodedFields4zvjb
	missingFieldsLeft4zvjb := maxFields4zvjb - totalEncodedFields4zvjb

	var nextMiss4zvjb int32 = -1
	var found4zvjb [maxFields4zvjb]bool
	var curField4zvjb string

doneWithStruct4zvjb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zvjb > 0 || missingFieldsLeft4zvjb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zvjb, missingFieldsLeft4zvjb, msgp.ShowFound(found4zvjb[:]), unmarshalMsgFieldOrder4zvjb)
		if encodedFieldsLeft4zvjb > 0 {
			encodedFieldsLeft4zvjb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zvjb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zvjb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zvjb = 0
			}
			for nextMiss4zvjb < maxFields4zvjb && (found4zvjb[nextMiss4zvjb] || unmarshalMsgFieldSkip4zvjb[nextMiss4zvjb]) {
				nextMiss4zvjb++
			}
			if nextMiss4zvjb == maxFields4zvjb {
				// filled all the empty fields!
				break doneWithStruct4zvjb
			}
			missingFieldsLeft4zvjb--
			curField4zvjb = unmarshalMsgFieldOrder4zvjb[nextMiss4zvjb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zvjb)
		switch curField4zvjb {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found4zvjb[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zuey uint32
				zuey, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zuey) {
					z.Chld = (z.Chld)[:zuey]
				} else {
					z.Chld = make([]Tree, zuey)
				}
				for zdpj := range z.Chld {
					bts, err = z.Chld[zdpj].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found4zvjb[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found4zvjb[2] = true
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
	if nextMiss4zvjb != -1 {
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
var unmarshalMsgFieldOrder4zvjb = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip4zvjb = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zdpj := range z.Chld {
		s += z.Chld[zdpj].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
