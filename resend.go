package dotray

import (
	"encoding/gob"
	"net"
	"time"
)

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

// periodically resend the messages
func resend(node *Node) {
	for {
		now := time.Now().Unix()
		lock.Lock()
		for rid, ps := range sendPackets {
			// if the message stays too long,we will delete it directly
			if now-(rid/1e9) > maxResendStayTime {
				delete(sendPackets, rid)
				delete(sendDatas, rid)
				continue
			}
			// the message must stays for some time to resend
			if now-(rid/1e9) > minResendStayTime {
				r, ok := sendDatas[rid]
				if ok {
					for i, p := range ps {
						conn := getConnByAddr(p.Addr, node)
						if conn == nil {
							// the conn is empty,delete the message
							ps = append(ps[:i], ps[i+1:]...)
							continue
						}
						encoder := gob.NewEncoder(conn)
						err := encoder.Encode(r)
						if err != nil {
							// the conn is broken, delete the message
							ps = append(ps[:i], ps[i+1:]...)
							continue
						}
					}
				}
			}

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
