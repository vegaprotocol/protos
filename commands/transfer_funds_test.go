package commands_test

import (
	"testing"

	"code.vegaprotocol.io/protos/commands"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"github.com/stretchr/testify/assert"
)

func TestNilTransferFundsFails(t *testing.T) {
	err := checkTransfer(nil)

	assert.Contains(t, err.Get("transfer_funds"), commands.ErrIsRequired)
}

func TestTransferFunds(t *testing.T) {
	var cases = []struct {
		transfer  commandspb.Transfer
		errString string
	}{
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "1",
				Reference: "testing",
			},
		},
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "1",
				Reference: "",
			},
		},
		{
			transfer: commandspb.Transfer{
				To:        "",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "1",
				Reference: "testing",
			},
			errString: "transfer_funds.to (is required)",
		},
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "",
				Amount:    "1",
				Reference: "testing",
			},
			errString: "transfer_funds.asset (is required)",
		},
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "",
				Reference: "testing",
			},
			errString: "transfer_funds.amount (is required)",
		},
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "-1",
				Reference: "testing",
			},
			errString: "transfer_funds.amount (must be positive)",
		},
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "0",
				Reference: "testing",
			},
			errString: "transfer_funds.amount (is required)",
		},
		{
			transfer: commandspb.Transfer{
				To:        "84e2b15102a8d6c1c6b4bdf40af8a0dc21b040eaaa1c94cd10d17604b75fdc35",
				Asset:     "080538b7cc2249de568cb4272a17f4d5e0b0a69a1a240acbf5119d816178daff",
				Amount:    "1",
				Reference: "testingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtest",
			},
			errString: "transfer_funds.reference (must be less than 100 characters)",
		},
		{
			transfer: commandspb.Transfer{
				To:        "",
				Asset:     "",
				Amount:    "",
				Reference: "",
			},
			errString: "transfer_funds.amount (is required), transfer_funds.asset (is required), transfer_funds.to (is required)",
		},
	}

	for _, c := range cases {
		err := commands.CheckTransfer(&c.transfer)
		if len(c.errString) <= 0 {
			assert.NoError(t, err)
			continue
		}
		assert.Error(t, err)
		assert.EqualError(t, err, c.errString)
	}

}

func checkTransfer(cmd *commandspb.Transfer) commands.Errors {
	err := commands.CheckTransfer(cmd)

	e, ok := err.(commands.Errors)
	if !ok {
		return commands.NewErrors()
	}

	return e
}
