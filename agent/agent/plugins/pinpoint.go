package plugins

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/containous/traefik/log"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/agent/misc"
	"github.com/mafanr/vgo/util"
	"github.com/openapm/openapm/collector/ahbase"
	"github.com/openapm/openapm/collector/proto"
	"go.uber.org/zap"
	"io"
	"net"
)

const (
	APPLICATION_SEND           int16 = 1
	APPLICATION_TRACE_SEND     int16 = 2
	APPLICATION_TRACE_SEND_ACK int16 = 3

	APPLICATION_REQUEST  int16 = 5
	APPLICATION_RESPONSE int16 = 6

	APPLICATION_STREAM_CREATE         int16 = 10
	APPLICATION_STREAM_CREATE_SUCCESS int16 = 12
	APPLICATION_STREAM_CREATE_FAIL    int16 = 14

	APPLICATION_STREAM_CLOSE int16 = 15

	APPLICATION_STREAM_PING int16 = 17
	APPLICATION_STREAM_PONG int16 = 18

	APPLICATION_STREAM_RESPONSE int16 = 20

	CONTROL_CLIENT_CLOSE int16 = 100
	CONTROL_SERVER_CLOSE int16 = 110

	// control packet
	CONTROL_HANDSHAKE          int16 = 150
	CONTROL_HANDSHAKE_RESPONSE int16 = 151

	// keep stay because of performance in case of ping and pong. others removed.
	CONTROL_PING int16 = 200
	CONTROL_PONG int16 = 201

	CONTROL_MECURY_PING int16 = 202
	CONTROL_MECURY_PONG int16 = 203

	UNKNOWN int16 = 500

	PACKET_TYPE_SIZE int16 = 2
)

type Pinpoint struct {
	isRun   bool
	infoCon net.Conn
}

func NewPinpoint() *Pinpoint {
	return &Pinpoint{}
}

func (p *Pinpoint) Start() error {
	go p.agentInfo()
	return nil
}

func (p *Pinpoint) Close() error {
	return nil
}

// agentInfo (默认tcp 9994) agent配置信息、Api信息、String信息、Exception信息
func (p *Pinpoint) agentInfo() error {
	defer func() {
		if err := recover(); err != nil {
			g.L.Warn("agentInfo:.", zap.Stack("server"), zap.Any("err", err))
		}
		if p.infoCon != nil {
			p.infoCon.Close()
		}
		//p.infoCon.Close()
		p.isRun = false
	}()

	ln, err := net.Listen("tcp", misc.Conf.Pinpoint.AgentInfoAddr)
	if err != nil {
		g.L.Warn("agentInfo:net.Listen", zap.Stack("server"), zap.Any("err", err))
		return err
	}
	g.L.Info("agentInfo:net.Listen", zap.String("addr", ln.Addr().String()))

	defer ln.Close()

ReStart:
	// 防止socket泄漏
	if p.infoCon != nil {
		p.infoCon.Close()
	}
	p.infoCon, err = ln.Accept()
	if err != nil {
		if p.infoCon != nil {
			p.infoCon.Close()
		}
	}

	reader := bufio.NewReaderSize(p.infoCon, util.MaxMessageSize)
	for {
		packetType, body, err := p.tcpRead(reader)
		if err != nil {
			g.L.Warn("agentInfo:p.infoRead", zap.String("err", err.Error()))
			goto ReStart
		}
		log.Println(packetType, body, err)
	}

	return nil
}

func (p *Pinpoint) tcpRead(reader io.Reader) (int16, []byte, error) {
	buf := make([]byte, util.MaxMessageSize)
	if _, err := io.ReadFull(reader, buf[:2]); err != nil {
		g.L.Warn("infoRead:io.ReadFull", zap.String("err", err.Error()))
		return 0, nil, err
	}
	packetType := int16(binary.BigEndian.Uint16(buf[0:2]))

	log.Println("接收到的PacketType为", packetType)

	// 根据不同的报文读取后续内容
	switch packetType {
	case CONTROL_HANDSHAKE:
		body, err := read4RidAndLenAndBody(reader, buf, CONTROL_HANDSHAKE)
		return packetType, body, err
	default:
		g.L.Warn("infoRead:default", zap.Int16("packetType", packetType))
		return packetType, nil, fmt.Errorf("unknow type %d", packetType)
	}
}

func read4RidAndLenAndBody(reader io.Reader, b []byte, ty int16) ([]byte, error) {
	// 读取body长度
	if _, err := io.ReadFull(reader, b[2:10]); err != nil {
		g.L.Warn("read4RidAndLenAndBody:io.ReadFull", zap.String("err", err.Error()))
		return nil, err
	}

	length := int(binary.BigEndian.Uint32(b[6:10]))

	// 读取body
	if length >= 10240000 {
		return b[:10], nil
	}

	//
	if length+10 > len(b) {
		b = appendByteByLen(b, length+10)
	}

	if _, err := io.ReadFull(reader, b[10:length+10]); err != nil {
		g.L.Warn("read4RidAndLenAndBody:io.ReadFull", zap.String("err", err.Error()))
		return nil, err
	}

	if ty == CONTROL_HANDSHAKE {
		packageBuf, _, err := proto.ReadControlHandshakeBufBuffer(CONTROL_HANDSHAKE, b[:10+length])
		if err != nil {
			g.L.Warn("read4RidAndLenAndBody:proto.ReadControlHandshakeBufBuffer", zap.String("err", err.Error()))
			return nil, err
		}
		appLifeCycle := ahbase.NewAppLifeCycle()
		if err := json.Unmarshal(packageBuf.GetPayload(), appLifeCycle); err != nil {
			g.L.Warn("read4RidAndLenAndBody:json.Unmarshal", zap.String("err", err.Error()))
			return nil, err
		}
		log.Println("完美解决： get appName is", appLifeCycle.AppId, ", agentID is", appLifeCycle.AppName)
	}
	return b[:10+length], nil
}

//长度超过一定限度的时候用到
func appendByteByLen(b []byte, l int) []byte {
	num := 1
	lb := len(b)
	for {
		if lb+num*util.MaxMessageSize < l {
			num++
		} else {
			break
		}
	}
	s := make([][]byte, num+1)
	s[0] = b
	for index := 1; index <= num; index++ {
		s[index] = make([]byte, util.MaxMessageSize)
	}
	return bytes.Join(s, nil)
}

// readInfo ...
func (p *Pinpoint) readInfo() error {

	return nil
}

// jvmState (默认udp 9995）JVM内存使用情况，线程相关信息，30秒一个包，每个包为6条信息
func (p *Pinpoint) jvmState() error {
	return nil
}

// spanStream (默认Udp 9996)数据最多的,全链路Trace 报文
func (p *Pinpoint) spanStream() error {
	return nil
}
