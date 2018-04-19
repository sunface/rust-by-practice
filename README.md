# Dotray(http://dotray.io)
A modern peer to peer(P2P) network library over TCP written in Go, there is no third party deps in dotray,so it's extremely easy to use in your project.

We implemented  intelligent node discovery、data routing with DAG like algorithm, avoid data redundancy and connection redundancy.

so it's very suitable for simple blockchain project and teaching case.

## Features
- Based on TCP
- Pure Go
- Node discovery
- Intelligent data routing
- resend when failed
- ping and pong
- Cluster Partition avoidance（Todo)
- Nodes query for grpc or other actions(just query nodes info from p2p, the actions followed has noting to do with p2p).(Todo)
- Detail docs and code comments

## How to use
#### demo code
```go
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/uulesson/dotray"
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

    // send message to other node
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

    // receive message from other node
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