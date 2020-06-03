package og

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *BouncerMessage) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Value, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z BouncerMessage) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Value)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BouncerMessage) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendInt(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BouncerMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	z.Value, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Value")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z BouncerMessage) Msgsize() (s int) {
	s = 1 + msgp.IntSize
	return
}
