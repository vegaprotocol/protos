package commands

import commandspb "code.vegaprotocol.io/protos/vega/commands/v1"

func CheckKeyRotateSubmission(cmd *commandspb.KeyRotateSubmission) error {
	return checkKeyRotateSubmission(cmd).ErrorOrNil()
}

func checkKeyRotateSubmission(cmd *commandspb.KeyRotateSubmission) Errors {
	errs := NewErrors()

	if cmd == nil {
		return errs.FinalAddForProperty("key_rotate_submission", ErrIsRequired)
	}

	if len(cmd.NewPubKey) <= 0 {
		errs.AddForProperty("key_rotate_submission.new_pub_key", ErrIsRequired)
	}

	if cmd.KeyNumber == 0 {
		errs.AddForProperty("key_rotate_submission.key_number", ErrIsRequired)
	}

	return errs
}
