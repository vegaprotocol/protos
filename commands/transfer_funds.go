package commands

import (
	"math/big"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
)

func CheckTransfer(cmd *commandspb.Transfer) error {
	return checkTransfer(cmd).ErrorOrNil()
}

func checkTransfer(cmd *commandspb.Transfer) Errors {
	errs := NewErrors()

	if cmd == nil {
		return errs.FinalAddForProperty("transfer_funds", ErrIsRequired)
	}

	if len(cmd.Amount) <= 0 {
		errs.AddForProperty("transfer_funds.amount", ErrIsRequired)
	} else {
		if amount, ok := big.NewInt(0).SetString(cmd.Amount, 10); !ok {
			errs.AddForProperty("transfer_funds.amount", ErrNotAValidInteger)
		} else {
			if amount.Cmp(big.NewInt(0)) == 0 {
				errs.AddForProperty("transfer_funds.amount", ErrIsRequired)
			}
			if amount.Cmp(big.NewInt(0)) == -1 {
				errs.AddForProperty("transfer_funds.amount", ErrMustBePositive)
			}
		}
	}

	if len(cmd.To) <= 0 {
		errs.AddForProperty("transfer_funds.to", ErrIsRequired)
	}

	if len(cmd.Asset) <= 0 {
		errs.AddForProperty("transfer_funds.asset", ErrIsRequired)
	}

	// arbitrary 100 char length for now
	if len(cmd.Reference) > 100 {
		errs.AddForProperty("transfer_funds.reference", ErrMustBeLessThan100Chars)
	}

	return errs
}
