// Code generated from .//agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // agpml

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by agpmlParser.
type agpmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by agpmlParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by agpmlParser#headerConfigs.
	VisitHeaderConfigs(ctx *HeaderConfigsContext) interface{}

	// Visit a parse tree produced by agpmlParser#varConfigs.
	VisitVarConfigs(ctx *VarConfigsContext) interface{}

	// Visit a parse tree produced by agpmlParser#dataConfigs.
	VisitDataConfigs(ctx *DataConfigsContext) interface{}
}
