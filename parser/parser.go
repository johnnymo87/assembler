package parser

import (
	"io/ioutil"
	"strings"
)

type Parser struct {
	Lines []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewParser(filename string) *Parser {
	content, err := ioutil.ReadFile(filename)
	check(err)
	return &Parser{Lines: strings.Split(string(content), "\n")}
}

//func hasMoreCommands() bool {
