// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pocket/session/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/pokt-network/poktroll/x/application/types"
	types1 "github.com/pokt-network/poktroll/x/shared/types"
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

// SessionHeader is a lightweight header for a session that can be passed around.
// It is the minimal amount of data required to hydrate & retrieve all data relevant to the session.
type SessionHeader struct {
	ApplicationAddress string `protobuf:"bytes,1,opt,name=application_address,json=applicationAddress,proto3" json:"application_address,omitempty"`
	ServiceId          string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// NOTE: session_id can be derived from the above values using onchain but is included in the header for convenience
	SessionId               string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionStartBlockHeight int64  `protobuf:"varint,4,opt,name=session_start_block_height,json=sessionStartBlockHeight,proto3" json:"session_start_block_height,omitempty"`
	// Note that`session_end_block_height` is a derivative of (`start` + `num_blocks_per_session`)
	// as goverened by onchain params at the time of the session start.
	// It is stored as an additional field to simplofy business logic in case
	// the number of blocks_per_session changes during the session.
	SessionEndBlockHeight int64 `protobuf:"varint,5,opt,name=session_end_block_height,json=sessionEndBlockHeight,proto3" json:"session_end_block_height,omitempty"`
}

func (m *SessionHeader) Reset()         { *m = SessionHeader{} }
func (m *SessionHeader) String() string { return proto.CompactTextString(m) }
func (*SessionHeader) ProtoMessage()    {}
func (*SessionHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_bab703d2ebb9acfb, []int{0}
}
func (m *SessionHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SessionHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionHeader.Merge(m, src)
}
func (m *SessionHeader) XXX_Size() int {
	return m.Size()
}
func (m *SessionHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SessionHeader proto.InternalMessageInfo

func (m *SessionHeader) GetApplicationAddress() string {
	if m != nil {
		return m.ApplicationAddress
	}
	return ""
}

func (m *SessionHeader) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SessionHeader) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *SessionHeader) GetSessionStartBlockHeight() int64 {
	if m != nil {
		return m.SessionStartBlockHeight
	}
	return 0
}

func (m *SessionHeader) GetSessionEndBlockHeight() int64 {
	if m != nil {
		return m.SessionEndBlockHeight
	}
	return 0
}

// Session is a fully hydrated session object that contains all the information for the Session
// and its parcipants.
type Session struct {
	Header              *SessionHeader     `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SessionId           string             `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionNumber       int64              `protobuf:"varint,3,opt,name=session_number,json=sessionNumber,proto3" json:"session_number,omitempty"`
	NumBlocksPerSession int64              `protobuf:"varint,4,opt,name=num_blocks_per_session,json=numBlocksPerSession,proto3" json:"num_blocks_per_session,omitempty"`
	Application         *types.Application `protobuf:"bytes,5,opt,name=application,proto3" json:"application,omitempty"`
	Suppliers           []*types1.Supplier `protobuf:"bytes,6,rep,name=suppliers,proto3" json:"suppliers,omitempty"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_bab703d2ebb9acfb, []int{1}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetHeader() *SessionHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetSessionNumber() int64 {
	if m != nil {
		return m.SessionNumber
	}
	return 0
}

func (m *Session) GetNumBlocksPerSession() int64 {
	if m != nil {
		return m.NumBlocksPerSession
	}
	return 0
}

func (m *Session) GetApplication() *types.Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Session) GetSuppliers() []*types1.Supplier {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionHeader)(nil), "pocket.session.SessionHeader")
	proto.RegisterType((*Session)(nil), "pocket.session.Session")
}

func init() { proto.RegisterFile("pocket/session/types.proto", fileDescriptor_bab703d2ebb9acfb) }

var fileDescriptor_bab703d2ebb9acfb = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0x16, 0x8a, 0xea, 0x6a, 0x3b, 0x78, 0x83, 0x85, 0xc0, 0x42, 0x35, 0x09, 0xa9,
	0x97, 0x25, 0xa8, 0xd3, 0xc4, 0x81, 0x53, 0x2b, 0x21, 0xad, 0x97, 0x09, 0xa5, 0x37, 0x2e, 0x51,
	0x1a, 0xbf, 0x4a, 0xa2, 0x36, 0x71, 0x64, 0x3b, 0xfc, 0xf9, 0x12, 0x88, 0x0f, 0x03, 0xdf, 0x81,
	0xe3, 0xc4, 0x69, 0x47, 0xd4, 0x7e, 0x11, 0x14, 0xfb, 0xcd, 0x96, 0xe6, 0x16, 0xbf, 0xbf, 0xf7,
	0xf1, 0xfb, 0x3c, 0xb6, 0x43, 0x9c, 0x92, 0xc7, 0x1b, 0x50, 0xbe, 0x04, 0x29, 0x33, 0x5e, 0xf8,
	0xea, 0x7b, 0x09, 0xd2, 0x2b, 0x05, 0x57, 0x9c, 0x1e, 0x1b, 0xe6, 0x21, 0x73, 0x5e, 0xc6, 0x5c,
	0xe6, 0x5c, 0x86, 0x9a, 0xfa, 0x66, 0x61, 0x5a, 0x9d, 0x57, 0xcd, 0x36, 0x69, 0x24, 0x80, 0xf9,
	0x12, 0xc4, 0x97, 0x2c, 0x06, 0x84, 0x2e, 0xc2, 0xa8, 0x2c, 0xb7, 0x59, 0x1c, 0xa9, 0xce, 0x1c,
	0xe7, 0x75, 0x47, 0x5c, 0xd5, 0x7d, 0x20, 0x90, 0x9e, 0x26, 0x3c, 0xe1, 0x66, 0x64, 0xfd, 0x65,
	0xaa, 0x17, 0x3f, 0xfa, 0xe4, 0x68, 0x65, 0x7c, 0xdd, 0x40, 0xc4, 0x40, 0xd0, 0x25, 0x39, 0x69,
	0x0d, 0x08, 0x23, 0xc6, 0x04, 0x48, 0x69, 0x5b, 0x13, 0x6b, 0x3a, 0x5a, 0xd8, 0x7f, 0x7f, 0x5d,
	0x9e, 0xa2, 0xe3, 0xb9, 0x21, 0x2b, 0x25, 0xb2, 0x22, 0x09, 0x68, 0x4b, 0x84, 0x84, 0x9e, 0x13,
	0x82, 0x09, 0xc2, 0x8c, 0xd9, 0xfd, 0x7a, 0x87, 0x60, 0x84, 0x95, 0x25, 0x33, 0x58, 0x8f, 0xae,
	0xf1, 0xa0, 0xc1, 0xba, 0xb2, 0x64, 0xf4, 0x03, 0x71, 0x1a, 0x2c, 0x55, 0x24, 0x54, 0xb8, 0xde,
	0xf2, 0x78, 0x13, 0xa6, 0x90, 0x25, 0xa9, 0xb2, 0x9f, 0x4c, 0xac, 0xe9, 0x20, 0x38, 0xc3, 0x8e,
	0x55, 0xdd, 0xb0, 0xa8, 0xf9, 0x8d, 0xc6, 0xf4, 0x3d, 0xb1, 0x1b, 0x31, 0x14, 0xec, 0x50, 0xfa,
	0x54, 0x4b, 0x9f, 0x23, 0xff, 0x58, 0xb0, 0x96, 0xf0, 0xe2, 0x77, 0x9f, 0x3c, 0xc3, 0x03, 0xa1,
	0xd7, 0x64, 0x98, 0xea, 0x43, 0xd1, 0xe9, 0xc7, 0xb3, 0x73, 0xef, 0xf0, 0x26, 0xbd, 0x83, 0x93,
	0x0b, 0xb0, 0xb9, 0x93, 0xab, 0xdf, 0xcd, 0xf5, 0x96, 0x1c, 0x37, 0xb8, 0xa8, 0xf2, 0x35, 0x08,
	0x1d, 0x7d, 0x10, 0x1c, 0x61, 0xf5, 0x56, 0x17, 0xe9, 0x15, 0x79, 0x51, 0x54, 0xb9, 0x71, 0x2e,
	0xc3, 0x12, 0x44, 0x88, 0x1c, 0xa3, 0x9f, 0x14, 0x55, 0xae, 0x8d, 0xcb, 0x4f, 0x20, 0x1a, 0xc7,
	0x73, 0x32, 0x6e, 0xdd, 0x83, 0x4e, 0x3a, 0x9e, 0xbd, 0x69, 0x6c, 0xb7, 0x90, 0x37, 0x7f, 0xfc,
	0x0e, 0xda, 0x1a, 0x7a, 0x4d, 0x46, 0xcd, 0xcb, 0x91, 0xf6, 0x70, 0x32, 0x98, 0x8e, 0x67, 0x67,
	0x0f, 0xb9, 0xf5, 0xcb, 0xf2, 0x56, 0xc8, 0x83, 0xc7, 0xce, 0xc5, 0xed, 0x9f, 0x9d, 0x6b, 0xdd,
	0xed, 0x5c, 0xeb, 0x7e, 0xe7, 0x5a, 0xff, 0x76, 0xae, 0xf5, 0x73, 0xef, 0xf6, 0xee, 0xf6, 0x6e,
	0xef, 0x7e, 0xef, 0xf6, 0x3e, 0xbf, 0x4b, 0x32, 0x95, 0x56, 0x6b, 0x2f, 0xe6, 0xb9, 0x5f, 0xf2,
	0x8d, 0xba, 0x2c, 0x40, 0x7d, 0xe5, 0x62, 0xa3, 0x17, 0x82, 0x6f, 0xb7, 0xfe, 0xb7, 0xc3, 0x5f,
	0x67, 0x3d, 0xd4, 0xef, 0xf3, 0xea, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x4d, 0x3c, 0xa0,
	0x59, 0x03, 0x00, 0x00,
}

func (m *SessionHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SessionEndBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SessionEndBlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.SessionStartBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SessionStartBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SessionId) > 0 {
		i -= len(m.SessionId)
		copy(dAtA[i:], m.SessionId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SessionId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ServiceId) > 0 {
		i -= len(m.ServiceId)
		copy(dAtA[i:], m.ServiceId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ServiceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ApplicationAddress) > 0 {
		i -= len(m.ApplicationAddress)
		copy(dAtA[i:], m.ApplicationAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ApplicationAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Suppliers) > 0 {
		for iNdEx := len(m.Suppliers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Suppliers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Application != nil {
		{
			size, err := m.Application.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.NumBlocksPerSession != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NumBlocksPerSession))
		i--
		dAtA[i] = 0x20
	}
	if m.SessionNumber != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SessionNumber))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SessionId) > 0 {
		i -= len(m.SessionId)
		copy(dAtA[i:], m.SessionId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SessionId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SessionHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ApplicationAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SessionStartBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.SessionStartBlockHeight))
	}
	if m.SessionEndBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.SessionEndBlockHeight))
	}
	return n
}

func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SessionNumber != 0 {
		n += 1 + sovTypes(uint64(m.SessionNumber))
	}
	if m.NumBlocksPerSession != 0 {
		n += 1 + sovTypes(uint64(m.NumBlocksPerSession))
	}
	if m.Application != nil {
		l = m.Application.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Suppliers) > 0 {
		for _, e := range m.Suppliers {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SessionHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionStartBlockHeight", wireType)
			}
			m.SessionStartBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionStartBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionEndBlockHeight", wireType)
			}
			m.SessionEndBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionEndBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &SessionHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionNumber", wireType)
			}
			m.SessionNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocksPerSession", wireType)
			}
			m.NumBlocksPerSession = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocksPerSession |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Application == nil {
				m.Application = &types.Application{}
			}
			if err := m.Application.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suppliers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Suppliers = append(m.Suppliers, &types1.Supplier{})
			if err := m.Suppliers[len(m.Suppliers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
