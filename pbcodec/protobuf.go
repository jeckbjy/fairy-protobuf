package pbcodec

import (
	"github.com/jeckbjy/fairy"

	"github.com/golang/protobuf/proto"
)

// New create protobuf codec
func New() *ProtobufCodec {
	codec := &ProtobufCodec{}
	return codec
}

type ProtobufCodec struct {
}

func (*ProtobufCodec) Encode(buffer *fairy.Buffer, obj interface{}) error {
	msg := obj.(proto.Message)
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	buffer.Append(data)
	return nil
}

func (*ProtobufCodec) Decode(buffer *fairy.Buffer, obj interface{}) error {
	data := buffer.ReadToEnd()
	msg := obj.(proto.Message)
	return proto.Unmarshal(data, msg)
}
