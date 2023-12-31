// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testnet/testnet/stored_poll.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StoredPoll struct {
	Index    string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Question string   `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Voters   []string `protobuf:"bytes,3,rep,name=voters,proto3" json:"voters,omitempty"`
}

func (m *StoredPoll) Reset()         { *m = StoredPoll{} }
func (m *StoredPoll) String() string { return proto.CompactTextString(m) }
func (*StoredPoll) ProtoMessage()    {}
func (*StoredPoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9f81d5d19d140d2, []int{0}
}
func (m *StoredPoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredPoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredPoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredPoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredPoll.Merge(m, src)
}
func (m *StoredPoll) XXX_Size() int {
	return m.Size()
}
func (m *StoredPoll) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredPoll.DiscardUnknown(m)
}

var xxx_messageInfo_StoredPoll proto.InternalMessageInfo

func (m *StoredPoll) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredPoll) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *StoredPoll) GetVoters() []string {
	if m != nil {
		return m.Voters
	}
	return nil
}

func init() {
	proto.RegisterType((*StoredPoll)(nil), "testnet.testnet.StoredPoll")
}

func init() { proto.RegisterFile("testnet/testnet/stored_poll.proto", fileDescriptor_c9f81d5d19d140d2) }

var fileDescriptor_c9f81d5d19d140d2 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x49, 0x2d, 0x2e,
	0xc9, 0x4b, 0x2d, 0xd1, 0x87, 0xd1, 0xc5, 0x25, 0xf9, 0x45, 0xa9, 0x29, 0xf1, 0x05, 0xf9, 0x39,
	0x39, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xfc, 0x50, 0x29, 0x3d, 0x28, 0xad, 0x14, 0xc6,
	0xc5, 0x15, 0x0c, 0x56, 0x15, 0x90, 0x9f, 0x93, 0x23, 0x24, 0xc2, 0xc5, 0x9a, 0x99, 0x97, 0x92,
	0x5a, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x49, 0x71, 0x71, 0x14, 0x96,
	0xa6, 0x16, 0x97, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x81, 0x25, 0xe0, 0x7c, 0x21, 0x31, 0x2e, 0xb6,
	0xb2, 0xfc, 0x92, 0xd4, 0xa2, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x28, 0xcf, 0xc9,
	0xf7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x8c, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x13, 0x8b, 0x33, 0x32, 0x8b, 0x33, 0xb2, 0x33, 0x4a,
	0x8b, 0x12, 0x33, 0x8b, 0x33, 0x2a, 0xc1, 0x0e, 0xd7, 0x05, 0xb9, 0xbc, 0x02, 0xee, 0x87, 0x92,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xf3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0xa8, 0x8a, 0x11, 0xe3, 0x00, 0x00, 0x00,
}

func (m *StoredPoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredPoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredPoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Voters) > 0 {
		for iNdEx := len(m.Voters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voters[iNdEx])
			copy(dAtA[i:], m.Voters[iNdEx])
			i = encodeVarintStoredPoll(dAtA, i, uint64(len(m.Voters[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintStoredPoll(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredPoll(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredPoll(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredPoll(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredPoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredPoll(uint64(l))
	}
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovStoredPoll(uint64(l))
	}
	if len(m.Voters) > 0 {
		for _, s := range m.Voters {
			l = len(s)
			n += 1 + l + sovStoredPoll(uint64(l))
		}
	}
	return n
}

func sovStoredPoll(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredPoll(x uint64) (n int) {
	return sovStoredPoll(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredPoll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredPoll
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoredPoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredPoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStoredPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStoredPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStoredPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voters = append(m.Voters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredPoll
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStoredPoll(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredPoll
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStoredPoll
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStoredPoll
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStoredPoll
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredPoll
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredPoll
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredPoll        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredPoll          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredPoll = fmt.Errorf("proto: unexpected end of group")
)
