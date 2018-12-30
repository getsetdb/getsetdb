package main

import (
	"log"
	"net"
)

const port string = ":4998"

func main() {

	ln, err := net.Listen("tcp4", port)
	check(err)

	log.Printf("tcp server running on 127.0.0.1%s\n", port)

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		check(err)
		go serve(conn)
	}

}
