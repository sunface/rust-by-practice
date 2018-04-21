# dotray
### http://dotray.io
A modern peer to peer(P2P) network library over TCP written in Go, there is no third party deps in dotray,so it's extremely easy to use in your project.

We have implemented  intelligent node discovery、data routing with DAG like algorithm, avoid data redundancy and connection redundancy.

so it's very suitable for simple blockchain project and teaching case.

Notice! This project is still under heavy development, for now,you can use it in simple blockchain or demos.

## Features
- Pure Go and Easy using
- Based on TCP/TLS
- Abstract Transport: Non-IP network
- Encryption(todo)
- NAT Tranversal(todo)
- Peer Discovery(to be optimized)
- Resource Discovery(todo)
- Intelligent message routing(to be optimized)
- Cluster Partition avoidance(to be optimized)
- Messaging by Multicast etc  (todo)
- Detail docs and code comments(Todo docs)

## How to use
#### demo code
```go
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dotray/dotray"
)

var laddr = flag.String("l", "", "")
var saddr = flag.String("s", "", "")
var messaging = flag.Bool("m", false, "")
var id = flag.String("i", "", "")

func main() {
	flag.Parse()

	send := make(chan interface{}, 1)
        recv := make(chan interface{}, 1)
    // start the p2p node
	go func() {
		err := dotray.StartNode(*laddr, *saddr, send, recv)
		if err != nil {
			panic("node start panic:" + err.Error())
		}
	}()


	// wait 1 second for p2p node started
	time.Sleep(1 * time.Second)

	// query 10 nodes address from p2p network
	addrs := dotray.QueryNodes(10)

    // excute some actions with the nodes address,like downloads blockchain from these nodes
    // all depends on yourself
    fmt.Println("query nodes:", addrs)
    
    // send message to all the other nodes
	if *messaging {
		data := "hello-" + *id
		go func() {
			for {
				send <- data
				time.Sleep(5 * time.Second)
				fmt.Println("send message：", data)
			}
		}()
	}

    // receive message from other nodes
	for {
		select {
		case r := <-recv:
			res := r.(*dotray.Request)
			fmt.Printf("receive message: %v from other node: \"%s\" \n", res.Data, res.From)
		}
	}

}
```

#### terminal
##### open the first terminal
```
go run main.go -l localhost:2000
```

##### open the second terminal
```
go run main.go -l localhost:2001 -s localhost:2000 
``` 

##### the third
```
go run main.go -l localhost:2002 -s localhost:2001 -m -i 3
```

##### the fourth
```
go run main.go -l localhost:2003 -s localhost:2002
```

##### the fifth
```
go run main.go -l localhost:2004 -s localhost:2002 -m -i 5
```

##### some explanation
- `-l` means the node will listen on this address
- `-s` means the node will first connect to this seed address
- `-m` means the node will send message to the p2p network
- `-i 5` just give the `-m` message a little variety

You can start as many nodes as you like, and you can kill any node as you like,