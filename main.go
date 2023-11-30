package main

import (
	"agpml/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

func main() {

	filePath := "./test.agpml"
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewAgpmlLexer(input)

	t := lexer.NextToken()
	for t.GetTokenType() != antlr.TokenEOF {
		fmt.Printf("<%s, %s>\n", t.GetText(), lexer.SymbolicNames[t.GetTokenType()])
		t = lexer.NextToken()
	}

}
