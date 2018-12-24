package main

import (
	"errors"
	"strings"
)

// list of all usable commands
var commands = []string{
	"new", // TODO creates a new database
	"bye", // closes the connection to db
	"drop", // TODO deletes an existing database
	"list", // TODO lists all existing databases
	"rename", // TODO renames and existing database
	"commands", // displays a list of commands
	"version", // TODO displays GSDB's running version
	"datetime", // TODO shows current date and time of database
}

// main executor through which all commands
// passed through the tcp server passes though
func executor(command string) (string, error) {

	switch command {
		case commands[5]:
			return cCommands(), nil
		default:
			return "", errors.New("command not recognized")
	}
}

// command type function
// lists all the commands available for
// the GSDB version that's running on system
func cCommands() string {
	return strings.Join(commands, "\n")
}
