package util

// SkywalkingPacket ...
type SkywalkingPacket struct {
	Type    uint16 `msg:"type"`
	Payload []byte `msg:"payload"`
}

// AppRegister ...
type AppRegister struct {
	Name string `msg:"n"`
	Code int32  `msg:"c"`
}
