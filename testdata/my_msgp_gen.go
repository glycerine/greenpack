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
	const maxFields0zsza = 6

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields0zsza uint32
	if !nbs.AlwaysNil {
		totalEncodedFields0zsza, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft0zsza := totalEncodedFields0zsza
	missingFieldsLeft0zsza := maxFields0zsza - totalEncodedFields0zsza

	var nextMiss0zsza int32 = -1
	var found0zsza [maxFields0zsza]bool
	var curField0zsza string

doneWithStruct0zsza:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zsza > 0 || missingFieldsLeft0zsza > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zsza, missingFieldsLeft0zsza, msgp.ShowFound(found0zsza[:]), unmarshalMsgFieldOrder0zsza)
		if encodedFieldsLeft0zsza > 0 {
			encodedFieldsLeft0zsza--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField0zsza = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss0zsza < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss0zsza = 0
			}
			for nextMiss0zsza < maxFields0zsza && (found0zsza[nextMiss0zsza] || unmarshalMsgFieldSkip0zsza[nextMiss0zsza]) {
				nextMiss0zsza++
			}
			if nextMiss0zsza == maxFields0zsza {
				// filled all the empty fields!
				break doneWithStruct0zsza
			}
			missingFieldsLeft0zsza--
			curField0zsza = unmarshalMsgFieldOrder0zsza[nextMiss0zsza]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zsza)
		switch curField0zsza {
		// -- templateUnmarshalMsg ends here --

		case "name":
			found0zsza[0] = true
			z.Name, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Bday":
			found0zsza[1] = true
			z.Bday, bts, err = nbs.ReadTimeBytes(bts)

			if err != nil {
				return
			}
		case "phone":
			found0zsza[2] = true
			z.Phone, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Sibs":
			found0zsza[3] = true
			z.Sibs, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case "GPA":
			found0zsza[4] = true
			z.GPA, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case "Friend":
			found0zsza[5] = true
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
	if nextMiss0zsza != -1 {
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
var unmarshalMsgFieldOrder0zsza = []string{"name", "Bday", "phone", "Sibs", "GPA", "Friend"}

var unmarshalMsgFieldSkip0zsza = []bool{false, false, false, false, false, false}

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
	for zmlr := range z.Slice {
		o, err = z.Slice[zmlr].MSGPMarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Transform"
	o = append(o, 0xa9, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendMapHeader(o, uint32(len(z.Transform)))
	for zewr, zrif := range z.Transform {
		o = msgp.AppendInt(o, zewr)
		if zrif == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = zrif.MSGPMarshalMsg(o)
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
	for ztag := range z.Myarray {
		o = msgp.AppendString(o, z.Myarray[ztag])
	}
	// string "MySlice"
	o = append(o, 0xa7, 0x4d, 0x79, 0x53, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MySlice)))
	for zgpw := range z.MySlice {
		o = msgp.AppendString(o, z.MySlice[zgpw])
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
	const maxFields1zgfy = 5

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields1zgfy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zgfy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zgfy := totalEncodedFields1zgfy
	missingFieldsLeft1zgfy := maxFields1zgfy - totalEncodedFields1zgfy

	var nextMiss1zgfy int32 = -1
	var found1zgfy [maxFields1zgfy]bool
	var curField1zgfy string

doneWithStruct1zgfy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgfy > 0 || missingFieldsLeft1zgfy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgfy, missingFieldsLeft1zgfy, msgp.ShowFound(found1zgfy[:]), unmarshalMsgFieldOrder1zgfy)
		if encodedFieldsLeft1zgfy > 0 {
			encodedFieldsLeft1zgfy--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField1zgfy = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss1zgfy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zgfy = 0
			}
			for nextMiss1zgfy < maxFields1zgfy && (found1zgfy[nextMiss1zgfy] || unmarshalMsgFieldSkip1zgfy[nextMiss1zgfy]) {
				nextMiss1zgfy++
			}
			if nextMiss1zgfy == maxFields1zgfy {
				// filled all the empty fields!
				break doneWithStruct1zgfy
			}
			missingFieldsLeft1zgfy--
			curField1zgfy = unmarshalMsgFieldOrder1zgfy[nextMiss1zgfy]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgfy)
		switch curField1zgfy {
		// -- templateUnmarshalMsg ends here --

		case "Slice":
			found1zgfy[0] = true
			if nbs.AlwaysNil {
				(z.Slice) = (z.Slice)[:0]
			} else {

				var zzhd uint32
				zzhd, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Slice) >= int(zzhd) {
					z.Slice = (z.Slice)[:zzhd]
				} else {
					z.Slice = make([]S2, zzhd)
				}
				for zmlr := range z.Slice {
					bts, err = z.Slice[zmlr].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Transform":
			found1zgfy[1] = true
			if nbs.AlwaysNil {
				if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}

			} else {

				var zfvw uint32
				zfvw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Transform == nil && zfvw > 0 {
					z.Transform = make(map[int]*S2, zfvw)
				} else if len(z.Transform) > 0 {
					for key, _ := range z.Transform {
						delete(z.Transform, key)
					}
				}
				for zfvw > 0 {
					var zewr int
					var zrif *S2
					zfvw--
					zewr, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zrif != nil {
							zrif.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zrif {
								zrif.MSGPUnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zrif == nil {
								zrif = new(S2)
							}
							bts, err = zrif.MSGPUnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Transform[zewr] = zrif
				}
			}
		case "Myptr":
			found1zgfy[2] = true
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
			found1zgfy[3] = true
			var zmwk uint32
			zmwk, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zmwk != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: zmwk}
				return
			}
			for ztag := range z.Myarray {
				z.Myarray[ztag], bts, err = nbs.ReadStringBytes(bts)

				if err != nil {
					return
				}
			}
		case "MySlice":
			found1zgfy[4] = true
			if nbs.AlwaysNil {
				(z.MySlice) = (z.MySlice)[:0]
			} else {

				var zddq uint32
				zddq, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.MySlice) >= int(zddq) {
					z.MySlice = (z.MySlice)[:zddq]
				} else {
					z.MySlice = make([]string, zddq)
				}
				for zgpw := range z.MySlice {
					z.MySlice[zgpw], bts, err = nbs.ReadStringBytes(bts)

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
	if nextMiss1zgfy != -1 {
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
var unmarshalMsgFieldOrder1zgfy = []string{"Slice", "Transform", "Myptr", "Myarray", "MySlice"}

var unmarshalMsgFieldSkip1zgfy = []bool{false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Big) MSGPMsgsize() (s int) {
	s = 1 + 6 + msgp.ArrayHeaderSize
	for zmlr := range z.Slice {
		s += z.Slice[zmlr].MSGPMsgsize()
	}
	s += 10 + msgp.MapHeaderSize
	if z.Transform != nil {
		for zewr, zrif := range z.Transform {
			_ = zrif
			_ = zewr
			s += msgp.IntSize
			if zrif == nil {
				s += msgp.NilSize
			} else {
				s += zrif.MSGPMsgsize()
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
	for ztag := range z.Myarray {
		s += msgp.StringPrefixSize + len(z.Myarray[ztag])
	}
	s += 8 + msgp.ArrayHeaderSize
	for zgpw := range z.MySlice {
		s += msgp.StringPrefixSize + len(z.MySlice[zgpw])
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
	for zjpd, zrse := range z.R {
		o = msgp.AppendString(o, zjpd)
		o = msgp.AppendUint8(o, zrse)
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
	for zuuy := range z.Arr {
		o = msgp.AppendFloat64(o, z.Arr[zuuy])
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
	const maxFields2zonu = 8

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields2zonu uint32
	if !nbs.AlwaysNil {
		totalEncodedFields2zonu, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft2zonu := totalEncodedFields2zonu
	missingFieldsLeft2zonu := maxFields2zonu - totalEncodedFields2zonu

	var nextMiss2zonu int32 = -1
	var found2zonu [maxFields2zonu]bool
	var curField2zonu string

doneWithStruct2zonu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zonu > 0 || missingFieldsLeft2zonu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zonu, missingFieldsLeft2zonu, msgp.ShowFound(found2zonu[:]), unmarshalMsgFieldOrder2zonu)
		if encodedFieldsLeft2zonu > 0 {
			encodedFieldsLeft2zonu--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField2zonu = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss2zonu < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss2zonu = 0
			}
			for nextMiss2zonu < maxFields2zonu && (found2zonu[nextMiss2zonu] || unmarshalMsgFieldSkip2zonu[nextMiss2zonu]) {
				nextMiss2zonu++
			}
			if nextMiss2zonu == maxFields2zonu {
				// filled all the empty fields!
				break doneWithStruct2zonu
			}
			missingFieldsLeft2zonu--
			curField2zonu = unmarshalMsgFieldOrder2zonu[nextMiss2zonu]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zonu)
		switch curField2zonu {
		// -- templateUnmarshalMsg ends here --

		case "beta":
			found2zonu[1] = true
			z.B, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "ralph":
			found2zonu[2] = true
			if nbs.AlwaysNil {
				if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}

			} else {

				var zziw uint32
				zziw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.R == nil && zziw > 0 {
					z.R = make(map[string]uint8, zziw)
				} else if len(z.R) > 0 {
					for key, _ := range z.R {
						delete(z.R, key)
					}
				}
				for zziw > 0 {
					var zjpd string
					var zrse uint8
					zziw--
					zjpd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zrse, bts, err = nbs.ReadUint8Bytes(bts)

					if err != nil {
						return
					}
					z.R[zjpd] = zrse
				}
			}
		case "P":
			found2zonu[3] = true
			z.P, bts, err = nbs.ReadUint16Bytes(bts)

			if err != nil {
				return
			}
		case "Q":
			found2zonu[4] = true
			z.Q, bts, err = nbs.ReadUint32Bytes(bts)

			if err != nil {
				return
			}
		case "T":
			found2zonu[5] = true
			z.T, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case "arr":
			found2zonu[6] = true
			var zgtx uint32
			zgtx, bts, err = nbs.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if !nbs.IsNil(bts) && zgtx != 6 {
				err = msgp.ArrayError{Wanted: 6, Got: zgtx}
				return
			}
			for zuuy := range z.Arr {
				z.Arr[zuuy], bts, err = nbs.ReadFloat64Bytes(bts)

				if err != nil {
					return
				}
			}
		case "MyTree":
			found2zonu[7] = true
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
	if nextMiss2zonu != -1 {
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
var unmarshalMsgFieldOrder2zonu = []string{"alpha", "beta", "ralph", "P", "Q", "T", "arr", "MyTree"}

var unmarshalMsgFieldSkip2zonu = []bool{true, false, false, false, false, false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *S2) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.B) + 6 + msgp.MapHeaderSize
	if z.R != nil {
		for zjpd, zrse := range z.R {
			_ = zrse
			_ = zjpd
			s += msgp.StringPrefixSize + len(zjpd) + msgp.Uint8Size
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
	const maxFields3zlye = 1

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields3zlye uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zlye, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zlye := totalEncodedFields3zlye
	missingFieldsLeft3zlye := maxFields3zlye - totalEncodedFields3zlye

	var nextMiss3zlye int32 = -1
	var found3zlye [maxFields3zlye]bool
	var curField3zlye string

doneWithStruct3zlye:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zlye > 0 || missingFieldsLeft3zlye > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zlye, missingFieldsLeft3zlye, msgp.ShowFound(found3zlye[:]), unmarshalMsgFieldOrder3zlye)
		if encodedFieldsLeft3zlye > 0 {
			encodedFieldsLeft3zlye--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField3zlye = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss3zlye < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zlye = 0
			}
			for nextMiss3zlye < maxFields3zlye && (found3zlye[nextMiss3zlye] || unmarshalMsgFieldSkip3zlye[nextMiss3zlye]) {
				nextMiss3zlye++
			}
			if nextMiss3zlye == maxFields3zlye {
				// filled all the empty fields!
				break doneWithStruct3zlye
			}
			missingFieldsLeft3zlye--
			curField3zlye = unmarshalMsgFieldOrder3zlye[nextMiss3zlye]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zlye)
		switch curField3zlye {
		// -- templateUnmarshalMsg ends here --

		case "F":
			found3zlye[0] = true
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
	if nextMiss3zlye != -1 {
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
var unmarshalMsgFieldOrder3zlye = []string{"F"}

var unmarshalMsgFieldSkip3zlye = []bool{false}

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
	for zpxz := range z.Chld {
		o, err = z.Chld[zpxz].MSGPMarshalMsg(o)
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
	const maxFields4zsuh = 3

	// -- templateUnmarshalMsg starts here--
	var totalEncodedFields4zsuh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields4zsuh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft4zsuh := totalEncodedFields4zsuh
	missingFieldsLeft4zsuh := maxFields4zsuh - totalEncodedFields4zsuh

	var nextMiss4zsuh int32 = -1
	var found4zsuh [maxFields4zsuh]bool
	var curField4zsuh string

doneWithStruct4zsuh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zsuh > 0 || missingFieldsLeft4zsuh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zsuh, missingFieldsLeft4zsuh, msgp.ShowFound(found4zsuh[:]), unmarshalMsgFieldOrder4zsuh)
		if encodedFieldsLeft4zsuh > 0 {
			encodedFieldsLeft4zsuh--
			field, bts, err = nbs.ReadMapKeyZC(bts)
			if err != nil {
				return
			}
			curField4zsuh = msgp.UnsafeString(field)
		} else {
			//missing fields need handling
			if nextMiss4zsuh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss4zsuh = 0
			}
			for nextMiss4zsuh < maxFields4zsuh && (found4zsuh[nextMiss4zsuh] || unmarshalMsgFieldSkip4zsuh[nextMiss4zsuh]) {
				nextMiss4zsuh++
			}
			if nextMiss4zsuh == maxFields4zsuh {
				// filled all the empty fields!
				break doneWithStruct4zsuh
			}
			missingFieldsLeft4zsuh--
			curField4zsuh = unmarshalMsgFieldOrder4zsuh[nextMiss4zsuh]
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zsuh)
		switch curField4zsuh {
		// -- templateUnmarshalMsg ends here --

		case "Chld":
			found4zsuh[0] = true
			if nbs.AlwaysNil {
				(z.Chld) = (z.Chld)[:0]
			} else {

				var zbmz uint32
				zbmz, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Chld) >= int(zbmz) {
					z.Chld = (z.Chld)[:zbmz]
				} else {
					z.Chld = make([]Tree, zbmz)
				}
				for zpxz := range z.Chld {
					bts, err = z.Chld[zpxz].MSGPUnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
				}
			}
		case "Str":
			found4zsuh[1] = true
			z.Str, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case "Par":
			found4zsuh[2] = true
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
	if nextMiss4zsuh != -1 {
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
var unmarshalMsgFieldOrder4zsuh = []string{"Chld", "Str", "Par"}

var unmarshalMsgFieldSkip4zsuh = []bool{false, false, false}

// MSGPMsgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tree) MSGPMsgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for zpxz := range z.Chld {
		s += z.Chld[zpxz].MSGPMsgsize()
	}
	s += 4 + msgp.StringPrefixSize + len(z.Str) + 4
	if z.Par == nil {
		s += msgp.NilSize
	} else {
		s += z.Par.MSGPMsgsize()
	}
	return
}
