package util

const (
	TypeOfPinpoint = uint16(iota + 1) // 	监控数据
	TypeOfLog      = uint16(iota + 1) // 	日志数据
	TypeOfPing     = uint16(iota + 1) // 	日志数据
)

const (
	// MaxMessageSize max message size
	MaxMessageSize    int  = 65536
	TypeOfCompressYes byte = 1
	TypeOfCompressNo  byte = 2
)
