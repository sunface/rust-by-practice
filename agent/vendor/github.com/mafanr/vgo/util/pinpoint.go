package util

// PinpointData pinpointdata
type PinpointData struct {
	Type      string           `msg:"type"`
	AgentName string           `msg:"agentName"`
	AgentID   string           `msg:"agentID"`
	SpanTime  int64            `msg:"spanTime"`
	Payload   []*SpanDataModel `msg:"payload"`
}

// type DataType : SpanV2 SpanChunk AgentStat AgentStatBatch
// SpanDataModel data
type SpanDataModel struct {
	Type  string `msg:"type"`
	Spans []byte `msg:"spans"`
}
