package util

const (
	TypeOfSkywalking = uint16(iota + 1) // 	Skywalking 监控数据
	TypeOfPinpoint                      // 	Pinpoint 日志数据
	TypeOfCmd                           // 	Cmd 数据
	TypeOfLog                           // 	日志数据
)

const (
	TypeOfPing = uint16(iota + 1) // 	Skywalking 监控数据
)

const (
	// MaxMessageSize max message size
	MaxMessageSize    int  = 16 * 1024
	TypeOfCompressYes byte = 1
	TypeOfCompressNo  byte = 2
)
