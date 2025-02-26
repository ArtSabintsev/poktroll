// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/migration/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_21658240592266b6, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21658240592266b6, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

// MsgImportMorseClaimableAccounts is used to create the on-chain MorseClaimableAccounts ONLY AND EXACTLY ONCE (per network / re-genesis).
type MsgImportMorseClaimableAccounts struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// the account state derived from the Morse state export and the `poktrolld migrate collect-morse-accounts` command.
	MorseAccountState MorseAccountState `protobuf:"bytes,2,opt,name=morse_account_state,json=morseAccountState,proto3" json:"morse_account_state"`
	// Validates the morse_account_state sha256 hash:
	// - Transaction fails if hash doesn't match on-chain computation
	// - Off-chain social consensus should be reached off-chain before verification
	//
	// Verification can be done by comparing with locally derived Morse state like so:
	//   $ poktrolld migrate collect-morse-accounts $<(pocket util export-genesis-for-reset)
	//
	// Additional documentation:
	// - pocket util export-genesis-for-migration --help
	// - poktrolld migrate collect-morse-accounts --help
	MorseAccountStateHash []byte `protobuf:"bytes,3,opt,name=morse_account_state_hash,json=morseAccountStateHash,proto3" json:"morse_account_state_hash"`
}

func (m *MsgImportMorseClaimableAccounts) Reset()         { *m = MsgImportMorseClaimableAccounts{} }
func (m *MsgImportMorseClaimableAccounts) String() string { return proto.CompactTextString(m) }
func (*MsgImportMorseClaimableAccounts) ProtoMessage()    {}
func (*MsgImportMorseClaimableAccounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_21658240592266b6, []int{2}
}
func (m *MsgImportMorseClaimableAccounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgImportMorseClaimableAccounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgImportMorseClaimableAccounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgImportMorseClaimableAccounts.Merge(m, src)
}
func (m *MsgImportMorseClaimableAccounts) XXX_Size() int {
	return m.Size()
}
func (m *MsgImportMorseClaimableAccounts) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgImportMorseClaimableAccounts.DiscardUnknown(m)
}

var xxx_messageInfo_MsgImportMorseClaimableAccounts proto.InternalMessageInfo

func (m *MsgImportMorseClaimableAccounts) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgImportMorseClaimableAccounts) GetMorseAccountState() MorseAccountState {
	if m != nil {
		return m.MorseAccountState
	}
	return MorseAccountState{}
}

func (m *MsgImportMorseClaimableAccounts) GetMorseAccountStateHash() []byte {
	if m != nil {
		return m.MorseAccountStateHash
	}
	return nil
}

// MsgImportMorseClaimableAccountsResponse is returned from MsgImportMorseClaimableAccounts.
// It indicates the canonical hash of the imported MorseAccountState, and the number of claimable accounts which were imported.
type MsgImportMorseClaimableAccountsResponse struct {
	// On-chain computed sha256 hash of the morse_account_state provided in the corresponding MsgCreateMorseAccountState.
	StateHash []byte `protobuf:"bytes,1,opt,name=state_hash,json=stateHash,proto3" json:"state_hash"`
	// Number of claimable accounts (EOAs) collected from Morse state export.
	NumAccounts uint64 `protobuf:"varint,2,opt,name=num_accounts,json=numAccounts,proto3" json:"num_accounts"`
}

func (m *MsgImportMorseClaimableAccountsResponse) Reset() {
	*m = MsgImportMorseClaimableAccountsResponse{}
}
func (m *MsgImportMorseClaimableAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgImportMorseClaimableAccountsResponse) ProtoMessage()    {}
func (*MsgImportMorseClaimableAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21658240592266b6, []int{3}
}
func (m *MsgImportMorseClaimableAccountsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgImportMorseClaimableAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MsgImportMorseClaimableAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgImportMorseClaimableAccountsResponse.Merge(m, src)
}
func (m *MsgImportMorseClaimableAccountsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgImportMorseClaimableAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgImportMorseClaimableAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgImportMorseClaimableAccountsResponse proto.InternalMessageInfo

func (m *MsgImportMorseClaimableAccountsResponse) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *MsgImportMorseClaimableAccountsResponse) GetNumAccounts() uint64 {
	if m != nil {
		return m.NumAccounts
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "poktroll.migration.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "poktroll.migration.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgImportMorseClaimableAccounts)(nil), "poktroll.migration.MsgImportMorseClaimableAccounts")
	proto.RegisterType((*MsgImportMorseClaimableAccountsResponse)(nil), "poktroll.migration.MsgImportMorseClaimableAccountsResponse")
}

func init() { proto.RegisterFile("poktroll/migration/tx.proto", fileDescriptor_21658240592266b6) }

var fileDescriptor_21658240592266b6 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xbf, 0x6b, 0xdb, 0x40,
	0x14, 0xf6, 0x39, 0x6d, 0xc0, 0x17, 0x93, 0x36, 0x6a, 0x4a, 0x1c, 0x25, 0x48, 0xc6, 0xfd, 0x65,
	0x5c, 0x6c, 0x51, 0x1b, 0x5a, 0x48, 0xe9, 0x10, 0x75, 0x69, 0x07, 0x43, 0x50, 0xc8, 0xd2, 0xc5,
	0x3d, 0xcb, 0x42, 0x12, 0xf1, 0xdd, 0x89, 0xbb, 0x53, 0x9b, 0x6c, 0x6d, 0xc7, 0x42, 0x21, 0x7f,
	0x46, 0x47, 0x0f, 0xf9, 0x07, 0xba, 0x65, 0x0c, 0x9d, 0x32, 0x89, 0x62, 0x0f, 0x06, 0xff, 0x15,
	0x45, 0xbf, 0x1c, 0xdb, 0x51, 0xd2, 0x92, 0x45, 0xd2, 0xbd, 0xf7, 0xbd, 0xef, 0x7d, 0xef, 0xe3,
	0x9d, 0xe0, 0x96, 0x47, 0x0f, 0x05, 0xa3, 0xfd, 0xbe, 0x86, 0x5d, 0x9b, 0x21, 0xe1, 0x52, 0xa2,
	0x89, 0xa3, 0x86, 0xc7, 0xa8, 0xa0, 0x92, 0x94, 0x26, 0x1b, 0xd3, 0xa4, 0xbc, 0x86, 0xb0, 0x4b,
	0xa8, 0x16, 0x3d, 0x63, 0x98, 0xbc, 0x61, 0x52, 0x8e, 0x29, 0xd7, 0x30, 0xb7, 0xb5, 0x4f, 0x2f,
	0xc2, 0x57, 0x92, 0xd8, 0x8c, 0x13, 0x9d, 0xe8, 0xa4, 0xc5, 0x87, 0x24, 0xb5, 0x6e, 0x53, 0x9b,
	0xc6, 0xf1, 0xf0, 0x2b, 0x89, 0x3e, 0xcd, 0x50, 0x83, 0x29, 0xe3, 0x56, 0x87, 0x12, 0xd3, 0x41,
	0x2e, 0x49, 0x70, 0x6a, 0x06, 0xce, 0x43, 0x0c, 0xe1, 0x84, 0xbe, 0xf2, 0x0b, 0xc0, 0x7b, 0x6d,
	0x6e, 0x1f, 0x78, 0x3d, 0x24, 0xac, 0xbd, 0x28, 0x23, 0xbd, 0x84, 0x05, 0xe4, 0x0b, 0x87, 0x32,
	0x57, 0x1c, 0x97, 0x40, 0x19, 0x54, 0x0b, 0x7a, 0xe9, 0xf7, 0x69, 0x7d, 0x3d, 0xd1, 0xb5, 0xdb,
	0xeb, 0x31, 0x8b, 0xf3, 0x7d, 0xc1, 0x5c, 0x62, 0x1b, 0x97, 0x50, 0xe9, 0x0d, 0x5c, 0x8e, 0xb9,
	0x4b, 0xf9, 0x32, 0xa8, 0xae, 0x34, 0xe5, 0xc6, 0x55, 0x5b, 0x1a, 0x71, 0x0f, 0xbd, 0x70, 0x16,
	0xa8, 0xb9, 0x9f, 0xe3, 0x41, 0x0d, 0x18, 0x49, 0xd1, 0xce, 0xab, 0x6f, 0xe3, 0x41, 0xed, 0x92,
	0xee, 0xfb, 0x78, 0x50, 0x7b, 0x3c, 0x95, 0x7f, 0x34, 0x33, 0xc0, 0x82, 0xde, 0xca, 0x26, 0xdc,
	0x58, 0x08, 0x19, 0x16, 0xf7, 0x28, 0xe1, 0x56, 0xe5, 0x34, 0x0f, 0xd5, 0x36, 0xb7, 0xdf, 0x63,
	0x8f, 0x32, 0xd1, 0x0e, 0x0d, 0x7a, 0xdb, 0x47, 0x2e, 0x46, 0xdd, 0xbe, 0xb5, 0x6b, 0x9a, 0xd4,
	0x27, 0xe2, 0xf6, 0xe3, 0x32, 0xf8, 0x20, 0xb6, 0x1c, 0xc5, 0x4c, 0x1d, 0x2e, 0x90, 0xb0, 0x92,
	0xd9, 0x9f, 0x64, 0xcd, 0x1e, 0x09, 0x48, 0xfa, 0xee, 0x87, 0x60, 0x7d, 0x2b, 0xb4, 0x61, 0x12,
	0xa8, 0x59, 0x4c, 0xc6, 0x1a, 0x5e, 0xc4, 0x4b, 0x07, 0xb0, 0x94, 0x81, 0xec, 0x38, 0x88, 0x3b,
	0xa5, 0xa5, 0x32, 0xa8, 0x16, 0xf5, 0xed, 0x49, 0xa0, 0x5e, 0x8b, 0x31, 0x1e, 0x5e, 0xa1, 0x7c,
	0x87, 0xb8, 0xb3, 0xb3, 0x3a, 0x6f, 0x7d, 0xe5, 0x07, 0x80, 0xcf, 0xfe, 0x61, 0x5b, 0x6a, 0xb1,
	0x54, 0x87, 0x70, 0x46, 0x04, 0x88, 0x44, 0xac, 0x4e, 0x02, 0x75, 0x26, 0x6a, 0x14, 0x78, 0xda,
	0x4a, 0x6a, 0xc1, 0x22, 0xf1, 0x71, 0xaa, 0x2d, 0x5e, 0x95, 0x3b, 0xfa, 0xfd, 0x49, 0xa0, 0xce,
	0xc5, 0x8d, 0x15, 0xe2, 0xe3, 0xb4, 0x57, 0xf3, 0x6b, 0x1e, 0x2e, 0xb5, 0xb9, 0x2d, 0x7d, 0x84,
	0xc5, 0xb9, 0x4d, 0x7d, 0x94, 0xe9, 0xf2, 0xfc, 0x2e, 0xc8, 0xcf, 0xff, 0x03, 0x34, 0x9d, 0xe6,
	0x04, 0xc0, 0xed, 0x1b, 0xb7, 0xa5, 0x75, 0x0d, 0xdb, 0x4d, 0x45, 0xf2, 0xeb, 0x5b, 0x14, 0xa5,
	0x92, 0xe4, 0xbb, 0x5f, 0xc2, 0x6b, 0xa2, 0xef, 0x9d, 0x0d, 0x15, 0x70, 0x3e, 0x54, 0xc0, 0xc5,
	0x50, 0x01, 0x7f, 0x86, 0x0a, 0x38, 0x19, 0x29, 0xb9, 0xf3, 0x91, 0x92, 0xbb, 0x18, 0x29, 0xb9,
	0x0f, 0x4d, 0xdb, 0x15, 0x8e, 0xdf, 0x6d, 0x98, 0x14, 0x6b, 0x61, 0xaf, 0x3a, 0xb1, 0xc4, 0x67,
	0xca, 0x0e, 0xb5, 0xcc, 0x1b, 0x24, 0x8e, 0x3d, 0x8b, 0x77, 0x97, 0xa3, 0x5f, 0x40, 0xeb, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x2d, 0xb8, 0x70, 0xdb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	ImportMorseClaimableAccounts(ctx context.Context, in *MsgImportMorseClaimableAccounts, opts ...grpc.CallOption) (*MsgImportMorseClaimableAccountsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/poktroll.migration.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ImportMorseClaimableAccounts(ctx context.Context, in *MsgImportMorseClaimableAccounts, opts ...grpc.CallOption) (*MsgImportMorseClaimableAccountsResponse, error) {
	out := new(MsgImportMorseClaimableAccountsResponse)
	err := c.cc.Invoke(ctx, "/poktroll.migration.Msg/ImportMorseClaimableAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	ImportMorseClaimableAccounts(context.Context, *MsgImportMorseClaimableAccounts) (*MsgImportMorseClaimableAccountsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) ImportMorseClaimableAccounts(ctx context.Context, req *MsgImportMorseClaimableAccounts) (*MsgImportMorseClaimableAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportMorseClaimableAccounts not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.migration.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ImportMorseClaimableAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgImportMorseClaimableAccounts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ImportMorseClaimableAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.migration.Msg/ImportMorseClaimableAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ImportMorseClaimableAccounts(ctx, req.(*MsgImportMorseClaimableAccounts))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poktroll.migration.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "ImportMorseClaimableAccounts",
			Handler:    _Msg_ImportMorseClaimableAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poktroll/migration/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgImportMorseClaimableAccounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgImportMorseClaimableAccounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgImportMorseClaimableAccounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MorseAccountStateHash) > 0 {
		i -= len(m.MorseAccountStateHash)
		copy(dAtA[i:], m.MorseAccountStateHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MorseAccountStateHash)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.MorseAccountState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgImportMorseClaimableAccountsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgImportMorseClaimableAccountsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgImportMorseClaimableAccountsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumAccounts != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NumAccounts))
		i--
		dAtA[i] = 0x10
	}
	if len(m.StateHash) > 0 {
		i -= len(m.StateHash)
		copy(dAtA[i:], m.StateHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StateHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgImportMorseClaimableAccounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MorseAccountState.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.MorseAccountStateHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgImportMorseClaimableAccountsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StateHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.NumAccounts != 0 {
		n += 1 + sovTx(uint64(m.NumAccounts))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgImportMorseClaimableAccounts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgImportMorseClaimableAccounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgImportMorseClaimableAccounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MorseAccountState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MorseAccountState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MorseAccountStateHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MorseAccountStateHash = append(m.MorseAccountStateHash[:0], dAtA[iNdEx:postIndex]...)
			if m.MorseAccountStateHash == nil {
				m.MorseAccountStateHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgImportMorseClaimableAccountsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgImportMorseClaimableAccountsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgImportMorseClaimableAccountsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateHash = append(m.StateHash[:0], dAtA[iNdEx:postIndex]...)
			if m.StateHash == nil {
				m.StateHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumAccounts", wireType)
			}
			m.NumAccounts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumAccounts |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
