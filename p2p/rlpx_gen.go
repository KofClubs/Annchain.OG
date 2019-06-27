package p2p

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AuthMsgV4) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	err = dc.ReadExactBytes((z.Signature)[:])
	if err != nil {
		return
	}
	err = dc.ReadExactBytes((z.InitiatorPubkey)[:])
	if err != nil {
		return
	}
	err = dc.ReadExactBytes((z.Nonce)[:])
	if err != nil {
		return
	}
	z.Version, err = dc.ReadUint32()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AuthMsgV4) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Signature)[:])
	if err != nil {
		return
	}
	err = en.WriteBytes((z.InitiatorPubkey)[:])
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Nonce)[:])
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AuthMsgV4) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendBytes(o, (z.Signature)[:])
	o = msgp.AppendBytes(o, (z.InitiatorPubkey)[:])
	o = msgp.AppendBytes(o, (z.Nonce)[:])
	o = msgp.AppendUint32(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AuthMsgV4) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	bts, err = msgp.ReadExactBytes(bts, (z.Signature)[:])
	if err != nil {
		return
	}
	bts, err = msgp.ReadExactBytes(bts, (z.InitiatorPubkey)[:])
	if err != nil {
		return
	}
	bts, err = msgp.ReadExactBytes(bts, (z.Nonce)[:])
	if err != nil {
		return
	}
	z.Version, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AuthMsgV4) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (sigLen * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (pubLen * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (shaLen * (msgp.ByteSize)) + msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *AuthRespV4) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	err = dc.ReadExactBytes((z.RandomPubkey)[:])
	if err != nil {
		return
	}
	err = dc.ReadExactBytes((z.Nonce)[:])
	if err != nil {
		return
	}
	z.Version, err = dc.ReadUint32()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AuthRespV4) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.RandomPubkey)[:])
	if err != nil {
		return
	}
	err = en.WriteBytes((z.Nonce)[:])
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Version)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AuthRespV4) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o = msgp.AppendBytes(o, (z.RandomPubkey)[:])
	o = msgp.AppendBytes(o, (z.Nonce)[:])
	o = msgp.AppendUint32(o, z.Version)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AuthRespV4) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	bts, err = msgp.ReadExactBytes(bts, (z.RandomPubkey)[:])
	if err != nil {
		return
	}
	bts, err = msgp.ReadExactBytes(bts, (z.Nonce)[:])
	if err != nil {
		return
	}
	z.Version, bts, err = msgp.ReadUint32Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AuthRespV4) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (pubLen * (msgp.ByteSize)) + msgp.ArrayHeaderSize + (shaLen * (msgp.ByteSize)) + msgp.Uint32Size
	return
}
