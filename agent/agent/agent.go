package agent

import (
	"github.com/containous/traefik/log"
	"github.com/mafanr/vgo/util"
)

type Agent struct {
	consumeC chan util.Original
}

func New() *Agent {
	return &Agent{
		consumeC: make(chan util.Original, 1000),
	}
}

func (a *Agent) Start() error {
	return nil
}

func (a *Agent) Close() error {
	return nil
}

func (a *Agent) initVgo() error {
	// connect vgo
	return nil
}

func (a *Agent) work() {
	// read
	// route
}

func (a *Agent) consume() {
	// select consume chan
	for {
		select {
		case msg, ok := <-a.consumeC:
			//g.L.Info(mok)
			log.Println(msg, ok)
		}
	}
}
