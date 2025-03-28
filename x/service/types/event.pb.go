// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pocket/service/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// EventRelayMiningDifficultyUpdated is an event emitted whenever the relay mining difficulty is updated
// for a given service.
type EventRelayMiningDifficultyUpdated struct {
	ServiceId                string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	PrevTargetHashHexEncoded string `protobuf:"bytes,2,opt,name=prev_target_hash_hex_encoded,json=prevTargetHashHexEncoded,proto3" json:"prev_target_hash_hex_encoded,omitempty"`
	NewTargetHashHexEncoded  string `protobuf:"bytes,3,opt,name=new_target_hash_hex_encoded,json=newTargetHashHexEncoded,proto3" json:"new_target_hash_hex_encoded,omitempty"`
	PrevNumRelaysEma         uint64 `protobuf:"varint,4,opt,name=prev_num_relays_ema,json=prevNumRelaysEma,proto3" json:"prev_num_relays_ema,omitempty"`
	NewNumRelaysEma          uint64 `protobuf:"varint,5,opt,name=new_num_relays_ema,json=newNumRelaysEma,proto3" json:"new_num_relays_ema,omitempty"`
}

func (m *EventRelayMiningDifficultyUpdated) Reset()         { *m = EventRelayMiningDifficultyUpdated{} }
func (m *EventRelayMiningDifficultyUpdated) String() string { return proto.CompactTextString(m) }
func (*EventRelayMiningDifficultyUpdated) ProtoMessage()    {}
func (*EventRelayMiningDifficultyUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_a38747b533ead694, []int{0}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.Merge(m, src)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventRelayMiningDifficultyUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRelayMiningDifficultyUpdated proto.InternalMessageInfo

func (m *EventRelayMiningDifficultyUpdated) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevTargetHashHexEncoded() string {
	if m != nil {
		return m.PrevTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetNewTargetHashHexEncoded() string {
	if m != nil {
		return m.NewTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevNumRelaysEma() uint64 {
	if m != nil {
		return m.PrevNumRelaysEma
	}
	return 0
}

func (m *EventRelayMiningDifficultyUpdated) GetNewNumRelaysEma() uint64 {
	if m != nil {
		return m.NewNumRelaysEma
	}
	return 0
}

func init() {
	proto.RegisterType((*EventRelayMiningDifficultyUpdated)(nil), "pocket.service.EventRelayMiningDifficultyUpdated")
}

func init() { proto.RegisterFile("pocket/service/event.proto", fileDescriptor_a38747b533ead694) }

var fileDescriptor_a38747b533ead694 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x97, 0x39, 0x85, 0xe5, 0xa0, 0x52, 0x05, 0xcb, 0xd4, 0x30, 0x3d, 0x0d, 0x64, 0xab,
	0xe0, 0x55, 0x3c, 0x88, 0x83, 0x79, 0x70, 0x87, 0xa1, 0x17, 0x2f, 0x21, 0x6b, 0xbf, 0xb5, 0x61,
	0x6d, 0x52, 0x92, 0x74, 0xdd, 0xde, 0xc2, 0x07, 0xf0, 0x81, 0x3c, 0xee, 0xb8, 0xa3, 0x74, 0x2f,
	0x22, 0xc9, 0x86, 0x88, 0xe8, 0x2d, 0xc9, 0xff, 0xf7, 0xcb, 0xc7, 0x3f, 0xc1, 0xad, 0x5c, 0x86,
	0x53, 0x30, 0x81, 0x06, 0x35, 0xe3, 0x21, 0x04, 0x30, 0x03, 0x61, 0x7a, 0xb9, 0x92, 0x46, 0x7a,
	0xfb, 0x9b, 0xac, 0xb7, 0xcd, 0x5a, 0xc7, 0xb1, 0x8c, 0xa5, 0x8b, 0x02, 0xbb, 0xda, 0x50, 0x97,
	0xef, 0x75, 0x7c, 0xd1, 0xb7, 0xd6, 0x08, 0x52, 0xb6, 0x78, 0xe2, 0x82, 0x8b, 0xf8, 0x81, 0x4f,
	0x26, 0x3c, 0x2c, 0x52, 0xb3, 0x78, 0xc9, 0x23, 0x66, 0x20, 0xf2, 0xce, 0x31, 0xde, 0x5e, 0x43,
	0x79, 0xe4, 0xa3, 0x36, 0xea, 0x34, 0x47, 0xcd, 0xed, 0xc9, 0x63, 0xe4, 0xdd, 0xe1, 0xb3, 0x5c,
	0xc1, 0x8c, 0x1a, 0xa6, 0x62, 0x30, 0x34, 0x61, 0x3a, 0xa1, 0x09, 0xcc, 0x29, 0x88, 0x50, 0x46,
	0x10, 0xf9, 0x75, 0x27, 0xf8, 0x96, 0x79, 0x76, 0xc8, 0x80, 0xe9, 0x64, 0x00, 0xf3, 0xfe, 0x26,
	0xf7, 0x6e, 0xf1, 0xa9, 0x80, 0xf2, 0x5f, 0x7d, 0xc7, 0xe9, 0x27, 0x02, 0xca, 0x3f, 0xed, 0x2e,
	0x3e, 0x72, 0xd3, 0x45, 0x91, 0x51, 0x65, 0x5b, 0x68, 0x0a, 0x19, 0xf3, 0x1b, 0x6d, 0xd4, 0x69,
	0x8c, 0x0e, 0x6d, 0x34, 0x2c, 0x32, 0x57, 0x4f, 0xf7, 0x33, 0xe6, 0x5d, 0x61, 0xcf, 0x0e, 0xfb,
	0x45, 0xef, 0x3a, 0xfa, 0x40, 0x40, 0xf9, 0x13, 0xbe, 0x1f, 0x7e, 0x54, 0x04, 0x2d, 0x2b, 0x82,
	0x56, 0x15, 0x41, 0x9f, 0x15, 0x41, 0x6f, 0x6b, 0x52, 0x5b, 0xae, 0x49, 0x6d, 0xb5, 0x26, 0xb5,
	0xd7, 0xeb, 0x98, 0x9b, 0xa4, 0x18, 0xf7, 0x42, 0x99, 0x05, 0xb9, 0x9c, 0x9a, 0xae, 0x00, 0x53,
	0x4a, 0x35, 0x75, 0x1b, 0x25, 0xd3, 0x34, 0x98, 0x7f, 0x7f, 0x8d, 0x59, 0xe4, 0xa0, 0xc7, 0x7b,
	0xee, 0xd5, 0x6f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x09, 0xb0, 0x48, 0xb9, 0x01, 0x00,
	0x00,
}

func (m *EventRelayMiningDifficultyUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRelayMiningDifficultyUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRelayMiningDifficultyUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NewNumRelaysEma))
		i--
		dAtA[i] = 0x28
	}
	if m.PrevNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PrevNumRelaysEma))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NewTargetHashHexEncoded) > 0 {
		i -= len(m.NewTargetHashHexEncoded)
		copy(dAtA[i:], m.NewTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.NewTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PrevTargetHashHexEncoded) > 0 {
		i -= len(m.PrevTargetHashHexEncoded)
		copy(dAtA[i:], m.PrevTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PrevTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceId) > 0 {
		i -= len(m.ServiceId)
		copy(dAtA[i:], m.ServiceId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ServiceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRelayMiningDifficultyUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PrevTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.NewTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.PrevNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.PrevNumRelaysEma))
	}
	if m.NewNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.NewNumRelaysEma))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRelayMiningDifficultyUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevNumRelaysEma", wireType)
			}
			m.PrevNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewNumRelaysEma", wireType)
			}
			m.NewNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
