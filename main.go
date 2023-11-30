package main

import (
	"agpml/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {

	filePath := "./test.agpml"
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewAgpmlLexer(input)

	commonTokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parserInstance := parser.NewAgpmlParser(commonTokenStream)
	parserInstance.Program()

}
