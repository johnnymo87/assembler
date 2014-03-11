package parser_test

import (
	"github.com/johnnymo87/assembler/parser/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func read()", func() {
	It("should read a file", func() {
		path := "fixtures/dummy_file.txt"
		p := parser.NewParser(path)
		p.Should(Equal(parser.Parser))
	})
})
