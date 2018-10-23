package util

import (
	"encoding/binary"
	"io"

	"github.com/golang/snappy"
	"github.com/mafanr/g"
	"go.uber.org/zap"
)

// VgoPacket 通用报文
type VgoPacket struct {
	Type       byte   `msgp:"t"`  // 类型
	Version    byte   `msgp:"v"`  // 版本
	IsSync     byte   `msgp:"is"` // 是否同步
	IsCompress byte   `msgp:"ic"` // 是否压缩
	ID         uint32 `msgp:"i"`  // 报文ID
	Len        uint32 `msgp:"l"`  // 长度
	PayLoad    []byte `msgp:"p"`  // 数据
}

// // NewVgoPacket ...
// func NewVgoPacket(pType byte, version byte, isSync byte, isCompress byte, id uint32, payload []byte) *VgoPacket {
// 	return &VgoPacket{
// 		Type:       pType,
// 		Version:    version,
// 		IsSync:     isSync,
// 		IsCompress: isCompress,
// 		ID:         id,
// 		PayLoad:    payload,
// 	}
// }

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
	buf := make([]byte, v.Len+12)

	buf[0] = v.Type
	buf[1] = v.Version
	buf[2] = v.IsSync
	buf[3] = v.IsCompress
	binary.BigEndian.PutUint32(buf[4:8], v.ID)
	binary.BigEndian.PutUint32(buf[8:12], v.Len)

	if v.Len > 0 {
		copy(buf[12:], v.PayLoad)
	}
	return buf

}

// Decode decode
func (v *VgoPacket) Decode(rdr io.Reader) error {
	buf := make([]byte, 12)
	if _, err := io.ReadFull(rdr, buf); err != nil {
		g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
		return err
	}

	v.Type = buf[0]
	v.Version = buf[1]
	v.IsSync = buf[2]
	v.IsCompress = buf[3]
	v.ID = binary.BigEndian.Uint32(buf[4:8])

	length := binary.BigEndian.Uint32(buf[8:12])
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
