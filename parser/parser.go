package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var Table = map[string]int{
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
}

var path string

func extractPath(asm string) string {
	regex := regexp.MustCompile(`(\S+)\.asm`)
	if !regex.MatchString(asm) {
		fmt.Printf("file '%v' is not of type *.asm\n", asm)
		panic(asm)
	}
	path := regex.FindStringSubmatch(asm)
	return path[len(path)-1]
}

func FirstPass(scanner *bufio.Scanner) []Command {
	counter := 0
	var lines []Command
	for scanner.Scan() {
		safe := comment.ReplaceAllString(scanner.Text(), "")
		command := Command(safe)
		typ, err := command.Type()
		if err == nil {
			if typ == "L_Command" {
				sym, _ := command.Symbol()
				Table[sym] = counter
			} else {
				counter += 1
				lines = append(lines, command)
			}
		}
	}
	if counter == 0 {
		fmt.Printf("failed to read any lines from %v\n")
		panic(lines)
	}
	return lines
}

func SecondPass(lines []Command) []Command {
	addr := 16
	var newlines []Command
	for _, command := range lines {
		typ, _ := command.Type()
		switch typ {
		case "A_Command":
			if digit.MatchString(string(command)) == false { // it's not just a constant like @3
				sym, _ := command.Symbol()
				value, present := Table[sym] // is it in the symbol Table?
				if present == true {
					newsym := "@" + strconv.Itoa(value) // if so, input the value
					newlines = append(newlines, Command(newsym))
				} else {
					Table[sym] = addr // if not, add it to the Table
					newsym := "@" + strconv.Itoa(addr)
					addr += 1
					newlines = append(newlines, Command(newsym))
				}
			} else {
				newlines = append(newlines, command)
			}
		case "C_Command":
			newlines = append(newlines, command)
		default:
			fmt.Println("How did '%v' get in here?\n", command)
			panic(command)
		}
	}
	return newlines
}

func ReadLines(filename string) []Command {
	path = extractPath(filename) + ".hack"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//scanner := openScanner(filename)
	lines := FirstPass(scanner)
	return SecondPass(lines)
}

func WriteLines(lines []string) error {
	// ".hack" is len 5
	if len(path) <= 5 {
		fmt.Printf("path '%v' does not seem correct\n")
		panic(path)
	}
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

var comment = regexp.MustCompile(`//(.*)$`)
var digit = regexp.MustCompile(`@(\d+)`)     // @3
var symbol = regexp.MustCompile(`@(\S+)`)    // @sum
var label = regexp.MustCompile(`\((\S+)\)$`) // (LOOP)
var dest = regexp.MustCompile(`(\S+)=`)      // D=D-M
var destcomp = regexp.MustCompile(`=(\S+)`)  // D=D-M
var compjump = regexp.MustCompile(`(\S+);`)  // D;JGT
var jump = regexp.MustCompile(`;(\S+)`)      // D;JGT

type Command string

func (c Command) Type() (string, error) {
	command := string(c)
	switch {
	case symbol.MatchString(command):
		return "A_Command", nil
	case destcomp.MatchString(command) || compjump.MatchString(command):
		return "C_Command", nil
	case label.MatchString(command):
		return "L_Command", nil
	default:
		return "", errors.New("unrecognized command type")
	}
}

func (c Command) Symbol() (string, error) {
	command := string(c)
	switch {
	case symbol.MatchString(command):
		result := symbol.FindStringSubmatch(command)
		return result[len(result)-1], nil
	case label.MatchString(command):
		result := label.FindStringSubmatch(command)
		return result[len(result)-1], nil
	default:
		return "", errors.New("command has no Symbol")
	}
}

func (c Command) Dest() string {
	command := string(c)
	switch {
	case dest.MatchString(command):
		result := dest.FindStringSubmatch(command)
		return result[len(result)-1]
	default:
		return "null"
	}

}

func (c Command) Comp() (string, error) {
	command := string(c)
	switch {
	case destcomp.MatchString(command):
		result := destcomp.FindStringSubmatch(command)
		return result[len(result)-1], nil
	case compjump.MatchString(command):
		result := compjump.FindStringSubmatch(command)
		return result[len(result)-1], nil
	default:
		return "", errors.New("command has no Comp")
	}
}

func (c Command) Jump() string {
	command := string(c)
	switch {
	case jump.MatchString(command):
		result := jump.FindStringSubmatch(command)
		return result[len(result)-1]
	default:
		return "null"
	}
}
