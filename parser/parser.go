package parser

import (
	"fmt"
	"github.com/jcmaunsell/rhesus/ast"
	"github.com/jcmaunsell/rhesus/lexer"
	"github.com/jcmaunsell/rhesus/token"
	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"
)

type Parser interface {
	Parse() (ast.Program, error)
}

type parser struct {
	lexer               lexer.Lexer
	curToken, peekToken token.Token
	log                 *logrus.Logger
}

func New(lex lexer.Lexer) Parser {
	p := &parser{lexer: lex, log: logrus.New()}

	// Initialize curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *parser) Parse() (ast.Program, error) {
	var errs error
	program := ast.Program{}
	for p.curToken != token.EndOfFile {
		statement, err := p.parseStatement()
		if err != nil {
			p.log.WithError(err).Error("Could not parse statement.")
			errs = multierr.Append(errs, err)
		} else {
			program = program.Append(statement)
		}
		p.nextToken()
	}
	return program, nil
}

func (p *parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *parser) parseStatement() (ast.Statement, error) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLet()
	default:
		return nil, fmt.Errorf("unsupported token: %v", p.curToken)
	}
}

func (p *parser) parseLet() (ast.Statement, error) {
	statement := ast.Let{}

	if p.curToken.Type != token.LET {
		return nil, fmt.Errorf("expected token type '%v', got '%v'", token.LET, p.curToken)
	}
	p.nextToken()

	statement.Name = ast.Identifier(p.curToken.Literal)
	if p.peekToken.Type != token.ASSIGN {
		return nil, fmt.Errorf("expected token type '%v', got '%v'", token.ASSIGN, p.curToken)
	}
	p.nextToken()

	// TODO(jcmaunsell): Parse expression
	// For now, just skip to the end of the statement
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return statement, nil
}
