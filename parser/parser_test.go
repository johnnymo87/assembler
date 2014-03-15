package parser_test

import (
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadLines()", func() {
	var lines = ReadLines("MaxL.asm")
	It("parses 16 commands from MaxL.asm", func() {
		Expect(len(lines)).To(Equal(16))
	})
	It("Only A & C commands are collected", func() {
		expectation := []string{"A_Command", "C_Command"}
		for _, command := range lines {
			comm, _ := command.Type()
			Expect(expectation).To(ContainElement(comm))
		}
	})
})

var _ = Describe("Command text parsing", func() {
	It("Symbol() parses symbols", func() {
		Expect(Command("@sum").Symbol()).To(Equal("sum"))
	})
	It("Symbol() parses labels", func() {
		Expect(Command("(LOOP)").Symbol()).To(Equal("LOOP"))
	})
	It("Symbol() returns errors", func() {
		_, err := Command("D=D-M").Symbol()
		Expect(err).To(HaveOccurred())
	})
	It("Dest() parses destinations", func() {
		Expect(Command("D=D-M").Dest()).To(Equal("D"))
	})
	It("Dest() can return 'null'", func() {
		Expect(Command("D;JGT").Dest()).To(Equal("null"))
	})
	It("Comp() parses computations when there's a destination", func() {
		Expect(Command("D=D-M").Comp()).To(Equal("D-M"))
	})
	It("Comp() parses computations when there's a jump", func() {
		Expect(Command("D;JGT").Comp()).To(Equal("D"))
	})
	It("Comp() returns errors", func() {
		_, err := Command("@sum").Comp()
		Expect(err).To(HaveOccurred())
	})
	It("Jump() parses jumps", func() {
		Expect(Command("D;JGT").Jump()).To(Equal("JGT"))
	})
	It("Jump() can return 'null'", func() {
		Expect(Command("D=D-M").Jump()).To(Equal("null"))
	})
})
