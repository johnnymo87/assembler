package parser

import (
	"bufio"
	"errors"
	"os"
	"regexp"
)

func Parse(filename string) []Command {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := Command(scanner.Text())
		typ, _ := command.Type()
		switch typ {
		case "A_Command":
			lines = append(lines, command)
		case "C_Command":
			lines = append(lines, command)
		case "L_Command":
			break
		default:
			break
		}
	}
	return lines
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
