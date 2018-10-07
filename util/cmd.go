package util

import (
	"encoding/binary"
	"github.com/mafanr/g"
	"go.uber.org/zap"
	"io"
)

type CMD struct {
	Type    uint16
	Len     uint32
	PayLoad []byte
}

func NewCMD() *CMD {
	return &CMD{}
}

// Encode encode
func (c *CMD) Encode() []byte {
	c.Len = uint32(len(c.PayLoad))
	buf := make([]byte, c.Len+6)
	binary.BigEndian.PutUint16(buf[0:2], c.Type)
	binary.BigEndian.PutUint32(buf[2:6], c.Len)
	if c.Len > 0 {
		copy(buf[6:], c.PayLoad)
	}
	return buf
}

// Decode decode
func (c *CMD) Decode(rdr io.Reader) error {
	buf := make([]byte, 6)
	if _, err := io.ReadFull(rdr, buf); err != nil {
		g.L.Warn("Decode", zap.String("err", err.Error()))
		return err
	}
	c.Type = binary.BigEndian.Uint16(buf[0:2])
	c.Len = binary.BigEndian.Uint32(buf[2:6])
	c.PayLoad = make([]byte, c.Len)

	if c.Len > 0 {
		_, err := io.ReadFull(rdr, c.PayLoad)
		if err != nil {
			g.L.Warn("Decode", zap.String("err", err.Error()))
			return err
		}
	}
	return nil
}
