package agent

import (
	"github.com/containous/traefik/log"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/util"
	"go.uber.org/zap"
	"net"
	"time"
)

type Agent struct {
	consumeC chan util.Original
	vgo      *Vgo
}

func New() *Agent {
	return &Agent{
		consumeC: make(chan util.Original, 1000),
		vgo:      NewVgo(),
	}
}

func (a *Agent) Start() error {
	// 启动接收本地插件信息 chan 携程
	// 初始化链接vgo
	// 启动本地接收采集信息端口

	return nil
}

func (a *Agent) Close() error {
	return nil
}

func (a *Agent) initVgo() error {
	var con net.Conn
	var err error
	var quitC chan bool
	isRestart := true
	defer func() {
		close(quitC)
		if con != nil {
			con.Close()
		}
		if err := recover(); err != nil {
			g.L.Warn("server panic", zap.Stack("server"), zap.Any("err", err))
			return
		}
		// 是否重启
		if isRestart {
			a.initVgo()
		}
	}()
	// connect vgo
	for {
		con, err = net.Dial("tcp", misc.Conf.Agent.VgoAddr)
		if err != nil {
			g.L.Warn("Connect Vgo", zap.String("err", err.Error()), zap.String("addr", misc.Conf.Agent.VgoAddr))
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	// quit chan
	// go read
	// go write
	// select <- dataC

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
