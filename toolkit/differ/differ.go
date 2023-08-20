package differ

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"
	antlrparser "github.com/rebelice/agt/parser/antlr"
)

type Differ struct {
	GrammarFile string
	Target      string

	Tree antlrparser.IGrammarSpecContext
}

// NewDiffer creates a new differ.
func NewDiffer(grammarFile string, target string) *Differ {
	return &Differ{
		GrammarFile: grammarFile,
		Target:      target,
	}
}

func (d *Differ) Run() error {
	if err := d.parseGrammar(); err != nil {
		return err
	}

	grammarRuleListener := &grammarRuleListener{
		target: d.Target,
	}

	antlr.ParseTreeWalkerDefault.Walk(grammarRuleListener, d.Tree)

	return nil
}

type parseErrorListener struct {
	err error
}

// SyntaxError is called when the parser encounters an error.
func (l *parseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if l.err == nil {
		l.err = errors.Errorf("line %d:%d %s", line, column, msg)
	}
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

// ReportAmbiguity reports an ambiguity.
func (*parseErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

// ReportAttemptingFullContext reports an attempting full context.
func (*parseErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

// ReportContextSensitivity reports a context sensitivity.
func (*parseErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func (d *Differ) parseGrammar() error {
	grammar, err := antlr.NewFileStream(d.GrammarFile)
	if err != nil {
		return err
	}

	lexer := antlrparser.NewANTLRv4Lexer(grammar)
	// Maybe this is the the bug in antlr4-go?
	// If we don't set lexer.Virt = lexer, the lexer.Virt will be the antlr.BaseLexer,
	// So the lexer.Virt.Emit() will call the antlr.BaseLexer.Emit() instead of antlrparser.ANTLRv4Lexer.Emit().
	lexer.Virt = lexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := antlrparser.NewANTLRv4Parser(stream)

	lexerErrorListener := &parseErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrorListener)

	parserErrorListener := &parseErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrorListener)

	p.BuildParseTrees = true
	d.Tree = p.GrammarSpec()

	if lexerErrorListener.err != nil {
		return lexerErrorListener.err
	}

	if parserErrorListener.err != nil {
		return parserErrorListener.err
	}

	return nil
}

type grammarRuleListener struct {
	*antlrparser.BaseANTLRv4ParserListener

	target string

	ruleName string
	ruleMap  map[string]bool
}

func (g *grammarRuleListener) EnterParserRuleSpec(ctx *antlrparser.ParserRuleSpecContext) {
	if ctx.RULE_REF().GetText() == g.target {
		g.ruleName = g.target
		fmt.Printf("Found target rule %s\n", g.target)
	}
}

func (g *grammarRuleListener) ExitParserRuleSpec(ctx *antlrparser.ParserRuleSpecContext) {
	g.ruleName = ""
}

func (g *grammarRuleListener) EnterElement(ctx *antlrparser.ElementContext) {
	if g.ruleName == "" {
		return
	}

	// TODO: handle labeledElement
	if ctx.LabeledElement() != nil {
		return
	}

	// TODO: handle atom
	if ctx.Atom() != nil {
		return
	}

	// handle ebnf
	if ctx.Ebnf() != nil {
		return
	}

	// TODO: handle actionBlock
	if ctx.ActionBlock() != nil {
		return
	}
}

func (g *grammarRuleListener) EnterAtom(ctx *antlrparser.AtomContext) {
	if g.ruleName == "" {
		return
	}

	if ctx.Ruleref() != nil {

	}
}
