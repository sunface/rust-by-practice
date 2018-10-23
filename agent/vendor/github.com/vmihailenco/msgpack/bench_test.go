package msgpack_test

import (
	"bytes"
	"io/ioutil"
	"math"
	"testing"
	"time"

	"github.com/vmihailenco/msgpack"
)

func BenchmarkDiscard(b *testing.B) {
	enc := msgpack.NewEncoder(ioutil.Discard)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := enc.Encode(nil); err != nil {
			b.Fatal(err)
		}
		if err := enc.Encode("hello"); err != nil {
			b.Fatal(err)
		}
	}
}

func benchmarkEncodeDecode(b *testing.B, src, dst interface{}) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	dec := msgpack.NewDecoder(&buf)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := enc.Encode(src); err != nil {
			b.Fatal(err)
		}
		if err := dec.Decode(dst); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBool(b *testing.B) {
	var dst bool
	benchmarkEncodeDecode(b, true, &dst)
}

func BenchmarkInt0(b *testing.B) {
	var dst int
	benchmarkEncodeDecode(b, 1, &dst)
}

func BenchmarkInt1(b *testing.B) {
	var dst int
	benchmarkEncodeDecode(b, -33, &dst)
}

func BenchmarkInt2(b *testing.B) {
	var dst int
	benchmarkEncodeDecode(b, 128, &dst)
}

func BenchmarkInt4(b *testing.B) {
	var dst int
	benchmarkEncodeDecode(b, 32768, &dst)
}

func BenchmarkInt8(b *testing.B) {
	var dst int
	benchmarkEncodeDecode(b, int64(2147483648), &dst)
}

func BenchmarkInt32(b *testing.B) {
	var dst int32
	benchmarkEncodeDecode(b, int32(0), &dst)
}

func BenchmarkTime(b *testing.B) {
	var dst time.Time
	benchmarkEncodeDecode(b, time.Now(), &dst)
}

func BenchmarkDuration(b *testing.B) {
	var dst time.Duration
	benchmarkEncodeDecode(b, time.Hour, &dst)
}

func BenchmarkByteSlice(b *testing.B) {
	src := make([]byte, 1024)
	var dst []byte
	benchmarkEncodeDecode(b, src, &dst)
}

func BenchmarkByteArray(b *testing.B) {
	var src [1024]byte
	var dst [1024]byte
	benchmarkEncodeDecode(b, src, &dst)
}

func BenchmarkByteArrayPtr(b *testing.B) {
	var src [1024]byte
	var dst [1024]byte
	benchmarkEncodeDecode(b, &src, &dst)
}

func BenchmarkMapStringString(b *testing.B) {
	src := map[string]string{
		"hello": "world",
		"foo":   "bar",
	}
	var dst map[string]string
	benchmarkEncodeDecode(b, src, &dst)
}

func BenchmarkMapStringStringPtr(b *testing.B) {
	src := map[string]string{
		"hello": "world",
		"foo":   "bar",
	}
	var dst map[string]string
	dstptr := &dst
	benchmarkEncodeDecode(b, src, &dstptr)
}

func BenchmarkMapStringInterface(b *testing.B) {
	src := map[string]interface{}{
		"hello": "world",
		"foo":   "bar",
	}
	var dst map[string]interface{}
	benchmarkEncodeDecode(b, src, &dst)
}

func BenchmarkMapIntInt(b *testing.B) {
	src := map[int]int{
		1: 10,
		2: 20,
	}
	var dst map[int]int
	benchmarkEncodeDecode(b, src, &dst)
}

func BenchmarkStringSlice(b *testing.B) {
	src := []string{"hello", "world"}
	var dst []string
	benchmarkEncodeDecode(b, src, &dst)
}

func BenchmarkStringSlicePtr(b *testing.B) {
	src := []string{"hello", "world"}
	var dst []string
	dstptr := &dst
	benchmarkEncodeDecode(b, src, &dstptr)
}

type benchmarkStruct struct {
	Name      string
	Age       int
	Colors    []string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type benchmarkStruct2 struct {
	Name      string
	Age       int
	Colors    []string
	Data      []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

var _ msgpack.CustomEncoder = (*benchmarkStruct2)(nil)
var _ msgpack.CustomDecoder = (*benchmarkStruct2)(nil)

func (s *benchmarkStruct2) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.EncodeMulti(
		s.Name,
		s.Colors,
		s.Age,
		s.Data,
		s.CreatedAt,
		s.UpdatedAt,
	)
}

func (s *benchmarkStruct2) DecodeMsgpack(dec *msgpack.Decoder) error {
	return dec.DecodeMulti(
		&s.Name,
		&s.Colors,
		&s.Age,
		&s.Data,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
}

func structForBenchmark() *benchmarkStruct {
	return &benchmarkStruct{
		Name:      "Hello World",
		Colors:    []string{"red", "orange", "yellow", "green", "blue", "violet"},
		Age:       math.MaxInt32,
		Data:      make([]byte, 1024),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func structForBenchmark2() *benchmarkStruct2 {
	return &benchmarkStruct2{
		Name:      "Hello World",
		Colors:    []string{"red", "orange", "yellow", "green", "blue", "violet"},
		Age:       math.MaxInt32,
		Data:      make([]byte, 1024),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func BenchmarkStructVmihailencoMsgpack(b *testing.B) {
	in := structForBenchmark()
	out := new(benchmarkStruct)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf, err := msgpack.Marshal(in)
		if err != nil {
			b.Fatal(err)
		}

		err = msgpack.Unmarshal(buf, out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStructMarshal(b *testing.B) {
	in := structForBenchmark()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := msgpack.Marshal(in)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStructUnmarshal(b *testing.B) {
	in := structForBenchmark()
	buf, err := msgpack.Marshal(in)
	if err != nil {
		b.Fatal(err)
	}
	out := new(benchmarkStruct)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = msgpack.Unmarshal(buf, out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStructManual(b *testing.B) {
	in := structForBenchmark2()
	out := new(benchmarkStruct2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf, err := msgpack.Marshal(in)
		if err != nil {
			b.Fatal(err)
		}

		err = msgpack.Unmarshal(buf, out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

type benchmarkStructPartially struct {
	Name string
	Age  int
}

func BenchmarkStructUnmarshalPartially(b *testing.B) {
	in := structForBenchmark()
	buf, err := msgpack.Marshal(in)
	if err != nil {
		b.Fatal(err)
	}
	out := new(benchmarkStructPartially)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err = msgpack.Unmarshal(buf, out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkQuery(b *testing.B) {
	var records []map[string]interface{}
	for i := 0; i < 1000; i++ {
		record := map[string]interface{}{
			"id":    i,
			"attrs": map[string]interface{}{"phone": i},
		}
		records = append(records, record)
	}

	bs, err := msgpack.Marshal(records)
	if err != nil {
		b.Fatal(err)
	}

	dec := msgpack.NewDecoder(bytes.NewBuffer(bs))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dec.Reset(bytes.NewBuffer(bs))

		values, err := dec.Query("10.attrs.phone")
		if err != nil {
			b.Fatal(err)
		}
		if values[0].(int8) != 10 {
			b.Fatalf("%v != %d", values[0], 10)
		}
	}
}
