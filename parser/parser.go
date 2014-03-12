package parser

import (
	"bufio"
	"errors"
	"os"
	"regexp"
)

func Parse(filename string) []*Command {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []*Command
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		com := NewCommand(scanner.Text())
		typ, _ := com.Type()
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
	return lines
}

type Command struct {
	Text string
}

func NewCommand(text string) *Command {
	return &Command{Text: text}
}

var symbol = regexp.MustCompile(`^@(\S+)`)    // @sum
var label = regexp.MustCompile(`^\((\S+)\)$`) // (LOOP)
var dest = regexp.MustCompile(`^(\S+)=`)      // D=D-M
var destcomp = regexp.MustCompile(`=(\S+)`)   // D=D-M
var compjump = regexp.MustCompile(`(\S+);`)   // D;JGT
var jump = regexp.MustCompile(`;(\S+)`)       // D;JGT

func (c *Command) Type() (string, error) {
	switch {
	case symbol.MatchString(c.Text):
		return "A_Command", nil
	case destcomp.MatchString(c.Text) || compjump.MatchString(c.Text):
		return "C_Command", nil
	case label.MatchString(c.Text):
		return "L_Command", nil
	default:
		return "", errors.New("unrecognized command type")
	}
}

func (c *Command) Symbol() (string, error) {
	switch {
	case symbol.MatchString(c.Text):
		result := symbol.FindStringSubmatch(c.Text)
		return result[len(result)-1], nil
	case label.MatchString(c.Text):
		result := label.FindStringSubmatch(c.Text)
		return result[len(result)-1], nil
	default:
		return "", errors.New("command has no Symbol")
	}
}

func (c *Command) Dest() (string, error) {
	switch {
	case dest.MatchString(c.Text):
		result := dest.FindStringSubmatch(c.Text)
		return result[len(result)-1], nil
	default:
		return "", errors.New("command has no Dest")
	}

}

func (c *Command) Comp() (string, error) {
	switch {
	case destcomp.MatchString(c.Text):
		result := destcomp.FindStringSubmatch(c.Text)
		return result[len(result)-1], nil
	case compjump.MatchString(c.Text):
		result := compjump.FindStringSubmatch(c.Text)
		return result[len(result)-1], nil
	default:
		return "", errors.New("command has no Comp")
	}
}

func (c *Command) Jump() (string, error) {
	switch {
	case jump.MatchString(c.Text):
		result := jump.FindStringSubmatch(c.Text)
		return result[len(result)-1], nil
	default:
		return "", errors.New("command has no Jump")
	}
}
