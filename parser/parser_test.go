package parser_test

import (
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var lines = Parse("MaxL.asm")
	It("Only A & C commands are collected", func() {
		expectation := []string{"A_Command", "C_Command"}
		for _, command := range lines {
			comm, _ := command.Type()
			Expect(expectation).To(ContainElement(comm))
		}
	})
	It("parses symbols", func() {
		Expect(NewCommand("@sum").Symbol()).To(Equal("sum"))
	})
	It("parses labels", func() {
		Expect(NewCommand("(LOOP)").Symbol()).To(Equal("LOOP"))
	})
	It("parses destinations", func() {
		Expect(NewCommand("D=D-M").Dest()).To(Equal("D"))
	})
	It("parses computations when there's a destination", func() {
		Expect(NewCommand("D=D-M").Comp()).To(Equal("D-M"))
	})
	It("parses computations when there's a jump", func() {
		Expect(NewCommand("D;JGT").Comp()).To(Equal("D"))
	})
	It("parses jumps", func() {
		Expect(NewCommand("D;JGT").Jump()).To(Equal("JGT"))
	})
})
