package util

// CMD ...
type CMD struct {
	Type    uint16 `msg:"t"`
	PayLoad []byte `msg:"p"`
}

// NewCMD ...
func NewCMD() *CMD {
	return &CMD{}
}

// // Encode encode
// func (c *CMD) Encode() []byte {
// 	c.Len = uint32(len(c.PayLoad))
// 	buf := make([]byte, c.Len+6)
// 	binary.BigEndian.PutUint16(buf[0:2], c.Type)
// 	binary.BigEndian.PutUint32(buf[2:6], c.Len)
// 	if c.Len > 0 {
// 		copy(buf[6:], c.PayLoad)
// 	}
// 	return buf
// }

// // Decode decode
// func (c *CMD) Decode(rdr io.Reader) error {
// 	buf := make([]byte, 6)
// 	if _, err := io.ReadFull(rdr, buf); err != nil {
// 		g.L.Warn("Decode", zap.String("err", err.Error()))
// 		return err
// 	}
// 	c.Type = binary.BigEndian.Uint16(buf[0:2])
// 	c.Len = binary.BigEndian.Uint32(buf[2:6])
// 	c.PayLoad = make([]byte, c.Len)

// 	if c.Len > 0 {
// 		_, err := io.ReadFull(rdr, c.PayLoad)
// 		if err != nil {
// 			g.L.Warn("Decode:io.ReadFull", zap.String("err", err.Error()))
// 			return err
// 		}
// 	}
// 	return nil
// }

// // AgentInfo ...
// type AgentInfo struct {
// 	AgentName string `msg:"an"`
// 	Host      string `msg:"h"`
// 	AppName   string `msg:"an"`
// }

// Ping ...
type Ping struct {
	// Name string `msg:"n"`
}

// NewPing ...
func NewPing() *Ping {
	return &Ping{}
}
