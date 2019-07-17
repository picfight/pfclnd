// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signrpc/signer.proto

package signrpc // import "github.com/picfight/pfclnd/lnrpc/signrpc"

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

type KeyLocator struct {
	// / The family of key being identified.
	KeyFamily int32 `protobuf:"varint,1,opt,name=key_family,json=keyFamily,proto3" json:"key_family,omitempty"`
	// / The precise index of the key being identified.
	KeyIndex             int32    `protobuf:"varint,2,opt,name=key_index,json=keyIndex,proto3" json:"key_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyLocator) Reset()         { *m = KeyLocator{} }
func (m *KeyLocator) String() string { return proto.CompactTextString(m) }
func (*KeyLocator) ProtoMessage()    {}
func (*KeyLocator) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{0}
}
func (m *KeyLocator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyLocator.Unmarshal(m, b)
}
func (m *KeyLocator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyLocator.Marshal(b, m, deterministic)
}
func (dst *KeyLocator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyLocator.Merge(dst, src)
}
func (m *KeyLocator) XXX_Size() int {
	return xxx_messageInfo_KeyLocator.Size(m)
}
func (m *KeyLocator) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyLocator.DiscardUnknown(m)
}

var xxx_messageInfo_KeyLocator proto.InternalMessageInfo

func (m *KeyLocator) GetKeyFamily() int32 {
	if m != nil {
		return m.KeyFamily
	}
	return 0
}

func (m *KeyLocator) GetKeyIndex() int32 {
	if m != nil {
		return m.KeyIndex
	}
	return 0
}

type KeyDescriptor struct {
	// *
	// The raw bytes of the key being identified. Either this or the KeyLocator
	// must be specified.
	RawKeyBytes []byte `protobuf:"bytes,1,opt,name=raw_key_bytes,json=rawKeyBytes,proto3" json:"raw_key_bytes,omitempty"`
	// *
	// The key locator that identifies which key to use for signing. Either this
	// or the raw bytes of the target key must be specified.
	KeyLoc               *KeyLocator `protobuf:"bytes,2,opt,name=key_loc,json=keyLoc,proto3" json:"key_loc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *KeyDescriptor) Reset()         { *m = KeyDescriptor{} }
func (m *KeyDescriptor) String() string { return proto.CompactTextString(m) }
func (*KeyDescriptor) ProtoMessage()    {}
func (*KeyDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{1}
}
func (m *KeyDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyDescriptor.Unmarshal(m, b)
}
func (m *KeyDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyDescriptor.Marshal(b, m, deterministic)
}
func (dst *KeyDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDescriptor.Merge(dst, src)
}
func (m *KeyDescriptor) XXX_Size() int {
	return xxx_messageInfo_KeyDescriptor.Size(m)
}
func (m *KeyDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDescriptor proto.InternalMessageInfo

func (m *KeyDescriptor) GetRawKeyBytes() []byte {
	if m != nil {
		return m.RawKeyBytes
	}
	return nil
}

func (m *KeyDescriptor) GetKeyLoc() *KeyLocator {
	if m != nil {
		return m.KeyLoc
	}
	return nil
}

type TxOut struct {
	// / The value of the output being spent.
	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	// / The script of the output being spent.
	PkScript             []byte   `protobuf:"bytes,2,opt,name=pk_script,json=pkScript,proto3" json:"pk_script,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxOut) Reset()         { *m = TxOut{} }
func (m *TxOut) String() string { return proto.CompactTextString(m) }
func (*TxOut) ProtoMessage()    {}
func (*TxOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{2}
}
func (m *TxOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxOut.Unmarshal(m, b)
}
func (m *TxOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxOut.Marshal(b, m, deterministic)
}
func (dst *TxOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxOut.Merge(dst, src)
}
func (m *TxOut) XXX_Size() int {
	return xxx_messageInfo_TxOut.Size(m)
}
func (m *TxOut) XXX_DiscardUnknown() {
	xxx_messageInfo_TxOut.DiscardUnknown(m)
}

var xxx_messageInfo_TxOut proto.InternalMessageInfo

func (m *TxOut) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TxOut) GetPkScript() []byte {
	if m != nil {
		return m.PkScript
	}
	return nil
}

type SignDescriptor struct {
	// *
	// A descriptor that precisely describes *which* key to use for signing. This
	// may provide the raw public key directly, or require the Signer to re-derive
	// the key according to the populated derivation path.
	KeyDesc *KeyDescriptor `protobuf:"bytes,1,opt,name=key_desc,json=keyDesc,proto3" json:"key_desc,omitempty"`
	// *
	// A scalar value that will be added to the private key corresponding to the
	// above public key to obtain the private key to be used to sign this input.
	// This value is typically derived via the following computation:
	//
	// derivedKey = privkey + sha256(perCommitmentPoint || pubKey) mod N
	SingleTweak []byte `protobuf:"bytes,2,opt,name=single_tweak,json=singleTweak,proto3" json:"single_tweak,omitempty"`
	// *
	// A private key that will be used in combination with its corresponding
	// private key to derive the private key that is to be used to sign the target
	// input. Within the Lightning protocol, this value is typically the
	// commitment secret from a previously revoked commitment transaction. This
	// value is in combination with two hash values, and the original private key
	// to derive the private key to be used when signing.
	//
	// k = (privKey*sha256(pubKey || tweakPub) +
	// tweakPriv*sha256(tweakPub || pubKey)) mod N
	DoubleTweak []byte `protobuf:"bytes,3,opt,name=double_tweak,json=doubleTweak,proto3" json:"double_tweak,omitempty"`
	// *
	// The full script required to properly redeem the output.  This field will
	// only be populated if a p2wsh or a p2sh output is being signed.
	WitnessScript []byte `protobuf:"bytes,4,opt,name=witness_script,json=witnessScript,proto3" json:"witness_script,omitempty"`
	// *
	// A description of the output being spent. The value and script MUST be provided.
	Output *TxOut `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	// *
	// The target sighash type that should be used when generating the final
	// sighash, and signature.
	Sighash uint32 `protobuf:"varint,7,opt,name=sighash,proto3" json:"sighash,omitempty"`
	// *
	// The target input within the transaction that should be signed.
	InputIndex           int32    `protobuf:"varint,8,opt,name=input_index,json=inputIndex,proto3" json:"input_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignDescriptor) Reset()         { *m = SignDescriptor{} }
func (m *SignDescriptor) String() string { return proto.CompactTextString(m) }
func (*SignDescriptor) ProtoMessage()    {}
func (*SignDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{3}
}
func (m *SignDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignDescriptor.Unmarshal(m, b)
}
func (m *SignDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignDescriptor.Marshal(b, m, deterministic)
}
func (dst *SignDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignDescriptor.Merge(dst, src)
}
func (m *SignDescriptor) XXX_Size() int {
	return xxx_messageInfo_SignDescriptor.Size(m)
}
func (m *SignDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SignDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SignDescriptor proto.InternalMessageInfo

func (m *SignDescriptor) GetKeyDesc() *KeyDescriptor {
	if m != nil {
		return m.KeyDesc
	}
	return nil
}

func (m *SignDescriptor) GetSingleTweak() []byte {
	if m != nil {
		return m.SingleTweak
	}
	return nil
}

func (m *SignDescriptor) GetDoubleTweak() []byte {
	if m != nil {
		return m.DoubleTweak
	}
	return nil
}

func (m *SignDescriptor) GetWitnessScript() []byte {
	if m != nil {
		return m.WitnessScript
	}
	return nil
}

func (m *SignDescriptor) GetOutput() *TxOut {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *SignDescriptor) GetSighash() uint32 {
	if m != nil {
		return m.Sighash
	}
	return 0
}

func (m *SignDescriptor) GetInputIndex() int32 {
	if m != nil {
		return m.InputIndex
	}
	return 0
}

type SignReq struct {
	// / The raw bytes of the transaction to be signed.
	RawTxBytes []byte `protobuf:"bytes,1,opt,name=raw_tx_bytes,json=rawTxBytes,proto3" json:"raw_tx_bytes,omitempty"`
	// / A set of sign descriptors, for each input to be signed.
	SignDescs            []*SignDescriptor `protobuf:"bytes,2,rep,name=sign_descs,json=signDescs,proto3" json:"sign_descs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SignReq) Reset()         { *m = SignReq{} }
func (m *SignReq) String() string { return proto.CompactTextString(m) }
func (*SignReq) ProtoMessage()    {}
func (*SignReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{4}
}
func (m *SignReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignReq.Unmarshal(m, b)
}
func (m *SignReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignReq.Marshal(b, m, deterministic)
}
func (dst *SignReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignReq.Merge(dst, src)
}
func (m *SignReq) XXX_Size() int {
	return xxx_messageInfo_SignReq.Size(m)
}
func (m *SignReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignReq proto.InternalMessageInfo

func (m *SignReq) GetRawTxBytes() []byte {
	if m != nil {
		return m.RawTxBytes
	}
	return nil
}

func (m *SignReq) GetSignDescs() []*SignDescriptor {
	if m != nil {
		return m.SignDescs
	}
	return nil
}

type SignResp struct {
	// *
	// A set of signatures realized in a fixed 64-byte format ordered in ascending
	// input order.
	RawSigs              [][]byte `protobuf:"bytes,1,rep,name=raw_sigs,json=rawSigs,proto3" json:"raw_sigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignResp) Reset()         { *m = SignResp{} }
func (m *SignResp) String() string { return proto.CompactTextString(m) }
func (*SignResp) ProtoMessage()    {}
func (*SignResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{5}
}
func (m *SignResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignResp.Unmarshal(m, b)
}
func (m *SignResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignResp.Marshal(b, m, deterministic)
}
func (dst *SignResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResp.Merge(dst, src)
}
func (m *SignResp) XXX_Size() int {
	return xxx_messageInfo_SignResp.Size(m)
}
func (m *SignResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignResp proto.InternalMessageInfo

func (m *SignResp) GetRawSigs() [][]byte {
	if m != nil {
		return m.RawSigs
	}
	return nil
}

type InputScript struct {
	// / The serializes witness stack for the specified input.
	Witness [][]byte `protobuf:"bytes,1,rep,name=witness,proto3" json:"witness,omitempty"`
	// **
	// The optional sig script for the specified witness that will only be set if
	// the input specified is a nested p2sh witness program.
	SigScript            []byte   `protobuf:"bytes,2,opt,name=sig_script,json=sigScript,proto3" json:"sig_script,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputScript) Reset()         { *m = InputScript{} }
func (m *InputScript) String() string { return proto.CompactTextString(m) }
func (*InputScript) ProtoMessage()    {}
func (*InputScript) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{6}
}
func (m *InputScript) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputScript.Unmarshal(m, b)
}
func (m *InputScript) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputScript.Marshal(b, m, deterministic)
}
func (dst *InputScript) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputScript.Merge(dst, src)
}
func (m *InputScript) XXX_Size() int {
	return xxx_messageInfo_InputScript.Size(m)
}
func (m *InputScript) XXX_DiscardUnknown() {
	xxx_messageInfo_InputScript.DiscardUnknown(m)
}

var xxx_messageInfo_InputScript proto.InternalMessageInfo

func (m *InputScript) GetWitness() [][]byte {
	if m != nil {
		return m.Witness
	}
	return nil
}

func (m *InputScript) GetSigScript() []byte {
	if m != nil {
		return m.SigScript
	}
	return nil
}

type InputScriptResp struct {
	// / The set of fully valid input scripts requested.
	InputScripts         []*InputScript `protobuf:"bytes,1,rep,name=input_scripts,json=inputScripts,proto3" json:"input_scripts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InputScriptResp) Reset()         { *m = InputScriptResp{} }
func (m *InputScriptResp) String() string { return proto.CompactTextString(m) }
func (*InputScriptResp) ProtoMessage()    {}
func (*InputScriptResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_signer_8dc402df3ea9a6a7, []int{7}
}
func (m *InputScriptResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputScriptResp.Unmarshal(m, b)
}
func (m *InputScriptResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputScriptResp.Marshal(b, m, deterministic)
}
func (dst *InputScriptResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputScriptResp.Merge(dst, src)
}
func (m *InputScriptResp) XXX_Size() int {
	return xxx_messageInfo_InputScriptResp.Size(m)
}
func (m *InputScriptResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InputScriptResp.DiscardUnknown(m)
}

var xxx_messageInfo_InputScriptResp proto.InternalMessageInfo

func (m *InputScriptResp) GetInputScripts() []*InputScript {
	if m != nil {
		return m.InputScripts
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyLocator)(nil), "signrpc.KeyLocator")
	proto.RegisterType((*KeyDescriptor)(nil), "signrpc.KeyDescriptor")
	proto.RegisterType((*TxOut)(nil), "signrpc.TxOut")
	proto.RegisterType((*SignDescriptor)(nil), "signrpc.SignDescriptor")
	proto.RegisterType((*SignReq)(nil), "signrpc.SignReq")
	proto.RegisterType((*SignResp)(nil), "signrpc.SignResp")
	proto.RegisterType((*InputScript)(nil), "signrpc.InputScript")
	proto.RegisterType((*InputScriptResp)(nil), "signrpc.InputScriptResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignerClient is the client API for Signer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignerClient interface {
	// *
	// SignOutputRaw is a method that can be used to generated a signature for a
	// set of inputs/outputs to a transaction. Each request specifies details
	// concerning how the outputs should be signed, which keys they should be
	// signed with, and also any optional tweaks. The return value is a fixed
	// 64-byte signature (the same format as we use on the wire in Lightning).
	//
	// If we are  unable to sign using the specified keys, then an error will be
	// returned.
	SignOutputRaw(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignResp, error)
	// *
	// ComputeInputScript generates a complete InputIndex for the passed
	// transaction with the signature as defined within the passed SignDescriptor.
	// This method should be capable of generating the proper input script for
	// both regular p2wkh output and p2wkh outputs nested within a regular p2sh
	// output.
	//
	// Note that when using this method to sign inputs belonging to the wallet,
	// the only items of the SignDescriptor that need to be populated are pkScript
	// in the TxOut field, the value in that same field, and finally the input
	// index.
	ComputeInputScript(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*InputScriptResp, error)
}

type signerClient struct {
	cc *grpc.ClientConn
}

func NewSignerClient(cc *grpc.ClientConn) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) SignOutputRaw(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignResp, error) {
	out := new(SignResp)
	err := c.cc.Invoke(ctx, "/signrpc.Signer/SignOutputRaw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signerClient) ComputeInputScript(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*InputScriptResp, error) {
	out := new(InputScriptResp)
	err := c.cc.Invoke(ctx, "/signrpc.Signer/ComputeInputScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignerServer is the server API for Signer service.
type SignerServer interface {
	// *
	// SignOutputRaw is a method that can be used to generated a signature for a
	// set of inputs/outputs to a transaction. Each request specifies details
	// concerning how the outputs should be signed, which keys they should be
	// signed with, and also any optional tweaks. The return value is a fixed
	// 64-byte signature (the same format as we use on the wire in Lightning).
	//
	// If we are  unable to sign using the specified keys, then an error will be
	// returned.
	SignOutputRaw(context.Context, *SignReq) (*SignResp, error)
	// *
	// ComputeInputScript generates a complete InputIndex for the passed
	// transaction with the signature as defined within the passed SignDescriptor.
	// This method should be capable of generating the proper input script for
	// both regular p2wkh output and p2wkh outputs nested within a regular p2sh
	// output.
	//
	// Note that when using this method to sign inputs belonging to the wallet,
	// the only items of the SignDescriptor that need to be populated are pkScript
	// in the TxOut field, the value in that same field, and finally the input
	// index.
	ComputeInputScript(context.Context, *SignReq) (*InputScriptResp, error)
}

func RegisterSignerServer(s *grpc.Server, srv SignerServer) {
	s.RegisterService(&_Signer_serviceDesc, srv)
}

func _Signer_SignOutputRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).SignOutputRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signrpc.Signer/SignOutputRaw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).SignOutputRaw(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signer_ComputeInputScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).ComputeInputScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/signrpc.Signer/ComputeInputScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).ComputeInputScript(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "signrpc.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignOutputRaw",
			Handler:    _Signer_SignOutputRaw_Handler,
		},
		{
			MethodName: "ComputeInputScript",
			Handler:    _Signer_ComputeInputScript_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signrpc/signer.proto",
}

func init() { proto.RegisterFile("signrpc/signer.proto", fileDescriptor_signer_8dc402df3ea9a6a7) }

var fileDescriptor_signer_8dc402df3ea9a6a7 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4f, 0x8f, 0xd3, 0x3e,
	0x10, 0x55, 0xb7, 0xbf, 0x36, 0xdd, 0x49, 0xd2, 0x1f, 0x98, 0x0a, 0x02, 0x08, 0x51, 0x22, 0x2d,
	0xea, 0x01, 0x5a, 0x51, 0x10, 0x12, 0x9c, 0xd0, 0x82, 0x56, 0xac, 0xba, 0xd2, 0x4a, 0x6e, 0x4f,
	0x5c, 0xa2, 0x34, 0x35, 0xa9, 0x95, 0x34, 0xf1, 0xc6, 0x0e, 0x69, 0x6e, 0x7c, 0x07, 0xbe, 0x30,
	0x1a, 0x3b, 0xfd, 0x07, 0x9c, 0x9a, 0xf7, 0x3c, 0x33, 0xef, 0x79, 0x5e, 0x0d, 0x03, 0xc9, 0xe3,
	0xac, 0x10, 0xd1, 0x04, 0x7f, 0x59, 0x31, 0x16, 0x45, 0xae, 0x72, 0x62, 0x35, 0xac, 0xff, 0x15,
	0x60, 0xc6, 0xea, 0x9b, 0x3c, 0x0a, 0x55, 0x5e, 0x90, 0x67, 0x00, 0x09, 0xab, 0x83, 0xef, 0xe1,
	0x86, 0xa7, 0xb5, 0xd7, 0x1a, 0xb6, 0x46, 0x1d, 0x7a, 0x9e, 0xb0, 0xfa, 0x4a, 0x13, 0xe4, 0x29,
	0x20, 0x08, 0x78, 0xb6, 0x62, 0x5b, 0xef, 0x4c, 0x9f, 0xf6, 0x12, 0x56, 0x5f, 0x23, 0xf6, 0x43,
	0x70, 0x67, 0xac, 0xfe, 0xc2, 0x64, 0x54, 0x70, 0x81, 0xc3, 0x7c, 0x70, 0x8b, 0xb0, 0x0a, 0xb0,
	0x63, 0x59, 0x2b, 0x26, 0xf5, 0x3c, 0x87, 0xda, 0x45, 0x58, 0xcd, 0x58, 0x7d, 0x89, 0x14, 0x79,
	0x05, 0x16, 0x9e, 0xa7, 0x79, 0xa4, 0xe7, 0xd9, 0xd3, 0x07, 0xe3, 0xc6, 0xd9, 0xf8, 0x60, 0x8b,
	0x76, 0x13, 0xfd, 0xed, 0x7f, 0x84, 0xce, 0x62, 0x7b, 0x5b, 0x2a, 0x32, 0x80, 0xce, 0x8f, 0x30,
	0x2d, 0x99, 0x1e, 0xd9, 0xa6, 0x06, 0xa0, 0x3d, 0x91, 0x04, 0x46, 0x5f, 0x8f, 0x73, 0x68, 0x4f,
	0x24, 0x73, 0x8d, 0xfd, 0x5f, 0x67, 0xd0, 0x9f, 0xf3, 0x38, 0x3b, 0x32, 0xf8, 0x06, 0xd0, 0x7d,
	0xb0, 0x62, 0x32, 0xd2, 0x83, 0xec, 0xe9, 0xc3, 0x63, 0xf5, 0x43, 0x25, 0x45, 0x93, 0x08, 0xc9,
	0x0b, 0x70, 0x24, 0xcf, 0xe2, 0x94, 0x05, 0xaa, 0x62, 0x61, 0xd2, 0xa8, 0xd8, 0x86, 0x5b, 0x20,
	0x85, 0x25, 0xab, 0xbc, 0x5c, 0xee, 0x4b, 0xda, 0xa6, 0xc4, 0x70, 0xa6, 0xe4, 0x02, 0xfa, 0x15,
	0x57, 0x19, 0x93, 0x72, 0xe7, 0xf6, 0x3f, 0x5d, 0xe4, 0x36, 0xac, 0xb1, 0x4c, 0x5e, 0x42, 0x37,
	0x2f, 0x95, 0x28, 0x95, 0xd7, 0xd1, 0xee, 0xfa, 0x7b, 0x77, 0x7a, 0x0b, 0xb4, 0x39, 0x25, 0x1e,
	0x60, 0x9c, 0xeb, 0x50, 0xae, 0x3d, 0x6b, 0xd8, 0x1a, 0xb9, 0x74, 0x07, 0xc9, 0x73, 0xb0, 0x79,
	0x26, 0x4a, 0xd5, 0x44, 0xd6, 0xd3, 0x91, 0x81, 0xa6, 0x4c, 0x68, 0x11, 0x58, 0xb8, 0x14, 0xca,
	0xee, 0xc8, 0x10, 0x1c, 0x8c, 0x4b, 0x6d, 0x4f, 0xd2, 0x82, 0x22, 0xac, 0x16, 0x5b, 0x13, 0xd6,
	0x7b, 0x00, 0x34, 0xa0, 0x17, 0x26, 0xbd, 0xb3, 0x61, 0x7b, 0x64, 0x4f, 0x1f, 0xed, 0x3d, 0x9d,
	0x2e, 0x97, 0x9e, 0xcb, 0x06, 0x4b, 0xff, 0x02, 0x7a, 0x46, 0x44, 0x0a, 0xf2, 0x18, 0x7a, 0xa8,
	0x22, 0x79, 0x8c, 0x0a, 0xed, 0x91, 0x43, 0xad, 0x22, 0xac, 0xe6, 0x3c, 0x96, 0xfe, 0x15, 0xd8,
	0xd7, 0xe8, 0xac, 0xb9, 0xbd, 0x07, 0x56, 0xb3, 0x8e, 0x5d, 0x61, 0x03, 0xf1, 0x5f, 0x2a, 0x79,
	0x7c, 0x1a, 0x34, 0xca, 0x35, 0x49, 0xdf, 0xc0, 0xff, 0x47, 0x73, 0xb4, 0xea, 0x07, 0x70, 0xcd,
	0x1e, 0x4c, 0x8f, 0x99, 0x68, 0x4f, 0x07, 0x7b, 0xf3, 0xc7, 0x0d, 0x0e, 0x3f, 0x00, 0x39, 0xfd,
	0xd9, 0x82, 0xee, 0x5c, 0x3f, 0x1d, 0xf2, 0x0e, 0x5c, 0xfc, 0xba, 0xd5, 0x5b, 0xa7, 0x61, 0x45,
	0xee, 0x9d, 0x5c, 0x9e, 0xb2, 0xbb, 0x27, 0xf7, 0xff, 0x60, 0xa4, 0x20, 0x9f, 0x80, 0x7c, 0xce,
	0x37, 0xa2, 0x54, 0xec, 0xf8, 0x76, 0x7f, 0xb7, 0x7a, 0xff, 0x34, 0xc3, 0xa4, 0xb8, 0x9c, 0x7c,
	0x7b, 0x1d, 0x73, 0xb5, 0x2e, 0x97, 0xe3, 0x28, 0xdf, 0x4c, 0x52, 0x1e, 0xaf, 0x55, 0xc6, 0xb3,
	0x38, 0x63, 0xaa, 0xca, 0x8b, 0x64, 0x92, 0x66, 0xab, 0x49, 0xba, 0x7f, 0xe2, 0x85, 0x88, 0x96,
	0x5d, 0xfd, 0xc8, 0xdf, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x80, 0x01, 0xce, 0xe1, 0xfc, 0x03,
	0x00, 0x00,
}
