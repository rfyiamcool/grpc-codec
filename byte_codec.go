package codec

import (
	"errors"
	"fmt"
)

type byteCodec struct {
	payload []byte
}

func NewByteCodec() *byteCodec {
	return &byteCodec{}
}

func (r *byteCodec) Marshal(v interface{}) ([]byte, error) {
	switch v.(type) {
	case []byte:
		{
			r.payload = v.([]byte)
			return r.payload, nil
		}
	default:
		{
			return nil, errors.New("Invalid data type, must bytes")
		}
	}
}

func (r *byteCodec) Unmarshal(data []byte, v interface{}) error {
	switch v.(type) {
	case *[]byte:
		{
			out := v.(*[]byte)
			if out != nil {
				*out = append(*out, data...)
			}
			return nil
		}
	default:
		{
			return errors.New("Invalid data type, must bytes")
		}
	}
}

func (r *byteCodec) String() string {
	return fmt.Sprintf("%s", string(r.payload))
}
