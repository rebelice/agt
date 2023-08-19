package antlr

import (
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

var (
	_ antlr.Lexer = (*LexerAdaptor)(nil)
)

const (
	PREQUEL_CONSTRUCT = -10
	OPTIONS_CONSTRUCT = -11
)

type LexerAdaptor struct {
	*antlr.BaseLexer

	currentRuleType int
}

func (l *LexerAdaptor) handleBeginArgument() {
	if l.inLexerRule() {
		l.PushMode(ANTLRv4LexerLexerCharSet)
		l.More()
	} else {
		l.PushMode(ANTLRv4LexerArgument)
	}
}

func (l *LexerAdaptor) handleEndArgument() {
	l.PopMode()
	if l.ModeStackLength() > 0 {
		l.SetType(ANTLRv4LexerARGUMENT_CONTENT)
	}
}

func (l *LexerAdaptor) handleEndAction() {
	oldMode := l.GetMode()
	newMode := l.PopMode()
	isActionWithinAction := l.ModeStackLength() > 0 && newMode == ANTLRv4LexerTargetLanguageAction && oldMode == newMode
	if isActionWithinAction {
		l.SetType(ANTLRv4LexerACTION_CONTENT)
	}
}

func (l *LexerAdaptor) Emit() antlr.Token {
	if (l.GetType() == ANTLRv4LexerOPTIONS || l.GetType() == ANTLRv4LexerTOKENS || l.GetType() == ANTLRv4LexerCHANNELS) && l.currentRuleType == antlr.TokenInvalidType {
		l.currentRuleType = PREQUEL_CONSTRUCT
	} else if l.GetType() == ANTLRv4LexerOPTIONS && l.currentRuleType == ANTLRv4LexerTOKEN_REF {
		l.currentRuleType = OPTIONS_CONSTRUCT
	} else if l.GetType() == ANTLRv4LexerRBRACE && l.currentRuleType == PREQUEL_CONSTRUCT {
		l.currentRuleType = antlr.TokenInvalidType
	} else if l.GetType() == ANTLRv4LexerRBRACE && l.currentRuleType == OPTIONS_CONSTRUCT {
		l.currentRuleType = ANTLRv4LexerTOKEN_REF
	} else if l.GetType() == ANTLRv4LexerAT && l.currentRuleType == antlr.TokenInvalidType {
		l.currentRuleType = ANTLRv4LexerAT
	} else if l.GetType() == ANTLRv4LexerSEMI && l.currentRuleType == OPTIONS_CONSTRUCT {
		// ';' in options { .... }. Don't change anything.
	} else if l.GetType() == ANTLRv4LexerEND_ACTION && l.currentRuleType == ANTLRv4LexerAT {
		l.currentRuleType = antlr.TokenInvalidType
	} else if l.GetType() == ANTLRv4LexerID {
		firstChar := l.GetInputStream().GetText(l.TokenStartCharIndex, l.TokenStartCharIndex)[0]
		if unicode.IsUpper(rune(firstChar)) {
			l.SetType(ANTLRv4LexerTOKEN_REF)
		} else {
			l.SetType(ANTLRv4LexerRULE_REF)
		}

		if l.currentRuleType == antlr.TokenInvalidType {
			l.currentRuleType = l.GetType()
		}
	} else if l.GetType() == ANTLRv4LexerSEMI {
		l.currentRuleType = antlr.TokenInvalidType
	}
	return l.BaseLexer.Emit()
}

func (l *LexerAdaptor) inLexerRule() bool {
	return l.currentRuleType == ANTLRv4LexerTOKEN_REF
}
