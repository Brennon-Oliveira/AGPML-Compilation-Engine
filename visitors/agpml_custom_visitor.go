// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package visitors // Agpml

import (
	"agpml/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
)

type BaseAgpmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func SemanticError(line int, message string) {
	fmt.Printf("SemanticError on line %d: %s\n", line, message)
	os.Exit(0)
}

func Visit(ctx antlr.ParseTree) {
	visitor := &BaseAgpmlVisitor{}
	visitor.VisitProgram(ctx.(*parser.ProgramContext))
}

func (v *BaseAgpmlVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	println("VisitProgram")
	if ctx.HeaderConfigs() != nil {
		v.VisitHeaderConfigs(ctx.HeaderConfigs().(*parser.HeaderConfigsContext))
	}
	if ctx.VarConfigs() != nil {
		v.VisitVarConfigs(ctx.VarConfigs().(*parser.VarConfigsContext))
	}
	return nil
}

func (v *BaseAgpmlVisitor) VisitBody(ctx *parser.BodyContext) interface{} {
	println("VisitBody")
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitIdGroup(ctx *parser.IdGroupContext) interface{} {
	println("VisitIdGroup")
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitClassGroup(ctx *parser.ClassGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitElement(ctx *parser.ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitLine(ctx *parser.LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitStyle(ctx *parser.StyleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitStyleConfig(ctx *parser.StyleConfigContext) interface{} {
	return v.VisitChildren(ctx)
}
