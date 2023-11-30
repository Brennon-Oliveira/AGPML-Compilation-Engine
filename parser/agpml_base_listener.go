// Code generated from .//agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // agpml

import "github.com/antlr4-go/antlr/v4"

// BaseagpmlListener is a complete listener for a parse tree produced by agpmlParser.
type BaseagpmlListener struct{}

var _ agpmlListener = &BaseagpmlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseagpmlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseagpmlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseagpmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseagpmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseagpmlListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseagpmlListener) ExitProgram(ctx *ProgramContext) {}

// EnterHeaderConfigs is called when production headerConfigs is entered.
func (s *BaseagpmlListener) EnterHeaderConfigs(ctx *HeaderConfigsContext) {}

// ExitHeaderConfigs is called when production headerConfigs is exited.
func (s *BaseagpmlListener) ExitHeaderConfigs(ctx *HeaderConfigsContext) {}

// EnterVarConfigs is called when production varConfigs is entered.
func (s *BaseagpmlListener) EnterVarConfigs(ctx *VarConfigsContext) {}

// ExitVarConfigs is called when production varConfigs is exited.
func (s *BaseagpmlListener) ExitVarConfigs(ctx *VarConfigsContext) {}

// EnterDataConfigs is called when production dataConfigs is entered.
func (s *BaseagpmlListener) EnterDataConfigs(ctx *DataConfigsContext) {}

// ExitDataConfigs is called when production dataConfigs is exited.
func (s *BaseagpmlListener) ExitDataConfigs(ctx *DataConfigsContext) {}
