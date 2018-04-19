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
	NormalRequest         = 0 // outer application's data
	NormalRequestReceived = 1 // ack
	ServerPing            = 2 // ping to the seed
	ServerPong            = 3 // pong to the ping
	BackupSeeds           = 4 // return the backup seeds
	SyncBackupSeeds       = 5 // query for the backup seeds
)

func (r *Request) handle(node *Node, conn net.Conn) (string, error) {
	switch r.Command {
	case NormalRequestReceived:
		// delete the message from resend queue
		deleteResend(r.ID, r.From)
	case NormalRequest:
		// route the received message to other nodes and outer application
		routeSend(node, r)

		// response with ack
		encoder := gob.NewEncoder(conn)
		encoder.Encode(Request{
			ID:      r.ID,
			Command: NormalRequestReceived,
			From:    node.nodeAddr,
		})
	case SyncBackupSeeds:
		// the address of the requester
		fromAddr := r.Data.(string)

		// filter the adjacent nodes from current seed and the downstream nodes
		// Avoid forming a dead loop, seed addr mustn't be equal to from addr
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

		for _, addr1 := range addrs {
			if addr1 == "" {
				continue
			}
			// the strategy of seeds update
			// if the upper limit of the seedBackup is not reached, we can append the new addr to the seedBackup
			// otherwise, we need to replace those nodes whose connection retries bigger than the maxRetry,
			// with the new seed
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

			if !exist {
				if len(node.seedBackup) >= maxBackupSeedLen {
					if maxRetry <= seedMaxRetry {
						break
					}
					for i, seed := range node.seedBackup {
						if seed.retry > seedMaxRetry {
							node.seedBackup[i] = &Seed{
								addr:  addr1,
								retry: 0,
							}
						}
					}
				} else {
					node.seedBackup = append(node.seedBackup, &Seed{
						addr:  addr1,
						retry: 0,
					})
				}
			}

		}

		fmt.Printf("source seed: %s,current seed：%s,backup seeds：%v,downsteam：%v\n", node.sourceAddr, node.seedAddr, getSeedAddrs(node.seedBackup), node.downstreams)
	case ServerPing:
		// a downstream node sends its address to its seed node(our node)
		addr, ok := r.Data.(string)
		if ok {
			// we need to add the downstream node
			lock.Lock()
			node.downstreams[addr] = conn
			lock.Unlock()
			// a node can't appear in downstream and seedBackup at the same time
			for i, seed := range node.seedBackup {
				if seed.addr == addr {
					node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
					break
				}
			}

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
		fmt.Println("unrecognized message type：", r.Command)
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
