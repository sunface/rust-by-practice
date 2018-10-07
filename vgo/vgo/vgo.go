package vgo

import (
	"bufio"
	"github.com/containous/traefik/log"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/util"
	"github.com/mafanr/vgo/vgo/misc"
	"go.uber.org/zap"
	"net"
	"time"
)

type Vgo struct {
}

func New() *Vgo {
	return &Vgo{}
}

func (v *Vgo) Start() error {
	if err := v.init(); err != nil {
		g.L.Fatal("Start", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (v *Vgo) init() error {
	// start web ser

	// start stats
	// init vgo
	v.acceptAgent()

	return nil
}

func (v *Vgo) acceptAgent() error {
	ln, err := net.Listen("tcp", misc.Conf.Vgo.ListenAddr)
	if err != nil {
		g.L.Fatal("init", zap.String("msg", err.Error()), zap.String("addr", misc.Conf.Vgo.ListenAddr))
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				g.L.Fatal("Accept", zap.String("msg", err.Error()), zap.String("addr", misc.Conf.Vgo.ListenAddr))
			}
			log.Println(conn.RemoteAddr())
			conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
			go v.agentWork(conn)
		}

	}()

	return nil
}

func (v *Vgo) agentWork(conn net.Conn) {
	quitC := make(chan bool, 1)

	msgC := make(chan *util.BatchAPMPacket, 1000)

	defer func() {
		if err := recover(); err != nil {
			g.L.Error("work", zap.Any("msg", err))
			return
		}
	}()

	defer func() {
		close(quitC)
		close(msgC)
	}()

	go v.agentRead(conn, msgC, quitC)

	for {
		select {
		case <-quitC:
			g.L.Info("Quit")
			return
		case msg, ok := <-msgC:
			if ok {
				log.Println(msg)
			}
		}
	}
}

func (v *Vgo) agentRead(conn net.Conn, msgC chan *util.BatchAPMPacket, quitC chan bool) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	defer func() {
		quitC <- true
	}()
	reader := bufio.NewReaderSize(conn, util.MaxMessageSize)
	for {
		msg := &util.BatchAPMPacket{}
		if err := msg.Decode(reader); err != nil {
			g.L.Warn("agentRead", zap.String("err", err.Error()))
			return
		}
		msgC <- msg
		// 设置超时时间
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	}
}

func (v *Vgo) Close() error {
	return nil
}

//
//package cluster
//
//import (
//"encoding/binary"
//"fmt"
//"io"
//)
//
//type Message interface {
//	Type() uint16
//	Decode([]byte) error
//	Encode() ([]byte, error)
//}
//
//// Packet ... node transfer packet
//type Packet struct {
//	TypeOf  uint16
//	Sizeof  uint32
//	PayLoad []byte
//}
//
//// Encode encode
//func (p *Packet) Encode() ([]byte, error) {
//	p.Sizeof = uint32(len(p.PayLoad))
//	buf := make([]byte, p.Sizeof+6)
//	binary.BigEndian.PutUint16(buf[0:2], p.TypeOf)
//	binary.BigEndian.PutUint32(buf[2:6], p.Sizeof)
//	if p.Sizeof > 0 {
//		copy(buf[6:], p.PayLoad)
//	}
//	return buf, nil
//}

//
//func (cluster *Cluster) nodeRead(conn net.Conn, messageC chan Message, quitC chan bool) {
//	defer func() {
//		if err := recover(); err != nil {
//			logger.Warn("server panic", zap.Stack("server"), zap.Any("err", err))
//			return
//		}
//	}()
//	defer func() {
//		quitC <- true
//	}()
//	reader := bufio.NewReaderSize(conn, MaxMessageSize)
//	for {
//		msg, err := DecodePacket(reader)
//		if err != nil {
//			logger.Warn("DecodePacket", zap.String("err", err.Error()), zap.Int("keepAlive", cluster.keepAlive))
//			return
//		}
//		messageC <- msg
//		// 设置超时时间
//		conn.SetReadDeadline(time.Now().Add(time.Duration(cluster.keepAlive) * time.Second))
//	}
//}

//
//// DecodePacket decodes the packet from the provided reader.
//func DecodePacket(rdr io.Reader) (Message, error) {
//	p := &Packet{}
//	messageType, sizeOf, err := p.Decode(rdr)
//	if err != nil {
//		return nil, err
//	}
//
//	// 数据长度异常
//	if int(sizeOf) > MaxMessageSize {
//		return nil, fmt.Errorf("Message size is too large")
//	}
//
//	buffer := make([]byte, sizeOf)
//	_, err = io.ReadFull(rdr, buffer)
//	if err != nil {
//		return nil, err
//	}
//
//	// Decode the body
//	var msg Message
//	switch messageType {
//	case TypeOfConnect:
//		msg, err = decodeConnect(buffer)
//	default:
//		return nil, fmt.Errorf("Invalid type packet with type %d", messageType)
//	}
//
//	return msg, nil
//}
//
//// Decode decode
//func (p *Packet) Decode(rdr io.Reader) (uint16, uint32, error) {
//	buf := make([]byte, 6)
//	if _, err := io.ReadFull(rdr, buf); err != nil {
//		return 0, 0, err
//	}
//	p.TypeOf = binary.BigEndian.Uint16(buf)
//	p.Sizeof = binary.BigEndian.Uint32(buf[2:6])
//	return p.TypeOf, p.Sizeof, nil
//}
