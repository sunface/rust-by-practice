package main

import (
	"log"
	"net"
)

func main() {
	var con net.Conn
	if con != nil {
		con.Close()
	} else {
		log.Println("con is nil")
	}

}
