// functions and variables that
package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

// slice of all commands
// for input validation as
// well as for documentation
// purposes
var databaseCommands = []string{
	"get",
	"set",
	"del",
	"help",
	"info",
}

func databaseExecutor(command string) (string, error) {

	var databases []string
	database := extractFirstTerm(command)

	files, err := ioutil.ReadDir("/tmp/gsdb/")
	check(err)

	// loop over database files and
	// store them without their
	// extension in `databases` slice
	for _, f := range files {
		databases = append(databases, removeExtension(f.Name()))
	}

	// check if database entered exists
	// in the slice of database names
	if !stringInSlice(database, databases) {
		return "", errors.New(database + " does not exist")
	}

	databaseCommand := extractCommandFromDatabaseCommand(command)

	if len(databaseCommand) == 0 {
		return "", errors.New("command not specified for database `" + database + "`")
	}

	// return commandError if database
	// command is not found right after
	// the database name specifier
	if !stringInSlice(databaseCommand[0], databaseCommands) {
		return "", commandError(databaseCommand[0])
	}

	return strings.Join(databaseCommand, " "), nil

}
