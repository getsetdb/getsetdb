// interpreter for a REPL
// like feature for gsdb
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// no customisation allowed
// because `> ` is the perfect
// prompt for everything
const prompt string = "> "

// sends a command to
// the tcp server and returns
// with a response by it
func sendToServer(command string, c *net.Conn) string {
	conn := *c

	_, _ = fmt.Fprintf(conn, command+"\n")

	response := make([]byte, 1024)

	_, _ = conn.Read(response)

	return string(response)

}

func input(c *net.Conn) (string, error) {
	// initialise the reader
	reader := bufio.NewReader(os.Stdin)

	// display the prompt
	fmt.Print(prompt)

	// store command in `text`
	// right after the user hits
	// the return key - entering
	// (no pun intended) a new line
	text, err := reader.ReadString('\n')

	if err == io.EOF {
		fmt.Println("bye")
		x := *c
		_ = x.Close()
		os.Exit(0)
	}

	// for some reason, `text`
	// is returned with a \n
	// trailing at the end which
	// is why this hack of returning
	// the whole text upto the
	// second last line has to
	// be in place to cut the \n
	return text[:len(text)-1], nil

}

// single REPL function
// for the entire functionality
// of the full getsetdb
func clientShell() {
	c, err := net.Dial("tcp4", port)
	check(err)
	for {
		command, err := input(&c)

		if err != nil {
			log.Println(err)
		}

		switch command {
			case "exit":
				log.Println("bye")
				_ = c.Close()
				return
			default:
				fmt.Print(sendToServer(command, &c))
				continue
		}
	}
}
