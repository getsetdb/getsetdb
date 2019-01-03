// functions and variables that
// are used to interact with
// specified database for RUD
// proper validation for this
// is extremely important for
// these commands as other
// clients can issue erroneous
// directory to a database
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// slice of all commands
// for input validation as
// well as for documentation
// purposes
var databaseCommands = []string{
	"get",   // gets value from a key
	"set",   // sets value to a key
	"del",   // deleted a pair
	"all",   // lists all pairs
	"info",  // returns size and path of database
	"count", // returns number of pairs in database
}

// databaseExecutor for validating
// and executing commands for databases
// existing on the running system
func databaseExecutor(command string) (string, error) {

	// slice that will store
	// names of database which
	// will be extracted from
	// path /tmp/gsdb/ containing
	// the database files for it
	// by removing the ".gsdb"
	// extension from the end of
	// the file name
	var databases []string
	databaseName := extractFirstTerm(command)

	// read directory for
	// listing all files
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
	if !stringInSlice(databaseName, databases) {
		return "", errors.New(databaseName + " does not exist")
	}

	// extract database command
	// by getting all strings in
	// a slice after the first term
	// which will be the database
	// specifier that has been used
	// for existence validation
	databaseCommand := extractCommandFromDatabaseCommand(command)

	// if no strings were provided
	// then the databaseCommand slice
	// will have a length of 0 raising
	// a command type error
	if len(databaseCommand) == 0 {
		return "", errors.New("command not specified for database `" + databaseName + "`")
	}

	// return commandError if database
	// command is not found right after
	// the database name specifier
	if !stringInSlice(databaseCommand[0], databaseCommands) {
		return "", commandError(databaseCommand[0])
	}

	switch databaseCommand[0] {
		case databaseCommands[0]: // `get`
			if len(databaseCommand) < 2 {
				return "", errors.New("key value not specified for database `" + databaseName + "`")
			}
			return dGet(databaseName, databaseCommand[1])
		case databaseCommands[1]: // `set`
			if len(databaseCommand) < 3 {
				if len(databaseCommand) < 2 {
					return "", errors.New("key value not specified for database `" + databaseName + "`")
				}
				return "", errors.New("value for key `" + databaseCommand[1] + "` not specified for database `" + databaseName + "`")
			}
			return dSet(databaseName, databaseCommand[1], strings.Join(databaseCommand[2:], " "))
		case databaseCommands[2]: // `del`
			if len(databaseCommand) < 2 {
				return "", errors.New("value to be deleted not specified for database `" + databaseName + "`")
			}
			return dDel(databaseName, databaseCommand[1])
		case databaseCommands[3]: // `all`
			return dAll(databaseName)
		case databaseCommands[4]: // `info`
			return dInfo(databaseName)
		case databaseCommands[5]: // `count`
			return dCount(databaseName)
		default:
			return "", errors.New("command `" + databaseCommand[0] + "` not recognized")
	}

}

/************** COMMANDS **************/

// initialises Parrington to
// read database file and
// extract value from key provided
func dGet(databaseName string, key string) (string, error) {

	if _, hasKey := pairs[databaseName + "_" + key]; hasKey {
		return pairs[databaseName + "_" + key], nil
	}

	p := Parrington{databasePath: path(databaseName)}
	p.writeToBody()
	p.writeToPairs()

	// gets value from key
	value, err := p.getValue(key)

	if err == nil {
		pairs[databaseName + "_" + key] = value
	}

	return value, err

}

// sets a value to
// the key provided
func dSet(databaseName string, key string, value string) (string, error) {
	file, err := os.OpenFile(path(databaseName), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)

	defer file.Close()

	pair := fmt.Sprintf("%s : %s", key, value)

	if _, err = file.WriteString(pair + "\n"); err != nil {
		return "", err
	}

	return pair, nil
}

// deletes pair
func dDel(databaseName, key string) (string, error) {
	p := Parrington{databasePath: path(databaseName)}
	p.writeToBody()
	p.writeToPairs()

	return p.delPair(key)
}

// returns all keys
// existing in database
func dAll(databaseName string) (string, error) {
	p := Parrington{databasePath: path(databaseName)}
	p.writeToBody()
	p.writeToPairs()

	return p.getKeys()
}

// gets info for database file
// in the getsetdb running
func dInfo(databaseName string) (string, error) {

	// open and read file
	// info into `file`
	file, err := os.Stat(path(databaseName))

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("size : %d bytes\npath : %s", file.Size(), path(databaseName)), nil

}

// returns total number of lines
// in database - 1 to give number
// of pairs in the line
func dCount(databaseName string) (string, error) {
	r, _ := os.Open(path(databaseName))
	return "count : " + strconv.Itoa(lineCounter(r)-1), nil
}

// defunct version for getting
// number of total pairs in database
// since it'll take too long to read
// to file and parser pairs to `p.paris`
func _dCount(databaseName string) (string, error) {
	p := Parrington{databasePath: path(databaseName)}
	p.writeToBody()
	p.writeToPairs()

	return strconv.Itoa(len(p.pairs)), nil
}
