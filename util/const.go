package util

const (
	TypeOfPinpoint = uint8(iota + 1) // 	监控数据
	TypeOfLog      = uint8(iota + 1) // 	日志数据
)

// MaxMessageSize max message size
const MaxMessageSize int = 65536
