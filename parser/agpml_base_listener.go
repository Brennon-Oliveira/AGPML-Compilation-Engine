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

// EnterHeaderConfigs is called when production header_configs is entered.
func (s *BaseAgpmlListener) EnterHeaderConfigs(ctx *HeaderConfigsContext) {}

// ExitHeaderConfigs is called when production header_configs is exited.
func (s *BaseAgpmlListener) ExitHeaderConfigs(ctx *HeaderConfigsContext) {}

// EnterHeaderConfig is called when production headerConfig is entered.
func (s *BaseAgpmlListener) EnterHeaderConfig(ctx *HeaderConfigContext) {}

// ExitHeaderConfig is called when production headerConfig is exited.
func (s *BaseAgpmlListener) ExitHeaderConfig(ctx *HeaderConfigContext) {}

// EnterVarConfigs is called when production varConfigs is entered.
func (s *BaseAgpmlListener) EnterVarConfigs(ctx *VarConfigsContext) {}

// ExitVarConfigs is called when production varConfigs is exited.
func (s *BaseAgpmlListener) ExitVarConfigs(ctx *VarConfigsContext) {}

// EnterVarConfig is called when production varConfig is entered.
func (s *BaseAgpmlListener) EnterVarConfig(ctx *VarConfigContext) {}

// ExitVarConfig is called when production varConfig is exited.
func (s *BaseAgpmlListener) ExitVarConfig(ctx *VarConfigContext) {}

// EnterBody is called when production body is entered.
func (s *BaseAgpmlListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseAgpmlListener) ExitBody(ctx *BodyContext) {}

// EnterIdGroup is called when production idGroup is entered.
func (s *BaseAgpmlListener) EnterIdGroup(ctx *IdGroupContext) {}

// ExitIdGroup is called when production idGroup is exited.
func (s *BaseAgpmlListener) ExitIdGroup(ctx *IdGroupContext) {}

// EnterClassGroup is called when production classGroup is entered.
func (s *BaseAgpmlListener) EnterClassGroup(ctx *ClassGroupContext) {}

// ExitClassGroup is called when production classGroup is exited.
func (s *BaseAgpmlListener) ExitClassGroup(ctx *ClassGroupContext) {}

// EnterElement is called when production element is entered.
func (s *BaseAgpmlListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseAgpmlListener) ExitElement(ctx *ElementContext) {}

// EnterLine is called when production line is entered.
func (s *BaseAgpmlListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseAgpmlListener) ExitLine(ctx *LineContext) {}

// EnterStyle is called when production style is entered.
func (s *BaseAgpmlListener) EnterStyle(ctx *StyleContext) {}

// ExitStyle is called when production style is exited.
func (s *BaseAgpmlListener) ExitStyle(ctx *StyleContext) {}

// EnterStyleConfig is called when production styleConfig is entered.
func (s *BaseAgpmlListener) EnterStyleConfig(ctx *StyleConfigContext) {}

// ExitStyleConfig is called when production styleConfig is exited.
func (s *BaseAgpmlListener) ExitStyleConfig(ctx *StyleConfigContext) {}
