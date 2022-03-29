package commands

import (
	"errors"
	"math/big"

	"code.vegaprotocol.io/protos/vega"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
)

var (
	ErrMustBeAfterStartEpoch = errors.New("must be after start_epoch")
)

func CheckTransfer(cmd *commandspb.Transfer) error {
	return checkTransfer(cmd).ErrorOrNil()
}

func checkTransfer(cmd *commandspb.Transfer) Errors {
	errs := NewErrors()

	if cmd == nil {
		return errs.FinalAddForProperty("transfer", ErrIsRequired)
	}

	if len(cmd.Amount) <= 0 {
		errs.AddForProperty("transfer.amount", ErrIsRequired)
	} else {
		if amount, ok := big.NewInt(0).SetString(cmd.Amount, 10); !ok {
			errs.AddForProperty("transfer.amount", ErrNotAValidInteger)
		} else {
			if amount.Cmp(big.NewInt(0)) == 0 {
				errs.AddForProperty("transfer.amount", ErrIsRequired)
			}
			if amount.Cmp(big.NewInt(0)) == -1 {
				errs.AddForProperty("transfer.amount", ErrMustBePositive)
			}
		}
	}

	if len(cmd.To) <= 0 {
		errs.AddForProperty("transfer.to", ErrIsRequired)
	}

	if cmd.ToAccountType == vega.AccountType_ACCOUNT_TYPE_UNSPECIFIED {
		errs.AddForProperty("transfer.to_account_type", ErrIsNotValid)
	} else if _, ok := vega.AccountType_name[int32(cmd.ToAccountType)]; !ok {
		errs.AddForProperty("transfer.to_account_type", ErrIsNotValid)
	}

	if cmd.FromAccountType == vega.AccountType_ACCOUNT_TYPE_UNSPECIFIED {
		errs.AddForProperty("transfer.from_account_type", ErrIsNotValid)
	} else if _, ok := vega.AccountType_name[int32(cmd.FromAccountType)]; !ok {
		errs.AddForProperty("transfer.from_account_type", ErrIsNotValid)
	}

	if len(cmd.Asset) <= 0 {
		errs.AddForProperty("transfer.asset", ErrIsRequired)
	}

	// arbitrary 100 char length for now
	if len(cmd.Reference) > 100 {
		errs.AddForProperty("transfer.reference", ErrMustBeLessThan100Chars)
	}

	if cmd.Kind == nil {
		errs.AddForProperty("transfer.kind", ErrIsRequired)
	} else {
		switch k := cmd.Kind.(type) {
		case *commandspb.Transfer_OneOff:
			if k.OneOff.GetDeliverOn() < 0 {
				errs.AddForProperty("transfer.kind.deliver_on", ErrMustBePositiveOrZero)
			}
		case *commandspb.Transfer_Recurring:
			if k.Recurring.EndEpoch != nil && *k.Recurring.EndEpoch <= 0 {
				errs.AddForProperty("transfer.kind.end_epoch", ErrMustBePositive)
			}
			if k.Recurring.StartEpoch == 0 {
				errs.AddForProperty("transfer.kind.start_epoch", ErrMustBePositive)
			}
			if k.Recurring.EndEpoch != nil && k.Recurring.StartEpoch > *k.Recurring.EndEpoch {
				errs.AddForProperty("transfer.kind.end_epoch", ErrMustBeAfterStartEpoch)
			}
			if f, ok := big.NewFloat(0).SetString(k.Recurring.Factor); !ok {
				errs.AddForProperty("transfer.kind.factor", ErrNotAValidFloat)
			} else {
				if f.Cmp(big.NewFloat(0)) <= 0 {
					errs.AddForProperty("transfer.kind.factor", ErrMustBePositive)
				}
			}
		default:
			errs.AddForProperty("transfer.kind", ErrIsNotSupported)
		}
	}

	return errs
}
