package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var pairs = map[string]string{}

// opens a new connection
// every time a client is
// connected to tcp port
// of 127.0.0.1:4998
func serve(c net.Conn) {
	// log a message of connection
	fmt.Print("\t")
	log.Println("established connection with", c.RemoteAddr().String())

	for {
		// get data from
		// client connection
		remoteData, err := bufio.NewReader(c).ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			_ = c.Close()
			break
		}

		// remove all unnecessary spaces
		command := strings.TrimSpace(string(remoteData))

		// if only the return
		// key was pressed as
		// input then no response
		// will be received by
		// the client connected
		if len(command) == 0 {
			continue
		}

		// close tcp connection
		// to the client
		if command == "bye" || command == "exit" {
			_, _ = c.Write([]byte("bye\n"))
			break
		}

		// get response as an
		// error response
		// or a success response
		response, err := executor(command)

		// if error was received
		// as a response then
		// an error type response
		// is issued to the client
		if err != nil {
			_, _ = c.Write([]byte("error : " + err.Error()))
		}

		// send back response
		// with successful execution
		_, _ = c.Write([]byte(response + "\n"))

	}

	fmt.Print("\t")
	log.Println("terminated connection with", c.RemoteAddr().String())

	_ = c.Close()

}
