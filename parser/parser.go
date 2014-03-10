package parser

import (
	"bufio"
	"github.com/johnnymo87/assembler/commands"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Parser interface {
	HasMoreCommands() bool
	ReadCommand() *Command
}

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner(filename string) *Scanner {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()
	return &Parser{scanner: bufio.NewScanner(file)}
}

func (s *Scanner) HasMoreCommands() bool {
	return s.scanner.Scan()
}

func (s *Scanner) ReadCommand() *Command {
	return commands.NewCommand(s.scanner.Text())
}

//func NewParser(filename string) *Parser {
//	file, err := os.Open(filename)
//	check(err)
//	defer file.Close()
//	//var lines []string
//	//for scanner.Scan() {
//	//	lines = append(lines, scanner.Text())
//	//}
//	//check(scanner.Err())
//	return &Parser{scanner: bufio.NewScanner(file)}
//}

//func hasMoreCommands() bool {
