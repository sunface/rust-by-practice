package util

// Packet 通用报文
type APMPacket struct {
	Type       uint8  `msg:"type"`
	IsCompress bool   `msg:"is_compress"`
	Len        int    `msg:"len"`
	Payload    []byte `msg:"payload"`
}
