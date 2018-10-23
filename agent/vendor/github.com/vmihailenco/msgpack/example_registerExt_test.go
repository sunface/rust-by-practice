package msgpack_test

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/vmihailenco/msgpack"
)

func init() {
	msgpack.RegisterExt(0, (*EventTime)(nil))
}

// https://github.com/fluent/fluentd/wiki/Forward-Protocol-Specification-v1#eventtime-ext-format
type EventTime struct {
	time.Time
}

var _ msgpack.Marshaler = (*EventTime)(nil)
var _ msgpack.Unmarshaler = (*EventTime)(nil)

func (tm *EventTime) MarshalMsgpack() ([]byte, error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, uint32(tm.Unix()))
	binary.BigEndian.PutUint32(b[4:], uint32(tm.Nanosecond()))
	return b, nil
}

func (tm *EventTime) UnmarshalMsgpack(b []byte) error {
	if len(b) != 8 {
		return fmt.Errorf("invalid data length: got %d, wanted 8", len(b))
	}
	sec := binary.BigEndian.Uint32(b)
	usec := binary.BigEndian.Uint32(b[4:])
	tm.Time = time.Unix(int64(sec), int64(usec))
	return nil
}

func ExampleRegisterExt() {
	b, err := msgpack.Marshal(&EventTime{time.Unix(123456789, 123)})
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = msgpack.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}
	fmt.Println(v.(*EventTime).UTC())

	tm := new(EventTime)
	err = msgpack.Unmarshal(b, tm)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm.UTC())

	// Output: 1973-11-29 21:33:09.000000123 +0000 UTC
	// 1973-11-29 21:33:09.000000123 +0000 UTC
}
