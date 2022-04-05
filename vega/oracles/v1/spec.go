package v1

import (
	"encoding/hex"
	"strings"

	"code.vegaprotocol.io/protos/crypto"
)

func (OracleSpec) IsEvent() {}

func NewOracleSpec(pubKeys []string, filters []*Filter) *OracleSpec {
	return &OracleSpec{
		Id:      newID(pubKeys, filters),
		PubKeys: pubKeys,
		Filters: filters,
	}
}

func newID(pubKeys []string, filters []*Filter) string {
	buf := []byte{}
	for _, filter := range filters {
		s := filter.Key.Name + filter.Key.Type.NonDeterministicString()
		for _, c := range filter.Conditions {
			s += c.Operator.NonDeterministicString() + c.Value
		}

		buf = append(buf, []byte(s)...)
	}
	buf = append(buf, []byte(strings.Join(pubKeys, ""))...)

	return hex.EncodeToString(crypto.Hash(buf))
}
