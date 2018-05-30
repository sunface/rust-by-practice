# dotray

Dotray is a tool for building distributed applications, similar to a network composed with dot and ray 

Dotray implements a [gossip protocol](https://en.wikipedia.org/wiki/Gossip_protocol)
that provide membership, unicast, and broadcast functionality
with [eventually-consistent semantics](https://en.wikipedia.org/wiki/Eventual_consistency).
In CAP terms, it is AP: highly-available and partition-tolerant.

Dotray works in a wide variety of network setups, including thru NAT and firewalls, and across clouds and datacenters.
It works in situations where there is only partial connectivity,
 i.e. data is transparently routed across multiple hops when there is no direct connection between peers.
It copes with partitions and partial network failure.
It can be easily bootstrapped, typically only requiring knowledge of a single existing peer in the dotray to join.
It has built-in shared-secret authentication and encryption.
It scales to on the order of 100 peers, and has no dependencies.

## Using

Dotray is currently distributed as a Rust package.


## Status
Still under developing.

## Feedback
Your feedback is always welcome!
