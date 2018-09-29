package agent

type Agent struct {
}

func New() *Agent {
	return &Agent{}
}

func (a *Agent) Start() error {
	return nil
}

func (a *Agent) Close() error {
	return nil
}