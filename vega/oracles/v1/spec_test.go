package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeterminisitcOracleID(t *testing.T) {
	filter := &Filter{
		Key: &PropertyKey{
			Name: "trading.terminated",
			Type: PropertyKey_TYPE_BOOLEAN,
		},
		Conditions: []*Condition{
			{
				Operator: Condition_OPERATOR_EQUALS,
				Value:    "17",
			},
		},
	}
	os := NewOracleSpec([]string{"0xDEADBEEF"}, []*Filter{filter})
	assert.Equal(t, "2fd2f1fbcb2d9ecc102524b6ef48b189ed020eb1fecb613f35ad20471a0e9a96", os.Id)
}
