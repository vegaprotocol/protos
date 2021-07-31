package vega

import (
	"github.com/golang/protobuf/proto"
)

// MarshalSnapshot is a convenience func to avoid weird issues when marshalling in other modules
func MarshalSnapshot(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}
