// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 205,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 7, 13, 96, 10, 13, 12, 13, 14, 13, 99, 11, 13,
	3, 14, 6, 14, 102, 10, 14, 13, 14, 14, 14, 103, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 7, 15, 112, 10, 15, 12, 15, 14, 15, 115, 11, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 126, 10,
	16, 12, 16, 14, 16, 129, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7,
	17, 136, 10, 17, 12, 17, 14, 17, 139, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 5, 18, 146, 10, 18, 3, 18, 6, 18, 149, 10, 18, 13, 18, 14, 18, 150,
	3, 18, 5, 18, 154, 10, 18, 5, 18, 156, 10, 18, 3, 18, 5, 18, 159, 10, 18,
	3, 19, 3, 19, 7, 19, 163, 10, 19, 12, 19, 14, 19, 166, 11, 19, 3, 19, 5,
	19, 169, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 177,
	10, 21, 3, 21, 5, 21, 180, 10, 21, 3, 21, 3, 21, 3, 21, 6, 21, 185, 10,
	21, 13, 21, 14, 21, 186, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 194,
	10, 21, 3, 22, 3, 22, 5, 22, 198, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5,
	23, 204, 10, 23, 3, 113, 2, 24, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 3, 2, 17, 5, 2, 11, 12,
	14, 15, 34, 34, 4, 2, 12, 12, 15, 15, 6, 2, 12, 12, 15, 15, 36, 36, 94,
	94, 3, 2, 51, 59, 4, 2, 78, 78, 110, 110, 3, 2, 50, 59, 4, 2, 50, 59, 97,
	97, 5, 2, 50, 59, 67, 72, 99, 104, 10, 2, 36, 36, 41, 41, 94, 94, 100,
	100, 104, 104, 112, 112, 116, 116, 118, 118, 3, 2, 50, 53, 3, 2, 50, 57,
	6, 2, 38, 38, 67, 92, 97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2,
	55298, 56321, 3, 2, 56322, 57345, 2, 220, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	3, 47, 3, 2, 2, 2, 5, 49, 3, 2, 2, 2, 7, 51, 3, 2, 2, 2, 9, 53, 3, 2, 2,
	2, 11, 55, 3, 2, 2, 2, 13, 57, 3, 2, 2, 2, 15, 59, 3, 2, 2, 2, 17, 67,
	3, 2, 2, 2, 19, 74, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2, 23, 88, 3, 2, 2, 2,
	25, 93, 3, 2, 2, 2, 27, 101, 3, 2, 2, 2, 29, 107, 3, 2, 2, 2, 31, 121,
	3, 2, 2, 2, 33, 132, 3, 2, 2, 2, 35, 155, 3, 2, 2, 2, 37, 160, 3, 2, 2,
	2, 39, 170, 3, 2, 2, 2, 41, 193, 3, 2, 2, 2, 43, 197, 3, 2, 2, 2, 45, 203,
	3, 2, 2, 2, 47, 48, 7, 42, 2, 2, 48, 4, 3, 2, 2, 2, 49, 50, 7, 43, 2, 2,
	50, 6, 3, 2, 2, 2, 51, 52, 7, 125, 2, 2, 52, 8, 3, 2, 2, 2, 53, 54, 7,
	127, 2, 2, 54, 10, 3, 2, 2, 2, 55, 56, 7, 46, 2, 2, 56, 12, 3, 2, 2, 2,
	57, 58, 7, 63, 2, 2, 58, 14, 3, 2, 2, 2, 59, 60, 7, 114, 2, 2, 60, 61,
	7, 99, 2, 2, 61, 62, 7, 101, 2, 2, 62, 63, 7, 109, 2, 2, 63, 64, 7, 99,
	2, 2, 64, 65, 7, 105, 2, 2, 65, 66, 7, 103, 2, 2, 66, 16, 3, 2, 2, 2, 67,
	68, 7, 107, 2, 2, 68, 69, 7, 111, 2, 2, 69, 70, 7, 114, 2, 2, 70, 71, 7,
	113, 2, 2, 71, 72, 7, 116, 2, 2, 72, 73, 7, 118, 2, 2, 73, 18, 3, 2, 2,
	2, 74, 75, 7, 117, 2, 2, 75, 76, 7, 118, 2, 2, 76, 77, 7, 116, 2, 2, 77,
	78, 7, 119, 2, 2, 78, 79, 7, 101, 2, 2, 79, 80, 7, 118, 2, 2, 80, 20, 3,
	2, 2, 2, 81, 82, 7, 111, 2, 2, 82, 83, 7, 103, 2, 2, 83, 84, 7, 111, 2,
	2, 84, 85, 7, 100, 2, 2, 85, 86, 7, 103, 2, 2, 86, 87, 7, 116, 2, 2, 87,
	22, 3, 2, 2, 2, 88, 89, 7, 104, 2, 2, 89, 90, 7, 119, 2, 2, 90, 91, 7,
	112, 2, 2, 91, 92, 7, 101, 2, 2, 92, 24, 3, 2, 2, 2, 93, 97, 5, 45, 23,
	2, 94, 96, 5, 43, 22, 2, 95, 94, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95,
	3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 26, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2,
	100, 102, 9, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 106,
	8, 14, 2, 2, 106, 28, 3, 2, 2, 2, 107, 108, 7, 49, 2, 2, 108, 109, 7, 44,
	2, 2, 109, 113, 3, 2, 2, 2, 110, 112, 11, 2, 2, 2, 111, 110, 3, 2, 2, 2,
	112, 115, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114,
	116, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117, 7, 44, 2, 2, 117, 118,
	7, 49, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 8, 15, 2, 2, 120, 30, 3, 2,
	2, 2, 121, 122, 7, 49, 2, 2, 122, 123, 7, 49, 2, 2, 123, 127, 3, 2, 2,
	2, 124, 126, 10, 3, 2, 2, 125, 124, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127,
	125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 130, 3, 2, 2, 2, 129, 127,
	3, 2, 2, 2, 130, 131, 8, 16, 2, 2, 131, 32, 3, 2, 2, 2, 132, 137, 7, 36,
	2, 2, 133, 136, 10, 4, 2, 2, 134, 136, 5, 41, 21, 2, 135, 133, 3, 2, 2,
	2, 135, 134, 3, 2, 2, 2, 136, 139, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 140, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 141,
	7, 36, 2, 2, 141, 34, 3, 2, 2, 2, 142, 156, 7, 50, 2, 2, 143, 153, 9, 5,
	2, 2, 144, 146, 5, 37, 19, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2, 2,
	2, 146, 154, 3, 2, 2, 2, 147, 149, 7, 97, 2, 2, 148, 147, 3, 2, 2, 2, 149,
	150, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 152,
	3, 2, 2, 2, 152, 154, 5, 37, 19, 2, 153, 145, 3, 2, 2, 2, 153, 148, 3,
	2, 2, 2, 154, 156, 3, 2, 2, 2, 155, 142, 3, 2, 2, 2, 155, 143, 3, 2, 2,
	2, 156, 158, 3, 2, 2, 2, 157, 159, 9, 6, 2, 2, 158, 157, 3, 2, 2, 2, 158,
	159, 3, 2, 2, 2, 159, 36, 3, 2, 2, 2, 160, 168, 9, 7, 2, 2, 161, 163, 9,
	8, 2, 2, 162, 161, 3, 2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 162, 3, 2, 2,
	2, 164, 165, 3, 2, 2, 2, 165, 167, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167,
	169, 9, 7, 2, 2, 168, 164, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 38, 3,
	2, 2, 2, 170, 171, 9, 9, 2, 2, 171, 40, 3, 2, 2, 2, 172, 173, 7, 94, 2,
	2, 173, 194, 9, 10, 2, 2, 174, 179, 7, 94, 2, 2, 175, 177, 9, 11, 2, 2,
	176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178,
	180, 9, 12, 2, 2, 179, 176, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181,
	3, 2, 2, 2, 181, 194, 9, 12, 2, 2, 182, 184, 7, 94, 2, 2, 183, 185, 7,
	119, 2, 2, 184, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 184, 3, 2,
	2, 2, 186, 187, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 5, 39, 20,
	2, 189, 190, 5, 39, 20, 2, 190, 191, 5, 39, 20, 2, 191, 192, 5, 39, 20,
	2, 192, 194, 3, 2, 2, 2, 193, 172, 3, 2, 2, 2, 193, 174, 3, 2, 2, 2, 193,
	182, 3, 2, 2, 2, 194, 42, 3, 2, 2, 2, 195, 198, 5, 45, 23, 2, 196, 198,
	9, 7, 2, 2, 197, 195, 3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 44, 3, 2,
	2, 2, 199, 204, 9, 13, 2, 2, 200, 204, 10, 14, 2, 2, 201, 202, 9, 15, 2,
	2, 202, 204, 9, 16, 2, 2, 203, 199, 3, 2, 2, 2, 203, 200, 3, 2, 2, 2, 203,
	201, 3, 2, 2, 2, 204, 46, 3, 2, 2, 2, 22, 2, 97, 103, 113, 127, 135, 137,
	145, 150, 153, 155, 158, 164, 168, 176, 179, 186, 193, 197, 203, 3, 2,
	3, 2,
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
	"", "'('", "')'", "'{'", "'}'", "','", "'='", "'package'", "'import'",
	"'struct'", "'member'", "'func'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "PACKAGE", "IMPORT", "DATA_STRUCT", "MEMBER",
	"FUNCTION", "IDENTIFIER", "WS", "COMMENT", "LINE_COMMENT", "STRING_LITERAL",
	"DECIMAL_LITERAL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "PACKAGE", "IMPORT", "DATA_STRUCT",
	"MEMBER", "FUNCTION", "IDENTIFIER", "WS", "COMMENT", "LINE_COMMENT", "STRING_LITERAL",
	"DECIMAL_LITERAL", "Digits", "HexDigit", "EscapeSequence", "LetterOrDigit",
	"Letter",
}

type CodeLexer struct {
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

func NewCodeLexer(input antlr.CharStream) *CodeLexer {

	l := new(CodeLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Code.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CodeLexer tokens.
const (
	CodeLexerT__0            = 1
	CodeLexerT__1            = 2
	CodeLexerT__2            = 3
	CodeLexerT__3            = 4
	CodeLexerT__4            = 5
	CodeLexerT__5            = 6
	CodeLexerPACKAGE         = 7
	CodeLexerIMPORT          = 8
	CodeLexerDATA_STRUCT     = 9
	CodeLexerMEMBER          = 10
	CodeLexerFUNCTION        = 11
	CodeLexerIDENTIFIER      = 12
	CodeLexerWS              = 13
	CodeLexerCOMMENT         = 14
	CodeLexerLINE_COMMENT    = 15
	CodeLexerSTRING_LITERAL  = 16
	CodeLexerDECIMAL_LITERAL = 17
)
