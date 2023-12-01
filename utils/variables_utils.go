package utils

import "agpml/parser"

func GetVariableSetValue(ctx *parser.VarConfigContext) string {
	if ctx.BOOLEAN() != nil {
		return ctx.BOOLEAN().GetText()
	}
	if ctx.STRING() != nil {
		return ctx.STRING().GetText()
	}

	if ctx.COLOR() != nil {
		return ctx.COLOR().GetText()
	}

	if ctx.NUMBER() != nil {
		return ctx.NUMBER().GetText()
	}

	return ""
}
