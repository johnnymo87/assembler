package main

import (
	. "github.com/johnnymo87/assembler/parser"
)

func main() {
	scanner, err := NewScanner("tests/fixtures/dummy_file.txt")
	if err != nil {
		panic(err)
	}
	var lines []*Command
	for scanner.Scan() {
		com := NewCommand(scanner.Text())
		typ, err := com.Type()
		switch typ {
		case "A_Command":
			lines = append(lines, NewCommand(scanner.Text()))
		case "C_Command":
			lines = append(lines, NewCommand(scanner.Text()))
		case "L_Command":
			break
		default:
			break
		}
	}
}
