package util

// Original 通用原始报文
type Original struct {
	Type    uint8  `msg:"type"`
	Payload []byte `msg:"payload"`
}
