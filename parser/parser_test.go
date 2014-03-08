package parser_test

import (
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("should read a file", func() {
		path := "dummy_file.txt"
		p := NewParser(path)
		Expect(p.Lines).To(Equal([]string{"0"}))
	})
})
