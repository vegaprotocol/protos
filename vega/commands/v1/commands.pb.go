// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/commands/v1/commands.proto

package v1

import (
	vega "code.vegaprotocol.io/protos/vega"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
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

type UndelegateSubmission_Method int32

const (
	UndelegateSubmission_METHOD_UNSPECIFIED     UndelegateSubmission_Method = 0
	UndelegateSubmission_METHOD_NOW             UndelegateSubmission_Method = 1
	UndelegateSubmission_METHOD_AT_END_OF_EPOCH UndelegateSubmission_Method = 2
	UndelegateSubmission_METHOD_IN_ANGER        UndelegateSubmission_Method = 3
)

var UndelegateSubmission_Method_name = map[int32]string{
	0: "METHOD_UNSPECIFIED",
	1: "METHOD_NOW",
	2: "METHOD_AT_END_OF_EPOCH",
	3: "METHOD_IN_ANGER",
}

var UndelegateSubmission_Method_value = map[string]int32{
	"METHOD_UNSPECIFIED":     0,
	"METHOD_NOW":             1,
	"METHOD_AT_END_OF_EPOCH": 2,
	"METHOD_IN_ANGER":        3,
}

func (x UndelegateSubmission_Method) String() string {
	return proto.EnumName(UndelegateSubmission_Method_name, int32(x))
}

func (UndelegateSubmission_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{10, 0}
}

// An order submission is a request to submit or create a new order on Vega
type OrderSubmission struct {
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Price for the order, the price is an integer, for example `123456` is a correctly
	// formatted price of `1.23456` assuming market configured to 5 decimal places,
	// , required field for limit orders, however it is not required for market orders
	Price string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	// Size for the order, for example, in a futures market the size equals the number of contracts, cannot be negative
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// Side for the order, e.g. SIDE_BUY or SIDE_SELL, required field - See [`Side`](#vega.Side)
	Side vega.Side `protobuf:"varint,4,opt,name=side,proto3,enum=vega.Side" json:"side,omitempty"`
	// Time in force indicates how long an order will remain active before it is executed or expires, required field
	// - See [`Order.TimeInForce`](#vega.Order.TimeInForce)
	TimeInForce vega.Order_TimeInForce `protobuf:"varint,5,opt,name=time_in_force,json=timeInForce,proto3,enum=vega.Order_TimeInForce" json:"time_in_force,omitempty"`
	// Timestamp for when the order will expire, in nanoseconds since the epoch,
	// required field only for [`Order.TimeInForce`](#vega.Order.TimeInForce)`.TIME_IN_FORCE_GTT`
	// - See [`VegaTimeResponse`](#api.VegaTimeResponse).`timestamp`
	ExpiresAt int64 `protobuf:"varint,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Type for the order, required field - See [`Order.Type`](#vega.Order.Type)
	Type vega.Order_Type `protobuf:"varint,7,opt,name=type,proto3,enum=vega.Order_Type" json:"type,omitempty"`
	// Reference given for the order, this is typically used to retrieve an order submitted through consensus, currently
	// set internally by the node to return a unique reference identifier for the order submission
	Reference string `protobuf:"bytes,8,opt,name=reference,proto3" json:"reference,omitempty"`
	// Used to specify the details for a pegged order
	// - See [`PeggedOrder`](#vega.PeggedOrder)
	PeggedOrder          *vega.PeggedOrder `protobuf:"bytes,9,opt,name=pegged_order,json=peggedOrder,proto3" json:"pegged_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderSubmission) Reset()         { *m = OrderSubmission{} }
func (m *OrderSubmission) String() string { return proto.CompactTextString(m) }
func (*OrderSubmission) ProtoMessage()    {}
func (*OrderSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{0}
}

func (m *OrderSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderSubmission.Unmarshal(m, b)
}
func (m *OrderSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderSubmission.Marshal(b, m, deterministic)
}
func (m *OrderSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderSubmission.Merge(m, src)
}
func (m *OrderSubmission) XXX_Size() int {
	return xxx_messageInfo_OrderSubmission.Size(m)
}
func (m *OrderSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_OrderSubmission proto.InternalMessageInfo

func (m *OrderSubmission) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderSubmission) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *OrderSubmission) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *OrderSubmission) GetSide() vega.Side {
	if m != nil {
		return m.Side
	}
	return vega.Side_SIDE_UNSPECIFIED
}

func (m *OrderSubmission) GetTimeInForce() vega.Order_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return vega.Order_TIME_IN_FORCE_UNSPECIFIED
}

func (m *OrderSubmission) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *OrderSubmission) GetType() vega.Order_Type {
	if m != nil {
		return m.Type
	}
	return vega.Order_TYPE_UNSPECIFIED
}

func (m *OrderSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *OrderSubmission) GetPeggedOrder() *vega.PeggedOrder {
	if m != nil {
		return m.PeggedOrder
	}
	return nil
}

// An order cancellation is a request to cancel an existing order on Vega
type OrderCancellation struct {
	// Unique identifier for the order (set by the system after consensus), required field
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Market identifier for the order, required field
	MarketId             string   `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderCancellation) Reset()         { *m = OrderCancellation{} }
func (m *OrderCancellation) String() string { return proto.CompactTextString(m) }
func (*OrderCancellation) ProtoMessage()    {}
func (*OrderCancellation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{1}
}

func (m *OrderCancellation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderCancellation.Unmarshal(m, b)
}
func (m *OrderCancellation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderCancellation.Marshal(b, m, deterministic)
}
func (m *OrderCancellation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderCancellation.Merge(m, src)
}
func (m *OrderCancellation) XXX_Size() int {
	return xxx_messageInfo_OrderCancellation.Size(m)
}
func (m *OrderCancellation) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderCancellation.DiscardUnknown(m)
}

var xxx_messageInfo_OrderCancellation proto.InternalMessageInfo

func (m *OrderCancellation) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderCancellation) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

// An order amendment is a request to amend or update an existing order on Vega
type OrderAmendment struct {
	// Order identifier, this is required to find the order and will not be updated, required field
	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// Market identifier, this is required to find the order and will not be updated
	MarketId string `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Amend the price for the order, if the Price value is set, otherwise price will remain unchanged - See [`Price`](#vega.Price)
	Price *vega.Price `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	// Amend the size for the order by the delta specified:
	// - To reduce the size from the current value set a negative integer value
	// - To increase the size from the current value, set a positive integer value
	// - To leave the size unchanged set a value of zero
	SizeDelta int64 `protobuf:"varint,4,opt,name=size_delta,json=sizeDelta,proto3" json:"size_delta,omitempty"`
	// Amend the expiry time for the order, if the Timestamp value is set, otherwise expiry time will remain unchanged
	// - See [`VegaTimeResponse`](#api.VegaTimeResponse).`timestamp`
	ExpiresAt *vega.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Amend the time in force for the order, set to TIME_IN_FORCE_UNSPECIFIED to remain unchanged
	// - See [`TimeInForce`](#api.VegaTimeResponse).`timestamp`
	TimeInForce vega.Order_TimeInForce `protobuf:"varint,6,opt,name=time_in_force,json=timeInForce,proto3,enum=vega.Order_TimeInForce" json:"time_in_force,omitempty"`
	// Amend the pegged order offset for the order
	PeggedOffset string `protobuf:"bytes,7,opt,name=pegged_offset,json=peggedOffset,proto3" json:"pegged_offset,omitempty"`
	// Amend the pegged order reference for the order
	// - See [`PeggedReference`](#vega.PeggedReference)
	PeggedReference      vega.PeggedReference `protobuf:"varint,8,opt,name=pegged_reference,json=peggedReference,proto3,enum=vega.PeggedReference" json:"pegged_reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderAmendment) Reset()         { *m = OrderAmendment{} }
func (m *OrderAmendment) String() string { return proto.CompactTextString(m) }
func (*OrderAmendment) ProtoMessage()    {}
func (*OrderAmendment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{2}
}

func (m *OrderAmendment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderAmendment.Unmarshal(m, b)
}
func (m *OrderAmendment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderAmendment.Marshal(b, m, deterministic)
}
func (m *OrderAmendment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderAmendment.Merge(m, src)
}
func (m *OrderAmendment) XXX_Size() int {
	return xxx_messageInfo_OrderAmendment.Size(m)
}
func (m *OrderAmendment) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderAmendment.DiscardUnknown(m)
}

var xxx_messageInfo_OrderAmendment proto.InternalMessageInfo

func (m *OrderAmendment) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *OrderAmendment) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *OrderAmendment) GetPrice() *vega.Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *OrderAmendment) GetSizeDelta() int64 {
	if m != nil {
		return m.SizeDelta
	}
	return 0
}

func (m *OrderAmendment) GetExpiresAt() *vega.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (m *OrderAmendment) GetTimeInForce() vega.Order_TimeInForce {
	if m != nil {
		return m.TimeInForce
	}
	return vega.Order_TIME_IN_FORCE_UNSPECIFIED
}

func (m *OrderAmendment) GetPeggedOffset() string {
	if m != nil {
		return m.PeggedOffset
	}
	return ""
}

func (m *OrderAmendment) GetPeggedReference() vega.PeggedReference {
	if m != nil {
		return m.PeggedReference
	}
	return vega.PeggedReference_PEGGED_REFERENCE_UNSPECIFIED
}

// A liquidity provision submitted for a given market
type LiquidityProvisionSubmission struct {
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Specified as a unitless number that represents the amount of settlement asset of the market
	CommitmentAmount string `protobuf:"bytes,2,opt,name=commitment_amount,json=commitmentAmount,proto3" json:"commitment_amount,omitempty"`
	// Nominated liquidity fee factor, which is an input to the calculation of taker fees on the market, as per seeting fees and rewarding liquidity providers
	Fee string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	// A set of liquidity sell orders to meet the liquidity provision obligation
	Sells []*vega.LiquidityOrder `protobuf:"bytes,4,rep,name=sells,proto3" json:"sells,omitempty"`
	// A set of liquidity buy orders to meet the liquidity provision obligation
	Buys []*vega.LiquidityOrder `protobuf:"bytes,5,rep,name=buys,proto3" json:"buys,omitempty"`
	// A reference to be added to every order created out of this liquidityProvisionSubmission
	Reference            string   `protobuf:"bytes,6,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiquidityProvisionSubmission) Reset()         { *m = LiquidityProvisionSubmission{} }
func (m *LiquidityProvisionSubmission) String() string { return proto.CompactTextString(m) }
func (*LiquidityProvisionSubmission) ProtoMessage()    {}
func (*LiquidityProvisionSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{3}
}

func (m *LiquidityProvisionSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquidityProvisionSubmission.Unmarshal(m, b)
}
func (m *LiquidityProvisionSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquidityProvisionSubmission.Marshal(b, m, deterministic)
}
func (m *LiquidityProvisionSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityProvisionSubmission.Merge(m, src)
}
func (m *LiquidityProvisionSubmission) XXX_Size() int {
	return xxx_messageInfo_LiquidityProvisionSubmission.Size(m)
}
func (m *LiquidityProvisionSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityProvisionSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityProvisionSubmission proto.InternalMessageInfo

func (m *LiquidityProvisionSubmission) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *LiquidityProvisionSubmission) GetCommitmentAmount() string {
	if m != nil {
		return m.CommitmentAmount
	}
	return ""
}

func (m *LiquidityProvisionSubmission) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *LiquidityProvisionSubmission) GetSells() []*vega.LiquidityOrder {
	if m != nil {
		return m.Sells
	}
	return nil
}

func (m *LiquidityProvisionSubmission) GetBuys() []*vega.LiquidityOrder {
	if m != nil {
		return m.Buys
	}
	return nil
}

func (m *LiquidityProvisionSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

// Cancel a liquidity provision request
type LiquidityProvisionCancellation struct {
	MarketId             string   `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiquidityProvisionCancellation) Reset()         { *m = LiquidityProvisionCancellation{} }
func (m *LiquidityProvisionCancellation) String() string { return proto.CompactTextString(m) }
func (*LiquidityProvisionCancellation) ProtoMessage()    {}
func (*LiquidityProvisionCancellation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{4}
}

func (m *LiquidityProvisionCancellation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquidityProvisionCancellation.Unmarshal(m, b)
}
func (m *LiquidityProvisionCancellation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquidityProvisionCancellation.Marshal(b, m, deterministic)
}
func (m *LiquidityProvisionCancellation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityProvisionCancellation.Merge(m, src)
}
func (m *LiquidityProvisionCancellation) XXX_Size() int {
	return xxx_messageInfo_LiquidityProvisionCancellation.Size(m)
}
func (m *LiquidityProvisionCancellation) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityProvisionCancellation.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityProvisionCancellation proto.InternalMessageInfo

func (m *LiquidityProvisionCancellation) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

// Amend a liquidity provision request
type LiquidityProvisionAmendment struct {
	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// From here at least one of the following is required to consider the command valid
	CommitmentAmount     string                 `protobuf:"bytes,2,opt,name=commitment_amount,json=commitmentAmount,proto3" json:"commitment_amount,omitempty"`
	Fee                  string                 `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Sells                []*vega.LiquidityOrder `protobuf:"bytes,4,rep,name=sells,proto3" json:"sells,omitempty"`
	Buys                 []*vega.LiquidityOrder `protobuf:"bytes,5,rep,name=buys,proto3" json:"buys,omitempty"`
	Reference            string                 `protobuf:"bytes,6,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LiquidityProvisionAmendment) Reset()         { *m = LiquidityProvisionAmendment{} }
func (m *LiquidityProvisionAmendment) String() string { return proto.CompactTextString(m) }
func (*LiquidityProvisionAmendment) ProtoMessage()    {}
func (*LiquidityProvisionAmendment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{5}
}

func (m *LiquidityProvisionAmendment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiquidityProvisionAmendment.Unmarshal(m, b)
}
func (m *LiquidityProvisionAmendment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiquidityProvisionAmendment.Marshal(b, m, deterministic)
}
func (m *LiquidityProvisionAmendment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityProvisionAmendment.Merge(m, src)
}
func (m *LiquidityProvisionAmendment) XXX_Size() int {
	return xxx_messageInfo_LiquidityProvisionAmendment.Size(m)
}
func (m *LiquidityProvisionAmendment) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityProvisionAmendment.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityProvisionAmendment proto.InternalMessageInfo

func (m *LiquidityProvisionAmendment) GetMarketId() string {
	if m != nil {
		return m.MarketId
	}
	return ""
}

func (m *LiquidityProvisionAmendment) GetCommitmentAmount() string {
	if m != nil {
		return m.CommitmentAmount
	}
	return ""
}

func (m *LiquidityProvisionAmendment) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *LiquidityProvisionAmendment) GetSells() []*vega.LiquidityOrder {
	if m != nil {
		return m.Sells
	}
	return nil
}

func (m *LiquidityProvisionAmendment) GetBuys() []*vega.LiquidityOrder {
	if m != nil {
		return m.Buys
	}
	return nil
}

func (m *LiquidityProvisionAmendment) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

// Represents the submission request to withdraw funds for a party on Vega
type WithdrawSubmission struct {
	// The amount to be withdrawn
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The asset we want to withdraw
	Asset string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	// Foreign chain specifics
	Ext                  *vega.WithdrawExt `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WithdrawSubmission) Reset()         { *m = WithdrawSubmission{} }
func (m *WithdrawSubmission) String() string { return proto.CompactTextString(m) }
func (*WithdrawSubmission) ProtoMessage()    {}
func (*WithdrawSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{6}
}

func (m *WithdrawSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawSubmission.Unmarshal(m, b)
}
func (m *WithdrawSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawSubmission.Marshal(b, m, deterministic)
}
func (m *WithdrawSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawSubmission.Merge(m, src)
}
func (m *WithdrawSubmission) XXX_Size() int {
	return xxx_messageInfo_WithdrawSubmission.Size(m)
}
func (m *WithdrawSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawSubmission proto.InternalMessageInfo

func (m *WithdrawSubmission) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *WithdrawSubmission) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *WithdrawSubmission) GetExt() *vega.WithdrawExt {
	if m != nil {
		return m.Ext
	}
	return nil
}

// A command to submit a new proposal for the
// vega network governance
type ProposalSubmission struct {
	// Proposal reference
	Reference string `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	// Proposal configuration and the actual change that is meant to be executed when proposal is enacted
	Terms                *vega.ProposalTerms `protobuf:"bytes,2,opt,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProposalSubmission) Reset()         { *m = ProposalSubmission{} }
func (m *ProposalSubmission) String() string { return proto.CompactTextString(m) }
func (*ProposalSubmission) ProtoMessage()    {}
func (*ProposalSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{7}
}

func (m *ProposalSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalSubmission.Unmarshal(m, b)
}
func (m *ProposalSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalSubmission.Marshal(b, m, deterministic)
}
func (m *ProposalSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalSubmission.Merge(m, src)
}
func (m *ProposalSubmission) XXX_Size() int {
	return xxx_messageInfo_ProposalSubmission.Size(m)
}
func (m *ProposalSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalSubmission proto.InternalMessageInfo

func (m *ProposalSubmission) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *ProposalSubmission) GetTerms() *vega.ProposalTerms {
	if m != nil {
		return m.Terms
	}
	return nil
}

// A command to submit a new vote for a governance
// proposal.
type VoteSubmission struct {
	// The ID of the proposal to vote for.
	ProposalId string `protobuf:"bytes,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	// The actual value of the vote
	Value                vega.Vote_Value `protobuf:"varint,2,opt,name=value,proto3,enum=vega.Vote_Value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VoteSubmission) Reset()         { *m = VoteSubmission{} }
func (m *VoteSubmission) String() string { return proto.CompactTextString(m) }
func (*VoteSubmission) ProtoMessage()    {}
func (*VoteSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{8}
}

func (m *VoteSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoteSubmission.Unmarshal(m, b)
}
func (m *VoteSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoteSubmission.Marshal(b, m, deterministic)
}
func (m *VoteSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteSubmission.Merge(m, src)
}
func (m *VoteSubmission) XXX_Size() int {
	return xxx_messageInfo_VoteSubmission.Size(m)
}
func (m *VoteSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_VoteSubmission proto.InternalMessageInfo

func (m *VoteSubmission) GetProposalId() string {
	if m != nil {
		return m.ProposalId
	}
	return ""
}

func (m *VoteSubmission) GetValue() vega.Vote_Value {
	if m != nil {
		return m.Value
	}
	return vega.Vote_VALUE_UNSPECIFIED
}

// A command to submit an instruction to delegate some stake to a node
type DelegateSubmission struct {
	// The ID for the node to delegate to
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// The amount of stake to delegate
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegateSubmission) Reset()         { *m = DelegateSubmission{} }
func (m *DelegateSubmission) String() string { return proto.CompactTextString(m) }
func (*DelegateSubmission) ProtoMessage()    {}
func (*DelegateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{9}
}

func (m *DelegateSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegateSubmission.Unmarshal(m, b)
}
func (m *DelegateSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegateSubmission.Marshal(b, m, deterministic)
}
func (m *DelegateSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegateSubmission.Merge(m, src)
}
func (m *DelegateSubmission) XXX_Size() int {
	return xxx_messageInfo_DelegateSubmission.Size(m)
}
func (m *DelegateSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegateSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_DelegateSubmission proto.InternalMessageInfo

func (m *DelegateSubmission) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *DelegateSubmission) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type UndelegateSubmission struct {
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// optional, if not specified = ALL
	Amount               string                      `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Method               UndelegateSubmission_Method `protobuf:"varint,3,opt,name=method,proto3,enum=vega.commands.v1.UndelegateSubmission_Method" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UndelegateSubmission) Reset()         { *m = UndelegateSubmission{} }
func (m *UndelegateSubmission) String() string { return proto.CompactTextString(m) }
func (*UndelegateSubmission) ProtoMessage()    {}
func (*UndelegateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{10}
}

func (m *UndelegateSubmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndelegateSubmission.Unmarshal(m, b)
}
func (m *UndelegateSubmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndelegateSubmission.Marshal(b, m, deterministic)
}
func (m *UndelegateSubmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndelegateSubmission.Merge(m, src)
}
func (m *UndelegateSubmission) XXX_Size() int {
	return xxx_messageInfo_UndelegateSubmission.Size(m)
}
func (m *UndelegateSubmission) XXX_DiscardUnknown() {
	xxx_messageInfo_UndelegateSubmission.DiscardUnknown(m)
}

var xxx_messageInfo_UndelegateSubmission proto.InternalMessageInfo

func (m *UndelegateSubmission) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *UndelegateSubmission) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *UndelegateSubmission) GetMethod() UndelegateSubmission_Method {
	if m != nil {
		return m.Method
	}
	return UndelegateSubmission_METHOD_UNSPECIFIED
}

// A command that loads the state from a given checkpoint
type RestoreSnapshot struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreSnapshot) Reset()         { *m = RestoreSnapshot{} }
func (m *RestoreSnapshot) String() string { return proto.CompactTextString(m) }
func (*RestoreSnapshot) ProtoMessage()    {}
func (*RestoreSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{11}
}

func (m *RestoreSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreSnapshot.Unmarshal(m, b)
}
func (m *RestoreSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreSnapshot.Marshal(b, m, deterministic)
}
func (m *RestoreSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreSnapshot.Merge(m, src)
}
func (m *RestoreSnapshot) XXX_Size() int {
	return xxx_messageInfo_RestoreSnapshot.Size(m)
}
func (m *RestoreSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreSnapshot proto.InternalMessageInfo

func (m *RestoreSnapshot) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("vega.commands.v1.UndelegateSubmission_Method", UndelegateSubmission_Method_name, UndelegateSubmission_Method_value)
	proto.RegisterType((*OrderSubmission)(nil), "vega.commands.v1.OrderSubmission")
	proto.RegisterType((*OrderCancellation)(nil), "vega.commands.v1.OrderCancellation")
	proto.RegisterType((*OrderAmendment)(nil), "vega.commands.v1.OrderAmendment")
	proto.RegisterType((*LiquidityProvisionSubmission)(nil), "vega.commands.v1.LiquidityProvisionSubmission")
	proto.RegisterType((*LiquidityProvisionCancellation)(nil), "vega.commands.v1.LiquidityProvisionCancellation")
	proto.RegisterType((*LiquidityProvisionAmendment)(nil), "vega.commands.v1.LiquidityProvisionAmendment")
	proto.RegisterType((*WithdrawSubmission)(nil), "vega.commands.v1.WithdrawSubmission")
	proto.RegisterType((*ProposalSubmission)(nil), "vega.commands.v1.ProposalSubmission")
	proto.RegisterType((*VoteSubmission)(nil), "vega.commands.v1.VoteSubmission")
	proto.RegisterType((*DelegateSubmission)(nil), "vega.commands.v1.DelegateSubmission")
	proto.RegisterType((*UndelegateSubmission)(nil), "vega.commands.v1.UndelegateSubmission")
	proto.RegisterType((*RestoreSnapshot)(nil), "vega.commands.v1.RestoreSnapshot")
}

func init() { proto.RegisterFile("vega/commands/v1/commands.proto", fileDescriptor_9c3a53c66a940c51) }

var fileDescriptor_9c3a53c66a940c51 = []byte{
	// 962 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdf, 0x6f, 0x1b, 0x45,
	0x10, 0xee, 0xf9, 0x57, 0xe3, 0x71, 0x6b, 0x3b, 0xdb, 0x34, 0x1c, 0x49, 0x9b, 0xb8, 0x17, 0x7e,
	0x58, 0x40, 0xce, 0x8a, 0x41, 0xbc, 0x20, 0x24, 0xd2, 0xc4, 0xa1, 0x16, 0x24, 0x8e, 0x36, 0x69,
	0x8b, 0x78, 0x39, 0x6d, 0x7c, 0xe3, 0xcb, 0xaa, 0x77, 0xb7, 0xc7, 0xdd, 0xda, 0x49, 0xf8, 0x63,
	0x41, 0xe2, 0x81, 0x37, 0x5e, 0x91, 0x78, 0x43, 0xbb, 0x7b, 0x76, 0xec, 0x18, 0x55, 0x54, 0xe2,
	0x85, 0xb7, 0x9d, 0x6f, 0x66, 0xbe, 0xdd, 0x99, 0x6f, 0x76, 0xef, 0x60, 0x7b, 0x82, 0x01, 0xeb,
	0x0c, 0x45, 0x14, 0xb1, 0xd8, 0xcf, 0x3a, 0x93, 0xbd, 0xd9, 0xda, 0x4d, 0x52, 0x21, 0x05, 0x69,
	0xaa, 0x00, 0x77, 0x06, 0x4e, 0xf6, 0x36, 0x1e, 0xeb, 0x94, 0x40, 0x4c, 0x30, 0x8d, 0x59, 0x3c,
	0x44, 0x13, 0xb8, 0xd1, 0xd0, 0xb0, 0x8e, 0x36, 0xc0, 0x56, 0x20, 0x44, 0x10, 0x62, 0x47, 0x5b,
	0x17, 0xe3, 0x51, 0xe7, 0x2a, 0x65, 0x49, 0x82, 0x69, 0xce, 0xbc, 0xf1, 0x65, 0xc0, 0xe5, 0xe5,
	0xf8, 0x42, 0x71, 0x77, 0xa2, 0x2b, 0x2e, 0xdf, 0x88, 0xab, 0x4e, 0x20, 0x76, 0xb5, 0x73, 0x77,
	0xc2, 0x42, 0xee, 0x33, 0x29, 0xd2, 0xac, 0x33, 0x5b, 0x9a, 0x3c, 0xe7, 0x97, 0x02, 0x34, 0x06,
	0xa9, 0x8f, 0xe9, 0xd9, 0xf8, 0x22, 0xe2, 0x59, 0xc6, 0x45, 0x4c, 0x76, 0xa0, 0x1a, 0xb1, 0xf4,
	0x0d, 0x4a, 0x8f, 0xfb, 0xb6, 0xd5, 0xb2, 0xda, 0xd5, 0xe7, 0x95, 0xdf, 0x7e, 0xdd, 0x2e, 0xfc,
	0x60, 0xd1, 0x15, 0xe3, 0xe8, 0xfb, 0x64, 0x0d, 0xca, 0x49, 0xca, 0x87, 0x68, 0x17, 0x54, 0x00,
	0x35, 0x06, 0xd9, 0x80, 0x52, 0xc6, 0x7f, 0x46, 0xbb, 0xd8, 0xb2, 0xda, 0x25, 0x93, 0xd5, 0xbc,
	0x47, 0x35, 0x46, 0xb6, 0x94, 0xcf, 0x47, 0xbb, 0xd4, 0xb2, 0xda, 0xf5, 0x2e, 0xb8, 0xba, 0xba,
	0x33, 0xee, 0x23, 0xd5, 0x38, 0xf9, 0x0a, 0x1e, 0x4a, 0x1e, 0xa1, 0xc7, 0x63, 0x6f, 0x24, 0xd2,
	0x21, 0xda, 0x65, 0x1d, 0xf8, 0x9e, 0x09, 0xd4, 0x87, 0x74, 0xcf, 0x79, 0x84, 0xfd, 0xf8, 0x48,
	0xb9, 0x69, 0x4d, 0xde, 0x1a, 0xe4, 0x29, 0x00, 0x5e, 0x27, 0x3c, 0xc5, 0xcc, 0x63, 0xd2, 0xae,
	0xb4, 0xac, 0x76, 0x91, 0x56, 0x73, 0x64, 0x5f, 0x92, 0x0f, 0xa0, 0x24, 0x6f, 0x12, 0xb4, 0xef,
	0x6b, 0xca, 0xe6, 0x02, 0xe5, 0x4d, 0x82, 0x54, 0x7b, 0xc9, 0x13, 0xa8, 0xa6, 0x38, 0xc2, 0x14,
	0xe3, 0x21, 0xda, 0x2b, 0xba, 0xae, 0x5b, 0x80, 0x7c, 0x01, 0x0f, 0x12, 0x0c, 0x02, 0xf4, 0x3d,
	0xa1, 0x12, 0xed, 0x6a, 0xcb, 0x6a, 0xd7, 0xba, 0xab, 0x86, 0xeb, 0x54, 0x7b, 0x34, 0x23, 0xad,
	0x25, 0xb7, 0x86, 0xf3, 0x1d, 0xac, 0xea, 0xc5, 0x81, 0x52, 0x37, 0x0c, 0x99, 0x54, 0x1d, 0x7e,
	0x1f, 0x56, 0x34, 0xc7, 0xac, 0xc1, 0xf4, 0xbe, 0xb6, 0xfb, 0x3e, 0xd9, 0x9c, 0x6f, 0xbe, 0xe9,
	0xed, 0xac, 0xe9, 0xce, 0xef, 0x05, 0xa8, 0x6b, 0xb6, 0xfd, 0x08, 0x63, 0x3f, 0xc2, 0x58, 0x92,
	0x67, 0x77, 0xa9, 0x66, 0x5a, 0xfd, 0x2b, 0x4a, 0xf2, 0x6c, 0xaa, 0x63, 0x51, 0x97, 0x53, 0xcb,
	0xcb, 0x51, 0xd0, 0x54, 0xd4, 0xa7, 0x00, 0x4a, 0x40, 0xcf, 0xc7, 0x50, 0x32, 0x2d, 0x5f, 0x91,
	0x56, 0x15, 0x72, 0xa8, 0x00, 0xe2, 0x2e, 0xb4, 0xbe, 0xac, 0x69, 0x1a, 0x86, 0x46, 0xc9, 0x95,
	0x49, 0x16, 0x25, 0xf3, 0x5a, 0x2c, 0xe9, 0x5c, 0x79, 0x07, 0x9d, 0x77, 0xe0, 0xe1, 0x54, 0x84,
	0xd1, 0x28, 0x43, 0xa9, 0x15, 0xad, 0xd2, 0x5c, 0x99, 0x81, 0xc6, 0xc8, 0x37, 0xd0, 0xcc, 0x83,
	0x16, 0xe5, 0xac, 0x77, 0x1f, 0xcf, 0xab, 0x45, 0xa7, 0x4e, 0xda, 0x48, 0x16, 0x01, 0xe7, 0x4f,
	0x0b, 0x9e, 0x7c, 0xcf, 0x7f, 0x1a, 0x73, 0x9f, 0xcb, 0x9b, 0xd3, 0x54, 0x4c, 0xb8, 0xba, 0x19,
	0xef, 0x7a, 0x47, 0x3e, 0x85, 0x55, 0x75, 0xd7, 0xb9, 0x54, 0x4a, 0x79, 0x2c, 0x12, 0xe3, 0x58,
	0xe6, 0x02, 0x34, 0x6f, 0x1d, 0xfb, 0x1a, 0x27, 0x4d, 0x28, 0x8e, 0xd0, 0xc8, 0x50, 0xa5, 0x6a,
	0x49, 0x3e, 0x81, 0x72, 0x86, 0x61, 0x98, 0xd9, 0xa5, 0x56, 0xb1, 0x5d, 0xeb, 0xae, 0x99, 0xb3,
	0xcf, 0x8e, 0x65, 0x86, 0xcd, 0x84, 0x90, 0x36, 0x94, 0x2e, 0xc6, 0x37, 0x99, 0x5d, 0x7e, 0x4b,
	0xa8, 0x8e, 0x58, 0x1c, 0xf2, 0xca, 0x9d, 0x21, 0x77, 0xbe, 0x86, 0xad, 0xe5, 0xba, 0x17, 0x66,
	0x77, 0x73, 0xa9, 0xf2, 0xb9, 0x01, 0xfd, 0xc3, 0x82, 0xcd, 0xe5, 0xfc, 0xdb, 0x69, 0x7d, 0x5b,
	0xf2, 0xff, 0xb1, 0x5d, 0x01, 0x90, 0xd7, 0x5c, 0x5e, 0xfa, 0x29, 0xbb, 0x9a, 0x1b, 0x8e, 0x75,
	0xa8, 0xe4, 0xa7, 0x37, 0x25, 0xe6, 0x96, 0x7a, 0x33, 0x59, 0xa6, 0x86, 0x36, 0x7f, 0x33, 0xb5,
	0x41, 0x76, 0xa0, 0x88, 0xd7, 0x32, 0xbf, 0x7f, 0xf9, 0x73, 0x32, 0x25, 0xed, 0x5d, 0x4b, 0xaa,
	0xbc, 0x0e, 0x02, 0x39, 0x4d, 0x45, 0x22, 0x32, 0x16, 0xce, 0x6d, 0xb4, 0x70, 0x38, 0xeb, 0xee,
	0x83, 0xb5, 0x07, 0x65, 0x89, 0x69, 0x94, 0xe9, 0xed, 0x6a, 0xdd, 0x47, 0xd3, 0xab, 0x6d, 0x68,
	0xce, 0x95, 0xcb, 0x0c, 0x6d, 0xcb, 0xa2, 0x26, 0xd2, 0x61, 0x50, 0x7f, 0x25, 0x24, 0xce, 0x6d,
	0xf1, 0x31, 0xd4, 0x92, 0x3c, 0x63, 0x79, 0xd4, 0x61, 0xea, 0xea, 0xfb, 0xe4, 0x23, 0x28, 0x4f,
	0x58, 0x38, 0x36, 0x1f, 0x84, 0xd9, 0x1b, 0xab, 0xd8, 0xdc, 0x57, 0x0a, 0xa7, 0xc6, 0xed, 0x1c,
	0x03, 0x39, 0xc4, 0x10, 0x03, 0xb6, 0xb0, 0xcd, 0x36, 0xdc, 0x8f, 0x85, 0x8f, 0xcb, 0x5b, 0x54,
	0x14, 0xdc, 0xf7, 0xe7, 0x7a, 0x5a, 0x98, 0xef, 0xa9, 0xf3, 0x97, 0x05, 0x6b, 0x2f, 0x63, 0xff,
	0xbf, 0x63, 0x24, 0x3d, 0xa8, 0x44, 0x28, 0x2f, 0x85, 0xaf, 0x25, 0xa9, 0x77, 0x77, 0xdd, 0xbb,
	0x5f, 0x6d, 0xf7, 0x9f, 0x36, 0x74, 0x8f, 0x75, 0x12, 0xcd, 0x93, 0x1d, 0x06, 0x15, 0x83, 0x90,
	0x75, 0x20, 0xc7, 0xbd, 0xf3, 0x17, 0x83, 0x43, 0xef, 0xe5, 0xc9, 0xd9, 0x69, 0xef, 0xa0, 0x7f,
	0xd4, 0xef, 0x1d, 0x36, 0xef, 0x91, 0x3a, 0x40, 0x8e, 0x9f, 0x0c, 0x5e, 0x37, 0x2d, 0xb2, 0x01,
	0xeb, 0xb9, 0xbd, 0x7f, 0xee, 0xf5, 0x4e, 0x0e, 0xbd, 0xc1, 0x91, 0xd7, 0x3b, 0x1d, 0x1c, 0xbc,
	0x68, 0x16, 0xc8, 0x23, 0x68, 0xe4, 0xbe, 0xfe, 0x89, 0xb7, 0x7f, 0xf2, 0x6d, 0x8f, 0x36, 0x8b,
	0xce, 0x87, 0xd0, 0xa0, 0x98, 0x49, 0x91, 0xe2, 0x59, 0xcc, 0x92, 0xec, 0x52, 0x48, 0x42, 0xa0,
	0xe4, 0x33, 0xc9, 0x74, 0xc9, 0x0f, 0xa8, 0x5e, 0x3f, 0x77, 0x7f, 0xfc, 0x6c, 0x28, 0x7c, 0xd4,
	0x65, 0xe8, 0xaf, 0xfe, 0x50, 0x84, 0x2e, 0x17, 0xe6, 0x5f, 0x22, 0xeb, 0xdc, 0xfd, 0x69, 0xb9,
	0xa8, 0x68, 0xc7, 0xe7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x14, 0xff, 0xf7, 0xac, 0xcf, 0x08,
	0x00, 0x00,
}
