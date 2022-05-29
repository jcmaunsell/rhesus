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
	program := ast.Program{}
	for p.curToken != token.EndOfFile {
		if statement := p.parseStatement(); statement != nil {
			program = program.With(statement)
		}
		p.nextToken()
	}
	return program
}

func (p *parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLet()
	default:
		return nil
	}
}

func (p *parser) parseLet() ast.Statement {
	statement := ast.Let{}
	if p.curToken.Type != token.LET {
		return nil
	}
	statement.Name = ast.Identifier(p.curToken.Literal)
	if p.peekToken.Type != token.ASSIGN {
		return nil
	}
	// TODO(jcmaunsell): Parse expression
	// For now, just skip to the end of the statement
	for p.curToken.Type != token.Semicolon {
		p.nextToken()
	}
	return statement
}
