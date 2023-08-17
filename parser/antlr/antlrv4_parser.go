// Code generated from ./parser/antlr/ANTLRv4Parser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlr // ANTLRv4Parser
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

type ANTLRv4Parser struct {
	*antlr.BaseParser
}

var ANTLRv4ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func antlrv4parserParserInit() {
	staticData := &ANTLRv4ParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'import'",
		"'fragment'", "'lexer'", "'parser'", "'grammar'", "'protected'", "'public'",
		"'private'", "'returns'", "'locals'", "'throws'", "'catch'", "'finally'",
		"'mode'",
	}
	staticData.SymbolicNames = []string{
		"", "TOKEN_REF", "RULE_REF", "LEXER_CHAR_SET", "DOC_COMMENT", "BLOCK_COMMENT",
		"LINE_COMMENT", "INT", "STRING_LITERAL", "UNTERMINATED_STRING_LITERAL",
		"BEGIN_ARGUMENT", "BEGIN_ACTION", "OPTIONS", "TOKENS", "CHANNELS", "IMPORT",
		"FRAGMENT", "LEXER", "PARSER", "GRAMMAR", "PROTECTED", "PUBLIC", "PRIVATE",
		"RETURNS", "LOCALS", "THROWS", "CATCH", "FINALLY", "MODE", "COLON",
		"COLONCOLON", "COMMA", "SEMI", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"RARROW", "LT", "GT", "ASSIGN", "QUESTION", "STAR", "PLUS_ASSIGN", "PLUS",
		"OR", "DOLLAR", "RANGE", "DOT", "AT", "POUND", "NOT", "ID", "WS", "ERRCHAR",
		"END_ARGUMENT", "UNTERMINATED_ARGUMENT", "ARGUMENT_CONTENT", "END_ACTION",
		"UNTERMINATED_ACTION", "ACTION_CONTENT", "UNTERMINATED_CHAR_SET",
	}
	staticData.RuleNames = []string{
		"grammarSpec", "grammarDecl", "grammarType", "prequelConstruct", "optionsSpec",
		"option", "optionValue", "delegateGrammars", "delegateGrammar", "tokensSpec",
		"channelsSpec", "idList", "action_", "actionScopeName", "actionBlock",
		"argActionBlock", "modeSpec", "rules", "ruleSpec", "parserRuleSpec",
		"exceptionGroup", "exceptionHandler", "finallyClause", "rulePrequel",
		"ruleReturns", "throwsSpec", "localsSpec", "ruleAction", "ruleModifiers",
		"ruleModifier", "ruleBlock", "ruleAltList", "labeledAlt", "lexerRuleSpec",
		"lexerRuleBlock", "lexerAltList", "lexerAlt", "lexerElements", "lexerElement",
		"lexerBlock", "lexerCommands", "lexerCommand", "lexerCommandName", "lexerCommandExpr",
		"altList", "alternative", "element", "labeledElement", "ebnf", "blockSuffix",
		"ebnfSuffix", "lexerAtom", "atom", "notSet", "blockSet", "setElement",
		"block", "ruleref", "characterRange", "terminalNode", "elementOptions",
		"elementOption", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 617, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 1,
		0, 1, 0, 5, 0, 129, 8, 0, 10, 0, 12, 0, 132, 9, 0, 1, 0, 1, 0, 5, 0, 136,
		8, 0, 10, 0, 12, 0, 139, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 152, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 159, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 165, 8, 4, 10, 4, 12, 4,
		168, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6,
		179, 8, 6, 10, 6, 12, 6, 182, 9, 6, 1, 6, 1, 6, 1, 6, 3, 6, 187, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 193, 8, 7, 10, 7, 12, 7, 196, 9, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 205, 8, 8, 1, 9, 1, 9, 3, 9,
		209, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 215, 8, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 5, 11, 222, 8, 11, 10, 11, 12, 11, 225, 9, 11, 1, 11,
		3, 11, 228, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 234, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 242, 8, 13, 1, 14, 1, 14, 5,
		14, 246, 8, 14, 10, 14, 12, 14, 249, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		5, 15, 255, 8, 15, 10, 15, 12, 15, 258, 9, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 266, 8, 16, 10, 16, 12, 16, 269, 9, 16, 1, 17,
		5, 17, 272, 8, 17, 10, 17, 12, 17, 275, 9, 17, 1, 18, 1, 18, 3, 18, 279,
		8, 18, 1, 19, 3, 19, 282, 8, 19, 1, 19, 1, 19, 3, 19, 286, 8, 19, 1, 19,
		3, 19, 289, 8, 19, 1, 19, 3, 19, 292, 8, 19, 1, 19, 3, 19, 295, 8, 19,
		1, 19, 5, 19, 298, 8, 19, 10, 19, 12, 19, 301, 9, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 20, 5, 20, 309, 8, 20, 10, 20, 12, 20, 312, 9, 20,
		1, 20, 3, 20, 315, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 3, 23, 326, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 25, 5, 25, 335, 8, 25, 10, 25, 12, 25, 338, 9, 25, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 4, 28, 348, 8, 28, 11, 28,
		12, 28, 349, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 5, 31, 359,
		8, 31, 10, 31, 12, 31, 362, 9, 31, 1, 32, 1, 32, 1, 32, 3, 32, 367, 8,
		32, 1, 33, 3, 33, 370, 8, 33, 1, 33, 1, 33, 3, 33, 374, 8, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 5, 35, 385, 8, 35,
		10, 35, 12, 35, 388, 9, 35, 1, 36, 1, 36, 3, 36, 392, 8, 36, 1, 36, 3,
		36, 395, 8, 36, 1, 37, 4, 37, 398, 8, 37, 11, 37, 12, 37, 399, 1, 37, 3,
		37, 403, 8, 37, 1, 38, 1, 38, 3, 38, 407, 8, 38, 1, 38, 1, 38, 3, 38, 411,
		8, 38, 1, 38, 1, 38, 3, 38, 415, 8, 38, 3, 38, 417, 8, 38, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40, 427, 8, 40, 10, 40, 12,
		40, 430, 9, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 438, 8,
		41, 1, 42, 1, 42, 3, 42, 442, 8, 42, 1, 43, 1, 43, 3, 43, 446, 8, 43, 1,
		44, 1, 44, 1, 44, 5, 44, 451, 8, 44, 10, 44, 12, 44, 454, 9, 44, 1, 45,
		3, 45, 457, 8, 45, 1, 45, 4, 45, 460, 8, 45, 11, 45, 12, 45, 461, 1, 45,
		3, 45, 465, 8, 45, 1, 46, 1, 46, 1, 46, 3, 46, 470, 8, 46, 1, 46, 1, 46,
		1, 46, 3, 46, 475, 8, 46, 1, 46, 1, 46, 1, 46, 3, 46, 480, 8, 46, 3, 46,
		482, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 488, 8, 47, 1, 48, 1, 48,
		3, 48, 492, 8, 48, 1, 49, 1, 49, 1, 50, 1, 50, 3, 50, 498, 8, 50, 1, 50,
		1, 50, 3, 50, 502, 8, 50, 1, 50, 1, 50, 3, 50, 506, 8, 50, 3, 50, 508,
		8, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 516, 8, 51, 3,
		51, 518, 8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 525, 8, 52, 3,
		52, 527, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53, 3, 53, 533, 8, 53, 1, 54, 1,
		54, 1, 54, 1, 54, 5, 54, 539, 8, 54, 10, 54, 12, 54, 542, 9, 54, 1, 54,
		1, 54, 1, 55, 1, 55, 3, 55, 548, 8, 55, 1, 55, 1, 55, 3, 55, 552, 8, 55,
		1, 55, 1, 55, 3, 55, 556, 8, 55, 1, 56, 1, 56, 3, 56, 560, 8, 56, 1, 56,
		5, 56, 563, 8, 56, 10, 56, 12, 56, 566, 9, 56, 1, 56, 3, 56, 569, 8, 56,
		1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 3, 57, 576, 8, 57, 1, 57, 3, 57, 579,
		8, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 3, 59, 587, 8, 59, 1,
		59, 1, 59, 3, 59, 591, 8, 59, 3, 59, 593, 8, 59, 1, 60, 1, 60, 1, 60, 1,
		60, 5, 60, 599, 8, 60, 10, 60, 12, 60, 602, 9, 60, 1, 60, 1, 60, 1, 61,
		1, 61, 1, 61, 1, 61, 1, 61, 3, 61, 611, 8, 61, 3, 61, 613, 8, 61, 1, 62,
		1, 62, 1, 62, 0, 0, 63, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
		98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 0,
		3, 2, 0, 16, 16, 20, 22, 2, 0, 40, 40, 43, 43, 1, 0, 1, 2, 653, 0, 126,
		1, 0, 0, 0, 2, 142, 1, 0, 0, 0, 4, 151, 1, 0, 0, 0, 6, 158, 1, 0, 0, 0,
		8, 160, 1, 0, 0, 0, 10, 171, 1, 0, 0, 0, 12, 186, 1, 0, 0, 0, 14, 188,
		1, 0, 0, 0, 16, 204, 1, 0, 0, 0, 18, 206, 1, 0, 0, 0, 20, 212, 1, 0, 0,
		0, 22, 218, 1, 0, 0, 0, 24, 229, 1, 0, 0, 0, 26, 241, 1, 0, 0, 0, 28, 243,
		1, 0, 0, 0, 30, 252, 1, 0, 0, 0, 32, 261, 1, 0, 0, 0, 34, 273, 1, 0, 0,
		0, 36, 278, 1, 0, 0, 0, 38, 281, 1, 0, 0, 0, 40, 310, 1, 0, 0, 0, 42, 316,
		1, 0, 0, 0, 44, 320, 1, 0, 0, 0, 46, 325, 1, 0, 0, 0, 48, 327, 1, 0, 0,
		0, 50, 330, 1, 0, 0, 0, 52, 339, 1, 0, 0, 0, 54, 342, 1, 0, 0, 0, 56, 347,
		1, 0, 0, 0, 58, 351, 1, 0, 0, 0, 60, 353, 1, 0, 0, 0, 62, 355, 1, 0, 0,
		0, 64, 363, 1, 0, 0, 0, 66, 369, 1, 0, 0, 0, 68, 379, 1, 0, 0, 0, 70, 381,
		1, 0, 0, 0, 72, 394, 1, 0, 0, 0, 74, 402, 1, 0, 0, 0, 76, 416, 1, 0, 0,
		0, 78, 418, 1, 0, 0, 0, 80, 422, 1, 0, 0, 0, 82, 437, 1, 0, 0, 0, 84, 441,
		1, 0, 0, 0, 86, 445, 1, 0, 0, 0, 88, 447, 1, 0, 0, 0, 90, 464, 1, 0, 0,
		0, 92, 481, 1, 0, 0, 0, 94, 483, 1, 0, 0, 0, 96, 489, 1, 0, 0, 0, 98, 493,
		1, 0, 0, 0, 100, 507, 1, 0, 0, 0, 102, 517, 1, 0, 0, 0, 104, 526, 1, 0,
		0, 0, 106, 532, 1, 0, 0, 0, 108, 534, 1, 0, 0, 0, 110, 555, 1, 0, 0, 0,
		112, 557, 1, 0, 0, 0, 114, 573, 1, 0, 0, 0, 116, 580, 1, 0, 0, 0, 118,
		592, 1, 0, 0, 0, 120, 594, 1, 0, 0, 0, 122, 612, 1, 0, 0, 0, 124, 614,
		1, 0, 0, 0, 126, 130, 3, 2, 1, 0, 127, 129, 3, 6, 3, 0, 128, 127, 1, 0,
		0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0,
		131, 133, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 137, 3, 34, 17, 0, 134,
		136, 3, 32, 16, 0, 135, 134, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139, 137, 1, 0,
		0, 0, 140, 141, 5, 0, 0, 1, 141, 1, 1, 0, 0, 0, 142, 143, 3, 4, 2, 0, 143,
		144, 3, 124, 62, 0, 144, 145, 5, 32, 0, 0, 145, 3, 1, 0, 0, 0, 146, 147,
		5, 17, 0, 0, 147, 152, 5, 19, 0, 0, 148, 149, 5, 18, 0, 0, 149, 152, 5,
		19, 0, 0, 150, 152, 5, 19, 0, 0, 151, 146, 1, 0, 0, 0, 151, 148, 1, 0,
		0, 0, 151, 150, 1, 0, 0, 0, 152, 5, 1, 0, 0, 0, 153, 159, 3, 8, 4, 0, 154,
		159, 3, 14, 7, 0, 155, 159, 3, 18, 9, 0, 156, 159, 3, 20, 10, 0, 157, 159,
		3, 24, 12, 0, 158, 153, 1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 158, 155, 1,
		0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 157, 1, 0, 0, 0, 159, 7, 1, 0, 0, 0,
		160, 166, 5, 12, 0, 0, 161, 162, 3, 10, 5, 0, 162, 163, 5, 32, 0, 0, 163,
		165, 1, 0, 0, 0, 164, 161, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164,
		1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168, 166, 1, 0,
		0, 0, 169, 170, 5, 36, 0, 0, 170, 9, 1, 0, 0, 0, 171, 172, 3, 124, 62,
		0, 172, 173, 5, 40, 0, 0, 173, 174, 3, 12, 6, 0, 174, 11, 1, 0, 0, 0, 175,
		180, 3, 124, 62, 0, 176, 177, 5, 48, 0, 0, 177, 179, 3, 124, 62, 0, 178,
		176, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181,
		1, 0, 0, 0, 181, 187, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 187, 5, 8,
		0, 0, 184, 187, 3, 28, 14, 0, 185, 187, 5, 7, 0, 0, 186, 175, 1, 0, 0,
		0, 186, 183, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187,
		13, 1, 0, 0, 0, 188, 189, 5, 15, 0, 0, 189, 194, 3, 16, 8, 0, 190, 191,
		5, 31, 0, 0, 191, 193, 3, 16, 8, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1,
		0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0,
		0, 196, 194, 1, 0, 0, 0, 197, 198, 5, 32, 0, 0, 198, 15, 1, 0, 0, 0, 199,
		200, 3, 124, 62, 0, 200, 201, 5, 40, 0, 0, 201, 202, 3, 124, 62, 0, 202,
		205, 1, 0, 0, 0, 203, 205, 3, 124, 62, 0, 204, 199, 1, 0, 0, 0, 204, 203,
		1, 0, 0, 0, 205, 17, 1, 0, 0, 0, 206, 208, 5, 13, 0, 0, 207, 209, 3, 22,
		11, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 211, 5, 36, 0, 0, 211, 19, 1, 0, 0, 0, 212, 214, 5, 14, 0, 0, 213,
		215, 3, 22, 11, 0, 214, 213, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216,
		1, 0, 0, 0, 216, 217, 5, 36, 0, 0, 217, 21, 1, 0, 0, 0, 218, 223, 3, 124,
		62, 0, 219, 220, 5, 31, 0, 0, 220, 222, 3, 124, 62, 0, 221, 219, 1, 0,
		0, 0, 222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0,
		224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 228, 5, 31, 0, 0, 227,
		226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 23, 1, 0, 0, 0, 229, 233, 5,
		49, 0, 0, 230, 231, 3, 26, 13, 0, 231, 232, 5, 30, 0, 0, 232, 234, 1, 0,
		0, 0, 233, 230, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0,
		235, 236, 3, 124, 62, 0, 236, 237, 3, 28, 14, 0, 237, 25, 1, 0, 0, 0, 238,
		242, 3, 124, 62, 0, 239, 242, 5, 17, 0, 0, 240, 242, 5, 18, 0, 0, 241,
		238, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 27, 1,
		0, 0, 0, 243, 247, 5, 11, 0, 0, 244, 246, 5, 60, 0, 0, 245, 244, 1, 0,
		0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0,
		248, 250, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 251, 5, 58, 0, 0, 251,
		29, 1, 0, 0, 0, 252, 256, 5, 10, 0, 0, 253, 255, 5, 57, 0, 0, 254, 253,
		1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256, 257, 1, 0,
		0, 0, 257, 259, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 260, 5, 55, 0, 0,
		260, 31, 1, 0, 0, 0, 261, 262, 5, 28, 0, 0, 262, 263, 3, 124, 62, 0, 263,
		267, 5, 32, 0, 0, 264, 266, 3, 66, 33, 0, 265, 264, 1, 0, 0, 0, 266, 269,
		1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 33, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 270, 272, 3, 36, 18, 0, 271, 270, 1, 0, 0,
		0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274,
		35, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 279, 3, 38, 19, 0, 277, 279,
		3, 66, 33, 0, 278, 276, 1, 0, 0, 0, 278, 277, 1, 0, 0, 0, 279, 37, 1, 0,
		0, 0, 280, 282, 3, 56, 28, 0, 281, 280, 1, 0, 0, 0, 281, 282, 1, 0, 0,
		0, 282, 283, 1, 0, 0, 0, 283, 285, 5, 2, 0, 0, 284, 286, 3, 30, 15, 0,
		285, 284, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 288, 1, 0, 0, 0, 287,
		289, 3, 48, 24, 0, 288, 287, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291,
		1, 0, 0, 0, 290, 292, 3, 50, 25, 0, 291, 290, 1, 0, 0, 0, 291, 292, 1,
		0, 0, 0, 292, 294, 1, 0, 0, 0, 293, 295, 3, 52, 26, 0, 294, 293, 1, 0,
		0, 0, 294, 295, 1, 0, 0, 0, 295, 299, 1, 0, 0, 0, 296, 298, 3, 46, 23,
		0, 297, 296, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299,
		300, 1, 0, 0, 0, 300, 302, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 303,
		5, 29, 0, 0, 303, 304, 3, 60, 30, 0, 304, 305, 5, 32, 0, 0, 305, 306, 3,
		40, 20, 0, 306, 39, 1, 0, 0, 0, 307, 309, 3, 42, 21, 0, 308, 307, 1, 0,
		0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0,
		311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 315, 3, 44, 22, 0, 314,
		313, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 41, 1, 0, 0, 0, 316, 317, 5,
		26, 0, 0, 317, 318, 3, 30, 15, 0, 318, 319, 3, 28, 14, 0, 319, 43, 1, 0,
		0, 0, 320, 321, 5, 27, 0, 0, 321, 322, 3, 28, 14, 0, 322, 45, 1, 0, 0,
		0, 323, 326, 3, 8, 4, 0, 324, 326, 3, 54, 27, 0, 325, 323, 1, 0, 0, 0,
		325, 324, 1, 0, 0, 0, 326, 47, 1, 0, 0, 0, 327, 328, 5, 23, 0, 0, 328,
		329, 3, 30, 15, 0, 329, 49, 1, 0, 0, 0, 330, 331, 5, 25, 0, 0, 331, 336,
		3, 124, 62, 0, 332, 333, 5, 31, 0, 0, 333, 335, 3, 124, 62, 0, 334, 332,
		1, 0, 0, 0, 335, 338, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0,
		0, 0, 337, 51, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339, 340, 5, 24, 0, 0,
		340, 341, 3, 30, 15, 0, 341, 53, 1, 0, 0, 0, 342, 343, 5, 49, 0, 0, 343,
		344, 3, 124, 62, 0, 344, 345, 3, 28, 14, 0, 345, 55, 1, 0, 0, 0, 346, 348,
		3, 58, 29, 0, 347, 346, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 347, 1,
		0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 57, 1, 0, 0, 0, 351, 352, 7, 0, 0,
		0, 352, 59, 1, 0, 0, 0, 353, 354, 3, 62, 31, 0, 354, 61, 1, 0, 0, 0, 355,
		360, 3, 64, 32, 0, 356, 357, 5, 45, 0, 0, 357, 359, 3, 64, 32, 0, 358,
		356, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 361,
		1, 0, 0, 0, 361, 63, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0, 363, 366, 3, 90,
		45, 0, 364, 365, 5, 50, 0, 0, 365, 367, 3, 124, 62, 0, 366, 364, 1, 0,
		0, 0, 366, 367, 1, 0, 0, 0, 367, 65, 1, 0, 0, 0, 368, 370, 5, 16, 0, 0,
		369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371,
		373, 5, 1, 0, 0, 372, 374, 3, 8, 4, 0, 373, 372, 1, 0, 0, 0, 373, 374,
		1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 5, 29, 0, 0, 376, 377, 3, 68,
		34, 0, 377, 378, 5, 32, 0, 0, 378, 67, 1, 0, 0, 0, 379, 380, 3, 70, 35,
		0, 380, 69, 1, 0, 0, 0, 381, 386, 3, 72, 36, 0, 382, 383, 5, 45, 0, 0,
		383, 385, 3, 72, 36, 0, 384, 382, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386,
		384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 71, 1, 0, 0, 0, 388, 386, 1,
		0, 0, 0, 389, 391, 3, 74, 37, 0, 390, 392, 3, 80, 40, 0, 391, 390, 1, 0,
		0, 0, 391, 392, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 395, 1, 0, 0, 0,
		394, 389, 1, 0, 0, 0, 394, 393, 1, 0, 0, 0, 395, 73, 1, 0, 0, 0, 396, 398,
		3, 76, 38, 0, 397, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 397, 1,
		0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 403, 1, 0, 0, 0, 401, 403, 1, 0, 0,
		0, 402, 397, 1, 0, 0, 0, 402, 401, 1, 0, 0, 0, 403, 75, 1, 0, 0, 0, 404,
		406, 3, 102, 51, 0, 405, 407, 3, 100, 50, 0, 406, 405, 1, 0, 0, 0, 406,
		407, 1, 0, 0, 0, 407, 417, 1, 0, 0, 0, 408, 410, 3, 78, 39, 0, 409, 411,
		3, 100, 50, 0, 410, 409, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 417, 1,
		0, 0, 0, 412, 414, 3, 28, 14, 0, 413, 415, 5, 41, 0, 0, 414, 413, 1, 0,
		0, 0, 414, 415, 1, 0, 0, 0, 415, 417, 1, 0, 0, 0, 416, 404, 1, 0, 0, 0,
		416, 408, 1, 0, 0, 0, 416, 412, 1, 0, 0, 0, 417, 77, 1, 0, 0, 0, 418, 419,
		5, 33, 0, 0, 419, 420, 3, 70, 35, 0, 420, 421, 5, 34, 0, 0, 421, 79, 1,
		0, 0, 0, 422, 423, 5, 37, 0, 0, 423, 428, 3, 82, 41, 0, 424, 425, 5, 31,
		0, 0, 425, 427, 3, 82, 41, 0, 426, 424, 1, 0, 0, 0, 427, 430, 1, 0, 0,
		0, 428, 426, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 81, 1, 0, 0, 0, 430,
		428, 1, 0, 0, 0, 431, 432, 3, 84, 42, 0, 432, 433, 5, 33, 0, 0, 433, 434,
		3, 86, 43, 0, 434, 435, 5, 34, 0, 0, 435, 438, 1, 0, 0, 0, 436, 438, 3,
		84, 42, 0, 437, 431, 1, 0, 0, 0, 437, 436, 1, 0, 0, 0, 438, 83, 1, 0, 0,
		0, 439, 442, 3, 124, 62, 0, 440, 442, 5, 28, 0, 0, 441, 439, 1, 0, 0, 0,
		441, 440, 1, 0, 0, 0, 442, 85, 1, 0, 0, 0, 443, 446, 3, 124, 62, 0, 444,
		446, 5, 7, 0, 0, 445, 443, 1, 0, 0, 0, 445, 444, 1, 0, 0, 0, 446, 87, 1,
		0, 0, 0, 447, 452, 3, 90, 45, 0, 448, 449, 5, 45, 0, 0, 449, 451, 3, 90,
		45, 0, 450, 448, 1, 0, 0, 0, 451, 454, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0,
		452, 453, 1, 0, 0, 0, 453, 89, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 455, 457,
		3, 120, 60, 0, 456, 455, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 459, 1,
		0, 0, 0, 458, 460, 3, 92, 46, 0, 459, 458, 1, 0, 0, 0, 460, 461, 1, 0,
		0, 0, 461, 459, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 465, 1, 0, 0, 0,
		463, 465, 1, 0, 0, 0, 464, 456, 1, 0, 0, 0, 464, 463, 1, 0, 0, 0, 465,
		91, 1, 0, 0, 0, 466, 469, 3, 94, 47, 0, 467, 470, 3, 100, 50, 0, 468, 470,
		1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469, 468, 1, 0, 0, 0, 470, 482, 1, 0,
		0, 0, 471, 474, 3, 104, 52, 0, 472, 475, 3, 100, 50, 0, 473, 475, 1, 0,
		0, 0, 474, 472, 1, 0, 0, 0, 474, 473, 1, 0, 0, 0, 475, 482, 1, 0, 0, 0,
		476, 482, 3, 96, 48, 0, 477, 479, 3, 28, 14, 0, 478, 480, 5, 41, 0, 0,
		479, 478, 1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 482, 1, 0, 0, 0, 481,
		466, 1, 0, 0, 0, 481, 471, 1, 0, 0, 0, 481, 476, 1, 0, 0, 0, 481, 477,
		1, 0, 0, 0, 482, 93, 1, 0, 0, 0, 483, 484, 3, 124, 62, 0, 484, 487, 7,
		1, 0, 0, 485, 488, 3, 104, 52, 0, 486, 488, 3, 112, 56, 0, 487, 485, 1,
		0, 0, 0, 487, 486, 1, 0, 0, 0, 488, 95, 1, 0, 0, 0, 489, 491, 3, 112, 56,
		0, 490, 492, 3, 98, 49, 0, 491, 490, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0,
		492, 97, 1, 0, 0, 0, 493, 494, 3, 100, 50, 0, 494, 99, 1, 0, 0, 0, 495,
		497, 5, 41, 0, 0, 496, 498, 5, 41, 0, 0, 497, 496, 1, 0, 0, 0, 497, 498,
		1, 0, 0, 0, 498, 508, 1, 0, 0, 0, 499, 501, 5, 42, 0, 0, 500, 502, 5, 41,
		0, 0, 501, 500, 1, 0, 0, 0, 501, 502, 1, 0, 0, 0, 502, 508, 1, 0, 0, 0,
		503, 505, 5, 44, 0, 0, 504, 506, 5, 41, 0, 0, 505, 504, 1, 0, 0, 0, 505,
		506, 1, 0, 0, 0, 506, 508, 1, 0, 0, 0, 507, 495, 1, 0, 0, 0, 507, 499,
		1, 0, 0, 0, 507, 503, 1, 0, 0, 0, 508, 101, 1, 0, 0, 0, 509, 518, 3, 116,
		58, 0, 510, 518, 3, 118, 59, 0, 511, 518, 3, 106, 53, 0, 512, 518, 5, 3,
		0, 0, 513, 515, 5, 48, 0, 0, 514, 516, 3, 120, 60, 0, 515, 514, 1, 0, 0,
		0, 515, 516, 1, 0, 0, 0, 516, 518, 1, 0, 0, 0, 517, 509, 1, 0, 0, 0, 517,
		510, 1, 0, 0, 0, 517, 511, 1, 0, 0, 0, 517, 512, 1, 0, 0, 0, 517, 513,
		1, 0, 0, 0, 518, 103, 1, 0, 0, 0, 519, 527, 3, 118, 59, 0, 520, 527, 3,
		114, 57, 0, 521, 527, 3, 106, 53, 0, 522, 524, 5, 48, 0, 0, 523, 525, 3,
		120, 60, 0, 524, 523, 1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 525, 527, 1, 0,
		0, 0, 526, 519, 1, 0, 0, 0, 526, 520, 1, 0, 0, 0, 526, 521, 1, 0, 0, 0,
		526, 522, 1, 0, 0, 0, 527, 105, 1, 0, 0, 0, 528, 529, 5, 51, 0, 0, 529,
		533, 3, 110, 55, 0, 530, 531, 5, 51, 0, 0, 531, 533, 3, 108, 54, 0, 532,
		528, 1, 0, 0, 0, 532, 530, 1, 0, 0, 0, 533, 107, 1, 0, 0, 0, 534, 535,
		5, 33, 0, 0, 535, 540, 3, 110, 55, 0, 536, 537, 5, 45, 0, 0, 537, 539,
		3, 110, 55, 0, 538, 536, 1, 0, 0, 0, 539, 542, 1, 0, 0, 0, 540, 538, 1,
		0, 0, 0, 540, 541, 1, 0, 0, 0, 541, 543, 1, 0, 0, 0, 542, 540, 1, 0, 0,
		0, 543, 544, 5, 34, 0, 0, 544, 109, 1, 0, 0, 0, 545, 547, 5, 1, 0, 0, 546,
		548, 3, 120, 60, 0, 547, 546, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 556,
		1, 0, 0, 0, 549, 551, 5, 8, 0, 0, 550, 552, 3, 120, 60, 0, 551, 550, 1,
		0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 556, 1, 0, 0, 0, 553, 556, 3, 116,
		58, 0, 554, 556, 5, 3, 0, 0, 555, 545, 1, 0, 0, 0, 555, 549, 1, 0, 0, 0,
		555, 553, 1, 0, 0, 0, 555, 554, 1, 0, 0, 0, 556, 111, 1, 0, 0, 0, 557,
		568, 5, 33, 0, 0, 558, 560, 3, 8, 4, 0, 559, 558, 1, 0, 0, 0, 559, 560,
		1, 0, 0, 0, 560, 564, 1, 0, 0, 0, 561, 563, 3, 54, 27, 0, 562, 561, 1,
		0, 0, 0, 563, 566, 1, 0, 0, 0, 564, 562, 1, 0, 0, 0, 564, 565, 1, 0, 0,
		0, 565, 567, 1, 0, 0, 0, 566, 564, 1, 0, 0, 0, 567, 569, 5, 29, 0, 0, 568,
		559, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570, 571,
		3, 88, 44, 0, 571, 572, 5, 34, 0, 0, 572, 113, 1, 0, 0, 0, 573, 575, 5,
		2, 0, 0, 574, 576, 3, 30, 15, 0, 575, 574, 1, 0, 0, 0, 575, 576, 1, 0,
		0, 0, 576, 578, 1, 0, 0, 0, 577, 579, 3, 120, 60, 0, 578, 577, 1, 0, 0,
		0, 578, 579, 1, 0, 0, 0, 579, 115, 1, 0, 0, 0, 580, 581, 5, 8, 0, 0, 581,
		582, 5, 47, 0, 0, 582, 583, 5, 8, 0, 0, 583, 117, 1, 0, 0, 0, 584, 586,
		5, 1, 0, 0, 585, 587, 3, 120, 60, 0, 586, 585, 1, 0, 0, 0, 586, 587, 1,
		0, 0, 0, 587, 593, 1, 0, 0, 0, 588, 590, 5, 8, 0, 0, 589, 591, 3, 120,
		60, 0, 590, 589, 1, 0, 0, 0, 590, 591, 1, 0, 0, 0, 591, 593, 1, 0, 0, 0,
		592, 584, 1, 0, 0, 0, 592, 588, 1, 0, 0, 0, 593, 119, 1, 0, 0, 0, 594,
		595, 5, 38, 0, 0, 595, 600, 3, 122, 61, 0, 596, 597, 5, 31, 0, 0, 597,
		599, 3, 122, 61, 0, 598, 596, 1, 0, 0, 0, 599, 602, 1, 0, 0, 0, 600, 598,
		1, 0, 0, 0, 600, 601, 1, 0, 0, 0, 601, 603, 1, 0, 0, 0, 602, 600, 1, 0,
		0, 0, 603, 604, 5, 39, 0, 0, 604, 121, 1, 0, 0, 0, 605, 613, 3, 124, 62,
		0, 606, 607, 3, 124, 62, 0, 607, 610, 5, 40, 0, 0, 608, 611, 3, 124, 62,
		0, 609, 611, 5, 8, 0, 0, 610, 608, 1, 0, 0, 0, 610, 609, 1, 0, 0, 0, 611,
		613, 1, 0, 0, 0, 612, 605, 1, 0, 0, 0, 612, 606, 1, 0, 0, 0, 613, 123,
		1, 0, 0, 0, 614, 615, 7, 2, 0, 0, 615, 125, 1, 0, 0, 0, 82, 130, 137, 151,
		158, 166, 180, 186, 194, 204, 208, 214, 223, 227, 233, 241, 247, 256, 267,
		273, 278, 281, 285, 288, 291, 294, 299, 310, 314, 325, 336, 349, 360, 366,
		369, 373, 386, 391, 394, 399, 402, 406, 410, 414, 416, 428, 437, 441, 445,
		452, 456, 461, 464, 469, 474, 479, 481, 487, 491, 497, 501, 505, 507, 515,
		517, 524, 526, 532, 540, 547, 551, 555, 559, 564, 568, 575, 578, 586, 590,
		592, 600, 610, 612,
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

// ANTLRv4ParserInit initializes any static state used to implement ANTLRv4Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewANTLRv4Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ANTLRv4ParserInit() {
	staticData := &ANTLRv4ParserParserStaticData
	staticData.once.Do(antlrv4parserParserInit)
}

// NewANTLRv4Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewANTLRv4Parser(input antlr.TokenStream) *ANTLRv4Parser {
	ANTLRv4ParserInit()
	this := new(ANTLRv4Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ANTLRv4ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ANTLRv4Parser.g4"

	return this
}

// ANTLRv4Parser tokens.
const (
	ANTLRv4ParserEOF                         = antlr.TokenEOF
	ANTLRv4ParserTOKEN_REF                   = 1
	ANTLRv4ParserRULE_REF                    = 2
	ANTLRv4ParserLEXER_CHAR_SET              = 3
	ANTLRv4ParserDOC_COMMENT                 = 4
	ANTLRv4ParserBLOCK_COMMENT               = 5
	ANTLRv4ParserLINE_COMMENT                = 6
	ANTLRv4ParserINT                         = 7
	ANTLRv4ParserSTRING_LITERAL              = 8
	ANTLRv4ParserUNTERMINATED_STRING_LITERAL = 9
	ANTLRv4ParserBEGIN_ARGUMENT              = 10
	ANTLRv4ParserBEGIN_ACTION                = 11
	ANTLRv4ParserOPTIONS                     = 12
	ANTLRv4ParserTOKENS                      = 13
	ANTLRv4ParserCHANNELS                    = 14
	ANTLRv4ParserIMPORT                      = 15
	ANTLRv4ParserFRAGMENT                    = 16
	ANTLRv4ParserLEXER                       = 17
	ANTLRv4ParserPARSER                      = 18
	ANTLRv4ParserGRAMMAR                     = 19
	ANTLRv4ParserPROTECTED                   = 20
	ANTLRv4ParserPUBLIC                      = 21
	ANTLRv4ParserPRIVATE                     = 22
	ANTLRv4ParserRETURNS                     = 23
	ANTLRv4ParserLOCALS                      = 24
	ANTLRv4ParserTHROWS                      = 25
	ANTLRv4ParserCATCH                       = 26
	ANTLRv4ParserFINALLY                     = 27
	ANTLRv4ParserMODE                        = 28
	ANTLRv4ParserCOLON                       = 29
	ANTLRv4ParserCOLONCOLON                  = 30
	ANTLRv4ParserCOMMA                       = 31
	ANTLRv4ParserSEMI                        = 32
	ANTLRv4ParserLPAREN                      = 33
	ANTLRv4ParserRPAREN                      = 34
	ANTLRv4ParserLBRACE                      = 35
	ANTLRv4ParserRBRACE                      = 36
	ANTLRv4ParserRARROW                      = 37
	ANTLRv4ParserLT                          = 38
	ANTLRv4ParserGT                          = 39
	ANTLRv4ParserASSIGN                      = 40
	ANTLRv4ParserQUESTION                    = 41
	ANTLRv4ParserSTAR                        = 42
	ANTLRv4ParserPLUS_ASSIGN                 = 43
	ANTLRv4ParserPLUS                        = 44
	ANTLRv4ParserOR                          = 45
	ANTLRv4ParserDOLLAR                      = 46
	ANTLRv4ParserRANGE                       = 47
	ANTLRv4ParserDOT                         = 48
	ANTLRv4ParserAT                          = 49
	ANTLRv4ParserPOUND                       = 50
	ANTLRv4ParserNOT                         = 51
	ANTLRv4ParserID                          = 52
	ANTLRv4ParserWS                          = 53
	ANTLRv4ParserERRCHAR                     = 54
	ANTLRv4ParserEND_ARGUMENT                = 55
	ANTLRv4ParserUNTERMINATED_ARGUMENT       = 56
	ANTLRv4ParserARGUMENT_CONTENT            = 57
	ANTLRv4ParserEND_ACTION                  = 58
	ANTLRv4ParserUNTERMINATED_ACTION         = 59
	ANTLRv4ParserACTION_CONTENT              = 60
	ANTLRv4ParserUNTERMINATED_CHAR_SET       = 61
)

// ANTLRv4Parser rules.
const (
	ANTLRv4ParserRULE_grammarSpec      = 0
	ANTLRv4ParserRULE_grammarDecl      = 1
	ANTLRv4ParserRULE_grammarType      = 2
	ANTLRv4ParserRULE_prequelConstruct = 3
	ANTLRv4ParserRULE_optionsSpec      = 4
	ANTLRv4ParserRULE_option           = 5
	ANTLRv4ParserRULE_optionValue      = 6
	ANTLRv4ParserRULE_delegateGrammars = 7
	ANTLRv4ParserRULE_delegateGrammar  = 8
	ANTLRv4ParserRULE_tokensSpec       = 9
	ANTLRv4ParserRULE_channelsSpec     = 10
	ANTLRv4ParserRULE_idList           = 11
	ANTLRv4ParserRULE_action_          = 12
	ANTLRv4ParserRULE_actionScopeName  = 13
	ANTLRv4ParserRULE_actionBlock      = 14
	ANTLRv4ParserRULE_argActionBlock   = 15
	ANTLRv4ParserRULE_modeSpec         = 16
	ANTLRv4ParserRULE_rules            = 17
	ANTLRv4ParserRULE_ruleSpec         = 18
	ANTLRv4ParserRULE_parserRuleSpec   = 19
	ANTLRv4ParserRULE_exceptionGroup   = 20
	ANTLRv4ParserRULE_exceptionHandler = 21
	ANTLRv4ParserRULE_finallyClause    = 22
	ANTLRv4ParserRULE_rulePrequel      = 23
	ANTLRv4ParserRULE_ruleReturns      = 24
	ANTLRv4ParserRULE_throwsSpec       = 25
	ANTLRv4ParserRULE_localsSpec       = 26
	ANTLRv4ParserRULE_ruleAction       = 27
	ANTLRv4ParserRULE_ruleModifiers    = 28
	ANTLRv4ParserRULE_ruleModifier     = 29
	ANTLRv4ParserRULE_ruleBlock        = 30
	ANTLRv4ParserRULE_ruleAltList      = 31
	ANTLRv4ParserRULE_labeledAlt       = 32
	ANTLRv4ParserRULE_lexerRuleSpec    = 33
	ANTLRv4ParserRULE_lexerRuleBlock   = 34
	ANTLRv4ParserRULE_lexerAltList     = 35
	ANTLRv4ParserRULE_lexerAlt         = 36
	ANTLRv4ParserRULE_lexerElements    = 37
	ANTLRv4ParserRULE_lexerElement     = 38
	ANTLRv4ParserRULE_lexerBlock       = 39
	ANTLRv4ParserRULE_lexerCommands    = 40
	ANTLRv4ParserRULE_lexerCommand     = 41
	ANTLRv4ParserRULE_lexerCommandName = 42
	ANTLRv4ParserRULE_lexerCommandExpr = 43
	ANTLRv4ParserRULE_altList          = 44
	ANTLRv4ParserRULE_alternative      = 45
	ANTLRv4ParserRULE_element          = 46
	ANTLRv4ParserRULE_labeledElement   = 47
	ANTLRv4ParserRULE_ebnf             = 48
	ANTLRv4ParserRULE_blockSuffix      = 49
	ANTLRv4ParserRULE_ebnfSuffix       = 50
	ANTLRv4ParserRULE_lexerAtom        = 51
	ANTLRv4ParserRULE_atom             = 52
	ANTLRv4ParserRULE_notSet           = 53
	ANTLRv4ParserRULE_blockSet         = 54
	ANTLRv4ParserRULE_setElement       = 55
	ANTLRv4ParserRULE_block            = 56
	ANTLRv4ParserRULE_ruleref          = 57
	ANTLRv4ParserRULE_characterRange   = 58
	ANTLRv4ParserRULE_terminalNode     = 59
	ANTLRv4ParserRULE_elementOptions   = 60
	ANTLRv4ParserRULE_elementOption    = 61
	ANTLRv4ParserRULE_identifier       = 62
)

// IGrammarSpecContext is an interface to support dynamic dispatch.
type IGrammarSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GrammarDecl() IGrammarDeclContext
	Rules() IRulesContext
	EOF() antlr.TerminalNode
	AllPrequelConstruct() []IPrequelConstructContext
	PrequelConstruct(i int) IPrequelConstructContext
	AllModeSpec() []IModeSpecContext
	ModeSpec(i int) IModeSpecContext

	// IsGrammarSpecContext differentiates from other interfaces.
	IsGrammarSpecContext()
}

type GrammarSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrammarSpecContext() *GrammarSpecContext {
	var p = new(GrammarSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_grammarSpec
	return p
}

func InitEmptyGrammarSpecContext(p *GrammarSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_grammarSpec
}

func (*GrammarSpecContext) IsGrammarSpecContext() {}

func NewGrammarSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrammarSpecContext {
	var p = new(GrammarSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_grammarSpec

	return p
}

func (s *GrammarSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *GrammarSpecContext) GrammarDecl() IGrammarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrammarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrammarDeclContext)
}

func (s *GrammarSpecContext) Rules() IRulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRulesContext)
}

func (s *GrammarSpecContext) EOF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserEOF, 0)
}

func (s *GrammarSpecContext) AllPrequelConstruct() []IPrequelConstructContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrequelConstructContext); ok {
			len++
		}
	}

	tst := make([]IPrequelConstructContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrequelConstructContext); ok {
			tst[i] = t.(IPrequelConstructContext)
			i++
		}
	}

	return tst
}

func (s *GrammarSpecContext) PrequelConstruct(i int) IPrequelConstructContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrequelConstructContext); ok {
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

	return t.(IPrequelConstructContext)
}

func (s *GrammarSpecContext) AllModeSpec() []IModeSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModeSpecContext); ok {
			len++
		}
	}

	tst := make([]IModeSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModeSpecContext); ok {
			tst[i] = t.(IModeSpecContext)
			i++
		}
	}

	return tst
}

func (s *GrammarSpecContext) ModeSpec(i int) IModeSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModeSpecContext); ok {
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

	return t.(IModeSpecContext)
}

func (s *GrammarSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrammarSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrammarSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterGrammarSpec(s)
	}
}

func (s *GrammarSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitGrammarSpec(s)
	}
}

func (s *GrammarSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitGrammarSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) GrammarSpec() (localctx IGrammarSpecContext) {
	localctx = NewGrammarSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ANTLRv4ParserRULE_grammarSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.GrammarDecl()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562949953482752) != 0 {
		{
			p.SetState(127)
			p.PrequelConstruct()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Rules()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserMODE {
		{
			p.SetState(134)
			p.ModeSpec()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(ANTLRv4ParserEOF)
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

// IGrammarDeclContext is an interface to support dynamic dispatch.
type IGrammarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GrammarType() IGrammarTypeContext
	Identifier() IIdentifierContext
	SEMI() antlr.TerminalNode

	// IsGrammarDeclContext differentiates from other interfaces.
	IsGrammarDeclContext()
}

type GrammarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrammarDeclContext() *GrammarDeclContext {
	var p = new(GrammarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_grammarDecl
	return p
}

func InitEmptyGrammarDeclContext(p *GrammarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_grammarDecl
}

func (*GrammarDeclContext) IsGrammarDeclContext() {}

func NewGrammarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrammarDeclContext {
	var p = new(GrammarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_grammarDecl

	return p
}

func (s *GrammarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *GrammarDeclContext) GrammarType() IGrammarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrammarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrammarTypeContext)
}

func (s *GrammarDeclContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *GrammarDeclContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSEMI, 0)
}

func (s *GrammarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrammarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrammarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterGrammarDecl(s)
	}
}

func (s *GrammarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitGrammarDecl(s)
	}
}

func (s *GrammarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitGrammarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) GrammarDecl() (localctx IGrammarDeclContext) {
	localctx = NewGrammarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ANTLRv4ParserRULE_grammarDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.GrammarType()
	}
	{
		p.SetState(143)
		p.Identifier()
	}
	{
		p.SetState(144)
		p.Match(ANTLRv4ParserSEMI)
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

// IGrammarTypeContext is an interface to support dynamic dispatch.
type IGrammarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEXER() antlr.TerminalNode
	GRAMMAR() antlr.TerminalNode
	PARSER() antlr.TerminalNode

	// IsGrammarTypeContext differentiates from other interfaces.
	IsGrammarTypeContext()
}

type GrammarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrammarTypeContext() *GrammarTypeContext {
	var p = new(GrammarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_grammarType
	return p
}

func InitEmptyGrammarTypeContext(p *GrammarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_grammarType
}

func (*GrammarTypeContext) IsGrammarTypeContext() {}

func NewGrammarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrammarTypeContext {
	var p = new(GrammarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_grammarType

	return p
}

func (s *GrammarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *GrammarTypeContext) LEXER() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLEXER, 0)
}

func (s *GrammarTypeContext) GRAMMAR() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserGRAMMAR, 0)
}

func (s *GrammarTypeContext) PARSER() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPARSER, 0)
}

func (s *GrammarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrammarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrammarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterGrammarType(s)
	}
}

func (s *GrammarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitGrammarType(s)
	}
}

func (s *GrammarTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitGrammarType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) GrammarType() (localctx IGrammarTypeContext) {
	localctx = NewGrammarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ANTLRv4ParserRULE_grammarType)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserLEXER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Match(ANTLRv4ParserLEXER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(ANTLRv4ParserGRAMMAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ANTLRv4ParserPARSER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(ANTLRv4ParserPARSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(ANTLRv4ParserGRAMMAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ANTLRv4ParserGRAMMAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Match(ANTLRv4ParserGRAMMAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IPrequelConstructContext is an interface to support dynamic dispatch.
type IPrequelConstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionsSpec() IOptionsSpecContext
	DelegateGrammars() IDelegateGrammarsContext
	TokensSpec() ITokensSpecContext
	ChannelsSpec() IChannelsSpecContext
	Action_() IAction_Context

	// IsPrequelConstructContext differentiates from other interfaces.
	IsPrequelConstructContext()
}

type PrequelConstructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrequelConstructContext() *PrequelConstructContext {
	var p = new(PrequelConstructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_prequelConstruct
	return p
}

func InitEmptyPrequelConstructContext(p *PrequelConstructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_prequelConstruct
}

func (*PrequelConstructContext) IsPrequelConstructContext() {}

func NewPrequelConstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrequelConstructContext {
	var p = new(PrequelConstructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_prequelConstruct

	return p
}

func (s *PrequelConstructContext) GetParser() antlr.Parser { return s.parser }

func (s *PrequelConstructContext) OptionsSpec() IOptionsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionsSpecContext)
}

func (s *PrequelConstructContext) DelegateGrammars() IDelegateGrammarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelegateGrammarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelegateGrammarsContext)
}

func (s *PrequelConstructContext) TokensSpec() ITokensSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITokensSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITokensSpecContext)
}

func (s *PrequelConstructContext) ChannelsSpec() IChannelsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChannelsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChannelsSpecContext)
}

func (s *PrequelConstructContext) Action_() IAction_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_Context)
}

func (s *PrequelConstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrequelConstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrequelConstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterPrequelConstruct(s)
	}
}

func (s *PrequelConstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitPrequelConstruct(s)
	}
}

func (s *PrequelConstructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitPrequelConstruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) PrequelConstruct() (localctx IPrequelConstructContext) {
	localctx = NewPrequelConstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ANTLRv4ParserRULE_prequelConstruct)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserOPTIONS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.OptionsSpec()
		}

	case ANTLRv4ParserIMPORT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.DelegateGrammars()
		}

	case ANTLRv4ParserTOKENS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.TokensSpec()
		}

	case ANTLRv4ParserCHANNELS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.ChannelsSpec()
		}

	case ANTLRv4ParserAT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Action_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IOptionsSpecContext is an interface to support dynamic dispatch.
type IOptionsSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTIONS() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllOption() []IOptionContext
	Option(i int) IOptionContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsOptionsSpecContext differentiates from other interfaces.
	IsOptionsSpecContext()
}

type OptionsSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionsSpecContext() *OptionsSpecContext {
	var p = new(OptionsSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_optionsSpec
	return p
}

func InitEmptyOptionsSpecContext(p *OptionsSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_optionsSpec
}

func (*OptionsSpecContext) IsOptionsSpecContext() {}

func NewOptionsSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionsSpecContext {
	var p = new(OptionsSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_optionsSpec

	return p
}

func (s *OptionsSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionsSpecContext) OPTIONS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserOPTIONS, 0)
}

func (s *OptionsSpecContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRBRACE, 0)
}

func (s *OptionsSpecContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *OptionsSpecContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
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

	return t.(IOptionContext)
}

func (s *OptionsSpecContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserSEMI)
}

func (s *OptionsSpecContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSEMI, i)
}

func (s *OptionsSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionsSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionsSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterOptionsSpec(s)
	}
}

func (s *OptionsSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitOptionsSpec(s)
	}
}

func (s *OptionsSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitOptionsSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) OptionsSpec() (localctx IOptionsSpecContext) {
	localctx = NewOptionsSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ANTLRv4ParserRULE_optionsSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(ANTLRv4ParserOPTIONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserTOKEN_REF || _la == ANTLRv4ParserRULE_REF {
		{
			p.SetState(161)
			p.Option()
		}
		{
			p.SetState(162)
			p.Match(ANTLRv4ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
		p.Match(ANTLRv4ParserRBRACE)
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

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	ASSIGN() antlr.TerminalNode
	OptionValue() IOptionValueContext

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *OptionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserASSIGN, 0)
}

func (s *OptionContext) OptionValue() IOptionValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionValueContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitOption(s)
	}
}

func (s *OptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ANTLRv4ParserRULE_option)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Identifier()
	}
	{
		p.SetState(172)
		p.Match(ANTLRv4ParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.OptionValue()
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

// IOptionValueContext is an interface to support dynamic dispatch.
type IOptionValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	ActionBlock() IActionBlockContext
	INT() antlr.TerminalNode

	// IsOptionValueContext differentiates from other interfaces.
	IsOptionValueContext()
}

type OptionValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueContext() *OptionValueContext {
	var p = new(OptionValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_optionValue
	return p
}

func InitEmptyOptionValueContext(p *OptionValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_optionValue
}

func (*OptionValueContext) IsOptionValueContext() {}

func NewOptionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueContext {
	var p = new(OptionValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_optionValue

	return p
}

func (s *OptionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *OptionValueContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *OptionValueContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserDOT)
}

func (s *OptionValueContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserDOT, i)
}

func (s *OptionValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSTRING_LITERAL, 0)
}

func (s *OptionValueContext) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *OptionValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserINT, 0)
}

func (s *OptionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterOptionValue(s)
	}
}

func (s *OptionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitOptionValue(s)
	}
}

func (s *OptionValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitOptionValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) OptionValue() (localctx IOptionValueContext) {
	localctx = NewOptionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ANTLRv4ParserRULE_optionValue)
	var _la int

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Identifier()
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ANTLRv4ParserDOT {
			{
				p.SetState(176)
				p.Match(ANTLRv4ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(177)
				p.Identifier()
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ANTLRv4ParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(ANTLRv4ParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ANTLRv4ParserBEGIN_ACTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.ActionBlock()
		}

	case ANTLRv4ParserINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(185)
			p.Match(ANTLRv4ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IDelegateGrammarsContext is an interface to support dynamic dispatch.
type IDelegateGrammarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	AllDelegateGrammar() []IDelegateGrammarContext
	DelegateGrammar(i int) IDelegateGrammarContext
	SEMI() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDelegateGrammarsContext differentiates from other interfaces.
	IsDelegateGrammarsContext()
}

type DelegateGrammarsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelegateGrammarsContext() *DelegateGrammarsContext {
	var p = new(DelegateGrammarsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_delegateGrammars
	return p
}

func InitEmptyDelegateGrammarsContext(p *DelegateGrammarsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_delegateGrammars
}

func (*DelegateGrammarsContext) IsDelegateGrammarsContext() {}

func NewDelegateGrammarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelegateGrammarsContext {
	var p = new(DelegateGrammarsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_delegateGrammars

	return p
}

func (s *DelegateGrammarsContext) GetParser() antlr.Parser { return s.parser }

func (s *DelegateGrammarsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserIMPORT, 0)
}

func (s *DelegateGrammarsContext) AllDelegateGrammar() []IDelegateGrammarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDelegateGrammarContext); ok {
			len++
		}
	}

	tst := make([]IDelegateGrammarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDelegateGrammarContext); ok {
			tst[i] = t.(IDelegateGrammarContext)
			i++
		}
	}

	return tst
}

func (s *DelegateGrammarsContext) DelegateGrammar(i int) IDelegateGrammarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelegateGrammarContext); ok {
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

	return t.(IDelegateGrammarContext)
}

func (s *DelegateGrammarsContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSEMI, 0)
}

func (s *DelegateGrammarsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserCOMMA)
}

func (s *DelegateGrammarsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOMMA, i)
}

func (s *DelegateGrammarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelegateGrammarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelegateGrammarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterDelegateGrammars(s)
	}
}

func (s *DelegateGrammarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitDelegateGrammars(s)
	}
}

func (s *DelegateGrammarsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitDelegateGrammars(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) DelegateGrammars() (localctx IDelegateGrammarsContext) {
	localctx = NewDelegateGrammarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ANTLRv4ParserRULE_delegateGrammars)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(ANTLRv4ParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.DelegateGrammar()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserCOMMA {
		{
			p.SetState(190)
			p.Match(ANTLRv4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.DelegateGrammar()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(197)
		p.Match(ANTLRv4ParserSEMI)
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

// IDelegateGrammarContext is an interface to support dynamic dispatch.
type IDelegateGrammarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	ASSIGN() antlr.TerminalNode

	// IsDelegateGrammarContext differentiates from other interfaces.
	IsDelegateGrammarContext()
}

type DelegateGrammarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelegateGrammarContext() *DelegateGrammarContext {
	var p = new(DelegateGrammarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_delegateGrammar
	return p
}

func InitEmptyDelegateGrammarContext(p *DelegateGrammarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_delegateGrammar
}

func (*DelegateGrammarContext) IsDelegateGrammarContext() {}

func NewDelegateGrammarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelegateGrammarContext {
	var p = new(DelegateGrammarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_delegateGrammar

	return p
}

func (s *DelegateGrammarContext) GetParser() antlr.Parser { return s.parser }

func (s *DelegateGrammarContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *DelegateGrammarContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *DelegateGrammarContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserASSIGN, 0)
}

func (s *DelegateGrammarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelegateGrammarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelegateGrammarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterDelegateGrammar(s)
	}
}

func (s *DelegateGrammarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitDelegateGrammar(s)
	}
}

func (s *DelegateGrammarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitDelegateGrammar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) DelegateGrammar() (localctx IDelegateGrammarContext) {
	localctx = NewDelegateGrammarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ANTLRv4ParserRULE_delegateGrammar)
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Identifier()
		}
		{
			p.SetState(200)
			p.Match(ANTLRv4ParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Identifier()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// ITokensSpecContext is an interface to support dynamic dispatch.
type ITokensSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TOKENS() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IdList() IIdListContext

	// IsTokensSpecContext differentiates from other interfaces.
	IsTokensSpecContext()
}

type TokensSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokensSpecContext() *TokensSpecContext {
	var p = new(TokensSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_tokensSpec
	return p
}

func InitEmptyTokensSpecContext(p *TokensSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_tokensSpec
}

func (*TokensSpecContext) IsTokensSpecContext() {}

func NewTokensSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokensSpecContext {
	var p = new(TokensSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_tokensSpec

	return p
}

func (s *TokensSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TokensSpecContext) TOKENS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserTOKENS, 0)
}

func (s *TokensSpecContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRBRACE, 0)
}

func (s *TokensSpecContext) IdList() IIdListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *TokensSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokensSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokensSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterTokensSpec(s)
	}
}

func (s *TokensSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitTokensSpec(s)
	}
}

func (s *TokensSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitTokensSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) TokensSpec() (localctx ITokensSpecContext) {
	localctx = NewTokensSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ANTLRv4ParserRULE_tokensSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(ANTLRv4ParserTOKENS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserTOKEN_REF || _la == ANTLRv4ParserRULE_REF {
		{
			p.SetState(207)
			p.IdList()
		}

	}
	{
		p.SetState(210)
		p.Match(ANTLRv4ParserRBRACE)
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

// IChannelsSpecContext is an interface to support dynamic dispatch.
type IChannelsSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHANNELS() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IdList() IIdListContext

	// IsChannelsSpecContext differentiates from other interfaces.
	IsChannelsSpecContext()
}

type ChannelsSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChannelsSpecContext() *ChannelsSpecContext {
	var p = new(ChannelsSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_channelsSpec
	return p
}

func InitEmptyChannelsSpecContext(p *ChannelsSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_channelsSpec
}

func (*ChannelsSpecContext) IsChannelsSpecContext() {}

func NewChannelsSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChannelsSpecContext {
	var p = new(ChannelsSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_channelsSpec

	return p
}

func (s *ChannelsSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ChannelsSpecContext) CHANNELS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCHANNELS, 0)
}

func (s *ChannelsSpecContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRBRACE, 0)
}

func (s *ChannelsSpecContext) IdList() IIdListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *ChannelsSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChannelsSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChannelsSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterChannelsSpec(s)
	}
}

func (s *ChannelsSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitChannelsSpec(s)
	}
}

func (s *ChannelsSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitChannelsSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ChannelsSpec() (localctx IChannelsSpecContext) {
	localctx = NewChannelsSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ANTLRv4ParserRULE_channelsSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(ANTLRv4ParserCHANNELS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserTOKEN_REF || _la == ANTLRv4ParserRULE_REF {
		{
			p.SetState(213)
			p.IdList()
		}

	}
	{
		p.SetState(216)
		p.Match(ANTLRv4ParserRBRACE)
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

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_idList
	return p
}

func InitEmptyIdListContext(p *IdListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_idList
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *IdListContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *IdListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserCOMMA)
}

func (s *IdListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOMMA, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterIdList(s)
	}
}

func (s *IdListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitIdList(s)
	}
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) IdList() (localctx IIdListContext) {
	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ANTLRv4ParserRULE_idList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Identifier()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(219)
				p.Match(ANTLRv4ParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(220)
				p.Identifier()
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserCOMMA {
		{
			p.SetState(226)
			p.Match(ANTLRv4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IAction_Context is an interface to support dynamic dispatch.
type IAction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	Identifier() IIdentifierContext
	ActionBlock() IActionBlockContext
	ActionScopeName() IActionScopeNameContext
	COLONCOLON() antlr.TerminalNode

	// IsAction_Context differentiates from other interfaces.
	IsAction_Context()
}

type Action_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_Context() *Action_Context {
	var p = new(Action_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_action_
	return p
}

func InitEmptyAction_Context(p *Action_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_action_
}

func (*Action_Context) IsAction_Context() {}

func NewAction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_Context {
	var p = new(Action_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_action_

	return p
}

func (s *Action_Context) GetParser() antlr.Parser { return s.parser }

func (s *Action_Context) AT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserAT, 0)
}

func (s *Action_Context) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Action_Context) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *Action_Context) ActionScopeName() IActionScopeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionScopeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionScopeNameContext)
}

func (s *Action_Context) COLONCOLON() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOLONCOLON, 0)
}

func (s *Action_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterAction_(s)
	}
}

func (s *Action_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitAction_(s)
	}
}

func (s *Action_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitAction_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Action_() (localctx IAction_Context) {
	localctx = NewAction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ANTLRv4ParserRULE_action_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(ANTLRv4ParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(230)
			p.ActionScopeName()
		}
		{
			p.SetState(231)
			p.Match(ANTLRv4ParserCOLONCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(235)
		p.Identifier()
	}
	{
		p.SetState(236)
		p.ActionBlock()
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

// IActionScopeNameContext is an interface to support dynamic dispatch.
type IActionScopeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	LEXER() antlr.TerminalNode
	PARSER() antlr.TerminalNode

	// IsActionScopeNameContext differentiates from other interfaces.
	IsActionScopeNameContext()
}

type ActionScopeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionScopeNameContext() *ActionScopeNameContext {
	var p = new(ActionScopeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_actionScopeName
	return p
}

func InitEmptyActionScopeNameContext(p *ActionScopeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_actionScopeName
}

func (*ActionScopeNameContext) IsActionScopeNameContext() {}

func NewActionScopeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionScopeNameContext {
	var p = new(ActionScopeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_actionScopeName

	return p
}

func (s *ActionScopeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionScopeNameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ActionScopeNameContext) LEXER() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLEXER, 0)
}

func (s *ActionScopeNameContext) PARSER() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPARSER, 0)
}

func (s *ActionScopeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionScopeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionScopeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterActionScopeName(s)
	}
}

func (s *ActionScopeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitActionScopeName(s)
	}
}

func (s *ActionScopeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitActionScopeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ActionScopeName() (localctx IActionScopeNameContext) {
	localctx = NewActionScopeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ANTLRv4ParserRULE_actionScopeName)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Identifier()
		}

	case ANTLRv4ParserLEXER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.Match(ANTLRv4ParserLEXER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ANTLRv4ParserPARSER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.Match(ANTLRv4ParserPARSER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IActionBlockContext is an interface to support dynamic dispatch.
type IActionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGIN_ACTION() antlr.TerminalNode
	END_ACTION() antlr.TerminalNode
	AllACTION_CONTENT() []antlr.TerminalNode
	ACTION_CONTENT(i int) antlr.TerminalNode

	// IsActionBlockContext differentiates from other interfaces.
	IsActionBlockContext()
}

type ActionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionBlockContext() *ActionBlockContext {
	var p = new(ActionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_actionBlock
	return p
}

func InitEmptyActionBlockContext(p *ActionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_actionBlock
}

func (*ActionBlockContext) IsActionBlockContext() {}

func NewActionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionBlockContext {
	var p = new(ActionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_actionBlock

	return p
}

func (s *ActionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionBlockContext) BEGIN_ACTION() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserBEGIN_ACTION, 0)
}

func (s *ActionBlockContext) END_ACTION() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserEND_ACTION, 0)
}

func (s *ActionBlockContext) AllACTION_CONTENT() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserACTION_CONTENT)
}

func (s *ActionBlockContext) ACTION_CONTENT(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserACTION_CONTENT, i)
}

func (s *ActionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterActionBlock(s)
	}
}

func (s *ActionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitActionBlock(s)
	}
}

func (s *ActionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitActionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ActionBlock() (localctx IActionBlockContext) {
	localctx = NewActionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ANTLRv4ParserRULE_actionBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(ANTLRv4ParserBEGIN_ACTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserACTION_CONTENT {
		{
			p.SetState(244)
			p.Match(ANTLRv4ParserACTION_CONTENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(250)
		p.Match(ANTLRv4ParserEND_ACTION)
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

// IArgActionBlockContext is an interface to support dynamic dispatch.
type IArgActionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BEGIN_ARGUMENT() antlr.TerminalNode
	END_ARGUMENT() antlr.TerminalNode
	AllARGUMENT_CONTENT() []antlr.TerminalNode
	ARGUMENT_CONTENT(i int) antlr.TerminalNode

	// IsArgActionBlockContext differentiates from other interfaces.
	IsArgActionBlockContext()
}

type ArgActionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgActionBlockContext() *ArgActionBlockContext {
	var p = new(ArgActionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_argActionBlock
	return p
}

func InitEmptyArgActionBlockContext(p *ArgActionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_argActionBlock
}

func (*ArgActionBlockContext) IsArgActionBlockContext() {}

func NewArgActionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgActionBlockContext {
	var p = new(ArgActionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_argActionBlock

	return p
}

func (s *ArgActionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgActionBlockContext) BEGIN_ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserBEGIN_ARGUMENT, 0)
}

func (s *ArgActionBlockContext) END_ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserEND_ARGUMENT, 0)
}

func (s *ArgActionBlockContext) AllARGUMENT_CONTENT() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserARGUMENT_CONTENT)
}

func (s *ArgActionBlockContext) ARGUMENT_CONTENT(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserARGUMENT_CONTENT, i)
}

func (s *ArgActionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgActionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgActionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterArgActionBlock(s)
	}
}

func (s *ArgActionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitArgActionBlock(s)
	}
}

func (s *ArgActionBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitArgActionBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ArgActionBlock() (localctx IArgActionBlockContext) {
	localctx = NewArgActionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ANTLRv4ParserRULE_argActionBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(ANTLRv4ParserBEGIN_ARGUMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserARGUMENT_CONTENT {
		{
			p.SetState(253)
			p.Match(ANTLRv4ParserARGUMENT_CONTENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(259)
		p.Match(ANTLRv4ParserEND_ARGUMENT)
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

// IModeSpecContext is an interface to support dynamic dispatch.
type IModeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODE() antlr.TerminalNode
	Identifier() IIdentifierContext
	SEMI() antlr.TerminalNode
	AllLexerRuleSpec() []ILexerRuleSpecContext
	LexerRuleSpec(i int) ILexerRuleSpecContext

	// IsModeSpecContext differentiates from other interfaces.
	IsModeSpecContext()
}

type ModeSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModeSpecContext() *ModeSpecContext {
	var p = new(ModeSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_modeSpec
	return p
}

func InitEmptyModeSpecContext(p *ModeSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_modeSpec
}

func (*ModeSpecContext) IsModeSpecContext() {}

func NewModeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModeSpecContext {
	var p = new(ModeSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_modeSpec

	return p
}

func (s *ModeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ModeSpecContext) MODE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserMODE, 0)
}

func (s *ModeSpecContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ModeSpecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSEMI, 0)
}

func (s *ModeSpecContext) AllLexerRuleSpec() []ILexerRuleSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILexerRuleSpecContext); ok {
			len++
		}
	}

	tst := make([]ILexerRuleSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILexerRuleSpecContext); ok {
			tst[i] = t.(ILexerRuleSpecContext)
			i++
		}
	}

	return tst
}

func (s *ModeSpecContext) LexerRuleSpec(i int) ILexerRuleSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerRuleSpecContext); ok {
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

	return t.(ILexerRuleSpecContext)
}

func (s *ModeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterModeSpec(s)
	}
}

func (s *ModeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitModeSpec(s)
	}
}

func (s *ModeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitModeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ModeSpec() (localctx IModeSpecContext) {
	localctx = NewModeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ANTLRv4ParserRULE_modeSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(ANTLRv4ParserMODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Identifier()
	}
	{
		p.SetState(263)
		p.Match(ANTLRv4ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserTOKEN_REF || _la == ANTLRv4ParserFRAGMENT {
		{
			p.SetState(264)
			p.LexerRuleSpec()
		}

		p.SetState(269)
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

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRuleSpec() []IRuleSpecContext
	RuleSpec(i int) IRuleSpecContext

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_rules
	return p
}

func InitEmptyRulesContext(p *RulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_rules
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) AllRuleSpec() []IRuleSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleSpecContext); ok {
			len++
		}
	}

	tst := make([]IRuleSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleSpecContext); ok {
			tst[i] = t.(IRuleSpecContext)
			i++
		}
	}

	return tst
}

func (s *RulesContext) RuleSpec(i int) IRuleSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleSpecContext); ok {
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

	return t.(IRuleSpecContext)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRules(s)
	}
}

func (s *RulesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRules(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ANTLRv4ParserRULE_rules)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7405574) != 0 {
		{
			p.SetState(270)
			p.RuleSpec()
		}

		p.SetState(275)
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

// IRuleSpecContext is an interface to support dynamic dispatch.
type IRuleSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParserRuleSpec() IParserRuleSpecContext
	LexerRuleSpec() ILexerRuleSpecContext

	// IsRuleSpecContext differentiates from other interfaces.
	IsRuleSpecContext()
}

type RuleSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleSpecContext() *RuleSpecContext {
	var p = new(RuleSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleSpec
	return p
}

func InitEmptyRuleSpecContext(p *RuleSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleSpec
}

func (*RuleSpecContext) IsRuleSpecContext() {}

func NewRuleSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleSpecContext {
	var p = new(RuleSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleSpec

	return p
}

func (s *RuleSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleSpecContext) ParserRuleSpec() IParserRuleSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParserRuleSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParserRuleSpecContext)
}

func (s *RuleSpecContext) LexerRuleSpec() ILexerRuleSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerRuleSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerRuleSpecContext)
}

func (s *RuleSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleSpec(s)
	}
}

func (s *RuleSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleSpec(s)
	}
}

func (s *RuleSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleSpec() (localctx IRuleSpecContext) {
	localctx = NewRuleSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ANTLRv4ParserRULE_ruleSpec)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.ParserRuleSpec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.LexerRuleSpec()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IParserRuleSpecContext is an interface to support dynamic dispatch.
type IParserRuleSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RULE_REF() antlr.TerminalNode
	COLON() antlr.TerminalNode
	RuleBlock() IRuleBlockContext
	SEMI() antlr.TerminalNode
	ExceptionGroup() IExceptionGroupContext
	RuleModifiers() IRuleModifiersContext
	ArgActionBlock() IArgActionBlockContext
	RuleReturns() IRuleReturnsContext
	ThrowsSpec() IThrowsSpecContext
	LocalsSpec() ILocalsSpecContext
	AllRulePrequel() []IRulePrequelContext
	RulePrequel(i int) IRulePrequelContext

	// IsParserRuleSpecContext differentiates from other interfaces.
	IsParserRuleSpecContext()
}

type ParserRuleSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParserRuleSpecContext() *ParserRuleSpecContext {
	var p = new(ParserRuleSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_parserRuleSpec
	return p
}

func InitEmptyParserRuleSpecContext(p *ParserRuleSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_parserRuleSpec
}

func (*ParserRuleSpecContext) IsParserRuleSpecContext() {}

func NewParserRuleSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParserRuleSpecContext {
	var p = new(ParserRuleSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_parserRuleSpec

	return p
}

func (s *ParserRuleSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ParserRuleSpecContext) RULE_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRULE_REF, 0)
}

func (s *ParserRuleSpecContext) COLON() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOLON, 0)
}

func (s *ParserRuleSpecContext) RuleBlock() IRuleBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleBlockContext)
}

func (s *ParserRuleSpecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSEMI, 0)
}

func (s *ParserRuleSpecContext) ExceptionGroup() IExceptionGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExceptionGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExceptionGroupContext)
}

func (s *ParserRuleSpecContext) RuleModifiers() IRuleModifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleModifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleModifiersContext)
}

func (s *ParserRuleSpecContext) ArgActionBlock() IArgActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgActionBlockContext)
}

func (s *ParserRuleSpecContext) RuleReturns() IRuleReturnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleReturnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleReturnsContext)
}

func (s *ParserRuleSpecContext) ThrowsSpec() IThrowsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThrowsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThrowsSpecContext)
}

func (s *ParserRuleSpecContext) LocalsSpec() ILocalsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalsSpecContext)
}

func (s *ParserRuleSpecContext) AllRulePrequel() []IRulePrequelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRulePrequelContext); ok {
			len++
		}
	}

	tst := make([]IRulePrequelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRulePrequelContext); ok {
			tst[i] = t.(IRulePrequelContext)
			i++
		}
	}

	return tst
}

func (s *ParserRuleSpecContext) RulePrequel(i int) IRulePrequelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulePrequelContext); ok {
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

	return t.(IRulePrequelContext)
}

func (s *ParserRuleSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParserRuleSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParserRuleSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterParserRuleSpec(s)
	}
}

func (s *ParserRuleSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitParserRuleSpec(s)
	}
}

func (s *ParserRuleSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitParserRuleSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ParserRuleSpec() (localctx IParserRuleSpecContext) {
	localctx = NewParserRuleSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ANTLRv4ParserRULE_parserRuleSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7405568) != 0 {
		{
			p.SetState(280)
			p.RuleModifiers()
		}

	}
	{
		p.SetState(283)
		p.Match(ANTLRv4ParserRULE_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserBEGIN_ARGUMENT {
		{
			p.SetState(284)
			p.ArgActionBlock()
		}

	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserRETURNS {
		{
			p.SetState(287)
			p.RuleReturns()
		}

	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserTHROWS {
		{
			p.SetState(290)
			p.ThrowsSpec()
		}

	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserLOCALS {
		{
			p.SetState(293)
			p.LocalsSpec()
		}

	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserOPTIONS || _la == ANTLRv4ParserAT {
		{
			p.SetState(296)
			p.RulePrequel()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(302)
		p.Match(ANTLRv4ParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.RuleBlock()
	}
	{
		p.SetState(304)
		p.Match(ANTLRv4ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.ExceptionGroup()
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

// IExceptionGroupContext is an interface to support dynamic dispatch.
type IExceptionGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExceptionHandler() []IExceptionHandlerContext
	ExceptionHandler(i int) IExceptionHandlerContext
	FinallyClause() IFinallyClauseContext

	// IsExceptionGroupContext differentiates from other interfaces.
	IsExceptionGroupContext()
}

type ExceptionGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExceptionGroupContext() *ExceptionGroupContext {
	var p = new(ExceptionGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_exceptionGroup
	return p
}

func InitEmptyExceptionGroupContext(p *ExceptionGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_exceptionGroup
}

func (*ExceptionGroupContext) IsExceptionGroupContext() {}

func NewExceptionGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptionGroupContext {
	var p = new(ExceptionGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_exceptionGroup

	return p
}

func (s *ExceptionGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ExceptionGroupContext) AllExceptionHandler() []IExceptionHandlerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExceptionHandlerContext); ok {
			len++
		}
	}

	tst := make([]IExceptionHandlerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExceptionHandlerContext); ok {
			tst[i] = t.(IExceptionHandlerContext)
			i++
		}
	}

	return tst
}

func (s *ExceptionGroupContext) ExceptionHandler(i int) IExceptionHandlerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExceptionHandlerContext); ok {
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

	return t.(IExceptionHandlerContext)
}

func (s *ExceptionGroupContext) FinallyClause() IFinallyClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFinallyClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFinallyClauseContext)
}

func (s *ExceptionGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExceptionGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExceptionGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterExceptionGroup(s)
	}
}

func (s *ExceptionGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitExceptionGroup(s)
	}
}

func (s *ExceptionGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitExceptionGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ExceptionGroup() (localctx IExceptionGroupContext) {
	localctx = NewExceptionGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ANTLRv4ParserRULE_exceptionGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserCATCH {
		{
			p.SetState(307)
			p.ExceptionHandler()
		}

		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserFINALLY {
		{
			p.SetState(313)
			p.FinallyClause()
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

// IExceptionHandlerContext is an interface to support dynamic dispatch.
type IExceptionHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CATCH() antlr.TerminalNode
	ArgActionBlock() IArgActionBlockContext
	ActionBlock() IActionBlockContext

	// IsExceptionHandlerContext differentiates from other interfaces.
	IsExceptionHandlerContext()
}

type ExceptionHandlerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExceptionHandlerContext() *ExceptionHandlerContext {
	var p = new(ExceptionHandlerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_exceptionHandler
	return p
}

func InitEmptyExceptionHandlerContext(p *ExceptionHandlerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_exceptionHandler
}

func (*ExceptionHandlerContext) IsExceptionHandlerContext() {}

func NewExceptionHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptionHandlerContext {
	var p = new(ExceptionHandlerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_exceptionHandler

	return p
}

func (s *ExceptionHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *ExceptionHandlerContext) CATCH() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCATCH, 0)
}

func (s *ExceptionHandlerContext) ArgActionBlock() IArgActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgActionBlockContext)
}

func (s *ExceptionHandlerContext) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *ExceptionHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExceptionHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExceptionHandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterExceptionHandler(s)
	}
}

func (s *ExceptionHandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitExceptionHandler(s)
	}
}

func (s *ExceptionHandlerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitExceptionHandler(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ExceptionHandler() (localctx IExceptionHandlerContext) {
	localctx = NewExceptionHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ANTLRv4ParserRULE_exceptionHandler)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(ANTLRv4ParserCATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.ArgActionBlock()
	}
	{
		p.SetState(318)
		p.ActionBlock()
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

// IFinallyClauseContext is an interface to support dynamic dispatch.
type IFinallyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FINALLY() antlr.TerminalNode
	ActionBlock() IActionBlockContext

	// IsFinallyClauseContext differentiates from other interfaces.
	IsFinallyClauseContext()
}

type FinallyClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinallyClauseContext() *FinallyClauseContext {
	var p = new(FinallyClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_finallyClause
	return p
}

func InitEmptyFinallyClauseContext(p *FinallyClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_finallyClause
}

func (*FinallyClauseContext) IsFinallyClauseContext() {}

func NewFinallyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinallyClauseContext {
	var p = new(FinallyClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_finallyClause

	return p
}

func (s *FinallyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FinallyClauseContext) FINALLY() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserFINALLY, 0)
}

func (s *FinallyClauseContext) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *FinallyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinallyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinallyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterFinallyClause(s)
	}
}

func (s *FinallyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitFinallyClause(s)
	}
}

func (s *FinallyClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitFinallyClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) FinallyClause() (localctx IFinallyClauseContext) {
	localctx = NewFinallyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ANTLRv4ParserRULE_finallyClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(ANTLRv4ParserFINALLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.ActionBlock()
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

// IRulePrequelContext is an interface to support dynamic dispatch.
type IRulePrequelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionsSpec() IOptionsSpecContext
	RuleAction() IRuleActionContext

	// IsRulePrequelContext differentiates from other interfaces.
	IsRulePrequelContext()
}

type RulePrequelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulePrequelContext() *RulePrequelContext {
	var p = new(RulePrequelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_rulePrequel
	return p
}

func InitEmptyRulePrequelContext(p *RulePrequelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_rulePrequel
}

func (*RulePrequelContext) IsRulePrequelContext() {}

func NewRulePrequelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulePrequelContext {
	var p = new(RulePrequelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_rulePrequel

	return p
}

func (s *RulePrequelContext) GetParser() antlr.Parser { return s.parser }

func (s *RulePrequelContext) OptionsSpec() IOptionsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionsSpecContext)
}

func (s *RulePrequelContext) RuleAction() IRuleActionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleActionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleActionContext)
}

func (s *RulePrequelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulePrequelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulePrequelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRulePrequel(s)
	}
}

func (s *RulePrequelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRulePrequel(s)
	}
}

func (s *RulePrequelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRulePrequel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RulePrequel() (localctx IRulePrequelContext) {
	localctx = NewRulePrequelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ANTLRv4ParserRULE_rulePrequel)
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserOPTIONS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.OptionsSpec()
		}

	case ANTLRv4ParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.RuleAction()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IRuleReturnsContext is an interface to support dynamic dispatch.
type IRuleReturnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURNS() antlr.TerminalNode
	ArgActionBlock() IArgActionBlockContext

	// IsRuleReturnsContext differentiates from other interfaces.
	IsRuleReturnsContext()
}

type RuleReturnsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleReturnsContext() *RuleReturnsContext {
	var p = new(RuleReturnsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleReturns
	return p
}

func InitEmptyRuleReturnsContext(p *RuleReturnsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleReturns
}

func (*RuleReturnsContext) IsRuleReturnsContext() {}

func NewRuleReturnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleReturnsContext {
	var p = new(RuleReturnsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleReturns

	return p
}

func (s *RuleReturnsContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleReturnsContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRETURNS, 0)
}

func (s *RuleReturnsContext) ArgActionBlock() IArgActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgActionBlockContext)
}

func (s *RuleReturnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleReturnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleReturnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleReturns(s)
	}
}

func (s *RuleReturnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleReturns(s)
	}
}

func (s *RuleReturnsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleReturns(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleReturns() (localctx IRuleReturnsContext) {
	localctx = NewRuleReturnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ANTLRv4ParserRULE_ruleReturns)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(ANTLRv4ParserRETURNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.ArgActionBlock()
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

// IThrowsSpecContext is an interface to support dynamic dispatch.
type IThrowsSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	THROWS() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsThrowsSpecContext differentiates from other interfaces.
	IsThrowsSpecContext()
}

type ThrowsSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowsSpecContext() *ThrowsSpecContext {
	var p = new(ThrowsSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_throwsSpec
	return p
}

func InitEmptyThrowsSpecContext(p *ThrowsSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_throwsSpec
}

func (*ThrowsSpecContext) IsThrowsSpecContext() {}

func NewThrowsSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowsSpecContext {
	var p = new(ThrowsSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_throwsSpec

	return p
}

func (s *ThrowsSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowsSpecContext) THROWS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserTHROWS, 0)
}

func (s *ThrowsSpecContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ThrowsSpecContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *ThrowsSpecContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserCOMMA)
}

func (s *ThrowsSpecContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOMMA, i)
}

func (s *ThrowsSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowsSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowsSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterThrowsSpec(s)
	}
}

func (s *ThrowsSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitThrowsSpec(s)
	}
}

func (s *ThrowsSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitThrowsSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ThrowsSpec() (localctx IThrowsSpecContext) {
	localctx = NewThrowsSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ANTLRv4ParserRULE_throwsSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(ANTLRv4ParserTHROWS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Identifier()
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserCOMMA {
		{
			p.SetState(332)
			p.Match(ANTLRv4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Identifier()
		}

		p.SetState(338)
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

// ILocalsSpecContext is an interface to support dynamic dispatch.
type ILocalsSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOCALS() antlr.TerminalNode
	ArgActionBlock() IArgActionBlockContext

	// IsLocalsSpecContext differentiates from other interfaces.
	IsLocalsSpecContext()
}

type LocalsSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalsSpecContext() *LocalsSpecContext {
	var p = new(LocalsSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_localsSpec
	return p
}

func InitEmptyLocalsSpecContext(p *LocalsSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_localsSpec
}

func (*LocalsSpecContext) IsLocalsSpecContext() {}

func NewLocalsSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalsSpecContext {
	var p = new(LocalsSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_localsSpec

	return p
}

func (s *LocalsSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalsSpecContext) LOCALS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLOCALS, 0)
}

func (s *LocalsSpecContext) ArgActionBlock() IArgActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgActionBlockContext)
}

func (s *LocalsSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalsSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalsSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLocalsSpec(s)
	}
}

func (s *LocalsSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLocalsSpec(s)
	}
}

func (s *LocalsSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLocalsSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LocalsSpec() (localctx ILocalsSpecContext) {
	localctx = NewLocalsSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ANTLRv4ParserRULE_localsSpec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(ANTLRv4ParserLOCALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.ArgActionBlock()
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

// IRuleActionContext is an interface to support dynamic dispatch.
type IRuleActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	Identifier() IIdentifierContext
	ActionBlock() IActionBlockContext

	// IsRuleActionContext differentiates from other interfaces.
	IsRuleActionContext()
}

type RuleActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleActionContext() *RuleActionContext {
	var p = new(RuleActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleAction
	return p
}

func InitEmptyRuleActionContext(p *RuleActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleAction
}

func (*RuleActionContext) IsRuleActionContext() {}

func NewRuleActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleActionContext {
	var p = new(RuleActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleAction

	return p
}

func (s *RuleActionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleActionContext) AT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserAT, 0)
}

func (s *RuleActionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RuleActionContext) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *RuleActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleAction(s)
	}
}

func (s *RuleActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleAction(s)
	}
}

func (s *RuleActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleAction() (localctx IRuleActionContext) {
	localctx = NewRuleActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ANTLRv4ParserRULE_ruleAction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(ANTLRv4ParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Identifier()
	}
	{
		p.SetState(344)
		p.ActionBlock()
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

// IRuleModifiersContext is an interface to support dynamic dispatch.
type IRuleModifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRuleModifier() []IRuleModifierContext
	RuleModifier(i int) IRuleModifierContext

	// IsRuleModifiersContext differentiates from other interfaces.
	IsRuleModifiersContext()
}

type RuleModifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleModifiersContext() *RuleModifiersContext {
	var p = new(RuleModifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleModifiers
	return p
}

func InitEmptyRuleModifiersContext(p *RuleModifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleModifiers
}

func (*RuleModifiersContext) IsRuleModifiersContext() {}

func NewRuleModifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleModifiersContext {
	var p = new(RuleModifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleModifiers

	return p
}

func (s *RuleModifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleModifiersContext) AllRuleModifier() []IRuleModifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleModifierContext); ok {
			len++
		}
	}

	tst := make([]IRuleModifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleModifierContext); ok {
			tst[i] = t.(IRuleModifierContext)
			i++
		}
	}

	return tst
}

func (s *RuleModifiersContext) RuleModifier(i int) IRuleModifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleModifierContext); ok {
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

	return t.(IRuleModifierContext)
}

func (s *RuleModifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleModifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleModifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleModifiers(s)
	}
}

func (s *RuleModifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleModifiers(s)
	}
}

func (s *RuleModifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleModifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleModifiers() (localctx IRuleModifiersContext) {
	localctx = NewRuleModifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ANTLRv4ParserRULE_ruleModifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7405568) != 0) {
		{
			p.SetState(346)
			p.RuleModifier()
		}

		p.SetState(349)
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

// IRuleModifierContext is an interface to support dynamic dispatch.
type IRuleModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUBLIC() antlr.TerminalNode
	PRIVATE() antlr.TerminalNode
	PROTECTED() antlr.TerminalNode
	FRAGMENT() antlr.TerminalNode

	// IsRuleModifierContext differentiates from other interfaces.
	IsRuleModifierContext()
}

type RuleModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleModifierContext() *RuleModifierContext {
	var p = new(RuleModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleModifier
	return p
}

func InitEmptyRuleModifierContext(p *RuleModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleModifier
}

func (*RuleModifierContext) IsRuleModifierContext() {}

func NewRuleModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleModifierContext {
	var p = new(RuleModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleModifier

	return p
}

func (s *RuleModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleModifierContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPUBLIC, 0)
}

func (s *RuleModifierContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPRIVATE, 0)
}

func (s *RuleModifierContext) PROTECTED() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPROTECTED, 0)
}

func (s *RuleModifierContext) FRAGMENT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserFRAGMENT, 0)
}

func (s *RuleModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleModifier(s)
	}
}

func (s *RuleModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleModifier(s)
	}
}

func (s *RuleModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleModifier() (localctx IRuleModifierContext) {
	localctx = NewRuleModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ANTLRv4ParserRULE_ruleModifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7405568) != 0) {
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

// IRuleBlockContext is an interface to support dynamic dispatch.
type IRuleBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RuleAltList() IRuleAltListContext

	// IsRuleBlockContext differentiates from other interfaces.
	IsRuleBlockContext()
}

type RuleBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleBlockContext() *RuleBlockContext {
	var p = new(RuleBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleBlock
	return p
}

func InitEmptyRuleBlockContext(p *RuleBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleBlock
}

func (*RuleBlockContext) IsRuleBlockContext() {}

func NewRuleBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleBlockContext {
	var p = new(RuleBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleBlock

	return p
}

func (s *RuleBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleBlockContext) RuleAltList() IRuleAltListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleAltListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleAltListContext)
}

func (s *RuleBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleBlock(s)
	}
}

func (s *RuleBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleBlock(s)
	}
}

func (s *RuleBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleBlock() (localctx IRuleBlockContext) {
	localctx = NewRuleBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ANTLRv4ParserRULE_ruleBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.RuleAltList()
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

// IRuleAltListContext is an interface to support dynamic dispatch.
type IRuleAltListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabeledAlt() []ILabeledAltContext
	LabeledAlt(i int) ILabeledAltContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsRuleAltListContext differentiates from other interfaces.
	IsRuleAltListContext()
}

type RuleAltListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleAltListContext() *RuleAltListContext {
	var p = new(RuleAltListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleAltList
	return p
}

func InitEmptyRuleAltListContext(p *RuleAltListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleAltList
}

func (*RuleAltListContext) IsRuleAltListContext() {}

func NewRuleAltListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleAltListContext {
	var p = new(RuleAltListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleAltList

	return p
}

func (s *RuleAltListContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleAltListContext) AllLabeledAlt() []ILabeledAltContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabeledAltContext); ok {
			len++
		}
	}

	tst := make([]ILabeledAltContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabeledAltContext); ok {
			tst[i] = t.(ILabeledAltContext)
			i++
		}
	}

	return tst
}

func (s *RuleAltListContext) LabeledAlt(i int) ILabeledAltContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabeledAltContext); ok {
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

	return t.(ILabeledAltContext)
}

func (s *RuleAltListContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserOR)
}

func (s *RuleAltListContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserOR, i)
}

func (s *RuleAltListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleAltListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleAltListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleAltList(s)
	}
}

func (s *RuleAltListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleAltList(s)
	}
}

func (s *RuleAltListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleAltList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) RuleAltList() (localctx IRuleAltListContext) {
	localctx = NewRuleAltListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ANTLRv4ParserRULE_ruleAltList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.LabeledAlt()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserOR {
		{
			p.SetState(356)
			p.Match(ANTLRv4ParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.LabeledAlt()
		}

		p.SetState(362)
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

// ILabeledAltContext is an interface to support dynamic dispatch.
type ILabeledAltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Alternative() IAlternativeContext
	POUND() antlr.TerminalNode
	Identifier() IIdentifierContext

	// IsLabeledAltContext differentiates from other interfaces.
	IsLabeledAltContext()
}

type LabeledAltContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeledAltContext() *LabeledAltContext {
	var p = new(LabeledAltContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_labeledAlt
	return p
}

func InitEmptyLabeledAltContext(p *LabeledAltContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_labeledAlt
}

func (*LabeledAltContext) IsLabeledAltContext() {}

func NewLabeledAltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabeledAltContext {
	var p = new(LabeledAltContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_labeledAlt

	return p
}

func (s *LabeledAltContext) GetParser() antlr.Parser { return s.parser }

func (s *LabeledAltContext) Alternative() IAlternativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlternativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlternativeContext)
}

func (s *LabeledAltContext) POUND() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPOUND, 0)
}

func (s *LabeledAltContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LabeledAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabeledAltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabeledAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLabeledAlt(s)
	}
}

func (s *LabeledAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLabeledAlt(s)
	}
}

func (s *LabeledAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLabeledAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LabeledAlt() (localctx ILabeledAltContext) {
	localctx = NewLabeledAltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ANTLRv4ParserRULE_labeledAlt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Alternative()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserPOUND {
		{
			p.SetState(364)
			p.Match(ANTLRv4ParserPOUND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Identifier()
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

// ILexerRuleSpecContext is an interface to support dynamic dispatch.
type ILexerRuleSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TOKEN_REF() antlr.TerminalNode
	COLON() antlr.TerminalNode
	LexerRuleBlock() ILexerRuleBlockContext
	SEMI() antlr.TerminalNode
	FRAGMENT() antlr.TerminalNode
	OptionsSpec() IOptionsSpecContext

	// IsLexerRuleSpecContext differentiates from other interfaces.
	IsLexerRuleSpecContext()
}

type LexerRuleSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerRuleSpecContext() *LexerRuleSpecContext {
	var p = new(LexerRuleSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerRuleSpec
	return p
}

func InitEmptyLexerRuleSpecContext(p *LexerRuleSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerRuleSpec
}

func (*LexerRuleSpecContext) IsLexerRuleSpecContext() {}

func NewLexerRuleSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerRuleSpecContext {
	var p = new(LexerRuleSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerRuleSpec

	return p
}

func (s *LexerRuleSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerRuleSpecContext) TOKEN_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserTOKEN_REF, 0)
}

func (s *LexerRuleSpecContext) COLON() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOLON, 0)
}

func (s *LexerRuleSpecContext) LexerRuleBlock() ILexerRuleBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerRuleBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerRuleBlockContext)
}

func (s *LexerRuleSpecContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSEMI, 0)
}

func (s *LexerRuleSpecContext) FRAGMENT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserFRAGMENT, 0)
}

func (s *LexerRuleSpecContext) OptionsSpec() IOptionsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionsSpecContext)
}

func (s *LexerRuleSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerRuleSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerRuleSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerRuleSpec(s)
	}
}

func (s *LexerRuleSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerRuleSpec(s)
	}
}

func (s *LexerRuleSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerRuleSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerRuleSpec() (localctx ILexerRuleSpecContext) {
	localctx = NewLexerRuleSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ANTLRv4ParserRULE_lexerRuleSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserFRAGMENT {
		{
			p.SetState(368)
			p.Match(ANTLRv4ParserFRAGMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(371)
		p.Match(ANTLRv4ParserTOKEN_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserOPTIONS {
		{
			p.SetState(372)
			p.OptionsSpec()
		}

	}
	{
		p.SetState(375)
		p.Match(ANTLRv4ParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.LexerRuleBlock()
	}
	{
		p.SetState(377)
		p.Match(ANTLRv4ParserSEMI)
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

// ILexerRuleBlockContext is an interface to support dynamic dispatch.
type ILexerRuleBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LexerAltList() ILexerAltListContext

	// IsLexerRuleBlockContext differentiates from other interfaces.
	IsLexerRuleBlockContext()
}

type LexerRuleBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerRuleBlockContext() *LexerRuleBlockContext {
	var p = new(LexerRuleBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerRuleBlock
	return p
}

func InitEmptyLexerRuleBlockContext(p *LexerRuleBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerRuleBlock
}

func (*LexerRuleBlockContext) IsLexerRuleBlockContext() {}

func NewLexerRuleBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerRuleBlockContext {
	var p = new(LexerRuleBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerRuleBlock

	return p
}

func (s *LexerRuleBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerRuleBlockContext) LexerAltList() ILexerAltListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerAltListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerAltListContext)
}

func (s *LexerRuleBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerRuleBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerRuleBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerRuleBlock(s)
	}
}

func (s *LexerRuleBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerRuleBlock(s)
	}
}

func (s *LexerRuleBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerRuleBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerRuleBlock() (localctx ILexerRuleBlockContext) {
	localctx = NewLexerRuleBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ANTLRv4ParserRULE_lexerRuleBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.LexerAltList()
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

// ILexerAltListContext is an interface to support dynamic dispatch.
type ILexerAltListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLexerAlt() []ILexerAltContext
	LexerAlt(i int) ILexerAltContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLexerAltListContext differentiates from other interfaces.
	IsLexerAltListContext()
}

type LexerAltListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerAltListContext() *LexerAltListContext {
	var p = new(LexerAltListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerAltList
	return p
}

func InitEmptyLexerAltListContext(p *LexerAltListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerAltList
}

func (*LexerAltListContext) IsLexerAltListContext() {}

func NewLexerAltListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerAltListContext {
	var p = new(LexerAltListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerAltList

	return p
}

func (s *LexerAltListContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerAltListContext) AllLexerAlt() []ILexerAltContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILexerAltContext); ok {
			len++
		}
	}

	tst := make([]ILexerAltContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILexerAltContext); ok {
			tst[i] = t.(ILexerAltContext)
			i++
		}
	}

	return tst
}

func (s *LexerAltListContext) LexerAlt(i int) ILexerAltContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerAltContext); ok {
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

	return t.(ILexerAltContext)
}

func (s *LexerAltListContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserOR)
}

func (s *LexerAltListContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserOR, i)
}

func (s *LexerAltListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerAltListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerAltListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerAltList(s)
	}
}

func (s *LexerAltListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerAltList(s)
	}
}

func (s *LexerAltListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerAltList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerAltList() (localctx ILexerAltListContext) {
	localctx = NewLexerAltListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ANTLRv4ParserRULE_lexerAltList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.LexerAlt()
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserOR {
		{
			p.SetState(382)
			p.Match(ANTLRv4ParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.LexerAlt()
		}

		p.SetState(388)
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

// ILexerAltContext is an interface to support dynamic dispatch.
type ILexerAltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LexerElements() ILexerElementsContext
	LexerCommands() ILexerCommandsContext

	// IsLexerAltContext differentiates from other interfaces.
	IsLexerAltContext()
}

type LexerAltContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerAltContext() *LexerAltContext {
	var p = new(LexerAltContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerAlt
	return p
}

func InitEmptyLexerAltContext(p *LexerAltContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerAlt
}

func (*LexerAltContext) IsLexerAltContext() {}

func NewLexerAltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerAltContext {
	var p = new(LexerAltContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerAlt

	return p
}

func (s *LexerAltContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerAltContext) LexerElements() ILexerElementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerElementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerElementsContext)
}

func (s *LexerAltContext) LexerCommands() ILexerCommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerCommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerCommandsContext)
}

func (s *LexerAltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerAltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerAltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerAlt(s)
	}
}

func (s *LexerAltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerAlt(s)
	}
}

func (s *LexerAltContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerAlt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerAlt() (localctx ILexerAltContext) {
	localctx = NewLexerAltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ANTLRv4ParserRULE_lexerAlt)
	var _la int

	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.LexerElements()
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserRARROW {
			{
				p.SetState(390)
				p.LexerCommands()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// ILexerElementsContext is an interface to support dynamic dispatch.
type ILexerElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLexerElement() []ILexerElementContext
	LexerElement(i int) ILexerElementContext

	// IsLexerElementsContext differentiates from other interfaces.
	IsLexerElementsContext()
}

type LexerElementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerElementsContext() *LexerElementsContext {
	var p = new(LexerElementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerElements
	return p
}

func InitEmptyLexerElementsContext(p *LexerElementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerElements
}

func (*LexerElementsContext) IsLexerElementsContext() {}

func NewLexerElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerElementsContext {
	var p = new(LexerElementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerElements

	return p
}

func (s *LexerElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerElementsContext) AllLexerElement() []ILexerElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILexerElementContext); ok {
			len++
		}
	}

	tst := make([]ILexerElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILexerElementContext); ok {
			tst[i] = t.(ILexerElementContext)
			i++
		}
	}

	return tst
}

func (s *LexerElementsContext) LexerElement(i int) ILexerElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerElementContext); ok {
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

	return t.(ILexerElementContext)
}

func (s *LexerElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerElements(s)
	}
}

func (s *LexerElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerElements(s)
	}
}

func (s *LexerElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerElements() (localctx ILexerElementsContext) {
	localctx = NewLexerElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ANTLRv4ParserRULE_lexerElements)
	var _la int

	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserLEXER_CHAR_SET, ANTLRv4ParserSTRING_LITERAL, ANTLRv4ParserBEGIN_ACTION, ANTLRv4ParserLPAREN, ANTLRv4ParserDOT, ANTLRv4ParserNOT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2533283380332810) != 0) {
			{
				p.SetState(396)
				p.LexerElement()
			}

			p.SetState(399)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ANTLRv4ParserSEMI, ANTLRv4ParserRPAREN, ANTLRv4ParserRARROW, ANTLRv4ParserOR:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// ILexerElementContext is an interface to support dynamic dispatch.
type ILexerElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LexerAtom() ILexerAtomContext
	EbnfSuffix() IEbnfSuffixContext
	LexerBlock() ILexerBlockContext
	ActionBlock() IActionBlockContext
	QUESTION() antlr.TerminalNode

	// IsLexerElementContext differentiates from other interfaces.
	IsLexerElementContext()
}

type LexerElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerElementContext() *LexerElementContext {
	var p = new(LexerElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerElement
	return p
}

func InitEmptyLexerElementContext(p *LexerElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerElement
}

func (*LexerElementContext) IsLexerElementContext() {}

func NewLexerElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerElementContext {
	var p = new(LexerElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerElement

	return p
}

func (s *LexerElementContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerElementContext) LexerAtom() ILexerAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerAtomContext)
}

func (s *LexerElementContext) EbnfSuffix() IEbnfSuffixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEbnfSuffixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEbnfSuffixContext)
}

func (s *LexerElementContext) LexerBlock() ILexerBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerBlockContext)
}

func (s *LexerElementContext) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *LexerElementContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserQUESTION, 0)
}

func (s *LexerElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerElement(s)
	}
}

func (s *LexerElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerElement(s)
	}
}

func (s *LexerElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerElement() (localctx ILexerElementContext) {
	localctx = NewLexerElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ANTLRv4ParserRULE_lexerElement)
	var _la int

	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserLEXER_CHAR_SET, ANTLRv4ParserSTRING_LITERAL, ANTLRv4ParserDOT, ANTLRv4ParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.LexerAtom()
		}
		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&24189255811072) != 0 {
			{
				p.SetState(405)
				p.EbnfSuffix()
			}

		}

	case ANTLRv4ParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)
			p.LexerBlock()
		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&24189255811072) != 0 {
			{
				p.SetState(409)
				p.EbnfSuffix()
			}

		}

	case ANTLRv4ParserBEGIN_ACTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(412)
			p.ActionBlock()
		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserQUESTION {
			{
				p.SetState(413)
				p.Match(ANTLRv4ParserQUESTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// ILexerBlockContext is an interface to support dynamic dispatch.
type ILexerBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	LexerAltList() ILexerAltListContext
	RPAREN() antlr.TerminalNode

	// IsLexerBlockContext differentiates from other interfaces.
	IsLexerBlockContext()
}

type LexerBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerBlockContext() *LexerBlockContext {
	var p = new(LexerBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerBlock
	return p
}

func InitEmptyLexerBlockContext(p *LexerBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerBlock
}

func (*LexerBlockContext) IsLexerBlockContext() {}

func NewLexerBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerBlockContext {
	var p = new(LexerBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerBlock

	return p
}

func (s *LexerBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerBlockContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLPAREN, 0)
}

func (s *LexerBlockContext) LexerAltList() ILexerAltListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerAltListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerAltListContext)
}

func (s *LexerBlockContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRPAREN, 0)
}

func (s *LexerBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerBlock(s)
	}
}

func (s *LexerBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerBlock(s)
	}
}

func (s *LexerBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerBlock() (localctx ILexerBlockContext) {
	localctx = NewLexerBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ANTLRv4ParserRULE_lexerBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(ANTLRv4ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.LexerAltList()
	}
	{
		p.SetState(420)
		p.Match(ANTLRv4ParserRPAREN)
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

// ILexerCommandsContext is an interface to support dynamic dispatch.
type ILexerCommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RARROW() antlr.TerminalNode
	AllLexerCommand() []ILexerCommandContext
	LexerCommand(i int) ILexerCommandContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLexerCommandsContext differentiates from other interfaces.
	IsLexerCommandsContext()
}

type LexerCommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerCommandsContext() *LexerCommandsContext {
	var p = new(LexerCommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommands
	return p
}

func InitEmptyLexerCommandsContext(p *LexerCommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommands
}

func (*LexerCommandsContext) IsLexerCommandsContext() {}

func NewLexerCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerCommandsContext {
	var p = new(LexerCommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommands

	return p
}

func (s *LexerCommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerCommandsContext) RARROW() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRARROW, 0)
}

func (s *LexerCommandsContext) AllLexerCommand() []ILexerCommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILexerCommandContext); ok {
			len++
		}
	}

	tst := make([]ILexerCommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILexerCommandContext); ok {
			tst[i] = t.(ILexerCommandContext)
			i++
		}
	}

	return tst
}

func (s *LexerCommandsContext) LexerCommand(i int) ILexerCommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerCommandContext); ok {
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

	return t.(ILexerCommandContext)
}

func (s *LexerCommandsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserCOMMA)
}

func (s *LexerCommandsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOMMA, i)
}

func (s *LexerCommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerCommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerCommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerCommands(s)
	}
}

func (s *LexerCommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerCommands(s)
	}
}

func (s *LexerCommandsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerCommands(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerCommands() (localctx ILexerCommandsContext) {
	localctx = NewLexerCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ANTLRv4ParserRULE_lexerCommands)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(ANTLRv4ParserRARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.LexerCommand()
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserCOMMA {
		{
			p.SetState(424)
			p.Match(ANTLRv4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.LexerCommand()
		}

		p.SetState(430)
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

// ILexerCommandContext is an interface to support dynamic dispatch.
type ILexerCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LexerCommandName() ILexerCommandNameContext
	LPAREN() antlr.TerminalNode
	LexerCommandExpr() ILexerCommandExprContext
	RPAREN() antlr.TerminalNode

	// IsLexerCommandContext differentiates from other interfaces.
	IsLexerCommandContext()
}

type LexerCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerCommandContext() *LexerCommandContext {
	var p = new(LexerCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommand
	return p
}

func InitEmptyLexerCommandContext(p *LexerCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommand
}

func (*LexerCommandContext) IsLexerCommandContext() {}

func NewLexerCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerCommandContext {
	var p = new(LexerCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommand

	return p
}

func (s *LexerCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerCommandContext) LexerCommandName() ILexerCommandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerCommandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerCommandNameContext)
}

func (s *LexerCommandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLPAREN, 0)
}

func (s *LexerCommandContext) LexerCommandExpr() ILexerCommandExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILexerCommandExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILexerCommandExprContext)
}

func (s *LexerCommandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRPAREN, 0)
}

func (s *LexerCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerCommand(s)
	}
}

func (s *LexerCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerCommand(s)
	}
}

func (s *LexerCommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerCommand() (localctx ILexerCommandContext) {
	localctx = NewLexerCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ANTLRv4ParserRULE_lexerCommand)
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
			p.LexerCommandName()
		}
		{
			p.SetState(432)
			p.Match(ANTLRv4ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.LexerCommandExpr()
		}
		{
			p.SetState(434)
			p.Match(ANTLRv4ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.LexerCommandName()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// ILexerCommandNameContext is an interface to support dynamic dispatch.
type ILexerCommandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	MODE() antlr.TerminalNode

	// IsLexerCommandNameContext differentiates from other interfaces.
	IsLexerCommandNameContext()
}

type LexerCommandNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerCommandNameContext() *LexerCommandNameContext {
	var p = new(LexerCommandNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommandName
	return p
}

func InitEmptyLexerCommandNameContext(p *LexerCommandNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommandName
}

func (*LexerCommandNameContext) IsLexerCommandNameContext() {}

func NewLexerCommandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerCommandNameContext {
	var p = new(LexerCommandNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommandName

	return p
}

func (s *LexerCommandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerCommandNameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LexerCommandNameContext) MODE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserMODE, 0)
}

func (s *LexerCommandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerCommandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerCommandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerCommandName(s)
	}
}

func (s *LexerCommandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerCommandName(s)
	}
}

func (s *LexerCommandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerCommandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerCommandName() (localctx ILexerCommandNameContext) {
	localctx = NewLexerCommandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ANTLRv4ParserRULE_lexerCommandName)
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(439)
			p.Identifier()
		}

	case ANTLRv4ParserMODE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.Match(ANTLRv4ParserMODE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// ILexerCommandExprContext is an interface to support dynamic dispatch.
type ILexerCommandExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	INT() antlr.TerminalNode

	// IsLexerCommandExprContext differentiates from other interfaces.
	IsLexerCommandExprContext()
}

type LexerCommandExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerCommandExprContext() *LexerCommandExprContext {
	var p = new(LexerCommandExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommandExpr
	return p
}

func InitEmptyLexerCommandExprContext(p *LexerCommandExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommandExpr
}

func (*LexerCommandExprContext) IsLexerCommandExprContext() {}

func NewLexerCommandExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerCommandExprContext {
	var p = new(LexerCommandExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerCommandExpr

	return p
}

func (s *LexerCommandExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerCommandExprContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LexerCommandExprContext) INT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserINT, 0)
}

func (s *LexerCommandExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerCommandExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerCommandExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerCommandExpr(s)
	}
}

func (s *LexerCommandExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerCommandExpr(s)
	}
}

func (s *LexerCommandExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerCommandExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerCommandExpr() (localctx ILexerCommandExprContext) {
	localctx = NewLexerCommandExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ANTLRv4ParserRULE_lexerCommandExpr)
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)
			p.Identifier()
		}

	case ANTLRv4ParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)
			p.Match(ANTLRv4ParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IAltListContext is an interface to support dynamic dispatch.
type IAltListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAlternative() []IAlternativeContext
	Alternative(i int) IAlternativeContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsAltListContext differentiates from other interfaces.
	IsAltListContext()
}

type AltListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAltListContext() *AltListContext {
	var p = new(AltListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_altList
	return p
}

func InitEmptyAltListContext(p *AltListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_altList
}

func (*AltListContext) IsAltListContext() {}

func NewAltListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AltListContext {
	var p = new(AltListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_altList

	return p
}

func (s *AltListContext) GetParser() antlr.Parser { return s.parser }

func (s *AltListContext) AllAlternative() []IAlternativeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAlternativeContext); ok {
			len++
		}
	}

	tst := make([]IAlternativeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAlternativeContext); ok {
			tst[i] = t.(IAlternativeContext)
			i++
		}
	}

	return tst
}

func (s *AltListContext) Alternative(i int) IAlternativeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlternativeContext); ok {
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

	return t.(IAlternativeContext)
}

func (s *AltListContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserOR)
}

func (s *AltListContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserOR, i)
}

func (s *AltListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AltListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AltListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterAltList(s)
	}
}

func (s *AltListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitAltList(s)
	}
}

func (s *AltListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitAltList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) AltList() (localctx IAltListContext) {
	localctx = NewAltListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ANTLRv4ParserRULE_altList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.Alternative()
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserOR {
		{
			p.SetState(448)
			p.Match(ANTLRv4ParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Alternative()
		}

		p.SetState(454)
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

// IAlternativeContext is an interface to support dynamic dispatch.
type IAlternativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElementOptions() IElementOptionsContext
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsAlternativeContext differentiates from other interfaces.
	IsAlternativeContext()
}

type AlternativeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlternativeContext() *AlternativeContext {
	var p = new(AlternativeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_alternative
	return p
}

func InitEmptyAlternativeContext(p *AlternativeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_alternative
}

func (*AlternativeContext) IsAlternativeContext() {}

func NewAlternativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlternativeContext {
	var p = new(AlternativeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_alternative

	return p
}

func (s *AlternativeContext) GetParser() antlr.Parser { return s.parser }

func (s *AlternativeContext) ElementOptions() IElementOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementOptionsContext)
}

func (s *AlternativeContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *AlternativeContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *AlternativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlternativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterAlternative(s)
	}
}

func (s *AlternativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitAlternative(s)
	}
}

func (s *AlternativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitAlternative(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Alternative() (localctx IAlternativeContext) {
	localctx = NewAlternativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ANTLRv4ParserRULE_alternative)
	var _la int

	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF, ANTLRv4ParserSTRING_LITERAL, ANTLRv4ParserBEGIN_ACTION, ANTLRv4ParserLPAREN, ANTLRv4ParserLT, ANTLRv4ParserDOT, ANTLRv4ParserNOT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(455)
				p.ElementOptions()
			}

		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2533283380332806) != 0) {
			{
				p.SetState(458)
				p.Element()
			}

			p.SetState(461)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ANTLRv4ParserSEMI, ANTLRv4ParserRPAREN, ANTLRv4ParserOR, ANTLRv4ParserPOUND:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabeledElement() ILabeledElementContext
	EbnfSuffix() IEbnfSuffixContext
	Atom() IAtomContext
	Ebnf() IEbnfContext
	ActionBlock() IActionBlockContext
	QUESTION() antlr.TerminalNode

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_element
	return p
}

func InitEmptyElementContext(p *ElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_element
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) LabeledElement() ILabeledElementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabeledElementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabeledElementContext)
}

func (s *ElementContext) EbnfSuffix() IEbnfSuffixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEbnfSuffixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEbnfSuffixContext)
}

func (s *ElementContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ElementContext) Ebnf() IEbnfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEbnfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEbnfContext)
}

func (s *ElementContext) ActionBlock() IActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionBlockContext)
}

func (s *ElementContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserQUESTION, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ANTLRv4ParserRULE_element)
	var _la int

	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)
			p.LabeledElement()
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ANTLRv4ParserQUESTION, ANTLRv4ParserSTAR, ANTLRv4ParserPLUS:
			{
				p.SetState(467)
				p.EbnfSuffix()
			}

		case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF, ANTLRv4ParserSTRING_LITERAL, ANTLRv4ParserBEGIN_ACTION, ANTLRv4ParserSEMI, ANTLRv4ParserLPAREN, ANTLRv4ParserRPAREN, ANTLRv4ParserOR, ANTLRv4ParserDOT, ANTLRv4ParserPOUND, ANTLRv4ParserNOT:

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(471)
			p.Atom()
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ANTLRv4ParserQUESTION, ANTLRv4ParserSTAR, ANTLRv4ParserPLUS:
			{
				p.SetState(472)
				p.EbnfSuffix()
			}

		case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF, ANTLRv4ParserSTRING_LITERAL, ANTLRv4ParserBEGIN_ACTION, ANTLRv4ParserSEMI, ANTLRv4ParserLPAREN, ANTLRv4ParserRPAREN, ANTLRv4ParserOR, ANTLRv4ParserDOT, ANTLRv4ParserPOUND, ANTLRv4ParserNOT:

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(476)
			p.Ebnf()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(477)
			p.ActionBlock()
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserQUESTION {
			{
				p.SetState(478)
				p.Match(ANTLRv4ParserQUESTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// ILabeledElementContext is an interface to support dynamic dispatch.
type ILabeledElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	ASSIGN() antlr.TerminalNode
	PLUS_ASSIGN() antlr.TerminalNode
	Atom() IAtomContext
	Block() IBlockContext

	// IsLabeledElementContext differentiates from other interfaces.
	IsLabeledElementContext()
}

type LabeledElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeledElementContext() *LabeledElementContext {
	var p = new(LabeledElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_labeledElement
	return p
}

func InitEmptyLabeledElementContext(p *LabeledElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_labeledElement
}

func (*LabeledElementContext) IsLabeledElementContext() {}

func NewLabeledElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabeledElementContext {
	var p = new(LabeledElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_labeledElement

	return p
}

func (s *LabeledElementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabeledElementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LabeledElementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserASSIGN, 0)
}

func (s *LabeledElementContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPLUS_ASSIGN, 0)
}

func (s *LabeledElementContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *LabeledElementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LabeledElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabeledElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabeledElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLabeledElement(s)
	}
}

func (s *LabeledElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLabeledElement(s)
	}
}

func (s *LabeledElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLabeledElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LabeledElement() (localctx ILabeledElementContext) {
	localctx = NewLabeledElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ANTLRv4ParserRULE_labeledElement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Identifier()
	}
	{
		p.SetState(484)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ANTLRv4ParserASSIGN || _la == ANTLRv4ParserPLUS_ASSIGN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF, ANTLRv4ParserSTRING_LITERAL, ANTLRv4ParserDOT, ANTLRv4ParserNOT:
		{
			p.SetState(485)
			p.Atom()
		}

	case ANTLRv4ParserLPAREN:
		{
			p.SetState(486)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IEbnfContext is an interface to support dynamic dispatch.
type IEbnfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	BlockSuffix() IBlockSuffixContext

	// IsEbnfContext differentiates from other interfaces.
	IsEbnfContext()
}

type EbnfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEbnfContext() *EbnfContext {
	var p = new(EbnfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ebnf
	return p
}

func InitEmptyEbnfContext(p *EbnfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ebnf
}

func (*EbnfContext) IsEbnfContext() {}

func NewEbnfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EbnfContext {
	var p = new(EbnfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ebnf

	return p
}

func (s *EbnfContext) GetParser() antlr.Parser { return s.parser }

func (s *EbnfContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *EbnfContext) BlockSuffix() IBlockSuffixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockSuffixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockSuffixContext)
}

func (s *EbnfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EbnfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EbnfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterEbnf(s)
	}
}

func (s *EbnfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitEbnf(s)
	}
}

func (s *EbnfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitEbnf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Ebnf() (localctx IEbnfContext) {
	localctx = NewEbnfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ANTLRv4ParserRULE_ebnf)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Block()
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&24189255811072) != 0 {
		{
			p.SetState(490)
			p.BlockSuffix()
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

// IBlockSuffixContext is an interface to support dynamic dispatch.
type IBlockSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EbnfSuffix() IEbnfSuffixContext

	// IsBlockSuffixContext differentiates from other interfaces.
	IsBlockSuffixContext()
}

type BlockSuffixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockSuffixContext() *BlockSuffixContext {
	var p = new(BlockSuffixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_blockSuffix
	return p
}

func InitEmptyBlockSuffixContext(p *BlockSuffixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_blockSuffix
}

func (*BlockSuffixContext) IsBlockSuffixContext() {}

func NewBlockSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockSuffixContext {
	var p = new(BlockSuffixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_blockSuffix

	return p
}

func (s *BlockSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockSuffixContext) EbnfSuffix() IEbnfSuffixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEbnfSuffixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEbnfSuffixContext)
}

func (s *BlockSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterBlockSuffix(s)
	}
}

func (s *BlockSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitBlockSuffix(s)
	}
}

func (s *BlockSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitBlockSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) BlockSuffix() (localctx IBlockSuffixContext) {
	localctx = NewBlockSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ANTLRv4ParserRULE_blockSuffix)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		p.EbnfSuffix()
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

// IEbnfSuffixContext is an interface to support dynamic dispatch.
type IEbnfSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUESTION() []antlr.TerminalNode
	QUESTION(i int) antlr.TerminalNode
	STAR() antlr.TerminalNode
	PLUS() antlr.TerminalNode

	// IsEbnfSuffixContext differentiates from other interfaces.
	IsEbnfSuffixContext()
}

type EbnfSuffixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEbnfSuffixContext() *EbnfSuffixContext {
	var p = new(EbnfSuffixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ebnfSuffix
	return p
}

func InitEmptyEbnfSuffixContext(p *EbnfSuffixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ebnfSuffix
}

func (*EbnfSuffixContext) IsEbnfSuffixContext() {}

func NewEbnfSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EbnfSuffixContext {
	var p = new(EbnfSuffixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ebnfSuffix

	return p
}

func (s *EbnfSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *EbnfSuffixContext) AllQUESTION() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserQUESTION)
}

func (s *EbnfSuffixContext) QUESTION(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserQUESTION, i)
}

func (s *EbnfSuffixContext) STAR() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSTAR, 0)
}

func (s *EbnfSuffixContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserPLUS, 0)
}

func (s *EbnfSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EbnfSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EbnfSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterEbnfSuffix(s)
	}
}

func (s *EbnfSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitEbnfSuffix(s)
	}
}

func (s *EbnfSuffixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitEbnfSuffix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) EbnfSuffix() (localctx IEbnfSuffixContext) {
	localctx = NewEbnfSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ANTLRv4ParserRULE_ebnfSuffix)
	var _la int

	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserQUESTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(495)
			p.Match(ANTLRv4ParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserQUESTION {
			{
				p.SetState(496)
				p.Match(ANTLRv4ParserQUESTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ANTLRv4ParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(499)
			p.Match(ANTLRv4ParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserQUESTION {
			{
				p.SetState(500)
				p.Match(ANTLRv4ParserQUESTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case ANTLRv4ParserPLUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(503)
			p.Match(ANTLRv4ParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserQUESTION {
			{
				p.SetState(504)
				p.Match(ANTLRv4ParserQUESTION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// ILexerAtomContext is an interface to support dynamic dispatch.
type ILexerAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CharacterRange() ICharacterRangeContext
	TerminalNode() ITerminalNodeContext
	NotSet() INotSetContext
	LEXER_CHAR_SET() antlr.TerminalNode
	DOT() antlr.TerminalNode
	ElementOptions() IElementOptionsContext

	// IsLexerAtomContext differentiates from other interfaces.
	IsLexerAtomContext()
}

type LexerAtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexerAtomContext() *LexerAtomContext {
	var p = new(LexerAtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerAtom
	return p
}

func InitEmptyLexerAtomContext(p *LexerAtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_lexerAtom
}

func (*LexerAtomContext) IsLexerAtomContext() {}

func NewLexerAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LexerAtomContext {
	var p = new(LexerAtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_lexerAtom

	return p
}

func (s *LexerAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *LexerAtomContext) CharacterRange() ICharacterRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterRangeContext)
}

func (s *LexerAtomContext) TerminalNode() ITerminalNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminalNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminalNodeContext)
}

func (s *LexerAtomContext) NotSet() INotSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotSetContext)
}

func (s *LexerAtomContext) LEXER_CHAR_SET() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLEXER_CHAR_SET, 0)
}

func (s *LexerAtomContext) DOT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserDOT, 0)
}

func (s *LexerAtomContext) ElementOptions() IElementOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementOptionsContext)
}

func (s *LexerAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LexerAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LexerAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterLexerAtom(s)
	}
}

func (s *LexerAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitLexerAtom(s)
	}
}

func (s *LexerAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitLexerAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) LexerAtom() (localctx ILexerAtomContext) {
	localctx = NewLexerAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ANTLRv4ParserRULE_lexerAtom)
	var _la int

	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			p.CharacterRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.TerminalNode()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(511)
			p.NotSet()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(512)
			p.Match(ANTLRv4ParserLEXER_CHAR_SET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(513)
			p.Match(ANTLRv4ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(515)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(514)
				p.ElementOptions()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TerminalNode() ITerminalNodeContext
	Ruleref() IRulerefContext
	NotSet() INotSetContext
	DOT() antlr.TerminalNode
	ElementOptions() IElementOptionsContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) TerminalNode() ITerminalNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminalNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminalNodeContext)
}

func (s *AtomContext) Ruleref() IRulerefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRulerefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRulerefContext)
}

func (s *AtomContext) NotSet() INotSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotSetContext)
}

func (s *AtomContext) DOT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserDOT, 0)
}

func (s *AtomContext) ElementOptions() IElementOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementOptionsContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ANTLRv4ParserRULE_atom)
	var _la int

	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)
			p.TerminalNode()
		}

	case ANTLRv4ParserRULE_REF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
			p.Ruleref()
		}

	case ANTLRv4ParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(521)
			p.NotSet()
		}

	case ANTLRv4ParserDOT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(522)
			p.Match(ANTLRv4ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(523)
				p.ElementOptions()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// INotSetContext is an interface to support dynamic dispatch.
type INotSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	SetElement() ISetElementContext
	BlockSet() IBlockSetContext

	// IsNotSetContext differentiates from other interfaces.
	IsNotSetContext()
}

type NotSetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotSetContext() *NotSetContext {
	var p = new(NotSetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_notSet
	return p
}

func InitEmptyNotSetContext(p *NotSetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_notSet
}

func (*NotSetContext) IsNotSetContext() {}

func NewNotSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotSetContext {
	var p = new(NotSetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_notSet

	return p
}

func (s *NotSetContext) GetParser() antlr.Parser { return s.parser }

func (s *NotSetContext) NOT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserNOT, 0)
}

func (s *NotSetContext) SetElement() ISetElementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetElementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetElementContext)
}

func (s *NotSetContext) BlockSet() IBlockSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockSetContext)
}

func (s *NotSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterNotSet(s)
	}
}

func (s *NotSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitNotSet(s)
	}
}

func (s *NotSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitNotSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) NotSet() (localctx INotSetContext) {
	localctx = NewNotSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ANTLRv4ParserRULE_notSet)
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(528)
			p.Match(ANTLRv4ParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(529)
			p.SetElement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(530)
			p.Match(ANTLRv4ParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(531)
			p.BlockSet()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IBlockSetContext is an interface to support dynamic dispatch.
type IBlockSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllSetElement() []ISetElementContext
	SetElement(i int) ISetElementContext
	RPAREN() antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsBlockSetContext differentiates from other interfaces.
	IsBlockSetContext()
}

type BlockSetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockSetContext() *BlockSetContext {
	var p = new(BlockSetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_blockSet
	return p
}

func InitEmptyBlockSetContext(p *BlockSetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_blockSet
}

func (*BlockSetContext) IsBlockSetContext() {}

func NewBlockSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockSetContext {
	var p = new(BlockSetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_blockSet

	return p
}

func (s *BlockSetContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockSetContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLPAREN, 0)
}

func (s *BlockSetContext) AllSetElement() []ISetElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISetElementContext); ok {
			len++
		}
	}

	tst := make([]ISetElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISetElementContext); ok {
			tst[i] = t.(ISetElementContext)
			i++
		}
	}

	return tst
}

func (s *BlockSetContext) SetElement(i int) ISetElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetElementContext); ok {
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

	return t.(ISetElementContext)
}

func (s *BlockSetContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRPAREN, 0)
}

func (s *BlockSetContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserOR)
}

func (s *BlockSetContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserOR, i)
}

func (s *BlockSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterBlockSet(s)
	}
}

func (s *BlockSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitBlockSet(s)
	}
}

func (s *BlockSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitBlockSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) BlockSet() (localctx IBlockSetContext) {
	localctx = NewBlockSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ANTLRv4ParserRULE_blockSet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(ANTLRv4ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(535)
		p.SetElement()
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserOR {
		{
			p.SetState(536)
			p.Match(ANTLRv4ParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(537)
			p.SetElement()
		}

		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(543)
		p.Match(ANTLRv4ParserRPAREN)
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

// ISetElementContext is an interface to support dynamic dispatch.
type ISetElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TOKEN_REF() antlr.TerminalNode
	ElementOptions() IElementOptionsContext
	STRING_LITERAL() antlr.TerminalNode
	CharacterRange() ICharacterRangeContext
	LEXER_CHAR_SET() antlr.TerminalNode

	// IsSetElementContext differentiates from other interfaces.
	IsSetElementContext()
}

type SetElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetElementContext() *SetElementContext {
	var p = new(SetElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_setElement
	return p
}

func InitEmptySetElementContext(p *SetElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_setElement
}

func (*SetElementContext) IsSetElementContext() {}

func NewSetElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetElementContext {
	var p = new(SetElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_setElement

	return p
}

func (s *SetElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetElementContext) TOKEN_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserTOKEN_REF, 0)
}

func (s *SetElementContext) ElementOptions() IElementOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementOptionsContext)
}

func (s *SetElementContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSTRING_LITERAL, 0)
}

func (s *SetElementContext) CharacterRange() ICharacterRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharacterRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharacterRangeContext)
}

func (s *SetElementContext) LEXER_CHAR_SET() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLEXER_CHAR_SET, 0)
}

func (s *SetElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterSetElement(s)
	}
}

func (s *SetElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitSetElement(s)
	}
}

func (s *SetElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitSetElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) SetElement() (localctx ISetElementContext) {
	localctx = NewSetElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ANTLRv4ParserRULE_setElement)
	var _la int

	p.SetState(555)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 70, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(545)
			p.Match(ANTLRv4ParserTOKEN_REF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(546)
				p.ElementOptions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(549)
			p.Match(ANTLRv4ParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(550)
				p.ElementOptions()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
			p.CharacterRange()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(554)
			p.Match(ANTLRv4ParserLEXER_CHAR_SET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AltList() IAltListContext
	RPAREN() antlr.TerminalNode
	COLON() antlr.TerminalNode
	OptionsSpec() IOptionsSpecContext
	AllRuleAction() []IRuleActionContext
	RuleAction(i int) IRuleActionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLPAREN, 0)
}

func (s *BlockContext) AltList() IAltListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAltListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAltListContext)
}

func (s *BlockContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRPAREN, 0)
}

func (s *BlockContext) COLON() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOLON, 0)
}

func (s *BlockContext) OptionsSpec() IOptionsSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionsSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionsSpecContext)
}

func (s *BlockContext) AllRuleAction() []IRuleActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleActionContext); ok {
			len++
		}
	}

	tst := make([]IRuleActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleActionContext); ok {
			tst[i] = t.(IRuleActionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) RuleAction(i int) IRuleActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleActionContext); ok {
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

	return t.(IRuleActionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ANTLRv4ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		p.Match(ANTLRv4ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562950490296320) != 0 {
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserOPTIONS {
			{
				p.SetState(558)
				p.OptionsSpec()
			}

		}
		p.SetState(564)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ANTLRv4ParserAT {
			{
				p.SetState(561)
				p.RuleAction()
			}

			p.SetState(566)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(567)
			p.Match(ANTLRv4ParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(570)
		p.AltList()
	}
	{
		p.SetState(571)
		p.Match(ANTLRv4ParserRPAREN)
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

// IRulerefContext is an interface to support dynamic dispatch.
type IRulerefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RULE_REF() antlr.TerminalNode
	ArgActionBlock() IArgActionBlockContext
	ElementOptions() IElementOptionsContext

	// IsRulerefContext differentiates from other interfaces.
	IsRulerefContext()
}

type RulerefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulerefContext() *RulerefContext {
	var p = new(RulerefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleref
	return p
}

func InitEmptyRulerefContext(p *RulerefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_ruleref
}

func (*RulerefContext) IsRulerefContext() {}

func NewRulerefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulerefContext {
	var p = new(RulerefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_ruleref

	return p
}

func (s *RulerefContext) GetParser() antlr.Parser { return s.parser }

func (s *RulerefContext) RULE_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRULE_REF, 0)
}

func (s *RulerefContext) ArgActionBlock() IArgActionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgActionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgActionBlockContext)
}

func (s *RulerefContext) ElementOptions() IElementOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementOptionsContext)
}

func (s *RulerefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulerefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulerefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterRuleref(s)
	}
}

func (s *RulerefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitRuleref(s)
	}
}

func (s *RulerefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitRuleref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Ruleref() (localctx IRulerefContext) {
	localctx = NewRulerefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ANTLRv4ParserRULE_ruleref)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(ANTLRv4ParserRULE_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserBEGIN_ARGUMENT {
		{
			p.SetState(574)
			p.ArgActionBlock()
		}

	}
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ANTLRv4ParserLT {
		{
			p.SetState(577)
			p.ElementOptions()
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

// ICharacterRangeContext is an interface to support dynamic dispatch.
type ICharacterRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING_LITERAL() []antlr.TerminalNode
	STRING_LITERAL(i int) antlr.TerminalNode
	RANGE() antlr.TerminalNode

	// IsCharacterRangeContext differentiates from other interfaces.
	IsCharacterRangeContext()
}

type CharacterRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterRangeContext() *CharacterRangeContext {
	var p = new(CharacterRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_characterRange
	return p
}

func InitEmptyCharacterRangeContext(p *CharacterRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_characterRange
}

func (*CharacterRangeContext) IsCharacterRangeContext() {}

func NewCharacterRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterRangeContext {
	var p = new(CharacterRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_characterRange

	return p
}

func (s *CharacterRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterRangeContext) AllSTRING_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserSTRING_LITERAL)
}

func (s *CharacterRangeContext) STRING_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSTRING_LITERAL, i)
}

func (s *CharacterRangeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRANGE, 0)
}

func (s *CharacterRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterCharacterRange(s)
	}
}

func (s *CharacterRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitCharacterRange(s)
	}
}

func (s *CharacterRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitCharacterRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) CharacterRange() (localctx ICharacterRangeContext) {
	localctx = NewCharacterRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ANTLRv4ParserRULE_characterRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		p.Match(ANTLRv4ParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(581)
		p.Match(ANTLRv4ParserRANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(582)
		p.Match(ANTLRv4ParserSTRING_LITERAL)
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

// ITerminalNodeContext is an interface to support dynamic dispatch.
type ITerminalNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TOKEN_REF() antlr.TerminalNode
	ElementOptions() IElementOptionsContext
	STRING_LITERAL() antlr.TerminalNode

	// IsTerminalNodeContext differentiates from other interfaces.
	IsTerminalNodeContext()
}

type TerminalNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminalNodeContext() *TerminalNodeContext {
	var p = new(TerminalNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_terminalNode
	return p
}

func InitEmptyTerminalNodeContext(p *TerminalNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_terminalNode
}

func (*TerminalNodeContext) IsTerminalNodeContext() {}

func NewTerminalNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminalNodeContext {
	var p = new(TerminalNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_terminalNode

	return p
}

func (s *TerminalNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminalNodeContext) TOKEN_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserTOKEN_REF, 0)
}

func (s *TerminalNodeContext) ElementOptions() IElementOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementOptionsContext)
}

func (s *TerminalNodeContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSTRING_LITERAL, 0)
}

func (s *TerminalNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminalNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminalNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterTerminalNode(s)
	}
}

func (s *TerminalNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitTerminalNode(s)
	}
}

func (s *TerminalNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitTerminalNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) TerminalNode() (localctx ITerminalNodeContext) {
	localctx = NewTerminalNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ANTLRv4ParserRULE_terminalNode)
	var _la int

	p.SetState(592)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ANTLRv4ParserTOKEN_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(584)
			p.Match(ANTLRv4ParserTOKEN_REF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(586)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(585)
				p.ElementOptions()
			}

		}

	case ANTLRv4ParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(588)
			p.Match(ANTLRv4ParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(590)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ANTLRv4ParserLT {
			{
				p.SetState(589)
				p.ElementOptions()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IElementOptionsContext is an interface to support dynamic dispatch.
type IElementOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode
	AllElementOption() []IElementOptionContext
	ElementOption(i int) IElementOptionContext
	GT() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsElementOptionsContext differentiates from other interfaces.
	IsElementOptionsContext()
}

type ElementOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementOptionsContext() *ElementOptionsContext {
	var p = new(ElementOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_elementOptions
	return p
}

func InitEmptyElementOptionsContext(p *ElementOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_elementOptions
}

func (*ElementOptionsContext) IsElementOptionsContext() {}

func NewElementOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementOptionsContext {
	var p = new(ElementOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_elementOptions

	return p
}

func (s *ElementOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementOptionsContext) LT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserLT, 0)
}

func (s *ElementOptionsContext) AllElementOption() []IElementOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementOptionContext); ok {
			len++
		}
	}

	tst := make([]IElementOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementOptionContext); ok {
			tst[i] = t.(IElementOptionContext)
			i++
		}
	}

	return tst
}

func (s *ElementOptionsContext) ElementOption(i int) IElementOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementOptionContext); ok {
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

	return t.(IElementOptionContext)
}

func (s *ElementOptionsContext) GT() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserGT, 0)
}

func (s *ElementOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ANTLRv4ParserCOMMA)
}

func (s *ElementOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserCOMMA, i)
}

func (s *ElementOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterElementOptions(s)
	}
}

func (s *ElementOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitElementOptions(s)
	}
}

func (s *ElementOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitElementOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ElementOptions() (localctx IElementOptionsContext) {
	localctx = NewElementOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ANTLRv4ParserRULE_elementOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		p.Match(ANTLRv4ParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(595)
		p.ElementOption()
	}
	p.SetState(600)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ANTLRv4ParserCOMMA {
		{
			p.SetState(596)
			p.Match(ANTLRv4ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)
			p.ElementOption()
		}

		p.SetState(602)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(603)
		p.Match(ANTLRv4ParserGT)
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

// IElementOptionContext is an interface to support dynamic dispatch.
type IElementOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	ASSIGN() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsElementOptionContext differentiates from other interfaces.
	IsElementOptionContext()
}

type ElementOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementOptionContext() *ElementOptionContext {
	var p = new(ElementOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_elementOption
	return p
}

func InitEmptyElementOptionContext(p *ElementOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_elementOption
}

func (*ElementOptionContext) IsElementOptionContext() {}

func NewElementOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementOptionContext {
	var p = new(ElementOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_elementOption

	return p
}

func (s *ElementOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementOptionContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ElementOptionContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *ElementOptionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserASSIGN, 0)
}

func (s *ElementOptionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserSTRING_LITERAL, 0)
}

func (s *ElementOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterElementOption(s)
	}
}

func (s *ElementOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitElementOption(s)
	}
}

func (s *ElementOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitElementOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) ElementOption() (localctx IElementOptionContext) {
	localctx = NewElementOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ANTLRv4ParserRULE_elementOption)
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 81, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(605)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(606)
			p.Identifier()
		}
		{
			p.SetState(607)
			p.Match(ANTLRv4ParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(610)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ANTLRv4ParserTOKEN_REF, ANTLRv4ParserRULE_REF:
			{
				p.SetState(608)
				p.Identifier()
			}

		case ANTLRv4ParserSTRING_LITERAL:
			{
				p.SetState(609)
				p.Match(ANTLRv4ParserSTRING_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RULE_REF() antlr.TerminalNode
	TOKEN_REF() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ANTLRv4ParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ANTLRv4ParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) RULE_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserRULE_REF, 0)
}

func (s *IdentifierContext) TOKEN_REF() antlr.TerminalNode {
	return s.GetToken(ANTLRv4ParserTOKEN_REF, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ANTLRv4ParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ANTLRv4ParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ANTLRv4Parser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ANTLRv4ParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(614)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ANTLRv4ParserTOKEN_REF || _la == ANTLRv4ParserRULE_REF) {
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
