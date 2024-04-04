// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/mm2/v1/tx.proto

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

// MsgCreateMarkets defines a message carrying a payload for creating markets in
// the x/marketmap module.
type MsgCreateMarkets struct {
	// Authority is the signer of this transaction.  This authority must be
	// authorized by the module to execute the message.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// CreateMarkets is the list of all markets to be created for the given
	// transaction.
	CreateMarkets []Market `protobuf:"bytes,2,rep,name=create_markets,json=createMarkets,proto3" json:"create_markets"`
}

func (m *MsgCreateMarkets) Reset()         { *m = MsgCreateMarkets{} }
func (m *MsgCreateMarkets) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMarkets) ProtoMessage()    {}
func (*MsgCreateMarkets) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a86eaac7c34565b, []int{0}
}
func (m *MsgCreateMarkets) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMarkets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMarkets.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMarkets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMarkets.Merge(m, src)
}
func (m *MsgCreateMarkets) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMarkets) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMarkets.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMarkets proto.InternalMessageInfo

func (m *MsgCreateMarkets) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgCreateMarkets) GetCreateMarkets() []Market {
	if m != nil {
		return m.CreateMarkets
	}
	return nil
}

// MsgUpdateMarketMapResponse is the response message for MsgUpdateMarketMap.
type MsgCreateMarketsResponse struct {
}

func (m *MsgCreateMarketsResponse) Reset()         { *m = MsgCreateMarketsResponse{} }
func (m *MsgCreateMarketsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMarketsResponse) ProtoMessage()    {}
func (*MsgCreateMarketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a86eaac7c34565b, []int{1}
}
func (m *MsgCreateMarketsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMarketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMarketsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMarketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMarketsResponse.Merge(m, src)
}
func (m *MsgCreateMarketsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMarketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMarketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMarketsResponse proto.InternalMessageInfo

// MsgUpdateMarkets defines a message carrying a payload for updating the
// x/marketmap module.
type MsgUpdateMarkets struct {
	// Authority is the signer of this transaction.  This authority must be
	// authorized by the module to execute the message.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// UpdateMarkets is the list of all markets to be updated for the given
	// transaction.
	UpdateMarkets []Market `protobuf:"bytes,2,rep,name=update_markets,json=updateMarkets,proto3" json:"update_markets"`
}

func (m *MsgUpdateMarkets) Reset()         { *m = MsgUpdateMarkets{} }
func (m *MsgUpdateMarkets) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateMarkets) ProtoMessage()    {}
func (*MsgUpdateMarkets) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a86eaac7c34565b, []int{2}
}
func (m *MsgUpdateMarkets) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateMarkets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateMarkets.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateMarkets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateMarkets.Merge(m, src)
}
func (m *MsgUpdateMarkets) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateMarkets) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateMarkets.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateMarkets proto.InternalMessageInfo

func (m *MsgUpdateMarkets) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateMarkets) GetUpdateMarkets() []Market {
	if m != nil {
		return m.UpdateMarkets
	}
	return nil
}

// MsgUpdateMarketsResponse is the response message for MsgUpdateMarkets.
type MsgUpdateMarketsResponse struct {
}

func (m *MsgUpdateMarketsResponse) Reset()         { *m = MsgUpdateMarketsResponse{} }
func (m *MsgUpdateMarketsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateMarketsResponse) ProtoMessage()    {}
func (*MsgUpdateMarketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a86eaac7c34565b, []int{3}
}
func (m *MsgUpdateMarketsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateMarketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateMarketsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateMarketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateMarketsResponse.Merge(m, src)
}
func (m *MsgUpdateMarketsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateMarketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateMarketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateMarketsResponse proto.InternalMessageInfo

// MsgParams defines the Msg/Params request type. It contains the
// new parameters for the x/marketmap module.
type MsgParams struct {
	// Params defines the new parameters for the x/marketmap module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// Authority defines the authority that is updating the x/marketmap module
	// parameters.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgParams) Reset()         { *m = MsgParams{} }
func (m *MsgParams) String() string { return proto.CompactTextString(m) }
func (*MsgParams) ProtoMessage()    {}
func (*MsgParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a86eaac7c34565b, []int{4}
}
func (m *MsgParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgParams.Merge(m, src)
}
func (m *MsgParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgParams proto.InternalMessageInfo

func (m *MsgParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *MsgParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgParamsResponse defines the Msg/Params response type.
type MsgParamsResponse struct {
}

func (m *MsgParamsResponse) Reset()         { *m = MsgParamsResponse{} }
func (m *MsgParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgParamsResponse) ProtoMessage()    {}
func (*MsgParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a86eaac7c34565b, []int{5}
}
func (m *MsgParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgParamsResponse.Merge(m, src)
}
func (m *MsgParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateMarkets)(nil), "slinky.mm2.v1.MsgCreateMarkets")
	proto.RegisterType((*MsgCreateMarketsResponse)(nil), "slinky.mm2.v1.MsgCreateMarketsResponse")
	proto.RegisterType((*MsgUpdateMarkets)(nil), "slinky.mm2.v1.MsgUpdateMarkets")
	proto.RegisterType((*MsgUpdateMarketsResponse)(nil), "slinky.mm2.v1.MsgUpdateMarketsResponse")
	proto.RegisterType((*MsgParams)(nil), "slinky.mm2.v1.MsgParams")
	proto.RegisterType((*MsgParamsResponse)(nil), "slinky.mm2.v1.MsgParamsResponse")
}

func init() { proto.RegisterFile("slinky/mm2/v1/tx.proto", fileDescriptor_1a86eaac7c34565b) }

var fileDescriptor_1a86eaac7c34565b = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xce, 0xc9, 0xcc,
	0xcb, 0xae, 0xd4, 0xcf, 0xcd, 0x35, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x88, 0xeb, 0xe5, 0xe6, 0x1a, 0xe9, 0x95, 0x19, 0x4a, 0x89, 0x27,
	0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xeb, 0xe7, 0x16, 0xa7, 0x83, 0x94, 0xe5, 0x16, 0xa7, 0x43, 0xd4,
	0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x54, 0x54, 0x12, 0xa2, 0x3c,
	0x1e, 0x22, 0x01, 0xe1, 0x40, 0xa5, 0x04, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24, 0x54,
	0x48, 0x0a, 0xd5, 0x0d, 0xb9, 0x89, 0x45, 0xd9, 0xa9, 0x25, 0xd8, 0xe5, 0x0a, 0x12, 0x8b, 0x12,
	0x73, 0xa1, 0x46, 0x29, 0x9d, 0x62, 0xe4, 0x12, 0xf0, 0x2d, 0x4e, 0x77, 0x2e, 0x4a, 0x4d, 0x2c,
	0x49, 0xf5, 0x05, 0xeb, 0x2a, 0x16, 0x32, 0xe3, 0xe2, 0x4c, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca,
	0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x92, 0xb8, 0xb4, 0x45, 0x57, 0x04, 0xea,
	0x08, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0xf4, 0x20, 0x84,
	0x52, 0x21, 0x27, 0x2e, 0xbe, 0x64, 0xb0, 0x41, 0xf1, 0x10, 0xfb, 0x8b, 0x25, 0x98, 0x14, 0x98,
	0x35, 0xb8, 0x8d, 0x44, 0xf5, 0x50, 0x42, 0x42, 0x0f, 0x62, 0x8f, 0x13, 0xcb, 0x89, 0x7b, 0xf2,
	0x0c, 0x41, 0xbc, 0xc9, 0xc8, 0x76, 0x5b, 0x59, 0xbd, 0x58, 0x20, 0xcf, 0xd0, 0xf4, 0x7c, 0x83,
	0x16, 0xc2, 0xdc, 0xae, 0xe7, 0x1b, 0xb4, 0x94, 0xa1, 0x7e, 0xa8, 0x80, 0x7a, 0x2d, 0x37, 0xb1,
	0x40, 0x1f, 0xdd, 0xdd, 0x4a, 0x52, 0x5c, 0x12, 0xe8, 0x62, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x30, 0x8f, 0x86, 0x16, 0xa4, 0x50, 0xc7, 0xa3, 0xa5, 0x60, 0x83, 0x48, 0xf2, 0x68,
	0x29, 0xb2, 0xdd, 0x24, 0x7a, 0x14, 0xc5, 0xdd, 0x50, 0x8f, 0xa2, 0x88, 0xc1, 0x3d, 0xda, 0xc1,
	0xc8, 0xc5, 0xe9, 0x5b, 0x9c, 0x1e, 0x00, 0x8e, 0x65, 0x21, 0x63, 0x2e, 0x36, 0x48, 0x7c, 0x83,
	0xbd, 0x87, 0xe9, 0x42, 0x88, 0x32, 0xa8, 0x0b, 0xa1, 0x4a, 0x51, 0x83, 0x85, 0x89, 0xe8, 0x60,
	0xb1, 0xe2, 0x43, 0xf5, 0x8e, 0x92, 0x30, 0x97, 0x20, 0xdc, 0x25, 0x30, 0xf7, 0x19, 0x75, 0x30,
	0x71, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x45, 0x72, 0xf1, 0xa2, 0xa6, 0x3a, 0x79, 0xf4, 0xc0, 0x43,
	0x8b, 0x4a, 0x29, 0x75, 0x02, 0x0a, 0x60, 0x56, 0x80, 0x8c, 0x46, 0x8d, 0x67, 0x2c, 0x46, 0xa3,
	0x28, 0xc0, 0x66, 0x34, 0xd6, 0xd0, 0x15, 0x72, 0xe1, 0x62, 0x83, 0x86, 0xac, 0x04, 0xa6, 0x16,
	0x88, 0x8c, 0x94, 0x02, 0x2e, 0x19, 0x98, 0x29, 0x52, 0xac, 0x0d, 0xcf, 0x37, 0x68, 0x31, 0x3a,
	0x39, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5a, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x76, 0x66, 0x81, 0x6e, 0x6e, 0x6a, 0x99,
	0x3e, 0x22, 0x65, 0xe4, 0x1a, 0xe9, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x73, 0xb1,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x17, 0xf1, 0x96, 0x83, 0x04, 0x00, 0x00,
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
	// CreateMarkets creates markets from the given message.
	CreateMarkets(ctx context.Context, in *MsgCreateMarkets, opts ...grpc.CallOption) (*MsgCreateMarketsResponse, error)
	// UpdateMarkets updates markets from the given message.
	UpdateMarkets(ctx context.Context, in *MsgUpdateMarkets, opts ...grpc.CallOption) (*MsgUpdateMarketsResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateMarkets(ctx context.Context, in *MsgCreateMarkets, opts ...grpc.CallOption) (*MsgCreateMarketsResponse, error) {
	out := new(MsgCreateMarketsResponse)
	err := c.cc.Invoke(ctx, "/slinky.mm2.v1.Msg/CreateMarkets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateMarkets(ctx context.Context, in *MsgUpdateMarkets, opts ...grpc.CallOption) (*MsgUpdateMarketsResponse, error) {
	out := new(MsgUpdateMarketsResponse)
	err := c.cc.Invoke(ctx, "/slinky.mm2.v1.Msg/UpdateMarkets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Params(ctx context.Context, in *MsgParams, opts ...grpc.CallOption) (*MsgParamsResponse, error) {
	out := new(MsgParamsResponse)
	err := c.cc.Invoke(ctx, "/slinky.mm2.v1.Msg/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateMarkets creates markets from the given message.
	CreateMarkets(context.Context, *MsgCreateMarkets) (*MsgCreateMarketsResponse, error)
	// UpdateMarkets updates markets from the given message.
	UpdateMarkets(context.Context, *MsgUpdateMarkets) (*MsgUpdateMarketsResponse, error)
	// Params defines a method for updating the x/marketmap module parameters.
	Params(context.Context, *MsgParams) (*MsgParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateMarkets(ctx context.Context, req *MsgCreateMarkets) (*MsgCreateMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMarkets not implemented")
}
func (*UnimplementedMsgServer) UpdateMarkets(ctx context.Context, req *MsgUpdateMarkets) (*MsgUpdateMarketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarkets not implemented")
}
func (*UnimplementedMsgServer) Params(ctx context.Context, req *MsgParams) (*MsgParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMarkets)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.mm2.v1.Msg/CreateMarkets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMarkets(ctx, req.(*MsgCreateMarkets))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateMarkets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMarkets)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMarkets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.mm2.v1.Msg/UpdateMarkets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMarkets(ctx, req.(*MsgUpdateMarkets))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slinky.mm2.v1.Msg/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Params(ctx, req.(*MsgParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slinky.mm2.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMarkets",
			Handler:    _Msg_CreateMarkets_Handler,
		},
		{
			MethodName: "UpdateMarkets",
			Handler:    _Msg_UpdateMarkets_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Msg_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slinky/mm2/v1/tx.proto",
}

func (m *MsgCreateMarkets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMarkets) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMarkets) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreateMarkets) > 0 {
		for iNdEx := len(m.CreateMarkets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreateMarkets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateMarketsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMarketsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMarketsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateMarkets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateMarkets) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateMarkets) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UpdateMarkets) > 0 {
		for iNdEx := len(m.UpdateMarkets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UpdateMarkets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateMarketsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateMarketsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateMarketsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreateMarkets) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.CreateMarkets) > 0 {
		for _, e := range m.CreateMarkets {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateMarketsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateMarkets) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.UpdateMarkets) > 0 {
		for _, e := range m.UpdateMarkets {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdateMarketsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateMarkets) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMarkets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMarkets: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field CreateMarkets", wireType)
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
			m.CreateMarkets = append(m.CreateMarkets, Market{})
			if err := m.CreateMarkets[len(m.CreateMarkets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateMarketsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMarketsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMarketsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateMarkets) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateMarkets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateMarkets: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateMarkets", wireType)
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
			m.UpdateMarkets = append(m.UpdateMarkets, Market{})
			if err := m.UpdateMarkets[len(m.UpdateMarkets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateMarketsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateMarketsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateMarketsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
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
func (m *MsgParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
