package parser_test

import (
	"bufio"
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Parser", func() {
	BeforeEach(func() {
		scanner, _ := NewScanner("dummy_file.txt")
		var lines []*Command
		for scanner.Scan() {
			lines = append(lines, NewCommand(scanner.Text()))
		}
	})
	It("Parsed Lines are collected in a *Command slice", func() {
		scanner := NewScanner("dummy_file.txt")
		Expect(p.Lines).To(Equal([]string{"0 1", "a"}))
	})
})
