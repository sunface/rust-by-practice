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
	fmt.Printf("节点准备启动,监听地址：%s，种子地址：%s\n", laddr, saddr)
	// 若不传IP，则默认为本地Node
	if laddr == "" {
		return errors.New("请通过laddr参数输入监听地址")
	}

	node := &Node{
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

	fmt.Printf("节点监听成功\n")
	// 监听下游节点建立的连接
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("建立客户端连接错误：", err)
				continue
			}
			fmt.Printf("接收到新的下游连接：%s\n", conn.RemoteAddr().String())
			go node.receiveFrom(conn, false)
		}
	}()

	// 接收本地节点发来的消息，转发给种子节点和下游节点
	go func() {
		for {
			select {
			case raw := <-node.send:
				r := Request{
					Command: NormalRequest,
					Data:    raw,
					From:    node.nodeAddr,
				}
				if node.seedAddr != "" {
					// 发送给种子
					encoder := gob.NewEncoder(node.seedConn)
					encoder.Encode(r)
					fmt.Printf("发送数据%v给种子%s\n", raw, node.seedAddr)
				}

				// 发送给下游
				for addr, conn := range node.downstreams {
					encoder := gob.NewEncoder(conn)
					encoder.Encode(r)
					fmt.Printf("发送数据%v给下游%s\n", raw, addr)
				}
			}
		}
	}()

	// 种子节点管理和连接
	if saddr != "" {
		fmt.Printf("开始连接原始种子:%s\n", saddr)
		err := node.connectSeed(saddr)
		if err != nil {
			fmt.Printf("连接种子节点失败%s：%v\n", saddr, err)
			return err
		}

		fmt.Printf("原始种子连接断开:%s\n", saddr)

	SourceSeedTrye:
		// 原始种子断开，先继续尝试连接原始种子，超过上限，则访问备份种子
		n := 0
		for {
			if n > seedMaxRetry {
				break
			}

			fmt.Printf("继续尝试连接原始种子:%s\n", saddr)
			err := node.connectSeed(saddr)
			if err != nil {
				fmt.Printf("重新连接原始种子节点失败%s：%v\n", saddr, err)
				// 重连次数+1
				n++
				goto CONTINUE
			}
			// 之前的连接成功，又断开，重连计数归0
			n = 0
		CONTINUE:
			time.Sleep(3 * time.Second)
		}

		// 原始种子尝试彻底失败，连接备用种子列表
		for {
			fmt.Printf("准备连接备份种子列表：%v\n", node.seedBackup)
			if len(node.seedBackup) <= 0 {
				fmt.Printf("备份种子列表已经为空%v\n", node.seedBackup)
				break
			}

			for i, seed := range node.seedBackup {
				// 备份节点不得在下游节点中存在，若存在则删除
				exist := false
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
				fmt.Printf("开始连接新的种子节点:%v\n", seed)
				err = node.connectSeed(seed.addr)
				if err != nil {
					// 种子重试+1
					seed.retry++
					fmt.Printf("重新连接种子节点失败%v：%v\n", seed, err)
					goto CONTINUE1
				}
				// 种子连接成功后再断开，重试归0
				seed.retry = 0
				fmt.Printf("新的种子节点连接失效:%v\n", seed)

			CONTINUE1:
				time.Sleep(3 * time.Second)
			}
		}

		goto SourceSeedTrye
	}

	// 创世节点
	select {}
}

// 接收其它节点发来的数据
func (node *Node) receiveFrom(conn net.Conn, isSeed bool) {
	var addr string
	// 处理连接关闭
	defer func() {
		conn.Close()
		// 若是下游节点，则从下游列表中移除
		if addr != "" {
			fmt.Printf("下游节点%s关闭，准备从列表%v中删除\n", addr, node.downstreams)
			node.delete(addr)
			fmt.Printf("下游节点%s已经从列表%v中删除\n", addr, node.downstreams)
		}
	}()

	for {
		decoder := gob.NewDecoder(conn)
		r := &Request{}
		err := decoder.Decode(r)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("远程节点%s关闭了链接\n", addr)
				break
			}
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

// 通知seed节点，定时请求获取备份种子列表
func (node *Node) syncBackupSeed() {
	time.Sleep(100 * time.Millisecond)
	go func() {
		for {
			r := &Request{
				Command: SyncBackupSeeds,
				Data:    node.nodeAddr,
			}
			e := gob.NewEncoder(node.seedConn)
			e.Encode(r)

			time.Sleep(syncBackupSeedInterval * time.Second)
		}
	}()
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
	fmt.Printf("种子节点连接成功，准备ping\n")
	// ping 种子节点
	node.ping(node.seedConn)
	fmt.Printf("种子节点ping成功，准备获取备份种子列表\n")
	// 请求获取备份种子列表
	go node.syncBackupSeed()

	node.receiveFrom(node.seedConn, true)
	return nil
}
