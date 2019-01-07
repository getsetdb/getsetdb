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

// returns value from key provided
// directory from the Parrington.paris
// by performing a linear search
func (p Parrington) getValue(key string) (string, error) {
	for _, line := range p.pairs {
		piece := splitString(line, " ")

		if piece[0] == key {
			returnValue := strings.Join(piece[2:], " ")
			return dataInferer(returnValue) + " : " + returnValue, nil
		}
	}
	return "", errors.New("no value found for key `" + key + "`")
}

// replaces line where
// where key is found
// with a backspace '\b'
func (p Parrington) delPair(key string) (string, error) {

	input, err := ioutil.ReadFile(p.databasePath)
	check(err)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if splitString(line, " ")[0] == key {

			lines[i] = lines[len(lines)-1]
			lines = lines[:len(lines)-1]

			output := strings.Join(lines, "\n")

			err = ioutil.WriteFile(p.databasePath, []byte(output), 0644)
			check(err)

			return "deleted : " + key, nil
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

// returns all keys
// existing in database
// in a slice
func (p Parrington) getKeysSlice() ([]string, error) {

	var keys []string

	for _, line := range p.pairs {
		if len(line) == 0 {
			continue
		}
		piece := splitString(line, " ")
		keys = append(keys, piece[0])
	}

	if len(keys) != 0 {
		return keys, nil
	}

	return keys, errors.New("no keys in database")

}
