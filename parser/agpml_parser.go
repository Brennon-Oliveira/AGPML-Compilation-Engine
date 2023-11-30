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
		"", "'.header'", "'.vars'", "'.body'", "'1'", "'2'", "'3'", "'{'", "'}'",
		"", "", "", "", "'*'", "'>'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "STRING", "DIVIDER", "FLAG_TITLE",
		"FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "VARIABLE", "TEXT", "ATRIBUTION",
		"COLOR", "NUMBER", "DIGIT", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"program", "headerConfigs", "varConfigs", "dataConfigs", "style", "styleConfig",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 43, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 3, 0, 15, 8, 0, 1, 0, 1, 0, 3, 0, 19, 8, 0,
		3, 0, 21, 8, 0, 1, 0, 1, 0, 3, 0, 25, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 37, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 0, 0, 6, 0, 2, 4, 6, 8, 10, 0, 0, 41, 0, 12, 1, 0, 0, 0, 2, 28, 1,
		0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8, 34, 1, 0, 0, 0, 10, 40,
		1, 0, 0, 0, 12, 14, 5, 1, 0, 0, 13, 15, 3, 2, 1, 0, 14, 13, 1, 0, 0, 0,
		14, 15, 1, 0, 0, 0, 15, 20, 1, 0, 0, 0, 16, 18, 5, 2, 0, 0, 17, 19, 3,
		4, 2, 0, 18, 17, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 21, 1, 0, 0, 0, 20,
		16, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 24, 5, 3, 0,
		0, 23, 25, 3, 6, 3, 0, 24, 23, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26,
		1, 0, 0, 0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0, 0, 0, 28, 29, 5, 4, 0, 0,
		29, 3, 1, 0, 0, 0, 30, 31, 5, 5, 0, 0, 31, 5, 1, 0, 0, 0, 32, 33, 5, 6,
		0, 0, 33, 7, 1, 0, 0, 0, 34, 36, 5, 7, 0, 0, 35, 37, 3, 10, 5, 0, 36, 35,
		1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 5, 8, 0, 0,
		39, 9, 1, 0, 0, 0, 40, 41, 5, 4, 0, 0, 41, 11, 1, 0, 0, 0, 5, 14, 18, 20,
		24, 36,
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
	AgpmlParserEOF        = antlr.TokenEOF
	AgpmlParserT__0       = 1
	AgpmlParserT__1       = 2
	AgpmlParserT__2       = 3
	AgpmlParserT__3       = 4
	AgpmlParserT__4       = 5
	AgpmlParserT__5       = 6
	AgpmlParserT__6       = 7
	AgpmlParserT__7       = 8
	AgpmlParserSTRING     = 9
	AgpmlParserDIVIDER    = 10
	AgpmlParserFLAG_TITLE = 11
	AgpmlParserFLAG_OLIST = 12
	AgpmlParserFLAG_ULIST = 13
	AgpmlParserFLAG_QUOTE = 14
	AgpmlParserVARIABLE   = 15
	AgpmlParserTEXT       = 16
	AgpmlParserATRIBUTION = 17
	AgpmlParserCOLOR      = 18
	AgpmlParserNUMBER     = 19
	AgpmlParserDIGIT      = 20
	AgpmlParserCOMMENT    = 21
	AgpmlParserWS         = 22
)

// AgpmlParser rules.
const (
	AgpmlParserRULE_program       = 0
	AgpmlParserRULE_headerConfigs = 1
	AgpmlParserRULE_varConfigs    = 2
	AgpmlParserRULE_dataConfigs   = 3
	AgpmlParserRULE_style         = 4
	AgpmlParserRULE_styleConfig   = 5
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
		p.SetState(12)
		p.Match(AgpmlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__3 {
		{
			p.SetState(13)
			p.HeaderConfigs()
		}

	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__1 {
		{
			p.SetState(16)
			p.Match(AgpmlParserT__1)
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

		if _la == AgpmlParserT__4 {
			{
				p.SetState(17)
				p.VarConfigs()
			}

		}

	}
	{
		p.SetState(22)
		p.Match(AgpmlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AgpmlParserT__5 {
		{
			p.SetState(23)
			p.DataConfigs()
		}

	}
	{
		p.SetState(26)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
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

// IVarConfigsContext is an interface to support dynamic dispatch.
type IVarConfigsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.EnterRule(localctx, 4, AgpmlParserRULE_varConfigs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
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
	p.EnterRule(localctx, 6, AgpmlParserRULE_dataConfigs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(AgpmlParserT__5)
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
	p.EnterRule(localctx, 8, AgpmlParserRULE_style)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(AgpmlParserT__6)
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

	if _la == AgpmlParserT__3 {
		{
			p.SetState(35)
			p.StyleConfig()
		}

	}
	{
		p.SetState(38)
		p.Match(AgpmlParserT__7)
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
	p.EnterRule(localctx, 10, AgpmlParserRULE_styleConfig)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
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
