// every command passes
// through either the
// databaseExecutor or
// the spaceExecutor
package main

import (
	"net"
)

// main executor through which all spaceCommands
// passed through the tcp server passes though
func executor(command string, c *net.Conn) (string, error) {

	// first string of command
	// entered with a delemiter
	// of space to check whether
	// command issues is a space
	// or a database command
	firstTerm := extractFirstTerm(command)

	// if firstTerm matches a
	// spaceCommand then the
	// spaceExecutor is called
	if stringInSlice(firstTerm, spaceCommands) {
		return spaceExecutor(command, *&c)
	}

	// else the databaseExecutor
	// command is called since
	// the first term would've
	// been the name of the database
	return databaseExecutor(command)

}
