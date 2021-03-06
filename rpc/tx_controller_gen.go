package rpc

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *NewTxRequest) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	z.Nonce, err = dc.ReadUint64()
	if err != nil {
		return
	}
	z.From, err = dc.ReadString()
	if err != nil {
		return
	}
	z.To, err = dc.ReadString()
	if err != nil {
		return
	}
	z.Value, err = dc.ReadString()
	if err != nil {
		return
	}
	z.Data, err = dc.ReadString()
	if err != nil {
		return
	}
	z.CryptoType, err = dc.ReadString()
	if err != nil {
		return
	}
	z.Signature, err = dc.ReadString()
	if err != nil {
		return
	}
	z.Pubkey, err = dc.ReadString()
	if err != nil {
		return
	}
	z.TokenId, err = dc.ReadInt32()
	if err != nil {
		return
	}
	// 存证哈希
	z.OpHash, err = dc.ReadString()
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *NewTxRequest) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 9
	err = en.Append(0x99)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Nonce)
	if err != nil {
		return
	}
	err = en.WriteString(z.From)
	if err != nil {
		return
	}
	err = en.WriteString(z.To)
	if err != nil {
		return
	}
	err = en.WriteString(z.Value)
	if err != nil {
		return
	}
	err = en.WriteString(z.Data)
	if err != nil {
		return
	}
	err = en.WriteString(z.CryptoType)
	if err != nil {
		return
	}
	err = en.WriteString(z.Signature)
	if err != nil {
		return
	}
	err = en.WriteString(z.Pubkey)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.TokenId)
	if err != nil {
		return
	}
	// 存证哈希
	err = en.WriteString(z.OpHash)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NewTxRequest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 9
	o = append(o, 0x99)
	o = msgp.AppendUint64(o, z.Nonce)
	o = msgp.AppendString(o, z.From)
	o = msgp.AppendString(o, z.To)
	o = msgp.AppendString(o, z.Value)
	o = msgp.AppendString(o, z.Data)
	o = msgp.AppendString(o, z.CryptoType)
	o = msgp.AppendString(o, z.Signature)
	o = msgp.AppendString(o, z.Pubkey)
	o = msgp.AppendInt32(o, z.TokenId)
	// 存证哈希
	o = msgp.AppendString(o, z.OpHash)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NewTxRequest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 9 {
		err = msgp.ArrayError{Wanted: 9, Got: zb0001}
		return
	}
	z.Nonce, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	z.From, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.To, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.Value, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.Data, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.CryptoType, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.Signature, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.Pubkey, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.TokenId, bts, err = msgp.ReadInt32Bytes(bts)
	if err != nil {
		return
	}
	// 存证哈希
	z.OpHash, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NewTxRequest) Msgsize() (s int) {
	s = 1 + msgp.Uint64Size + msgp.StringPrefixSize + len(z.From) + msgp.StringPrefixSize + len(z.To) + msgp.StringPrefixSize + len(z.Value) + msgp.StringPrefixSize + len(z.Data) + msgp.StringPrefixSize + len(z.CryptoType) + msgp.StringPrefixSize + len(z.Signature) + msgp.StringPrefixSize + len(z.Pubkey) + msgp.Int32Size /* 存证哈希 */ + msgp.StringPrefixSize + len(z.OpHash)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NewTxsRequests) DecodeMsg(dc *msgp.Reader) (err error) {
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
		z.Txs = make([]NewTxRequest, zb0002)
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
func (z *NewTxsRequests) EncodeMsg(en *msgp.Writer) (err error) {
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
func (z *NewTxsRequests) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *NewTxsRequests) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		z.Txs = make([]NewTxRequest, zb0002)
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
func (z *NewTxsRequests) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize
	for za0001 := range z.Txs {
		s += z.Txs[za0001].Msgsize()
	}
	return
}
