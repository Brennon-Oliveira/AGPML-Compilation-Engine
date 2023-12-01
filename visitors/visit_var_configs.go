package visitors

import (
	"agpml/parser"
	"agpml/utils"
	"agpml/variables"
)

func (v *BaseAgpmlVisitor) VisitVarConfigs(ctx *parser.VarConfigsContext) interface{} {
	println("VisitVarConfigs")
	for _, child := range ctx.GetChildren() {
		v.VisitVarConfig(child.(*parser.VarConfigContext))
	}
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitVarConfig(ctx *parser.VarConfigContext) interface{} {
	println("VisitVarConfig")
	variable := &variables.Variable{}
	if ctx.VARIABLE(0) == nil {
		SemanticError(ctx.GetStart().GetLine(), "Invalid variable")
	}

	if ctx.ATRIBUTION() == nil {
		SemanticError(ctx.GetStart().GetLine(), "Need a '=' to set a value to variable")
	}

	variable.Name = ctx.VARIABLE(0).GetText()
	valueToSet := utils.GetVariableSetValue(ctx)

	if valueToSet == "" {
		if ctx.VARIABLE(1) == nil {
			SemanticError(ctx.GetStart().GetLine(), "Variable need a value")
		}

		if _, ok := variables.Variables[ctx.VARIABLE(1).GetText()]; !ok {
			SemanticError(ctx.GetStart().GetLine(), "Variable "+ctx.VARIABLE(1).GetText()+" doesn't exist")
		}

		variable.Value = variables.Variables[ctx.VARIABLE(1).GetText()].Value
	} else {
		variable.Value = valueToSet
	}

	if _, ok := variables.Variables[variable.Name]; ok {
		SemanticError(ctx.GetStart().GetLine(), "Variable "+variable.Name+" already exist")
	}

	println("Variable: ", variable.Name, "Value: ", variable.Value)
	variables.Variables[variable.Name] = *variable

	return v.VisitChildren(ctx)
}
