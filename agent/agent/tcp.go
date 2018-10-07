package agent

import (
	"bufio"
	"github.com/containous/traefik/log"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/util"
	"go.uber.org/zap"
	"io"
	"net"
	"time"
)

// TcpClient tcp client
type TcpClient struct {
	conn net.Conn
}

// NewTcpClient ...
func NewTcpClient() *TcpClient {
	return &TcpClient{}
}

// Init ...
func (t *TcpClient) Init() error {
	var conn net.Conn
	var err error
	isRestart := true
	defer func() {
		if conn != nil {
			conn.Close()
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
		conn, err = net.Dial("tcp", misc.Conf.Agent.VgoAddr)
		if err != nil {
			g.L.Warn("Connect vgo", zap.String("err", err.Error()), zap.String("addr", misc.Conf.Agent.VgoAddr))
			time.Sleep(5 * time.Second)
			continue
		}
		t.conn = conn
		break
	}
	// 启动心跳
	go t.KeepLive()
	reader := bufio.NewReaderSize(t.conn, util.MaxMessageSize)
	for {
		cmdPacket, err := t.ReadPacket(reader)
		if err != nil {
			g.L.Warn("ReadPacket", zap.Error(err))
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
func (t *TcpClient) ReadPacket(rdr io.Reader) (*util.CMD, error) {
	cmd := util.NewCMD ()
	if err:= cmd.Decode(rdr); err!=nil {
		g.L.Warn("ReadPacket", zap.String("error", err.Error()))
		return nil, err
	}
	return cmd, nil
}

// WritePacket ...
func (t *TcpClient) WritePacket(p *util.APMPacket) error {
	return nil
}

// Close ....
func (t *TcpClient) Close() error {
	if t.conn != nil {
		t.conn.Close()
	}
	return nil
}
