package util

// Packet 通用报文
type APMPacket struct {
	Pinpoints []*PinpointData `msg:"tp"`
	//Logs    []*LogPacket    `msg:"lp"`
	//Systems []*SystemPacket `msg:"sp"`
}

func (ap *APMPacket) Len() int {
	return len(ap.Pinpoints)
}

func (ap *APMPacket) Clear() {
	ap.Pinpoints = ap.Pinpoints[:0]
}

func NewAPMPacket() *APMPacket {
	return &APMPacket{}
}

type BatchAPMPacket struct {
	Len        int
	IsCompress byte
	PayLoad    []byte
}
