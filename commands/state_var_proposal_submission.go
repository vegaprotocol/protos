package commands

import commandspb "code.vegaprotocol.io/protos/vega/commands/v1"

func CheckStateVariableProposal(cmd *commandspb.StateVariableProposal) error {
	return checkStateVariableProposal(cmd).ErrorOrNil()
}

func checkStateVariableProposal(cmd *commandspb.StateVariableProposal) Errors {
	errs := NewErrors()
	if cmd == nil {
		return errs.FinalAddForProperty("state_variable_proposal", ErrIsRequired)
	}

	if cmd.Proposal == nil {
		return errs.FinalAddForProperty("state_variable_proposal.missing_proposal", ErrIsRequired)
	}

	if len(cmd.Proposal.EventId) == 0 {
		errs.AddForProperty("state_variable_proposal.event_id", ErrIsRequired)
	}
	if len(cmd.Proposal.Kvb) == 0 {
		errs.AddForProperty("state_variable_proposal.key_value_bundle", ErrIsRequired)
	}

	return errs
}
