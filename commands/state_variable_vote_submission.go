package commands

import commandspb "code.vegaprotocol.io/protos/vega/commands/v1"

func CheckStateVariableVote(cmd *commandspb.StateVariableConsensusVote) error {
	return checkStateVariableVote(cmd).ErrorOrNil()
}

func checkStateVariableVote(cmd *commandspb.StateVariableConsensusVote) Errors {
	errs := NewErrors()
	if cmd == nil {
		return errs.FinalAddForProperty("state_variable_vote", ErrIsRequired)
	}

	if cmd.Vote == nil {
		return errs.FinalAddForProperty("state_variable_vote.vote", ErrIsRequired)
	}

	if cmd.Vote.StateVarId == "" {
		errs.AddForProperty("state_variable_vote.id", ErrIsRequired)
	}

	if cmd.Vote.EventId <= 0 {
		errs.AddForProperty("state_variable_vote.event_id", ErrMustBePositive)
	}

	if cmd.Vote.Round < 0 {
		errs.AddForProperty("state_variable_vote.round", ErrMustBePositiveOrZero)
	}

	return errs
}
