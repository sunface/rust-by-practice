package service

import (
	"bufio"
	"io"
	"net"
	"time"

	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/util"
	"go.uber.org/zap"
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
	//var conn net.Conn
	var err error
	isRestart := true
	quitC := make(chan bool, 1)
	// 定时器
	keepLiveTc := time.NewTicker(time.Duration(misc.Conf.Agent.KeepLiveInterval) * time.Second)

	defer func() {
		if err := recover(); err != nil {
			g.L.Warn("Init:.", zap.Stack("server"), zap.Any("err", err))
		}
		// 是否重启
		if isRestart {
			t.Init()
		}
	}()

	defer func() {
		close(quitC)
		t.conn.Close()
		keepLiveTc.Stop()
	}()

	// connect vgo
	for {
		t.conn, err = net.Dial("tcp", misc.Conf.Agent.VgoAddr)
		if err != nil {
			g.L.Warn("Init:net.Dial", zap.String("err", err.Error()), zap.String("addr", misc.Conf.Agent.VgoAddr))
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	// 启动心跳
	go func() {
		for {
			select {
			case <-keepLiveTc.C:
				if err := t.KeepLive(); err != nil {
					g.L.Warn("Init:t.KeepLive", zap.String("error", err.Error()))
				}
				break
			case <-quitC:
				return
			}
		}
	}()
	reader := bufio.NewReaderSize(t.conn, util.MaxMessageSize)
	for {
		packet, err := t.ReadPacket(reader)
		if err != nil {
			g.L.Warn("Init:t.ReadPacket", zap.Error(err))
			return err
		}
		g.L.Info("cmd", zap.Any("cmd", packet))
		// 发给上层处理
		gAgent.downloadC <- packet
	}
}

// KeepLive ...
func (t *TcpClient) KeepLive() error {
	// ping := util.NewCMD()
	// ping.Type = util.TypeOfPing

	// p := util.NewAPMPacket()
	// p.Cmd = []*util.CMD{ping}
	// if err := t.WritePacket(p, util.TypeOfCompressNo); err != nil {
	// 	g.L.Warn("KeepLive:t.WritePacket", zap.String("error", err.Error()))
	// 	return err
	// }
	return nil
}

// ReadPacket ...
func (t *TcpClient) ReadPacket(rdr io.Reader) (*util.VgoPacket, error) {
	packet := util.NewVgoPacket()
	if err := packet.Decode(rdr); err != nil {
		g.L.Warn("ReadPacket:packet.Decode", zap.String("error", err.Error()))
		return nil, err
	}
	return packet, nil
}

// WritePacket ...
func (t *TcpClient) WritePacket(packet *util.VgoPacket) error {
	body := packet.Encode()
	if t.conn != nil {
		_, err := t.conn.Write(body)
		if err != nil {
			g.L.Warn("WritePacket:t.conn.Write", zap.String("error", err.Error()))
			return err
		}
	}
	return nil
}

// Close ....
func (t *TcpClient) Close() error {
	if t.conn != nil {
		t.conn.Close()
	}
	return nil
}
