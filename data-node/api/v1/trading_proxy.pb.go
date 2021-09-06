// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data-node/api/v1/trading_proxy.proto

package v1

import (
	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Blockchain transaction type
type SubmitTransactionRequest_Type int32

const (
	SubmitTransactionRequest_TYPE_UNSPECIFIED SubmitTransactionRequest_Type = 0
	// The transaction will be submitted without waiting for response
	SubmitTransactionRequest_TYPE_ASYNC SubmitTransactionRequest_Type = 1
	// The transaction will be submitted, and blocking until the
	// tendermint mempool return a response
	SubmitTransactionRequest_TYPE_SYNC SubmitTransactionRequest_Type = 2
	// The transaction will submitted, and blocking until the tendermint
	// network will have committed it into a block
	SubmitTransactionRequest_TYPE_COMMIT SubmitTransactionRequest_Type = 3
)

var SubmitTransactionRequest_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_ASYNC",
	2: "TYPE_SYNC",
	3: "TYPE_COMMIT",
}

var SubmitTransactionRequest_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_ASYNC":       1,
	"TYPE_SYNC":        2,
	"TYPE_COMMIT":      3,
}

func (x SubmitTransactionRequest_Type) String() string {
	return proto.EnumName(SubmitTransactionRequest_Type_name, int32(x))
}

func (SubmitTransactionRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c134d37351a048e1, []int{0, 0}
}

// Request for submitting a transaction V2 on Vega
type SubmitTransactionRequest struct {
	// A bundle of signed payload and signature, to form a transaction that will be submitted to the Vega blockchain
	Tx *v1.Transaction `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	// Type of transaction request, for example ASYNC, meaning the transaction will be submitted and not block on a response
	Type                 SubmitTransactionRequest_Type `protobuf:"varint,2,opt,name=type,proto3,enum=datanode.api.v1.SubmitTransactionRequest_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubmitTransactionRequest) Reset()         { *m = SubmitTransactionRequest{} }
func (m *SubmitTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionRequest) ProtoMessage()    {}
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c134d37351a048e1, []int{0}
}

func (m *SubmitTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionRequest.Unmarshal(m, b)
}
func (m *SubmitTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionRequest.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionRequest.Merge(m, src)
}
func (m *SubmitTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionRequest.Size(m)
}
func (m *SubmitTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionRequest proto.InternalMessageInfo

func (m *SubmitTransactionRequest) GetTx() *v1.Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *SubmitTransactionRequest) GetType() SubmitTransactionRequest_Type {
	if m != nil {
		return m.Type
	}
	return SubmitTransactionRequest_TYPE_UNSPECIFIED
}

// Response for submitting a transaction V2 on Vega
type SubmitTransactionResponse struct {
	// Success will be true if the transaction was accepted by the node,
	// **Important** - success does not mean that the event is confirmed by consensus
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitTransactionResponse) Reset()         { *m = SubmitTransactionResponse{} }
func (m *SubmitTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitTransactionResponse) ProtoMessage()    {}
func (*SubmitTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c134d37351a048e1, []int{1}
}

func (m *SubmitTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitTransactionResponse.Unmarshal(m, b)
}
func (m *SubmitTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitTransactionResponse.Marshal(b, m, deterministic)
}
func (m *SubmitTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitTransactionResponse.Merge(m, src)
}
func (m *SubmitTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitTransactionResponse.Size(m)
}
func (m *SubmitTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitTransactionResponse proto.InternalMessageInfo

func (m *SubmitTransactionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// A request to get the height of the very last block processed
// by tendermint
type LastBlockHeightRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastBlockHeightRequest) Reset()         { *m = LastBlockHeightRequest{} }
func (m *LastBlockHeightRequest) String() string { return proto.CompactTextString(m) }
func (*LastBlockHeightRequest) ProtoMessage()    {}
func (*LastBlockHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c134d37351a048e1, []int{2}
}

func (m *LastBlockHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastBlockHeightRequest.Unmarshal(m, b)
}
func (m *LastBlockHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastBlockHeightRequest.Marshal(b, m, deterministic)
}
func (m *LastBlockHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastBlockHeightRequest.Merge(m, src)
}
func (m *LastBlockHeightRequest) XXX_Size() int {
	return xxx_messageInfo_LastBlockHeightRequest.Size(m)
}
func (m *LastBlockHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LastBlockHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LastBlockHeightRequest proto.InternalMessageInfo

// A response with the height of the last block processed by
// tendermint
type LastBlockHeightResponse struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastBlockHeightResponse) Reset()         { *m = LastBlockHeightResponse{} }
func (m *LastBlockHeightResponse) String() string { return proto.CompactTextString(m) }
func (*LastBlockHeightResponse) ProtoMessage()    {}
func (*LastBlockHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c134d37351a048e1, []int{3}
}

func (m *LastBlockHeightResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastBlockHeightResponse.Unmarshal(m, b)
}
func (m *LastBlockHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastBlockHeightResponse.Marshal(b, m, deterministic)
}
func (m *LastBlockHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastBlockHeightResponse.Merge(m, src)
}
func (m *LastBlockHeightResponse) XXX_Size() int {
	return xxx_messageInfo_LastBlockHeightResponse.Size(m)
}
func (m *LastBlockHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LastBlockHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LastBlockHeightResponse proto.InternalMessageInfo

func (m *LastBlockHeightResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("datanode.api.v1.SubmitTransactionRequest_Type", SubmitTransactionRequest_Type_name, SubmitTransactionRequest_Type_value)
	proto.RegisterType((*SubmitTransactionRequest)(nil), "datanode.api.v1.SubmitTransactionRequest")
	proto.RegisterType((*SubmitTransactionResponse)(nil), "datanode.api.v1.SubmitTransactionResponse")
	proto.RegisterType((*LastBlockHeightRequest)(nil), "datanode.api.v1.LastBlockHeightRequest")
	proto.RegisterType((*LastBlockHeightResponse)(nil), "datanode.api.v1.LastBlockHeightResponse")
}

func init() {
	proto.RegisterFile("data-node/api/v1/trading_proxy.proto", fileDescriptor_c134d37351a048e1)
}

var fileDescriptor_c134d37351a048e1 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x61, 0xaf, 0xd2, 0x30,
	0x14, 0x75, 0x93, 0x3c, 0xf5, 0xbe, 0xf8, 0x98, 0xd5, 0x3c, 0xe7, 0x12, 0xe3, 0xcb, 0x62, 0x22,
	0x1a, 0x5f, 0x97, 0x61, 0xfc, 0x01, 0x82, 0x18, 0x49, 0x1e, 0x48, 0xb6, 0xf9, 0x01, 0xbf, 0x90,
	0xd2, 0x35, 0x50, 0x85, 0x75, 0xae, 0x65, 0x81, 0xdf, 0xea, 0x0f, 0xf0, 0x6f, 0x98, 0x76, 0x43,
	0x0d, 0xc3, 0xc8, 0xb7, 0x9e, 0xde, 0x73, 0xee, 0x39, 0xed, 0xbd, 0xf0, 0x3c, 0x25, 0x8a, 0x5c,
	0x67, 0x22, 0x65, 0x01, 0xc9, 0x79, 0x50, 0x86, 0x81, 0x2a, 0x48, 0xca, 0xb3, 0xc5, 0x2c, 0x2f,
	0xc4, 0x76, 0x87, 0xf3, 0x42, 0x28, 0x81, 0xda, 0x9a, 0xa5, 0x49, 0x98, 0xe4, 0x1c, 0x97, 0xa1,
	0xf7, 0xac, 0x64, 0x0b, 0x12, 0x50, 0xb1, 0x5e, 0x93, 0x2c, 0x95, 0x5a, 0xb6, 0x3f, 0x57, 0x0a,
	0xcf, 0x6f, 0x10, 0x54, 0x41, 0x32, 0x49, 0xa8, 0xe2, 0x22, 0xab, 0x38, 0xfe, 0x0f, 0x0b, 0xdc,
	0x78, 0x33, 0x5f, 0x73, 0x95, 0xfc, 0xa9, 0x45, 0xec, 0xfb, 0x86, 0x49, 0x85, 0xae, 0xc1, 0x56,
	0x5b, 0xd7, 0xba, 0xb2, 0x3a, 0xe7, 0xdd, 0xa7, 0x58, 0x77, 0xc3, 0xbf, 0x2d, 0xca, 0x10, 0xff,
	0xad, 0xb0, 0xd5, 0x16, 0xf5, 0xa0, 0xa5, 0x76, 0x39, 0x73, 0xed, 0x2b, 0xab, 0x73, 0xd1, 0xc5,
	0xf8, 0x20, 0x30, 0xfe, 0x97, 0x0f, 0x4e, 0x76, 0x39, 0x8b, 0x8c, 0xd6, 0xbf, 0x81, 0x96, 0x46,
	0xe8, 0x11, 0x38, 0xc9, 0x74, 0x32, 0x98, 0x7d, 0x1e, 0xc7, 0x93, 0x41, 0x7f, 0xf8, 0x61, 0x38,
	0x78, 0xef, 0xdc, 0x42, 0x17, 0x00, 0xe6, 0xf6, 0x5d, 0x3c, 0x1d, 0xf7, 0x1d, 0x0b, 0xdd, 0x87,
	0x7b, 0x06, 0x1b, 0x68, 0xa3, 0x36, 0x9c, 0x1b, 0xd8, 0xff, 0x34, 0x1a, 0x0d, 0x13, 0xe7, 0xb6,
	0xff, 0x16, 0x9e, 0x1c, 0x31, 0x95, 0xb9, 0xc8, 0x24, 0x43, 0x2e, 0xdc, 0x91, 0x1b, 0x4a, 0x99,
	0x94, 0xe6, 0x89, 0x77, 0xa3, 0x3d, 0xf4, 0x5d, 0xb8, 0xbc, 0x21, 0x52, 0xf5, 0x56, 0x82, 0x7e,
	0xfb, 0xc8, 0xf8, 0x62, 0xa9, 0xea, 0xa4, 0x7e, 0x08, 0x8f, 0x1b, 0x95, 0xba, 0xdd, 0x25, 0x9c,
	0x2d, 0xcd, 0x8d, 0xe9, 0xd6, 0x8a, 0x6a, 0xd4, 0xfd, 0x69, 0xc1, 0xc3, 0xa4, 0x9a, 0xe7, 0x44,
	0x8f, 0x33, 0x66, 0x45, 0xc9, 0x29, 0x43, 0x5f, 0xe1, 0x41, 0x23, 0x1b, 0x7a, 0x79, 0xf2, 0xa7,
	0x79, 0xaf, 0x4e, 0xa1, 0xd6, 0xd9, 0x52, 0x68, 0x1f, 0xc4, 0x46, 0x2f, 0x1a, 0xf2, 0xe3, 0x4f,
	0xf6, 0x3a, 0xff, 0x27, 0x56, 0x2e, 0x3d, 0xfc, 0xe5, 0x35, 0xd5, 0x34, 0xbd, 0x28, 0x66, 0xbb,
	0xa8, 0x58, 0x61, 0x2e, 0x02, 0x73, 0x96, 0xc1, 0xe1, 0x96, 0xcf, 0xcf, 0x4c, 0xe1, 0xcd, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x58, 0xb1, 0xb1, 0x00, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TradingProxyServiceClient is the client API for TradingProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TradingProxyServiceClient interface {
	// Submit a signed transaction
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// Get the height of the very last block processed by tendermint
	LastBlockHeight(ctx context.Context, in *LastBlockHeightRequest, opts ...grpc.CallOption) (*LastBlockHeightResponse, error)
}

type tradingProxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewTradingProxyServiceClient(cc *grpc.ClientConn) TradingProxyServiceClient {
	return &tradingProxyServiceClient{cc}
}

func (c *tradingProxyServiceClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, "/datanode.api.v1.TradingProxyService/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradingProxyServiceClient) LastBlockHeight(ctx context.Context, in *LastBlockHeightRequest, opts ...grpc.CallOption) (*LastBlockHeightResponse, error) {
	out := new(LastBlockHeightResponse)
	err := c.cc.Invoke(ctx, "/datanode.api.v1.TradingProxyService/LastBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingProxyServiceServer is the server API for TradingProxyService service.
type TradingProxyServiceServer interface {
	// Submit a signed transaction
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// Get the height of the very last block processed by tendermint
	LastBlockHeight(context.Context, *LastBlockHeightRequest) (*LastBlockHeightResponse, error)
}

func RegisterTradingProxyServiceServer(s *grpc.Server, srv TradingProxyServiceServer) {
	s.RegisterService(&_TradingProxyService_serviceDesc, srv)
}

func _TradingProxyService_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingProxyServiceServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datanode.api.v1.TradingProxyService/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingProxyServiceServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradingProxyService_LastBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingProxyServiceServer).LastBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datanode.api.v1.TradingProxyService/LastBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingProxyServiceServer).LastBlockHeight(ctx, req.(*LastBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TradingProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datanode.api.v1.TradingProxyService",
	HandlerType: (*TradingProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _TradingProxyService_SubmitTransaction_Handler,
		},
		{
			MethodName: "LastBlockHeight",
			Handler:    _TradingProxyService_LastBlockHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data-node/api/v1/trading_proxy.proto",
}
