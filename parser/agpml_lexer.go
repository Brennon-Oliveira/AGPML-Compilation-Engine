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
		"T__0", "T__1", "T__2", "BOOLEAN", "TEXT_CONTENT", "STRING", "DIVIDER",
		"FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "ATTRIBUTE",
		"VARIABLE", "ATRIBUTION", "DEFINER", "SEPARATOR", "ESC_SEQ", "COLOR",
		"NUMBER", "DIGIT", "COMMENT", "WS", "OPEN_CURLY_BRACKET", "CLOSE_CURLY_BRACKET",
		"OPEN_CLASS", "CLOSE_CLASS", "OPEN_ID", "CLOSE_ID", "PARAGRAPH",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 236, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 89, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 94, 8, 4, 10, 4, 12, 4, 97, 9,
		4, 1, 4, 3, 4, 100, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 105, 8, 5, 10, 5, 12,
		5, 108, 9, 5, 1, 5, 1, 5, 1, 6, 4, 6, 113, 8, 6, 11, 6, 12, 6, 114, 1,
		7, 1, 7, 1, 7, 1, 7, 4, 7, 121, 8, 7, 11, 7, 12, 7, 122, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 136, 8, 11,
		10, 11, 12, 11, 139, 9, 11, 1, 12, 1, 12, 1, 12, 5, 12, 144, 8, 12, 10,
		12, 12, 12, 147, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 159, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 168, 8, 17, 1, 18, 4, 18, 171, 8, 18, 11, 18,
		12, 18, 172, 1, 18, 1, 18, 4, 18, 177, 8, 18, 11, 18, 12, 18, 178, 3, 18,
		181, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 189, 8, 20,
		10, 20, 12, 20, 192, 9, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 4, 21, 199,
		8, 21, 11, 21, 12, 21, 200, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 214, 8, 24, 10, 24, 12, 24, 217,
		9, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 227,
		8, 26, 10, 26, 12, 26, 230, 9, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		190, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 0, 35, 17, 37,
		18, 39, 19, 41, 20, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55,
		27, 57, 28, 1, 0, 9, 3, 0, 10, 10, 59, 59, 92, 92, 2, 1, 10, 10, 59, 59,
		2, 0, 34, 34, 92, 92, 1, 0, 45, 45, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57,
		65, 70, 97, 102, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 48, 57,
		65, 90, 97, 122, 253, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0,
		39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0,
		0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0,
		0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 67, 1, 0,
		0, 0, 5, 73, 1, 0, 0, 0, 7, 88, 1, 0, 0, 0, 9, 90, 1, 0, 0, 0, 11, 101,
		1, 0, 0, 0, 13, 112, 1, 0, 0, 0, 15, 116, 1, 0, 0, 0, 17, 126, 1, 0, 0,
		0, 19, 129, 1, 0, 0, 0, 21, 131, 1, 0, 0, 0, 23, 133, 1, 0, 0, 0, 25, 140,
		1, 0, 0, 0, 27, 148, 1, 0, 0, 0, 29, 150, 1, 0, 0, 0, 31, 152, 1, 0, 0,
		0, 33, 158, 1, 0, 0, 0, 35, 160, 1, 0, 0, 0, 37, 170, 1, 0, 0, 0, 39, 182,
		1, 0, 0, 0, 41, 184, 1, 0, 0, 0, 43, 198, 1, 0, 0, 0, 45, 204, 1, 0, 0,
		0, 47, 206, 1, 0, 0, 0, 49, 208, 1, 0, 0, 0, 51, 218, 1, 0, 0, 0, 53, 221,
		1, 0, 0, 0, 55, 231, 1, 0, 0, 0, 57, 234, 1, 0, 0, 0, 59, 60, 5, 46, 0,
		0, 60, 61, 5, 104, 0, 0, 61, 62, 5, 101, 0, 0, 62, 63, 5, 97, 0, 0, 63,
		64, 5, 100, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66, 5, 114, 0, 0, 66, 2, 1,
		0, 0, 0, 67, 68, 5, 46, 0, 0, 68, 69, 5, 118, 0, 0, 69, 70, 5, 97, 0, 0,
		70, 71, 5, 114, 0, 0, 71, 72, 5, 115, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74,
		5, 46, 0, 0, 74, 75, 5, 98, 0, 0, 75, 76, 5, 111, 0, 0, 76, 77, 5, 100,
		0, 0, 77, 78, 5, 121, 0, 0, 78, 6, 1, 0, 0, 0, 79, 80, 5, 116, 0, 0, 80,
		81, 5, 114, 0, 0, 81, 82, 5, 117, 0, 0, 82, 89, 5, 101, 0, 0, 83, 84, 5,
		102, 0, 0, 84, 85, 5, 97, 0, 0, 85, 86, 5, 108, 0, 0, 86, 87, 5, 115, 0,
		0, 87, 89, 5, 101, 0, 0, 88, 79, 1, 0, 0, 0, 88, 83, 1, 0, 0, 0, 89, 8,
		1, 0, 0, 0, 90, 95, 5, 59, 0, 0, 91, 94, 3, 33, 16, 0, 92, 94, 8, 0, 0,
		0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93,
		1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0,
		98, 100, 7, 1, 0, 0, 99, 98, 1, 0, 0, 0, 100, 10, 1, 0, 0, 0, 101, 106,
		5, 34, 0, 0, 102, 105, 3, 33, 16, 0, 103, 105, 8, 2, 0, 0, 104, 102, 1,
		0, 0, 0, 104, 103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0,
		0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109,
		110, 5, 34, 0, 0, 110, 12, 1, 0, 0, 0, 111, 113, 7, 3, 0, 0, 112, 111,
		1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 14, 1, 0, 0, 0, 116, 117, 5, 35, 0, 0, 117, 118, 5, 91, 0, 0,
		118, 120, 1, 0, 0, 0, 119, 121, 3, 39, 19, 0, 120, 119, 1, 0, 0, 0, 121,
		122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 124,
		1, 0, 0, 0, 124, 125, 5, 93, 0, 0, 125, 16, 1, 0, 0, 0, 126, 127, 3, 39,
		19, 0, 127, 128, 5, 46, 0, 0, 128, 18, 1, 0, 0, 0, 129, 130, 5, 42, 0,
		0, 130, 20, 1, 0, 0, 0, 131, 132, 5, 62, 0, 0, 132, 22, 1, 0, 0, 0, 133,
		137, 7, 4, 0, 0, 134, 136, 7, 4, 0, 0, 135, 134, 1, 0, 0, 0, 136, 139,
		1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 24, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 145, 5, 36, 0, 0, 141, 144, 7, 4, 0, 0,
		142, 144, 3, 39, 19, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144,
		147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 26, 1,
		0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 61, 0, 0, 149, 28, 1, 0, 0,
		0, 150, 151, 5, 58, 0, 0, 151, 30, 1, 0, 0, 0, 152, 153, 5, 44, 0, 0, 153,
		32, 1, 0, 0, 0, 154, 155, 5, 92, 0, 0, 155, 159, 5, 34, 0, 0, 156, 157,
		5, 92, 0, 0, 157, 159, 5, 59, 0, 0, 158, 154, 1, 0, 0, 0, 158, 156, 1,
		0, 0, 0, 159, 34, 1, 0, 0, 0, 160, 161, 5, 35, 0, 0, 161, 162, 7, 5, 0,
		0, 162, 163, 7, 5, 0, 0, 163, 167, 7, 5, 0, 0, 164, 165, 7, 5, 0, 0, 165,
		166, 7, 5, 0, 0, 166, 168, 7, 5, 0, 0, 167, 164, 1, 0, 0, 0, 167, 168,
		1, 0, 0, 0, 168, 36, 1, 0, 0, 0, 169, 171, 3, 39, 19, 0, 170, 169, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 180, 1, 0, 0, 0, 174, 176, 5, 46, 0, 0, 175, 177, 3, 39, 19, 0, 176,
		175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179,
		1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180, 174, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 38, 1, 0, 0, 0, 182, 183, 7, 6, 0, 0, 183, 40, 1, 0, 0, 0, 184,
		185, 5, 47, 0, 0, 185, 186, 5, 47, 0, 0, 186, 190, 1, 0, 0, 0, 187, 189,
		9, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 190, 188, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0,
		193, 194, 5, 10, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 6, 20, 0, 0, 196,
		42, 1, 0, 0, 0, 197, 199, 7, 7, 0, 0, 198, 197, 1, 0, 0, 0, 199, 200, 1,
		0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0, 0,
		0, 202, 203, 6, 21, 0, 0, 203, 44, 1, 0, 0, 0, 204, 205, 5, 123, 0, 0,
		205, 46, 1, 0, 0, 0, 206, 207, 5, 125, 0, 0, 207, 48, 1, 0, 0, 0, 208,
		209, 5, 46, 0, 0, 209, 210, 5, 46, 0, 0, 210, 211, 1, 0, 0, 0, 211, 215,
		7, 4, 0, 0, 212, 214, 7, 8, 0, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0,
		0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 50, 1, 0, 0, 0,
		217, 215, 1, 0, 0, 0, 218, 219, 5, 46, 0, 0, 219, 220, 5, 46, 0, 0, 220,
		52, 1, 0, 0, 0, 221, 222, 5, 35, 0, 0, 222, 223, 5, 35, 0, 0, 223, 224,
		1, 0, 0, 0, 224, 228, 7, 4, 0, 0, 225, 227, 7, 8, 0, 0, 226, 225, 1, 0,
		0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 54, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 232, 5, 35, 0, 0, 232,
		233, 5, 35, 0, 0, 233, 56, 1, 0, 0, 0, 234, 235, 5, 59, 0, 0, 235, 58,
		1, 0, 0, 0, 21, 0, 88, 93, 95, 99, 104, 106, 114, 122, 137, 143, 145, 158,
		167, 172, 178, 180, 190, 200, 215, 228, 1, 6, 0, 0,
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
	AgpmlLexerBOOLEAN             = 4
	AgpmlLexerTEXT_CONTENT        = 5
	AgpmlLexerSTRING              = 6
	AgpmlLexerDIVIDER             = 7
	AgpmlLexerFLAG_TITLE          = 8
	AgpmlLexerFLAG_OLIST          = 9
	AgpmlLexerFLAG_ULIST          = 10
	AgpmlLexerFLAG_QUOTE          = 11
	AgpmlLexerATTRIBUTE           = 12
	AgpmlLexerVARIABLE            = 13
	AgpmlLexerATRIBUTION          = 14
	AgpmlLexerDEFINER             = 15
	AgpmlLexerSEPARATOR           = 16
	AgpmlLexerCOLOR               = 17
	AgpmlLexerNUMBER              = 18
	AgpmlLexerDIGIT               = 19
	AgpmlLexerCOMMENT             = 20
	AgpmlLexerWS                  = 21
	AgpmlLexerOPEN_CURLY_BRACKET  = 22
	AgpmlLexerCLOSE_CURLY_BRACKET = 23
	AgpmlLexerOPEN_CLASS          = 24
	AgpmlLexerCLOSE_CLASS         = 25
	AgpmlLexerOPEN_ID             = 26
	AgpmlLexerCLOSE_ID            = 27
	AgpmlLexerPARAGRAPH           = 28
)
