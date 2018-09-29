package alert

type Alert struct {
}

func New() *Alert {
	return &Alert{}
}

func (a *Alert) Start() error {
	return nil
}

func (a *Alert) Close() error {
	return nil
}
