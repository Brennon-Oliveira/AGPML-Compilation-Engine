// Code generated from .//agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // agpml

import "github.com/antlr4-go/antlr/v4"

type BaseagpmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseagpmlVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagpmlVisitor) VisitHeaderConfigs(ctx *HeaderConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagpmlVisitor) VisitVarConfigs(ctx *VarConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseagpmlVisitor) VisitDataConfigs(ctx *DataConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}
