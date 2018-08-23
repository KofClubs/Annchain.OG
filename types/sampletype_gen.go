package types

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Foo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "bar":
			z.Bar, err = dc.ReadString()
			if err != nil {
				return
			}
		case "baz":
			z.Baz, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "address":
			err = z.Address.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "parents":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Parents) >= int(zb0002) {
				z.Parents = (z.Parents)[:zb0002]
			} else {
				z.Parents = make([]Hash, zb0002)
			}
			for za0001 := range z.Parents {
				err = z.Parents[za0001].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "kv":
			var zb0003 uint32
			zb0003, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.KV == nil {
				z.KV = make(map[string]float64, zb0003)
			} else if len(z.KV) > 0 {
				for key := range z.KV {
					delete(z.KV, key)
				}
			}
			for zb0003 > 0 {
				zb0003--
				var za0002 string
				var za0003 float64
				za0002, err = dc.ReadString()
				if err != nil {
					return
				}
				za0003, err = dc.ReadFloat64()
				if err != nil {
					return
				}
				z.KV[za0002] = za0003
			}
		case "seq":
			err = z.Seq.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "tx":
			err = z.TxInner.DecodeMsg(dc)
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
func (z *Foo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "bar"
	err = en.Append(0x87, 0xa3, 0x62, 0x61, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bar)
	if err != nil {
		return
	}
	// write "baz"
	err = en.Append(0xa3, 0x62, 0x61, 0x7a)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Baz)
	if err != nil {
		return
	}
	// write "address"
	err = en.Append(0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	if err != nil {
		return
	}
	err = z.Address.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "parents"
	err = en.Append(0xa7, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Parents)))
	if err != nil {
		return
	}
	for za0001 := range z.Parents {
		err = z.Parents[za0001].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "kv"
	err = en.Append(0xa2, 0x6b, 0x76)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.KV)))
	if err != nil {
		return
	}
	for za0002, za0003 := range z.KV {
		err = en.WriteString(za0002)
		if err != nil {
			return
		}
		err = en.WriteFloat64(za0003)
		if err != nil {
			return
		}
	}
	// write "seq"
	err = en.Append(0xa3, 0x73, 0x65, 0x71)
	if err != nil {
		return
	}
	err = z.Seq.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "tx"
	err = en.Append(0xa2, 0x74, 0x78)
	if err != nil {
		return
	}
	err = z.TxInner.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Foo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "bar"
	o = append(o, 0x87, 0xa3, 0x62, 0x61, 0x72)
	o = msgp.AppendString(o, z.Bar)
	// string "baz"
	o = append(o, 0xa3, 0x62, 0x61, 0x7a)
	o = msgp.AppendFloat64(o, z.Baz)
	// string "address"
	o = append(o, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o, err = z.Address.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "parents"
	o = append(o, 0xa7, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Parents)))
	for za0001 := range z.Parents {
		o, err = z.Parents[za0001].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "kv"
	o = append(o, 0xa2, 0x6b, 0x76)
	o = msgp.AppendMapHeader(o, uint32(len(z.KV)))
	for za0002, za0003 := range z.KV {
		o = msgp.AppendString(o, za0002)
		o = msgp.AppendFloat64(o, za0003)
	}
	// string "seq"
	o = append(o, 0xa3, 0x73, 0x65, 0x71)
	o, err = z.Seq.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "tx"
	o = append(o, 0xa2, 0x74, 0x78)
	o, err = z.TxInner.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Foo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "bar":
			z.Bar, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "baz":
			z.Baz, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "address":
			bts, err = z.Address.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "parents":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Parents) >= int(zb0002) {
				z.Parents = (z.Parents)[:zb0002]
			} else {
				z.Parents = make([]Hash, zb0002)
			}
			for za0001 := range z.Parents {
				bts, err = z.Parents[za0001].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "kv":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.KV == nil {
				z.KV = make(map[string]float64, zb0003)
			} else if len(z.KV) > 0 {
				for key := range z.KV {
					delete(z.KV, key)
				}
			}
			for zb0003 > 0 {
				var za0002 string
				var za0003 float64
				zb0003--
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				za0003, bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
				z.KV[za0002] = za0003
			}
		case "seq":
			bts, err = z.Seq.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "tx":
			bts, err = z.TxInner.UnmarshalMsg(bts)
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
func (z *Foo) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Bar) + 4 + msgp.Float64Size + 8 + z.Address.Msgsize() + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Parents {
		s += z.Parents[za0001].Msgsize()
	}
	s += 3 + msgp.MapHeaderSize
	if z.KV != nil {
		for za0002, za0003 := range z.KV {
			_ = za0003
			s += msgp.StringPrefixSize + len(za0002) + msgp.Float64Size
		}
	}
	s += 4 + z.Seq.Msgsize() + 3 + z.TxInner.Msgsize()
	return
}
