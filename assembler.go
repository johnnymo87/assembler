package main

import (
	. "github.com/johnnymo87/assembler/parser"
)

func main() {
	s := NewScanner("tests/fixtures/dummy_file.txt")
	var p Parser
	p = s //equivalent to "p = Parser(s)"
	var lines []Command
	for p.HasMoreCommands {
		lines = append(lines, ReadCommand())
	}
}
