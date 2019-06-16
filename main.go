package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	port string = ":4998"
	version string = "1.0"
)

func main() {

	if _, err := os.Stat("/tmp/gsdb/"); os.IsNotExist(err) {
		_ = os.Mkdir("/tmp/gsdb/", 0755)
	}

	fmt.Print(`
	   ___     _   ___      _   ___  ___ 
	  / __|___| |_/ __| ___| |_|   \| _ )
	 | (_ / -_)  _\__ \/ -_)  _| |) | _ \
	  \___\___|\__|___/\___|\__|___/|___/

	author  : mentix02
	version : 1.0.1

	`)

	if len(os.Args) != 1 {
		if os.Args[1][0] == 'c' {
			clientShell()
			os.Exit(0)
		} else if os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Println("getsetdb-" + version)
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
