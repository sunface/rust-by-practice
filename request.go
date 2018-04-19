package dotray

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// Request 节点之间交换的数据结构
type Request struct {
	Command int
	Data    interface{}
	From    string
}

const (
	NormalRequest   = 1 // 正常请求，发送数据报文
	ServersRequest  = 2 // 请求服务器列表
	ServerResponse  = 3 // 返回服务器列表
	ServerPing      = 4 // 发送节点的监听地址
	BackupSeeds     = 5 // 备用种子节点列表
	SyncBackupSeeds = 6 // 请求获取备用种子列表
)

func (r *Request) handle(node *Node, conn net.Conn) (string, error) {
	switch r.Command {
	case NormalRequest:
		fmt.Printf("接收到其它节点%s的数据:%v,\n", r.From, r.Data)
		// 接收到远程节点发来的消息
		//发送给本地节点处理
		//转发给种子节点和下游节点

		//若当前种子节点和来源节点一致，则忽略
		newR := Request{
			Command: NormalRequest,
			Data:    r.Data,
			From:    node.nodeAddr,
		}

		if r.From != node.seedAddr && node.seedAddr != "" {
			encoder := gob.NewEncoder(node.seedConn)
			encoder.Encode(newR)
			fmt.Printf("发送数据%v给种子%s\n", r.Data, node.seedAddr)
		}

		// 转发给下游节点
		for addr, conn := range node.downstreams {
			if r.From != addr && addr != "" {
				encoder := gob.NewEncoder(conn)
				encoder.Encode(newR)
				fmt.Printf("发送数据%v给下游%s\n", r.Data, addr)
			}
		}
		// 发送给本地节点处理
		node.recv <- r
	case SyncBackupSeeds:
		// 请求节点的地址
		fromAddr := r.Data.(string)
		fmt.Printf("收到远程节点：%s获取备份种子列表的请求\n", fromAddr)
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

		fmt.Printf("处理完成远程节点：%s获取备份种子列表的请求，返回种子列表：%v\n", fromAddr, addrs)
		encoder := gob.NewEncoder(conn)
		encoder.Encode(Request{
			Command: BackupSeeds,
			Data:    addrs,
		})
	case BackupSeeds:
		addrs := r.Data.([]string)
		fmt.Printf("收到之前请求的备份种子列表:%v\n", addrs)
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
		fmt.Printf("当前种子：%s\n", node.seedAddr)
		fmt.Printf("最新的备份种子:%v\n", getSeedAddrs(node.seedBackup))
		fmt.Printf("下游节点：%v\n", node.downstreams)
	case ServerPing:
		// 下游节点发送它的监听地址
		addr, ok := r.Data.(string)
		fmt.Printf("收到下游节点%s的ping请求\n", addr)
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
			return addr, nil
		}
		fmt.Printf("最新的下游节点列表：%v\n", node.downstreams)
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
