package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Sequencer) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBase":
			err = z.TxBase.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "Id":
			z.Id, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "Issuer":
			err = z.Issuer.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "ContractHashOrder":
			err = z.ContractHashOrder.DecodeMsg(dc)
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
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Sequencer) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "TxBase"
	err = en.Append(0x84, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
	if err != nil {
		return
	}
	err = z.TxBase.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "Id"
	err = en.Append(0xa2, 0x49, 0x64)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Id)
	if err != nil {
		return
	}
	// write "Issuer"
	err = en.Append(0xa6, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.Issuer.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "ContractHashOrder"
	err = en.Append(0xb1, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.ContractHashOrder.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Sequencer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "TxBase"
	o = append(o, 0x84, 0xa6, 0x54, 0x78, 0x42, 0x61, 0x73, 0x65)
	o, err = z.TxBase.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "Id"
	o = append(o, 0xa2, 0x49, 0x64)
	o = msgp.AppendUint64(o, z.Id)
	// string "Issuer"
	o = append(o, 0xa6, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72)
	o, err = z.Issuer.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "ContractHashOrder"
	o = append(o, 0xb1, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x61, 0x73, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72)
	o, err = z.ContractHashOrder.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sequencer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "TxBase":
			bts, err = z.TxBase.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "Id":
			z.Id, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "Issuer":
			bts, err = z.Issuer.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "ContractHashOrder":
			bts, err = z.ContractHashOrder.UnmarshalMsg(bts)
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
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Sequencer) Msgsize() (s int) {
	s = 1 + 7 + z.TxBase.Msgsize() + 3 + msgp.Uint64Size + 7 + z.Issuer.Msgsize() + 18 + z.ContractHashOrder.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Sequencers) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Sequencers, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Sequencer)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Sequencers) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Sequencers) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Sequencers) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Sequencers, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Sequencer)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Sequencers) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
