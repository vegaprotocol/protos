// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data-node/api/v2/trading_data.proto

package v2

import (
	vega "code.vegaprotocol.io/protos/vega"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A list of the properties of an account, used for grouping
type AccountField int32

const (
	AccountField_ACCOUNT_FIELD_UNSPECIFIED AccountField = 0
	AccountField_ACCOUNT_FIELD_ID          AccountField = 1
	AccountField_ACCOUNT_FIELD_PARTY_ID    AccountField = 2
	AccountField_ACCOUNT_FIELD_ASSET_ID    AccountField = 3
	AccountField_ACCOUNT_FIELD_MARKET_ID   AccountField = 4
	AccountField_ACCOUNT_FIELD_TYPE        AccountField = 5
)

var AccountField_name = map[int32]string{
	0: "ACCOUNT_FIELD_UNSPECIFIED",
	1: "ACCOUNT_FIELD_ID",
	2: "ACCOUNT_FIELD_PARTY_ID",
	3: "ACCOUNT_FIELD_ASSET_ID",
	4: "ACCOUNT_FIELD_MARKET_ID",
	5: "ACCOUNT_FIELD_TYPE",
}

var AccountField_value = map[string]int32{
	"ACCOUNT_FIELD_UNSPECIFIED": 0,
	"ACCOUNT_FIELD_ID":          1,
	"ACCOUNT_FIELD_PARTY_ID":    2,
	"ACCOUNT_FIELD_ASSET_ID":    3,
	"ACCOUNT_FIELD_MARKET_ID":   4,
	"ACCOUNT_FIELD_TYPE":        5,
}

func (x AccountField) String() string {
	return proto.EnumName(AccountField_name, int32(x))
}

func (AccountField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28fa50d8f66c85e8, []int{0}
}

type BalanceQueryRequest struct {
	// Limit the accounts considered according to the filter supplied
	Filter *AccountFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// By default the net balances of all the accounts specified by the filter are returned.
	// If a list if fields is given in group_by, split out those balances by the supplied crietera.
	GroupBy              []AccountField `protobuf:"varint,2,rep,packed,name=group_by,json=groupBy,proto3,enum=datanode.api.v2.AccountField" json:"group_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BalanceQueryRequest) Reset()         { *m = BalanceQueryRequest{} }
func (m *BalanceQueryRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceQueryRequest) ProtoMessage()    {}
func (*BalanceQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28fa50d8f66c85e8, []int{0}
}

func (m *BalanceQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceQueryRequest.Unmarshal(m, b)
}
func (m *BalanceQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceQueryRequest.Marshal(b, m, deterministic)
}
func (m *BalanceQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceQueryRequest.Merge(m, src)
}
func (m *BalanceQueryRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceQueryRequest.Size(m)
}
func (m *BalanceQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceQueryRequest proto.InternalMessageInfo

func (m *BalanceQueryRequest) GetFilter() *AccountFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *BalanceQueryRequest) GetGroupBy() []AccountField {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

type BalanceQueryResponse struct {
	Balances             []*AggregatedBalance `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BalanceQueryResponse) Reset()         { *m = BalanceQueryResponse{} }
func (m *BalanceQueryResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceQueryResponse) ProtoMessage()    {}
func (*BalanceQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_28fa50d8f66c85e8, []int{1}
}

func (m *BalanceQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceQueryResponse.Unmarshal(m, b)
}
func (m *BalanceQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceQueryResponse.Marshal(b, m, deterministic)
}
func (m *BalanceQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceQueryResponse.Merge(m, src)
}
func (m *BalanceQueryResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceQueryResponse.Size(m)
}
func (m *BalanceQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceQueryResponse proto.InternalMessageInfo

func (m *BalanceQueryResponse) GetBalances() []*AggregatedBalance {
	if m != nil {
		return m.Balances
	}
	return nil
}

type AggregatedBalance struct {
	// Timestamp to of block the balance is referring to, in nanoseconds since the epoch
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The balance of the set of requested accounts at the time above
	Balance string `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	// If grouping by account ID, the account ID
	AccountId *wrappers.StringValue `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// If grouping by party, the party ID
	PartyId *wrappers.StringValue `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// If grouping by asset, the asset ID
	AssetId *wrappers.StringValue `protobuf:"bytes,5,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// If grouping by market, the market ID
	MarketId *wrappers.StringValue `protobuf:"bytes,6,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// If grouping by account type, the account type
	AccountType          vega.AccountType `protobuf:"varint,7,opt,name=account_type,json=accountType,proto3,enum=vega.AccountType" json:"account_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AggregatedBalance) Reset()         { *m = AggregatedBalance{} }
func (m *AggregatedBalance) String() string { return proto.CompactTextString(m) }
func (*AggregatedBalance) ProtoMessage()    {}
func (*AggregatedBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_28fa50d8f66c85e8, []int{2}
}

func (m *AggregatedBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatedBalance.Unmarshal(m, b)
}
func (m *AggregatedBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatedBalance.Marshal(b, m, deterministic)
}
func (m *AggregatedBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedBalance.Merge(m, src)
}
func (m *AggregatedBalance) XXX_Size() int {
	return xxx_messageInfo_AggregatedBalance.Size(m)
}
func (m *AggregatedBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedBalance proto.InternalMessageInfo

func (m *AggregatedBalance) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AggregatedBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *AggregatedBalance) GetAccountId() *wrappers.StringValue {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *AggregatedBalance) GetPartyId() *wrappers.StringValue {
	if m != nil {
		return m.PartyId
	}
	return nil
}

func (m *AggregatedBalance) GetAssetId() *wrappers.StringValue {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *AggregatedBalance) GetMarketId() *wrappers.StringValue {
	if m != nil {
		return m.MarketId
	}
	return nil
}

func (m *AggregatedBalance) GetAccountType() vega.AccountType {
	if m != nil {
		return m.AccountType
	}
	return vega.AccountType_ACCOUNT_TYPE_UNSPECIFIED
}

type AccountFilter struct {
	// Restrict accounts to those holding balances in this asset ID
	AssetId string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// Restrict accounts to those owned by the parties in this list (pass an empty list for no filter)
	PartyIds []string `protobuf:"bytes,2,rep,name=party_ids,json=partyIds,proto3" json:"party_ids,omitempty"`
	// Restrict accounts to those connected to the marketsin this list (pass an empty list for no filter)
	MarketIds []string `protobuf:"bytes,3,rep,name=market_ids,json=marketIds,proto3" json:"market_ids,omitempty"`
	// Restrict accounts to those connected to any of the types in this list (pass an empty list for no filter)
	AccountTypes         []vega.AccountType `protobuf:"varint,4,rep,packed,name=account_types,json=accountTypes,proto3,enum=vega.AccountType" json:"account_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AccountFilter) Reset()         { *m = AccountFilter{} }
func (m *AccountFilter) String() string { return proto.CompactTextString(m) }
func (*AccountFilter) ProtoMessage()    {}
func (*AccountFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_28fa50d8f66c85e8, []int{3}
}

func (m *AccountFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountFilter.Unmarshal(m, b)
}
func (m *AccountFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountFilter.Marshal(b, m, deterministic)
}
func (m *AccountFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountFilter.Merge(m, src)
}
func (m *AccountFilter) XXX_Size() int {
	return xxx_messageInfo_AccountFilter.Size(m)
}
func (m *AccountFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountFilter.DiscardUnknown(m)
}

var xxx_messageInfo_AccountFilter proto.InternalMessageInfo

func (m *AccountFilter) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *AccountFilter) GetPartyIds() []string {
	if m != nil {
		return m.PartyIds
	}
	return nil
}

func (m *AccountFilter) GetMarketIds() []string {
	if m != nil {
		return m.MarketIds
	}
	return nil
}

func (m *AccountFilter) GetAccountTypes() []vega.AccountType {
	if m != nil {
		return m.AccountTypes
	}
	return nil
}

func init() {
	proto.RegisterEnum("datanode.api.v2.AccountField", AccountField_name, AccountField_value)
	proto.RegisterType((*BalanceQueryRequest)(nil), "datanode.api.v2.BalanceQueryRequest")
	proto.RegisterType((*BalanceQueryResponse)(nil), "datanode.api.v2.BalanceQueryResponse")
	proto.RegisterType((*AggregatedBalance)(nil), "datanode.api.v2.AggregatedBalance")
	proto.RegisterType((*AccountFilter)(nil), "datanode.api.v2.AccountFilter")
}

func init() {
	proto.RegisterFile("data-node/api/v2/trading_data.proto", fileDescriptor_28fa50d8f66c85e8)
}

var fileDescriptor_28fa50d8f66c85e8 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x7f, 0x6b, 0xd3, 0x40,
	0x18, 0x5e, 0xda, 0x6d, 0x6d, 0xde, 0x75, 0x5b, 0x77, 0x8e, 0x99, 0x75, 0x3f, 0x28, 0x55, 0xa1,
	0x88, 0x26, 0x10, 0x65, 0x2a, 0x82, 0xd0, 0xad, 0x1d, 0x04, 0x75, 0xce, 0xb4, 0x1b, 0x4c, 0x90,
	0x72, 0x4d, 0x6e, 0x21, 0x98, 0x25, 0xf1, 0xee, 0x52, 0xc9, 0x27, 0xf0, 0x5b, 0xf8, 0x19, 0xfc,
	0x5e, 0x7e, 0x09, 0xc9, 0x5d, 0xb2, 0x35, 0xeb, 0xc4, 0xfd, 0x53, 0x7a, 0xf7, 0x3c, 0xcf, 0xdd,
	0xf3, 0xbe, 0x79, 0xde, 0x83, 0x47, 0x2e, 0xe6, 0xf8, 0x79, 0x18, 0xb9, 0xc4, 0xc0, 0xb1, 0x6f,
	0x4c, 0x4d, 0x83, 0x53, 0xec, 0xfa, 0xa1, 0x37, 0xce, 0x00, 0x3d, 0xa6, 0x11, 0x8f, 0xd0, 0x7a,
	0xf6, 0x3f, 0xe3, 0xe8, 0x38, 0xf6, 0xf5, 0xa9, 0xd9, 0x5a, 0x9f, 0x12, 0x0f, 0x1b, 0xd9, 0x8f,
	0x64, 0xb4, 0xf6, 0xbd, 0x28, 0xf2, 0x02, 0x62, 0x88, 0xd5, 0x24, 0xb9, 0x34, 0x7e, 0x50, 0x1c,
	0xc7, 0x84, 0x32, 0x89, 0x77, 0x7e, 0x2a, 0xf0, 0xe0, 0x10, 0x07, 0x38, 0x74, 0xc8, 0xe7, 0x84,
	0xd0, 0xd4, 0x26, 0xdf, 0x13, 0xc2, 0x38, 0x3a, 0x80, 0xe5, 0x4b, 0x3f, 0xe0, 0x84, 0x6a, 0x4a,
	0x5b, 0xe9, 0xae, 0x98, 0xfb, 0xfa, 0xad, 0xab, 0xf4, 0x9e, 0xe3, 0x44, 0x49, 0xc8, 0x8f, 0x05,
	0xcb, 0xce, 0xd9, 0xe8, 0x35, 0xd4, 0x3d, 0x1a, 0x25, 0xf1, 0x78, 0x92, 0x6a, 0x95, 0x76, 0xb5,
	0xbb, 0x66, 0xee, 0xfd, 0x5b, 0x49, 0x02, 0xd7, 0xae, 0x09, 0xfa, 0x61, 0xda, 0x39, 0x87, 0xcd,
	0xb2, 0x11, 0x16, 0x47, 0x21, 0x23, 0xe8, 0x1d, 0xd4, 0x27, 0x72, 0x9f, 0x69, 0x4a, 0xbb, 0xda,
	0x5d, 0x31, 0x3b, 0xf3, 0x27, 0x7a, 0x1e, 0x25, 0x1e, 0xe6, 0xc4, 0xcd, 0x8f, 0xb0, 0xaf, 0x35,
	0x9d, 0x3f, 0x15, 0xd8, 0x98, 0xc3, 0xd1, 0x2e, 0xa8, 0xdc, 0xbf, 0x22, 0x8c, 0xe3, 0xab, 0x58,
	0x94, 0x58, 0xb5, 0x6f, 0x36, 0x90, 0x06, 0xb5, 0x5c, 0xaf, 0x55, 0xda, 0x4a, 0x57, 0xb5, 0x8b,
	0x25, 0x7a, 0x0b, 0x80, 0xa5, 0xfd, 0xb1, 0xef, 0x6a, 0x55, 0xd1, 0x9b, 0x5d, 0x5d, 0x36, 0x59,
	0x2f, 0x9a, 0xac, 0x0f, 0x39, 0xf5, 0x43, 0xef, 0x1c, 0x07, 0x09, 0xb1, 0xd5, 0x9c, 0x6f, 0xb9,
	0xe8, 0x15, 0xd4, 0x63, 0x4c, 0x79, 0x9a, 0x49, 0x17, 0xef, 0x21, 0xad, 0x09, 0xb6, 0x14, 0x62,
	0xc6, 0x88, 0xb8, 0x73, 0xe9, 0x3e, 0x42, 0xc1, 0xb6, 0x5c, 0xf4, 0x06, 0xd4, 0x2b, 0x4c, 0xbf,
	0x49, 0xe5, 0xf2, 0x3d, 0x94, 0x75, 0x49, 0xb7, 0x5c, 0xf4, 0x12, 0x1a, 0x45, 0xa5, 0x3c, 0x8d,
	0x89, 0x56, 0x6b, 0x2b, 0xdd, 0x35, 0x73, 0x43, 0x17, 0xe1, 0xca, 0x3f, 0xe1, 0x28, 0x8d, 0x89,
	0xbd, 0x82, 0x6f, 0x16, 0x9d, 0x5f, 0x0a, 0xac, 0x96, 0x92, 0x81, 0xb6, 0x67, 0xbc, 0x2b, 0xb2,
	0x99, 0x85, 0xbb, 0x1d, 0x50, 0x8b, 0x7e, 0x30, 0x91, 0x16, 0xd5, 0xae, 0xe7, 0x25, 0x33, 0xb4,
	0x07, 0x70, 0x6d, 0x9d, 0x69, 0x55, 0x81, 0xaa, 0x85, 0x3b, 0x86, 0x0e, 0x60, 0x75, 0xd6, 0x1e,
	0xd3, 0x16, 0x45, 0xda, 0xee, 0xf0, 0xd7, 0x98, 0xf1, 0xc7, 0x9e, 0xfe, 0x56, 0xa0, 0x31, 0x1b,
	0x40, 0xb4, 0x07, 0xdb, 0xbd, 0xa3, 0xa3, 0x4f, 0x67, 0x27, 0xa3, 0xf1, 0xb1, 0x35, 0xf8, 0xd0,
	0x1f, 0x9f, 0x9d, 0x0c, 0x4f, 0x07, 0x47, 0xd6, 0xb1, 0x35, 0xe8, 0x37, 0x17, 0xd0, 0x26, 0x34,
	0xcb, 0xb0, 0xd5, 0x6f, 0x2a, 0xa8, 0x05, 0x5b, 0xe5, 0xdd, 0xd3, 0x9e, 0x3d, 0xba, 0xc8, 0xb0,
	0xca, 0x3c, 0xd6, 0x1b, 0x0e, 0x07, 0xa3, 0x0c, 0xab, 0xa2, 0x1d, 0x78, 0x58, 0xc6, 0x3e, 0xf6,
	0xec, 0xf7, 0x12, 0x5c, 0x44, 0x5b, 0x80, 0xca, 0xe0, 0xe8, 0xe2, 0x74, 0xd0, 0x5c, 0x32, 0x19,
	0xa0, 0x91, 0x9c, 0xfd, 0x3e, 0xe6, 0x78, 0x48, 0xe8, 0xd4, 0x77, 0x08, 0xfa, 0x0a, 0x8d, 0xd9,
	0x79, 0x41, 0x8f, 0xe7, 0xa6, 0xe2, 0x8e, 0xb9, 0x6e, 0x3d, 0xf9, 0x0f, 0x4b, 0x0e, 0x5d, 0x67,
	0xe1, 0x50, 0xff, 0xf2, 0xcc, 0xc9, 0x58, 0x59, 0x3b, 0x45, 0x52, 0x9c, 0x28, 0xd0, 0xfd, 0x48,
	0x3e, 0x24, 0xcc, 0xb8, 0xfd, 0x3a, 0x4d, 0x96, 0x05, 0xf0, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x21, 0x0d, 0x49, 0x0b, 0xb8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TradingDataServiceClient is the client API for TradingDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TradingDataServiceClient interface {
	// Get an aggregated list of the changes in balances in a set of accounts over time
	BalanceQuery(ctx context.Context, in *BalanceQueryRequest, opts ...grpc.CallOption) (*BalanceQueryResponse, error)
}

type tradingDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewTradingDataServiceClient(cc *grpc.ClientConn) TradingDataServiceClient {
	return &tradingDataServiceClient{cc}
}

func (c *tradingDataServiceClient) BalanceQuery(ctx context.Context, in *BalanceQueryRequest, opts ...grpc.CallOption) (*BalanceQueryResponse, error) {
	out := new(BalanceQueryResponse)
	err := c.cc.Invoke(ctx, "/datanode.api.v2.TradingDataService/BalanceQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradingDataServiceServer is the server API for TradingDataService service.
type TradingDataServiceServer interface {
	// Get an aggregated list of the changes in balances in a set of accounts over time
	BalanceQuery(context.Context, *BalanceQueryRequest) (*BalanceQueryResponse, error)
}

func RegisterTradingDataServiceServer(s *grpc.Server, srv TradingDataServiceServer) {
	s.RegisterService(&_TradingDataService_serviceDesc, srv)
}

func _TradingDataService_BalanceQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradingDataServiceServer).BalanceQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datanode.api.v2.TradingDataService/BalanceQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradingDataServiceServer).BalanceQuery(ctx, req.(*BalanceQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TradingDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datanode.api.v2.TradingDataService",
	HandlerType: (*TradingDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BalanceQuery",
			Handler:    _TradingDataService_BalanceQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data-node/api/v2/trading_data.proto",
}
