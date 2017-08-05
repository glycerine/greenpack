package testdata

// NOTE: THIS FILE WAS PRODUCED BY THE
// GREENPACK CODE GENERATION TOOL (github.com/glycerine/greenpack)
// DO NOT EDIT

import (
	"github.com/glycerine/greenpack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *TupleByDefaultTestStruct) DecodeMsg(dc *msgp.Reader) (err error) {

	var zgensym_f43bd773bf01f77_0 uint32
	zgensym_f43bd773bf01f77_0, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zgensym_f43bd773bf01f77_0 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zgensym_f43bd773bf01f77_0}
		return
	}
	z.S, err = dc.ReadString()
	if err != nil {
		return
	}
	z.N, err = dc.ReadInt64()
	if err != nil {
		return
	}
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// EncodeMsg implements msgp.Encodable
func (z TupleByDefaultTestStruct) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return err
	}
	err = en.WriteString(z.S)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.N)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TupleByDefaultTestStruct) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.S)
	o = msgp.AppendInt64(o, z.N)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TupleByDefaultTestStruct) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.UnmarshalMsgWithCfg(bts, nil)
}
func (z *TupleByDefaultTestStruct) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var zgensym_f43bd773bf01f77_1 uint32
	zgensym_f43bd773bf01f77_1, bts, err = nbs.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zgensym_f43bd773bf01f77_1 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zgensym_f43bd773bf01f77_1}
		return
	}
	z.S, bts, err = nbs.ReadStringBytes(bts)

	if err != nil {
		return
	}
	z.N, bts, err = nbs.ReadInt64Bytes(bts)

	if err != nil {
		return
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

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TupleByDefaultTestStruct) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.S) + msgp.Int64Size
	return
}
