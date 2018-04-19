# p2p
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
- Detail docs and code comments

## How to use
#### demo code
```go
package main
import (
    "github.com/uulesson/dotray"
)

func main() {

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
```go run main.go -l localhost:2004 -s localhost:2002 -m -i 5
````

##### some explanation
- `-l` means the node will listen on this address
- `-s` means the node will first connect to this seed address
- `-m` means the node will send message to the p2p network
- `-i 5` just give the `-m` message a little variety

You can start as many nodes as you like, and you can kill any node as you like,