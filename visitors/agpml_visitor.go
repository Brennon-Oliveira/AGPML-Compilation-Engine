// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package visitors // Agpml

import (
	"agpml/parser"
	"github.com/antlr4-go/antlr/v4"
)

// A complete Visitor for a parse tree produced by AgpmlParser.
type AgpmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AgpmlParser#program.
	VisitProgram(ctx *parser.ProgramContext) interface{}

	// Visit a parse tree produced by AgpmlParser#header_configs.
	VisitHeaderConfigs(ctx *parser.HeaderConfigsContext) interface{}

	// Visit a parse tree produced by AgpmlParser#headerConfig.
	VisitHeaderConfig(ctx *parser.HeaderConfigContext) interface{}

	// Visit a parse tree produced by AgpmlParser#varConfigs.
	VisitVarConfigs(ctx *parser.VarConfigsContext) interface{}

	// Visit a parse tree produced by AgpmlParser#varConfig.
	VisitVarConfig(ctx *parser.VarConfigContext) interface{}

	// Visit a parse tree produced by AgpmlParser#body.
	VisitBody(ctx *parser.BodyContext) interface{}

	// Visit a parse tree produced by AgpmlParser#idGroup.
	VisitIdGroup(ctx *parser.IdGroupContext) interface{}

	// Visit a parse tree produced by AgpmlParser#classGroup.
	VisitClassGroup(ctx *parser.ClassGroupContext) interface{}

	// Visit a parse tree produced by AgpmlParser#element.
	VisitElement(ctx *parser.ElementContext) interface{}

	// Visit a parse tree produced by AgpmlParser#line.
	VisitLine(ctx *parser.LineContext) interface{}

	// Visit a parse tree produced by AgpmlParser#style.
	VisitStyle(ctx *parser.StyleContext) interface{}

	// Visit a parse tree produced by AgpmlParser#styleConfig.
	VisitStyleConfig(ctx *parser.StyleConfigContext) interface{}
}
