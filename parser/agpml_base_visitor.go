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

func (v *BaseAgpmlVisitor) VisitHeaderConfig(ctx *HeaderConfigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitVarConfigs(ctx *VarConfigsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitVarConfig(ctx *VarConfigContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitIdGroup(ctx *IdGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitClassGroup(ctx *ClassGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitStyle(ctx *StyleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitStyleConfig(ctx *StyleConfigContext) interface{} {
	return v.VisitChildren(ctx)
}
