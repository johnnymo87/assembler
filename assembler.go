package main

import (
	"flag"
	. "github.com/johnnymo87/assembler/code"
	. "github.com/johnnymo87/assembler/parser"
)

func main() {
	filename := flag.String("filename", "", "a string *.asm")
	flag.Parse()
	lines := ReadLines(*filename)
	bLines := ToBinary(lines)
	err := WriteLines(bLines)
	if err != nil {
		panic(err)
	}
}
