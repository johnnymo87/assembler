package parser_test

import (
	. "github.com/johnnymo87/assembler/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"bufio"
)

var _ = Describe("Packages", func() {
	It("ioutil.Readfile", func() {
		file, _ := os.Open("dummy_file.txt")
		defer file.Close()
		var lines []string
		scanner := bufio.NewScanner(file)
		
		Expect(os.Open("dummy_file.txt")).To(Equal

	})
	It("strings.Split", func() {
		content, _ := ioutil.ReadFile("dummy_file.txt")
		Expect(strings.Fields(string(content))).To(Equal([]string{"0", "a"}))
	})

})
var _ = Describe("Parser", func() {
	It("NewParser should read file lines into a slice", func() {
		p := NewParser("dummy_file.txt")
		Expect(p.Lines).To(Equal([]string{"0 1", "a"}))
	})
})
