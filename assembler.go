package main

import (
	"github.com/johnnymo87/assembler/commands"
	"github.com/johnnymo87/assembler/parser"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	s := parser.NewScanner("tests/fixtures/dummy_file.txt")
	var p Parser
	p = s //equivalent to "p = Parser(s)"
	var lines []Command
	for p.HasCommands {
		lines = append(lines, ReadCommand())
	}
}
