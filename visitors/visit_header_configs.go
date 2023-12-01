package visitors

import (
	"agpml/header_configs"
	"agpml/parser"
	"regexp"
	"strings"
)

func receivedHeaderConfigValue(ctx *parser.HeaderConfigContext) string {
	if ctx.BOOLEAN() != nil {
		return "BOOLEAN"
	} else if ctx.STRING() != nil {
		return "STRING"
	} else if ctx.COLOR() != nil {
		return "COLOR"
	} else if ctx.NUMBER() != nil {
		return "NUMBER"
	}
	return ""
}

func (v *BaseAgpmlVisitor) VisitHeaderConfigs(ctx *parser.HeaderConfigsContext) interface{} {
	println("VisitHeaderConfigs")
	for _, child := range ctx.GetChildren() {
		v.VisitHeaderConfig(child.(*parser.HeaderConfigContext))
	}
	return v.VisitChildren(ctx)
}

func (v *BaseAgpmlVisitor) VisitHeaderConfig(ctx *parser.HeaderConfigContext) interface{} {
	println("VisitHeaderConfig")

	attribute := strings.TrimSpace(ctx.ATTRIBUTE().GetText())
	if !header_configs.IsAcceptedHeaderConfig(attribute) {

	}

	switch attribute {
	case "inlineCss":
		if ctx.BOOLEAN() == nil {
			SemanticError(ctx.GetStart().GetLine(), "InlineCss must be a boolean, received "+receivedHeaderConfigValue(ctx))
		}
		header_configs.InlineCss = true
	case "title":
		if ctx.STRING() == nil {
			SemanticError(ctx.GetStart().GetLine(), "Title must be a string, received "+receivedHeaderConfigValue(ctx))
		}
		expr, _ := regexp.Compile(`^".*"$`)
		title := expr.FindString(ctx.STRING().GetText())
		if title == "" {
			SemanticError(ctx.GetStart().GetLine(), "Invalid title: "+ctx.STRING().GetText())
		}
		header_configs.Title = strings.TrimSpace(title)
	default:
		SemanticError(ctx.GetStart().GetLine(), "Invalid header config: "+attribute)
	}

	return v.VisitChildren(ctx)
}
