package proto

import (
	"github.com/go-sohunjug/kratos-transport/codec"
	"google.golang.org/protobuf/proto"
)

type Marshaler struct{}

func (Marshaler) Name() string {
	return "proto"
}

func (Marshaler) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, codec.ErrInvalidMessage
	}

	return proto.Marshal(pb)
}

func (Marshaler) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return codec.ErrInvalidMessage
	}

	return proto.Unmarshal(data, pb)
}
