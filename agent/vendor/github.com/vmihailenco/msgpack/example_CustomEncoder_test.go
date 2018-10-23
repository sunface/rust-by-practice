package msgpack_test

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type customStruct struct {
	S string
	N int
}

var _ msgpack.CustomEncoder = (*customStruct)(nil)
var _ msgpack.CustomDecoder = (*customStruct)(nil)

func (s *customStruct) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.EncodeMulti(s.S, s.N)
}

func (s *customStruct) DecodeMsgpack(dec *msgpack.Decoder) error {
	return dec.DecodeMulti(&s.S, &s.N)
}

func ExampleCustomEncoder() {
	b, err := msgpack.Marshal(&customStruct{S: "hello", N: 42})
	if err != nil {
		panic(err)
	}

	var v customStruct
	err = msgpack.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", v)
	// Output: msgpack_test.customStruct{S:"hello", N:42}
}
