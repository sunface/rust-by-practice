package dotray

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

// Node is the local server
type Node struct {
	nodeAddr string

	// the source seed addr
	sourceAddr string

	// the current seed
	seedAddr string
	seedConn net.Conn

	// ping、pong status with the current seed
	// when ping failed a few times, we need to connect to another seed node
	pinged bool

	// backup seeds list
	// when failed to connect to source seed,will try backup seed
	seedBackup []*Seed

	// the nodes which use our node as the current seed
	downstreams map[string]net.Conn

	// outer application channels
	send chan interface{}
	recv chan interface{}
}

// Seed structure
type Seed struct {
	addr  string
	retry int
}

/*
laddr: our node's listen addr
saddr: the source seed addr
send: outer application pushes messages to this channel
recv: outer application receives messages from this channel
*/
func StartNode(laddr, saddr string, send, recv chan interface{}) error {
	if laddr == "" {
		return errors.New("please use -l to specify our node's listen addr")
	}

	node = &Node{
		nodeAddr:    laddr,
		sourceAddr:  saddr,
		downstreams: make(map[string]net.Conn),
		send:        send,
		recv:        recv,
	}

	// start tcp listening
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		return err
	}

	// wait for downstream nodes to connect
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("accept downstream node error：", err)
				continue
			}

			go node.receiveFrom(conn, false, false)
		}
	}()

	// receive outer application's message,and route to the seed node and the downstream nodes
	go localSend(node)

	// resend the unsent messages(these messages didn't receive a matched ack from target node)
	go resend(node)

	// the main logic of seed manage
	if saddr != "" {
		// start to ping with the seed
		//when ping failed a few times, we need to connect to another seed node
		go node.ping()

		err := node.connectSeed(saddr)
		if err != nil {
			fmt.Printf("failed to connecto the seed%s：%v\n", saddr, err)
			return err
		}

		// start to sync the backup seed from current seed
		//the backup seeds are those nodes who directly connected with the current seed
		go node.syncBackupSeed()

		// start to receive messages from the current seed
		node.receiveFrom(node.seedConn, true, false)

	SourceSeedTrye:
		// although we disconnected with the source seed
		// but,here we want retry source seed for a few times(n) first
		n := 0
		for {
			if n > seedMaxRetry {
				break
			}

			err := node.connectSeed(saddr)
			if err != nil {
				n++
				goto CONTINUE
			}
			node.receiveFrom(node.seedConn, true, false)
			//when successfully connected, the counter will be reset to 0
			n = 0
		CONTINUE:
			time.Sleep(3 * time.Second)
		}

		// after retry several times with the source seed,now we want connect with our backup seeds
		for {
			if len(node.seedBackup) <= 0 {
				// if there is no backup seed,we will go back to the source seed
				fmt.Printf("no backup seed exist now\n")
				break
			}

			// here is one important thing to notice
			//if stepBack is setted to 'true', we will go back to source seed retrys again

			// why?
			//because, at times, the big cluster will divided into few smaller clusters, the smaller
			// ones will not perceive each other, so we need a way to combine smaller ones to a
			// larger one, this is why we will go back to retry the source seed after some time.

			//and this stepBack action only happend when we has connected to backup seeds
			stepBack := node.connectBackSeeds()
			if stepBack {
				fmt.Println("step back to the source seed")
				goto SourceSeedTrye
			}
		}

		// go back to try source seed
		goto SourceSeedTrye
	}

	select {}
}

// receive messages from remote node
func (node *Node) receiveFrom(conn net.Conn, isSeed bool, needStepBack bool) bool {
	var addr string
	// close the connection
	defer func() {
		conn.Close()
		// if the node is in downstream, then remove
		if addr != "" {
			fmt.Printf("remote downstream node %s close the connection\n", addr)
			node.delete(addr)
		}

		// if the node is the seed, then reset
		if isSeed {
			fmt.Printf("remote seed node %s close the connection\n", node.seedAddr)
			node.seedConn = nil
			node.seedAddr = ""
		}
	}()

	// the step back has been mentioned above
	start := time.Now().Unix()
	for {
		if needStepBack {
			now := time.Now().Unix()
			// A connection which connected to backup seed ,will maintain no more than 240 second
			if now-start > maxBackupSeedAlive {
				return true
			}
		}

		decoder := gob.NewDecoder(conn)
		r := &Request{}
		err := decoder.Decode(r)
		if err != nil {
			if err != io.EOF {
				fmt.Println("decode message error：", err)
			}
			break
		}
		a, err := r.handle(node, conn)
		if err != nil {
			fmt.Println("handle message error：", err)
			break
		}
		// update the downstream node's listen addr
		if a != "" {
			addr = a
		}
	}

	return false
}

// delete node from downstream
func (node *Node) delete(addr string) {
	lock.Lock()
	delete(node.downstreams, addr)
	lock.Unlock()
}

// dial to remote node
func (node *Node) dialToNode(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// connect to the seed node
func (node *Node) connectSeed(addr string) error {
	conn, err := node.dialToNode(addr)
	if err != nil {
		return err
	}

	node.seedConn = conn
	node.seedAddr = addr
	fmt.Printf("connect to the seed %s successfully\n", addr)
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
		if n >= maxPingAllowed {
			// when the ping failed several times, we will choose another seed to connnect
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
		time.Sleep(pingInterval * time.Second)
	}
}

// sync backup seed from current seed
// the backup seeds are those nodes who directly connected with the current seed
func (node *Node) syncBackupSeed() {
	// waiting for node's initialization
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
			fmt.Println("a critical error happens when connecto to backup seed", err)
		}
	}()

	for i, seed := range node.seedBackup {
		exist := false
		var err error
		var stepBack bool

		// a node can't  appear in seedBackup and downstream at the same time
		for addr := range node.downstreams {
			if addr == seed.addr {
				node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
				exist = true
			}
		}
		if exist {
			fmt.Printf("a conflict between backupSeeds and downstream,so the backup seed is deleted：%s\n", seed.addr)
			continue
		}

		// seed connection retries can't exceed the upper limit
		if seed.retry > seedMaxRetry {
			fmt.Printf("seed %sretries exceed the limit\n", seed.addr)
			node.seedBackup = append(node.seedBackup[:i], node.seedBackup[i+1:]...)
			goto CONTINUE1
		}
		err = node.connectSeed(seed.addr)
		if err != nil {
			seed.retry++
			fmt.Printf("reconnect to seed %v error: %v\n", seed, err)
			goto CONTINUE1
		}

		stepBack = node.receiveFrom(node.seedConn, true, true)
		// go back to source seed
		if stepBack {
			return true
		}
		// if a seed was successfully connected, the retry counter will be reset to 0
		seed.retry = 0
	CONTINUE1:
		time.Sleep(3 * time.Second)
	}

	return false
}
