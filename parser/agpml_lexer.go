// Code generated from .//Agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AgpmlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AgpmlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agpmllexerLexerInit() {
	staticData := &AgpmlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "STRING",
		"DIVIDER", "FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "VARIABLE",
		"TEXT", "ATRIBUTION", "ESC_SEQ", "COLOR", "NUMBER", "DIGIT", "COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 178, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 5, 8, 81, 8, 8, 10, 8, 12, 8, 84, 9, 8, 1, 8, 1, 8, 1, 9,
		4, 9, 89, 8, 9, 11, 9, 12, 9, 90, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		4, 10, 99, 8, 10, 11, 10, 12, 10, 100, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 115, 8, 14, 10,
		14, 12, 14, 118, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 124, 8, 15,
		11, 15, 12, 15, 125, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 142, 8, 18, 1, 19,
		4, 19, 145, 8, 19, 11, 19, 12, 19, 146, 1, 19, 1, 19, 4, 19, 151, 8, 19,
		11, 19, 12, 19, 152, 3, 19, 155, 8, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 5, 21, 163, 8, 21, 10, 21, 12, 21, 166, 9, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 22, 4, 22, 173, 8, 22, 11, 22, 12, 22, 174, 1, 22, 1,
		22, 1, 164, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		0, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 1, 0, 7, 2, 0, 34, 34, 92, 92,
		2, 0, 65, 90, 97, 122, 9, 0, 9, 10, 13, 13, 32, 32, 34, 36, 42, 42, 46,
		46, 49, 49, 62, 62, 92, 92, 3, 0, 10, 10, 13, 13, 36, 36, 3, 0, 48, 57,
		65, 70, 97, 102, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 190, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 55, 1,
		0, 0, 0, 5, 61, 1, 0, 0, 0, 7, 67, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71,
		1, 0, 0, 0, 13, 73, 1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0,
		19, 88, 1, 0, 0, 0, 21, 94, 1, 0, 0, 0, 23, 104, 1, 0, 0, 0, 25, 107, 1,
		0, 0, 0, 27, 109, 1, 0, 0, 0, 29, 111, 1, 0, 0, 0, 31, 119, 1, 0, 0, 0,
		33, 129, 1, 0, 0, 0, 35, 131, 1, 0, 0, 0, 37, 134, 1, 0, 0, 0, 39, 144,
		1, 0, 0, 0, 41, 156, 1, 0, 0, 0, 43, 158, 1, 0, 0, 0, 45, 172, 1, 0, 0,
		0, 47, 48, 5, 46, 0, 0, 48, 49, 5, 104, 0, 0, 49, 50, 5, 101, 0, 0, 50,
		51, 5, 97, 0, 0, 51, 52, 5, 100, 0, 0, 52, 53, 5, 101, 0, 0, 53, 54, 5,
		114, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56, 5, 46, 0, 0, 56, 57, 5, 118, 0, 0,
		57, 58, 5, 97, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60, 5, 115, 0, 0, 60, 4,
		1, 0, 0, 0, 61, 62, 5, 46, 0, 0, 62, 63, 5, 98, 0, 0, 63, 64, 5, 111, 0,
		0, 64, 65, 5, 100, 0, 0, 65, 66, 5, 121, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68,
		5, 49, 0, 0, 68, 8, 1, 0, 0, 0, 69, 70, 5, 50, 0, 0, 70, 10, 1, 0, 0, 0,
		71, 72, 5, 51, 0, 0, 72, 12, 1, 0, 0, 0, 73, 74, 5, 123, 0, 0, 74, 14,
		1, 0, 0, 0, 75, 76, 5, 125, 0, 0, 76, 16, 1, 0, 0, 0, 77, 82, 5, 34, 0,
		0, 78, 81, 3, 35, 17, 0, 79, 81, 8, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 79,
		1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 85, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 34, 0, 0, 86, 18, 1,
		0, 0, 0, 87, 89, 5, 45, 0, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90,
		88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 4, 9, 0,
		0, 93, 20, 1, 0, 0, 0, 94, 95, 5, 35, 0, 0, 95, 96, 5, 91, 0, 0, 96, 98,
		1, 0, 0, 0, 97, 99, 3, 41, 20, 0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102,
		103, 5, 93, 0, 0, 103, 22, 1, 0, 0, 0, 104, 105, 3, 41, 20, 0, 105, 106,
		5, 46, 0, 0, 106, 24, 1, 0, 0, 0, 107, 108, 5, 42, 0, 0, 108, 26, 1, 0,
		0, 0, 109, 110, 5, 62, 0, 0, 110, 28, 1, 0, 0, 0, 111, 116, 5, 36, 0, 0,
		112, 115, 7, 1, 0, 0, 113, 115, 3, 41, 20, 0, 114, 112, 1, 0, 0, 0, 114,
		113, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 30, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 123, 8, 2,
		0, 0, 120, 121, 5, 92, 0, 0, 121, 124, 5, 36, 0, 0, 122, 124, 8, 3, 0,
		0, 123, 120, 1, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128,
		4, 15, 1, 0, 128, 32, 1, 0, 0, 0, 129, 130, 5, 61, 0, 0, 130, 34, 1, 0,
		0, 0, 131, 132, 5, 92, 0, 0, 132, 133, 5, 34, 0, 0, 133, 36, 1, 0, 0, 0,
		134, 135, 5, 35, 0, 0, 135, 136, 7, 4, 0, 0, 136, 137, 7, 4, 0, 0, 137,
		141, 7, 4, 0, 0, 138, 139, 7, 4, 0, 0, 139, 140, 7, 4, 0, 0, 140, 142,
		7, 4, 0, 0, 141, 138, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 38, 1, 0,
		0, 0, 143, 145, 3, 41, 20, 0, 144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0,
		0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 154, 1, 0, 0, 0, 148,
		150, 5, 46, 0, 0, 149, 151, 3, 41, 20, 0, 150, 149, 1, 0, 0, 0, 151, 152,
		1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 155, 1, 0,
		0, 0, 154, 148, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 40, 1, 0, 0, 0,
		156, 157, 7, 5, 0, 0, 157, 42, 1, 0, 0, 0, 158, 159, 5, 47, 0, 0, 159,
		160, 5, 47, 0, 0, 160, 164, 1, 0, 0, 0, 161, 163, 9, 0, 0, 0, 162, 161,
		1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 164, 162, 1, 0,
		0, 0, 165, 167, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 168, 5, 10, 0, 0,
		168, 169, 1, 0, 0, 0, 169, 170, 6, 21, 0, 0, 170, 44, 1, 0, 0, 0, 171,
		173, 7, 6, 0, 0, 172, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 172,
		1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 6, 22,
		0, 0, 177, 46, 1, 0, 0, 0, 15, 0, 80, 82, 90, 100, 114, 116, 123, 125,
		141, 146, 152, 154, 164, 174, 1, 6, 0, 0,
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

// AgpmlLexerInit initializes any static state used to implement AgpmlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAgpmlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgpmlLexerInit() {
	staticData := &AgpmlLexerLexerStaticData
	staticData.once.Do(agpmllexerLexerInit)
}

// NewAgpmlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAgpmlLexer(input antlr.CharStream) *AgpmlLexer {
	AgpmlLexerInit()
	l := new(AgpmlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AgpmlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Agpml.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AgpmlLexer tokens.
const (
	AgpmlLexerT__0       = 1
	AgpmlLexerT__1       = 2
	AgpmlLexerT__2       = 3
	AgpmlLexerT__3       = 4
	AgpmlLexerT__4       = 5
	AgpmlLexerT__5       = 6
	AgpmlLexerT__6       = 7
	AgpmlLexerT__7       = 8
	AgpmlLexerSTRING     = 9
	AgpmlLexerDIVIDER    = 10
	AgpmlLexerFLAG_TITLE = 11
	AgpmlLexerFLAG_OLIST = 12
	AgpmlLexerFLAG_ULIST = 13
	AgpmlLexerFLAG_QUOTE = 14
	AgpmlLexerVARIABLE   = 15
	AgpmlLexerTEXT       = 16
	AgpmlLexerATRIBUTION = 17
	AgpmlLexerCOLOR      = 18
	AgpmlLexerNUMBER     = 19
	AgpmlLexerDIGIT      = 20
	AgpmlLexerCOMMENT    = 21
	AgpmlLexerWS         = 22
)

func (l *AgpmlLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		return l.DIVIDER_Sempred(localctx, predIndex)

	case 15:
		return l.TEXT_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AgpmlLexer) DIVIDER_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetInputStream().LA(1) == '\n'

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AgpmlLexer) TEXT_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.GetInputStream().LA(1) != '\n'

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
