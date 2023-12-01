// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package visitors // Agpml

import (
	"agpml/header_configs"
	"agpml/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"strings"
)

type BaseAgpmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func SemanticError(line int, message string) {
	fmt.Printf("SemanticError on line %d: %s\n", line, message)
	println(message)
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
	return nil
}

func (v *BaseAgpmlVisitor) VisitHeaderConfigs(ctx *parser.HeaderConfigsContext) interface{} {
	println("VisitHeaderConfigs")
	for _, child := range ctx.GetChildren() {
		v.VisitHeaderConfig(child.(*parser.HeaderConfigContext))
	}
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitHeaderConfig(ctx *parser.HeaderConfigContext) interface{} {
	println("VisitHeaderConfig")

	attribute := strings.TrimSpace(ctx.ATTRIBUTE().GetText())
	if !header_configs.IsAcceptedHeaderConfig(attribute) {
		SemanticError(ctx.GetStart().GetLine(), "Invalid header config: "+attribute)
		os.Exit(0)
	}

	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitVarConfigs(ctx *parser.VarConfigsContext) interface{} {
	println("VisitVarConfigs")
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitVarConfig(ctx *parser.VarConfigContext) interface{} {
	println("VisitVarConfig")
	return v.VisitChildren(ctx)
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
