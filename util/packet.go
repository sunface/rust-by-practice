package util

import (
	"encoding/binary"

	"io"

	"github.com/golang/snappy"
	"github.com/mafanr/g"

	"go.uber.org/zap"
)

// VgoPacket ...
type VgoPacket struct {
	Type       uint16
	Len        uint32
	IsCompress byte
	PayLoad    []byte
}

// NewVgoPacket ...
func NewVgoPacket() *VgoPacket {
	return &VgoPacket{}
}

// Encode encode
func (v *VgoPacket) Encode() []byte {
	// 压缩
	if v.IsCompress == TypeOfCompressYes {
		if len(v.PayLoad) > 0 {
			compressBuf := snappy.Encode(nil, v.PayLoad)
			v.PayLoad = compressBuf
		}
	}

	v.Len = uint32(len(v.PayLoad))
	buf := make([]byte, v.Len+7)

	binary.BigEndian.PutUint16(buf[:2], v.Type)
	buf[2] = v.IsCompress
	binary.BigEndian.PutUint32(buf[3:7], v.Len)

	if v.Len > 0 {
		copy(buf[7:], v.PayLoad)
	}
	return buf
}

// Decode decode
func (v *VgoPacket) Decode(rdr io.Reader) error {
	buf := make([]byte, 7)
	if _, err := io.ReadFull(rdr, buf); err != nil {
		g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
		return err
	}
	v.Type = binary.BigEndian.Uint16(buf[:2])
	v.IsCompress = buf[2]

	length := binary.BigEndian.Uint32(buf[3:7])

	payLoad := make([]byte, length)
	if length > 0 {
		_, err := io.ReadFull(rdr, payLoad)
		if err != nil {
			g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
			return err
		}

		// 解压
		if v.IsCompress == TypeOfCompressYes {
			v.PayLoad, err = snappy.Decode(nil, payLoad)
			if err != nil {
				g.L.Warn("Decode:snappy.Decode", zap.String("error", err.Error()))
				return err
			}
		} else {
			v.PayLoad = payLoad
		}
		v.Len = uint32(len(v.PayLoad))
	}
	return nil
}

//
//// Packet 通用报文
//type APMPacket struct {
//	Cmd       []*CMD          `msg:"cmd"`
//	Pinpoints []*PinpointData `msg:"tp"`
//	Other     []byte          `msg:"ot"`
//	//Logs    []*LogPacket    `msg:"lp"`
//	//Systems []*SystemPacket `msg:"sp"`
//}
//
//func (ap *APMPacket) Len() int {
//	return len(ap.Pinpoints) + len(ap.Cmd)
//}
//
//func (ap *APMPacket) Clear() {
//	ap.Cmd = ap.Cmd[:0]
//	ap.Pinpoints = ap.Pinpoints[:0]
//}
//
//func NewAPMPacket() *APMPacket {
//	return &APMPacket{}
//}
//
//// BatchAPMPacket ... node transfer packet
//type BatchAPMPacket struct {
//	IsCompress byte
//	Len        uint32
//	PayLoad    []byte
//}
//
//// Encode encode
//func (b *BatchAPMPacket) Encode() []byte {
//	b.Len = uint32(len(b.PayLoad))
//	buf := make([]byte, b.Len+5)
//	buf[0] = b.IsCompress
//	binary.BigEndian.PutUint32(buf[1:5], b.Len)
//	if b.Len > 0 {
//		copy(buf[5:], b.PayLoad)
//	}
//	return buf
//}
//
//// Decode decode
//func (b *BatchAPMPacket) Decode(rdr io.Reader) error {
//	buf := make([]byte, 5)
//	if _, err := io.ReadFull(rdr, buf); err != nil {
//		g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
//		return err
//	}
//	b.IsCompress = buf[0]
//	b.Len = binary.BigEndian.Uint32(buf[1:5])
//
//	b.PayLoad = make([]byte, b.Len)
//	if b.Len > 0 {
//		_, err := io.ReadFull(rdr, b.PayLoad)
//		if err != nil {
//			g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
//			return err
//		}
//	}
//	return nil
//}
