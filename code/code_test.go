package code_test

import (
	. "github.com/johnnymo87/assembler/code"
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binary conversion", func() {
	It("A_Binary()", func() {
		a1 := Command("@16")
		Ω(A_Binary(a1)).Should(Equal("0000000000010000"))
	})
	It("C_Binary()", func() {
		c1 := Command("D=D-M")
		c2 := Command("D;JGT")
		Ω(C_Binary(c1)).Should(Equal("1111010011010000"))
		Ω(C_Binary(c2)).Should(Equal("1110001100000001"))
	})
})
