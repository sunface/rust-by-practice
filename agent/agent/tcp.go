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

// TcpClient tcp client
type TcpClient struct {
	con net.Conn
}

// NewTcpClient ...
func NewTcpClient() *TcpClient {
	return &TcpClient{}
}

// Init ...
func (t *TcpClient) Init() error {
	var con net.Conn
	var err error
	isRestart := true
	defer func() {
		if con != nil {
			con.Close()
		}
		if err := recover(); err != nil {
			g.L.Warn("server panic", zap.Stack("server"), zap.Any("err", err))
		}
		// 是否重启
		if isRestart {
			t.Init()
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
		t.con = con
		break
	}
	// 启动心跳
	go t.KeepLive()

	for {
		cmdPacket, err := t.ReadPacket()
		if err != nil {
			g.L.Warn("readAlert", zap.Error(err))
			return err
		}
		// 发给上层处理
		gAgent.cmdC <- cmdPacket
	}

	return nil
}

// KeepLive ...
func (t *TcpClient) KeepLive() {
	for {
		log.Println("I'm Ping !")
		time.Sleep(time.Second * 10)
	}
}

// ReadPacket ...
func (t *TcpClient) ReadPacket() (*util.CMD, error) {
	return &util.CMD{}, nil
}

// WritePacket ...
func (t *TcpClient) WritePacket(p *util.APMPacket) error {
	return nil
}

// Close ....
func (t *TcpClient) Close() error {
	if t.con != nil {
		t.con.Close()
	}
	return nil
}
