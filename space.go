package main

import (
	"errors"
	"net"
	"os"
	"strings"
	"time"
)

// list of all usable spaceCommands
var spaceCommands = []string{
	"new",           // TODO creates a new database
	"drop",          // TODO deletes an existing database
	"list",          // TODO lists all existing databases
	"rename",        // TODO renames and existing database
	"commands", 	 // displays a list of spaceCommands
	"version",       // TODO displays GSDB's running version
	"datetime",      // TODO shows current date and time of database
}

func spaceExecutor(command string, c *net.Conn) (string, error) {

	firstTerm := extractFirstTerm(command)

	switch firstTerm {
		case spaceCommands[0]: // new
			return sNew(command, &*c)
		case spaceCommands[4]: // commands
			return sCommands()
		case spaceCommands[6]: // datetime
			return sDatetime()
		default:
			return "", commandError(command)
	}
}

/************** COMMANDS **************/

// command type function
// creates new database
// inside of the GSDB space
func sNew(command string, c *net.Conn) (string, error) {

	commandSlice := splitString(command, " ")

	if len(commandSlice) < 2 {
		return "", databaseNameError(commandSlice[0])
	}

	for _, databaseName := range commandSlice[1:] {
		x := *c
		_, _ = x.Write([]byte("creating database `" + databaseName + "`... "))
		_, err := os.Stat(path(databaseName))

		if os.IsExist(err) {
			return "", errors.New("database `" + databaseName + "` already exists")
		}

		file, err := os.Create(path(databaseName))
		check(err)

		_ = file.Close()
		_, _ = x.Write([]byte("done\n"))
	}

	if len(commandSlice) > 2 {
		return "created databases " + strings.Join(commandSlice[1:], ", "), nil
	}

	return "created database", nil

}

// command type function
// returns datetime
// of the GSDB database
func sDatetime() (string, error) {
	return time.Now().String(), nil
}

// command type function
// lists all the spaceCommands available for
// the GSDB version that's running on system
func sCommands() (string, error) {
	return strings.Join(spaceCommands, "\n"), nil
}
