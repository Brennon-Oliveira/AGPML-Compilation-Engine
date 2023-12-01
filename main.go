package main

import (
	"agpml/parser"
	"agpml/visitors"
	"github.com/antlr4-go/antlr/v4"
)

func main() {

	filePath := "./test.agpml"
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewAgpmlLexer(input)

	//token := lexer.NextToken()
	//for token.GetTokenType() != antlr.TokenEOF {
	//	fmt.Printf("<%s, %s>\n", token.GetText(), lexer.SymbolicNames[token.GetTokenType()])
	//	token = lexer.NextToken()
	//}

	commonTokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parserInstance := parser.NewAgpmlParser(commonTokenStream)
	parserInstance.BuildParseTrees = true
	tree := parserInstance.Program()
	visitors.Visit(tree)
}
