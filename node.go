package dotray

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

// Node 代表了本地的服务器节点
type Node struct {
	// 节点地址
	nodeAddr string

	// 原始种子列表
	sourceAddr string

	// 当前种子
	seedAddr string
	seedConn net.Conn
	pinged   bool // ping失败一定次数后，认为种子节点失败，重新连接

	// 备用种子列表
	seedBackup []*Seed

	// 下流节点列表
	downstreams map[string]net.Conn

	send chan interface{}
	recv chan interface{}
}

// Seed 远程种子节点
type Seed struct {
	addr  string
	retry int
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

	node = &Node{
		nodeAddr:    laddr,
		sourceAddr:  saddr,
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

			go node.receiveFrom(conn, false, false)
		}
	}()

	// 接收本地节点发来的消息，转发给种子节点和下游节点
	go localSend(node)

	// 启动重新发送队列
	go resend(node)

	// 种子节点管理和连接
	if saddr != "" {
		// ping 种子节点
		go node.ping()

		err := node.connectSeed(saddr)
		if err != nil {
			fmt.Printf("连接种子节点失败%s：%v\n", saddr, err)
			return err
		}

		// 请求获取备份种子列表
		go node.syncBackupSeed()
		node.receiveFrom(node.seedConn, true, false)

	SourceSeedTrye:
		// 原始种子断开，先继续尝试连接原始种子，超过上限，则访问备份种子
		n := 0
		for {
			if n > seedMaxRetry {
				break
			}

			err := node.connectSeed(saddr)
			if err != nil {
				// 重连次数+1
				n++
				goto CONTINUE
			}
			node.receiveFrom(node.seedConn, true, false)
			// 之前的连接成功，又断开，重连计数归0
			n = 0
		CONTINUE:
			time.Sleep(3 * time.Second)
		}

		// 原始种子尝试彻底失败，连接备用种子列表
		for {
			if len(node.seedBackup) <= 0 {
				fmt.Printf("备份种子列表已经为空%v\n", node.seedBackup)
				break
			}
			stepBack := node.connectBackSeeds()
			if stepBack {
				fmt.Println("回到原始种子，重新连接")
				goto SourceSeedTrye
			}
		}

		goto SourceSeedTrye
	}

	// 创世节点
	select {}
}

// 接收其它节点发来的数据
func (node *Node) receiveFrom(conn net.Conn, isSeed bool, needStepBack bool) bool {
	var addr string
	// 处理连接关闭
	defer func() {
		conn.Close()
		// 若是下游节点，则从下游列表中移除
		if addr != "" {
			fmt.Printf("远程节点%s关闭了连接\n", addr)
			node.delete(addr)
		}

		if isSeed {
			fmt.Printf("远程种子节点%s关闭了连接\n", node.seedAddr)
			node.seedConn = nil
			node.seedAddr = ""
		}
	}()

	// 为了防止脑裂，当前连接持续一段时间后，需要重新连接原始种子进行尝试
	start := time.Now().Unix()
	for {
		if needStepBack {
			now := time.Now().Unix()
			if now-start > 60 {
				fmt.Println("当前连接持续时间过长，需要重新连接原始种子")
				return true
			}
		}

		decoder := gob.NewDecoder(conn)
		r := &Request{}
		err := decoder.Decode(r)
		if err != nil {
			if err != io.EOF {
				fmt.Println("接收其它节点数据错误：", err)
			}
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

	return false
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

func (node *Node) connectSeed(addr string) error {
	conn, err := node.dialToNode(addr)
	if err != nil {
		return err
	}

	node.seedConn = conn
	node.seedAddr = addr
	fmt.Printf("种子节点%s连接成功\n", addr)
	return nil
}

func (node *Node) ping() {
	n := 0
	for {
		if node.pinged {
			n = 0
			node.pinged = false
			continue
		}
		if n >= 6 {
			// 1分钟还ping不通，更换种子节点
			if node.seedConn != nil {
				node.seedConn.Close()
				node.seedConn = nil
				node.seedAddr = ""
			}
			n = 0
			continue
		}

		if node.seedConn != nil {
			r := &Request{
				Command: ServerPing,
				Data:    node.nodeAddr,
			}
			e := gob.NewEncoder(node.seedConn)
			e.Encode(r)
			n++
		}
		time.Sleep(15 * time.Second)
	}
}

// 通知seed节点，定时请求获取备份种子列表
func (node *Node) syncBackupSeed() {
	time.Sleep(100 * time.Millisecond)
	go func() {
		for {
			if node.seedConn != nil {
				r := &Request{
					Command: SyncBackupSeeds,
					Data:    node.nodeAddr,
				}
				e := gob.NewEncoder(node.seedConn)
				e.Encode(r)
			}
			time.Sleep(syncBackupSeedInterval * time.Second)
		}
	}()
}

func (node *Node) connectBackSeeds() bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("连接备份种子时，发生了错误：", err)
		}
	}()

	for i, seed := range node.seedBackup {
		// 备份节点不得在下游节点中存在，若存在则删除
		exist := false
		var err error
		var stepBack bool

		for addr := range node.downstreams {
			if addr == seed.addr {
				node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
				exist = true
			}
		}

		if exist {
			fmt.Printf("备份节点和下游节点冲突，删除备份节点：%s\n", seed.addr)
			continue
		}
		// 若种子的重试次数超过上限，则进行删除
		if seed.retry > seedMaxRetry {
			fmt.Printf("备份种子%s重连次数超过上限，从备份列表删除\n", seed.addr)
			node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
			goto CONTINUE1
		}
		err = node.connectSeed(seed.addr)
		if err != nil {
			// 种子重试+1
			seed.retry++
			fmt.Printf("重新连接种子节点失败%v：%v\n", seed, err)
			goto CONTINUE1
		}

		stepBack = node.receiveFrom(node.seedConn, true, true)
		// 返回原始种子重新尝试
		if stepBack {
			return true
		}
		// 种子连接成功后再断开，重试归0
		seed.retry = 0
	CONTINUE1:
		time.Sleep(3 * time.Second)
	}

	return false
}
