package parser_test

import (
	. "github.com/johnnymo87/go-assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var lines = Parse("MaxL.asm")
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
	It("parses symbols", func() {
		Expect(Command("@sum").Symbol()).To(Equal("sum"))
	})
	It("parses labels", func() {
		Expect(Command("(LOOP)").Symbol()).To(Equal("LOOP"))
	})
	It("parses destinations", func() {
		Expect(Command("D=D-M").Dest()).To(Equal("D"))
	})
	It("parses computations when there's a destination", func() {
		Expect(Command("D=D-M").Comp()).To(Equal("D-M"))
	})
	It("parses computations when there's a jump", func() {
		Expect(Command("D;JGT").Comp()).To(Equal("D"))
	})
	It("parses jumps", func() {
		Expect(Command("D;JGT").Jump()).To(Equal("JGT"))
	})
})
