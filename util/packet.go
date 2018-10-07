package util

import (
	"encoding/binary"
	"github.com/mafanr/g"
	"go.uber.org/zap"
	"io"
)

// Packet 通用报文
type APMPacket struct {
	Cmds      []*CMD          `msg:"cmd"`
	Pinpoints []*PinpointData `msg:"tp"`
	//Logs    []*LogPacket    `msg:"lp"`
	//Systems []*SystemPacket `msg:"sp"`
}

func (ap *APMPacket) Len() int {
	return len(ap.Pinpoints) + len(ap.Cmds)
}

func (ap *APMPacket) Clear() {
	ap.Cmds = ap.Cmds[:0]
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
		g.L.Warn("Decode", zap.String("err", err.Error()))
		return err
	}
	b.IsCompress = buf[0]
	b.Len = binary.BigEndian.Uint32(buf[1:5])

	b.PayLoad = make([]byte, b.Len)
	if b.Len > 0 {
		_, err := io.ReadFull(rdr, b.PayLoad)
		if err != nil {
			g.L.Warn("Decode", zap.String("err", err.Error()))
			return err
		}
	}
	return nil
}

//
//// DecodePacket decodes the packet from the provided reader.
//func DecodePacket(rdr io.Reader) (*BatchAPMPacket, error) {
//	b := &BatchAPMPacket{}
//	messageType, sizeOf, err := b.Decode(rdr)
//	if err != nil {
//		return nil, err
//	}
//
//	// 数据长度异常
//	if int(sizeOf) > MaxMessageSize {
//		return nil, fmt.Errorf("Packet size is too large")
//	}
//
//	buffer := make([]byte, sizeOf)
//	_, err = io.ReadFull(rdr, buffer)
//	if err != nil {
//		return nil, err
//	}
//
//	// Decode the body
//	var msg Message
//	switch messageType {
//	case TypeOfConnect:
//		msg, err = decodeConnect(buffer)
//	default:
//		return nil, fmt.Errorf("Invalid type packet with type %d", messageType)
//	}
//
//	return msg, nil
//}
