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
	quitC chan bool
	repC  chan *util.APMPacket
	serC  chan *util.APMPacket
}

func New() *Agent {
	return &Agent{
		quitC: make(chan bool, 1),
		repC:  make(chan *util.APMPacket, misc.Conf.Agent.ReportCacheLen),
		serC:  make(chan *util.APMPacket, misc.Conf.Agent.SerCacheLen),
	}
}

func (a *Agent) Start() error {
	// 启动report
	go a.report()

	// 初始化alert服务
	go a.initAlert()

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
	// 缓存列表
	reportList := make([]*util.APMPacket, 0, misc.Conf.Agent.ReportCacheLen+500)
	for {
		select {
		case msg, ok := <-a.repC:
			if ok {
				if len(a.repC) < misc.Conf.Agent.ReportCacheLen {
					reportList = append(reportList, msg)
				}
				if len(reportList) >= misc.Conf.Agent.ReportLen {
					// report
				}
				log.Println("time", reportList)
			}
		case <-tc.C:
			if len(reportList) > 0 {
				log.Println("time", reportList)
			}
		}
	}
}

func (a *Agent) initAlert() error {
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
		}
		// 是否重启
		if isRestart {
			a.initAlert()
		}
		return
	}()

	// connect alert
	for {
		con, err = net.Dial("tcp", misc.Conf.Agent.AlertAddr)
		if err != nil {
			g.L.Warn("Connect alert", zap.String("err", err.Error()), zap.String("addr", misc.Conf.Agent.AlertAddr))
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	for {
		apmPacket, err := a.readAlert(con, quitC)
		if err != nil {
			g.L.Warn("readAlert", zap.Error(err))
			return err
		}
		log.Println(apmPacket)
		a.serC<- apmPacket
		// 处理数据
	}

	return nil
}

func (a *Agent) readAlert(con net.Conn, quitC chan bool)(*util.APMPacket, error) {
	// read packet head
	// read packet body
	// compress
	// return
	return &util.APMPacket{}, nil
}


//
//func (a *Agent) work() {
//	// read
//	// route
//}
//
//func (a *Agent) collectRoute() {
//	// 缓存列表
//	cacheList := make([]util.APMPacket, 0, misc.Conf.Agent.CacheLen)
//	// 定时器
//	tc := time.NewTicker(time.Duration(misc.Conf.Agent.ReportInterval) * time.Second)
//	// select consume chan
//	for {
//		select {
//		//case msg, ok := <-a.Packet:
//		//	if ok {
//		//		if len(cacheList) < misc.Conf.Agent.CacheLen {
//		//			cacheList = append(cacheList, msg)
//		//		}
//		//		if len(cacheList) >= misc.Conf.Agent.ReportLen {
//		//			// report
//		//		}
//		//	}
//		//	log.Println(msg, ok)
//		//case <-a.quitC:
//		//	return
//		case <-tc.C:
//			log.Println("report msg", cacheList)
//		}
//	}
//}
