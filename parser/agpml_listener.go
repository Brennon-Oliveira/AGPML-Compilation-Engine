// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Agpml

import "github.com/antlr4-go/antlr/v4"

// AgpmlListener is a complete listener for a parse tree produced by AgpmlParser.
type AgpmlListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterHeaderConfigs is called when entering the headerConfigs production.
	EnterHeaderConfigs(c *HeaderConfigsContext)

	// EnterHeaderConfig is called when entering the headerConfig production.
	EnterHeaderConfig(c *HeaderConfigContext)

	// EnterVarConfigs is called when entering the varConfigs production.
	EnterVarConfigs(c *VarConfigsContext)

	// EnterVarConfig is called when entering the varConfig production.
	EnterVarConfig(c *VarConfigContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterIdGroup is called when entering the idGroup production.
	EnterIdGroup(c *IdGroupContext)

	// EnterClassGroup is called when entering the classGroup production.
	EnterClassGroup(c *ClassGroupContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterStyle is called when entering the style production.
	EnterStyle(c *StyleContext)

	// EnterStyleConfig is called when entering the styleConfig production.
	EnterStyleConfig(c *StyleConfigContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitHeaderConfigs is called when exiting the headerConfigs production.
	ExitHeaderConfigs(c *HeaderConfigsContext)

	// ExitHeaderConfig is called when exiting the headerConfig production.
	ExitHeaderConfig(c *HeaderConfigContext)

	// ExitVarConfigs is called when exiting the varConfigs production.
	ExitVarConfigs(c *VarConfigsContext)

	// ExitVarConfig is called when exiting the varConfig production.
	ExitVarConfig(c *VarConfigContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitIdGroup is called when exiting the idGroup production.
	ExitIdGroup(c *IdGroupContext)

	// ExitClassGroup is called when exiting the classGroup production.
	ExitClassGroup(c *ClassGroupContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitStyle is called when exiting the style production.
	ExitStyle(c *StyleContext)

	// ExitStyleConfig is called when exiting the styleConfig production.
	ExitStyleConfig(c *StyleConfigContext)
}
