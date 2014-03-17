package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var table = map[string]int{
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

//func openScanner(filename string) *bufio.Scanner {
//	path = extractPath(filename) + ".hack"
//	file, err := os.Open(filename)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	return bufio.NewScanner(file)
//}

func ReadLines(filename string) []Command {
	path = extractPath(filename) + ".hack"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []Command
	scanner := bufio.NewScanner(file)
	counter := 0
	addr := 16
	for scanner.Scan() {
		command := Command(scanner.Text())
		//fmt.Println(string(command))
		typ, err := command.Type()
		if err == nil {
			if typ == "L_Command" {
				sym, _ := command.Symbol()
				table[sym] = counter
			} else {
				counter += 1
			}
			if typ == "A_Command" {
				txt, _ := command.Symbol()
				value, present := table[txt]
				if present == true {
					newtxt := "@" + strconv.Itoa(value)
					lines = append(lines, Command(newtxt))
				} else {
					newtxt := "@" + strconv.Itoa(addr)
					lines = append(lines, Command(newtxt))
					table[txt] = addr
					addr += 1
				}
			} else if typ == "C_Command" {
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

var symbol = regexp.MustCompile(`^@(\S+)`)    // @sum
var label = regexp.MustCompile(`^\((\S+)\)$`) // (LOOP)
var dest = regexp.MustCompile(`^(\S+)=`)      // D=D-M
var destcomp = regexp.MustCompile(`=(\S+)`)   // D=D-M
var compjump = regexp.MustCompile(`(\S+);`)   // D;JGT
var jump = regexp.MustCompile(`;(\S+)`)       // D;JGT

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
