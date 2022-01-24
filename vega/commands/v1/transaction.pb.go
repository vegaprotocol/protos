// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vega/commands/v1/transaction.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type InputData struct {
	// A number to provide uniqueness to prevent accidental replays and,
	// in combination with `block_height`, deliberate attacks.
	// A nonce provides uniqueness for otherwise identical transactions,
	// ensuring that the transaction hash uniquely identifies a specific transaction.
	// Granted all other fields are equal, the nonce can either be a counter
	// or generated at random to submit multiple transactions within the same
	// block (see below), without being identified as replays.
	// Please note that Protocol Buffers do not have a canonical, unique encoding
	// and therefore different libraries or binaries may encode the same message
	// slightly differently, causing a different hash.
	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The block height at which the transaction was made.
	// This should be the current block height. The transaction will be valid
	// from the block and up to the `tolerance` block height.
	// Example: If the network has a tolerance of 150 blocks and `block_height`
	// is set to `200`, then the transaction will be valid until block `350`.
	// Note that a `block_height` that is ahead of the real block height will be
	// rejected. The tolerance can be queried from the chain's network parameters.
	// `block_height` prevents replay attacks in conjunction with `nonce` (see above).
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Types that are valid to be assigned to Command:
	//	*InputData_OrderSubmission
	//	*InputData_OrderCancellation
	//	*InputData_OrderAmendment
	//	*InputData_WithdrawSubmission
	//	*InputData_ProposalSubmission
	//	*InputData_VoteSubmission
	//	*InputData_LiquidityProvisionSubmission
	//	*InputData_DelegateSubmission
	//	*InputData_UndelegateSubmission
	//	*InputData_LiquidityProvisionCancellation
	//	*InputData_LiquidityProvisionAmendment
	//	*InputData_Transfer
	//	*InputData_CancelTransfer
	//	*InputData_NodeRegistration
	//	*InputData_NodeVote
	//	*InputData_NodeSignature
	//	*InputData_ChainEvent
	//	*InputData_KeyRotateSubmission
	//	*InputData_StateVariableProposal
	//	*InputData_OracleDataSubmission
	//	*InputData_RestoreSnapshotSubmission
	Command              isInputData_Command `protobuf_oneof:"command"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *InputData) Reset()         { *m = InputData{} }
func (m *InputData) String() string { return proto.CompactTextString(m) }
func (*InputData) ProtoMessage()    {}
func (*InputData) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d2b9785d3e1386, []int{0}
}

func (m *InputData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputData.Unmarshal(m, b)
}
func (m *InputData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputData.Marshal(b, m, deterministic)
}
func (m *InputData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputData.Merge(m, src)
}
func (m *InputData) XXX_Size() int {
	return xxx_messageInfo_InputData.Size(m)
}
func (m *InputData) XXX_DiscardUnknown() {
	xxx_messageInfo_InputData.DiscardUnknown(m)
}

var xxx_messageInfo_InputData proto.InternalMessageInfo

func (m *InputData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *InputData) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type isInputData_Command interface {
	isInputData_Command()
}

type InputData_OrderSubmission struct {
	OrderSubmission *OrderSubmission `protobuf:"bytes,1001,opt,name=order_submission,json=orderSubmission,proto3,oneof"`
}

type InputData_OrderCancellation struct {
	OrderCancellation *OrderCancellation `protobuf:"bytes,1002,opt,name=order_cancellation,json=orderCancellation,proto3,oneof"`
}

type InputData_OrderAmendment struct {
	OrderAmendment *OrderAmendment `protobuf:"bytes,1003,opt,name=order_amendment,json=orderAmendment,proto3,oneof"`
}

type InputData_WithdrawSubmission struct {
	WithdrawSubmission *WithdrawSubmission `protobuf:"bytes,1004,opt,name=withdraw_submission,json=withdrawSubmission,proto3,oneof"`
}

type InputData_ProposalSubmission struct {
	ProposalSubmission *ProposalSubmission `protobuf:"bytes,1005,opt,name=proposal_submission,json=proposalSubmission,proto3,oneof"`
}

type InputData_VoteSubmission struct {
	VoteSubmission *VoteSubmission `protobuf:"bytes,1006,opt,name=vote_submission,json=voteSubmission,proto3,oneof"`
}

type InputData_LiquidityProvisionSubmission struct {
	LiquidityProvisionSubmission *LiquidityProvisionSubmission `protobuf:"bytes,1007,opt,name=liquidity_provision_submission,json=liquidityProvisionSubmission,proto3,oneof"`
}

type InputData_DelegateSubmission struct {
	DelegateSubmission *DelegateSubmission `protobuf:"bytes,1008,opt,name=delegate_submission,json=delegateSubmission,proto3,oneof"`
}

type InputData_UndelegateSubmission struct {
	UndelegateSubmission *UndelegateSubmission `protobuf:"bytes,1009,opt,name=undelegate_submission,json=undelegateSubmission,proto3,oneof"`
}

type InputData_LiquidityProvisionCancellation struct {
	LiquidityProvisionCancellation *LiquidityProvisionCancellation `protobuf:"bytes,1010,opt,name=liquidity_provision_cancellation,json=liquidityProvisionCancellation,proto3,oneof"`
}

type InputData_LiquidityProvisionAmendment struct {
	LiquidityProvisionAmendment *LiquidityProvisionAmendment `protobuf:"bytes,1011,opt,name=liquidity_provision_amendment,json=liquidityProvisionAmendment,proto3,oneof"`
}

type InputData_Transfer struct {
	Transfer *Transfer `protobuf:"bytes,1012,opt,name=transfer,proto3,oneof"`
}

type InputData_CancelTransfer struct {
	CancelTransfer *CancelTransfer `protobuf:"bytes,1013,opt,name=cancel_transfer,json=cancelTransfer,proto3,oneof"`
}

type InputData_NodeRegistration struct {
	NodeRegistration *NodeRegistration `protobuf:"bytes,2001,opt,name=node_registration,json=nodeRegistration,proto3,oneof"`
}

type InputData_NodeVote struct {
	NodeVote *NodeVote `protobuf:"bytes,2002,opt,name=node_vote,json=nodeVote,proto3,oneof"`
}

type InputData_NodeSignature struct {
	NodeSignature *NodeSignature `protobuf:"bytes,2003,opt,name=node_signature,json=nodeSignature,proto3,oneof"`
}

type InputData_ChainEvent struct {
	ChainEvent *ChainEvent `protobuf:"bytes,2004,opt,name=chain_event,json=chainEvent,proto3,oneof"`
}

type InputData_KeyRotateSubmission struct {
	KeyRotateSubmission *KeyRotateSubmission `protobuf:"bytes,2005,opt,name=key_rotate_submission,json=keyRotateSubmission,proto3,oneof"`
}

type InputData_StateVariableProposal struct {
	StateVariableProposal *StateVariableProposal `protobuf:"bytes,2006,opt,name=state_variable_proposal,json=stateVariableProposal,proto3,oneof"`
}

type InputData_OracleDataSubmission struct {
	OracleDataSubmission *OracleDataSubmission `protobuf:"bytes,3001,opt,name=oracle_data_submission,json=oracleDataSubmission,proto3,oneof"`
}

type InputData_RestoreSnapshotSubmission struct {
	RestoreSnapshotSubmission *RestoreSnapshot `protobuf:"bytes,4001,opt,name=restore_snapshot_submission,json=restoreSnapshotSubmission,proto3,oneof"`
}

func (*InputData_OrderSubmission) isInputData_Command() {}

func (*InputData_OrderCancellation) isInputData_Command() {}

func (*InputData_OrderAmendment) isInputData_Command() {}

func (*InputData_WithdrawSubmission) isInputData_Command() {}

func (*InputData_ProposalSubmission) isInputData_Command() {}

func (*InputData_VoteSubmission) isInputData_Command() {}

func (*InputData_LiquidityProvisionSubmission) isInputData_Command() {}

func (*InputData_DelegateSubmission) isInputData_Command() {}

func (*InputData_UndelegateSubmission) isInputData_Command() {}

func (*InputData_LiquidityProvisionCancellation) isInputData_Command() {}

func (*InputData_LiquidityProvisionAmendment) isInputData_Command() {}

func (*InputData_Transfer) isInputData_Command() {}

func (*InputData_CancelTransfer) isInputData_Command() {}

func (*InputData_NodeRegistration) isInputData_Command() {}

func (*InputData_NodeVote) isInputData_Command() {}

func (*InputData_NodeSignature) isInputData_Command() {}

func (*InputData_ChainEvent) isInputData_Command() {}

func (*InputData_KeyRotateSubmission) isInputData_Command() {}

func (*InputData_StateVariableProposal) isInputData_Command() {}

func (*InputData_OracleDataSubmission) isInputData_Command() {}

func (*InputData_RestoreSnapshotSubmission) isInputData_Command() {}

func (m *InputData) GetCommand() isInputData_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *InputData) GetOrderSubmission() *OrderSubmission {
	if x, ok := m.GetCommand().(*InputData_OrderSubmission); ok {
		return x.OrderSubmission
	}
	return nil
}

func (m *InputData) GetOrderCancellation() *OrderCancellation {
	if x, ok := m.GetCommand().(*InputData_OrderCancellation); ok {
		return x.OrderCancellation
	}
	return nil
}

func (m *InputData) GetOrderAmendment() *OrderAmendment {
	if x, ok := m.GetCommand().(*InputData_OrderAmendment); ok {
		return x.OrderAmendment
	}
	return nil
}

func (m *InputData) GetWithdrawSubmission() *WithdrawSubmission {
	if x, ok := m.GetCommand().(*InputData_WithdrawSubmission); ok {
		return x.WithdrawSubmission
	}
	return nil
}

func (m *InputData) GetProposalSubmission() *ProposalSubmission {
	if x, ok := m.GetCommand().(*InputData_ProposalSubmission); ok {
		return x.ProposalSubmission
	}
	return nil
}

func (m *InputData) GetVoteSubmission() *VoteSubmission {
	if x, ok := m.GetCommand().(*InputData_VoteSubmission); ok {
		return x.VoteSubmission
	}
	return nil
}

func (m *InputData) GetLiquidityProvisionSubmission() *LiquidityProvisionSubmission {
	if x, ok := m.GetCommand().(*InputData_LiquidityProvisionSubmission); ok {
		return x.LiquidityProvisionSubmission
	}
	return nil
}

func (m *InputData) GetDelegateSubmission() *DelegateSubmission {
	if x, ok := m.GetCommand().(*InputData_DelegateSubmission); ok {
		return x.DelegateSubmission
	}
	return nil
}

func (m *InputData) GetUndelegateSubmission() *UndelegateSubmission {
	if x, ok := m.GetCommand().(*InputData_UndelegateSubmission); ok {
		return x.UndelegateSubmission
	}
	return nil
}

func (m *InputData) GetLiquidityProvisionCancellation() *LiquidityProvisionCancellation {
	if x, ok := m.GetCommand().(*InputData_LiquidityProvisionCancellation); ok {
		return x.LiquidityProvisionCancellation
	}
	return nil
}

func (m *InputData) GetLiquidityProvisionAmendment() *LiquidityProvisionAmendment {
	if x, ok := m.GetCommand().(*InputData_LiquidityProvisionAmendment); ok {
		return x.LiquidityProvisionAmendment
	}
	return nil
}

func (m *InputData) GetTransfer() *Transfer {
	if x, ok := m.GetCommand().(*InputData_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *InputData) GetCancelTransfer() *CancelTransfer {
	if x, ok := m.GetCommand().(*InputData_CancelTransfer); ok {
		return x.CancelTransfer
	}
	return nil
}

func (m *InputData) GetNodeRegistration() *NodeRegistration {
	if x, ok := m.GetCommand().(*InputData_NodeRegistration); ok {
		return x.NodeRegistration
	}
	return nil
}

func (m *InputData) GetNodeVote() *NodeVote {
	if x, ok := m.GetCommand().(*InputData_NodeVote); ok {
		return x.NodeVote
	}
	return nil
}

func (m *InputData) GetNodeSignature() *NodeSignature {
	if x, ok := m.GetCommand().(*InputData_NodeSignature); ok {
		return x.NodeSignature
	}
	return nil
}

func (m *InputData) GetChainEvent() *ChainEvent {
	if x, ok := m.GetCommand().(*InputData_ChainEvent); ok {
		return x.ChainEvent
	}
	return nil
}

func (m *InputData) GetKeyRotateSubmission() *KeyRotateSubmission {
	if x, ok := m.GetCommand().(*InputData_KeyRotateSubmission); ok {
		return x.KeyRotateSubmission
	}
	return nil
}

func (m *InputData) GetStateVariableProposal() *StateVariableProposal {
	if x, ok := m.GetCommand().(*InputData_StateVariableProposal); ok {
		return x.StateVariableProposal
	}
	return nil
}

func (m *InputData) GetOracleDataSubmission() *OracleDataSubmission {
	if x, ok := m.GetCommand().(*InputData_OracleDataSubmission); ok {
		return x.OracleDataSubmission
	}
	return nil
}

func (m *InputData) GetRestoreSnapshotSubmission() *RestoreSnapshot {
	if x, ok := m.GetCommand().(*InputData_RestoreSnapshotSubmission); ok {
		return x.RestoreSnapshotSubmission
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InputData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InputData_OrderSubmission)(nil),
		(*InputData_OrderCancellation)(nil),
		(*InputData_OrderAmendment)(nil),
		(*InputData_WithdrawSubmission)(nil),
		(*InputData_ProposalSubmission)(nil),
		(*InputData_VoteSubmission)(nil),
		(*InputData_LiquidityProvisionSubmission)(nil),
		(*InputData_DelegateSubmission)(nil),
		(*InputData_UndelegateSubmission)(nil),
		(*InputData_LiquidityProvisionCancellation)(nil),
		(*InputData_LiquidityProvisionAmendment)(nil),
		(*InputData_Transfer)(nil),
		(*InputData_CancelTransfer)(nil),
		(*InputData_NodeRegistration)(nil),
		(*InputData_NodeVote)(nil),
		(*InputData_NodeSignature)(nil),
		(*InputData_ChainEvent)(nil),
		(*InputData_KeyRotateSubmission)(nil),
		(*InputData_StateVariableProposal)(nil),
		(*InputData_OracleDataSubmission)(nil),
		(*InputData_RestoreSnapshotSubmission)(nil),
	}
}

// Represents a transaction to be sent to Vega.
type Transaction struct {
	// One of the set of Vega commands (proto marshalled).
	InputData []byte `protobuf:"bytes,1,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"`
	// The signature of the inputData.
	Signature *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// The sender of the transaction.
	// Any of the following would be valid:
	//
	// Types that are valid to be assigned to From:
	//	*Transaction_Address
	//	*Transaction_PubKey
	From isTransaction_From `protobuf_oneof:"from"`
	// A version of the transaction, to be used in the future in case changes are implemented
	// to the Transaction format.
	Version              uint32   `protobuf:"varint,2000,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d2b9785d3e1386, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetInputData() []byte {
	if m != nil {
		return m.InputData
	}
	return nil
}

func (m *Transaction) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type isTransaction_From interface {
	isTransaction_From()
}

type Transaction_Address struct {
	Address string `protobuf:"bytes,1001,opt,name=address,proto3,oneof"`
}

type Transaction_PubKey struct {
	PubKey string `protobuf:"bytes,1002,opt,name=pub_key,json=pubKey,proto3,oneof"`
}

func (*Transaction_Address) isTransaction_From() {}

func (*Transaction_PubKey) isTransaction_From() {}

func (m *Transaction) GetFrom() isTransaction_From {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetAddress() string {
	if x, ok := m.GetFrom().(*Transaction_Address); ok {
		return x.Address
	}
	return ""
}

func (m *Transaction) GetPubKey() string {
	if x, ok := m.GetFrom().(*Transaction_PubKey); ok {
		return x.PubKey
	}
	return ""
}

func (m *Transaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transaction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transaction_Address)(nil),
		(*Transaction_PubKey)(nil),
	}
}

// A signature to authenticate a transaction and to be verified by the Vega
// network.
type Signature struct {
	// The bytes of the signature (hex-encoded).
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// The algorithm used to create the signature.
	Algo string `protobuf:"bytes,2,opt,name=algo,proto3" json:"algo,omitempty"`
	// The version of the signature used to create the signature.
	Version              uint32   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d2b9785d3e1386, []int{2}
}

func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Signature) GetAlgo() string {
	if m != nil {
		return m.Algo
	}
	return ""
}

func (m *Signature) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*InputData)(nil), "vega.commands.v1.InputData")
	proto.RegisterType((*Transaction)(nil), "vega.commands.v1.Transaction")
	proto.RegisterType((*Signature)(nil), "vega.commands.v1.Signature")
}

func init() {
	proto.RegisterFile("vega/commands/v1/transaction.proto", fileDescriptor_73d2b9785d3e1386)
}

var fileDescriptor_73d2b9785d3e1386 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc7, 0x9d, 0x92, 0x66, 0xf1, 0x49, 0x9b, 0xa4, 0xd3, 0x84, 0xba, 0x49, 0x9b, 0xa4, 0xcb,
	0x57, 0x91, 0x60, 0x97, 0xc2, 0x55, 0xb9, 0x82, 0x52, 0x24, 0x57, 0xad, 0x48, 0x35, 0x69, 0x0b,
	0x42, 0x08, 0x6b, 0xd6, 0x9e, 0xee, 0x8e, 0xe2, 0x9d, 0x31, 0xe3, 0xb1, 0xa3, 0xbd, 0xe0, 0x55,
	0x90, 0x78, 0x0c, 0xae, 0x79, 0x01, 0xbe, 0xdf, 0x81, 0xef, 0xef, 0x7b, 0x34, 0x33, 0xf6, 0xae,
	0xbd, 0xb6, 0xab, 0xde, 0xd9, 0xff, 0xf9, 0x9f, 0xdf, 0xf1, 0x39, 0x9e, 0xf1, 0x31, 0xf4, 0x73,
	0x3a, 0x26, 0xc3, 0x50, 0x4c, 0xa7, 0x84, 0x47, 0xe9, 0x30, 0xbf, 0x31, 0x54, 0x92, 0xf0, 0x94,
	0x84, 0x8a, 0x09, 0x3e, 0x48, 0xa4, 0x50, 0x02, 0x6d, 0x69, 0xcf, 0xa0, 0xf4, 0x0c, 0xf2, 0x1b,
	0xbb, 0x07, 0x8d, 0xa8, 0xf9, 0xaa, 0x09, 0xd9, 0x7d, 0xa5, 0x61, 0xc8, 0x49, 0xcc, 0x22, 0xa2,
	0x84, 0x0c, 0x96, 0xac, 0xfb, 0x0d, 0xab, 0x90, 0x24, 0x8c, 0x69, 0xb1, 0xde, 0xff, 0x7c, 0x13,
	0xdc, 0x3b, 0x3c, 0xc9, 0xd4, 0x6d, 0xa2, 0x08, 0xda, 0x86, 0xb3, 0x5c, 0xf0, 0x90, 0x7a, 0x2b,
	0x87, 0x2b, 0xd7, 0x57, 0xb1, 0xbd, 0x41, 0xd7, 0xe0, 0xdc, 0x28, 0x16, 0xe1, 0x49, 0x30, 0xa1,
	0x6c, 0x3c, 0x51, 0xde, 0x19, 0xb3, 0xb8, 0x6e, 0x34, 0xdf, 0x48, 0xe8, 0x08, 0xb6, 0x84, 0x8c,
	0xa8, 0x0c, 0xd2, 0x6c, 0x34, 0x65, 0x69, 0xca, 0x04, 0xf7, 0x7e, 0xea, 0x1d, 0xae, 0x5c, 0x5f,
	0x7f, 0xe3, 0xda, 0x60, 0xb9, 0xc0, 0xc1, 0x91, 0xb6, 0x1e, 0xcf, 0x9d, 0xbe, 0x83, 0x37, 0x45,
	0x5d, 0x42, 0x0f, 0x01, 0x59, 0x60, 0x48, 0x78, 0x48, 0xe3, 0x98, 0xe8, 0x8e, 0x79, 0x3f, 0x5b,
	0xe4, 0xf3, 0x1d, 0xc8, 0x77, 0x2b, 0x5e, 0xdf, 0xc1, 0x17, 0xc4, 0xb2, 0x88, 0xee, 0x81, 0xcd,
	0x14, 0x90, 0x29, 0xe5, 0xd1, 0x94, 0x72, 0xe5, 0xfd, 0x62, 0x99, 0x87, 0x1d, 0xcc, 0x77, 0x4a,
	0xa3, 0xef, 0xe0, 0x0d, 0x51, 0x53, 0xd0, 0x87, 0x70, 0xf1, 0x94, 0xa9, 0x49, 0x24, 0xc9, 0x69,
	0xb5, 0xf0, 0x5f, 0x2d, 0xf1, 0x85, 0x26, 0xf1, 0x83, 0xc2, 0x5d, 0xab, 0x1d, 0x9d, 0x36, 0x54,
	0x4d, 0x4e, 0xa4, 0x48, 0x44, 0x4a, 0xe2, 0x2a, 0xf9, 0xb7, 0x4e, 0xf2, 0xfd, 0xc2, 0x5d, 0x27,
	0x27, 0x0d, 0x55, 0x77, 0x20, 0x17, 0x8a, 0x56, 0xa9, 0xbf, 0x77, 0x76, 0xe0, 0x91, 0x50, 0xb4,
	0x46, 0xdc, 0xc8, 0x6b, 0x0a, 0x3a, 0x85, 0xfd, 0x98, 0x7d, 0x9a, 0xb1, 0x88, 0xa9, 0x59, 0x90,
	0x48, 0x91, 0x33, 0x2d, 0x57, 0xe1, 0x7f, 0x58, 0xf8, 0xa0, 0x09, 0xbf, 0x57, 0x06, 0xde, 0x2f,
	0xe3, 0x6a, 0xa9, 0xae, 0xc4, 0x4f, 0x58, 0xd7, 0x0d, 0x8a, 0x68, 0x4c, 0xc7, 0xa4, 0x5e, 0xca,
	0x9f, 0x9d, 0x0d, 0xba, 0x5d, 0xb8, 0xeb, 0x0d, 0x8a, 0x1a, 0x2a, 0xfa, 0x04, 0x76, 0x32, 0xde,
	0xc6, 0xfe, 0xcb, 0xb2, 0x5f, 0x6a, 0xb2, 0x1f, 0xf2, 0xa8, 0x8d, 0xbe, 0x9d, 0xb5, 0xe8, 0xe8,
	0x33, 0x38, 0x6c, 0x6b, 0x59, 0x6d, 0x9f, 0xff, 0x6d, 0x53, 0xbd, 0xfe, 0x34, 0x4d, 0x5b, 0xda,
	0xf4, 0xfb, 0xf1, 0x13, 0x1d, 0x48, 0xc1, 0xd5, 0xb6, 0xf4, 0x8b, 0xf3, 0xf0, 0x8f, 0xcd, 0xfd,
	0xda, 0xd3, 0xe4, 0xae, 0x1e, 0x8e, 0xbd, 0xb8, 0x7b, 0x19, 0xdd, 0x84, 0x67, 0xcd, 0x97, 0xef,
	0x31, 0x95, 0xde, 0xbf, 0x36, 0xc1, 0x6e, 0x33, 0xc1, 0x83, 0xc2, 0xe2, 0x3b, 0x78, 0x6e, 0xd7,
	0x1b, 0xd6, 0xf6, 0x26, 0x98, 0x13, 0xfe, 0xeb, 0xdc, 0xb0, 0xb6, 0xd4, 0x0a, 0x67, 0x23, 0xac,
	0x29, 0x08, 0xc3, 0x05, 0x2e, 0x22, 0x1a, 0x48, 0x3a, 0x66, 0xa9, 0x92, 0xb6, 0xdd, 0xdf, 0x6c,
	0x1a, 0x5e, 0xbf, 0xc9, 0x7b, 0x5f, 0x44, 0x14, 0x57, 0xac, 0xbe, 0x83, 0xb7, 0xf8, 0x92, 0x86,
	0xde, 0x02, 0xd7, 0x30, 0xf5, 0xd9, 0xf0, 0xbe, 0xdd, 0xec, 0xaa, 0x4e, 0xb3, 0xf4, 0x81, 0xd2,
	0xd5, 0xf1, 0xe2, 0x1a, 0xdd, 0x81, 0x0d, 0x13, 0x9b, 0xb2, 0x31, 0x27, 0x2a, 0x93, 0xd4, 0xfb,
	0xce, 0x02, 0x0e, 0xda, 0x01, 0xc7, 0xa5, 0xcf, 0x77, 0xf0, 0x79, 0x5e, 0x15, 0xd0, 0xdb, 0xb0,
	0x1e, 0x4e, 0x08, 0xe3, 0x01, 0xcd, 0xf5, 0x7b, 0xfc, 0xde, 0x72, 0xae, 0xb4, 0x34, 0x49, 0xbb,
	0xde, 0xcb, 0xed, 0x6b, 0x83, 0x70, 0x7e, 0x87, 0x3e, 0x86, 0x9d, 0x13, 0x3a, 0x0b, 0xa4, 0x50,
	0x4b, 0x5b, 0xff, 0x07, 0xcb, 0x7a, 0xb1, 0xc9, 0xba, 0x4b, 0x67, 0xd8, 0xd8, 0x6b, 0x3b, 0xff,
	0xe2, 0x49, 0x53, 0x46, 0x23, 0xb8, 0x94, 0x1a, 0x70, 0x4e, 0x24, 0x23, 0xa3, 0x98, 0x06, 0xe5,
	0xe7, 0xc9, 0xfb, 0xd1, 0xf2, 0x5f, 0x6e, 0xf2, 0x8f, 0x75, 0xc4, 0xa3, 0x22, 0xa0, 0xfc, 0xc8,
	0xf9, 0x0e, 0xde, 0x49, 0xdb, 0x16, 0x50, 0x00, 0xcf, 0xd9, 0xf9, 0x16, 0x44, 0x44, 0x91, 0x6a,
	0x09, 0x5f, 0x5e, 0xea, 0x3a, 0xbd, 0x47, 0x26, 0x40, 0xcf, 0xbf, 0xfa, 0xe9, 0x15, 0x2d, 0x3a,
	0x8a, 0x60, 0x4f, 0xd2, 0x54, 0x09, 0x49, 0x83, 0x94, 0x93, 0x24, 0x9d, 0x08, 0x55, 0xcd, 0xf2,
	0xc5, 0x41, 0xd7, 0xcc, 0xc3, 0x36, 0xea, 0xb8, 0x08, 0xf2, 0x1d, 0x7c, 0x59, 0xd6, 0xa5, 0x45,
	0x96, 0x5b, 0x2e, 0xf4, 0x8a, 0xd8, 0xfe, 0x57, 0x2b, 0xb0, 0xfe, 0x60, 0xf1, 0xd3, 0x80, 0xae,
	0x02, 0x30, 0x3d, 0xaf, 0x4d, 0x81, 0x66, 0x4e, 0x9f, 0xc3, 0x2e, 0x9b, 0x4f, 0xf0, 0x9b, 0xe0,
	0x2e, 0xb6, 0xd2, 0x19, 0xf3, 0x30, 0x7b, 0x2d, 0x5d, 0x2d, 0x2d, 0x78, 0xe1, 0x46, 0x7b, 0xd0,
	0x23, 0x51, 0x24, 0x69, 0x9a, 0xda, 0xd1, 0xed, 0xfa, 0x0e, 0x2e, 0x15, 0xb4, 0x0b, 0xbd, 0x24,
	0x1b, 0x05, 0x27, 0x74, 0x66, 0x87, 0xb0, 0x5e, 0x5c, 0x4b, 0xb2, 0xd1, 0x5d, 0x3a, 0x43, 0x97,
	0xa1, 0x97, 0x53, 0x69, 0xea, 0xff, 0x5a, 0xbf, 0xc8, 0xf3, 0xb8, 0xbc, 0xbf, 0xb5, 0x06, 0xab,
	0x8f, 0xa5, 0x98, 0xf6, 0x8f, 0xc0, 0x5d, 0x6c, 0xd4, 0x6d, 0x38, 0x9b, 0x93, 0x38, 0xb3, 0x7f,
	0x19, 0x2e, 0xb6, 0x37, 0x08, 0xc1, 0x2a, 0x89, 0xc7, 0xc2, 0x3c, 0xb4, 0x8b, 0xcd, 0x35, 0xf2,
	0x16, 0xe4, 0x67, 0xea, 0xe0, 0xc1, 0x47, 0xaf, 0x86, 0x22, 0xa2, 0xa6, 0x34, 0xf3, 0x27, 0x13,
	0x8a, 0x78, 0xc0, 0xc4, 0xd0, 0x5c, 0xa7, 0xc3, 0xe5, 0xbf, 0x9e, 0xd1, 0x9a, 0x59, 0x78, 0xf3,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x1d, 0x75, 0xcc, 0x92, 0x09, 0x00, 0x00,
}
