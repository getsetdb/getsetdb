package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// Parrington a GSDB parser written for
// internal usage inside the
// GSDB server for reading
// and deriving values from
// keys via methods for it
type Parrington struct {
	body         string
	pairs        []string
	databasePath string
}

// EXTREMELY IMPORTANT FUNCTION
// to be called every time a
// variable of type Parrington
// is declared - it reads the
// file of the body and then
// fills the body parameter
// of the struct by referencing
// the variable by via a pointer
func (p *Parrington) writeToBody() {

	// read file into `data`
	data, err := ioutil.ReadFile(p.databasePath)
	check(err)

	// store inside local
	// parameter of self
	p.body = string(data)
}

// returns all key value
// paris of database raw
func (p Parrington) readBody() string {
	return p.body
}

// EXTREMELY IMPORTANT FUNCTION
// fills the pairs[] parameter
// of the Parrington struct for
// easier access to reading lines
func (p *Parrington) writeToPairs() {
	p.pairs = splitString(p.readBody(), "\n")
}

// infers the datatype from
// string provided as a value
// of a key pair on a single line
func dataInferer(value string) string {

	if !charInString(value, " ") { // check if there's a space in the value
		return "List"
	} else if _, err := strconv.Atoi(value); err == nil { // check if the value can be converted to a number
		return "Number"
	} else { // return String as default value
		return "String"
	}

}

// returns value from key provided
// directory from the Parrington.paris
// by performing a linear search
func (p Parrington) getValue(key string) (string, error) {
	for _, line := range p.pairs {
		piece := splitString(line, " ")

		if piece[0] == key {
			return strings.Join(piece[2:], " "), nil
		}
	}
	return "", errors.New("no value found for key `" + key + "`")
}

// returns all keys
// existing in database
func (p Parrington) getKeys() (string, error) {

	var keys []string

	for index, line := range p.pairs {
		if len(line) == 0 {
			continue
		}
		piece := splitString(line, " ")
		keys = append(keys, strconv.Itoa(index+1)+" : "+piece[0])
	}

	if len(keys) != 0 {
		return strings.Join(keys, "\n"), nil
	}

	return "", errors.New("no keys in database")
}
