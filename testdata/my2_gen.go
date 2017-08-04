package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
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
	const maxFields0zbvl = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zbvl uint32
	totalEncodedFields0zbvl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zbvl := totalEncodedFields0zbvl
	missingFieldsLeft0zbvl := maxFields0zbvl - totalEncodedFields0zbvl

	var nextMiss0zbvl int = -1
	var found0zbvl [maxFields0zbvl]bool
	var curField0zbvl int

doneWithStruct0zbvl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zbvl > 0 || missingFieldsLeft0zbvl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zbvl, missingFieldsLeft0zbvl, msgp.ShowFound(found0zbvl[:]), decodeMsgFieldOrder0zbvl)
		if encodedFieldsLeft0zbvl > 0 {
			encodedFieldsLeft0zbvl--
			curField0zbvl, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zbvl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zbvl = 0
			}
			for nextMiss0zbvl < maxFields0zbvl && (found0zbvl[nextMiss0zbvl] || decodeMsgFieldSkip0zbvl[nextMiss0zbvl]) {
				nextMiss0zbvl++
			}
			if nextMiss0zbvl == maxFields0zbvl {
				// filled all the empty fields!
				break doneWithStruct0zbvl
			}
			missingFieldsLeft0zbvl--
			curField0zbvl = nextMiss0zbvl
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zbvl)
		switch curField0zbvl {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found0zbvl[0] = true
			z.U, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found0zbvl[2] = true
			z.Nt, err = dc.ReadInt()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found0zbvl[3] = true
			var zoad uint32
			zoad, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Snot == nil && zoad > 0 {
				z.Snot = make(map[string]bool, zoad)
			} else if len(z.Snot) > 0 {
				for key, _ := range z.Snot {
					delete(z.Snot, key)
				}
			}
			for zoad > 0 {
				zoad--
				var zfhk string
				var zoku bool
				zfhk, err = dc.ReadString()
				if err != nil {
					return
				}
				zoku, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Snot[zfhk] = zoku
			}
		case 3:
			// zid 3 for "Notyet"
			found0zbvl[4] = true
			var zunc uint32
			zunc, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Notyet == nil && zunc > 0 {
				z.Notyet = make(map[string]bool, zunc)
			} else if len(z.Notyet) > 0 {
				for key, _ := range z.Notyet {
					delete(z.Notyet, key)
				}
			}
			for zunc > 0 {
				zunc--
				var zbtd string
				var zytj bool
				zbtd, err = dc.ReadString()
				if err != nil {
					return
				}
				zytj, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.Notyet[zbtd] = zytj
			}
		case 4:
			// zid 4 for "Setm"
			found0zbvl[5] = true
			var zfcv uint32
			zfcv, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Setm) >= int(zfcv) {
				z.Setm = (z.Setm)[:zfcv]
			} else {
				z.Setm = make([]*inn, zfcv)
			}
			for zmpe := range z.Setm {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if z.Setm[zmpe] != nil {
						dc.PushAlwaysNil()
						err = z.Setm[zmpe].DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if z.Setm[zmpe] == nil {
						z.Setm[zmpe] = new(inn)
					}
					err = z.Setm[zmpe].DecodeMsg(dc)
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
	if nextMiss0zbvl != -1 {
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
var decodeMsgFieldOrder0zbvl = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var decodeMsgFieldSkip0zbvl = []bool{false, true, false, false, false, false}

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
	var empty_zfqf [6]bool
	fieldsInUse_zyxh := z.fieldsNotEmpty(empty_zfqf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zyxh)
	if err != nil {
		return err
	}

	if !empty_zfqf[0] {
		// zid 0 for "U"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteString(z.U)
		if err != nil {
			return
		}
	}

	if !empty_zfqf[2] {
		// zid 1 for "Nt"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Nt)
		if err != nil {
			return
		}
	}

	if !empty_zfqf[3] {
		// zid 2 for "Snot"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Snot)))
		if err != nil {
			return
		}
		for zfhk, zoku := range z.Snot {
			err = en.WriteString(zfhk)
			if err != nil {
				return
			}
			err = en.WriteBool(zoku)
			if err != nil {
				return
			}
		}
	}

	if !empty_zfqf[4] {
		// zid 3 for "Notyet"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Notyet)))
		if err != nil {
			return
		}
		for zbtd, zytj := range z.Notyet {
			err = en.WriteString(zbtd)
			if err != nil {
				return
			}
			err = en.WriteBool(zytj)
			if err != nil {
				return
			}
		}
	}

	if !empty_zfqf[5] {
		// zid 4 for "Setm"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteArrayHeader(uint32(len(z.Setm)))
		if err != nil {
			return
		}
		for zmpe := range z.Setm {
			if z.Setm[zmpe] == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = z.Setm[zmpe].EncodeMsg(en)
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
		// zid 0 for "U"
		o = append(o, 0x0)
		o = msgp.AppendString(o, z.U)
	}

	if !empty[2] {
		// zid 1 for "Nt"
		o = append(o, 0x1)
		o = msgp.AppendInt(o, z.Nt)
	}

	if !empty[3] {
		// zid 2 for "Snot"
		o = append(o, 0x2)
		o = msgp.AppendMapHeader(o, uint32(len(z.Snot)))
		for zfhk, zoku := range z.Snot {
			o = msgp.AppendString(o, zfhk)
			o = msgp.AppendBool(o, zoku)
		}
	}

	if !empty[4] {
		// zid 3 for "Notyet"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.Notyet)))
		for zbtd, zytj := range z.Notyet {
			o = msgp.AppendString(o, zbtd)
			o = msgp.AppendBool(o, zytj)
		}
	}

	if !empty[5] {
		// zid 4 for "Setm"
		o = append(o, 0x4)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Setm)))
		for zmpe := range z.Setm {
			if z.Setm[zmpe] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Setm[zmpe].MarshalMsg(o)
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
	const maxFields1zlex = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zlex uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zlex, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zlex := totalEncodedFields1zlex
	missingFieldsLeft1zlex := maxFields1zlex - totalEncodedFields1zlex

	var nextMiss1zlex int = -1
	var found1zlex [maxFields1zlex]bool
	var curField1zlex int

doneWithStruct1zlex:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zlex > 0 || missingFieldsLeft1zlex > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zlex, missingFieldsLeft1zlex, msgp.ShowFound(found1zlex[:]), unmarshalMsgFieldOrder1zlex)
		if encodedFieldsLeft1zlex > 0 {
			encodedFieldsLeft1zlex--
			curField1zlex, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zlex < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zlex = 0
			}
			for nextMiss1zlex < maxFields1zlex && (found1zlex[nextMiss1zlex] || unmarshalMsgFieldSkip1zlex[nextMiss1zlex]) {
				nextMiss1zlex++
			}
			if nextMiss1zlex == maxFields1zlex {
				// filled all the empty fields!
				break doneWithStruct1zlex
			}
			missingFieldsLeft1zlex--
			curField1zlex = nextMiss1zlex
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zlex)
		switch curField1zlex {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "U"
			found1zlex[0] = true
			z.U, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Nt"
			found1zlex[2] = true
			z.Nt, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Snot"
			found1zlex[3] = true
			if nbs.AlwaysNil {
				if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}

			} else {

				var zrkz uint32
				zrkz, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Snot == nil && zrkz > 0 {
					z.Snot = make(map[string]bool, zrkz)
				} else if len(z.Snot) > 0 {
					for key, _ := range z.Snot {
						delete(z.Snot, key)
					}
				}
				for zrkz > 0 {
					var zfhk string
					var zoku bool
					zrkz--
					zfhk, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zoku, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Snot[zfhk] = zoku
				}
			}
		case 3:
			// zid 3 for "Notyet"
			found1zlex[4] = true
			if nbs.AlwaysNil {
				if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}

			} else {

				var zibp uint32
				zibp, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Notyet == nil && zibp > 0 {
					z.Notyet = make(map[string]bool, zibp)
				} else if len(z.Notyet) > 0 {
					for key, _ := range z.Notyet {
						delete(z.Notyet, key)
					}
				}
				for zibp > 0 {
					var zbtd string
					var zytj bool
					zibp--
					zbtd, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zytj, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.Notyet[zbtd] = zytj
				}
			}
		case 4:
			// zid 4 for "Setm"
			found1zlex[5] = true
			if nbs.AlwaysNil {
				(z.Setm) = (z.Setm)[:0]
			} else {

				var zzoj uint32
				zzoj, bts, err = nbs.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.Setm) >= int(zzoj) {
					z.Setm = (z.Setm)[:zzoj]
				} else {
					z.Setm = make([]*inn, zzoj)
				}
				for zmpe := range z.Setm {
					if nbs.AlwaysNil {
						if z.Setm[zmpe] != nil {
							z.Setm[zmpe].UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != z.Setm[zmpe] {
								z.Setm[zmpe].UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if z.Setm[zmpe] == nil {
								z.Setm[zmpe] = new(inn)
							}
							bts, err = z.Setm[zmpe].UnmarshalMsg(bts)
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
	if nextMiss1zlex != -1 {
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
var unmarshalMsgFieldOrder1zlex = []string{"U", "-", "Nt", "Snot", "Notyet", "Setm"}

var unmarshalMsgFieldSkip1zlex = []bool{false, true, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Tr) Msgsize() (s int) {
	s = 1 + 2 + msgp.StringPrefixSize + len(z.U) + 3 + msgp.IntSize + 5 + msgp.MapHeaderSize
	if z.Snot != nil {
		for zfhk, zoku := range z.Snot {
			_ = zoku
			_ = zfhk
			s += msgp.StringPrefixSize + len(zfhk) + msgp.BoolSize
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Notyet != nil {
		for zbtd, zytj := range z.Notyet {
			_ = zytj
			_ = zbtd
			s += msgp.StringPrefixSize + len(zbtd) + msgp.BoolSize
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for zmpe := range z.Setm {
		if z.Setm[zmpe] == nil {
			s += msgp.NilSize
		} else {
			s += z.Setm[zmpe].Msgsize()
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
	const maxFields2znzs = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2znzs uint32
	totalEncodedFields2znzs, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2znzs := totalEncodedFields2znzs
	missingFieldsLeft2znzs := maxFields2znzs - totalEncodedFields2znzs

	var nextMiss2znzs int = -1
	var found2znzs [maxFields2znzs]bool
	var curField2znzs int

doneWithStruct2znzs:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2znzs > 0 || missingFieldsLeft2znzs > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2znzs, missingFieldsLeft2znzs, msgp.ShowFound(found2znzs[:]), decodeMsgFieldOrder2znzs)
		if encodedFieldsLeft2znzs > 0 {
			encodedFieldsLeft2znzs--
			curField2znzs, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2znzs < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2znzs = 0
			}
			for nextMiss2znzs < maxFields2znzs && (found2znzs[nextMiss2znzs] || decodeMsgFieldSkip2znzs[nextMiss2znzs]) {
				nextMiss2znzs++
			}
			if nextMiss2znzs == maxFields2znzs {
				// filled all the empty fields!
				break doneWithStruct2znzs
			}
			missingFieldsLeft2znzs--
			curField2znzs = nextMiss2znzs
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2znzs)
		switch curField2znzs {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found2znzs[0] = true
			var zvcw uint32
			zvcw, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.j == nil && zvcw > 0 {
				z.j = make(map[string]bool, zvcw)
			} else if len(z.j) > 0 {
				for key, _ := range z.j {
					delete(z.j, key)
				}
			}
			for zvcw > 0 {
				zvcw--
				var zafj string
				var zbes bool
				zafj, err = dc.ReadString()
				if err != nil {
					return
				}
				zbes, err = dc.ReadBool()
				if err != nil {
					return
				}
				z.j[zafj] = zbes
			}
		case 1:
			// zid 1 for "e"
			found2znzs[1] = true
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
	if nextMiss2znzs != -1 {
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
var decodeMsgFieldOrder2znzs = []string{"j", "e"}

var decodeMsgFieldSkip2znzs = []bool{false, false}

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
	var empty_zjvi [2]bool
	fieldsInUse_zksy := z.fieldsNotEmpty(empty_zjvi[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zksy)
	if err != nil {
		return err
	}

	if !empty_zjvi[0] {
		// zid 0 for "j"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.j)))
		if err != nil {
			return
		}
		for zafj, zbes := range z.j {
			err = en.WriteString(zafj)
			if err != nil {
				return
			}
			err = en.WriteBool(zbes)
			if err != nil {
				return
			}
		}
	}

	if !empty_zjvi[1] {
		// zid 1 for "e"
		err = en.Append(0x1)
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
		// zid 0 for "j"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.j)))
		for zafj, zbes := range z.j {
			o = msgp.AppendString(o, zafj)
			o = msgp.AppendBool(o, zbes)
		}
	}

	if !empty[1] {
		// zid 1 for "e"
		o = append(o, 0x1)
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
	const maxFields3zagv = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zagv uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zagv, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zagv := totalEncodedFields3zagv
	missingFieldsLeft3zagv := maxFields3zagv - totalEncodedFields3zagv

	var nextMiss3zagv int = -1
	var found3zagv [maxFields3zagv]bool
	var curField3zagv int

doneWithStruct3zagv:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zagv > 0 || missingFieldsLeft3zagv > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zagv, missingFieldsLeft3zagv, msgp.ShowFound(found3zagv[:]), unmarshalMsgFieldOrder3zagv)
		if encodedFieldsLeft3zagv > 0 {
			encodedFieldsLeft3zagv--
			curField3zagv, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zagv < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zagv = 0
			}
			for nextMiss3zagv < maxFields3zagv && (found3zagv[nextMiss3zagv] || unmarshalMsgFieldSkip3zagv[nextMiss3zagv]) {
				nextMiss3zagv++
			}
			if nextMiss3zagv == maxFields3zagv {
				// filled all the empty fields!
				break doneWithStruct3zagv
			}
			missingFieldsLeft3zagv--
			curField3zagv = nextMiss3zagv
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zagv)
		switch curField3zagv {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "j"
			found3zagv[0] = true
			if nbs.AlwaysNil {
				if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}

			} else {

				var ztxa uint32
				ztxa, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.j == nil && ztxa > 0 {
					z.j = make(map[string]bool, ztxa)
				} else if len(z.j) > 0 {
					for key, _ := range z.j {
						delete(z.j, key)
					}
				}
				for ztxa > 0 {
					var zafj string
					var zbes bool
					ztxa--
					zafj, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					zbes, bts, err = nbs.ReadBoolBytes(bts)

					if err != nil {
						return
					}
					z.j[zafj] = zbes
				}
			}
		case 1:
			// zid 1 for "e"
			found3zagv[1] = true
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
	if nextMiss3zagv != -1 {
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
var unmarshalMsgFieldOrder3zagv = []string{"j", "e"}

var unmarshalMsgFieldSkip3zagv = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *inn) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.j != nil {
		for zafj, zbes := range z.j {
			_ = zbes
			_ = zafj
			s += msgp.StringPrefixSize + len(zafj) + msgp.BoolSize
		}
	}
	s += 2 + msgp.Int64Size
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
	const maxFields4zuiq = 3

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zuiq uint32
	totalEncodedFields4zuiq, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zuiq := totalEncodedFields4zuiq
	missingFieldsLeft4zuiq := maxFields4zuiq - totalEncodedFields4zuiq

	var nextMiss4zuiq int = -1
	var found4zuiq [maxFields4zuiq]bool
	var curField4zuiq int

doneWithStruct4zuiq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zuiq > 0 || missingFieldsLeft4zuiq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zuiq, missingFieldsLeft4zuiq, msgp.ShowFound(found4zuiq[:]), decodeMsgFieldOrder4zuiq)
		if encodedFieldsLeft4zuiq > 0 {
			encodedFieldsLeft4zuiq--
			curField4zuiq, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zuiq < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zuiq = 0
			}
			for nextMiss4zuiq < maxFields4zuiq && (found4zuiq[nextMiss4zuiq] || decodeMsgFieldSkip4zuiq[nextMiss4zuiq]) {
				nextMiss4zuiq++
			}
			if nextMiss4zuiq == maxFields4zuiq {
				// filled all the empty fields!
				break doneWithStruct4zuiq
			}
			missingFieldsLeft4zuiq--
			curField4zuiq = nextMiss4zuiq
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zuiq)
		switch curField4zuiq {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found4zuiq[0] = true
			var zkjp uint32
			zkjp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.m == nil && zkjp > 0 {
				z.m = make(map[string]*Tr, zkjp)
			} else if len(z.m) > 0 {
				for key, _ := range z.m {
					delete(z.m, key)
				}
			}
			for zkjp > 0 {
				zkjp--
				var zmxe string
				var zvyl *Tr
				zmxe, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zvyl != nil {
						dc.PushAlwaysNil()
						err = zvyl.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zvyl == nil {
						zvyl = new(Tr)
					}
					err = zvyl.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.m[zmxe] = zvyl
			}
		case 1:
			// zid 1 for "s"
			found4zuiq[1] = true
			z.s, err = dc.ReadString()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found4zuiq[2] = true
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
	if nextMiss4zuiq != -1 {
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
var decodeMsgFieldOrder4zuiq = []string{"m", "s", "n"}

var decodeMsgFieldSkip4zuiq = []bool{false, false, false}

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
	var empty_zldm [3]bool
	fieldsInUse_zhei := z.fieldsNotEmpty(empty_zldm[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zhei)
	if err != nil {
		return err
	}

	if !empty_zldm[0] {
		// zid 0 for "m"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.m)))
		if err != nil {
			return
		}
		for zmxe, zvyl := range z.m {
			err = en.WriteString(zmxe)
			if err != nil {
				return
			}
			if zvyl == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zvyl.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zldm[1] {
		// zid 1 for "s"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.s)
		if err != nil {
			return
		}
	}

	if !empty_zldm[2] {
		// zid 2 for "n"
		err = en.Append(0x2)
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
		// zid 0 for "m"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.m)))
		for zmxe, zvyl := range z.m {
			o = msgp.AppendString(o, zmxe)
			if zvyl == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zvyl.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[1] {
		// zid 1 for "s"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.s)
	}

	if !empty[2] {
		// zid 2 for "n"
		o = append(o, 0x2)
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
	const maxFields5zbbe = 3

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zbbe uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zbbe, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zbbe := totalEncodedFields5zbbe
	missingFieldsLeft5zbbe := maxFields5zbbe - totalEncodedFields5zbbe

	var nextMiss5zbbe int = -1
	var found5zbbe [maxFields5zbbe]bool
	var curField5zbbe int

doneWithStruct5zbbe:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zbbe > 0 || missingFieldsLeft5zbbe > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zbbe, missingFieldsLeft5zbbe, msgp.ShowFound(found5zbbe[:]), unmarshalMsgFieldOrder5zbbe)
		if encodedFieldsLeft5zbbe > 0 {
			encodedFieldsLeft5zbbe--
			curField5zbbe, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zbbe < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zbbe = 0
			}
			for nextMiss5zbbe < maxFields5zbbe && (found5zbbe[nextMiss5zbbe] || unmarshalMsgFieldSkip5zbbe[nextMiss5zbbe]) {
				nextMiss5zbbe++
			}
			if nextMiss5zbbe == maxFields5zbbe {
				// filled all the empty fields!
				break doneWithStruct5zbbe
			}
			missingFieldsLeft5zbbe--
			curField5zbbe = nextMiss5zbbe
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zbbe)
		switch curField5zbbe {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "m"
			found5zbbe[0] = true
			if nbs.AlwaysNil {
				if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}

			} else {

				var zzat uint32
				zzat, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.m == nil && zzat > 0 {
					z.m = make(map[string]*Tr, zzat)
				} else if len(z.m) > 0 {
					for key, _ := range z.m {
						delete(z.m, key)
					}
				}
				for zzat > 0 {
					var zmxe string
					var zvyl *Tr
					zzat--
					zmxe, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zvyl != nil {
							zvyl.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zvyl {
								zvyl.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zvyl == nil {
								zvyl = new(Tr)
							}
							bts, err = zvyl.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.m[zmxe] = zvyl
				}
			}
		case 1:
			// zid 1 for "s"
			found5zbbe[1] = true
			z.s, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "n"
			found5zbbe[2] = true
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
	if nextMiss5zbbe != -1 {
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
var unmarshalMsgFieldOrder5zbbe = []string{"m", "s", "n"}

var unmarshalMsgFieldSkip5zbbe = []bool{false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *u) Msgsize() (s int) {
	s = 1 + 2 + msgp.MapHeaderSize
	if z.m != nil {
		for zmxe, zvyl := range z.m {
			_ = zvyl
			_ = zmxe
			s += msgp.StringPrefixSize + len(zmxe)
			if zvyl == nil {
				s += msgp.NilSize
			} else {
				s += zvyl.Msgsize()
			}
		}
	}
	s += 2 + msgp.StringPrefixSize + len(z.s) + 2 + msgp.Int64Size
	return
}

// FileMy2 holds ZebraPack schema from file 'my2.go'
type FileMy2 struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 1636 bytes
func (FileMy2) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xa6, 0x6d, 0x79, 0x32, 0x2e, 0x67, 0x6f, 0xad,
		0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b,
		0x61, 0x67, 0x65, 0xa8, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
		0x74, 0x61, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x00, 0xa7, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x73, 0x83, 0xa3, 0x69, 0x6e, 0x6e,
		0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0xa3, 0x69, 0x6e, 0x6e, 0xa6, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x73, 0x92, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0xa1, 0x6a, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6a, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d,
		0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61,
		0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12,
		0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x87,
		0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x65, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa1, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xa1, 0x75, 0x82, 0xaa, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x75, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x93, 0x86,
		0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6d, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa1, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xae, 0x6d, 0x61,
		0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a,
		0x54, 0x72, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53,
		0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02,
		0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50,
		0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d,
		0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16,
		0xa3, 0x53, 0x74, 0x72, 0xa2, 0x54, 0x72, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x73, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa1, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x6e,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0xa2, 0x54, 0x72, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x54, 0x72, 0xa6,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa1, 0x55, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa1, 0x55, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x85, 0xa3, 0x5a, 0x69, 0x64,
		0xff, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa2, 0x65, 0x74, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa1,
		0x2d, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x1c, 0xa4, 0x53, 0x6b, 0x69,
		0x70, 0xc3, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa2, 0x4e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa2, 0x4e, 0x74, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0d, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0d, 0xa3, 0x53,
		0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa4, 0x53, 0x6e, 0x6f, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xaf, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0x5d, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53,
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa3, 0x5a,
		0x69, 0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74, 0x79,
		0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x4e, 0x6f, 0x74, 0x79,
		0x65, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xaf, 0x6d, 0x61, 0x70, 0x5b,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x62, 0x6f, 0x6f,
		0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74,
		0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61,
		0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4, 0x62, 0x6f,
		0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa4, 0x53, 0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x53,
		0x65, 0x74, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x5b, 0x5d, 0x2a,
		0x69, 0x6e, 0x6e, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1a, 0xa3,
		0x53, 0x74, 0x72, 0xa5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0xa6,
		0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x83, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f,
		0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61,
		0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x6e, 0xa7, 0x49, 0x6d,
		0x70, 0x6f, 0x72, 0x74, 0x73, 0x91, 0xbd, 0x22, 0x67, 0x69,
		0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
		0x6c, 0x79, 0x63, 0x65, 0x72, 0x69, 0x6e, 0x65, 0x2f, 0x72,
		0x62, 0x74, 0x72, 0x65, 0x65, 0x22,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 2153 bytes
func (FileMy2) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"my2.go","SourcePackage":"testdata","ZebraSchemaId":0,"Structs":{"inn":{"StructName":"inn","Fields":[{"Zid":0,"FieldGoName":"j","FieldTagName":"j","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":1,"FieldGoName":"e","FieldTagName":"e","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"u":{"StructName":"u","Fields":[{"Zid":0,"FieldGoName":"m","FieldTagName":"m","FieldTypeStr":"map[string]*Tr","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"Tr"}}}},{"Zid":1,"FieldGoName":"s","FieldTagName":"s","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"n","FieldTagName":"n","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Tr":{"StructName":"Tr","Fields":[{"Zid":0,"FieldGoName":"U","FieldTagName":"U","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":-1,"FieldGoName":"et","FieldTagName":"-","FieldCategory":28,"Skip":true},{"Zid":1,"FieldGoName":"Nt","FieldTagName":"Nt","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}},{"Zid":2,"FieldGoName":"Snot","FieldTagName":"Snot","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":3,"FieldGoName":"Notyet","FieldTagName":"Notyet","FieldTypeStr":"map[string]bool","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":18,"Str":"bool"}}},{"Zid":4,"FieldGoName":"Setm","FieldTagName":"Setm","FieldTypeStr":"[]*inn","FieldCategory":26,"FieldFullType":{"Kind":26,"Str":"Slice","Domain":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"inn"}}}}]}},"Imports":["\"github.com/glycerine/rbtree\""]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 6280 bytes
func (FileMy2) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "my2.go",
    "SourcePackage": "testdata",
    "ZebraSchemaId": 0,
    "Structs": {
        "inn": {
            "StructName": "inn",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "j",
                    "FieldTagName": "j",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "e",
                    "FieldTagName": "e",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        },
        "u": {
            "StructName": "u",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "m",
                    "FieldTagName": "m",
                    "FieldTypeStr": "map[string]*Tr",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "Tr"
                            }
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "s",
                    "FieldTagName": "s",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "n",
                    "FieldTagName": "n",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        },
        "Tr": {
            "StructName": "Tr",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "U",
                    "FieldTagName": "U",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": -1,
                    "FieldGoName": "et",
                    "FieldTagName": "-",
                    "FieldCategory": 28,
                    "Skip": true
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Nt",
                    "FieldTagName": "Nt",
                    "FieldTypeStr": "int",
                    "FieldCategory": 23,
                    "FieldPrimitive": 13,
                    "FieldFullType": {
                        "Kind": 13,
                        "Str": "int"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Snot",
                    "FieldTagName": "Snot",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 3,
                    "FieldGoName": "Notyet",
                    "FieldTagName": "Notyet",
                    "FieldTypeStr": "map[string]bool",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 18,
                            "Str": "bool"
                        }
                    }
                },
                {
                    "Zid": 4,
                    "FieldGoName": "Setm",
                    "FieldTagName": "Setm",
                    "FieldTypeStr": "[]*inn",
                    "FieldCategory": 26,
                    "FieldFullType": {
                        "Kind": 26,
                        "Str": "Slice",
                        "Domain": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "inn"
                            }
                        }
                    }
                }
            ]
        }
    },
    "Imports": [
        "\"github.com/glycerine/rbtree\""
    ]
}`)
}
