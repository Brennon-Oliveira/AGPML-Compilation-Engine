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
		"T__0", "T__1", "T__2", "T__3", "T__4", "BOOLEAN", "TEXT_CONTENT", "STRING",
		"DIVIDER", "FLAG_TITLE", "FLAG_OLIST", "FLAG_ULIST", "FLAG_QUOTE", "ATTRIBUTE",
		"VARIABLE", "ATRIBUTION", "DEFINER", "SEPARATOR", "ESC_SEQ", "COLOR",
		"NUMBER", "DIGIT", "COMMENT", "WS", "OPEN_CURLY_BRACKET", "CLOSE_CURLY_BRACKET",
		"OPEN_CLASS", "CLOSE_CLASS", "OPEN_ID", "CLOSE_ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 240, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 95, 8, 5, 1, 6, 1, 6, 1, 6,
		5, 6, 100, 8, 6, 10, 6, 12, 6, 103, 9, 6, 1, 6, 3, 6, 106, 8, 6, 1, 7,
		1, 7, 1, 7, 5, 7, 111, 8, 7, 10, 7, 12, 7, 114, 9, 7, 1, 7, 1, 7, 1, 8,
		4, 8, 119, 8, 8, 11, 8, 12, 8, 120, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 127,
		8, 9, 11, 9, 12, 9, 128, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 142, 8, 13, 10, 13, 12, 13, 145, 9,
		13, 1, 14, 1, 14, 1, 14, 5, 14, 150, 8, 14, 10, 14, 12, 14, 153, 9, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 165, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19,
		174, 8, 19, 1, 20, 4, 20, 177, 8, 20, 11, 20, 12, 20, 178, 1, 20, 1, 20,
		4, 20, 183, 8, 20, 11, 20, 12, 20, 184, 3, 20, 187, 8, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 195, 8, 22, 10, 22, 12, 22, 198, 9,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 4, 23, 205, 8, 23, 11, 23, 12, 23,
		206, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 5, 26, 220, 8, 26, 10, 26, 12, 26, 223, 9, 26, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 233, 8, 28, 10, 28, 12,
		28, 236, 9, 28, 1, 29, 1, 29, 1, 29, 1, 196, 0, 30, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 0, 39, 19, 41, 20, 43, 21, 45,
		22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 1, 0, 9, 3,
		0, 10, 10, 59, 59, 92, 92, 2, 1, 10, 10, 59, 59, 2, 0, 34, 34, 92, 92,
		1, 0, 45, 45, 2, 0, 65, 90, 97, 122, 3, 0, 48, 57, 65, 70, 97, 102, 1,
		0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 48, 57, 65, 90, 97, 122,
		257, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0,
		0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0,
		0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 1, 61, 1, 0, 0, 0, 3, 69, 1,
		0, 0, 0, 5, 75, 1, 0, 0, 0, 7, 81, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 94,
		1, 0, 0, 0, 13, 96, 1, 0, 0, 0, 15, 107, 1, 0, 0, 0, 17, 118, 1, 0, 0,
		0, 19, 122, 1, 0, 0, 0, 21, 132, 1, 0, 0, 0, 23, 135, 1, 0, 0, 0, 25, 137,
		1, 0, 0, 0, 27, 139, 1, 0, 0, 0, 29, 146, 1, 0, 0, 0, 31, 154, 1, 0, 0,
		0, 33, 156, 1, 0, 0, 0, 35, 158, 1, 0, 0, 0, 37, 164, 1, 0, 0, 0, 39, 166,
		1, 0, 0, 0, 41, 176, 1, 0, 0, 0, 43, 188, 1, 0, 0, 0, 45, 190, 1, 0, 0,
		0, 47, 204, 1, 0, 0, 0, 49, 210, 1, 0, 0, 0, 51, 212, 1, 0, 0, 0, 53, 214,
		1, 0, 0, 0, 55, 224, 1, 0, 0, 0, 57, 227, 1, 0, 0, 0, 59, 237, 1, 0, 0,
		0, 61, 62, 5, 46, 0, 0, 62, 63, 5, 104, 0, 0, 63, 64, 5, 101, 0, 0, 64,
		65, 5, 97, 0, 0, 65, 66, 5, 100, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5,
		114, 0, 0, 68, 2, 1, 0, 0, 0, 69, 70, 5, 46, 0, 0, 70, 71, 5, 118, 0, 0,
		71, 72, 5, 97, 0, 0, 72, 73, 5, 114, 0, 0, 73, 74, 5, 115, 0, 0, 74, 4,
		1, 0, 0, 0, 75, 76, 5, 46, 0, 0, 76, 77, 5, 98, 0, 0, 77, 78, 5, 111, 0,
		0, 78, 79, 5, 100, 0, 0, 79, 80, 5, 121, 0, 0, 80, 6, 1, 0, 0, 0, 81, 82,
		5, 51, 0, 0, 82, 8, 1, 0, 0, 0, 83, 84, 5, 49, 0, 0, 84, 10, 1, 0, 0, 0,
		85, 86, 5, 116, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 117, 0, 0, 88, 95,
		5, 101, 0, 0, 89, 90, 5, 102, 0, 0, 90, 91, 5, 97, 0, 0, 91, 92, 5, 108,
		0, 0, 92, 93, 5, 115, 0, 0, 93, 95, 5, 101, 0, 0, 94, 85, 1, 0, 0, 0, 94,
		89, 1, 0, 0, 0, 95, 12, 1, 0, 0, 0, 96, 101, 5, 59, 0, 0, 97, 100, 3, 37,
		18, 0, 98, 100, 8, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100,
		103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 105, 1,
		0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 106, 7, 1, 0, 0, 105, 104, 1, 0, 0,
		0, 106, 14, 1, 0, 0, 0, 107, 112, 5, 34, 0, 0, 108, 111, 3, 37, 18, 0,
		109, 111, 8, 2, 0, 0, 110, 108, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111,
		114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 115,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 34, 0, 0, 116, 16, 1, 0,
		0, 0, 117, 119, 7, 3, 0, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0,
		120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 18, 1, 0, 0, 0, 122, 123,
		5, 35, 0, 0, 123, 124, 5, 91, 0, 0, 124, 126, 1, 0, 0, 0, 125, 127, 3,
		43, 21, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0,
		0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 5, 93, 0, 0,
		131, 20, 1, 0, 0, 0, 132, 133, 3, 43, 21, 0, 133, 134, 5, 46, 0, 0, 134,
		22, 1, 0, 0, 0, 135, 136, 5, 42, 0, 0, 136, 24, 1, 0, 0, 0, 137, 138, 5,
		62, 0, 0, 138, 26, 1, 0, 0, 0, 139, 143, 7, 4, 0, 0, 140, 142, 7, 4, 0,
		0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 28, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 151, 5,
		36, 0, 0, 147, 150, 7, 4, 0, 0, 148, 150, 3, 43, 21, 0, 149, 147, 1, 0,
		0, 0, 149, 148, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 30, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 155,
		5, 61, 0, 0, 155, 32, 1, 0, 0, 0, 156, 157, 5, 58, 0, 0, 157, 34, 1, 0,
		0, 0, 158, 159, 5, 44, 0, 0, 159, 36, 1, 0, 0, 0, 160, 161, 5, 92, 0, 0,
		161, 165, 5, 34, 0, 0, 162, 163, 5, 92, 0, 0, 163, 165, 5, 59, 0, 0, 164,
		160, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 38, 1, 0, 0, 0, 166, 167, 5,
		35, 0, 0, 167, 168, 7, 5, 0, 0, 168, 169, 7, 5, 0, 0, 169, 173, 7, 5, 0,
		0, 170, 171, 7, 5, 0, 0, 171, 172, 7, 5, 0, 0, 172, 174, 7, 5, 0, 0, 173,
		170, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 40, 1, 0, 0, 0, 175, 177, 3,
		43, 21, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 186, 1, 0, 0, 0, 180, 182, 5, 46, 0, 0,
		181, 183, 3, 43, 21, 0, 182, 181, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 187, 1, 0, 0, 0, 186, 180,
		1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 42, 1, 0, 0, 0, 188, 189, 7, 6,
		0, 0, 189, 44, 1, 0, 0, 0, 190, 191, 5, 47, 0, 0, 191, 192, 5, 47, 0, 0,
		192, 196, 1, 0, 0, 0, 193, 195, 9, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195,
		198, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 199,
		1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 5, 10, 0, 0, 200, 201, 1, 0,
		0, 0, 201, 202, 6, 22, 0, 0, 202, 46, 1, 0, 0, 0, 203, 205, 7, 7, 0, 0,
		204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 6, 23, 0, 0, 209, 48,
		1, 0, 0, 0, 210, 211, 5, 123, 0, 0, 211, 50, 1, 0, 0, 0, 212, 213, 5, 125,
		0, 0, 213, 52, 1, 0, 0, 0, 214, 215, 5, 46, 0, 0, 215, 216, 5, 46, 0, 0,
		216, 217, 1, 0, 0, 0, 217, 221, 7, 4, 0, 0, 218, 220, 7, 8, 0, 0, 219,
		218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222,
		1, 0, 0, 0, 222, 54, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 225, 5, 46,
		0, 0, 225, 226, 5, 46, 0, 0, 226, 56, 1, 0, 0, 0, 227, 228, 5, 35, 0, 0,
		228, 229, 5, 35, 0, 0, 229, 230, 1, 0, 0, 0, 230, 234, 7, 4, 0, 0, 231,
		233, 7, 8, 0, 0, 232, 231, 1, 0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 58, 1, 0, 0, 0, 236, 234, 1, 0,
		0, 0, 237, 238, 5, 35, 0, 0, 238, 239, 5, 35, 0, 0, 239, 60, 1, 0, 0, 0,
		21, 0, 94, 99, 101, 105, 110, 112, 120, 128, 143, 149, 151, 164, 173, 178,
		184, 186, 196, 206, 221, 234, 1, 6, 0, 0,
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
	AgpmlLexerBOOLEAN             = 6
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
	AgpmlLexerOPEN_CLASS          = 26
	AgpmlLexerCLOSE_CLASS         = 27
	AgpmlLexerOPEN_ID             = 28
	AgpmlLexerCLOSE_ID            = 29
)
