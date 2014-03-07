package parser_spec

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/johnnymo87/assembler/lib/parser"
)

var _ = Describe("func read()", func() {
	It("should read a file", func() {
		path := 
