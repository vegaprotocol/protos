// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/commands/v1/commands.proto

package v1

import (
	vega "code.vegaprotocol.io/protos/vega"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
	UndelegateSubmission_Now          UndelegateSubmission_Method = 0
	UndelegateSubmission_AtEndOfEpoch UndelegateSubmission_Method = 1
	UndelegateSubmission_InAnger      UndelegateSubmission_Method = 2
)

var UndelegateSubmission_Method_name = map[int32]string{
	0: "Now",
	1: "AtEndOfEpoch",
	2: "InAnger",
}

var UndelegateSubmission_Method_value = map[string]int32{
	"Now":          0,
	"AtEndOfEpoch": 1,
	"InAnger":      2,
}

func (x UndelegateSubmission_Method) String() string {
	return proto.EnumName(UndelegateSubmission_Method_name, int32(x))
}

func (UndelegateSubmission_Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{8, 0}
}

// An order submission is a request to submit or create a new order on Vega
type OrderSubmission struct {
	// Market identifier for the order, required field
	MarketId string `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// Price for the order, the price is an integer, for example `123456` is a correctly
	// formatted price of `1.23456` assuming market configured to 5 decimal places,
	// , required field for limit orders, however it is not required for market orders
	Price uint64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
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

func (m *OrderSubmission) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
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
	PeggedOffset *wrapperspb.Int64Value `protobuf:"bytes,7,opt,name=pegged_offset,json=peggedOffset,proto3" json:"pegged_offset,omitempty"`
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

func (m *OrderAmendment) GetPeggedOffset() *wrapperspb.Int64Value {
	if m != nil {
		return m.PeggedOffset
	}
	return nil
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
	CommitmentAmount uint64 `protobuf:"varint,2,opt,name=commitment_amount,json=commitmentAmount,proto3" json:"commitment_amount,omitempty"`
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

func (m *LiquidityProvisionSubmission) GetCommitmentAmount() uint64 {
	if m != nil {
		return m.CommitmentAmount
	}
	return 0
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

// Represents the submission request to withdraw funds for a party on Vega
type WithdrawSubmission struct {
	// The amount to be withdrawn
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
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
	return fileDescriptor_9c3a53c66a940c51, []int{4}
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

func (m *WithdrawSubmission) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
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
	return fileDescriptor_9c3a53c66a940c51, []int{5}
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
	return fileDescriptor_9c3a53c66a940c51, []int{6}
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
	Amount               uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegateSubmission) Reset()         { *m = DelegateSubmission{} }
func (m *DelegateSubmission) String() string { return proto.CompactTextString(m) }
func (*DelegateSubmission) ProtoMessage()    {}
func (*DelegateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{7}
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

func (m *DelegateSubmission) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type UndelegateSubmission struct {
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// optional, if not specified = ALL
	Amount               uint64                      `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Method               UndelegateSubmission_Method `protobuf:"varint,3,opt,name=method,proto3,enum=vega.commands.v1.UndelegateSubmission_Method" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *UndelegateSubmission) Reset()         { *m = UndelegateSubmission{} }
func (m *UndelegateSubmission) String() string { return proto.CompactTextString(m) }
func (*UndelegateSubmission) ProtoMessage()    {}
func (*UndelegateSubmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c3a53c66a940c51, []int{8}
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

func (m *UndelegateSubmission) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UndelegateSubmission) GetMethod() UndelegateSubmission_Method {
	if m != nil {
		return m.Method
	}
	return UndelegateSubmission_Now
}

func init() {
	proto.RegisterEnum("vega.commands.v1.UndelegateSubmission_Method", UndelegateSubmission_Method_name, UndelegateSubmission_Method_value)
	proto.RegisterType((*OrderSubmission)(nil), "vega.commands.v1.OrderSubmission")
	proto.RegisterType((*OrderCancellation)(nil), "vega.commands.v1.OrderCancellation")
	proto.RegisterType((*OrderAmendment)(nil), "vega.commands.v1.OrderAmendment")
	proto.RegisterType((*LiquidityProvisionSubmission)(nil), "vega.commands.v1.LiquidityProvisionSubmission")
	proto.RegisterType((*WithdrawSubmission)(nil), "vega.commands.v1.WithdrawSubmission")
	proto.RegisterType((*ProposalSubmission)(nil), "vega.commands.v1.ProposalSubmission")
	proto.RegisterType((*VoteSubmission)(nil), "vega.commands.v1.VoteSubmission")
	proto.RegisterType((*DelegateSubmission)(nil), "vega.commands.v1.DelegateSubmission")
	proto.RegisterType((*UndelegateSubmission)(nil), "vega.commands.v1.UndelegateSubmission")
}

func init() { proto.RegisterFile("vega/commands/v1/commands.proto", fileDescriptor_9c3a53c66a940c51) }

var fileDescriptor_9c3a53c66a940c51 = []byte{
	// 888 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0x1b, 0x45,
	0x14, 0xed, 0xfa, 0x33, 0xbe, 0xdb, 0x3a, 0x9b, 0x21, 0x85, 0x25, 0xfd, 0x88, 0xbb, 0x45, 0x60,
	0x01, 0x59, 0x13, 0x53, 0xf5, 0x85, 0x97, 0xba, 0x34, 0x48, 0x16, 0x94, 0x44, 0xdb, 0x52, 0x10,
	0x2f, 0xd6, 0xd8, 0x73, 0xbd, 0x19, 0x75, 0x77, 0x67, 0x99, 0x1d, 0xdb, 0x09, 0xaf, 0xfc, 0x39,
	0x7e, 0x05, 0x48, 0xfc, 0x08, 0x9e, 0xd1, 0xcc, 0xac, 0x1d, 0xdb, 0x41, 0x88, 0x48, 0x7d, 0x9b,
	0x39, 0xe7, 0xde, 0x73, 0x67, 0xee, 0x99, 0xbb, 0x0b, 0x87, 0x73, 0x8c, 0x69, 0x6f, 0x22, 0xd2,
	0x94, 0x66, 0xac, 0xe8, 0xcd, 0x8f, 0x57, 0xeb, 0x30, 0x97, 0x42, 0x09, 0xe2, 0xe9, 0x80, 0x70,
	0x05, 0xce, 0x8f, 0x0f, 0xee, 0x9a, 0x94, 0x58, 0xcc, 0x51, 0x66, 0x34, 0x9b, 0xa0, 0x0d, 0x3c,
	0xd8, 0x35, 0xb0, 0x89, 0xb6, 0xc0, 0xc3, 0x58, 0x88, 0x38, 0xc1, 0x9e, 0xd9, 0x8d, 0x67, 0xd3,
	0xde, 0x42, 0xd2, 0x3c, 0x47, 0x59, 0x2a, 0x1f, 0x3c, 0x8d, 0xb9, 0x3a, 0x9f, 0x8d, 0xb5, 0x76,
	0x2f, 0x5d, 0x70, 0xf5, 0x56, 0x2c, 0x7a, 0xb1, 0x38, 0x32, 0xe4, 0xd1, 0x9c, 0x26, 0x9c, 0x51,
	0x25, 0x64, 0xd1, 0x5b, 0x2d, 0x6d, 0x5e, 0xf0, 0x47, 0x05, 0x76, 0x4f, 0x25, 0x43, 0xf9, 0x6a,
	0x36, 0x4e, 0x79, 0x51, 0x70, 0x91, 0x91, 0xc7, 0xd0, 0x4a, 0xa9, 0x7c, 0x8b, 0x6a, 0xc4, 0x99,
	0xef, 0x74, 0x9c, 0x6e, 0xeb, 0x79, 0xe3, 0xaf, 0x3f, 0x0f, 0x2b, 0x3f, 0x39, 0xd1, 0x8e, 0x25,
	0x86, 0x8c, 0xec, 0x43, 0x3d, 0x97, 0x7c, 0x82, 0x7e, 0xa5, 0xe3, 0x74, 0x6b, 0x91, 0xdd, 0x90,
	0x03, 0xa8, 0x15, 0xfc, 0x57, 0xf4, 0xab, 0x1a, 0xb4, 0x59, 0xde, 0xad, 0xc8, 0x60, 0xe4, 0xa1,
	0xe6, 0x18, 0xfa, 0xb5, 0x8e, 0xd3, 0x6d, 0xf7, 0x21, 0x34, 0xb7, 0x7b, 0xc5, 0x19, 0x46, 0x06,
	0x27, 0x5f, 0xc1, 0x1d, 0xc5, 0x53, 0x1c, 0xf1, 0x6c, 0x34, 0x15, 0x72, 0x82, 0x7e, 0xdd, 0x04,
	0x7e, 0x60, 0x03, 0xcd, 0x21, 0xc3, 0xd7, 0x3c, 0xc5, 0x61, 0xf6, 0x8d, 0xa6, 0x23, 0x57, 0x5d,
	0x6d, 0xc8, 0x03, 0x00, 0xbc, 0xc8, 0xb9, 0xc4, 0x62, 0x44, 0x95, 0xdf, 0xe8, 0x38, 0xdd, 0x6a,
	0xd4, 0x2a, 0x91, 0x81, 0x22, 0x1f, 0x41, 0x4d, 0x5d, 0xe6, 0xe8, 0x37, 0x8d, 0xa4, 0xb7, 0x21,
	0x79, 0x99, 0x63, 0x64, 0x58, 0x72, 0x1f, 0x5a, 0x12, 0xa7, 0x28, 0x31, 0x9b, 0xa0, 0xbf, 0xa3,
	0x2f, 0x1e, 0x5d, 0x01, 0xe4, 0x09, 0xdc, 0xce, 0x31, 0x8e, 0x91, 0x8d, 0x84, 0x4e, 0xf4, 0x5b,
	0x1d, 0xa7, 0xeb, 0xf6, 0xf7, 0xac, 0xd6, 0x99, 0x61, 0x8c, 0x62, 0xe4, 0xe6, 0x57, 0x9b, 0xe0,
	0x5b, 0xd8, 0x33, 0x8b, 0xaf, 0xb5, 0xbb, 0x49, 0x42, 0x95, 0xee, 0xf0, 0x87, 0xb0, 0x63, 0x34,
	0x56, 0x0d, 0x8e, 0x9a, 0x66, 0x3f, 0x64, 0xe4, 0xde, 0x7a, 0xf3, 0x2b, 0x86, 0x5b, 0x35, 0x3d,
	0xf8, 0xad, 0x0a, 0x6d, 0xa3, 0x36, 0x48, 0x31, 0x63, 0x29, 0x66, 0x8a, 0x3c, 0xda, 0x96, 0x5a,
	0x79, 0xf5, 0xbf, 0x24, 0xc9, 0xa3, 0xa5, 0x8f, 0x55, 0x73, 0x1d, 0xb7, 0xbc, 0x8e, 0x86, 0x96,
	0xa6, 0x3e, 0x00, 0xd0, 0x06, 0x8e, 0x18, 0x26, 0x8a, 0x1a, 0xfb, 0xaa, 0x51, 0x4b, 0x23, 0x2f,
	0x34, 0x40, 0xc2, 0x8d, 0xd6, 0xd7, 0x8d, 0xcc, 0xae, 0x95, 0xd1, 0x76, 0x15, 0x8a, 0xa6, 0xf9,
	0xba, 0x17, 0xd7, 0x7c, 0x6e, 0xdc, 0xc0, 0xe7, 0x67, 0x70, 0x67, 0x69, 0xc2, 0x74, 0x5a, 0xa0,
	0x32, 0x8e, 0xba, 0xfd, 0x7b, 0xa1, 0x9d, 0x8f, 0x70, 0x39, 0x1f, 0xe1, 0x30, 0x53, 0x4f, 0x9f,
	0xbc, 0xa1, 0xc9, 0x0c, 0xa3, 0xd2, 0xb6, 0x53, 0x93, 0x40, 0x9e, 0x81, 0x57, 0x2a, 0x6c, 0x7a,
	0xdd, 0xee, 0xdf, 0x5d, 0xb7, 0x32, 0x5a, 0x92, 0xd1, 0x6e, 0xbe, 0x09, 0x04, 0x7f, 0x3b, 0x70,
	0xff, 0x3b, 0xfe, 0xcb, 0x8c, 0x33, 0xae, 0x2e, 0xcf, 0xa4, 0x98, 0x73, 0x3d, 0x36, 0x37, 0x1d,
	0xa0, 0xcf, 0x60, 0x4f, 0x7f, 0x08, 0xb8, 0xd2, 0x36, 0x8e, 0x68, 0x2a, 0x66, 0x99, 0x2a, 0x87,
	0xc9, 0xbb, 0x22, 0x06, 0x06, 0x27, 0x1e, 0x54, 0xa7, 0x68, 0x3d, 0x6a, 0x45, 0x7a, 0x49, 0x3e,
	0x85, 0x7a, 0x81, 0x49, 0x52, 0xf8, 0xb5, 0x4e, 0xb5, 0xeb, 0xf6, 0xf7, 0xed, 0xd9, 0x57, 0xc7,
	0xb2, 0x2f, 0xd1, 0x86, 0x90, 0x2e, 0xd4, 0xc6, 0xb3, 0xcb, 0xc2, 0xaf, 0xff, 0x47, 0xa8, 0x89,
	0xd8, 0x9c, 0x80, 0xc6, 0xd6, 0x04, 0x04, 0x31, 0x90, 0x1f, 0xb9, 0x3a, 0x67, 0x92, 0x2e, 0xd6,
	0x6e, 0xfb, 0x3e, 0x34, 0xca, 0xd3, 0x3b, 0xe6, 0xf4, 0xe5, 0x4e, 0x7f, 0x21, 0x68, 0xa1, 0x2d,
	0xb2, 0x4f, 0xce, 0x6e, 0xc8, 0x63, 0xa8, 0xe2, 0x85, 0x2a, 0x5f, 0x5b, 0x39, 0x3c, 0x4b, 0xd1,
	0x93, 0x0b, 0x15, 0x69, 0x36, 0x40, 0x20, 0x67, 0x52, 0xe4, 0xa2, 0xa0, 0xc9, 0x5a, 0xa1, 0x8d,
	0xc3, 0x39, 0xdb, 0xe3, 0x79, 0x0c, 0x75, 0x85, 0x32, 0x2d, 0x4c, 0x39, 0xb7, 0xff, 0xde, 0xf2,
	0x21, 0x5b, 0x99, 0xd7, 0x9a, 0xb2, 0x2e, 0x74, 0x9c, 0xc8, 0x46, 0x06, 0x14, 0xda, 0x6f, 0x84,
	0xc2, 0xb5, 0x12, 0x9f, 0x80, 0x9b, 0x97, 0x19, 0xd7, 0xbd, 0x83, 0x25, 0x35, 0x64, 0xe4, 0x63,
	0xa8, 0xcf, 0xf5, 0xe3, 0x32, 0xd5, 0x56, 0x5f, 0x14, 0xad, 0x16, 0xda, 0x47, 0x67, 0xe9, 0xe0,
	0x25, 0x90, 0x17, 0x98, 0x60, 0x4c, 0x37, 0xca, 0x1c, 0x42, 0x33, 0x13, 0x0c, 0xaf, 0x97, 0x68,
	0x68, 0x78, 0xc8, 0xd6, 0x7a, 0x5a, 0x59, 0xef, 0x69, 0xf0, 0xbb, 0x03, 0xfb, 0x3f, 0x64, 0xec,
	0xdd, 0x29, 0x92, 0x13, 0x68, 0xa4, 0xa8, 0xce, 0x05, 0x33, 0x96, 0xb4, 0xfb, 0x47, 0xe1, 0xf6,
	0x3f, 0x2a, 0xfc, 0xb7, 0x82, 0xe1, 0x4b, 0x93, 0x14, 0x95, 0xc9, 0xc1, 0x17, 0xd0, 0xb0, 0x08,
	0x69, 0x42, 0xf5, 0x7b, 0xb1, 0xf0, 0x6e, 0x11, 0x0f, 0x6e, 0x0f, 0xd4, 0x49, 0xc6, 0x4e, 0xa7,
	0x27, 0xb9, 0x98, 0x9c, 0x7b, 0x0e, 0x71, 0xa1, 0x39, 0xcc, 0x06, 0x59, 0x8c, 0xd2, 0xab, 0x3c,
	0x0f, 0x7f, 0xfe, 0x7c, 0x22, 0x18, 0x9a, 0x72, 0x66, 0x6a, 0x27, 0x22, 0x09, 0xb9, 0xb0, 0x7f,
	0xb8, 0xa2, 0xb7, 0xfd, 0x2b, 0x1d, 0x37, 0x0c, 0xf1, 0xe5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3a, 0x1d, 0x54, 0x52, 0x65, 0x07, 0x00, 0x00,
}
