package enr

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Pair) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.K, err = dc.ReadString()
	if err != nil {
		return
	}
	z.V, err = dc.ReadBytes(z.V)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Pair) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteString(z.K)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.V)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Pair) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendString(o, z.K)
	o = msgp.AppendBytes(o, z.V)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Pair) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.K, bts, err = msgp.ReadStringBytes(bts)
	if err != nil {
		return
	}
	z.V, bts, err = msgp.ReadBytesBytes(bts, z.V)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Pair) Msgsize() (s int) {
	s = 1 + msgp.StringPrefixSize + len(z.K) + msgp.BytesPrefixSize + len(z.V)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Record) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.Seq, err = dc.ReadUint64()
	if err != nil {
		return
	}
	z.Signature, err = dc.ReadBytes(z.Signature)
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap(z.Pairs) >= int(zb0002) {
		z.Pairs = (z.Pairs)[:zb0002]
	} else {
		z.Pairs = make([]Pair, zb0002)
	}
	for za0001 := range z.Pairs {
		var zb0003 uint32
		zb0003, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		z.Pairs[za0001].K, err = dc.ReadString()
		if err != nil {
			return
		}
		z.Pairs[za0001].V, err = dc.ReadBytes(z.Pairs[za0001].V)
		if err != nil {
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Record) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 3
	err = en.Append(0x93)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Seq)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Signature)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Pairs)))
	if err != nil {
		return
	}
	for za0001 := range z.Pairs {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = en.WriteString(z.Pairs[za0001].K)
		if err != nil {
			return
		}
		err = en.WriteBytes(z.Pairs[za0001].V)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Record) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o = msgp.AppendUint64(o, z.Seq)
	o = msgp.AppendBytes(o, z.Signature)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Pairs)))
	for za0001 := range z.Pairs {
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendString(o, z.Pairs[za0001].K)
		o = msgp.AppendBytes(o, z.Pairs[za0001].V)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Record) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.Seq, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		return
	}
	z.Signature, bts, err = msgp.ReadBytesBytes(bts, z.Signature)
	if err != nil {
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap(z.Pairs) >= int(zb0002) {
		z.Pairs = (z.Pairs)[:zb0002]
	} else {
		z.Pairs = make([]Pair, zb0002)
	}
	for za0001 := range z.Pairs {
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			return
		}
		if zb0003 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0003}
			return
		}
		z.Pairs[za0001].K, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			return
		}
		z.Pairs[za0001].V, bts, err = msgp.ReadBytesBytes(bts, z.Pairs[za0001].V)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Record) Msgsize() (s int) {
	s = 1 + msgp.Uint64Size + msgp.BytesPrefixSize + len(z.Signature) + msgp.ArrayHeaderSize
	for za0001 := range z.Pairs {
		s += 1 + msgp.StringPrefixSize + len(z.Pairs[za0001].K) + msgp.BytesPrefixSize + len(z.Pairs[za0001].V)
	}
	return
}
