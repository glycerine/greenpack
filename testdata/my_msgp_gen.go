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

	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Bday"
	o = append(o, 0xa4, 0x42, 0x64, 0x61, 0x79)
	o = msgp.AppendTime(o, z.Bday)
	if !empty[2] {
		// string "phone"
		o = append(o, 0xa5, 0x70, 0x68, 0x6f, 0x6e, 0x65)
		o = msgp.AppendString(o, z.Phone)
	}

	if !empty[3] {
		// string "Sibs"
		o = append(o, 0xa4, 0x53, 0x69, 0x62, 0x73)
		o = msgp.AppendInt(o, z.Sibs)
	}

	// string "GPA"
	o = append(o, 0xa3, 0x47, 0x50, 0x41)
	o = msgp.AppendFloat64(o, z.GPA)
	// string "Friend"
	o = append(o, 0xa6, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64)
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
	const maxFields0zgqj = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zgqj uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zgqj, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zgqj := totalEncodedFields0zgqj
	missingFieldsLeft0zgqj := maxFields0zgqj - totalEncodedFields0zgqj

	var nextMiss0zgqj int32 = -1
	var found0zgqj [maxFields0zgqj]bool
	var curField0zgqj string

doneWithStruct0zgqj:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zgqj > 0 || missingFieldsLeft0zgqj > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zgqj, missingFieldsLeft0zgqj, msgp.ShowFound(found0zgqj[:]), unmarshalMsgFieldOrder0zgqj)
		if encodedFieldsLeft0zgqj > 0 {
			encodedFieldsLeft0zgqj--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zgqj = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zgqj < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zgqj = 0
			}
			for nextMiss0zgqj < maxFields0zgqj && (found0zgqj[nextMiss0zgqj] || unmarshalMsgFieldSkip0zgqj[nextMiss0zgqj]) {
				nextMiss0zgqj++
			}
			if nextMiss0zgqj == maxFields0zgqj {
				// filled all the empty fields!
				break doneWithStruct0zgqj
			}
			missingFieldsLeft0zgqj--
			curField0zgqj = unmarshalMsgFieldOrder0zgqj[nextMiss0zgqj]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zgqj)
		switch curField0zgqj {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zgqj[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday":
			found0zgqj[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone":
			found0zgqj[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs":
			found0zgqj[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA":
			found0zgqj[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend":
			found0zgqj[5] = true
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
	if nextMiss0zgqj != -1 {
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
var unmarshalMsgFieldOrder0zgqj = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zgqj = []bool{false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *A) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.TimeSize + 6 + msgp.StringPrefixSize + len(z.Phone) + 5 + msgp.IntSize + 4 + msgp.Float64Size + 7 + msgp.BoolSize
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
	// string "Slice"
	o = append(o, 0x85, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Slice)))
	for zvbk := range z.Slice {
		o, err = z.Slice[zvbk].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zrep, zgoz := range z.Transform {
		o = msgp.AppendInt(o, zrep)
		if zgoz == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zgoz.MSGPMarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Myptr"
	o = append(o, 0xa5, 0x4d, 0x79, 0x70, 0x74, 0x72)
	if z.Myptr == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Myptr.MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Myarray"
	o = append(o, 0xa7, 0x4d, 0x79, 0x61, 0x72, 0x72, 0x61, 0x79)
	o = msgp.AppendArrayHeader(o, 3)
	for zzpe := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[zzpe])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zatn := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zatn])
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
	const maxFields1zptq = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zptq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zptq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zptq := totalEncodedFields1zptq
	missingFieldsLeft1zptq := maxFields1zptq - totalEncodedFields1zptq

	var nextMiss1zptq int32 = -1
	var found1zptq [maxFields1zptq]bool
	var curField1zptq string

doneWithStruct1zptq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zptq > 0 || missingFieldsLeft1zptq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zptq, missingFieldsLeft1zptq, msgp.ShowFound(found1zptq[:]), unmarshalMsgFieldOrder1zptq)
		if encodedFieldsLeft1zptq > 0 {
			encodedFieldsLeft1zptq--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zptq = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zptq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zptq = 0
			}
			for nextMiss1zptq < maxFields1zptq && (found1zptq[nextMiss1zptq] || unmarshalMsgFieldSkip1zptq[nextMiss1zptq]) {
				nextMiss1zptq++
			}
			if nextMiss1zptq == maxFields1zptq {
				// filled all the empty fields!
				break doneWithStruct1zptq
			}
			missingFieldsLeft1zptq--
			curField1zptq = unmarshalMsgFieldOrder1zptq[nextMiss1zptq]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zptq)
		switch curField1zptq {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zptq[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var ztka uint32
				ztka, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(ztka) {
					z.Slice = (z.Slice)[:ztka]
				} else {
					z.Slice = make([]S2, ztka)
				}
				for zvbk := range z.Slice {
					bts, err = z.Slice[zvbk].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform":
			found1zptq[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var ztgj uint32
				ztgj, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && ztgj > 0 {
					z.Transform = make(map[int]*S2, ztgj)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for ztgj > 0 {
					var zrep int
					var zgoz *S2
					ztgj--
					zrep, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zgoz != nil {
							zgoz.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zgoz {
								zgoz.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zgoz == nil {
								zgoz = new(S2)
							}
							bts, err = zgoz.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zrep] = zgoz
				}
			}
		case "Myptr":
			found1zptq[2] = true
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
		case "Myarray":
			found1zptq[3] = true
			var zplw uint32
			zplw, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zplw != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zplw}
				return
			}
			for zzpe := range z.Myarray {
				z.Myarray[zzpe], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice":
			found1zptq[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zhwd uint32
				zhwd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zhwd) {
					z.MySlice = (z.MySlice)[:zhwd]
				} else {
					z.MySlice = make([]string, zhwd)
				}
				for zatn := range z.MySlice {
					z.MySlice[zatn], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zptq != -1 {
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
var unmarshalMsgFieldOrder1zptq = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zptq = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zvbk := range z.Slice {
		s += z.Slice[zvbk].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zrep, zgoz := range z.Transform {
			_ = zgoz
			_ = zrep
			s += msgp.IntSize
			if zgoz == nil {
				s += msgp.NilSize
			} else {
				s += zgoz.MSGPMsgsize()
			}
		}
	}
	s += 6
	if z.Myptr == nil {
		s += msgp.NilSize
	} else {
		s += z.Myptr.MSGPMsgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for zzpe := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[zzpe])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zatn := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zatn])
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
	// string "beta"
	o = append(o, 0x87, 0xa4, 0x62, 0x65, 0x74, 0x61)
	o = msgp.AppendString(o, z.B)
	// string "ralph"
	o = append(o, 0xa5, 0x72, 0x61, 0x6c, 0x70, 0x68)
	o = msgp.AppendMapHeader(o, uint32(len(z.R)))
	for ztwu, zutu := range z.R {
		o = msgp.AppendString(o, ztwu)
		o = msgp.AppendUint8(o, zutu)
	}
	// string "P"
	o = append(o, 0xa1, 0x50)
	o = msgp.AppendUint16(o, z.P)
	// string "Q"
	o = append(o, 0xa1, 0x51)
	o = msgp.AppendUint32(o, z.Q)
	// string "T"
	o = append(o, 0xa1, 0x54)
	o = msgp.AppendInt64(o, z.T)
	// string "arr"
	o = append(o, 0xa3, 0x61, 0x72, 0x72)
	o = msgp.AppendArrayHeader(o, 6)
	for zklo := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zklo])
	}
	// string "MyTree"
	o = append(o, 0xa6, 0x4d, 0x79, 0x54, 0x72, 0x65, 0x65)
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
	const maxFields2zrqf = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zrqf uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zrqf, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zrqf := totalEncodedFields2zrqf
	missingFieldsLeft2zrqf := maxFields2zrqf - totalEncodedFields2zrqf

	var nextMiss2zrqf int32 = -1
	var found2zrqf [maxFields2zrqf]bool
	var curField2zrqf string

doneWithStruct2zrqf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zrqf > 0 || missingFieldsLeft2zrqf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zrqf, missingFieldsLeft2zrqf, msgp.ShowFound(found2zrqf[:]), unmarshalMsgFieldOrder2zrqf)
		if encodedFieldsLeft2zrqf > 0 {
			encodedFieldsLeft2zrqf--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zrqf = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zrqf < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zrqf = 0
			}
			for nextMiss2zrqf < maxFields2zrqf && (found2zrqf[nextMiss2zrqf] || unmarshalMsgFieldSkip2zrqf[nextMiss2zrqf]) {
				nextMiss2zrqf++
			}
			if nextMiss2zrqf == maxFields2zrqf {
				// filled all the empty fields!
				break doneWithStruct2zrqf
			}
			missingFieldsLeft2zrqf--
			curField2zrqf = unmarshalMsgFieldOrder2zrqf[nextMiss2zrqf]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zrqf)
		switch curField2zrqf {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zrqf[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph":
			found2zrqf[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zymo uint32
				zymo, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zymo > 0 {
					z.R = make(map[string]uint8, zymo)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zymo > 0 {
					var ztwu string
					var zutu uint8
					zymo--
					ztwu, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zutu, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[ztwu] = zutu
				}
			}
		case "P":
			found2zrqf[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q":
			found2zrqf[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T":
			found2zrqf[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr":
			found2zrqf[6] = true
			var zzow uint32
			zzow, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zzow != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zzow}
				return
			}
			for zklo := range z.Arr {
				z.Arr[zklo], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree":
			found2zrqf[7] = true
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
	if nextMiss2zrqf != -1 {
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
var unmarshalMsgFieldOrder2zrqf = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zrqf = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for ztwu, zutu := range z.R {
			_ = zutu
			_ = ztwu
			s += msgp.StringPrefixSize + len(ztwu) + msgp.Uint8Size
		}
	}
	s += 2 + msgp.Uint16Size + 2 + msgp.Uint32Size + 2 + msgp.Int64Size + 4 + msgp.ArrayHeaderSize + (6 * (msgp.Float64Size)) + 7
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
	// string "F"
	o = append(o, 0x81, 0xa1, 0x46)
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
	const maxFields3zwix = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zwix uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zwix, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zwix := totalEncodedFields3zwix
	missingFieldsLeft3zwix := maxFields3zwix - totalEncodedFields3zwix

	var nextMiss3zwix int32 = -1
	var found3zwix [maxFields3zwix]bool
	var curField3zwix string

doneWithStruct3zwix:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zwix > 0 || missingFieldsLeft3zwix > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zwix, missingFieldsLeft3zwix, msgp.ShowFound(found3zwix[:]), unmarshalMsgFieldOrder3zwix)
		if encodedFieldsLeft3zwix > 0 {
			encodedFieldsLeft3zwix--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zwix = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zwix < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zwix = 0
			}
			for nextMiss3zwix < maxFields3zwix && (found3zwix[nextMiss3zwix] || unmarshalMsgFieldSkip3zwix[nextMiss3zwix]) {
				nextMiss3zwix++
			}
			if nextMiss3zwix == maxFields3zwix {
				// filled all the empty fields!
				break doneWithStruct3zwix
			}
			missingFieldsLeft3zwix--
			curField3zwix = unmarshalMsgFieldOrder3zwix[nextMiss3zwix]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zwix)
		switch curField3zwix {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zwix[0] = true
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
	if nextMiss3zwix != -1 {
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
var unmarshalMsgFieldOrder3zwix = []string{"F"}

var unmarshalMsgFieldSkip3zwix = []bool{false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sys) MSGPMsgsize() (s int) {
	s = 1 + 2 + msgp.GuessSize(z.F)
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
	// string "Chld"
	o = append(o, 0x83, 0xa4, 0x43, 0x68, 0x6c, 0x64)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Chld)))
	for zxsr := range z.Chld {
		o, err = z.Chld[zxsr].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Str"
	o = append(o, 0xa3, 0x53, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	// string "Par"
	o = append(o, 0xa3, 0x50, 0x61, 0x72)
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
	const maxFields4zmgh = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zmgh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zmgh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zmgh := totalEncodedFields4zmgh
	missingFieldsLeft4zmgh := maxFields4zmgh - totalEncodedFields4zmgh

	var nextMiss4zmgh int32 = -1
	var found4zmgh [maxFields4zmgh]bool
	var curField4zmgh string

doneWithStruct4zmgh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zmgh > 0 || missingFieldsLeft4zmgh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zmgh, missingFieldsLeft4zmgh, msgp.ShowFound(found4zmgh[:]), unmarshalMsgFieldOrder4zmgh)
		if encodedFieldsLeft4zmgh > 0 {
			encodedFieldsLeft4zmgh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zmgh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zmgh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zmgh = 0
			}
			for nextMiss4zmgh < maxFields4zmgh && (found4zmgh[nextMiss4zmgh] || unmarshalMsgFieldSkip4zmgh[nextMiss4zmgh]) {
				nextMiss4zmgh++
			}
			if nextMiss4zmgh == maxFields4zmgh {
				// filled all the empty fields!
				break doneWithStruct4zmgh
			}
			missingFieldsLeft4zmgh--
			curField4zmgh = unmarshalMsgFieldOrder4zmgh[nextMiss4zmgh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zmgh)
		switch curField4zmgh {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zmgh[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zanv uint32
				zanv, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zanv) {
					z.Chld = (z.Chld)[:zanv]
				} else {
					z.Chld = make([]Tree, zanv)
				}
				for zxsr := range z.Chld {
					bts, err = z.Chld[zxsr].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str":
			found4zmgh[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par":
			found4zmgh[2] = true
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
	if nextMiss4zmgh != -1 {
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
var unmarshalMsgFieldOrder4zmgh = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zmgh = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zxsr := range z.Chld {
		s += z.Chld[zxsr].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
