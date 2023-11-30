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
		"", "'.header'", "'.vars'", "'.body'", "'3'", "'1'", "", "", "", "",
		"", "", "'*'", "'>'", "", "", "'='", "':'", "','", "", "", "", "", "",
		"'{'", "'}'", "", "'..'", "", "'##'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "BOOLEAN", "TEXT_CONTENT", "STRING", "DIVIDER",
		"FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "ATTRIBUTE",
		"VARIABLE", "ATRIBUTION", "DEFINER", "SEPARATOR", "COLOR", "NUMBER",
		"DIGIT", "COMMENT", "WS", "OPEN_CURLY_BRACKET", "CLOSE_CURLY_BRACKET",
		"OPEN_CLASS", "CLOSE_CLASS", "OPEN_ID", "CLOSE_ID",
	}
	staticData.RuleNames = []string{
		"program", "headerConfigs", "headerConfig", "varConfigs", "varConfig",
		"dataConfigs", "style", "styleConfig",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 61, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 3, 0, 19, 8, 0, 1, 0,
		1, 0, 3, 0, 23, 8, 0, 3, 0, 25, 8, 0, 1, 0, 1, 0, 3, 0, 29, 8, 0, 1, 0,
		1, 0, 1, 1, 4, 1, 34, 8, 1, 11, 1, 12, 1, 35, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 4, 3, 43, 8, 3, 11, 3, 12, 3, 44, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 3, 6, 55, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2,
		4, 6, 8, 10, 12, 14, 0, 2, 3, 0, 6, 6, 8, 8, 19, 20, 4, 0, 6, 6, 8, 8,
		15, 15, 19, 20, 59, 0, 16, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4, 37, 1, 0,
		0, 0, 6, 42, 1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 50, 1, 0, 0, 0, 12, 52,
		1, 0, 0, 0, 14, 58, 1, 0, 0, 0, 16, 18, 5, 1, 0, 0, 17, 19, 3, 2, 1, 0,
		18, 17, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 24, 1, 0, 0, 0, 20, 22, 5,
		2, 0, 0, 21, 23, 3, 6, 3, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23,
		25, 1, 0, 0, 0, 24, 20, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 1, 0, 0,
		0, 26, 28, 5, 3, 0, 0, 27, 29, 3, 10, 5, 0, 28, 27, 1, 0, 0, 0, 28, 29,
		1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 5, 0, 0, 1, 31, 1, 1, 0, 0, 0,
		32, 34, 3, 4, 2, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 3, 1, 0, 0, 0, 37, 38, 5, 14, 0, 0, 38,
		39, 5, 16, 0, 0, 39, 40, 7, 0, 0, 0, 40, 5, 1, 0, 0, 0, 41, 43, 3, 8, 4,
		0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45,
		1, 0, 0, 0, 45, 7, 1, 0, 0, 0, 46, 47, 5, 15, 0, 0, 47, 48, 5, 16, 0, 0,
		48, 49, 7, 1, 0, 0, 49, 9, 1, 0, 0, 0, 50, 51, 5, 4, 0, 0, 51, 11, 1, 0,
		0, 0, 52, 54, 5, 24, 0, 0, 53, 55, 3, 14, 7, 0, 54, 53, 1, 0, 0, 0, 54,
		55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 5, 25, 0, 0, 57, 13, 1, 0,
		0, 0, 58, 59, 5, 5, 0, 0, 59, 15, 1, 0, 0, 0, 7, 18, 22, 24, 28, 35, 44,
		54,
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
	AgpmlParserT__3                = 4
	AgpmlParserT__4                = 5
	AgpmlParserBOOLEAN             = 6
	AgpmlParserTEXT_CONTENT        = 7
	AgpmlParserSTRING              = 8
	AgpmlParserDIVIDER             = 9
	AgpmlParserFLAG_TITLE          = 10
	AgpmlParserFLAG_OLIST          = 11
	AgpmlParserFLAG_ULIST          = 12
	AgpmlParserFLAG_QUOTE          = 13
	AgpmlParserATTRIBUTE           = 14
	AgpmlParserVARIABLE            = 15
	AgpmlParserATRIBUTION          = 16
	AgpmlParserDEFINER             = 17
	AgpmlParserSEPARATOR           = 18
	AgpmlParserCOLOR               = 19
	AgpmlParserNUMBER              = 20
	AgpmlParserDIGIT               = 21
	AgpmlParserCOMMENT             = 22
	AgpmlParserWS                  = 23
	AgpmlParserOPEN_CURLY_BRACKET  = 24
	AgpmlParserCLOSE_CURLY_BRACKET = 25
	AgpmlParserOPEN_CLASS          = 26
	AgpmlParserCLOSE_CLASS         = 27
	AgpmlParserOPEN_ID             = 28
	AgpmlParserCLOSE_ID            = 29
)

// AgpmlParser rules.
const (
	AgpmlParserRULE_program       = 0
	AgpmlParserRULE_headerConfigs = 1
	AgpmlParserRULE_headerConfig  = 2
	AgpmlParserRULE_varConfigs    = 3
	AgpmlParserRULE_varConfig     = 4
	AgpmlParserRULE_dataConfigs   = 5
	AgpmlParserRULE_style         = 6
	AgpmlParserRULE_styleConfig   = 7
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	HeaderConfigs() IHeaderConfigsContext
	DataConfigs() IDataConfigsContext
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

func (s *ProgramContext) DataConfigs() IDataConfigsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataConfigsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataConfigsContext)
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
		p.SetState(16)
		p.Match(AgpmlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserATTRIBUTE {
		{
			p.SetState(17)
			p.HeaderConfigs()
		}

	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__1 {
		{
			p.SetState(20)
			p.Match(AgpmlParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AgpmlParserVARIABLE {
			{
				p.SetState(21)
				p.VarConfigs()
			}

		}

	}
	{
		p.SetState(26)
		p.Match(AgpmlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__3 {
		{
			p.SetState(27)
			p.DataConfigs()
		}

	}
	{
		p.SetState(30)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AgpmlParserATTRIBUTE {
		{
			p.SetState(32)
			p.HeaderConfig()
		}

		p.SetState(35)
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
		p.SetState(37)
		p.Match(AgpmlParserATTRIBUTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(AgpmlParserATRIBUTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1573184) != 0) {
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AgpmlParserVARIABLE {
		{
			p.SetState(41)
			p.VarConfig()
		}

		p.SetState(44)
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
		p.SetState(46)
		p.Match(AgpmlParserVARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(AgpmlParserATRIBUTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1605952) != 0) {
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

// IDataConfigsContext is an interface to support dynamic dispatch.
type IDataConfigsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDataConfigsContext differentiates from other interfaces.
	IsDataConfigsContext()
}

type DataConfigsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataConfigsContext() *DataConfigsContext {
	var p = new(DataConfigsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_dataConfigs
	return p
}

func InitEmptyDataConfigsContext(p *DataConfigsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AgpmlParserRULE_dataConfigs
}

func (*DataConfigsContext) IsDataConfigsContext() {}

func NewDataConfigsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataConfigsContext {
	var p = new(DataConfigsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgpmlParserRULE_dataConfigs

	return p
}

func (s *DataConfigsContext) GetParser() antlr.Parser { return s.parser }
func (s *DataConfigsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataConfigsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataConfigsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.EnterDataConfigs(s)
	}
}

func (s *DataConfigsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgpmlListener); ok {
		listenerT.ExitDataConfigs(s)
	}
}

func (s *DataConfigsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AgpmlVisitor:
		return t.VisitDataConfigs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AgpmlParser) DataConfigs() (localctx IDataConfigsContext) {
	localctx = NewDataConfigsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AgpmlParserRULE_dataConfigs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(AgpmlParserT__3)
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

// IStyleContext is an interface to support dynamic dispatch.
type IStyleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_CURLY_BRACKET() antlr.TerminalNode
	CLOSE_CURLY_BRACKET() antlr.TerminalNode
	StyleConfig() IStyleConfigContext

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

func (s *StyleContext) StyleConfig() IStyleConfigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStyleConfigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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
	p.EnterRule(localctx, 12, AgpmlParserRULE_style)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(AgpmlParserOPEN_CURLY_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__4 {
		{
			p.SetState(53)
			p.StyleConfig()
		}

	}
	{
		p.SetState(56)
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
	p.EnterRule(localctx, 14, AgpmlParserRULE_styleConfig)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(AgpmlParserT__4)
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
