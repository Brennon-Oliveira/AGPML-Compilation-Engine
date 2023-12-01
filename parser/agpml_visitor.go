// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Agpml

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by AgpmlParser.
type AgpmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AgpmlParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by AgpmlParser#headerConfigs.
	VisitHeaderConfigs(ctx *HeaderConfigsContext) interface{}

	// Visit a parse tree produced by AgpmlParser#headerConfig.
	VisitHeaderConfig(ctx *HeaderConfigContext) interface{}

	// Visit a parse tree produced by AgpmlParser#varConfigs.
	VisitVarConfigs(ctx *VarConfigsContext) interface{}

	// Visit a parse tree produced by AgpmlParser#varConfig.
	VisitVarConfig(ctx *VarConfigContext) interface{}

	// Visit a parse tree produced by AgpmlParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by AgpmlParser#idGroup.
	VisitIdGroup(ctx *IdGroupContext) interface{}

	// Visit a parse tree produced by AgpmlParser#classGroup.
	VisitClassGroup(ctx *ClassGroupContext) interface{}

	// Visit a parse tree produced by AgpmlParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by AgpmlParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by AgpmlParser#style.
	VisitStyle(ctx *StyleContext) interface{}

	// Visit a parse tree produced by AgpmlParser#styleConfig.
	VisitStyleConfig(ctx *StyleConfigContext) interface{}
}
