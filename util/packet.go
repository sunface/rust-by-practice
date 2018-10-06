package util

// Packet 通用报文
type APMPacket struct {
	Traces  []*TracePacket  `msg:"tp"`
	//Logs    []*LogPacket    `msg:"lp"`
	//Systems []*SystemPacket `msg:"sp"`
}

func (ap*APMPacket)Len() int {
	return len(ap.Traces)
}


func (ap*APMPacket)Clear()  {
	ap.Traces = ap.Traces[:0]
}

func NewAPMPacket() *APMPacket {
	return &APMPacket{
	}
}

type BatchAPMPacket struct {
	Len        int
	IsCompress byte
	PayLoad    []byte
}
