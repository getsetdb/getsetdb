package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func serve(c net.Conn) {
	log.Println("established connection with", c.RemoteAddr().String())

	for {
		remoteData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			break
		}

		command := strings.TrimSpace(string(remoteData))

		if command == "bye" {
			break
		}

		databaseName := strings.Split(command, " ")[0]
		fmt.Println(databaseName)

		_, _ = c.Write([]byte("you entered " + databaseName + "\n"))

	}

	log.Println("terminated connection with", c.RemoteAddr().String())

	_ = c.Close()

}
