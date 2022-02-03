package commands

import (
	"encoding/hex"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
)

func CheckNodeRegistration(cmd *commandspb.NodeRegistration) error {
	return checkNodeRegistration(cmd).ErrorOrNil()
}

func checkNodeRegistration(cmd *commandspb.NodeRegistration) Errors {
	errs := NewErrors()

	if cmd == nil {
		return errs.FinalAddForProperty("node_registration", ErrIsRequired)
	}

	if len(cmd.VegaPubKey) == 0 {
		errs.AddForProperty("node_registration.vega_pub_key", ErrIsRequired)
	} else {
		_, err := hex.DecodeString(cmd.VegaPubKey)
		if err != nil {
			errs.AddForProperty("node_registration.vega_pub_key", ErrShouldBeHexEncoded)
		}
	}

	if len(cmd.EthereumAddress) == 0 {
		errs.AddForProperty("node_registration.ethereum_address", ErrIsRequired)
	}

	if len(cmd.ChainPubKey) == 0 {
		errs.AddForProperty("node_registration.chain_pub_key", ErrIsRequired)
	}

	return errs
}
