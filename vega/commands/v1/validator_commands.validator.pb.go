// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vega/commands/v1/validator_commands.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/protos/vega"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NodeRegistration) Validate() error {
	if this.VegaPubKey == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("VegaPubKey", fmt.Errorf(`value '%v' must not be an empty string`, this.VegaPubKey))
	}
	if this.EthereumAddress == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EthereumAddress", fmt.Errorf(`value '%v' must not be an empty string`, this.EthereumAddress))
	}
	if this.ChainPubKey == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ChainPubKey", fmt.Errorf(`value '%v' must not be an empty string`, this.ChainPubKey))
	}
	if this.InfoUrl == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InfoUrl", fmt.Errorf(`value '%v' must not be an empty string`, this.InfoUrl))
	}
	if this.Country == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Country", fmt.Errorf(`value '%v' must not be an empty string`, this.Country))
	}
	return nil
}
func (this *NodeVote) Validate() error {
	if this.Reference == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Reference", fmt.Errorf(`value '%v' must not be an empty string`, this.Reference))
	}
	return nil
}
func (this *NodeSignature) Validate() error {
	return nil
}
func (this *ChainEvent) Validate() error {
	if oneOfNester, ok := this.GetEvent().(*ChainEvent_Builtin); ok {
		if oneOfNester.Builtin != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Builtin); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Builtin", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*ChainEvent_Erc20); ok {
		if oneOfNester.Erc20 != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Erc20); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Erc20", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*ChainEvent_Btc); ok {
		if oneOfNester.Btc != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Btc); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Btc", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*ChainEvent_Validator); ok {
		if oneOfNester.Validator != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Validator); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Validator", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*ChainEvent_StakingEvent); ok {
		if oneOfNester.StakingEvent != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StakingEvent); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StakingEvent", err)
			}
		}
	}
	return nil
}
func (this *KeyRotateSubmission) Validate() error {
	return nil
}
func (this *StateVariableConsensusVote) Validate() error {
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *StateVariableProposal) Validate() error {
	if this.Proposal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proposal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
		}
	}
	return nil
}
