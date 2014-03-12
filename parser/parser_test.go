package parser_test

import (
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("Only A & C commands are collected", func() {
		lines := Parse("MaxL.asm")
		expectation := []string{"A_Command", "C_Command"}
		for _, command := range lines {
			comm, _ := command.Type()
			Expect(expectation).To(ContainElement(comm))
		}
	})
})
