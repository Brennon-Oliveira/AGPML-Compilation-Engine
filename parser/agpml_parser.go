// Code generated from .//agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // agpml

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

type agpmlParser struct {
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
		"", "'.header'", "'.vars'", "'.body'", "'1'", "'2'", "'3'", "", "",
		"", "", "", "", "'\\n*'", "'\\n>'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "TEXT_CONTENT", "CADEIA", "FLAG", "DIVIDER",
		"FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "VARIABLE",
		"ATTRIBUTE_NAME", "ATRIBUTION", "COLOR", "LETTER", "DIGIT", "COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "headerConfigs", "varConfigs", "dataConfigs",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 31, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 3, 0, 11, 8, 0, 1, 0, 1, 0, 3, 0, 15, 8, 0, 3, 0, 17, 8, 0, 1, 0, 1,
		0, 3, 0, 21, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		0, 0, 4, 0, 2, 4, 6, 0, 0, 30, 0, 8, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4,
		26, 1, 0, 0, 0, 6, 28, 1, 0, 0, 0, 8, 10, 5, 1, 0, 0, 9, 11, 3, 2, 1, 0,
		10, 9, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 16, 1, 0, 0, 0, 12, 14, 5, 2,
		0, 0, 13, 15, 3, 4, 2, 0, 14, 13, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 17,
		1, 0, 0, 0, 16, 12, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0,
		18, 20, 5, 3, 0, 0, 19, 21, 3, 6, 3, 0, 20, 19, 1, 0, 0, 0, 20, 21, 1,
		0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 5, 0, 0, 1, 23, 1, 1, 0, 0, 0, 24,
		25, 5, 4, 0, 0, 25, 3, 1, 0, 0, 0, 26, 27, 5, 5, 0, 0, 27, 5, 1, 0, 0,
		0, 28, 29, 5, 6, 0, 0, 29, 7, 1, 0, 0, 0, 4, 10, 14, 16, 20,
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

// agpmlParserInit initializes any static state used to implement agpmlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewagpmlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgpmlParserInit() {
	staticData := &AgpmlParserStaticData
	staticData.once.Do(agpmlParserInit)
}

// NewagpmlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewagpmlParser(input antlr.TokenStream) *agpmlParser {
	AgpmlParserInit()
	this := new(agpmlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AgpmlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "agpml.g4"

	return this
}

// agpmlParser tokens.
const (
	agpmlParserEOF            = antlr.TokenEOF
	agpmlParserT__0           = 1
	agpmlParserT__1           = 2
	agpmlParserT__2           = 3
	agpmlParserT__3           = 4
	agpmlParserT__4           = 5
	agpmlParserT__5           = 6
	agpmlParserTEXT_CONTENT   = 7
	agpmlParserCADEIA         = 8
	agpmlParserFLAG           = 9
	agpmlParserDIVIDER        = 10
	agpmlParserFLAG_TITLE     = 11
	agpmlParserFLAG_OLIST     = 12
	agpmlParserFLAG_ULIST     = 13
	agpmlParserFLAG_QUOTE     = 14
	agpmlParserVARIABLE       = 15
	agpmlParserATTRIBUTE_NAME = 16
	agpmlParserATRIBUTION     = 17
	agpmlParserCOLOR          = 18
	agpmlParserLETTER         = 19
	agpmlParserDIGIT          = 20
	agpmlParserCOMMENT        = 21
	agpmlParserWS             = 22
)

// agpmlParser rules.
const (
	agpmlParserRULE_program       = 0
	agpmlParserRULE_headerConfigs = 1
	agpmlParserRULE_varConfigs    = 2
	agpmlParserRULE_dataConfigs   = 3
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
	p.RuleIndex = agpmlParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = agpmlParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = agpmlParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(agpmlParserEOF, 0)
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
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case agpmlVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *agpmlParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, agpmlParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Match(agpmlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == agpmlParserT__3 {
		{
			p.SetState(9)
			p.HeaderConfigs()
		}

	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == agpmlParserT__1 {
		{
			p.SetState(12)
			p.Match(agpmlParserT__1)
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

		if _la == agpmlParserT__4 {
			{
				p.SetState(13)
				p.VarConfigs()
			}

		}

	}
	{
		p.SetState(18)
		p.Match(agpmlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == agpmlParserT__5 {
		{
			p.SetState(19)
			p.DataConfigs()
		}

	}
	{
		p.SetState(22)
		p.Match(agpmlParserEOF)
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
	p.RuleIndex = agpmlParserRULE_headerConfigs
	return p
}

func InitEmptyHeaderConfigsContext(p *HeaderConfigsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = agpmlParserRULE_headerConfigs
}

func (*HeaderConfigsContext) IsHeaderConfigsContext() {}

func NewHeaderConfigsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderConfigsContext {
	var p = new(HeaderConfigsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = agpmlParserRULE_headerConfigs

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
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.EnterHeaderConfigs(s)
	}
}

func (s *HeaderConfigsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.ExitHeaderConfigs(s)
	}
}

func (s *HeaderConfigsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case agpmlVisitor:
		return t.VisitHeaderConfigs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *agpmlParser) HeaderConfigs() (localctx IHeaderConfigsContext) {
	localctx = NewHeaderConfigsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, agpmlParserRULE_headerConfigs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(agpmlParserT__3)
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
	p.RuleIndex = agpmlParserRULE_varConfigs
	return p
}

func InitEmptyVarConfigsContext(p *VarConfigsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = agpmlParserRULE_varConfigs
}

func (*VarConfigsContext) IsVarConfigsContext() {}

func NewVarConfigsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarConfigsContext {
	var p = new(VarConfigsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = agpmlParserRULE_varConfigs

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
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.EnterVarConfigs(s)
	}
}

func (s *VarConfigsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.ExitVarConfigs(s)
	}
}

func (s *VarConfigsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case agpmlVisitor:
		return t.VisitVarConfigs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *agpmlParser) VarConfigs() (localctx IVarConfigsContext) {
	localctx = NewVarConfigsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, agpmlParserRULE_varConfigs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(agpmlParserT__4)
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
	p.RuleIndex = agpmlParserRULE_dataConfigs
	return p
}

func InitEmptyDataConfigsContext(p *DataConfigsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = agpmlParserRULE_dataConfigs
}

func (*DataConfigsContext) IsDataConfigsContext() {}

func NewDataConfigsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataConfigsContext {
	var p = new(DataConfigsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = agpmlParserRULE_dataConfigs

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
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.EnterDataConfigs(s)
	}
}

func (s *DataConfigsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(agpmlListener); ok {
		listenerT.ExitDataConfigs(s)
	}
}

func (s *DataConfigsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case agpmlVisitor:
		return t.VisitDataConfigs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *agpmlParser) DataConfigs() (localctx IDataConfigsContext) {
	localctx = NewDataConfigsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, agpmlParserRULE_dataConfigs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(agpmlParserT__5)
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
