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