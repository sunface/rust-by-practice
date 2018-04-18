package dotray

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"time"
)

type Node struct {
	// 节点地址
	nodeAddr string

	// 当前种子
	seedAddr string
	seedConn net.Conn

	// 备用种子列表
	seedBackup []string

	// 下流节点列表
	downstreams map[string]net.Conn

	send chan interface{}
	recv chan interface{}
}

/*
laddr: 本地节点监听地址
saddr: 种子节点地址
send: 本地节点发送数据通道
recv: 本地节点接收数据通道
*/
func StartNode(laddr, saddr string, send, recv chan interface{}) error {
	// 若不传IP，则默认为本地Node
	if laddr == "" {
		return errors.New("请通过laddr参数输入监听地址")
	}

	node := &Node{
		nodeAddr:    laddr,
		downstreams: make(map[string]net.Conn),
		send:        send,
		recv:        recv,
	}

	// 开始监听TCP端口
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		return err
	}

	// 监听下游节点建立的连接
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("建立客户端连接错误：", err)
				continue
			}

			go node.receiveFrom(conn, false)
		}
	}()

	// 种子节点管理和连接
	if saddr != "" {
		err := node.connectSeed(saddr)
		if err != nil {
			fmt.Printf("连接种子节点失败%s：%v\n", saddr, err)
			return err
		}

		// 种子节点连接断开，重新连接备用种子列表，直到成功
		for {
			for i, addr := range node.seedBackup {
				if addr != "" {
					err = node.connectSeed(addr)
					if err != nil {
						fmt.Printf("重新连接种子节点失败%s：%v\n", addr, err)
					}
					// 种子节点断开后，认为该种子已经失效，从备用种子列表删除
					if i == len(node.seedBackup)-1 {
						node.seedBackup = node.seedBackup[:i]
					} else {
						node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
					}
				}
				time.Sleep(100 * time.Millisecond)
			}
		}

	} else {
		select {}
	}
}

// 接收其它节点发来的数据
func (node *Node) receiveFrom(conn net.Conn, isSeed bool) {
	var addr string
	// 处理连接关闭
	defer func() {
		conn.Close()
		// 若断开的是种子节点，则需要重新连接到一个种子节点
		if isSeed {
			// todo,重新连接
			return
		}

		// 若是下游节点，则从下游列表中移除
		if addr != "" {
			node.delete(addr)
		}
	}()

	for {
		decoder := gob.NewDecoder(conn)
		r := &Request{}
		err := decoder.Decode(r)
		if err != nil {
			fmt.Println("接收其它节点数据错误：", err)
			break
		}
		a, err := r.handle(node, conn)
		if err != nil {
			fmt.Println("处理其它节点消息错误：", err)
			break
		}
		// 获取conn对应的下游节点的监听地址
		if a != "" {
			addr = a
		}
	}

}

func (node *Node) delete(addr string) {
	lock.Lock()
	delete(node.downstreams, addr)
	lock.Unlock()
}

// 跟远程节点建立连接
func (node *Node) dialToNode(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 通知seed节点，请求获取备份种子列表
func (node *Node) syncBackupSeed() {
	r := &Request{
		Command: SyncBackupSeeds,
		Data:    node.nodeAddr,
	}
	e := gob.NewEncoder(node.seedConn)
	e.Encode(r)
}

func (node *Node) ping(conn net.Conn) {
	r := &Request{
		Command: ServerPing,
		Data:    node.nodeAddr,
	}
	e := gob.NewEncoder(conn)
	e.Encode(r)
}

func (node *Node) connectSeed(addr string) error {
	conn, err := node.dialToNode(addr)
	if err != nil {
		return err
	}
	node.seedConn = conn
	node.seedAddr = addr
	// ping 种子节点
	node.ping(node.seedConn)

	// 请求获取备份种子列表
	node.syncBackupSeed()

	node.receiveFrom(node.seedConn, true)
	return nil
}
