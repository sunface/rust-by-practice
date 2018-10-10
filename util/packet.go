package util

import (
	"encoding/binary"
	"github.com/mafanr/g"
	"go.uber.org/zap"
	"io"
)

// Packet 通用报文
type APMPacket struct {
	Cmd       []*CMD          `msg:"cmd"`
	Pinpoints []*PinpointData `msg:"tp"`
	Other     []byte          `msg:"ot"`
	//Logs    []*LogPacket    `msg:"lp"`
	//Systems []*SystemPacket `msg:"sp"`
}

func (ap *APMPacket) Len() int {
	return len(ap.Pinpoints) + len(ap.Cmd)
}

func (ap *APMPacket) Clear() {
	ap.Cmd = ap.Cmd[:0]
	ap.Pinpoints = ap.Pinpoints[:0]
}

func NewAPMPacket() *APMPacket {
	return &APMPacket{}
}

// BatchAPMPacket ... node transfer packet
type BatchAPMPacket struct {
	IsCompress byte
	Len        uint32
	PayLoad    []byte
}

// Encode encode
func (b *BatchAPMPacket) Encode() []byte {
	b.Len = uint32(len(b.PayLoad))
	buf := make([]byte, b.Len+5)
	buf[0] = b.IsCompress
	binary.BigEndian.PutUint32(buf[1:5], b.Len)
	if b.Len > 0 {
		copy(buf[5:], b.PayLoad)
	}
	return buf
}

// Decode decode
func (b *BatchAPMPacket) Decode(rdr io.Reader) error {
	buf := make([]byte, 5)
	if _, err := io.ReadFull(rdr, buf); err != nil {
		g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
		return err
	}
	b.IsCompress = buf[0]
	b.Len = binary.BigEndian.Uint32(buf[1:5])

	b.PayLoad = make([]byte, b.Len)
	if b.Len > 0 {
		_, err := io.ReadFull(rdr, b.PayLoad)
		if err != nil {
			g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
			return err
		}
	}
	return nil
}
