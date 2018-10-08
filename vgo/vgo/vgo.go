package vgo

import (
	"bufio"
	"github.com/containous/traefik/log"
	"github.com/golang/snappy"
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/util"
	"github.com/mafanr/vgo/vgo/misc"
	"github.com/mafanr/vgo/vgo/stats"
	"github.com/shamaton/msgpack"
	"go.uber.org/zap"
	"net"
	"time"
)

type Vgo struct {
	stats *stats.Stats
}

func New() *Vgo {
	return &Vgo{
		stats: stats.New(),
	}
}

func (v *Vgo) Start() error {
	if err := v.init(); err != nil {
		g.L.Fatal("Start:v.init", zap.String("error", err.Error()))
		return err
	}
	return nil
}

func (v *Vgo) init() error {
	// start web ser

	// start stats
	if err := v.stats.Start(); err != nil {
		g.L.Warn("init:v.stats.Start", zap.String("error", err.Error()))
		return err
	}
	// init vgo
	v.acceptAgent()

	return nil
}

func (v *Vgo) acceptAgent() error {
	ln, err := net.Listen("tcp", misc.Conf.Vgo.ListenAddr)
	if err != nil {
		g.L.Fatal("acceptAgent:net.Listen", zap.String("msg", err.Error()), zap.String("addr", misc.Conf.Vgo.ListenAddr))
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				g.L.Fatal("acceptAgent:ln.Accept", zap.String("msg", err.Error()), zap.String("addr", misc.Conf.Vgo.ListenAddr))
			}
			conn.SetReadDeadline(time.Now().Add(time.Duration(misc.Conf.Vgo.AgentTimeout) * time.Second))
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
			g.L.Error("agentWork:.", zap.Any("msg", err))
			return
		}
	}()

	defer func() {
		close(quitC)
		close(msgC)
		conn.Close()
	}()

	go v.agentRead(conn, msgC, quitC)

	for {
		select {
		case <-quitC:
			g.L.Info("Quit")
			return
		case msg, ok := <-msgC:
			if ok {
				p := util.NewAPMPacket()
				var payload []byte
				var err error
				if msg.IsCompress == util.TypeOfCompressNo {
					payload = msg.PayLoad
				} else {
					payload, err = snappy.Decode(nil, msg.PayLoad)
					if err != nil {
						g.L.Warn("agentWork:snappy.Decode", zap.String("error", err.Error()))
						break
					}
				}
				if err := msgpack.Decode(payload, p); err != nil {
					g.L.Warn("agentWork:msgpack.Decode", zap.String("error", err.Error()))
					break
				}
				log.Println("接收到到报文", p)
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
			g.L.Warn("agentRead:msg.Decode", zap.String("err", err.Error()))
			return
		}
		msgC <- msg
		// 设置超时时间
		conn.SetReadDeadline(time.Now().Add(time.Duration(misc.Conf.Vgo.AgentTimeout) * time.Second))
	}
}

func (v *Vgo) Close() error {
	return nil
}
