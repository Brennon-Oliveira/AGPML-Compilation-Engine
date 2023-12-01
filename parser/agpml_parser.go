// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Agpml

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AgpmlParser struct {
	*antlr.BaseParser
}

var AgpmlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agpmlParserInit() {
	staticData := &AgpmlParserStaticData
	staticData.LiteralNames = []string{
		"", "'.header'", "'.vars'", "'.body'", "", "", "", "", "", "", "'*'",
		"'>'", "", "", "'='", "':'", "','", "", "", "", "", "", "'{'", "'}'",
		"", "'..'", "", "'##'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "BOOLEAN", "TEXT_CONTENT", "STRING", "DIVIDER", "FLAG_TITLE",
		"FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "ATTRIBUTE", "VARIABLE", "ATRIBUTION",
		"DEFINER", "SEPARATOR", "COLOR", "NUMBER", "DIGIT", "COMMENT", "WS",
		"OPEN_CURLY_BRACKET", "CLOSE_CURLY_BRACKET", "OPEN_CLASS", "CLOSE_CLASS",
		"OPEN_ID", "CLOSE_ID", "PARAGRAPH",
	}
	staticData.RuleNames = []string{
		"program", "headerConfigs", "headerConfig", "varConfigs", "varConfig",
		"body", "idGroup", "classGroup", "element", "line", "style", "styleConfig",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 116, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 3, 0, 27, 8, 0, 1, 0, 1, 0, 3, 0, 31, 8,
		0, 3, 0, 33, 8, 0, 1, 0, 1, 0, 3, 0, 37, 8, 0, 1, 0, 1, 0, 1, 1, 4, 1,
		42, 8, 1, 11, 1, 12, 1, 43, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 4, 3, 51, 8,
		3, 11, 3, 12, 3, 52, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 4, 5, 62,
		8, 5, 11, 5, 12, 5, 63, 1, 6, 1, 6, 5, 6, 68, 8, 6, 10, 6, 12, 6, 71, 9,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 77, 8, 7, 10, 7, 12, 7, 80, 9, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 3, 8, 86, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 94, 8, 9, 1, 9, 3, 9, 97, 8, 9, 1, 9, 3, 9, 100, 8, 9, 1, 10, 1, 10,
		4, 10, 104, 8, 10, 11, 10, 12, 10, 105, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 114, 8, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 0, 2, 3, 0, 4, 4, 6, 6, 17, 18, 4, 0, 4, 4, 6, 6, 13, 13,
		17, 18, 124, 0, 24, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 45, 1, 0, 0, 0, 6,
		50, 1, 0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12, 65, 1, 0, 0,
		0, 14, 74, 1, 0, 0, 0, 16, 85, 1, 0, 0, 0, 18, 93, 1, 0, 0, 0, 20, 101,
		1, 0, 0, 0, 22, 109, 1, 0, 0, 0, 24, 26, 5, 1, 0, 0, 25, 27, 3, 2, 1, 0,
		26, 25, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 32, 1, 0, 0, 0, 28, 30, 5,
		2, 0, 0, 29, 31, 3, 6, 3, 0, 30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31,
		33, 1, 0, 0, 0, 32, 28, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 36, 5, 3, 0, 0, 35, 37, 3, 10, 5, 0, 36, 35, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 5, 0, 0, 1, 39, 1, 1, 0, 0, 0,
		40, 42, 3, 4, 2, 0, 41, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 41, 1,
		0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 3, 1, 0, 0, 0, 45, 46, 5, 12, 0, 0, 46,
		47, 5, 14, 0, 0, 47, 48, 7, 0, 0, 0, 48, 5, 1, 0, 0, 0, 49, 51, 3, 8, 4,
		0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53,
		1, 0, 0, 0, 53, 7, 1, 0, 0, 0, 54, 55, 5, 13, 0, 0, 55, 56, 5, 14, 0, 0,
		56, 57, 7, 1, 0, 0, 57, 9, 1, 0, 0, 0, 58, 62, 3, 16, 8, 0, 59, 62, 3,
		12, 6, 0, 60, 62, 3, 14, 7, 0, 61, 58, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0,
		61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1,
		0, 0, 0, 64, 11, 1, 0, 0, 0, 65, 69, 5, 26, 0, 0, 66, 68, 3, 16, 8, 0,
		67, 66, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1,
		0, 0, 0, 70, 72, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 5, 27, 0, 0, 73,
		13, 1, 0, 0, 0, 74, 78, 5, 24, 0, 0, 75, 77, 3, 16, 8, 0, 76, 75, 1, 0,
		0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81,
		1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 5, 25, 0, 0, 82, 15, 1, 0, 0, 0,
		83, 86, 5, 7, 0, 0, 84, 86, 3, 18, 9, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1,
		0, 0, 0, 86, 17, 1, 0, 0, 0, 87, 94, 5, 9, 0, 0, 88, 94, 5, 10, 0, 0, 89,
		94, 5, 11, 0, 0, 90, 94, 5, 8, 0, 0, 91, 94, 5, 5, 0, 0, 92, 94, 3, 20,
		10, 0, 93, 87, 1, 0, 0, 0, 93, 88, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0, 93,
		90, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 96, 1, 0, 0,
		0, 95, 97, 3, 20, 10, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99,
		1, 0, 0, 0, 98, 100, 5, 5, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 19, 1, 0, 0, 0, 101, 103, 5, 22, 0, 0, 102, 104, 3, 22, 11, 0,
		103, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 5, 23, 0, 0, 108, 21,
		1, 0, 0, 0, 109, 110, 5, 12, 0, 0, 110, 111, 5, 15, 0, 0, 111, 113, 7,
		1, 0, 0, 112, 114, 5, 16, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0,
		0, 114, 23, 1, 0, 0, 0, 16, 26, 30, 32, 36, 43, 52, 61, 63, 69, 78, 85,
		93, 96, 99, 105, 113,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AgpmlParserInit initializes any static state used to implement AgpmlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAgpmlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgpmlParserInit() {
	staticData := &AgpmlParserStaticData
	staticData.once.Do(agpmlParserInit)
}

// NewAgpmlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAgpmlParser(input antlr.TokenStream) *AgpmlParser {
	AgpmlParserInit()
	this := new(AgpmlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AgpmlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Agpml.g4"

	return this
}

// AgpmlParser tokens.
const (
	AgpmlParserEOF                 = antlr.TokenEOF
	AgpmlParserT__0                = 1
	AgpmlParserT__1                = 2
	AgpmlParserT__2                = 3
	AgpmlParserBOOLEAN             = 4
	AgpmlParserTEXT_CONTENT        = 5
	AgpmlParserSTRING              = 6
	AgpmlParserDIVIDER             = 7
	AgpmlParserFLAG_TITLE          = 8
	AgpmlParserFLAG_OLIST          = 9
	AgpmlParserFLAG_ULIST          = 10
	AgpmlParserFLAG_QUOTE          = 11
	AgpmlParserATTRIBUTE           = 12
	AgpmlParserVARIABLE            = 13
	AgpmlParserATRIBUTION          = 14
	AgpmlParserDEFINER             = 15
	AgpmlParserSEPARATOR           = 16
	AgpmlParserCOLOR               = 17
	AgpmlParserNUMBER              = 18
	AgpmlParserDIGIT               = 19
	AgpmlParserCOMMENT             = 20
	AgpmlParserWS                  = 21
	AgpmlParserOPEN_CURLY_BRACKET  = 22
	AgpmlParserCLOSE_CURLY_BRACKET = 23
	AgpmlParserOPEN_CLASS          = 24
	AgpmlParserCLOSE_CLASS         = 25
	AgpmlParserOPEN_ID             = 26
	AgpmlParserCLOSE_ID            = 27
	AgpmlParserPARAGRAPH           = 28
)

// AgpmlParser rules.
const (
	AgpmlParserRULE_program       = 0
	AgpmlParserRULE_headerConfigs = 1
	AgpmlParserRULE_headerConfig  = 2
	AgpmlParserRULE_varConfigs    = 3
	AgpmlParserRULE_varConfig     = 4
	AgpmlParserRULE_body          = 5
	AgpmlParserRULE_idGroup       = 6
	AgpmlParserRULE_classGroup    = 7
	AgpmlParserRULE_element       = 8
	AgpmlParserRULE_line          = 9
	AgpmlParserRULE_style         = 10
	AgpmlParserRULE_styleConfig   = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	HeaderConfigs() IHeaderConfigsContext
	Body() IBodyContext
	VarConfigs() IVarConfigsContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(AgpmlParserEOF, 0)
}

func (s *ProgramContext) HeaderConfigs() IHeaderConfigsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderConfigsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderConfigsContext)
}

func (s *ProgramContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ProgramContext) VarConfigs() IVarConfigsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarConfigsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarConfigsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgpmlParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(AgpmlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserATTRIBUTE {
		{
			p.SetState(25)
			p.HeaderConfigs()
		}

	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__1 {
		{
			p.SetState(28)
			p.Match(AgpmlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AgpmlParserVARIABLE {
			{
				p.SetState(29)
				p.VarConfigs()
			}

		}

	}
	{
		p.SetState(34)
		p.Match(AgpmlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&88084384) != 0 {
		{
			p.SetState(35)
			p.Body()
		}

	}
	{
		p.SetState(38)
		p.Match(AgpmlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeaderConfigsContext is an interface to support dynamic dispatch.
type IHeaderConfigsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllHeaderConfig() []IHeaderConfigContext
	HeaderConfig(i int) IHeaderConfigContext

	// IsHeaderConfigsContext differentiates from other interfaces.
	IsHeaderConfigsContext()
}

type HeaderConfigsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderConfigsContext() *HeaderConfigsContext {
	var p = new(HeaderConfigsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_headerConfigs
	return p
}

func InitEmptyHeaderConfigsContext(p *HeaderConfigsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_headerConfigs
}

func (*HeaderConfigsContext) IsHeaderConfigsContext() {}

func NewHeaderConfigsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderConfigsContext {
	var p = new(HeaderConfigsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_headerConfigs

	return p
}

func (s *HeaderConfigsContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderConfigsContext) AllHeaderConfig() []IHeaderConfigContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeaderConfigContext); ok {
			len++
		}
	}

	tst := make([]IHeaderConfigContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeaderConfigContext); ok {
			tst[i] = t.(IHeaderConfigContext)
			i++
		}
	}

	return tst
}

func (s *HeaderConfigsContext) HeaderConfig(i int) IHeaderConfigContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderConfigContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderConfigContext)
}

func (s *HeaderConfigsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderConfigsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderConfigsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterHeaderConfigs(s)
	}
}

func (s *HeaderConfigsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitHeaderConfigs(s)
	}
}

func (s *HeaderConfigsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitHeaderConfigs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) HeaderConfigs() (localctx IHeaderConfigsContext) {
	localctx = NewHeaderConfigsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgpmlParserRULE_headerConfigs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AgpmlParserATTRIBUTE {
		{
			p.SetState(40)
			p.HeaderConfig()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeaderConfigContext is an interface to support dynamic dispatch.
type IHeaderConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRIBUTE() antlr.TerminalNode
	ATRIBUTION() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	COLOR() antlr.TerminalNode

	// IsHeaderConfigContext differentiates from other interfaces.
	IsHeaderConfigContext()
}

type HeaderConfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderConfigContext() *HeaderConfigContext {
	var p = new(HeaderConfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_headerConfig
	return p
}

func InitEmptyHeaderConfigContext(p *HeaderConfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_headerConfig
}

func (*HeaderConfigContext) IsHeaderConfigContext() {}

func NewHeaderConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderConfigContext {
	var p = new(HeaderConfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_headerConfig

	return p
}

func (s *HeaderConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderConfigContext) ATTRIBUTE() antlr.TerminalNode {
	return s.GetToken(AgpmlParserATTRIBUTE, 0)
}

func (s *HeaderConfigContext) ATRIBUTION() antlr.TerminalNode {
	return s.GetToken(AgpmlParserATRIBUTION, 0)
}

func (s *HeaderConfigContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(AgpmlParserBOOLEAN, 0)
}

func (s *HeaderConfigContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgpmlParserSTRING, 0)
}

func (s *HeaderConfigContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AgpmlParserNUMBER, 0)
}

func (s *HeaderConfigContext) COLOR() antlr.TerminalNode {
	return s.GetToken(AgpmlParserCOLOR, 0)
}

func (s *HeaderConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterHeaderConfig(s)
	}
}

func (s *HeaderConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitHeaderConfig(s)
	}
}

func (s *HeaderConfigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitHeaderConfig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) HeaderConfig() (localctx IHeaderConfigContext) {
	localctx = NewHeaderConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AgpmlParserRULE_headerConfig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(AgpmlParserATTRIBUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(AgpmlParserATRIBUTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&393296) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarConfigsContext is an interface to support dynamic dispatch.
type IVarConfigsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarConfig() []IVarConfigContext
	VarConfig(i int) IVarConfigContext

	// IsVarConfigsContext differentiates from other interfaces.
	IsVarConfigsContext()
}

type VarConfigsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarConfigsContext() *VarConfigsContext {
	var p = new(VarConfigsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_varConfigs
	return p
}

func InitEmptyVarConfigsContext(p *VarConfigsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_varConfigs
}

func (*VarConfigsContext) IsVarConfigsContext() {}

func NewVarConfigsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarConfigsContext {
	var p = new(VarConfigsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_varConfigs

	return p
}

func (s *VarConfigsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarConfigsContext) AllVarConfig() []IVarConfigContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarConfigContext); ok {
			len++
		}
	}

	tst := make([]IVarConfigContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarConfigContext); ok {
			tst[i] = t.(IVarConfigContext)
			i++
		}
	}

	return tst
}

func (s *VarConfigsContext) VarConfig(i int) IVarConfigContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarConfigContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarConfigContext)
}

func (s *VarConfigsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarConfigsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarConfigsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterVarConfigs(s)
	}
}

func (s *VarConfigsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitVarConfigs(s)
	}
}

func (s *VarConfigsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitVarConfigs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) VarConfigs() (localctx IVarConfigsContext) {
	localctx = NewVarConfigsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AgpmlParserRULE_varConfigs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AgpmlParserVARIABLE {
		{
			p.SetState(49)
			p.VarConfig()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarConfigContext is an interface to support dynamic dispatch.
type IVarConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVARIABLE() []antlr.TerminalNode
	VARIABLE(i int) antlr.TerminalNode
	ATRIBUTION() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	COLOR() antlr.TerminalNode

	// IsVarConfigContext differentiates from other interfaces.
	IsVarConfigContext()
}

type VarConfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarConfigContext() *VarConfigContext {
	var p = new(VarConfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_varConfig
	return p
}

func InitEmptyVarConfigContext(p *VarConfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_varConfig
}

func (*VarConfigContext) IsVarConfigContext() {}

func NewVarConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarConfigContext {
	var p = new(VarConfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_varConfig

	return p
}

func (s *VarConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *VarConfigContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(AgpmlParserVARIABLE)
}

func (s *VarConfigContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(AgpmlParserVARIABLE, i)
}

func (s *VarConfigContext) ATRIBUTION() antlr.TerminalNode {
	return s.GetToken(AgpmlParserATRIBUTION, 0)
}

func (s *VarConfigContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(AgpmlParserBOOLEAN, 0)
}

func (s *VarConfigContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgpmlParserSTRING, 0)
}

func (s *VarConfigContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AgpmlParserNUMBER, 0)
}

func (s *VarConfigContext) COLOR() antlr.TerminalNode {
	return s.GetToken(AgpmlParserCOLOR, 0)
}

func (s *VarConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterVarConfig(s)
	}
}

func (s *VarConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitVarConfig(s)
	}
}

func (s *VarConfigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitVarConfig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) VarConfig() (localctx IVarConfigContext) {
	localctx = NewVarConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AgpmlParserRULE_varConfig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(AgpmlParserVARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(AgpmlParserATRIBUTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&401488) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElement() []IElementContext
	Element(i int) IElementContext
	AllIdGroup() []IIdGroupContext
	IdGroup(i int) IIdGroupContext
	AllClassGroup() []IClassGroupContext
	ClassGroup(i int) IClassGroupContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *BodyContext) AllIdGroup() []IIdGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdGroupContext); ok {
			len++
		}
	}

	tst := make([]IIdGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdGroupContext); ok {
			tst[i] = t.(IIdGroupContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) IdGroup(i int) IIdGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdGroupContext)
}

func (s *BodyContext) AllClassGroup() []IClassGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassGroupContext); ok {
			len++
		}
	}

	tst := make([]IClassGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassGroupContext); ok {
			tst[i] = t.(IClassGroupContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) ClassGroup(i int) IClassGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassGroupContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AgpmlParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&88084384) != 0) {
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AgpmlParserTEXT_CONTENT, AgpmlParserDIVIDER, AgpmlParserFLAG_TITLE, AgpmlParserFLAG_OLIST, AgpmlParserFLAG_ULIST, AgpmlParserFLAG_QUOTE, AgpmlParserOPEN_CURLY_BRACKET:
			{
				p.SetState(58)
				p.Element()
			}

		case AgpmlParserOPEN_ID:
			{
				p.SetState(59)
				p.IdGroup()
			}

		case AgpmlParserOPEN_CLASS:
			{
				p.SetState(60)
				p.ClassGroup()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdGroupContext is an interface to support dynamic dispatch.
type IIdGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_ID() antlr.TerminalNode
	CLOSE_ID() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsIdGroupContext differentiates from other interfaces.
	IsIdGroupContext()
}

type IdGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdGroupContext() *IdGroupContext {
	var p = new(IdGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_idGroup
	return p
}

func InitEmptyIdGroupContext(p *IdGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_idGroup
}

func (*IdGroupContext) IsIdGroupContext() {}

func NewIdGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdGroupContext {
	var p = new(IdGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_idGroup

	return p
}

func (s *IdGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *IdGroupContext) OPEN_ID() antlr.TerminalNode {
	return s.GetToken(AgpmlParserOPEN_ID, 0)
}

func (s *IdGroupContext) CLOSE_ID() antlr.TerminalNode {
	return s.GetToken(AgpmlParserCLOSE_ID, 0)
}

func (s *IdGroupContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *IdGroupContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *IdGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterIdGroup(s)
	}
}

func (s *IdGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitIdGroup(s)
	}
}

func (s *IdGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitIdGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) IdGroup() (localctx IIdGroupContext) {
	localctx = NewIdGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AgpmlParserRULE_idGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(AgpmlParserOPEN_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4198304) != 0 {
		{
			p.SetState(66)
			p.Element()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(AgpmlParserCLOSE_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassGroupContext is an interface to support dynamic dispatch.
type IClassGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_CLASS() antlr.TerminalNode
	CLOSE_CLASS() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsClassGroupContext differentiates from other interfaces.
	IsClassGroupContext()
}

type ClassGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassGroupContext() *ClassGroupContext {
	var p = new(ClassGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_classGroup
	return p
}

func InitEmptyClassGroupContext(p *ClassGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_classGroup
}

func (*ClassGroupContext) IsClassGroupContext() {}

func NewClassGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassGroupContext {
	var p = new(ClassGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_classGroup

	return p
}

func (s *ClassGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassGroupContext) OPEN_CLASS() antlr.TerminalNode {
	return s.GetToken(AgpmlParserOPEN_CLASS, 0)
}

func (s *ClassGroupContext) CLOSE_CLASS() antlr.TerminalNode {
	return s.GetToken(AgpmlParserCLOSE_CLASS, 0)
}

func (s *ClassGroupContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *ClassGroupContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ClassGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterClassGroup(s)
	}
}

func (s *ClassGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitClassGroup(s)
	}
}

func (s *ClassGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitClassGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) ClassGroup() (localctx IClassGroupContext) {
	localctx = NewClassGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AgpmlParserRULE_classGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(AgpmlParserOPEN_CLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4198304) != 0 {
		{
			p.SetState(75)
			p.Element()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(AgpmlParserCLOSE_CLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIVIDER() antlr.TerminalNode
	Line() ILineContext

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_element
	return p
}

func InitEmptyElementContext(p *ElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_element
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) DIVIDER() antlr.TerminalNode {
	return s.GetToken(AgpmlParserDIVIDER, 0)
}

func (s *ElementContext) Line() ILineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AgpmlParserRULE_element)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AgpmlParserDIVIDER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(AgpmlParserDIVIDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgpmlParserTEXT_CONTENT, AgpmlParserFLAG_TITLE, AgpmlParserFLAG_OLIST, AgpmlParserFLAG_ULIST, AgpmlParserFLAG_QUOTE, AgpmlParserOPEN_CURLY_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Line()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLAG_OLIST() antlr.TerminalNode
	FLAG_ULIST() antlr.TerminalNode
	FLAG_QUOTE() antlr.TerminalNode
	FLAG_TITLE() antlr.TerminalNode
	AllTEXT_CONTENT() []antlr.TerminalNode
	TEXT_CONTENT(i int) antlr.TerminalNode
	AllStyle() []IStyleContext
	Style(i int) IStyleContext

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_line
	return p
}

func InitEmptyLineContext(p *LineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_line
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) FLAG_OLIST() antlr.TerminalNode {
	return s.GetToken(AgpmlParserFLAG_OLIST, 0)
}

func (s *LineContext) FLAG_ULIST() antlr.TerminalNode {
	return s.GetToken(AgpmlParserFLAG_ULIST, 0)
}

func (s *LineContext) FLAG_QUOTE() antlr.TerminalNode {
	return s.GetToken(AgpmlParserFLAG_QUOTE, 0)
}

func (s *LineContext) FLAG_TITLE() antlr.TerminalNode {
	return s.GetToken(AgpmlParserFLAG_TITLE, 0)
}

func (s *LineContext) AllTEXT_CONTENT() []antlr.TerminalNode {
	return s.GetTokens(AgpmlParserTEXT_CONTENT)
}

func (s *LineContext) TEXT_CONTENT(i int) antlr.TerminalNode {
	return s.GetToken(AgpmlParserTEXT_CONTENT, i)
}

func (s *LineContext) AllStyle() []IStyleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStyleContext); ok {
			len++
		}
	}

	tst := make([]IStyleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStyleContext); ok {
			tst[i] = t.(IStyleContext)
			i++
		}
	}

	return tst
}

func (s *LineContext) Style(i int) IStyleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStyleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStyleContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitLine(s)
	}
}

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AgpmlParserRULE_line)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AgpmlParserFLAG_OLIST:
		{
			p.SetState(87)
			p.Match(AgpmlParserFLAG_OLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgpmlParserFLAG_ULIST:
		{
			p.SetState(88)
			p.Match(AgpmlParserFLAG_ULIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgpmlParserFLAG_QUOTE:
		{
			p.SetState(89)
			p.Match(AgpmlParserFLAG_QUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgpmlParserFLAG_TITLE:
		{
			p.SetState(90)
			p.Match(AgpmlParserFLAG_TITLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgpmlParserTEXT_CONTENT:
		{
			p.SetState(91)
			p.Match(AgpmlParserTEXT_CONTENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AgpmlParserOPEN_CURLY_BRACKET:
		{
			p.SetState(92)
			p.Style()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(95)
			p.Style()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(98)
			p.Match(AgpmlParserTEXT_CONTENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStyleContext is an interface to support dynamic dispatch.
type IStyleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_CURLY_BRACKET() antlr.TerminalNode
	CLOSE_CURLY_BRACKET() antlr.TerminalNode
	AllStyleConfig() []IStyleConfigContext
	StyleConfig(i int) IStyleConfigContext

	// IsStyleContext differentiates from other interfaces.
	IsStyleContext()
}

type StyleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleContext() *StyleContext {
	var p = new(StyleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_style
	return p
}

func InitEmptyStyleContext(p *StyleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_style
}

func (*StyleContext) IsStyleContext() {}

func NewStyleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleContext {
	var p = new(StyleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_style

	return p
}

func (s *StyleContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleContext) OPEN_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(AgpmlParserOPEN_CURLY_BRACKET, 0)
}

func (s *StyleContext) CLOSE_CURLY_BRACKET() antlr.TerminalNode {
	return s.GetToken(AgpmlParserCLOSE_CURLY_BRACKET, 0)
}

func (s *StyleContext) AllStyleConfig() []IStyleConfigContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStyleConfigContext); ok {
			len++
		}
	}

	tst := make([]IStyleConfigContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStyleConfigContext); ok {
			tst[i] = t.(IStyleConfigContext)
			i++
		}
	}

	return tst
}

func (s *StyleContext) StyleConfig(i int) IStyleConfigContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStyleConfigContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStyleConfigContext)
}

func (s *StyleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterStyle(s)
	}
}

func (s *StyleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitStyle(s)
	}
}

func (s *StyleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitStyle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) Style() (localctx IStyleContext) {
	localctx = NewStyleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AgpmlParserRULE_style)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(AgpmlParserOPEN_CURLY_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AgpmlParserATTRIBUTE {
		{
			p.SetState(102)
			p.StyleConfig()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(AgpmlParserCLOSE_CURLY_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStyleConfigContext is an interface to support dynamic dispatch.
type IStyleConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRIBUTE() antlr.TerminalNode
	DEFINER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	COLOR() antlr.TerminalNode
	VARIABLE() antlr.TerminalNode
	SEPARATOR() antlr.TerminalNode

	// IsStyleConfigContext differentiates from other interfaces.
	IsStyleConfigContext()
}

type StyleConfigContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleConfigContext() *StyleConfigContext {
	var p = new(StyleConfigContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_styleConfig
	return p
}

func InitEmptyStyleConfigContext(p *StyleConfigContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_styleConfig
}

func (*StyleConfigContext) IsStyleConfigContext() {}

func NewStyleConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleConfigContext {
	var p = new(StyleConfigContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_styleConfig

	return p
}

func (s *StyleConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleConfigContext) ATTRIBUTE() antlr.TerminalNode {
	return s.GetToken(AgpmlParserATTRIBUTE, 0)
}

func (s *StyleConfigContext) DEFINER() antlr.TerminalNode {
	return s.GetToken(AgpmlParserDEFINER, 0)
}

func (s *StyleConfigContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(AgpmlParserBOOLEAN, 0)
}

func (s *StyleConfigContext) STRING() antlr.TerminalNode {
	return s.GetToken(AgpmlParserSTRING, 0)
}

func (s *StyleConfigContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AgpmlParserNUMBER, 0)
}

func (s *StyleConfigContext) COLOR() antlr.TerminalNode {
	return s.GetToken(AgpmlParserCOLOR, 0)
}

func (s *StyleConfigContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(AgpmlParserVARIABLE, 0)
}

func (s *StyleConfigContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(AgpmlParserSEPARATOR, 0)
}

func (s *StyleConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleConfigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterStyleConfig(s)
	}
}

func (s *StyleConfigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitStyleConfig(s)
	}
}

func (s *StyleConfigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitStyleConfig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) StyleConfig() (localctx IStyleConfigContext) {
	localctx = NewStyleConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AgpmlParserRULE_styleConfig)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(AgpmlParserATTRIBUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(AgpmlParserDEFINER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&401488) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserSEPARATOR {
		{
			p.SetState(112)
			p.Match(AgpmlParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
