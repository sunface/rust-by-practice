package service

type Vgo struct {
}

func New() *Vgo {
	return &Vgo{}
}

func (v *Vgo) Start() error {
	return nil
}

func (v *Vgo) Close() error {
	return nil
}
