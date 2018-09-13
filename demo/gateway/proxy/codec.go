package proxy

import "github.com/golang/protobuf/proto"

type Code struct{}

func (c *Code) Marshal(v interface{}) ([]byte, error) {
	frame, ok := v.(*Frame)
	if !ok {
		return proto.Marshal(v.(proto.Message))
	}

	return frame.Buf, nil
}

func (c *Code) Unmarshal(data []byte, v interface{}) error {
	frame, ok := v.(*Frame)
	if !ok {
		return proto.Unmarshal(data, v.(proto.Message))
	}
	frame.Buf = data
	return nil
}

func (c *Code) String() string {
	return "proto"
}
