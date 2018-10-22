package util

// SkywalkingPacket ...
type SkywalkingPacket struct {
	Type    uint16 `msg:"type"`
	Payload []byte `msg:"payload"`
}
