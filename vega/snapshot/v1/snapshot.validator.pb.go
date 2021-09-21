// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vega/snapshot/v1/snapshot.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/protos/vega"
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
func (this *Checkpoint) Validate() error {
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
