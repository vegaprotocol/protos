package commands_test

import (
	"testing"

	"code.vegaprotocol.io/protos/commands"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"

	"github.com/stretchr/testify/assert"
)

/**********************************************************************************/
/*                                   DELEGATION                                   */
/**********************************************************************************/
func TestSubmittingNoDelegateCommandFails(t *testing.T) {
	err := checkDelegateSubmission(nil)

	assert.Contains(t, err.Get("delegate_submission"), commands.ErrIsRequired)
}

func TestSubmittingNoDelegateNodeIdFails(t *testing.T) {
	cmd := &commandspb.DelegateSubmission{
		Amount: "1000",
	}
	err := checkDelegateSubmission(cmd)

	assert.Contains(t, err.Get("delegate_submission.node_id"), commands.ErrIsRequired)
}

func TestSubmittingNoDelegateAmountFails(t *testing.T) {
	cmd := &commandspb.DelegateSubmission{
		NodeId: "TestingNodeID",
	}
	err := checkDelegateSubmission(cmd)

	assert.Contains(t, err.Get("delegate_submission.amount"), commands.ErrIsRequired)
}

func checkDelegateSubmission(cmd *commandspb.DelegateSubmission) commands.Errors {
	err := commands.CheckDelegateSubmission(cmd)

	e, ok := err.(commands.Errors)
	if !ok {
		return commands.NewErrors()
	}
	return e
}

/**********************************************************************************/
/*                                  UNDELEGATION                                  */
/**********************************************************************************/
func TestSubmittingNoUndelegateCommandFails(t *testing.T) {
	err := checkUndelegateSubmission(nil)

	assert.Contains(t, err.Get("undelegate_submission"), commands.ErrIsRequired)
}

func TestSubmittingNoUndelegateNodeIdFails(t *testing.T) {
	cmd := &commandspb.UndelegateSubmission{
		Amount: "1000",
	}
	err := checkUndelegateSubmission(cmd)

	assert.Contains(t, err.Get("undelegate_submission.node_id"), commands.ErrIsRequired)
}

func TestSubmittingInvalidUndelegateMethod(t *testing.T) {
	invalidMethod := len(commandspb.UndelegateSubmission_Method_value)
	cmd := &commandspb.UndelegateSubmission{
		NodeId: "TestingNodeID",
		Method: commandspb.UndelegateSubmission_Method(invalidMethod),
	}
	err := checkUndelegateSubmission(cmd)

	assert.Contains(t, err.Get("undelegate_submission.method"), commands.ErrIsRequired)

	cmd = &commandspb.UndelegateSubmission{
		NodeId: "TestingNodeID",
	}
	err = checkUndelegateSubmission(cmd)

	assert.Contains(t, err.Get("undelegate_submission.method"), commands.ErrIsRequired)
}

func TestSubmittingNoUndelegateAmountSucceeds(t *testing.T) {
	cmd := &commandspb.UndelegateSubmission{
		NodeId: "TestingNodeID",
		Method: 1,
	}
	err := checkUndelegateSubmission(cmd)

	assert.Equal(t, 0, len(err))
}

func checkUndelegateSubmission(cmd *commandspb.UndelegateSubmission) commands.Errors {
	err := commands.CheckUndelegateSubmission(cmd)

	e, ok := err.(commands.Errors)
	if !ok {
		return commands.NewErrors()
	}
	return e
}
