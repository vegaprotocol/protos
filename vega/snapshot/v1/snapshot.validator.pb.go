// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vega/snapshot/v1/snapshot.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/protos/vega"
	_ "code.vegaprotocol.io/protos/vega/commands/v1"
	_ "code.vegaprotocol.io/protos/vega/events/v1"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Snapshot) Validate() error {
	return nil
}
func (this *NodeHash) Validate() error {
	return nil
}
func (this *Metadata) Validate() error {
	for _, item := range this.NodeHashes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeHashes", err)
			}
		}
	}
	return nil
}
func (this *Chunk) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *Payload) Validate() error {
	if oneOfNester, ok := this.GetData().(*Payload_ActiveAssets); ok {
		if oneOfNester.ActiveAssets != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ActiveAssets); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ActiveAssets", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_PendingAssets); ok {
		if oneOfNester.PendingAssets != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PendingAssets); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PendingAssets", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_BankingWithdrawals); ok {
		if oneOfNester.BankingWithdrawals != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.BankingWithdrawals); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BankingWithdrawals", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_BankingDeposits); ok {
		if oneOfNester.BankingDeposits != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.BankingDeposits); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BankingDeposits", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_BankingSeen); ok {
		if oneOfNester.BankingSeen != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.BankingSeen); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BankingSeen", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_BankingAssetActions); ok {
		if oneOfNester.BankingAssetActions != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.BankingAssetActions); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BankingAssetActions", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_Checkpoint); ok {
		if oneOfNester.Checkpoint != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Checkpoint); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Checkpoint", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_CollateralAccounts); ok {
		if oneOfNester.CollateralAccounts != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.CollateralAccounts); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CollateralAccounts", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_CollateralAssets); ok {
		if oneOfNester.CollateralAssets != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.CollateralAssets); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CollateralAssets", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_DelegationActive); ok {
		if oneOfNester.DelegationActive != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DelegationActive); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelegationActive", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_DelegationPending); ok {
		if oneOfNester.DelegationPending != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DelegationPending); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelegationPending", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_DelegationAuto); ok {
		if oneOfNester.DelegationAuto != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DelegationAuto); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelegationAuto", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_GovernanceActive); ok {
		if oneOfNester.GovernanceActive != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.GovernanceActive); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GovernanceActive", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_GovernanceEnacted); ok {
		if oneOfNester.GovernanceEnacted != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.GovernanceEnacted); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GovernanceEnacted", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_StakingAccounts); ok {
		if oneOfNester.StakingAccounts != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StakingAccounts); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StakingAccounts", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_MatchingBook); ok {
		if oneOfNester.MatchingBook != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MatchingBook); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MatchingBook", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_NetworkParameters); ok {
		if oneOfNester.NetworkParameters != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NetworkParameters); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NetworkParameters", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_ExecutionMarkets); ok {
		if oneOfNester.ExecutionMarkets != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ExecutionMarkets); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ExecutionMarkets", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_MarketPositions); ok {
		if oneOfNester.MarketPositions != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketPositions); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketPositions", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_AppState); ok {
		if oneOfNester.AppState != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.AppState); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AppState", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_Epoch); ok {
		if oneOfNester.Epoch != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Epoch); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Epoch", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_RewardsPendingPayouts); ok {
		if oneOfNester.RewardsPendingPayouts != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RewardsPendingPayouts); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RewardsPendingPayouts", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_GovernanceNode); ok {
		if oneOfNester.GovernanceNode != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.GovernanceNode); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("GovernanceNode", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_LimitState); ok {
		if oneOfNester.LimitState != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LimitState); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LimitState", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_VoteSpamPolicy); ok {
		if oneOfNester.VoteSpamPolicy != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.VoteSpamPolicy); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("VoteSpamPolicy", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_SimpleSpamPolicy); ok {
		if oneOfNester.SimpleSpamPolicy != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SimpleSpamPolicy); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SimpleSpamPolicy", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_Notary); ok {
		if oneOfNester.Notary != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Notary); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Notary", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_ReplayProtection); ok {
		if oneOfNester.ReplayProtection != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ReplayProtection); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ReplayProtection", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_FutureState); ok {
		if oneOfNester.FutureState != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.FutureState); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FutureState", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_EventForwarder); ok {
		if oneOfNester.EventForwarder != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.EventForwarder); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EventForwarder", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_StakeVerifierDeposited); ok {
		if oneOfNester.StakeVerifierDeposited != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StakeVerifierDeposited); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StakeVerifierDeposited", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_StakeVerifierRemoved); ok {
		if oneOfNester.StakeVerifierRemoved != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StakeVerifierRemoved); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StakeVerifierRemoved", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_Witness); ok {
		if oneOfNester.Witness != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Witness); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Witness", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_DelegationLastReconciliationTime); ok {
		if oneOfNester.DelegationLastReconciliationTime != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DelegationLastReconciliationTime); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelegationLastReconciliationTime", err)
			}
		}
	}
	if oneOfNester, ok := this.GetData().(*Payload_Topology); ok {
		if oneOfNester.Topology != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Topology); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Topology", err)
			}
		}
	}
	return nil
}
func (this *Witness) Validate() error {
	for _, item := range this.Resources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
			}
		}
	}
	return nil
}
func (this *Resource) Validate() error {
	return nil
}
func (this *EventForwarder) Validate() error {
	for _, item := range this.AckedEvents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AckedEvents", err)
			}
		}
	}
	return nil
}
func (this *CollateralAccounts) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *CollateralAssets) Validate() error {
	for _, item := range this.Assets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Assets", err)
			}
		}
	}
	return nil
}
func (this *ActiveAssets) Validate() error {
	for _, item := range this.Assets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Assets", err)
			}
		}
	}
	return nil
}
func (this *PendingAssets) Validate() error {
	for _, item := range this.Assets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Assets", err)
			}
		}
	}
	return nil
}
func (this *Withdrawal) Validate() error {
	if this.Withdrawal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Withdrawal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Withdrawal", err)
		}
	}
	return nil
}
func (this *Deposit) Validate() error {
	if this.Deposit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deposit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deposit", err)
		}
	}
	return nil
}
func (this *TxRef) Validate() error {
	return nil
}
func (this *AssetAction) Validate() error {
	if this.BuiltinDeposit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BuiltinDeposit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BuiltinDeposit", err)
		}
	}
	if this.Erc20Deposit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Erc20Deposit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Erc20Deposit", err)
		}
	}
	if this.AssetList != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssetList); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssetList", err)
		}
	}
	return nil
}
func (this *BankingWithdrawals) Validate() error {
	for _, item := range this.Withdrawals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Withdrawals", err)
			}
		}
	}
	return nil
}
func (this *BankingDeposits) Validate() error {
	for _, item := range this.Deposit {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deposit", err)
			}
		}
	}
	return nil
}
func (this *BankingSeen) Validate() error {
	for _, item := range this.Refs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Refs", err)
			}
		}
	}
	return nil
}
func (this *BankingAssetActions) Validate() error {
	for _, item := range this.AssetAction {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AssetAction", err)
			}
		}
	}
	return nil
}
func (this *Checkpoint) Validate() error {
	return nil
}
func (this *DelegationLastReconciliationTime) Validate() error {
	return nil
}
func (this *DelegationActive) Validate() error {
	for _, item := range this.Delegations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Delegations", err)
			}
		}
	}
	return nil
}
func (this *DelegationPending) Validate() error {
	for _, item := range this.Delegations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Delegations", err)
			}
		}
	}
	for _, item := range this.Undelegation {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Undelegation", err)
			}
		}
	}
	return nil
}
func (this *DelegationAuto) Validate() error {
	return nil
}
func (this *PendingProposal) Validate() error {
	if this.Proposal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proposal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
		}
	}
	for _, item := range this.Yes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Yes", err)
			}
		}
	}
	for _, item := range this.No {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("No", err)
			}
		}
	}
	for _, item := range this.Invalid {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Invalid", err)
			}
		}
	}
	return nil
}
func (this *GovernanceEnacted) Validate() error {
	for _, item := range this.Proposals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposals", err)
			}
		}
	}
	return nil
}
func (this *GovernanceActive) Validate() error {
	for _, item := range this.Proposals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposals", err)
			}
		}
	}
	return nil
}
func (this *GovernanceNode) Validate() error {
	for _, item := range this.Proposals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposals", err)
			}
		}
	}
	return nil
}
func (this *StakingAccount) Validate() error {
	for _, item := range this.Events {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Events", err)
			}
		}
	}
	return nil
}
func (this *StakingAccounts) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *MatchingBook) Validate() error {
	for _, item := range this.Buy {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Buy", err)
			}
		}
	}
	for _, item := range this.Sell {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sell", err)
			}
		}
	}
	return nil
}
func (this *NetParams) Validate() error {
	for _, item := range this.Params {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
			}
		}
	}
	return nil
}
func (this *DecimalMap) Validate() error {
	return nil
}
func (this *TimePrice) Validate() error {
	return nil
}
func (this *PriceVolume) Validate() error {
	return nil
}
func (this *PriceRange) Validate() error {
	return nil
}
func (this *PriceBound) Validate() error {
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *PriceRangeCache) Validate() error {
	if this.Bound != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Bound); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Bound", err)
		}
	}
	if this.Range != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Range); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Range", err)
		}
	}
	return nil
}
func (this *CurrentPrice) Validate() error {
	return nil
}
func (this *PastPrice) Validate() error {
	return nil
}
func (this *PriceMonitor) Validate() error {
	for _, item := range this.FpHorizons {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("FpHorizons", err)
			}
		}
	}
	for _, item := range this.Bounds {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Bounds", err)
			}
		}
	}
	for _, item := range this.PriceRangeCache {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PriceRangeCache", err)
			}
		}
	}
	for _, item := range this.PricesNow {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PricesNow", err)
			}
		}
	}
	for _, item := range this.PricesPast {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PricesPast", err)
			}
		}
	}
	for _, item := range this.RefPriceCache {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RefPriceCache", err)
			}
		}
	}
	return nil
}
func (this *AuctionState) Validate() error {
	if this.End != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.End); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("End", err)
		}
	}
	return nil
}
func (this *EquityShareLP) Validate() error {
	return nil
}
func (this *EquityShare) Validate() error {
	for _, item := range this.Lps {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Lps", err)
			}
		}
	}
	return nil
}
func (this *Market) Validate() error {
	if this.Market != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Market); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Market", err)
		}
	}
	if this.PriceMonitor != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PriceMonitor); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PriceMonitor", err)
		}
	}
	if this.AuctionState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AuctionState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AuctionState", err)
		}
	}
	for _, item := range this.PeggedOrders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PeggedOrders", err)
			}
		}
	}
	for _, item := range this.ExpiringOrders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ExpiringOrders", err)
			}
		}
	}
	if this.EquityShare != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EquityShare); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EquityShare", err)
		}
	}
	return nil
}
func (this *ExecutionMarkets) Validate() error {
	for _, item := range this.Markets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Markets", err)
			}
		}
	}
	return nil
}
func (this *Position) Validate() error {
	return nil
}
func (this *MarketPositions) Validate() error {
	for _, item := range this.Positions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Positions", err)
			}
		}
	}
	return nil
}
func (this *AppState) Validate() error {
	return nil
}
func (this *EpochState) Validate() error {
	return nil
}
func (this *RewardsPendingPayouts) Validate() error {
	for _, item := range this.ScheduledRewardsPayout {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ScheduledRewardsPayout", err)
			}
		}
	}
	return nil
}
func (this *ScheduledRewardsPayout) Validate() error {
	for _, item := range this.RewardsPayout {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RewardsPayout", err)
			}
		}
	}
	return nil
}
func (this *RewardsPayout) Validate() error {
	for _, item := range this.RewardPartyAmount {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RewardPartyAmount", err)
			}
		}
	}
	return nil
}
func (this *RewardsPartyAmount) Validate() error {
	return nil
}
func (this *LimitState) Validate() error {
	return nil
}
func (this *VoteSpamPolicy) Validate() error {
	for _, item := range this.PartyToVote {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PartyToVote", err)
			}
		}
	}
	for _, item := range this.BannedParties {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BannedParties", err)
			}
		}
	}
	for _, item := range this.TokenBalance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TokenBalance", err)
			}
		}
	}
	for _, item := range this.RecentBlocksRejectStats {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RecentBlocksRejectStats", err)
			}
		}
	}
	return nil
}
func (this *PartyProposalVoteCount) Validate() error {
	return nil
}
func (this *BannedParty) Validate() error {
	return nil
}
func (this *PartyTokenBalance) Validate() error {
	return nil
}
func (this *BlockRejectStats) Validate() error {
	return nil
}
func (this *SpamPartyTransactionCount) Validate() error {
	return nil
}
func (this *SimpleSpamPolicy) Validate() error {
	for _, item := range this.PartyToCount {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PartyToCount", err)
			}
		}
	}
	for _, item := range this.BannedParties {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("BannedParties", err)
			}
		}
	}
	for _, item := range this.TokenBalance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TokenBalance", err)
			}
		}
	}
	return nil
}
func (this *NotarySigs) Validate() error {
	return nil
}
func (this *Notary) Validate() error {
	for _, item := range this.NotarySigs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NotarySigs", err)
			}
		}
	}
	return nil
}
func (this *ReplayProtection) Validate() error {
	for _, item := range this.RecentBlocksTransactions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RecentBlocksTransactions", err)
			}
		}
	}
	return nil
}
func (this *RecentBlocksTransactions) Validate() error {
	return nil
}
func (this *FutureState) Validate() error {
	return nil
}
func (this *StakeVerifierDeposited) Validate() error {
	for _, item := range this.PendingDeposited {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PendingDeposited", err)
			}
		}
	}
	return nil
}
func (this *StakeVerifierRemoved) Validate() error {
	for _, item := range this.PendingRemoved {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PendingRemoved", err)
			}
		}
	}
	return nil
}
func (this *StakeVerifierPending) Validate() error {
	return nil
}
func (this *Topology) Validate() error {
	for _, item := range this.ValidatorData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ValidatorData", err)
			}
		}
	}
	return nil
}
