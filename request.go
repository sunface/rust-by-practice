package dotray

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Request struct {
	Command int
	Data    interface{}
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
	case SyncBackupSeeds:
		// 请求节点的地址
		fromAddr := r.Data.(string)

		// 请求获取备用种子列表
		//从当前种子 + 下游列表中选出合适的种子节点(根据负载)
		var addrs []string
		if node.seedAddr != "" {
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
			// 若当前备用种子节点数量达到上限，则忽视
			if len(node.seedBackup) >= maxBackupSeedLen {
				break
			}

			// 若新的备用种子节点地址不存在，则添加到备用种子节点列表重
			exist := false
			for _, addr2 := range node.seedBackup {
				if addr1 == addr2 {
					exist = true
					break
				}
			}
			if !exist {
				node.seedBackup = append(node.seedBackup, addr1)
			}
		}
	case ServerPing:
		// 下游节点发送它的监听地址
		addr, ok := r.Data.(string)
		if ok {
			// 添加进本节点下游节点列表
			lock.Lock()
			node.downstreams[addr] = conn
			lock.Unlock()
			return addr, nil
		}
	default:
		fmt.Println("未识别的消息类型：", r.Command)
	}

	return "", nil
}
