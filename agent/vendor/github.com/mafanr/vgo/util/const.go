package util

const (
	TypeOfSkywalking byte = 1 // 	Skywalking 监控数据 uint16(iota + 1)
	TypeOfPinpoint   byte = 2 // 	Pinpoint 日志数据
	TypeOfCmd        byte = 3 // 	指令包 数据
	TypeOfLog        byte = 4 // 	日志数据
)

const (
	TypeOfPing uint16 = 100 // 	Skywalking 监控数据 uint16(iota + 1) // 	Skywalking 监控数据
)

const (
	TypeOfAppRegister uint16 = 1 // TypeOfAppRegister
)

const (
	// MaxMessageSize max message size
	MaxMessageSize    int  = 16 * 1024
	TypeOfCompressYes byte = 1 // 数据压缩
	TypeOfCompressNo  byte = 2 // 数据不压缩
	TypeOfSyncYes     byte = 1 // 同步
	TypeOfSyncNo      byte = 2 // 非同步
)

const (
	VersionOf01 byte = 1
)
