// Code generated from /Users/anqiansong/goland/go/go-zero/anqiansong/mock/grammar/MockLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 30, 222,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3,
	2, 3, 2, 5, 2, 80, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 5, 7, 105, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24,
	6, 24, 144, 10, 24, 13, 24, 14, 24, 145, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 7, 25, 154, 10, 25, 12, 25, 14, 25, 157, 11, 25, 3, 26, 5, 26,
	160, 10, 26, 3, 26, 6, 26, 163, 10, 26, 13, 26, 14, 26, 164, 3, 27, 6,
	27, 168, 10, 27, 13, 27, 14, 27, 169, 3, 27, 5, 27, 173, 10, 27, 6, 27,
	175, 10, 27, 13, 27, 14, 27, 176, 3, 28, 3, 28, 6, 28, 181, 10, 28, 13,
	28, 14, 28, 182, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 2, 2, 39,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 2, 59,
	2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 30, 3, 2, 7, 5,
	2, 11, 12, 15, 15, 34, 34, 7, 2, 12, 12, 15, 15, 36, 36, 125, 125, 127,
	127, 3, 2, 50, 59, 3, 2, 99, 124, 4, 2, 67, 92, 99, 124, 2, 225, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2,
	2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3,
	2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 5, 81, 3, 2, 2, 2, 7, 86, 3, 2, 2, 2, 9,
	92, 3, 2, 2, 2, 11, 96, 3, 2, 2, 2, 13, 104, 3, 2, 2, 2, 15, 106, 3, 2,
	2, 2, 17, 108, 3, 2, 2, 2, 19, 110, 3, 2, 2, 2, 21, 112, 3, 2, 2, 2, 23,
	114, 3, 2, 2, 2, 25, 116, 3, 2, 2, 2, 27, 118, 3, 2, 2, 2, 29, 120, 3,
	2, 2, 2, 31, 122, 3, 2, 2, 2, 33, 124, 3, 2, 2, 2, 35, 126, 3, 2, 2, 2,
	37, 129, 3, 2, 2, 2, 39, 132, 3, 2, 2, 2, 41, 134, 3, 2, 2, 2, 43, 137,
	3, 2, 2, 2, 45, 140, 3, 2, 2, 2, 47, 143, 3, 2, 2, 2, 49, 149, 3, 2, 2,
	2, 51, 159, 3, 2, 2, 2, 53, 174, 3, 2, 2, 2, 55, 178, 3, 2, 2, 2, 57, 186,
	3, 2, 2, 2, 59, 189, 3, 2, 2, 2, 61, 192, 3, 2, 2, 2, 63, 199, 3, 2, 2,
	2, 65, 204, 3, 2, 2, 2, 67, 210, 3, 2, 2, 2, 69, 212, 3, 2, 2, 2, 71, 214,
	3, 2, 2, 2, 73, 216, 3, 2, 2, 2, 75, 218, 3, 2, 2, 2, 77, 80, 5, 63, 32,
	2, 78, 80, 5, 65, 33, 2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 4,
	3, 2, 2, 2, 81, 82, 7, 100, 2, 2, 82, 83, 7, 113, 2, 2, 83, 84, 7, 113,
	2, 2, 84, 85, 7, 110, 2, 2, 85, 6, 3, 2, 2, 2, 86, 87, 7, 117, 2, 2, 87,
	88, 7, 110, 2, 2, 88, 89, 7, 107, 2, 2, 89, 90, 7, 101, 2, 2, 90, 91, 7,
	103, 2, 2, 91, 8, 3, 2, 2, 2, 92, 93, 7, 111, 2, 2, 93, 94, 7, 99, 2, 2,
	94, 95, 7, 114, 2, 2, 95, 10, 3, 2, 2, 2, 96, 97, 7, 117, 2, 2, 97, 98,
	7, 114, 2, 2, 98, 99, 7, 103, 2, 2, 99, 100, 7, 101, 2, 2, 100, 12, 3,
	2, 2, 2, 101, 105, 5, 57, 29, 2, 102, 105, 5, 59, 30, 2, 103, 105, 5, 61,
	31, 2, 104, 101, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 103, 3, 2, 2, 2,
	105, 14, 3, 2, 2, 2, 106, 107, 7, 42, 2, 2, 107, 16, 3, 2, 2, 2, 108, 109,
	7, 43, 2, 2, 109, 18, 3, 2, 2, 2, 110, 111, 7, 125, 2, 2, 111, 20, 3, 2,
	2, 2, 112, 113, 7, 127, 2, 2, 113, 22, 3, 2, 2, 2, 114, 115, 7, 93, 2,
	2, 115, 24, 3, 2, 2, 2, 116, 117, 7, 95, 2, 2, 117, 26, 3, 2, 2, 2, 118,
	119, 7, 48, 2, 2, 119, 28, 3, 2, 2, 2, 120, 121, 7, 61, 2, 2, 121, 30,
	3, 2, 2, 2, 122, 123, 7, 46, 2, 2, 123, 32, 3, 2, 2, 2, 124, 125, 7, 63,
	2, 2, 125, 34, 3, 2, 2, 2, 126, 127, 7, 62, 2, 2, 127, 128, 7, 62, 2, 2,
	128, 36, 3, 2, 2, 2, 129, 130, 7, 64, 2, 2, 130, 131, 7, 64, 2, 2, 131,
	38, 3, 2, 2, 2, 132, 133, 7, 37, 2, 2, 133, 40, 3, 2, 2, 2, 134, 135, 7,
	47, 2, 2, 135, 136, 7, 64, 2, 2, 136, 42, 3, 2, 2, 2, 137, 138, 7, 62,
	2, 2, 138, 139, 7, 47, 2, 2, 139, 44, 3, 2, 2, 2, 140, 141, 7, 47, 2, 2,
	141, 46, 3, 2, 2, 2, 142, 144, 9, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 145,
	3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 3, 2,
	2, 2, 147, 148, 8, 24, 2, 2, 148, 48, 3, 2, 2, 2, 149, 155, 5, 69, 35,
	2, 150, 154, 5, 73, 37, 2, 151, 154, 5, 67, 34, 2, 152, 154, 5, 71, 36,
	2, 153, 150, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154,
	157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 50, 3,
	2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 160, 7, 47, 2, 2, 159, 158, 3, 2, 2,
	2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 163, 5, 67, 34, 2,
	162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164,
	165, 3, 2, 2, 2, 165, 52, 3, 2, 2, 2, 166, 168, 5, 69, 35, 2, 167, 166,
	3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2,
	2, 2, 170, 172, 3, 2, 2, 2, 171, 173, 5, 27, 14, 2, 172, 171, 3, 2, 2,
	2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174, 167, 3, 2, 2, 2, 175,
	176, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 54, 3,
	2, 2, 2, 178, 180, 7, 125, 2, 2, 179, 181, 10, 3, 2, 2, 180, 179, 3, 2,
	2, 2, 181, 182, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2,
	183, 184, 3, 2, 2, 2, 184, 185, 7, 127, 2, 2, 185, 56, 3, 2, 2, 2, 186,
	187, 7, 124, 2, 2, 187, 188, 7, 106, 2, 2, 188, 58, 3, 2, 2, 2, 189, 190,
	7, 103, 2, 2, 190, 191, 7, 112, 2, 2, 191, 60, 3, 2, 2, 2, 192, 193, 7,
	112, 2, 2, 193, 194, 7, 119, 2, 2, 194, 195, 7, 111, 2, 2, 195, 196, 7,
	100, 2, 2, 196, 197, 7, 103, 2, 2, 197, 198, 7, 116, 2, 2, 198, 62, 3,
	2, 2, 2, 199, 200, 7, 118, 2, 2, 200, 201, 7, 116, 2, 2, 201, 202, 7, 119,
	2, 2, 202, 203, 7, 103, 2, 2, 203, 64, 3, 2, 2, 2, 204, 205, 7, 104, 2,
	2, 205, 206, 7, 99, 2, 2, 206, 207, 7, 110, 2, 2, 207, 208, 7, 117, 2,
	2, 208, 209, 7, 103, 2, 2, 209, 66, 3, 2, 2, 2, 210, 211, 9, 4, 2, 2, 211,
	68, 3, 2, 2, 2, 212, 213, 9, 5, 2, 2, 213, 70, 3, 2, 2, 2, 214, 215, 7,
	97, 2, 2, 215, 72, 3, 2, 2, 2, 216, 217, 9, 6, 2, 2, 217, 74, 3, 2, 2,
	2, 218, 219, 11, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 8, 38, 2, 2,
	221, 76, 3, 2, 2, 2, 14, 2, 79, 104, 145, 153, 155, 159, 164, 169, 172,
	176, 182, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'bool'", "'slice'", "'map'", "'spec'", "", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "'.'", "';'", "','", "'='", "'<<'", "'>>'", "'#'",
	"'->'", "'<-'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "BOOL_VALUE", "BOOL", "SLICE", "MAP", "SPEC", "LANG", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LBRACK", "RBRACK", "DOT", "SMICOLON", "COMMA", "ASSIGN",
	"QUOTE_LEFT", "QUOTE_RIGHT", "SHARP", "TO", "PUT", "BAR", "WS", "IDENT",
	"NUMBER", "PACKAGE", "VALUE", "ERRCHAR",
}

var lexerRuleNames = []string{
	"BOOL_VALUE", "BOOL", "SLICE", "MAP", "SPEC", "LANG", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "LBRACK", "RBRACK", "DOT", "SMICOLON", "COMMA", "ASSIGN",
	"QUOTE_LEFT", "QUOTE_RIGHT", "SHARP", "TO", "PUT", "BAR", "WS", "IDENT",
	"NUMBER", "PACKAGE", "VALUE", "LANG_ZH", "LANG_EN", "LANG_NUMBER", "BOOL_TRUE",
	"BOOL_FALSE", "DIGIT", "LOWER_CASE", "UNDERSCORE", "ALPHA", "ERRCHAR",
}

type MockLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMockLexer(input antlr.CharStream) *MockLexer {

	l := new(MockLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MockLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MockLexer tokens.
const (
	MockLexerBOOL_VALUE  = 1
	MockLexerBOOL        = 2
	MockLexerSLICE       = 3
	MockLexerMAP         = 4
	MockLexerSPEC        = 5
	MockLexerLANG        = 6
	MockLexerLPAREN      = 7
	MockLexerRPAREN      = 8
	MockLexerLBRACE      = 9
	MockLexerRBRACE      = 10
	MockLexerLBRACK      = 11
	MockLexerRBRACK      = 12
	MockLexerDOT         = 13
	MockLexerSMICOLON    = 14
	MockLexerCOMMA       = 15
	MockLexerASSIGN      = 16
	MockLexerQUOTE_LEFT  = 17
	MockLexerQUOTE_RIGHT = 18
	MockLexerSHARP       = 19
	MockLexerTO          = 20
	MockLexerPUT         = 21
	MockLexerBAR         = 22
	MockLexerWS          = 23
	MockLexerIDENT       = 24
	MockLexerNUMBER      = 25
	MockLexerPACKAGE     = 26
	MockLexerVALUE       = 27
	MockLexerERRCHAR     = 28
)
