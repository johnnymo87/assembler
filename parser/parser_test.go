package parser_test

import (
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IO", func() {
	var lines = ReadLines("MaxL.asm")
	It("ReadLines() parses 16 commands from MaxL.asm", func() {
		Expect(len(lines)).To(Equal(16))
	})
	It("ReadLines() only A & C commands are collected", func() {
		expectation := []string{"A_Command", "C_Command"}
		for _, command := range lines {
			comm, _ := command.Type()
			Expect(expectation).To(ContainElement(comm))
		}
	})
	var bLines = []string{"00", "01", "10", "11"}
	It("WriteLines()", func() {
		err := WriteLines(bLines, "test/dummy_file")
		Expect(err).NotTo(HaveOccurred())
	})
})

var _ = Describe("Type()", func() {
	It("A_Command", func() {
		c, _ := Command("@sum").Type()
		Expect(c).To(Equal("A_Command"))
	})
	It("C_Command", func() {
		c1, _ := Command("D=D-M").Type()
		c2, _ := Command("D;JGT").Type()
		Expect(c1).To(Equal("C_Command"))
		Expect(c2).To(Equal("C_Command"))
	})
	It("L_Command", func() {
		c, _ := Command("(LOOP)").Type()
		Expect(c).To(Equal("L_Command"))
	})
	It("error", func() {
		_, err := Command("lulu").Type()
		Expect(err).To(HaveOccurred())
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
