package commands

import (
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
)

func CheckDelegateSubmission(cmd *commandspb.DelegateSubmission) error {
	return checkDelegateSubmission(cmd).ErrorOrNil()
}

func checkDelegateSubmission(cmd *commandspb.DelegateSubmission) Errors {
	errs := NewErrors()

	if cmd == nil {
		return errs.FinalAddForProperty("delegate_submission", ErrIsRequired)
	}

	if cmd.Amount <= 0 {
		errs.AddForProperty("delegate_submission.amount", ErrIsRequired)
	}

	if len(cmd.NodeId) <= 0 {
		errs.AddForProperty("delegate_submission.node_id", ErrIsRequired)
	}

	return errs
}

func CheckUndelegateSubmission(cmd *commandspb.UndelegateSubmission) error {
	return checkUndelegateSubmission(cmd).ErrorOrNil()
}

func checkUndelegateSubmission(cmd *commandspb.UndelegateSubmission) Errors {
	errs := NewErrors()

	if cmd == nil {
		return errs.FinalAddForProperty("undelegate_submission", ErrIsRequired)
	}

	if cmd.Method > commandspb.UndelegateSubmission_InAnger {
		errs.AddForProperty("undelegate_submission.method", ErrIsRequired)
	}

	if len(cmd.NodeId) <= 0 {
		errs.AddForProperty("undelegate_submission.node_id", ErrIsRequired)
	}

	return errs
}
