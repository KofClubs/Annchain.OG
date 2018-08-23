package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TxBase) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	z.Type, err = dc.ReadInt()
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.ParentsHash) >= int(zb0002) {
		z.ParentsHash = (z.ParentsHash)[:zb0002]
	} else {
		z.ParentsHash = make([]Hash, zb0002)
	}
	for za0001 := range z.ParentsHash {
		err = z.ParentsHash[za0001].DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	z.SequenceNonce, err = dc.ReadUint64()
	if err != nil {
		return
	}
	z.Height, err = dc.ReadUint64()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TxBase) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Type)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.ParentsHash)))
	if err != nil {
		return
	}
	for za0001 := range z.ParentsHash {
		err = z.ParentsHash[za0001].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	err = en.WriteUint64(z.SequenceNonce)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Height)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TxBase) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendInt(o, z.Type)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ParentsHash)))
	for za0001 := range z.ParentsHash {
		o, err = z.ParentsHash[za0001].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	o = msgp.AppendUint64(o, z.SequenceNonce)
	o = msgp.AppendUint64(o, z.Height)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TxBase) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	z.Type, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.ParentsHash) >= int(zb0002) {
		z.ParentsHash = (z.ParentsHash)[:zb0002]
	} else {
		z.ParentsHash = make([]Hash, zb0002)
	}
	for za0001 := range z.ParentsHash {
		bts, err = z.ParentsHash[za0001].UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	z.SequenceNonce, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	z.Height, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TxBase) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.ArrayHeaderSize
	for za0001 := range z.ParentsHash {
		s += z.ParentsHash[za0001].Msgsize()
	}
	s += msgp.Uint64Size + msgp.Uint64Size
	return
}
