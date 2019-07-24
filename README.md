# grpc-codec

custom codec for grpc server sidecar

`fix it`

```
... is not proto.Message: missing method ProtoMessage
```

## Usage

```
import (
	"github.com/rfyiamcool/grpc-codec"
)

gs := grpc.NewServer(
	grpc.MaxConcurrentStreams(3000),
	grpc.UnknownServiceHandler(defaultHandler),
	grpc.CustomCodec(codec.GetProxyCodec()),
)


func defaultHandler(srv interface{}, stream grpc.ServerStream) error {
	...
	...
	var (
		raw  = &codec.frame{}
	)

	err := stream.RecvMsg(raw)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(raw.payload)
	...
}

```
