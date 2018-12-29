package main

import (
	"bufio"
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

		if len(command) == 0 {
			continue
		}

		if command == "bye" || command == "exit" {
			_, _ = c.Write([]byte("bye\n"))
			break
		}

		response, err := executor(command, &c)

		if err != nil {
			_, _ = c.Write([]byte(err.Error()))
		}

		if len(response) == 0 {
			_, _ = c.Write([]byte("\n"))
			continue
		}

		_, _ = c.Write([]byte(response + "\n"))

	}

	log.Println("terminated connection with", c.RemoteAddr().String())

	_ = c.Close()

}
