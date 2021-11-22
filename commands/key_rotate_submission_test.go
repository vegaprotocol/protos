package commands_test

import (
	"testing"
	"time"

	"code.vegaprotocol.io/protos/commands"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"

	"github.com/stretchr/testify/assert"
)

func TestSubmittingNoKeyRotateSubmissionCommandFails(t *testing.T) {
	err := checkKeyRotateSubmission(nil)

	assert.Contains(t, err.Get("key_rotate_submission"), commands.ErrIsRequired)
}

func TestKeyRotateSubmissionSubmittingEmptyCommandFails(t *testing.T) {
	err := checkKeyRotateSubmission(&commandspb.KeyRotateSubmission{})

	assert.Contains(t, err.Get("key_rotate_submission.new_pub_key_hash"), commands.ErrIsRequired)
	assert.Contains(t, err.Get("key_rotate_submission.key_number"), commands.ErrIsRequired)
	assert.Contains(t, err.Get("key_rotate_submission.time"), commands.ErrNotAValidInteger)
}

func TestKeyRotateSubmissionMissingNewPubKeyHashFails(t *testing.T) {
	err := checkKeyRotateSubmission(&commandspb.KeyRotateSubmission{
		KeyNumber:   10,
		TargetBlock: 100,
		Time:        time.Now().UnixNano(),
	})

	assert.Contains(t, err.Get("key_rotate_submission.new_pub_key_hash"), commands.ErrIsRequired)
}

func TestKeyRotateSubmissionMissingKeyNumberFails(t *testing.T) {
	err := checkKeyRotateSubmission(&commandspb.KeyRotateSubmission{
		NewPubKeyHash: "123456789abcdef",
		TargetBlock:   100,
		Time:          time.Now().UnixNano(),
	})

	assert.Contains(t, err.Get("key_rotate_submission.key_number"), commands.ErrIsRequired)
}

func TestKeyRotateSubmissionMissingTimeFails(t *testing.T) {
	err := checkKeyRotateSubmission(&commandspb.KeyRotateSubmission{
		KeyNumber:     10,
		NewPubKeyHash: "123456789abcdef",
		TargetBlock:   100,
		Time:          -1,
	})

	assert.Contains(t, err.Get("key_rotate_submission.time"), commands.ErrNotAValidInteger)
}

func TestSubmittingEmptyCommandSuccess(t *testing.T) {
	err := checkKeyRotateSubmission(&commandspb.KeyRotateSubmission{
		KeyNumber:     10,
		NewPubKeyHash: "123456789abcdef",
		TargetBlock:   100,
		Time:          time.Now().UnixNano(),
	})

	assert.True(t, err.Empty())
}

func checkKeyRotateSubmission(cmd *commandspb.KeyRotateSubmission) commands.Errors {
	err := commands.CheckKeyRotateSubmission(cmd)

	e, ok := err.(commands.Errors)
	if !ok {
		return commands.NewErrors()
	}
	return e
}
