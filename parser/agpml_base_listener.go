// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Agpml

import "github.com/antlr4-go/antlr/v4"

// BaseAgpmlListener is a complete listener for a parse tree produced by AgpmlParser.
type BaseAgpmlListener struct{}

var _ AgpmlListener = &BaseAgpmlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAgpmlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAgpmlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAgpmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAgpmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseAgpmlListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseAgpmlListener) ExitProgram(ctx *ProgramContext) {}

// EnterHeaderConfigs is called when production headerConfigs is entered.
func (s *BaseAgpmlListener) EnterHeaderConfigs(ctx *HeaderConfigsContext) {}

// ExitHeaderConfigs is called when production headerConfigs is exited.
func (s *BaseAgpmlListener) ExitHeaderConfigs(ctx *HeaderConfigsContext) {}

// EnterVarConfigs is called when production varConfigs is entered.
func (s *BaseAgpmlListener) EnterVarConfigs(ctx *VarConfigsContext) {}

// ExitVarConfigs is called when production varConfigs is exited.
func (s *BaseAgpmlListener) ExitVarConfigs(ctx *VarConfigsContext) {}

// EnterDataConfigs is called when production dataConfigs is entered.
func (s *BaseAgpmlListener) EnterDataConfigs(ctx *DataConfigsContext) {}

// ExitDataConfigs is called when production dataConfigs is exited.
func (s *BaseAgpmlListener) ExitDataConfigs(ctx *DataConfigsContext) {}

// EnterStyle is called when production style is entered.
func (s *BaseAgpmlListener) EnterStyle(ctx *StyleContext) {}

// ExitStyle is called when production style is exited.
func (s *BaseAgpmlListener) ExitStyle(ctx *StyleContext) {}

// EnterStyleConfig is called when production styleConfig is entered.
func (s *BaseAgpmlListener) EnterStyleConfig(ctx *StyleConfigContext) {}

// ExitStyleConfig is called when production styleConfig is exited.
func (s *BaseAgpmlListener) ExitStyleConfig(ctx *StyleConfigContext) {}
