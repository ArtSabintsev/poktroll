// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pocket/gateway/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_918c541b0abffe4e, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_918c541b0abffe4e, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetGatewayRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryGetGatewayRequest) Reset()         { *m = QueryGetGatewayRequest{} }
func (m *QueryGetGatewayRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetGatewayRequest) ProtoMessage()    {}
func (*QueryGetGatewayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_918c541b0abffe4e, []int{2}
}
func (m *QueryGetGatewayRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetGatewayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryGetGatewayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetGatewayRequest.Merge(m, src)
}
func (m *QueryGetGatewayRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetGatewayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetGatewayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetGatewayRequest proto.InternalMessageInfo

func (m *QueryGetGatewayRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryGetGatewayResponse struct {
	Gateway Gateway `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway"`
}

func (m *QueryGetGatewayResponse) Reset()         { *m = QueryGetGatewayResponse{} }
func (m *QueryGetGatewayResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetGatewayResponse) ProtoMessage()    {}
func (*QueryGetGatewayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_918c541b0abffe4e, []int{3}
}
func (m *QueryGetGatewayResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetGatewayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryGetGatewayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetGatewayResponse.Merge(m, src)
}
func (m *QueryGetGatewayResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetGatewayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetGatewayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetGatewayResponse proto.InternalMessageInfo

func (m *QueryGetGatewayResponse) GetGateway() Gateway {
	if m != nil {
		return m.Gateway
	}
	return Gateway{}
}

type QueryAllGatewaysRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllGatewaysRequest) Reset()         { *m = QueryAllGatewaysRequest{} }
func (m *QueryAllGatewaysRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllGatewaysRequest) ProtoMessage()    {}
func (*QueryAllGatewaysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_918c541b0abffe4e, []int{4}
}
func (m *QueryAllGatewaysRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllGatewaysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllGatewaysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllGatewaysRequest.Merge(m, src)
}
func (m *QueryAllGatewaysRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllGatewaysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllGatewaysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllGatewaysRequest proto.InternalMessageInfo

func (m *QueryAllGatewaysRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllGatewaysResponse struct {
	Gateways   []Gateway           `protobuf:"bytes,1,rep,name=gateways,proto3" json:"gateways"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllGatewaysResponse) Reset()         { *m = QueryAllGatewaysResponse{} }
func (m *QueryAllGatewaysResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllGatewaysResponse) ProtoMessage()    {}
func (*QueryAllGatewaysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_918c541b0abffe4e, []int{5}
}
func (m *QueryAllGatewaysResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllGatewaysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *QueryAllGatewaysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllGatewaysResponse.Merge(m, src)
}
func (m *QueryAllGatewaysResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllGatewaysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllGatewaysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllGatewaysResponse proto.InternalMessageInfo

func (m *QueryAllGatewaysResponse) GetGateways() []Gateway {
	if m != nil {
		return m.Gateways
	}
	return nil
}

func (m *QueryAllGatewaysResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "pocket.gateway.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "pocket.gateway.QueryParamsResponse")
	proto.RegisterType((*QueryGetGatewayRequest)(nil), "pocket.gateway.QueryGetGatewayRequest")
	proto.RegisterType((*QueryGetGatewayResponse)(nil), "pocket.gateway.QueryGetGatewayResponse")
	proto.RegisterType((*QueryAllGatewaysRequest)(nil), "pocket.gateway.QueryAllGatewaysRequest")
	proto.RegisterType((*QueryAllGatewaysResponse)(nil), "pocket.gateway.QueryAllGatewaysResponse")
}

func init() { proto.RegisterFile("pocket/gateway/query.proto", fileDescriptor_918c541b0abffe4e) }

var fileDescriptor_918c541b0abffe4e = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x2d, 0x34, 0x74, 0x2b, 0x21, 0xb1, 0x54, 0x6d, 0x30, 0xc8, 0x20, 0x23, 0x92,
	0xa8, 0x52, 0xbd, 0x6d, 0x38, 0x00, 0x47, 0x72, 0x20, 0x37, 0x14, 0x7c, 0xe4, 0x82, 0x36, 0xe9,
	0xca, 0x58, 0x71, 0x3c, 0xae, 0xbd, 0xa1, 0x44, 0x08, 0x09, 0xf1, 0x04, 0x95, 0xe0, 0xcc, 0x99,
	0x23, 0x07, 0x1e, 0xa2, 0xc7, 0x4a, 0x5c, 0x7a, 0x42, 0x28, 0x41, 0xe2, 0x35, 0x90, 0x77, 0xc7,
	0x25, 0xb1, 0x0b, 0xe6, 0x92, 0x78, 0x76, 0xfe, 0x99, 0xf9, 0x66, 0x76, 0xb4, 0xd4, 0x8a, 0x61,
	0x38, 0x92, 0x8a, 0xfb, 0x42, 0xc9, 0x23, 0x31, 0xe5, 0x87, 0x13, 0x99, 0x4c, 0xdd, 0x38, 0x01,
	0x05, 0xec, 0xaa, 0xf1, 0xb9, 0xe8, 0xb3, 0xae, 0x89, 0x71, 0x10, 0x01, 0xd7, 0xbf, 0x46, 0x62,
	0xdd, 0x18, 0x42, 0x3a, 0x86, 0xf4, 0x85, 0xb6, 0xb8, 0x31, 0xd0, 0xb5, 0xe9, 0x83, 0x0f, 0xe6,
	0x3c, 0xfb, 0xc2, 0xd3, 0x5b, 0x3e, 0x80, 0x1f, 0x4a, 0x2e, 0xe2, 0x80, 0x8b, 0x28, 0x02, 0x25,
	0x54, 0x00, 0x51, 0x1e, 0xb3, 0x63, 0x32, 0xf0, 0x81, 0x48, 0xa5, 0x41, 0xe1, 0xaf, 0xf6, 0x07,
	0x52, 0x89, 0x7d, 0x1e, 0x0b, 0x3f, 0x88, 0xb4, 0x18, 0xb5, 0xf6, 0xa2, 0x36, 0x57, 0x0d, 0x21,
	0xc8, 0xfd, 0x37, 0x0b, 0x9d, 0xc5, 0x22, 0x11, 0xe3, 0xbc, 0x50, 0xb1, 0x6d, 0x35, 0x8d, 0x25,
	0xfa, 0x9c, 0x4d, 0xca, 0x9e, 0x65, 0xa5, 0xfb, 0x3a, 0xc0, 0x93, 0x87, 0x13, 0x99, 0x2a, 0xa7,
	0x4f, 0xaf, 0x2f, 0x9d, 0xa6, 0x31, 0x44, 0xa9, 0x64, 0x8f, 0xe8, 0x9a, 0x49, 0xdc, 0x20, 0x77,
	0x48, 0x7b, 0xa3, 0xb3, 0xe5, 0x2e, 0x0f, 0xcd, 0x35, 0xfa, 0xee, 0xfa, 0xc9, 0xf7, 0xdb, 0xb5,
	0xcf, 0xbf, 0xbe, 0xec, 0x10, 0x0f, 0x03, 0x9c, 0x0e, 0xdd, 0xd2, 0x19, 0x7b, 0x52, 0xf5, 0x8c,
	0x18, 0x6b, 0xb1, 0x06, 0xad, 0x8b, 0x83, 0x83, 0x44, 0xa6, 0x26, 0xeb, 0xba, 0x97, 0x9b, 0x8e,
	0x47, 0xb7, 0x4b, 0x31, 0x48, 0xf2, 0x80, 0xd6, 0xb1, 0x26, 0xa2, 0x6c, 0x17, 0x51, 0x30, 0xa2,
	0x7b, 0x29, 0x63, 0xf1, 0x72, 0xb5, 0x23, 0x30, 0xe7, 0xe3, 0x30, 0x44, 0x45, 0xde, 0x34, 0x7b,
	0x42, 0xe9, 0x9f, 0xb9, 0x63, 0xda, 0xa6, 0x8b, 0xd7, 0x9c, 0x0d, 0xde, 0x35, 0xfb, 0x82, 0xe3,
	0x77, 0xfb, 0xc2, 0x97, 0x18, 0xeb, 0x2d, 0x44, 0x3a, 0x9f, 0x08, 0x6d, 0x94, 0x6b, 0x9c, 0x8f,
	0xf0, 0x0a, 0xa2, 0x64, 0xed, 0xae, 0x56, 0x93, 0x9f, 0xcb, 0x59, 0x6f, 0x89, 0x6f, 0x45, 0xf3,
	0xb5, 0x2a, 0xf9, 0x4c, 0xdd, 0x45, 0xc0, 0xce, 0xd7, 0x55, 0x7a, 0x59, 0x03, 0xb2, 0x77, 0x84,
	0xae, 0x99, 0x3b, 0x63, 0x4e, 0x11, 0xa3, 0xbc, 0x16, 0xd6, 0xdd, 0x7f, 0x6a, 0x4c, 0x25, 0x67,
	0xf7, 0xfd, 0xb7, 0x9f, 0x1f, 0x56, 0x5a, 0xec, 0x1e, 0x8f, 0x61, 0xa4, 0x76, 0x23, 0xa9, 0x8e,
	0x20, 0x19, 0x69, 0x23, 0x81, 0x30, 0x2c, 0xac, 0x28, 0xfb, 0x48, 0x68, 0x1d, 0x3b, 0x66, 0xcd,
	0x0b, 0xf3, 0x97, 0x56, 0xc6, 0x6a, 0x55, 0xea, 0x90, 0xe5, 0xa1, 0x66, 0xe9, 0xb0, 0xbd, 0x0a,
	0x96, 0xfc, 0xff, 0x0d, 0xae, 0xde, 0x5b, 0x76, 0x4c, 0xe8, 0xc6, 0xc2, 0xfd, 0xb1, 0x8b, 0x4b,
	0x96, 0xb7, 0xc8, 0x6a, 0x57, 0x0b, 0x11, 0xce, 0xd5, 0x70, 0x6d, 0xd6, 0xfc, 0x3f, 0xb8, 0xee,
	0xd3, 0x93, 0x99, 0x4d, 0x4e, 0x67, 0x36, 0x39, 0x9b, 0xd9, 0xe4, 0xc7, 0xcc, 0x26, 0xc7, 0x73,
	0xbb, 0x76, 0x3a, 0xb7, 0x6b, 0x67, 0x73, 0xbb, 0xf6, 0x7c, 0xcf, 0x0f, 0xd4, 0xcb, 0xc9, 0xc0,
	0x1d, 0xc2, 0xf8, 0x2f, 0xf9, 0x5e, 0x2f, 0x3f, 0x00, 0x83, 0x35, 0xfd, 0x02, 0xdc, 0xff, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xf8, 0x29, 0x9a, 0x6b, 0x16, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Gateway items.
	Gateway(ctx context.Context, in *QueryGetGatewayRequest, opts ...grpc.CallOption) (*QueryGetGatewayResponse, error)
	AllGateways(ctx context.Context, in *QueryAllGatewaysRequest, opts ...grpc.CallOption) (*QueryAllGatewaysResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/pocket.gateway.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Gateway(ctx context.Context, in *QueryGetGatewayRequest, opts ...grpc.CallOption) (*QueryGetGatewayResponse, error) {
	out := new(QueryGetGatewayResponse)
	err := c.cc.Invoke(ctx, "/pocket.gateway.Query/Gateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllGateways(ctx context.Context, in *QueryAllGatewaysRequest, opts ...grpc.CallOption) (*QueryAllGatewaysResponse, error) {
	out := new(QueryAllGatewaysResponse)
	err := c.cc.Invoke(ctx, "/pocket.gateway.Query/AllGateways", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Gateway items.
	Gateway(context.Context, *QueryGetGatewayRequest) (*QueryGetGatewayResponse, error)
	AllGateways(context.Context, *QueryAllGatewaysRequest) (*QueryAllGatewaysResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Gateway(ctx context.Context, req *QueryGetGatewayRequest) (*QueryGetGatewayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gateway not implemented")
}
func (*UnimplementedQueryServer) AllGateways(ctx context.Context, req *QueryAllGatewaysRequest) (*QueryAllGatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllGateways not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pocket.gateway.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Gateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Gateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pocket.gateway.Query/Gateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Gateway(ctx, req.(*QueryGetGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllGateways_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllGateways(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pocket.gateway.Query/AllGateways",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllGateways(ctx, req.(*QueryAllGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pocket.gateway.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Gateway",
			Handler:    _Query_Gateway_Handler,
		},
		{
			MethodName: "AllGateways",
			Handler:    _Query_AllGateways_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pocket/gateway/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetGatewayRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetGatewayRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetGatewayRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetGatewayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetGatewayResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetGatewayResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Gateway.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllGatewaysRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllGatewaysRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllGatewaysRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllGatewaysResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllGatewaysResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllGatewaysResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Gateways) > 0 {
		for iNdEx := len(m.Gateways) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gateways[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetGatewayRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetGatewayResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Gateway.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllGatewaysRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllGatewaysResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Gateways) > 0 {
		for _, e := range m.Gateways {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetGatewayRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetGatewayRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetGatewayRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetGatewayResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetGatewayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetGatewayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateway", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Gateway.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllGatewaysRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllGatewaysRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllGatewaysRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllGatewaysResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllGatewaysResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllGatewaysResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateways", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gateways = append(m.Gateways, Gateway{})
			if err := m.Gateways[len(m.Gateways)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
