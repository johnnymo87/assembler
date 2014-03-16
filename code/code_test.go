package code_test

import (
	. "github.com/johnnymo87/assembler/code"
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binary conversion", func() {
	var a1 = Command("@16")
	var c1 = Command("D=D-M")
	var c2 = Command("D;JGT")
	var lines = []Command{a1, c1, c2}
	var bLines = []string{
		"0000000000010000",
		"1111010011010000",
		"1110001100000001",
	}

	It("A_Binary()", func() {
		Ω(A_Binary(a1)).Should(Equal(bLines[0]))
	})
	It("C_Binary()", func() {
		Ω(C_Binary(c1)).Should(Equal(bLines[1]))
		Ω(C_Binary(c2)).Should(Equal(bLines[2]))
	})
	It("ToBinary()", func() {
		Ω(ToBinary(lines)).Should(Equal(bLines))
	})
})
