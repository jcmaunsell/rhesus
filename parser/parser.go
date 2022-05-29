package parser

import (
	"github.com/jcmaunsell/rhesus/ast"
	"github.com/jcmaunsell/rhesus/lexer"
	"github.com/jcmaunsell/rhesus/token"
)

type Parser interface {
	Parse() ast.Program
}

type parser struct {
	lexer               lexer.Lexer
	curToken, peekToken token.Token
}

func New(lex lexer.Lexer) Parser {
	p := &parser{lexer: lex}

	// Initialize curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *parser) Parse() ast.Program {
	return nil
}

func (p *parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}