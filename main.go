package main

import (
	"net"
)

func main() {

	ln, err := net.Listen("tcp4", ":4444")
	check(err)

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		check(err)
		go serve(conn)
	}

}
