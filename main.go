package main

import (
	"log"
	"net"
	"os"
)

const port string = ":4998"
const version string = "0.9"

func main() {

	if _, err := os.Stat("/tmp/gsdb/"); os.IsNotExist(err) {
		err = os.Mkdir("/tmp/gsdb/", 0755)
		check(err)
	}

	if len(os.Args) != 1 {
		if os.Args[1][0] == 'c' {
			clientShell()
			os.Exit(0)
		}
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
