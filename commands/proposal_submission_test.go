package commands_test

import (
	"errors"
	"testing"

	"code.vegaprotocol.io/protos/commands"
	types "code.vegaprotocol.io/protos/vega"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"github.com/stretchr/testify/assert"
)

func TestCheckProposalSubmission(t *testing.T) {
	t.Run("Submitting a nil command fails", testNilProposalSubmissionFails)
	t.Run("Submitting a proposal without terms fails", testProposalSubmissionWithoutTermsFails)
	t.Run("Submitting a proposal with non-positive closing timestamp fails", testProposalSubmissionWithNonPositiveClosingTimestampFails)
	t.Run("Submitting a proposal with positive closing timestamp succeeds", testProposalSubmissionWithPositiveClosingTimestampSucceeds)
	t.Run("Submitting a proposal with non-positive enactment timestamp fails", testProposalSubmissionWithNonPositiveEnactmentTimestampFails)
	t.Run("Submitting a proposal with positive enactment timestamp succeeds", testProposalSubmissionWithPositiveEnactmentTimestampSucceeds)
	t.Run("Submitting a proposal with negative validation timestamp fails", testProposalSubmissionWithNegativeValidationTimestampFails)
	t.Run("Submitting a proposal with positive validation timestamp succeeds", testProposalSubmissionWithPositiveValidationTimestampSucceeds)
	t.Run("Submitting a proposal with closing timestamp after enactment timestamp fails", testProposalSubmissionWithClosingTimestampAfterEnactmentTimestampFails)
	t.Run("Submitting a proposal with closing timestamp before enactment timestamp succeeds", testProposalSubmissionWithClosingTimestampBeforeEnactmentTimestampSucceeds)
	t.Run("Submitting a proposal with closing timestamp at enactment timestamp succeeds", testProposalSubmissionWithClosingTimestampAtEnactmentTimestampSucceeds)
	t.Run("Submitting a proposal with validation timestamp after closing timestamp fails", testProposalSubmissionWithValidationTimestampAfterClosingTimestampFails)
	t.Run("Submitting a proposal with validation timestamp at closing timestamp succeeds", testProposalSubmissionWithValidationTimestampAtClosingTimestampFails)
	t.Run("Submitting a proposal with validation timestamp before closing timestamp fails", testProposalSubmissionWithValidationTimestampBeforeClosingTimestampSucceeds)
}

func testProposalSubmissionWithoutTermsFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{})

	assert.Contains(t, err.Get("proposal_submission.terms"), commands.ErrIsRequired)
}

func testNilProposalSubmissionFails(t *testing.T) {
	err := checkProposalSubmission(nil)

	assert.Contains(t, err.Get("proposal_submission"), commands.ErrIsRequired)
}

func testProposalSubmissionWithNonPositiveClosingTimestampFails(t *testing.T) {
	testCases := []struct {
		msg   string
		value int64
	}{
		{
			msg:   "with 0 as closing timestamp",
			value: 0,
		}, {
			msg:   "with negative closing timestamp",
			value: RandomNegativeI64(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			err := checkProposalSubmission(&commandspb.ProposalSubmission{
				Terms: &types.ProposalTerms{
					ClosingTimestamp: tc.value,
				},
			})

			assert.Contains(t, err.Get("proposal_submission.terms.closing_timestamp"), commands.ErrMustBePositive)
		})
	}
}

func testProposalSubmissionWithPositiveClosingTimestampSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp: RandomPositiveI64(),
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.closing_timestamp"), commands.ErrMustBePositive)
}

func testProposalSubmissionWithNonPositiveEnactmentTimestampFails(t *testing.T) {
	testCases := []struct {
		msg   string
		value int64
	}{
		{
			msg:   "with 0 as closing timestamp",
			value: 0,
		}, {
			msg:   "with negative closing timestamp",
			value: RandomNegativeI64(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			err := checkProposalSubmission(&commandspb.ProposalSubmission{
				Terms: &types.ProposalTerms{
					EnactmentTimestamp: tc.value,
				},
			})

			assert.Contains(t, err.Get("proposal_submission.terms.enactment_timestamp"), commands.ErrMustBePositive)
		})
	}
}

func testProposalSubmissionWithPositiveEnactmentTimestampSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			EnactmentTimestamp: RandomPositiveI64(),
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.enactment_timestamp"), commands.ErrMustBePositive)
}

func testProposalSubmissionWithNegativeValidationTimestampFails(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ValidationTimestamp: RandomNegativeI64(),
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.validation_timestamp"), commands.ErrMustBePositiveOrZero)
}

func testProposalSubmissionWithPositiveValidationTimestampSucceeds(t *testing.T) {
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ValidationTimestamp: RandomPositiveI64(),
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.validation_timestamp"), commands.ErrIsRequired)
}

func testProposalSubmissionWithClosingTimestampAfterEnactmentTimestampFails(t *testing.T) {
	closingTime := RandomPositiveI64()
	enactmentTime := RandomPositiveI64Before(closingTime)
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp:   closingTime,
			EnactmentTimestamp: enactmentTime,
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.closing_timestamp"),
		errors.New("cannot be after enactment time"),
	)
}

func testProposalSubmissionWithClosingTimestampBeforeEnactmentTimestampSucceeds(t *testing.T) {
	enactmentTime := RandomPositiveI64()
	closingTime := RandomPositiveI64Before(enactmentTime)

	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp:   closingTime,
			EnactmentTimestamp: enactmentTime,
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.closing_timestamp"),
		errors.New("cannot be after enactment time"),
	)
}

func testProposalSubmissionWithClosingTimestampAtEnactmentTimestampSucceeds(t *testing.T) {
	enactmentTime := RandomPositiveI64()

	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp:   enactmentTime,
			EnactmentTimestamp: enactmentTime,
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.closing_timestamp"),
		errors.New("cannot be after enactment time"),
	)
}

func testProposalSubmissionWithValidationTimestampAfterClosingTimestampFails(t *testing.T) {
	validationTime := RandomPositiveI64()
	closingTime := RandomPositiveI64Before(validationTime)
	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp:    closingTime,
			ValidationTimestamp: validationTime,
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.validation_timestamp"),
		errors.New("cannot be after or equal to closing time"),
	)
}

func testProposalSubmissionWithValidationTimestampAtClosingTimestampFails(t *testing.T) {
	validationTime := RandomPositiveI64()

	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp:    validationTime,
			ValidationTimestamp: validationTime,
		},
	})

	assert.Contains(t, err.Get("proposal_submission.terms.validation_timestamp"),
		errors.New("cannot be after or equal to closing time"),
	)

}

func testProposalSubmissionWithValidationTimestampBeforeClosingTimestampSucceeds(t *testing.T) {
	closingTime := RandomPositiveI64()
	validationTime := RandomPositiveI64Before(closingTime)

	err := checkProposalSubmission(&commandspb.ProposalSubmission{
		Terms: &types.ProposalTerms{
			ClosingTimestamp:    closingTime,
			ValidationTimestamp: validationTime,
		},
	})

	assert.NotContains(t, err.Get("proposal_submission.terms.validation_timestamp"),
		errors.New("cannot be after or equal to closing time"),
	)
}

func checkProposalSubmission(cmd *commandspb.ProposalSubmission) commands.Errors {
	err := commands.CheckProposalSubmission(cmd)

	e, ok := err.(commands.Errors)
	if !ok {
		return commands.NewErrors()
	}

	return e
}
