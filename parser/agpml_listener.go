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

	// EnterVarConfigs is called when entering the varConfigs production.
	EnterVarConfigs(c *VarConfigsContext)

	// EnterDataConfigs is called when entering the dataConfigs production.
	EnterDataConfigs(c *DataConfigsContext)

	// EnterStyle is called when entering the style production.
	EnterStyle(c *StyleContext)

	// EnterStyleConfig is called when entering the styleConfig production.
	EnterStyleConfig(c *StyleConfigContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitHeaderConfigs is called when exiting the headerConfigs production.
	ExitHeaderConfigs(c *HeaderConfigsContext)

	// ExitVarConfigs is called when exiting the varConfigs production.
	ExitVarConfigs(c *VarConfigsContext)

	// ExitDataConfigs is called when exiting the dataConfigs production.
	ExitDataConfigs(c *DataConfigsContext)

	// ExitStyle is called when exiting the style production.
	ExitStyle(c *StyleContext)

	// ExitStyleConfig is called when exiting the styleConfig production.
	ExitStyleConfig(c *StyleConfigContext)
}
