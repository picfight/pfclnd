// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilotrpc/autopilot.proto

package autopilotrpc // import "github.com/picfight/pfclnd/lnrpc/autopilotrpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_autopilot_45e6f1df6dc1d1df, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusResponse struct {
	// / Indicates whether the autopilot is active or not.
	Active               bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_autopilot_45e6f1df6dc1d1df, []int{1}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (dst *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(dst, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type ModifyStatusRequest struct {
	// / Whether the autopilot agent should be enabled or not.
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyStatusRequest) Reset()         { *m = ModifyStatusRequest{} }
func (m *ModifyStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyStatusRequest) ProtoMessage()    {}
func (*ModifyStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_autopilot_45e6f1df6dc1d1df, []int{2}
}
func (m *ModifyStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyStatusRequest.Unmarshal(m, b)
}
func (m *ModifyStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyStatusRequest.Marshal(b, m, deterministic)
}
func (dst *ModifyStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyStatusRequest.Merge(dst, src)
}
func (m *ModifyStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyStatusRequest.Size(m)
}
func (m *ModifyStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyStatusRequest proto.InternalMessageInfo

func (m *ModifyStatusRequest) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type ModifyStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyStatusResponse) Reset()         { *m = ModifyStatusResponse{} }
func (m *ModifyStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyStatusResponse) ProtoMessage()    {}
func (*ModifyStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_autopilot_45e6f1df6dc1d1df, []int{3}
}
func (m *ModifyStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyStatusResponse.Unmarshal(m, b)
}
func (m *ModifyStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyStatusResponse.Marshal(b, m, deterministic)
}
func (dst *ModifyStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyStatusResponse.Merge(dst, src)
}
func (m *ModifyStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyStatusResponse.Size(m)
}
func (m *ModifyStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyStatusResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StatusRequest)(nil), "autopilotrpc.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "autopilotrpc.StatusResponse")
	proto.RegisterType((*ModifyStatusRequest)(nil), "autopilotrpc.ModifyStatusRequest")
	proto.RegisterType((*ModifyStatusResponse)(nil), "autopilotrpc.ModifyStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AutopilotClient is the client API for Autopilot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutopilotClient interface {
	// *
	// Status returns whether the daemon's autopilot agent is active.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// *
	// ModifyStatus is used to modify the status of the autopilot agent, like
	// enabling or disabling it.
	ModifyStatus(ctx context.Context, in *ModifyStatusRequest, opts ...grpc.CallOption) (*ModifyStatusResponse, error)
}

type autopilotClient struct {
	cc *grpc.ClientConn
}

func NewAutopilotClient(cc *grpc.ClientConn) AutopilotClient {
	return &autopilotClient{cc}
}

func (c *autopilotClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/autopilotrpc.Autopilot/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autopilotClient) ModifyStatus(ctx context.Context, in *ModifyStatusRequest, opts ...grpc.CallOption) (*ModifyStatusResponse, error) {
	out := new(ModifyStatusResponse)
	err := c.cc.Invoke(ctx, "/autopilotrpc.Autopilot/ModifyStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutopilotServer is the server API for Autopilot service.
type AutopilotServer interface {
	// *
	// Status returns whether the daemon's autopilot agent is active.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// *
	// ModifyStatus is used to modify the status of the autopilot agent, like
	// enabling or disabling it.
	ModifyStatus(context.Context, *ModifyStatusRequest) (*ModifyStatusResponse, error)
}

func RegisterAutopilotServer(s *grpc.Server, srv AutopilotServer) {
	s.RegisterService(&_Autopilot_serviceDesc, srv)
}

func _Autopilot_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopilotServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopilotrpc.Autopilot/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopilotServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Autopilot_ModifyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutopilotServer).ModifyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/autopilotrpc.Autopilot/ModifyStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutopilotServer).ModifyStatus(ctx, req.(*ModifyStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Autopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "autopilotrpc.Autopilot",
	HandlerType: (*AutopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Autopilot_Status_Handler,
		},
		{
			MethodName: "ModifyStatus",
			Handler:    _Autopilot_ModifyStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "autopilotrpc/autopilot.proto",
}

func init() {
	proto.RegisterFile("autopilotrpc/autopilot.proto", fileDescriptor_autopilot_45e6f1df6dc1d1df)
}

var fileDescriptor_autopilot_45e6f1df6dc1d1df = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2c, 0x2d, 0xc9,
	0x2f, 0xc8, 0xcc, 0xc9, 0x2f, 0x29, 0x2a, 0x48, 0xd6, 0x87, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x78, 0x90, 0x65, 0x95, 0xf8, 0xb9, 0x78, 0x83, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x34, 0xb8, 0xf8, 0x60, 0x02, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0xc9, 0x25, 0x99, 0x65, 0xa9, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x1c, 0x41, 0x50, 0x9e, 0x92, 0x2e, 0x97, 0xb0, 0x6f, 0x7e, 0x4a, 0x66, 0x5a, 0x25, 0x8a,
	0x01, 0x20, 0xe5, 0xa9, 0x79, 0x89, 0x49, 0x39, 0x70, 0xe5, 0x10, 0x9e, 0x92, 0x18, 0x97, 0x08,
	0xaa, 0x72, 0x88, 0xf1, 0x46, 0xcb, 0x19, 0xb9, 0x38, 0x1d, 0x61, 0x4e, 0x12, 0x72, 0xe6, 0x62,
	0x83, 0xc8, 0x0b, 0x49, 0xeb, 0x21, 0x3b, 0x54, 0x0f, 0xc5, 0x12, 0x29, 0x19, 0xec, 0x92, 0x50,
	0x17, 0x87, 0x72, 0xf1, 0x20, 0x5b, 0x25, 0xa4, 0x88, 0xaa, 0x1a, 0x8b, 0xab, 0xa5, 0x94, 0xf0,
	0x29, 0x81, 0x18, 0xeb, 0x64, 0x12, 0x65, 0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x9f, 0x93, 0x99, 0x9e, 0x51, 0x92, 0x97, 0x99, 0x97, 0x9e, 0x97, 0x5a, 0x52, 0x9e,
	0x5f, 0x94, 0xad, 0x9f, 0x93, 0x97, 0xa2, 0x9f, 0x93, 0x87, 0x12, 0xe4, 0x45, 0x05, 0xc9, 0x49,
	0x6c, 0xe0, 0x60, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x71, 0xb3, 0xba, 0x96, 0x01,
	0x00, 0x00,
}
