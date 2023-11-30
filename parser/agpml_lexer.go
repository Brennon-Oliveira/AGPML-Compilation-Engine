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
		"", "'.header'", "'.vars'", "'.body'", "'1'", "'2'", "'3'", "", "",
		"", "", "", "'*'", "'>'", "", "", "'='", "':'", "','", "", "", "", "",
		"", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "TEXT_CONTENT", "STRING", "DIVIDER", "FLAG_TITLE",
		"FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "ATTRIBUTE", "VARIABLE", "ATRIBUTION",
		"DEFINER", "SEPARATOR", "COLOR", "NUMBER", "DIGIT", "COMMENT", "WS",
		"OPEN_CURLY_BRACKET", "CLOSE_CURLY_BRACKET",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "TEXT_CONTENT", "STRING",
		"DIVIDER", "FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "ATTRIBUTE",
		"VARIABLE", "ATRIBUTION", "DEFINER", "SEPARATOR", "ESC_SEQ", "COLOR",
		"NUMBER", "DIGIT", "COMMENT", "WS", "OPEN_CURLY_BRACKET", "CLOSE_CURLY_BRACKET",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 25, 197, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 83, 8, 6, 10, 6, 12, 6, 86, 9, 6, 1,
		6, 3, 6, 89, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 94, 8, 7, 10, 7, 12, 7, 97,
		9, 7, 1, 7, 1, 7, 1, 8, 4, 8, 102, 8, 8, 11, 8, 12, 8, 103, 1, 9, 1, 9,
		1, 9, 1, 9, 4, 9, 110, 8, 9, 11, 9, 12, 9, 111, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 125, 8, 13, 10,
		13, 12, 13, 128, 9, 13, 1, 14, 1, 14, 1, 14, 5, 14, 133, 8, 14, 10, 14,
		12, 14, 136, 9, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 148, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 3, 19, 157, 8, 19, 1, 20, 4, 20, 160, 8, 20, 11, 20, 12,
		20, 161, 1, 20, 1, 20, 4, 20, 166, 8, 20, 11, 20, 12, 20, 167, 3, 20, 170,
		8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 178, 8, 22, 10,
		22, 12, 22, 181, 9, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 4, 23, 188,
		8, 23, 11, 23, 12, 23, 189, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		179, 0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		0, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 1, 0, 8, 3,
		0, 10, 10, 59, 59, 92, 92, 2, 1, 10, 10, 59, 59, 2, 0, 34, 34, 92, 92,
		1, 0, 45, 45, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 70, 97, 102, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 211, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 1, 53, 1, 0, 0, 0, 3, 61, 1, 0, 0, 0, 5, 67, 1, 0, 0, 0, 7, 73, 1,
		0, 0, 0, 9, 75, 1, 0, 0, 0, 11, 77, 1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15,
		90, 1, 0, 0, 0, 17, 101, 1, 0, 0, 0, 19, 105, 1, 0, 0, 0, 21, 115, 1, 0,
		0, 0, 23, 118, 1, 0, 0, 0, 25, 120, 1, 0, 0, 0, 27, 122, 1, 0, 0, 0, 29,
		129, 1, 0, 0, 0, 31, 137, 1, 0, 0, 0, 33, 139, 1, 0, 0, 0, 35, 141, 1,
		0, 0, 0, 37, 147, 1, 0, 0, 0, 39, 149, 1, 0, 0, 0, 41, 159, 1, 0, 0, 0,
		43, 171, 1, 0, 0, 0, 45, 173, 1, 0, 0, 0, 47, 187, 1, 0, 0, 0, 49, 193,
		1, 0, 0, 0, 51, 195, 1, 0, 0, 0, 53, 54, 5, 46, 0, 0, 54, 55, 5, 104, 0,
		0, 55, 56, 5, 101, 0, 0, 56, 57, 5, 97, 0, 0, 57, 58, 5, 100, 0, 0, 58,
		59, 5, 101, 0, 0, 59, 60, 5, 114, 0, 0, 60, 2, 1, 0, 0, 0, 61, 62, 5, 46,
		0, 0, 62, 63, 5, 118, 0, 0, 63, 64, 5, 97, 0, 0, 64, 65, 5, 114, 0, 0,
		65, 66, 5, 115, 0, 0, 66, 4, 1, 0, 0, 0, 67, 68, 5, 46, 0, 0, 68, 69, 5,
		98, 0, 0, 69, 70, 5, 111, 0, 0, 70, 71, 5, 100, 0, 0, 71, 72, 5, 121, 0,
		0, 72, 6, 1, 0, 0, 0, 73, 74, 5, 49, 0, 0, 74, 8, 1, 0, 0, 0, 75, 76, 5,
		50, 0, 0, 76, 10, 1, 0, 0, 0, 77, 78, 5, 51, 0, 0, 78, 12, 1, 0, 0, 0,
		79, 84, 5, 59, 0, 0, 80, 83, 3, 37, 18, 0, 81, 83, 8, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0,
		84, 85, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89, 7,
		1, 0, 0, 88, 87, 1, 0, 0, 0, 89, 14, 1, 0, 0, 0, 90, 95, 5, 34, 0, 0, 91,
		94, 3, 37, 18, 0, 92, 94, 8, 2, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0,
		0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98,
		1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 99, 5, 34, 0, 0, 99, 16, 1, 0, 0, 0,
		100, 102, 7, 3, 0, 0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106, 5,
		35, 0, 0, 106, 107, 5, 91, 0, 0, 107, 109, 1, 0, 0, 0, 108, 110, 3, 43,
		21, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 5, 93, 0, 0, 114,
		20, 1, 0, 0, 0, 115, 116, 3, 43, 21, 0, 116, 117, 5, 46, 0, 0, 117, 22,
		1, 0, 0, 0, 118, 119, 5, 42, 0, 0, 119, 24, 1, 0, 0, 0, 120, 121, 5, 62,
		0, 0, 121, 26, 1, 0, 0, 0, 122, 126, 7, 4, 0, 0, 123, 125, 7, 4, 0, 0,
		124, 123, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126,
		127, 1, 0, 0, 0, 127, 28, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 134, 5,
		36, 0, 0, 130, 133, 7, 4, 0, 0, 131, 133, 3, 43, 21, 0, 132, 130, 1, 0,
		0, 0, 132, 131, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 30, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 138,
		5, 61, 0, 0, 138, 32, 1, 0, 0, 0, 139, 140, 5, 58, 0, 0, 140, 34, 1, 0,
		0, 0, 141, 142, 5, 44, 0, 0, 142, 36, 1, 0, 0, 0, 143, 144, 5, 92, 0, 0,
		144, 148, 5, 34, 0, 0, 145, 146, 5, 92, 0, 0, 146, 148, 5, 59, 0, 0, 147,
		143, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 38, 1, 0, 0, 0, 149, 150, 5,
		35, 0, 0, 150, 151, 7, 5, 0, 0, 151, 152, 7, 5, 0, 0, 152, 156, 7, 5, 0,
		0, 153, 154, 7, 5, 0, 0, 154, 155, 7, 5, 0, 0, 155, 157, 7, 5, 0, 0, 156,
		153, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 40, 1, 0, 0, 0, 158, 160, 3,
		43, 21, 0, 159, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 159, 1, 0,
		0, 0, 161, 162, 1, 0, 0, 0, 162, 169, 1, 0, 0, 0, 163, 165, 5, 46, 0, 0,
		164, 166, 3, 43, 21, 0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 163,
		1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 42, 1, 0, 0, 0, 171, 172, 7, 6,
		0, 0, 172, 44, 1, 0, 0, 0, 173, 174, 5, 47, 0, 0, 174, 175, 5, 47, 0, 0,
		175, 179, 1, 0, 0, 0, 176, 178, 9, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178,
		181, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 182,
		1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 183, 5, 10, 0, 0, 183, 184, 1, 0,
		0, 0, 184, 185, 6, 22, 0, 0, 185, 46, 1, 0, 0, 0, 186, 188, 7, 7, 0, 0,
		187, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189,
		190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 192, 6, 23, 0, 0, 192, 48,
		1, 0, 0, 0, 193, 194, 5, 123, 0, 0, 194, 50, 1, 0, 0, 0, 195, 196, 5, 125,
		0, 0, 196, 52, 1, 0, 0, 0, 18, 0, 82, 84, 88, 93, 95, 103, 111, 126, 132,
		134, 147, 156, 161, 167, 169, 179, 189, 1, 6, 0, 0,
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
	AgpmlLexerT__0                = 1
	AgpmlLexerT__1                = 2
	AgpmlLexerT__2                = 3
	AgpmlLexerT__3                = 4
	AgpmlLexerT__4                = 5
	AgpmlLexerT__5                = 6
	AgpmlLexerTEXT_CONTENT        = 7
	AgpmlLexerSTRING              = 8
	AgpmlLexerDIVIDER             = 9
	AgpmlLexerFLAG_TITLE          = 10
	AgpmlLexerFLAG_OLIST          = 11
	AgpmlLexerFLAG_ULIST          = 12
	AgpmlLexerFLAG_QUOTE          = 13
	AgpmlLexerATTRIBUTE           = 14
	AgpmlLexerVARIABLE            = 15
	AgpmlLexerATRIBUTION          = 16
	AgpmlLexerDEFINER             = 17
	AgpmlLexerSEPARATOR           = 18
	AgpmlLexerCOLOR               = 19
	AgpmlLexerNUMBER              = 20
	AgpmlLexerDIGIT               = 21
	AgpmlLexerCOMMENT             = 22
	AgpmlLexerWS                  = 23
	AgpmlLexerOPEN_CURLY_BRACKET  = 24
	AgpmlLexerCLOSE_CURLY_BRACKET = 25
)
