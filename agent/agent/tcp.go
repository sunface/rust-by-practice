package agent

import (
	"bufio"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/util"
	"github.com/vmihailenco/msgpack"
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
	//var conn net.Conn
	var err error
	isRestart := true
	quitC := make(chan bool, 1)
	// 定时器
	tc := time.NewTicker(time.Duration(misc.Conf.Agent.KeepliveInterval) * time.Second)

	defer func() {
		if err := recover(); err != nil {
			g.L.Warn("server panic", zap.Stack("server"), zap.Any("err", err))
		}
		// 是否重启
		if isRestart {
			t.Init()
		}
		return
	}()

	defer func() {
		close(quitC)
		t.conn.Close()
		tc.Stop()
	}()

	// connect alert
	for {
		t.conn, err = net.Dial("tcp", misc.Conf.Agent.VgoAddr)
		if err != nil {
			g.L.Warn("Connect vgo", zap.String("err", err.Error()), zap.String("addr", misc.Conf.Agent.VgoAddr))
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	// 启动心跳
	go func() {
		for ; ;  {
			select {
			case <-tc.C:
				t.KeepLive()
			break
			case <-quitC:
				return
			}
		}
	}()
	reader := bufio.NewReaderSize(t.conn, util.MaxMessageSize)
	for {
		cmdPacket, err := t.ReadPacket(reader)
		if err != nil {
			g.L.Warn("ReadPacket", zap.Error(err))
			return err
		}
		g.L.Info("cmd", zap.Any("cmd", cmdPacket))
		// 发给上层处理
		gAgent.cmdC <- cmdPacket
	}

	return nil
}

// KeepLive ...
func (t *TcpClient) KeepLive() error{
	var packet util.BatchAPMPacket
	ping:=util.NewCMD()
	ping.Type = util.TypeOfPing

	p := util.NewAPMPacket()
	p.Cmds = []*util.CMD{ping}

	payload, err:=msgpack.Marshal(p)
	if err!=nil {
		g.L.Warn("KeepLive", zap.String("error", err.Error()))
		return err
	}

	packet.IsCompress = util.TypeOfCompressNo
	packet.PayLoad = payload
	body :=packet.Encode()
	if t.conn != nil {
		_, err:= t.conn.Write(body)
		if err!=nil {
			g.L.Warn("KeepLive", zap.String("error", err.Error()))
			return err
		}
	}
	return nil
}

// ReadPacket ...
func (t *TcpClient) ReadPacket(rdr io.Reader) (*util.CMD, error) {
	cmd := util.NewCMD()
	if err := cmd.Decode(rdr); err != nil {
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
