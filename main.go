package main

import (
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp4", ":4444")
	check(err)

	log.Println("tcp server running on 127.0.0.1:4444")

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		check(err)
		go serve(conn)
	}

}
