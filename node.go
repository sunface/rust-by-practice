package dotray

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"sync"
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

type Request struct {
	Command int
	Data    interface{}
}

const (
	NormalRequest  = 1 // 正常请求，发送数据报文
	ServersRequest = 2 // 请求服务器列表
	ServerResponse = 3 // 返回服务器列表
	ServerPing     = 4 // 发送自己的监听地址
)

var lock = &sync.Mutex{}

func StartNode(laddr, saddr string, send, recv chan interface{}) error {
	// 若不传IP，则默认为本地Node
	if laddr == "" {
		return errors.New("请通过laddr参数输入监听地址")
	}

	node := &Node{
		nodeAddr:    laddr,
		seedAddr:    saddr,
		downstreams: make(map[string]net.Conn),
		send:        send,
		recv:        recv,
	}

	// 连接种子节点
	if saddr != "" {
		conn, err := node.dialToNode(saddr)
		if err != nil {
			return err
		}
		node.seedConn = conn
		node.save(saddr, node.seedConn)

		// 定期同步服务器列表
		node.sync()

		go node.receiveFrom(node.seedConn)
	}

	// 开始监听TCP端口
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		return err
	}

	// 监听客户端建立的连接
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("建立客户端连接错误：", err)
			continue
		}

		go node.receiveFrom(conn)
	}
}

// 在本地缓存所有建立连接的节点,key是远程节点的地址
func (node *Node) save(addr string, conn net.Conn) {
	lock.Lock()
	node.conns[addr] = conn
	lock.Unlock()
}

func (node *Node) delete(conn net.Conn) {
	lock.Lock()
	delete(node.conns, conn.RemoteAddr().String())
	lock.Unlock()
}

// 接收其它节点发来的数据
func (node *Node) receiveFrom(conn net.Conn) {
	// 处理连接关闭
	defer func() {
		conn.Close()

	}()

	for {
		decoder := gob.NewDecoder(conn)
		r := Request{}
		err := decoder.Decode(r)
		if err != nil {
			fmt.Println("接收其它节点数据错误：", err)
			break
		}

		switch r.Command {
		case ServersRequest:
			// 返回节点列表
			addrs := make([]string, len(node.conns))
			encoder := gob.NewEncoder(conn)
			resp := Request{
				Command: ServerResponse,
				Data:    addrs,
			}
			encoder.Encode(resp)
		case ServerResponse:
			// 更新本地节点列表
			remoteAddrs := r.Data.([]string)
			// 判断远程节点地址在本地是否已经存在，存在则跳过，不存在则建立连接
			for _, addr := range remoteAddrs {
				exist := false
				for addr1 := range node.conns {
					if addr1 == addr {
						exist = true
						break
					}
				}

				if !exist {
					c, err := node.dialToNode(addr)
					if err != nil {
						fmt.Println("收到服务器列表后，建立连接错误：", err)
						continue
					}

				}
			}
		}
	}

}

// 跟远程节点建立连接
func (node *Node) dialToNode(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// 通知seed节点，请求同步服务器列表
func (node *Node) sync() {
	r := &Request{
		Command: ServersRequest,
	}
	e := gob.NewEncoder(node.seedConn)
	e.Encode(r)
}
