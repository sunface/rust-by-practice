package dotray

import (
	"encoding/gob"
	"time"
)

// the outer application send messages
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
				// send to the seed
				encoder := gob.NewEncoder(node.seedConn)
				encoder.Encode(r)
				lock.Lock()
				sendPackets[r.ID] = append(sendPackets[r.ID], &Packet{
					Addr: node.seedAddr,
				})
				lock.Unlock()
				n++
			}

			// send to the downstream
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

			// nothing happend, do some sweeping work.
			if n == 0 {
				lock.Lock()
				delete(sendPackets, r.ID)
				delete(sendDatas, r.ID)
				lock.Unlock()
			}
		}
	}
}

// receive remote node's messages, and we will route to other nodes and the outer application
func routeSend(node *Node, r *Request) {
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

	// nothing happend, do some sweeping work.
	if n == 0 {
		lock.Lock()
		delete(sendPackets, newR.ID)
		delete(sendDatas, newR.ID)
		lock.Unlock()
	}

	// send to the outer application
	node.recv <- r
}
