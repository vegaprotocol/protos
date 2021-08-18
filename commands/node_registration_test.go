package commands_test

import (
	"encoding/hex"
	"testing"

	"code.vegaprotocol.io/protos/commands"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	"github.com/stretchr/testify/assert"
)

func TestCheckNodeRegistration(t *testing.T) {
	t.Run("Submitting a nil command fails", testNilNodeRegistrationFails)
	t.Run("Submitting a node registration without vega pub key fails", testNodeRegistrationWithoutVegaPubKeyFails)
	t.Run("Submitting a node registration with valid vega pub key succeeds", testNodeRegistrationWithValidVegaPubKeySucceeds)
	t.Run("Submitting a node registration with invalid encoding of vega pub key succeeds", testNodeRegistrationWithInvalidEncodingOfVegaPubKeyFails)
	t.Run("Submitting a node registration without ethereum pub key fails", testNodeRegistrationWithoutEthereumAddressFails)
	t.Run("Submitting a node registration with ethereum address succeeds", testNodeRegistrationWithEthereumAddressSucceeds)
	t.Run("Submitting a node registration without chain address fails", testNodeRegistrationWithoutChainPubKeyFails)
	t.Run("Submitting a node registration with chain pub key succeeds", testNodeRegistrationWithChainPubKeySucceeds)
}

func testNilNodeRegistrationFails(t *testing.T) {
	err := checkNodeRegistration(nil)

	assert.Error(t, err)
}

func testNodeRegistrationWithoutVegaPubKeyFails(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{})
	assert.Contains(t, err.Get("node_registration.vega_pub_key"), commands.ErrIsRequired)
}

func testNodeRegistrationWithValidVegaPubKeySucceeds(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{
		VegaPubKey: hex.EncodeToString([]byte("0xDEADBEEF")),
	})
	assert.NotContains(t, err.Get("node_registration.vega_pub_key"), commands.ErrIsRequired)
	assert.NotContains(t, err.Get("node_registration.vega_pub_key"), commands.ErrShouldBeHexEncoded)
}

func testNodeRegistrationWithInvalidEncodingOfVegaPubKeyFails(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{
		VegaPubKey: "invalid-hex-encoding",
	})
	assert.Contains(t, err.Get("node_registration.vega_pub_key"), commands.ErrShouldBeHexEncoded)
}

func testNodeRegistrationWithoutEthereumAddressFails(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{})
	assert.Contains(t, err.Get("node_registration.ethereum_address"), commands.ErrIsRequired)
}

func testNodeRegistrationWithEthereumAddressSucceeds(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{
		EthereumAddress: "0xDEADBEEF",
	})
	assert.NotContains(t, err.Get("node_registration.ethereum_address"), commands.ErrIsRequired)
}

func testNodeRegistrationWithoutChainPubKeyFails(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{})
	assert.Contains(t, err.Get("node_registration.chain_pub_key"), commands.ErrIsRequired)
}

func testNodeRegistrationWithChainPubKeySucceeds(t *testing.T) {
	err := checkNodeRegistration(&commandspb.NodeRegistration{
		ChainPubKey: "0xDEADBEEF",
	})
	assert.NotContains(t, err.Get("node_registration.chain_pub_key"), commands.ErrIsRequired)
}

func checkNodeRegistration(cmd *commandspb.NodeRegistration) commands.Errors {
	err := commands.CheckNodeRegistration(cmd)

	e, ok := err.(commands.Errors)
	if !ok {
		return commands.NewErrors()
	}

	return e
}
