package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack2/msgp"
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
	const maxFields0zixs = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zixs uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zixs, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zixs := totalEncodedFields0zixs
	missingFieldsLeft0zixs := maxFields0zixs - totalEncodedFields0zixs

	var nextMiss0zixs int32 = -1
	var found0zixs [maxFields0zixs]bool
	var curField0zixs string

doneWithStruct0zixs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zixs > 0 || missingFieldsLeft0zixs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zixs, missingFieldsLeft0zixs, msgp.ShowFound(found0zixs[:]), unmarshalMsgFieldOrder0zixs)
		if encodedFieldsLeft0zixs > 0 {
			encodedFieldsLeft0zixs--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zixs = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zixs < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zixs = 0
			}
			for nextMiss0zixs < maxFields0zixs && (found0zixs[nextMiss0zixs] || unmarshalMsgFieldSkip0zixs[nextMiss0zixs]) {
				nextMiss0zixs++
			}
			if nextMiss0zixs == maxFields0zixs {
				// filled all the empty fields!
				break doneWithStruct0zixs
			}
			missingFieldsLeft0zixs--
			curField0zixs = unmarshalMsgFieldOrder0zixs[nextMiss0zixs]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zixs)
		switch curField0zixs {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0zixs[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0zixs[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0zixs[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0zixs[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0zixs[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0zixs[5] = true
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
	if nextMiss0zixs != -1 {
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
var unmarshalMsgFieldOrder0zixs = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0zixs = []bool{false, false, false, false, false, false}

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
	for zoub := range z.Slice {
		o, err = z.Slice[zoub].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform_zid01_map"
	o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zwbr, zbml := range z.Transform {
		o = msgp.AppendInt(o, zwbr)
		if zbml == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zbml.MSGPMarshalMsg(o)
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
	for zokd := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zokd])
	}
	// string "MySlice_zid04_slc"
	o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for ztrj := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[ztrj])
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
	const maxFields1zavm = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zavm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zavm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zavm := totalEncodedFields1zavm
	missingFieldsLeft1zavm := maxFields1zavm - totalEncodedFields1zavm

	var nextMiss1zavm int32 = -1
	var found1zavm [maxFields1zavm]bool
	var curField1zavm string

doneWithStruct1zavm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zavm > 0 || missingFieldsLeft1zavm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zavm, missingFieldsLeft1zavm, msgp.ShowFound(found1zavm[:]), unmarshalMsgFieldOrder1zavm)
		if encodedFieldsLeft1zavm > 0 {
			encodedFieldsLeft1zavm--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zavm = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zavm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zavm = 0
			}
			for nextMiss1zavm < maxFields1zavm && (found1zavm[nextMiss1zavm] || unmarshalMsgFieldSkip1zavm[nextMiss1zavm]) {
				nextMiss1zavm++
			}
			if nextMiss1zavm == maxFields1zavm {
				// filled all the empty fields!
				break doneWithStruct1zavm
			}
			missingFieldsLeft1zavm--
			curField1zavm = unmarshalMsgFieldOrder1zavm[nextMiss1zavm]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zavm)
		switch curField1zavm {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found1zavm[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zgvz uint32
				zgvz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zgvz) {
					z.Slice = (z.Slice)[:zgvz]
				} else {
					z.Slice = make([]S2, zgvz)
				}
				for zoub := range z.Slice {
					bts, err = z.Slice[zoub].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found1zavm[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zmbo uint32
				zmbo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zmbo > 0 {
					z.Transform = make(map[int]*S2, zmbo)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zmbo > 0 {
					var zwbr int
					var zbml *S2
					zmbo--
					zwbr, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zbml != nil {
							zbml.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zbml {
								zbml.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zbml == nil {
								zbml = new(S2)
							}
							bts, err = zbml.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zwbr] = zbml
				}
			}
		case "Myptr_zid02_ptr":
			found1zavm[2] = true
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
			found1zavm[3] = true
			var zpqq uint32
			zpqq, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zpqq != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zpqq}
				return
			}
			for zokd := range z.Myarray {
				z.Myarray[zokd], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found1zavm[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zuhs uint32
				zuhs, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zuhs) {
					z.MySlice = (z.MySlice)[:zuhs]
				} else {
					z.MySlice = make([]string, zuhs)
				}
				for ztrj := range z.MySlice {
					z.MySlice[ztrj], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zavm != -1 {
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
var unmarshalMsgFieldOrder1zavm = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip1zavm = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zoub := range z.Slice {
		s += z.Slice[zoub].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zwbr, zbml := range z.Transform {
			_ = zbml
			_ = zwbr
			s += msgp.IntSize
			if zbml == nil {
				s += msgp.NilSize
			} else {
				s += zbml.MSGPMsgsize()
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
	for zokd := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zokd])
	}
	s += 18 + msgp.ArrayHeaderSize
	for ztrj := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[ztrj])
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
	for zlxi, zbqt := range z.R {
		o = msgp.AppendString(o, zlxi)
		o = msgp.AppendUint8(o, zbqt)
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
	for zuub := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zuub])
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
	const maxFields2zlgc = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zlgc uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zlgc, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zlgc := totalEncodedFields2zlgc
	missingFieldsLeft2zlgc := maxFields2zlgc - totalEncodedFields2zlgc

	var nextMiss2zlgc int32 = -1
	var found2zlgc [maxFields2zlgc]bool
	var curField2zlgc string

doneWithStruct2zlgc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zlgc > 0 || missingFieldsLeft2zlgc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zlgc, missingFieldsLeft2zlgc, msgp.ShowFound(found2zlgc[:]), unmarshalMsgFieldOrder2zlgc)
		if encodedFieldsLeft2zlgc > 0 {
			encodedFieldsLeft2zlgc--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zlgc = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zlgc < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zlgc = 0
			}
			for nextMiss2zlgc < maxFields2zlgc && (found2zlgc[nextMiss2zlgc] || unmarshalMsgFieldSkip2zlgc[nextMiss2zlgc]) {
				nextMiss2zlgc++
			}
			if nextMiss2zlgc == maxFields2zlgc {
				// filled all the empty fields!
				break doneWithStruct2zlgc
			}
			missingFieldsLeft2zlgc--
			curField2zlgc = unmarshalMsgFieldOrder2zlgc[nextMiss2zlgc]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zlgc)
		switch curField2zlgc {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found2zlgc[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found2zlgc[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zjgc uint32
				zjgc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zjgc > 0 {
					z.R = make(map[string]uint8, zjgc)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zjgc > 0 {
					var zlxi string
					var zbqt uint8
					zjgc--
					zlxi, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zbqt, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zlxi] = zbqt
				}
			}
		case "P_zid03_u16":
			found2zlgc[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found2zlgc[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found2zlgc[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found2zlgc[6] = true
			var zuxv uint32
			zuxv, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zuxv != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zuxv}
				return
			}
			for zuub := range z.Arr {
				z.Arr[zuub], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found2zlgc[7] = true
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
	if nextMiss2zlgc != -1 {
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
var unmarshalMsgFieldOrder2zlgc = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip2zlgc = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zlxi, zbqt := range z.R {
			_ = zbqt
			_ = zlxi
			s += msgp.StringPrefixSize + len(zlxi) + msgp.Uint8Size
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
	const maxFields3zfwb = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zfwb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zfwb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zfwb := totalEncodedFields3zfwb
	missingFieldsLeft3zfwb := maxFields3zfwb - totalEncodedFields3zfwb

	var nextMiss3zfwb int32 = -1
	var found3zfwb [maxFields3zfwb]bool
	var curField3zfwb string

doneWithStruct3zfwb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zfwb > 0 || missingFieldsLeft3zfwb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zfwb, missingFieldsLeft3zfwb, msgp.ShowFound(found3zfwb[:]), unmarshalMsgFieldOrder3zfwb)
		if encodedFieldsLeft3zfwb > 0 {
			encodedFieldsLeft3zfwb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zfwb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zfwb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zfwb = 0
			}
			for nextMiss3zfwb < maxFields3zfwb && (found3zfwb[nextMiss3zfwb] || unmarshalMsgFieldSkip3zfwb[nextMiss3zfwb]) {
				nextMiss3zfwb++
			}
			if nextMiss3zfwb == maxFields3zfwb {
				// filled all the empty fields!
				break doneWithStruct3zfwb
			}
			missingFieldsLeft3zfwb--
			curField3zfwb = unmarshalMsgFieldOrder3zfwb[nextMiss3zfwb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zfwb)
		switch curField3zfwb {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found3zfwb[0] = true
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
	if nextMiss3zfwb != -1 {
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
var unmarshalMsgFieldOrder3zfwb = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip3zfwb = []bool{false}

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
	for zsqs := range z.Chld {
		o, err = z.Chld[zsqs].MSGPMarshalMsg(o)
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
	const maxFields4zmxn = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zmxn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zmxn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zmxn := totalEncodedFields4zmxn
	missingFieldsLeft4zmxn := maxFields4zmxn - totalEncodedFields4zmxn

	var nextMiss4zmxn int32 = -1
	var found4zmxn [maxFields4zmxn]bool
	var curField4zmxn string

doneWithStruct4zmxn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zmxn > 0 || missingFieldsLeft4zmxn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zmxn, missingFieldsLeft4zmxn, msgp.ShowFound(found4zmxn[:]), unmarshalMsgFieldOrder4zmxn)
		if encodedFieldsLeft4zmxn > 0 {
			encodedFieldsLeft4zmxn--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zmxn = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zmxn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zmxn = 0
			}
			for nextMiss4zmxn < maxFields4zmxn && (found4zmxn[nextMiss4zmxn] || unmarshalMsgFieldSkip4zmxn[nextMiss4zmxn]) {
				nextMiss4zmxn++
			}
			if nextMiss4zmxn == maxFields4zmxn {
				// filled all the empty fields!
				break doneWithStruct4zmxn
			}
			missingFieldsLeft4zmxn--
			curField4zmxn = unmarshalMsgFieldOrder4zmxn[nextMiss4zmxn]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zmxn)
		switch curField4zmxn {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found4zmxn[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zhsq uint32
				zhsq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zhsq) {
					z.Chld = (z.Chld)[:zhsq]
				} else {
					z.Chld = make([]Tree, zhsq)
				}
				for zsqs := range z.Chld {
					bts, err = z.Chld[zsqs].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found4zmxn[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found4zmxn[2] = true
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
	if nextMiss4zmxn != -1 {
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
var unmarshalMsgFieldOrder4zmxn = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip4zmxn = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zsqs := range z.Chld {
		s += z.Chld[zsqs].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
