package dotray

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

// 待ACK数据包
type Packet struct {
	Addr   string
	retrys int
}

var sendPackets = make(map[int64][]*Packet)
var sendDatas = make(map[int64]Request)

func deleteResend(rid int64, from string) {
	lock.Lock()
	ps, ok := sendPackets[rid]
	lock.Unlock()
	if !ok {
		return
	}

	for i, p := range ps {
		if p.Addr == from {
			ps = append(ps[:i], ps[i+1:]...)
			break
		}
	}
	if len(ps) != 0 {
		lock.Lock()
		sendPackets[rid] = ps
		lock.Unlock()
		return
	}

	lock.Lock()
	delete(sendPackets, rid)
	delete(sendDatas, rid)
	lock.Unlock()
}

// 定时重发队列消息
// 大于10秒的重发，大于60秒的删除
func resend(node *Node) {
	for {
		now := time.Now().Unix()
		lock.Lock()
		for rid, ps := range sendPackets {
			if now-(rid/1e9) > 120 {
				delete(sendPackets, rid)
				delete(sendDatas, rid)
				continue
			}
			fmt.Println("old ps:", ps)
			if now-(rid/1e9) > 20 {
				r, ok := sendDatas[rid]
				if ok {
					for i, p := range ps {
						conn := getConnByAddr(p.Addr, node)
						if conn == nil {
							// 连接不存在，删除对应的包
							ps = append(ps[:i], ps[i+1:]...)
							continue
						}
						encoder := gob.NewEncoder(conn)
						err := encoder.Encode(r)
						if err != nil {
							ps = append(ps[:i], ps[i+1:]...)
							continue
						}
						fmt.Println("here1111: ", p.Addr)
					}
				}
			}

			fmt.Println("new ps:", ps)
			if len(ps) == 0 {
				delete(sendPackets, rid)
				delete(sendDatas, rid)
			} else {
				sendPackets[rid] = ps
			}
		}
		lock.Unlock()

		time.Sleep(10 * time.Second)
	}
}

func getConnByAddr(addr string, node *Node) net.Conn {
	if addr == node.seedAddr {
		return node.seedConn
	}

	conn, ok := node.downstreams[addr]
	if ok {
		return conn
	}

	return nil
}
