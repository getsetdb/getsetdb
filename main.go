package main

import (
	"log"
	"net"
	"os"
)

const port string = ":4998"

func main() {

	if _, err := os.Stat("/tmp/gsdb/"); os.IsNotExist(err) {
		err = os.Mkdir("/tmp/gsdb/", 0755)
		check(err)
	}

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
