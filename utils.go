// common utilities and functions
// used again and again in the code
package main

import (
	"errors"
	"path/filepath"
	"strings"
)

// commonly used error both
// for spaceCommands as well
// as for databaseCommands
func commandError(command string) error {
	return errors.New("command `" + command + "` not recognized")
}

// commonly used error
// for spaceCommands to
// inform for lack of
// database name in the
// command query
func databaseNameError(command string) error {
	return errors.New("database name not specified for command `" + command + "`")
}

func removeExtension(fileName string) string {
	return fileName[0:len(fileName) - len(filepath.Ext(fileName))]
}

// simply panic if
// error is found
// QUITS the server
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// function to split strings
// according to a delimiter
func splitString(command string, del string) []string {
	return strings.Split(command, del)
}

// a linear search function for
// iterating over a list of strings
// usually containing database names
// to compare with the database
// entered by the user to give
// commands on the specified database
func stringInSlice(command string, list []string) bool {
	for _, b := range list {
		if b == command {
			return true
		}
	}
	return false
}

// return full path of the datbase
// along with the file extension
func path(database string) string {
	return "/tmp/gsdb/" + database + ".gsdb"
}

// simply return the first string
// of the command string which
// would specify the database name
func extractFirstTerm(command string) string {
	return strings.Split(command, " ")[0]
}

// return array of strings of
// the command string splitted
// with delemiter of space except
// the first which is the database
// name on which the command is to
// be executed upon
func extractCommandFromDatabaseCommand(command string) []string {
	return strings.Split(command, " ")[1:]
}
