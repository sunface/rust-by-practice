package dotray

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// Request 节点之间交换的数据结构
type Request struct {
	ID      int64
	Command int
	Data    interface{}
	From    string
}

const (
	NormalRequest         = 0 // 正常请求，发送数据报文
	NormalRequestReceived = 1 // 消息已经接收ack
	ServersRequest        = 2 // 请求服务器列表
	ServerResponse        = 3 // 返回服务器列表
	ServerPing            = 4 // 发送节点的监听地址
	ServerPong            = 5 // ping请求的ack
	BackupSeeds           = 6 // 备用种子节点列表
	SyncBackupSeeds       = 7 // 请求获取备用种子列表
)

func (r *Request) handle(node *Node, conn net.Conn) (string, error) {
	switch r.Command {
	case NormalRequestReceived:
		// 从待重发队列中，删除消息
		deleteResend(r.ID, r.From)
	case NormalRequest:
		// 接收到远程节点发来的消息
		//发送给本地节点处理
		//转发给种子节点和下游节点
		routeSend(node, r)

		// 回复消息收到的ack
		encoder := gob.NewEncoder(conn)
		encoder.Encode(Request{
			ID:      r.ID,
			Command: NormalRequestReceived,
			From:    node.nodeAddr,
		})
	case SyncBackupSeeds:
		// 请求节点的地址
		fromAddr := r.Data.(string)
		// 请求获取备用种子列表
		//从当前种子 + 下游列表中选出合适的种子节点(根据负载)
		// 要避免互为种子的情况
		var addrs []string
		if node.seedAddr != "" && node.seedAddr != fromAddr {
			addrs = append(addrs, node.seedAddr)
		}
		for addr := range node.downstreams {
			if len(addrs) < maxBackupSeedLen && addr != fromAddr {
				addrs = append(addrs, addr)
			}
		}

		encoder := gob.NewEncoder(conn)
		encoder.Encode(Request{
			Command: BackupSeeds,
			Data:    addrs,
		})
	case BackupSeeds:
		addrs := r.Data.([]string)

		// 当前种子节点发送备用种子节点
		for _, addr1 := range addrs {
			if addr1 == "" {
				continue
			}
			// 种子更新策略
			//1.备用节点数未达上限，添加到列表
			//2.达到上限，替换重试次数超过seedMaxRetry,且从最高的开始替换
			exist := false
			maxRetry := 0
			for _, seed := range node.seedBackup {
				if seed.retry > maxRetry {
					maxRetry = seed.retry
				}
				if addr1 == seed.addr {
					exist = true
					break
				}
			}
			// 新的种子在当前的备用列表中不存在
			if !exist {
				if len(node.seedBackup) >= maxBackupSeedLen {
					// 若备用列表中的种子最大重试次数没有超过阀值，则不替换
					if maxRetry <= seedMaxRetry {
						break
					}
					// 替换第一个超过阀值的旧种子
					for i, seed := range node.seedBackup {
						if seed.retry > seedMaxRetry {
							node.seedBackup[i] = &Seed{
								addr:  addr1,
								retry: 0,
							}
						}
					}
				} else {
					// 添加新种子
					node.seedBackup = append(node.seedBackup, &Seed{
						addr:  addr1,
						retry: 0,
					})
				}
			}

		}
		// 打印当前种子、备份种子、下游节点
		// fmt.Printf("当前种子：%s,备份种子：%v,下游节点：%v\n", node.seedAddr, getSeedAddrs(node.seedBackup), node.downstreams)
	case ServerPing:
		// 下游节点发送它的监听地址
		addr, ok := r.Data.(string)
		if ok {
			// 添加进本节点下游节点列表
			lock.Lock()
			node.downstreams[addr] = conn
			lock.Unlock()
			// 添加进下游节点后，目标节点不得在种子备份列表中存在
			for i, seed := range node.seedBackup {
				if seed.addr == addr {
					node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
					break
				}
			}

			// 返回Pong
			encoder := gob.NewEncoder(conn)
			encoder.Encode(Request{
				Command: ServerPong,
				From:    node.nodeAddr,
			})

			return addr, nil
		}

	case ServerPong:
		node.pinged = true
	default:
		fmt.Println("未识别的消息类型：", r.Command)
	}

	return "", nil
}

func getSeedAddrs(seeds []*Seed) []string {
	addrs := make([]string, len(seeds))
	for i, seed := range seeds {
		addrs[i] = seed.addr + "/" + strconv.Itoa(seed.retry)
	}

	return addrs
}
