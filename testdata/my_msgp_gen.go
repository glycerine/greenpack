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
	const maxFields0zaby = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zaby uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zaby, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zaby := totalEncodedFields0zaby
	missingFieldsLeft0zaby := maxFields0zaby - totalEncodedFields0zaby

	var nextMiss0zaby int32 = -1
	var found0zaby [maxFields0zaby]bool
	var curField0zaby string

doneWithStruct0zaby:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaby > 0 || missingFieldsLeft0zaby > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaby, missingFieldsLeft0zaby, msgp.ShowFound(found0zaby[:]), unmarshalMsgFieldOrder0zaby)
		if encodedFieldsLeft0zaby > 0 {
			encodedFieldsLeft0zaby--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zaby = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zaby < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zaby = 0
			}
			for nextMiss0zaby < maxFields0zaby && (found0zaby[nextMiss0zaby] || unmarshalMsgFieldSkip0zaby[nextMiss0zaby]) {
				nextMiss0zaby++
			}
			if nextMiss0zaby == maxFields0zaby {
				// filled all the empty fields!
				break doneWithStruct0zaby
			}
			missingFieldsLeft0zaby--
			curField0zaby = unmarshalMsgFieldOrder0zaby[nextMiss0zaby]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaby)
		switch curField0zaby {
		// -- templateUnmarshalMsg ends here --

		case "name_zid00_str":
			found0zaby[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday_zid01_tim":
			found0zaby[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone_zid02_str":
			found0zaby[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs_zid03_int":
			found0zaby[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA_zid04_f64":
			found0zaby[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend_zid05_boo":
			found0zaby[5] = true
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
	if nextMiss0zaby != -1 {
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
var unmarshalMsgFieldOrder0zaby = []string{"name_zid00_str", "Bday_zid01_tim", "phone_zid02_str", "Sibs_zid03_int", "GPA_zid04_f64", "Friend_zid05_boo"}

var unmarshalMsgFieldSkip0zaby = []bool{false, false, false, false, false, false}

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
		for zmhv := range z.Slice {
			o, err = z.Slice[zmhv].MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// string "Transform_zid01_map"
		o = append(o, 0xb3, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x31, 0x5f, 0x6d, 0x61, 0x70)
		o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
		for zbar, zicr := range z.Transform {
			o = msgp.AppendInt(o, zbar)
			if zicr == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zicr.MSGPMarshalMsg(o)
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
		for zpgj := range z.Myarray {
			o = msgp.AppendString(o, z.Myarray[zpgj])
		}
	}

	if !empty[4] {
		// string "MySlice_zid04_slc"
		o = append(o, 0xb1, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x69, 0x64, 0x30, 0x34, 0x5f, 0x73, 0x6c, 0x63)
		o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
		for zdte := range z.MySlice {
			o = msgp.AppendString(o, z.MySlice[zdte])
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
	const maxFields1zkdl = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zkdl uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zkdl, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zkdl := totalEncodedFields1zkdl
	missingFieldsLeft1zkdl := maxFields1zkdl - totalEncodedFields1zkdl

	var nextMiss1zkdl int32 = -1
	var found1zkdl [maxFields1zkdl]bool
	var curField1zkdl string

doneWithStruct1zkdl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zkdl > 0 || missingFieldsLeft1zkdl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zkdl, missingFieldsLeft1zkdl, msgp.ShowFound(found1zkdl[:]), unmarshalMsgFieldOrder1zkdl)
		if encodedFieldsLeft1zkdl > 0 {
			encodedFieldsLeft1zkdl--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zkdl = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zkdl < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zkdl = 0
			}
			for nextMiss1zkdl < maxFields1zkdl && (found1zkdl[nextMiss1zkdl] || unmarshalMsgFieldSkip1zkdl[nextMiss1zkdl]) {
				nextMiss1zkdl++
			}
			if nextMiss1zkdl == maxFields1zkdl {
				// filled all the empty fields!
				break doneWithStruct1zkdl
			}
			missingFieldsLeft1zkdl--
			curField1zkdl = unmarshalMsgFieldOrder1zkdl[nextMiss1zkdl]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zkdl)
		switch curField1zkdl {
		// -- templateUnmarshalMsg ends here --

		case "Slice_zid00_slc":
			found1zkdl[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zwin uint32
				zwin, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zwin) {
					z.Slice = (z.Slice)[:zwin]
				} else {
					z.Slice = make([]S2, zwin)
				}
				for zmhv := range z.Slice {
					bts, err = z.Slice[zmhv].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform_zid01_map":
			found1zkdl[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var ztbz uint32
				ztbz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && ztbz > 0 {
					z.Transform = make(map[int]*S2, ztbz)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for ztbz > 0 {
					var zbar int
					var zicr *S2
					ztbz--
					zbar, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zicr != nil {
							zicr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zicr {
								zicr.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zicr == nil {
								zicr = new(S2)
							}
							bts, err = zicr.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zbar] = zicr
				}
			}
		case "Myptr_zid02_ptr":
			found1zkdl[2] = true
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
			found1zkdl[3] = true
			var zlmh uint32
			zlmh, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zlmh != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zlmh}
				return
			}
			for zpgj := range z.Myarray {
				z.Myarray[zpgj], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice_zid04_slc":
			found1zkdl[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zlhu uint32
				zlhu, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zlhu) {
					z.MySlice = (z.MySlice)[:zlhu]
				} else {
					z.MySlice = make([]string, zlhu)
				}
				for zdte := range z.MySlice {
					z.MySlice[zdte], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zkdl != -1 {
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
var unmarshalMsgFieldOrder1zkdl = []string{"Slice_zid00_slc", "Transform_zid01_map", "Myptr_zid02_ptr", "Myarray_zid03_ary", "MySlice_zid04_slc"}

var unmarshalMsgFieldSkip1zkdl = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 16 + msgp.ArrayHeaderSize
	for zmhv := range z.Slice {
		s += z.Slice[zmhv].MSGPMsgsize()
	}
	s += 20 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zbar, zicr := range z.Transform {
			_ = zicr
			_ = zbar
			s += msgp.IntSize
			if zicr == nil {
				s += msgp.NilSize
			} else {
				s += zicr.MSGPMsgsize()
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
	for zpgj := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zpgj])
	}
	s += 18 + msgp.ArrayHeaderSize
	for zdte := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zdte])
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
		for zrfn, zlam := range z.R {
			o = msgp.AppendString(o, zrfn)
			o = msgp.AppendUint8(o, zlam)
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
		for zudb := range z.Arr {
			o = msgp.AppendFloat64(o, z.Arr[zudb])
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
	const maxFields2zjkf = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zjkf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zjkf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zjkf := totalEncodedFields2zjkf
	missingFieldsLeft2zjkf := maxFields2zjkf - totalEncodedFields2zjkf

	var nextMiss2zjkf int32 = -1
	var found2zjkf [maxFields2zjkf]bool
	var curField2zjkf string

doneWithStruct2zjkf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zjkf > 0 || missingFieldsLeft2zjkf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zjkf, missingFieldsLeft2zjkf, msgp.ShowFound(found2zjkf[:]), unmarshalMsgFieldOrder2zjkf)
		if encodedFieldsLeft2zjkf > 0 {
			encodedFieldsLeft2zjkf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zjkf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zjkf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zjkf = 0
			}
			for nextMiss2zjkf < maxFields2zjkf && (found2zjkf[nextMiss2zjkf] || unmarshalMsgFieldSkip2zjkf[nextMiss2zjkf]) {
				nextMiss2zjkf++
			}
			if nextMiss2zjkf == maxFields2zjkf {
				// filled all the empty fields!
				break doneWithStruct2zjkf
			}
			missingFieldsLeft2zjkf--
			curField2zjkf = unmarshalMsgFieldOrder2zjkf[nextMiss2zjkf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zjkf)
		switch curField2zjkf {
		// -- templateUnmarshalMsg ends here --

		case "beta_zid01_str":
			found2zjkf[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph_zid02_map":
			found2zjkf[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zuap uint32
				zuap, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zuap > 0 {
					z.R = make(map[string]uint8, zuap)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zuap > 0 {
					var zrfn string
					var zlam uint8
					zuap--
					zrfn, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zlam, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zrfn] = zlam
				}
			}
		case "P_zid03_u16":
			found2zjkf[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q_zid04_u32":
			found2zjkf[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T_zid05_i64":
			found2zjkf[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr_zid06_ary":
			found2zjkf[6] = true
			var zhsj uint32
			zhsj, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zhsj != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zhsj}
				return
			}
			for zudb := range z.Arr {
				z.Arr[zudb], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree_zid07_ptr":
			found2zjkf[7] = true
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
	if nextMiss2zjkf != -1 {
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
var unmarshalMsgFieldOrder2zjkf = []string{"", "beta_zid01_str", "ralph_zid02_map", "P_zid03_u16", "Q_zid04_u32", "T_zid05_i64", "arr_zid06_ary", "MyTree_zid07_ptr"}

var unmarshalMsgFieldSkip2zjkf = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.B) + 16 + msgp.MapHeaderSize
	if z.R != nil {
		for zrfn, zlam := range z.R {
			_ = zlam
			_ = zrfn
			s += msgp.StringPrefixSize + len(zrfn) + msgp.Uint8Size
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
	const maxFields3zvlo = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zvlo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zvlo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zvlo := totalEncodedFields3zvlo
	missingFieldsLeft3zvlo := maxFields3zvlo - totalEncodedFields3zvlo

	var nextMiss3zvlo int32 = -1
	var found3zvlo [maxFields3zvlo]bool
	var curField3zvlo string

doneWithStruct3zvlo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zvlo > 0 || missingFieldsLeft3zvlo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zvlo, missingFieldsLeft3zvlo, msgp.ShowFound(found3zvlo[:]), unmarshalMsgFieldOrder3zvlo)
		if encodedFieldsLeft3zvlo > 0 {
			encodedFieldsLeft3zvlo--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zvlo = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zvlo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zvlo = 0
			}
			for nextMiss3zvlo < maxFields3zvlo && (found3zvlo[nextMiss3zvlo] || unmarshalMsgFieldSkip3zvlo[nextMiss3zvlo]) {
				nextMiss3zvlo++
			}
			if nextMiss3zvlo == maxFields3zvlo {
				// filled all the empty fields!
				break doneWithStruct3zvlo
			}
			missingFieldsLeft3zvlo--
			curField3zvlo = unmarshalMsgFieldOrder3zvlo[nextMiss3zvlo]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zvlo)
		switch curField3zvlo {
		// -- templateUnmarshalMsg ends here --

		case "F_zid00_ifc":
			found3zvlo[0] = true
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
	if nextMiss3zvlo != -1 {
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
var unmarshalMsgFieldOrder3zvlo = []string{"F_zid00_ifc"}

var unmarshalMsgFieldSkip3zvlo = []bool{false}

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
		for zula := range z.Chld {
			o, err = z.Chld[zula].MSGPMarshalMsg(o)
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
	const maxFields4zdtb = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zdtb uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zdtb, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zdtb := totalEncodedFields4zdtb
	missingFieldsLeft4zdtb := maxFields4zdtb - totalEncodedFields4zdtb

	var nextMiss4zdtb int32 = -1
	var found4zdtb [maxFields4zdtb]bool
	var curField4zdtb string

doneWithStruct4zdtb:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zdtb > 0 || missingFieldsLeft4zdtb > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zdtb, missingFieldsLeft4zdtb, msgp.ShowFound(found4zdtb[:]), unmarshalMsgFieldOrder4zdtb)
		if encodedFieldsLeft4zdtb > 0 {
			encodedFieldsLeft4zdtb--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zdtb = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zdtb < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zdtb = 0
			}
			for nextMiss4zdtb < maxFields4zdtb && (found4zdtb[nextMiss4zdtb] || unmarshalMsgFieldSkip4zdtb[nextMiss4zdtb]) {
				nextMiss4zdtb++
			}
			if nextMiss4zdtb == maxFields4zdtb {
				// filled all the empty fields!
				break doneWithStruct4zdtb
			}
			missingFieldsLeft4zdtb--
			curField4zdtb = unmarshalMsgFieldOrder4zdtb[nextMiss4zdtb]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zdtb)
		switch curField4zdtb {
		// -- templateUnmarshalMsg ends here --

		case "Chld_zid00_slc":
			found4zdtb[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zgos uint32
				zgos, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zgos) {
					z.Chld = (z.Chld)[:zgos]
				} else {
					z.Chld = make([]Tree, zgos)
				}
				for zula := range z.Chld {
					bts, err = z.Chld[zula].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str_zid01_str":
			found4zdtb[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par_zid02_ptr":
			found4zdtb[2] = true
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
	if nextMiss4zdtb != -1 {
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
var unmarshalMsgFieldOrder4zdtb = []string{"Chld_zid00_slc", "Str_zid01_str", "Par_zid02_ptr"}

var unmarshalMsgFieldSkip4zdtb = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 15 + msgp.ArrayHeaderSize
	for zula := range z.Chld {
		s += z.Chld[zula].MSGPMsgsize()
	}
	s += 14 + msgp.StringPrefixSize + len(z.Str) + 14
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
