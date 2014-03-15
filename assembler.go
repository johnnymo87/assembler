package main

import (
	"bufio"
	"fmt"
	. "github.com/johnnymo87/assembler/code"
	. "github.com/johnnymo87/assembler/parser"
	"os"
)

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func toBinary(lines []Command) []string {
	var bLines []string
	for _, command := range lines {
		typ, err := command.Type()
		if err != nil {
			panic(err)
		}
		switch typ {
		case "A_Command":
			symbol, err := command.Symbol()
			if err != nil {
				panic(err)
			}
			bin := fmt.Sprintf("%016b", symbol)
			bLines = append(bLines, bin)
		case "C_Command":
			bin, err := Encode(command)
			if err != nil {
				panic(err)
			}
			bLines = append(bLines, bin)
		default:
			panic(command)
		}
	}
	return bLines
}

func main() {
	lines := Parse("parser/MaxL.asm")
	fmt.Printf("%#v", lines)
}
