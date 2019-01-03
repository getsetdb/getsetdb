package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// list of all usable spaceCommands
var spaceCommands = []string{
	"new",           // creates new databases provided
	"del",           // deletes an existing database
	"list",          // lists all existing databases
	"rename",        // renames and existing database
	"commands", 	 // displays a list of spaceCommands
	"version",       // displays GSDB's running version
	"datetime",      // shows current date and time of database
}

func spaceExecutor(command string) (string, error) {

	firstTerm := extractFirstTerm(command)

	switch firstTerm {
		case spaceCommands[0]: // `new`
			return sNew(command)
		case spaceCommands[1]: // `del`
			return sDel(command)
		case spaceCommands[2]: // `list`
			return sList()
		case spaceCommands[3]: // `rename`
			return sRename(command)
		case spaceCommands[4]: // `commands`
			return sCommands()
		case spaceCommands[5]: // `version`
			return sVersion()
		case spaceCommands[6]: // `datetime`
			return sDatetime()
		default:
			return "", commandError(command)
	}
}

/************** COMMANDS **************/

// command type function
// creates new database
// inside of the GSDB space
func sNew(command string) (string, error) {

	commandSlice := splitString(command, " ")

	if len(commandSlice) < 2 {
		return "", databaseNameError(commandSlice[0])
	}

	for _, databaseName := range commandSlice[1:] {

		// check if database already
		// exists in the path for
		// gsdb databases - /tmp/gsdb/
		if _, err := os.Stat(path(databaseName)); !os.IsNotExist(err) {
			return "", errors.New("database `" + databaseName + "` already exists")
		}

		file, err := os.Create(path(databaseName))
		check(err)

		_ = file.Close()
	}

	if len(commandSlice) > 2 {
		return "created databases `" + strings.Join(commandSlice[1:], "`, `") + "`", nil
	}

	return "created database `" + commandSlice[1] + "`", nil

}

// command type function
// deleted existing database
// inside of the gsdb space
func sDel(command string) (string, error) {

	commandSlice := splitString(command, " ")

	if len(commandSlice) < 2 {
		return "", databaseNameError(commandSlice[0])
	}

	for _, databaseName := range commandSlice[1:] {

		// check if database exists
		// in path for gsdb space
		// - /tmp/gsdb/
		if _, err := os.Stat(path(databaseName)); !os.IsNotExist(err) {
			_ = os.Remove(path(databaseName))
		} else { // response with error if file doesn't exist
			return "", errors.New("database `" + databaseName + "` does not exist")
		}
	}

	if len(commandSlice) > 2 {
		return "deleted databases `" + strings.Join(commandSlice[1:], "`, `") + "`", nil
	}

	return "deleted database `" + commandSlice[1] + "`", nil

}

// command type function
// lists all database available
// on the gsdb path /tmp/gsdb/
func sList() (string, error) {
	databases, err := ioutil.ReadDir("/tmp/gsdb/")
	check(err)

	var files []string

	// loop over all database files
	// and append the string in format
	// "<num>. <databaseName>" to files
	for index, database := range databases {
		files = append(files, fmt.Sprintf("%d : %s", index+1, removeExtension(database.Name())))
	}

	// join array of list
	return strings.Join(files, "\n"), nil

}

// command type function
// renames existing database
func sRename(command string) (string, error) {

	commandSlice := splitString(command, " ")

	if len(commandSlice) < 3 {
		if len(commandSlice) < 2 {
			return "", errors.New(fmt.Sprintf("database name not provided"))
		}
		return "", errors.New(fmt.Sprintf("new name for database `%s` not provided", commandSlice[1]))
	}

	oldDatabaseName := commandSlice[1]
	newDatabaseName := commandSlice[2]

	// check if database exists
	// in path for gsdb space -
	// /tmp/gsdb/
	if _, err := os.Stat(path(oldDatabaseName)); !os.IsNotExist(err) { // if oldDatabaseName exists

		var databases []string

		// check for old
		// databases to be
		// named `newDatabaseName`
		files, err := ioutil.ReadDir("/tmp/gsdb/")
		check(err)

		// loops over all files
		// (databases) in the
		// gsdb path - /tmp/gsdb/
		// and append them to
		// `databases` without
		// their file extensions
		for _, f := range files {
			databases = append(databases, removeExtension(f.Name()))
		}

		// if `newDatabaseName` exists
		// on gsdb path on disk
		if stringInSlice(newDatabaseName, databases) {
			return "", errors.New(fmt.Sprintf("database `%s` already exists", newDatabaseName))
		}

		_ = os.Rename(path(oldDatabaseName), path(newDatabaseName))

		return fmt.Sprintf("renamed `%s` to `%s`", oldDatabaseName, newDatabaseName), nil
	} else { // if oldDatabaseName does not exist
		return "", errors.New("database `" + oldDatabaseName + "` does not exist")
	}

}

// command type function
// lists all the spaceCommands available for
// the GSDB version that's running on system
func sCommands() (string, error) {
	return strings.Join(spaceCommands, "\n"), nil
}

// command type function
// returns version of gsdb
// running on system
func sVersion() (string, error) {
	return version, nil
}

// command type function
// returns datetime
// of the GSDB database
func sDatetime() (string, error) {
	return time.Now().String(), nil
}
