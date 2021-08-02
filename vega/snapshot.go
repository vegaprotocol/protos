package vega

import (
	"github.com/golang/protobuf/proto"
)

// Marshal is a convenience func to avoid weird issues when marshalling in other modules
func Marshal(msg proto.Message) ([]byte, error) {
	return proto.Marshal(msg)
}

// Unmarshal is for convenience
func Unmarshal(data []byte, msg proto.Message) error {
	return proto.Unmarshal(data, msg)
}
