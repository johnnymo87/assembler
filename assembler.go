package main

import (
	"fmt"
	. "github.com/johnnymo87/assembler/parser"
)

func main() {
	lines := Parse("parser/MaxL.asm")
	fmt.Printf("%#v", lines)
}
