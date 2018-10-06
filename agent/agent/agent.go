package agent

import (
	"github.com/containous/traefik/log"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/util"
	"go.uber.org/zap"
	"time"
)

type Agent struct {
	quitC     chan bool
	pinpointC chan *util.PinpointData
	cmdC      chan *util.CMD
	client    *TcpClient
	//repC   chan *util.APMPacket
}

var gAgent *Agent

func New() *Agent {
	gAgent = &Agent{
		quitC:     make(chan bool, 1),
		pinpointC: make(chan *util.PinpointData, misc.Conf.Agent.PinpointCacheLen),
		cmdC:      make(chan *util.CMD, misc.Conf.Agent.CmdCacheLen),
		client:    NewTcpClient(),
	}
	return gAgent
}

func (a *Agent) Start() error {
	// 启动report
	go a.report()

	// 初始化处理下行命令等
	go a.dealCmdPacket()

	// 初始化tcp client
	a.client.Init()

	// 启动本地接收采集信息端口

	return nil
}

func (a *Agent) Close() error {

	return nil
}

func (a *Agent) report() {
	defer func() {
		if err := recover(); err != nil {
			g.L.Warn("collector panic", zap.Stack("server"), zap.Any("err", err))
		}
		return
	}()
	// 定时器
	tc := time.NewTicker(time.Duration(misc.Conf.Agent.ReportInterval) * time.Second)

	// 缓存
	apmPacket := util.NewAPMPacket()

	for {
		select {
		case t, ok := <-a.pinpointC:
			if ok {
				apmPacket.Pinpoints = append(apmPacket.Pinpoints, t)
				if apmPacket.Len() > misc.Conf.Agent.ReportLen {
					// report
					apmPacket.Clear()
				}
			}
			break
		case <-tc.C:
			if apmPacket.Len() > misc.Conf.Agent.ReportLen {
				// report
				apmPacket.Clear()
			}
			log.Println("time", apmPacket)
		}
	}
}

func (a *Agent) dealCmdPacket() {
	for {
		select {
		case p, ok := <-a.cmdC:
			if ok {
				log.Println(p)
			}
		case <-a.quitC:
			return
		}
	}
}
