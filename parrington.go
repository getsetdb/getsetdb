package main

import "io/ioutil"

type Parrington struct {
	body string
	databasePath string
}

func (p *Parrington) readToBody() {
	data, err := ioutil.ReadFile(p.databasePath)
	check(err)

	p.body = string(data)
}
