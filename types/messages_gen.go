package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MessageSyncRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Hashes) >= int(zb0002) {
		z.Hashes = (z.Hashes)[:zb0002]
	} else {
		z.Hashes = make([]Hash, zb0002)
	}
	for za0001 := range z.Hashes {
		err = z.Hashes[za0001].DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageSyncRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Hashes)))
	if err != nil {
		return
	}
	for za0001 := range z.Hashes {
		err = z.Hashes[za0001].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageSyncRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Hashes)))
	for za0001 := range z.Hashes {
		o, err = z.Hashes[za0001].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageSyncRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Hashes) >= int(zb0002) {
		z.Hashes = (z.Hashes)[:zb0002]
	} else {
		z.Hashes = make([]Hash, zb0002)
	}
	for za0001 := range z.Hashes {
		bts, err = z.Hashes[za0001].UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageSyncRequest) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize
	for za0001 := range z.Hashes {
		s += z.Hashes[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MessageSyncResponse) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Txs) >= int(zb0002) {
		z.Txs = (z.Txs)[:zb0002]
	} else {
		z.Txs = make([]Txi, zb0002)
	}
	for za0001 := range z.Txs {
		err = z.Txs[za0001].DecodeMsg(dc)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MessageSyncResponse) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 1
	err = en.Append(0x91)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Txs)))
	if err != nil {
		return
	}
	for za0001 := range z.Txs {
		err = z.Txs[za0001].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MessageSyncResponse) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 1
	o = append(o, 0x91)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Txs)))
	for za0001 := range z.Txs {
		o, err = z.Txs[za0001].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MessageSyncResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 1 {
		err = msgp.ArrayError{Wanted: 1, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Txs) >= int(zb0002) {
		z.Txs = (z.Txs)[:zb0002]
	} else {
		z.Txs = make([]Txi, zb0002)
	}
	for za0001 := range z.Txs {
		bts, err = z.Txs[za0001].UnmarshalMsg(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MessageSyncResponse) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize
	for za0001 := range z.Txs {
		s += z.Txs[za0001].Msgsize()
	}
	return
}
