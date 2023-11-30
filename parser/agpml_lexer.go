// Code generated from .//agpml.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type agpmlLexer struct {
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TEXT_CONTENT", "CADEIA",
		"FLAG", "DIVIDER", "FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE",
		"VARIABLE", "ATTRIBUTE_NAME", "ATRIBUTION", "ESC_SEQ", "COLOR", "LETTER",
		"DIGIT", "COMMENT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 201, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 4, 6, 76, 8, 6,
		11, 6, 12, 6, 77, 1, 6, 1, 6, 1, 6, 4, 6, 83, 8, 6, 11, 6, 12, 6, 84, 1,
		6, 1, 6, 1, 6, 1, 6, 4, 6, 91, 8, 6, 11, 6, 12, 6, 92, 1, 6, 3, 6, 96,
		8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 101, 8, 7, 10, 7, 12, 7, 104, 9, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 112, 8, 8, 1, 9, 1, 9, 5, 9, 116, 8,
		9, 10, 9, 12, 9, 119, 9, 9, 1, 9, 5, 9, 122, 8, 9, 10, 9, 12, 9, 125, 9,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 134, 8, 10, 11,
		10, 12, 10, 135, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 153, 8, 14, 10,
		14, 12, 14, 156, 9, 14, 1, 15, 1, 15, 1, 15, 5, 15, 161, 8, 15, 10, 15,
		12, 15, 164, 9, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 5,
		18, 173, 8, 18, 10, 18, 12, 18, 176, 9, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 186, 8, 21, 10, 21, 12, 21, 189, 9,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 4, 22, 196, 8, 22, 11, 22, 12, 22,
		197, 1, 22, 1, 22, 1, 187, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 0, 37, 18, 39, 19, 41, 20, 43, 21, 45, 22, 1, 0, 5, 3,
		0, 9, 10, 13, 13, 32, 32, 2, 0, 34, 34, 92, 92, 3, 0, 48, 57, 65, 70, 97,
		102, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 219, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 55, 1, 0, 0, 0, 5, 61, 1,
		0, 0, 0, 7, 67, 1, 0, 0, 0, 9, 69, 1, 0, 0, 0, 11, 71, 1, 0, 0, 0, 13,
		95, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 111, 1, 0, 0, 0, 19, 113, 1, 0,
		0, 0, 21, 128, 1, 0, 0, 0, 23, 139, 1, 0, 0, 0, 25, 143, 1, 0, 0, 0, 27,
		146, 1, 0, 0, 0, 29, 149, 1, 0, 0, 0, 31, 157, 1, 0, 0, 0, 33, 165, 1,
		0, 0, 0, 35, 167, 1, 0, 0, 0, 37, 170, 1, 0, 0, 0, 39, 177, 1, 0, 0, 0,
		41, 179, 1, 0, 0, 0, 43, 181, 1, 0, 0, 0, 45, 195, 1, 0, 0, 0, 47, 48,
		5, 46, 0, 0, 48, 49, 5, 104, 0, 0, 49, 50, 5, 101, 0, 0, 50, 51, 5, 97,
		0, 0, 51, 52, 5, 100, 0, 0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 114, 0, 0,
		54, 2, 1, 0, 0, 0, 55, 56, 5, 46, 0, 0, 56, 57, 5, 118, 0, 0, 57, 58, 5,
		97, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60, 5, 115, 0, 0, 60, 4, 1, 0, 0, 0,
		61, 62, 5, 46, 0, 0, 62, 63, 5, 98, 0, 0, 63, 64, 5, 111, 0, 0, 64, 65,
		5, 100, 0, 0, 65, 66, 5, 121, 0, 0, 66, 6, 1, 0, 0, 0, 67, 68, 5, 49, 0,
		0, 68, 8, 1, 0, 0, 0, 69, 70, 5, 50, 0, 0, 70, 10, 1, 0, 0, 0, 71, 72,
		5, 51, 0, 0, 72, 12, 1, 0, 0, 0, 73, 75, 5, 10, 0, 0, 74, 76, 8, 0, 0,
		0, 75, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 96, 5, 10, 0, 0, 80, 82, 3, 17, 8,
		0, 81, 83, 8, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 5, 10, 0, 0,
		87, 96, 1, 0, 0, 0, 88, 90, 5, 125, 0, 0, 89, 91, 8, 0, 0, 0, 90, 89, 1,
		0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93,
		94, 1, 0, 0, 0, 94, 96, 5, 10, 0, 0, 95, 73, 1, 0, 0, 0, 95, 80, 1, 0,
		0, 0, 95, 88, 1, 0, 0, 0, 96, 14, 1, 0, 0, 0, 97, 102, 5, 34, 0, 0, 98,
		101, 3, 35, 17, 0, 99, 101, 8, 1, 0, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1,
		0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0,
		0, 103, 105, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 34, 0, 0, 106,
		16, 1, 0, 0, 0, 107, 112, 3, 21, 10, 0, 108, 112, 3, 23, 11, 0, 109, 112,
		3, 25, 12, 0, 110, 112, 3, 27, 13, 0, 111, 107, 1, 0, 0, 0, 111, 108, 1,
		0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 18, 1, 0, 0,
		0, 113, 117, 5, 10, 0, 0, 114, 116, 5, 45, 0, 0, 115, 114, 1, 0, 0, 0,
		116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		123, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 122, 3, 45, 22, 0, 121, 120,
		1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 10, 0, 0,
		127, 20, 1, 0, 0, 0, 128, 129, 5, 10, 0, 0, 129, 130, 5, 35, 0, 0, 130,
		131, 5, 91, 0, 0, 131, 133, 1, 0, 0, 0, 132, 134, 3, 41, 20, 0, 133, 132,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0,
		0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 5, 93, 0, 0, 138, 22, 1, 0, 0, 0,
		139, 140, 5, 10, 0, 0, 140, 141, 3, 41, 20, 0, 141, 142, 5, 46, 0, 0, 142,
		24, 1, 0, 0, 0, 143, 144, 5, 10, 0, 0, 144, 145, 5, 42, 0, 0, 145, 26,
		1, 0, 0, 0, 146, 147, 5, 10, 0, 0, 147, 148, 5, 62, 0, 0, 148, 28, 1, 0,
		0, 0, 149, 154, 5, 36, 0, 0, 150, 153, 3, 39, 19, 0, 151, 153, 3, 41, 20,
		0, 152, 150, 1, 0, 0, 0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 30, 1, 0, 0, 0, 156, 154, 1,
		0, 0, 0, 157, 162, 3, 39, 19, 0, 158, 161, 3, 39, 19, 0, 159, 161, 3, 41,
		20, 0, 160, 158, 1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0,
		162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 32, 1, 0, 0, 0, 164, 162,
		1, 0, 0, 0, 165, 166, 5, 61, 0, 0, 166, 34, 1, 0, 0, 0, 167, 168, 5, 92,
		0, 0, 168, 169, 5, 34, 0, 0, 169, 36, 1, 0, 0, 0, 170, 174, 5, 35, 0, 0,
		171, 173, 7, 2, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174,
		172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 38, 1, 0, 0, 0, 176, 174, 1,
		0, 0, 0, 177, 178, 7, 3, 0, 0, 178, 40, 1, 0, 0, 0, 179, 180, 7, 4, 0,
		0, 180, 42, 1, 0, 0, 0, 181, 182, 5, 47, 0, 0, 182, 183, 5, 47, 0, 0, 183,
		187, 1, 0, 0, 0, 184, 186, 9, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 189,
		1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 190, 1, 0,
		0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 5, 10, 0, 0, 191, 192, 1, 0, 0, 0,
		192, 193, 6, 21, 0, 0, 193, 44, 1, 0, 0, 0, 194, 196, 7, 0, 0, 0, 195,
		194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198,
		1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 6, 22, 0, 0, 200, 46, 1, 0,
		0, 0, 18, 0, 77, 84, 92, 95, 100, 102, 111, 117, 123, 135, 152, 154, 160,
		162, 174, 187, 197, 1, 6, 0, 0,
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

// agpmlLexerInit initializes any static state used to implement agpmlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewagpmlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgpmlLexerInit() {
	staticData := &AgpmlLexerLexerStaticData
	staticData.once.Do(agpmllexerLexerInit)
}

// NewagpmlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewagpmlLexer(input antlr.CharStream) *agpmlLexer {
	AgpmlLexerInit()
	l := new(agpmlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AgpmlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "agpml.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// agpmlLexer tokens.
const (
	agpmlLexerT__0           = 1
	agpmlLexerT__1           = 2
	agpmlLexerT__2           = 3
	agpmlLexerT__3           = 4
	agpmlLexerT__4           = 5
	agpmlLexerT__5           = 6
	agpmlLexerTEXT_CONTENT   = 7
	agpmlLexerCADEIA         = 8
	agpmlLexerFLAG           = 9
	agpmlLexerDIVIDER        = 10
	agpmlLexerFLAG_TITLE     = 11
	agpmlLexerFLAG_OLIST     = 12
	agpmlLexerFLAG_ULIST     = 13
	agpmlLexerFLAG_QUOTE     = 14
	agpmlLexerVARIABLE       = 15
	agpmlLexerATTRIBUTE_NAME = 16
	agpmlLexerATRIBUTION     = 17
	agpmlLexerCOLOR          = 18
	agpmlLexerLETTER         = 19
	agpmlLexerDIGIT          = 20
	agpmlLexerCOMMENT        = 21
	agpmlLexerWS             = 22
)