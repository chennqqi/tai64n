// Code generated by protoc-gen-gogo.
// source: tai64n.proto
// DO NOT EDIT!

/*
	Package tai64n is a generated protocol buffer package.

	It is generated from these files:
		tai64n.proto

	It has these top-level messages:
		TAI64N
*/
package tai64n

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

import fmt1 "fmt"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TAI64N struct {
	Seconds          uint64 `protobuf:"varint,1,req,name=seconds" json:"seconds" codec:"seconds"`
	Nanoseconds      uint32 `protobuf:"varint,2,req,name=nanoseconds" json:"nanoseconds" codec:"nanoseconds"`
	XXX_unrecognized []byte `json:"-" codec:"-"`
}

func (m *TAI64N) Reset()         { *m = TAI64N{} }
func (m *TAI64N) String() string { return proto.CompactTextString(m) }
func (*TAI64N) ProtoMessage()    {}

func (m *TAI64N) GetSeconds() uint64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TAI64N) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

func init() {
}
func (m *TAI64N) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Seconds |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanoseconds", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Nanoseconds |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *TAI64N) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTai64N(uint64(m.Seconds))
	n += 1 + sovTai64N(uint64(m.Nanoseconds))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTai64N(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTai64N(x uint64) (n int) {
	return sovTai64N(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TAI64N) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TAI64N) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintTai64N(data, i, uint64(m.Seconds))
	data[i] = 0x10
	i++
	i = encodeVarintTai64N(data, i, uint64(m.Nanoseconds))
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Tai64N(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Tai64N(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTai64N(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *TAI64N) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt1.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*TAI64N)
	if !ok {
		return fmt1.Errorf("that is not of type *TAI64N")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt1.Errorf("that is type *TAI64N but is nil && this != nil")
	} else if this == nil {
		return fmt1.Errorf("that is type *TAI64Nbut is not nil && this == nil")
	}
	if this.Seconds != that1.Seconds {
		return fmt1.Errorf("Seconds this(%v) Not Equal that(%v)", this.Seconds, that1.Seconds)
	}
	if this.Nanoseconds != that1.Nanoseconds {
		return fmt1.Errorf("Nanoseconds this(%v) Not Equal that(%v)", this.Nanoseconds, that1.Nanoseconds)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt1.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *TAI64N) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TAI64N)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Seconds != that1.Seconds {
		return false
	}
	if this.Nanoseconds != that1.Nanoseconds {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}