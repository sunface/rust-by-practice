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

// Ping ...
type Ping struct {
	// Name string `msg:"n"`
}

// NewPing ...
func NewPing() *Ping {
	return &Ping{}
}
