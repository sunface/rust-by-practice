package stats

// Stats 离线计算
type Stats struct {
}

func New() *Stats {
	return &Stats{}
}

func (s *Stats) Start() error {
	return nil
}

func (s *Stats) Close() error {
	return nil
}
