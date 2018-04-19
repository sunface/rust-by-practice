package dotray

import (
	"encoding/gob"
	"time"
)

func localSend(node *Node) {
	for {
		select {
		case raw := <-node.send:
			now := time.Now().UnixNano()
			r := Request{
				ID:      now,
				Command: NormalRequest,
				Data:    raw,
				From:    node.nodeAddr,
			}

			lock.Lock()
			sendPackets[r.ID] = make([]*Packet, 0)
			sendDatas[r.ID] = r
			lock.Unlock()
			n := 0
			if node.seedAddr != "" {
				// 发送给种子
				encoder := gob.NewEncoder(node.seedConn)
				encoder.Encode(r)
				lock.Lock()
				sendPackets[r.ID] = append(sendPackets[r.ID], &Packet{
					Addr: node.seedAddr,
				})
				lock.Unlock()
				n++
			}

			// 发送给下游
			for addr, conn := range node.downstreams {
				encoder := gob.NewEncoder(conn)
				encoder.Encode(r)
				lock.Lock()
				sendPackets[r.ID] = append(sendPackets[r.ID], &Packet{
					Addr: addr,
				})
				lock.Unlock()
				n++
			}

			// 没有发送任何消息，删除待重发送列表
			if n == 0 {
				lock.Lock()
				delete(sendPackets, r.ID)
				delete(sendDatas, r.ID)
				lock.Unlock()
			}
		}
	}
}

func routeSend(node *Node, r *Request) {
	//若当前种子节点和来源节点一致，则忽略
	now := time.Now().UnixNano()
	newR := Request{
		ID:      now,
		Command: NormalRequest,
		Data:    r.Data,
		From:    node.nodeAddr,
	}

	lock.Lock()
	sendPackets[newR.ID] = make([]*Packet, 0)
	sendDatas[newR.ID] = newR
	lock.Unlock()

	n := 0
	if r.From != node.seedAddr && node.seedAddr != "" {
		encoder := gob.NewEncoder(node.seedConn)
		encoder.Encode(newR)
		lock.Lock()
		sendPackets[newR.ID] = append(sendPackets[newR.ID], &Packet{
			Addr: node.seedAddr,
		})
		lock.Unlock()
		n++
	}

	// 转发给下游节点
	for addr, conn := range node.downstreams {
		if r.From != addr && addr != "" {
			encoder := gob.NewEncoder(conn)
			encoder.Encode(newR)
			lock.Lock()
			sendPackets[newR.ID] = append(sendPackets[newR.ID], &Packet{
				Addr: addr,
			})
			lock.Unlock()
			n++
		}
	}

	// 没有发送任何消息，删除待重发送列表
	if n == 0 {
		lock.Lock()
		delete(sendPackets, newR.ID)
		delete(sendDatas, newR.ID)
		lock.Unlock()
	}

	// 发送给本地节点处理
	node.recv <- r
}
