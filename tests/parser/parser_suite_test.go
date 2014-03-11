package parser_test

import (
	"github.com/johnnymo87/assembler/parser/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParser(t *testing.T) {
	//fixtures.Initialize()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}
