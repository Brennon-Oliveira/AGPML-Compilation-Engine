// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Agpml

import "github.com/antlr4-go/antlr/v4"

type BaseAgpmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAgpmlVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitHeaderConfigs(ctx *HeaderConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitVarConfigs(ctx *VarConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitDataConfigs(ctx *DataConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitStyle(ctx *StyleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitStyleConfig(ctx *StyleConfigContext) interface{} {
	return v.VisitChildren(ctx)
}
