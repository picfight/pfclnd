// Code generated by protoc-gen-go. DO NOT EDIT.
// source: walletrpc/walletkit.proto

package walletrpc // import "github.com/picfight/pfclnd/lnrpc/walletrpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import signrpc "github.com/picfight/pfclnd/lnrpc/signrpc"

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

type KeyReq struct {
	// *
	// Is the key finger print of the root pubkey that this request is targeting.
	// This allows the WalletKit to possibly serve out keys for multiple HD chains
	// via public derivation.
	KeyFingerPrint int32 `protobuf:"varint,1,opt,name=key_finger_print,json=keyFingerPrint,proto3" json:"key_finger_print,omitempty"`
	// *
	// The target key family to derive a key from. In other contexts, this is
	// known as the "account".
	KeyFamily            int32    `protobuf:"varint,2,opt,name=key_family,json=keyFamily,proto3" json:"key_family,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyReq) Reset()         { *m = KeyReq{} }
func (m *KeyReq) String() string { return proto.CompactTextString(m) }
func (*KeyReq) ProtoMessage()    {}
func (*KeyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{0}
}
func (m *KeyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyReq.Unmarshal(m, b)
}
func (m *KeyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyReq.Marshal(b, m, deterministic)
}
func (dst *KeyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyReq.Merge(dst, src)
}
func (m *KeyReq) XXX_Size() int {
	return xxx_messageInfo_KeyReq.Size(m)
}
func (m *KeyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyReq.DiscardUnknown(m)
}

var xxx_messageInfo_KeyReq proto.InternalMessageInfo

func (m *KeyReq) GetKeyFingerPrint() int32 {
	if m != nil {
		return m.KeyFingerPrint
	}
	return 0
}

func (m *KeyReq) GetKeyFamily() int32 {
	if m != nil {
		return m.KeyFamily
	}
	return 0
}

type AddrRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddrRequest) Reset()         { *m = AddrRequest{} }
func (m *AddrRequest) String() string { return proto.CompactTextString(m) }
func (*AddrRequest) ProtoMessage()    {}
func (*AddrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{1}
}
func (m *AddrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddrRequest.Unmarshal(m, b)
}
func (m *AddrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddrRequest.Marshal(b, m, deterministic)
}
func (dst *AddrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrRequest.Merge(dst, src)
}
func (m *AddrRequest) XXX_Size() int {
	return xxx_messageInfo_AddrRequest.Size(m)
}
func (m *AddrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddrRequest proto.InternalMessageInfo

type AddrResponse struct {
	// *
	// The address encoded using a bech32 format.
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddrResponse) Reset()         { *m = AddrResponse{} }
func (m *AddrResponse) String() string { return proto.CompactTextString(m) }
func (*AddrResponse) ProtoMessage()    {}
func (*AddrResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{2}
}
func (m *AddrResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddrResponse.Unmarshal(m, b)
}
func (m *AddrResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddrResponse.Marshal(b, m, deterministic)
}
func (dst *AddrResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrResponse.Merge(dst, src)
}
func (m *AddrResponse) XXX_Size() int {
	return xxx_messageInfo_AddrResponse.Size(m)
}
func (m *AddrResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddrResponse proto.InternalMessageInfo

func (m *AddrResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Transaction struct {
	// *
	// The raw serialized transaction.
	TxHex                []byte   `protobuf:"bytes,1,opt,name=tx_hex,json=txHex,proto3" json:"tx_hex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{3}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTxHex() []byte {
	if m != nil {
		return m.TxHex
	}
	return nil
}

type PublishResponse struct {
	// *
	// If blank, then no error occurred and the transaction was successfully
	// published. If not the empty string, then a string representation of the
	// broadcast error.
	//
	// TODO(roasbeef): map to a proper enum type
	PublishError         string   `protobuf:"bytes,1,opt,name=publish_error,json=publishError,proto3" json:"publish_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{4}
}
func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (dst *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(dst, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func (m *PublishResponse) GetPublishError() string {
	if m != nil {
		return m.PublishError
	}
	return ""
}

type SendOutputsRequest struct {
	// *
	// The number of satoshis per kilo weight that should be used when crafting
	// this transaction.
	SatPerKw int64 `protobuf:"varint,1,opt,name=sat_per_kw,json=satPerKw,proto3" json:"sat_per_kw,omitempty"`
	// *
	// A slice of the outputs that should be created in the transaction produced.
	Outputs              []*signrpc.TxOut `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SendOutputsRequest) Reset()         { *m = SendOutputsRequest{} }
func (m *SendOutputsRequest) String() string { return proto.CompactTextString(m) }
func (*SendOutputsRequest) ProtoMessage()    {}
func (*SendOutputsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{5}
}
func (m *SendOutputsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendOutputsRequest.Unmarshal(m, b)
}
func (m *SendOutputsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendOutputsRequest.Marshal(b, m, deterministic)
}
func (dst *SendOutputsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendOutputsRequest.Merge(dst, src)
}
func (m *SendOutputsRequest) XXX_Size() int {
	return xxx_messageInfo_SendOutputsRequest.Size(m)
}
func (m *SendOutputsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendOutputsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendOutputsRequest proto.InternalMessageInfo

func (m *SendOutputsRequest) GetSatPerKw() int64 {
	if m != nil {
		return m.SatPerKw
	}
	return 0
}

func (m *SendOutputsRequest) GetOutputs() []*signrpc.TxOut {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type SendOutputsResponse struct {
	// *
	// The serialized transaction sent out on the network.
	RawTx                []byte   `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendOutputsResponse) Reset()         { *m = SendOutputsResponse{} }
func (m *SendOutputsResponse) String() string { return proto.CompactTextString(m) }
func (*SendOutputsResponse) ProtoMessage()    {}
func (*SendOutputsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{6}
}
func (m *SendOutputsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendOutputsResponse.Unmarshal(m, b)
}
func (m *SendOutputsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendOutputsResponse.Marshal(b, m, deterministic)
}
func (dst *SendOutputsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendOutputsResponse.Merge(dst, src)
}
func (m *SendOutputsResponse) XXX_Size() int {
	return xxx_messageInfo_SendOutputsResponse.Size(m)
}
func (m *SendOutputsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendOutputsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendOutputsResponse proto.InternalMessageInfo

func (m *SendOutputsResponse) GetRawTx() []byte {
	if m != nil {
		return m.RawTx
	}
	return nil
}

type EstimateFeeRequest struct {
	// *
	// The number of confirmations to shoot for when estimating the fee.
	ConfTarget           int32    `protobuf:"varint,1,opt,name=conf_target,json=confTarget,proto3" json:"conf_target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimateFeeRequest) Reset()         { *m = EstimateFeeRequest{} }
func (m *EstimateFeeRequest) String() string { return proto.CompactTextString(m) }
func (*EstimateFeeRequest) ProtoMessage()    {}
func (*EstimateFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{7}
}
func (m *EstimateFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateFeeRequest.Unmarshal(m, b)
}
func (m *EstimateFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateFeeRequest.Marshal(b, m, deterministic)
}
func (dst *EstimateFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateFeeRequest.Merge(dst, src)
}
func (m *EstimateFeeRequest) XXX_Size() int {
	return xxx_messageInfo_EstimateFeeRequest.Size(m)
}
func (m *EstimateFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateFeeRequest proto.InternalMessageInfo

func (m *EstimateFeeRequest) GetConfTarget() int32 {
	if m != nil {
		return m.ConfTarget
	}
	return 0
}

type EstimateFeeResponse struct {
	// *
	// The amount of satoshis per kw that should be used in order to reach the
	// confirmation target in the request.
	SatPerKw             int64    `protobuf:"varint,1,opt,name=sat_per_kw,json=satPerKw,proto3" json:"sat_per_kw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimateFeeResponse) Reset()         { *m = EstimateFeeResponse{} }
func (m *EstimateFeeResponse) String() string { return proto.CompactTextString(m) }
func (*EstimateFeeResponse) ProtoMessage()    {}
func (*EstimateFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_walletkit_ca4e27c2068154e3, []int{8}
}
func (m *EstimateFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateFeeResponse.Unmarshal(m, b)
}
func (m *EstimateFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateFeeResponse.Marshal(b, m, deterministic)
}
func (dst *EstimateFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateFeeResponse.Merge(dst, src)
}
func (m *EstimateFeeResponse) XXX_Size() int {
	return xxx_messageInfo_EstimateFeeResponse.Size(m)
}
func (m *EstimateFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateFeeResponse proto.InternalMessageInfo

func (m *EstimateFeeResponse) GetSatPerKw() int64 {
	if m != nil {
		return m.SatPerKw
	}
	return 0
}

func init() {
	proto.RegisterType((*KeyReq)(nil), "walletrpc.KeyReq")
	proto.RegisterType((*AddrRequest)(nil), "walletrpc.AddrRequest")
	proto.RegisterType((*AddrResponse)(nil), "walletrpc.AddrResponse")
	proto.RegisterType((*Transaction)(nil), "walletrpc.Transaction")
	proto.RegisterType((*PublishResponse)(nil), "walletrpc.PublishResponse")
	proto.RegisterType((*SendOutputsRequest)(nil), "walletrpc.SendOutputsRequest")
	proto.RegisterType((*SendOutputsResponse)(nil), "walletrpc.SendOutputsResponse")
	proto.RegisterType((*EstimateFeeRequest)(nil), "walletrpc.EstimateFeeRequest")
	proto.RegisterType((*EstimateFeeResponse)(nil), "walletrpc.EstimateFeeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletKitClient is the client API for WalletKit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletKitClient interface {
	// *
	// DeriveNextKey attempts to derive the *next* key within the key family
	// (account in BIP43) specified. This method should return the next external
	// child within this branch.
	DeriveNextKey(ctx context.Context, in *KeyReq, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error)
	// *
	// DeriveKey attempts to derive an arbitrary key specified by the passed
	// KeyLocator.
	DeriveKey(ctx context.Context, in *signrpc.KeyLocator, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error)
	// *
	// NextAddr returns the next unused address within the wallet.
	NextAddr(ctx context.Context, in *AddrRequest, opts ...grpc.CallOption) (*AddrResponse, error)
	// *
	// PublishTransaction attempts to publish the passed transaction to the
	// network. Once this returns without an error, the wallet will continually
	// attempt to re-broadcast the transaction on start up, until it enters the
	// chain.
	PublishTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PublishResponse, error)
	// *
	// SendOutputs is similar to the existing sendmany call in Picfightcoind, and
	// allows the caller to create a transaction that sends to several outputs at
	// once. This is ideal when wanting to batch create a set of transactions.
	SendOutputs(ctx context.Context, in *SendOutputsRequest, opts ...grpc.CallOption) (*SendOutputsResponse, error)
	// *
	// EstimateFee attempts to query the internal fee estimator of the wallet to
	// determine the fee (in sat/kw) to attach to a transaction in order to
	// achieve the confirmation target.
	EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error)
}

type walletKitClient struct {
	cc *grpc.ClientConn
}

func NewWalletKitClient(cc *grpc.ClientConn) WalletKitClient {
	return &walletKitClient{cc}
}

func (c *walletKitClient) DeriveNextKey(ctx context.Context, in *KeyReq, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error) {
	out := new(signrpc.KeyDescriptor)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/DeriveNextKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) DeriveKey(ctx context.Context, in *signrpc.KeyLocator, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error) {
	out := new(signrpc.KeyDescriptor)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/DeriveKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) NextAddr(ctx context.Context, in *AddrRequest, opts ...grpc.CallOption) (*AddrResponse, error) {
	out := new(AddrResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/NextAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) PublishTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/PublishTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) SendOutputs(ctx context.Context, in *SendOutputsRequest, opts ...grpc.CallOption) (*SendOutputsResponse, error) {
	out := new(SendOutputsResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/SendOutputs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error) {
	out := new(EstimateFeeResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/EstimateFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKitServer is the server API for WalletKit service.
type WalletKitServer interface {
	// *
	// DeriveNextKey attempts to derive the *next* key within the key family
	// (account in BIP43) specified. This method should return the next external
	// child within this branch.
	DeriveNextKey(context.Context, *KeyReq) (*signrpc.KeyDescriptor, error)
	// *
	// DeriveKey attempts to derive an arbitrary key specified by the passed
	// KeyLocator.
	DeriveKey(context.Context, *signrpc.KeyLocator) (*signrpc.KeyDescriptor, error)
	// *
	// NextAddr returns the next unused address within the wallet.
	NextAddr(context.Context, *AddrRequest) (*AddrResponse, error)
	// *
	// PublishTransaction attempts to publish the passed transaction to the
	// network. Once this returns without an error, the wallet will continually
	// attempt to re-broadcast the transaction on start up, until it enters the
	// chain.
	PublishTransaction(context.Context, *Transaction) (*PublishResponse, error)
	// *
	// SendOutputs is similar to the existing sendmany call in Picfightcoind, and
	// allows the caller to create a transaction that sends to several outputs at
	// once. This is ideal when wanting to batch create a set of transactions.
	SendOutputs(context.Context, *SendOutputsRequest) (*SendOutputsResponse, error)
	// *
	// EstimateFee attempts to query the internal fee estimator of the wallet to
	// determine the fee (in sat/kw) to attach to a transaction in order to
	// achieve the confirmation target.
	EstimateFee(context.Context, *EstimateFeeRequest) (*EstimateFeeResponse, error)
}

func RegisterWalletKitServer(s *grpc.Server, srv WalletKitServer) {
	s.RegisterService(&_WalletKit_serviceDesc, srv)
}

func _WalletKit_DeriveNextKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).DeriveNextKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/DeriveNextKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).DeriveNextKey(ctx, req.(*KeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_DeriveKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(signrpc.KeyLocator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).DeriveKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/DeriveKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).DeriveKey(ctx, req.(*signrpc.KeyLocator))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_NextAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).NextAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/NextAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).NextAddr(ctx, req.(*AddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_PublishTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).PublishTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/PublishTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).PublishTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_SendOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).SendOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/SendOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).SendOutputs(ctx, req.(*SendOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_EstimateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).EstimateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/EstimateFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).EstimateFee(ctx, req.(*EstimateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WalletKit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "walletrpc.WalletKit",
	HandlerType: (*WalletKitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeriveNextKey",
			Handler:    _WalletKit_DeriveNextKey_Handler,
		},
		{
			MethodName: "DeriveKey",
			Handler:    _WalletKit_DeriveKey_Handler,
		},
		{
			MethodName: "NextAddr",
			Handler:    _WalletKit_NextAddr_Handler,
		},
		{
			MethodName: "PublishTransaction",
			Handler:    _WalletKit_PublishTransaction_Handler,
		},
		{
			MethodName: "SendOutputs",
			Handler:    _WalletKit_SendOutputs_Handler,
		},
		{
			MethodName: "EstimateFee",
			Handler:    _WalletKit_EstimateFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "walletrpc/walletkit.proto",
}

func init() {
	proto.RegisterFile("walletrpc/walletkit.proto", fileDescriptor_walletkit_ca4e27c2068154e3)
}

var fileDescriptor_walletkit_ca4e27c2068154e3 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xd5, 0x56, 0x56, 0xd6, 0xdb, 0x76, 0x80, 0xcb, 0x46, 0x89, 0x18, 0x4c, 0x81, 0x87, 0x3e,
	0xa0, 0x54, 0x6c, 0x02, 0x21, 0x78, 0x02, 0x6d, 0xd3, 0xa4, 0x4e, 0xac, 0x84, 0x4a, 0x48, 0x08,
	0x29, 0x72, 0x93, 0xbb, 0xd4, 0x6a, 0x6a, 0x67, 0xce, 0x0d, 0x49, 0xfe, 0x0f, 0x3f, 0x14, 0xe5,
	0xa3, 0x5d, 0x4a, 0x19, 0x4f, 0x71, 0x8e, 0xcf, 0x3d, 0xbe, 0xd7, 0xe7, 0x18, 0x9e, 0x26, 0x3c,
	0x08, 0x90, 0x74, 0xe8, 0x0e, 0xcb, 0xd5, 0x5c, 0x90, 0x15, 0x6a, 0x45, 0x8a, 0xb5, 0x56, 0x5b,
	0xc6, 0xe3, 0x48, 0xf8, 0x32, 0xe7, 0xe4, 0x5f, 0xd4, 0x25, 0xc1, 0xfc, 0x0a, 0xcd, 0x11, 0x66,
	0x36, 0xde, 0xb0, 0x01, 0x3c, 0x9c, 0x63, 0xe6, 0x5c, 0x0b, 0xe9, 0xa3, 0x76, 0x42, 0x2d, 0x24,
	0xf5, 0xb7, 0x8e, 0xb6, 0x06, 0x3b, 0xf6, 0xde, 0x1c, 0xb3, 0xf3, 0x02, 0x1e, 0xe7, 0x28, 0x3b,
	0x04, 0x28, 0x98, 0x7c, 0x21, 0x82, 0xac, 0xbf, 0x5d, 0x70, 0x5a, 0x39, 0xa7, 0x00, 0xcc, 0x2e,
	0xb4, 0x3f, 0x79, 0x9e, 0xb6, 0xf1, 0x26, 0xc6, 0x88, 0x4c, 0x13, 0x3a, 0xe5, 0x6f, 0x14, 0x2a,
	0x19, 0x21, 0x63, 0x70, 0x8f, 0x7b, 0x9e, 0x2e, 0xb4, 0x5b, 0x76, 0xb1, 0x36, 0x5f, 0x41, 0x7b,
	0xa2, 0xb9, 0x8c, 0xb8, 0x4b, 0x42, 0x49, 0xb6, 0x0f, 0x4d, 0x4a, 0x9d, 0x19, 0xa6, 0x05, 0xa9,
	0x63, 0xef, 0x50, 0x7a, 0x81, 0xa9, 0xf9, 0x0e, 0x1e, 0x8c, 0xe3, 0x69, 0x20, 0xa2, 0xd9, 0x4a,
	0xec, 0x25, 0x74, 0xc3, 0x12, 0x72, 0x50, 0x6b, 0xb5, 0x54, 0xed, 0x54, 0xe0, 0x59, 0x8e, 0x99,
	0x3f, 0x81, 0x7d, 0x43, 0xe9, 0x5d, 0xc5, 0x14, 0xc6, 0x14, 0x55, 0x7d, 0xb1, 0x67, 0x00, 0x11,
	0x27, 0x27, 0x44, 0xed, 0xcc, 0x93, 0xa2, 0xae, 0x61, 0xef, 0x46, 0x9c, 0xc6, 0xa8, 0x47, 0x09,
	0x1b, 0xc0, 0x7d, 0x55, 0xf2, 0xfb, 0xdb, 0x47, 0x8d, 0x41, 0xfb, 0x78, 0xcf, 0xaa, 0xee, 0xcf,
	0x9a, 0xa4, 0x57, 0x31, 0xd9, 0xcb, 0x6d, 0xf3, 0x35, 0xf4, 0xd6, 0xd4, 0xab, 0xce, 0xf6, 0xa1,
	0xa9, 0x79, 0xe2, 0xd0, 0x6a, 0x06, 0xcd, 0x93, 0x49, 0x6a, 0xbe, 0x05, 0x76, 0x16, 0x91, 0x58,
	0x70, 0xc2, 0x73, 0xc4, 0x65, 0x2f, 0x2f, 0xa0, 0xed, 0x2a, 0x79, 0xed, 0x10, 0xd7, 0x3e, 0x2e,
	0xaf, 0x1d, 0x72, 0x68, 0x52, 0x20, 0xe6, 0x09, 0xf4, 0xd6, 0xca, 0xaa, 0x43, 0xfe, 0x3b, 0xc3,
	0xf1, 0xef, 0x06, 0xb4, 0xbe, 0x17, 0xfe, 0x8f, 0x04, 0xb1, 0x0f, 0xd0, 0x3d, 0x45, 0x2d, 0x7e,
	0xe1, 0x17, 0x4c, 0x69, 0x84, 0x19, 0x7b, 0x64, 0xad, 0xc2, 0x61, 0x95, 0x19, 0x30, 0x0e, 0x56,
	0x43, 0x8e, 0x30, 0x3b, 0xc5, 0xc8, 0xd5, 0x22, 0x24, 0xa5, 0xd9, 0x7b, 0x68, 0x95, 0xb5, 0x79,
	0x5d, 0xaf, 0x4e, 0xba, 0x54, 0x2e, 0x27, 0xa5, 0xef, 0xac, 0xfc, 0x08, 0xbb, 0xf9, 0x79, 0x79,
	0x02, 0xd8, 0x41, 0xed, 0xc0, 0x5a, 0x42, 0x8c, 0x27, 0x1b, 0x78, 0x35, 0xde, 0x05, 0xb0, 0xca,
	0xf0, 0x7a, 0x3a, 0xea, 0x32, 0x35, 0xdc, 0x30, 0x6a, 0xf8, 0xdf, 0x39, 0xb9, 0x84, 0x76, 0xcd,
	0x24, 0x76, 0x58, 0xa3, 0x6e, 0x46, 0xc3, 0x78, 0x7e, 0xd7, 0xf6, 0xad, 0x5a, 0xcd, 0x8d, 0x35,
	0xb5, 0x4d, 0x73, 0xd7, 0xd4, 0xfe, 0x61, 0xe2, 0xe7, 0x37, 0x3f, 0x86, 0xbe, 0xa0, 0x59, 0x3c,
	0xb5, 0x5c, 0xb5, 0x18, 0x06, 0xc2, 0x9f, 0x91, 0x14, 0xd2, 0x97, 0x48, 0x89, 0xd2, 0xf3, 0x61,
	0x20, 0xbd, 0x61, 0x20, 0x6f, 0x1f, 0xb7, 0x0e, 0xdd, 0x69, 0xb3, 0x78, 0xbc, 0x27, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x2d, 0xbb, 0xcd, 0x97, 0xfa, 0x03, 0x00, 0x00,
}
