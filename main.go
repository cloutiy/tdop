package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		buf, _ := ioutil.ReadFile(args[1])
		source := string(buf)

		l := tdopLexer(source)
		p := parser{lexer: l}

		stmts := p.statements()
		for _, stmt := range stmts {
			printAST(stmt, 0)
			fmt.Println()
		}
	}
}

func printAST(t *token, identation int) {
	fmt.Println()
	for i := 0; i < identation; i++ {
		fmt.Print(" ")
	}
	fmt.Print("(")
	fmt.Print(t.value)
	if len(t.children) > 0 {
		for _, c := range t.children {
			fmt.Print(" ")
			printAST(c, identation+4)
		}
	}
	fmt.Print(")")
}
