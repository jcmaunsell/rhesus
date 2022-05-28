package lexer

import (
	"github.com/jcmaunsell/rhesus/token"
)

type Lexer interface {
	NextToken() token.Token
}

type lexer struct {
	input string
	cur   int
	next  int
	char  byte
}

func New(input string) Lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) NextToken() token.Token {
	l.skipWhitespace()
	if isLetter(l.char) {
		return token.Identifier(l.readIdentifier())
	} else if isDigit(l.char) {
		return token.Integer(l.readNumber())
	} else {
		tok := token.FromChar(l.char)
		l.readChar()
		return tok
	}
}

func (l *lexer) readIdentifier() token.Literal {
	start := l.cur
	for isLetter(l.char) {
		l.readChar()
	}
	end := l.cur
	return token.Literal(l.input[start:end])
}

func (l *lexer) readNumber() token.Literal {
	start := l.cur
	for isDigit(l.char) {
		l.readChar()
	}
	end := l.cur
	return token.Literal(l.input[start:end])
}

func (l *lexer) readChar() {
	if l.next >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.next]
	}
	l.cur = l.next
	l.next += 1
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *lexer) skipWhitespace() {
	for {
		switch l.char {
		case ' ':
			fallthrough
		case '\t':
			fallthrough
		case '\r':
			fallthrough
		case '\n':
			l.readChar()
		default:
			return
		}
	}
}
