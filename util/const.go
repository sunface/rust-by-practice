package util

const (
	TypeOfSkywalking = uint16(iota + 1) // 	Pinpoint 监控数据
	TypeOfPinpoint                      // 	Skywalking 日志数据
	TypeOfPing                          // 	日志数据
	TypeOfLog                           // 	日志数据
)

const (
	// MaxMessageSize max message size
	MaxMessageSize    int  = 16 * 1024
	TypeOfCompressYes byte = 1
	TypeOfCompressNo  byte = 2
)
