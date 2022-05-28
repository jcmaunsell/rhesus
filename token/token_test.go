package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromChar(t *testing.T) {
	for _, tt := range []struct {
		name     string
		char     byte
		expected Token
	}{
		{
			"assign",
			'=',
			Token{ASSIGN, "="},
		},
		{
			"plus",
			'+',
			Token{PLUS, "+"},
		},
		{
			"lparen",
			'(',
			Token{LPAREN, "("},
		},
		{
			"rparen",
			')',
			Token{RPAREN, ")"},
		},
		{
			"lbrace",
			'{',
			Token{LBRACE, "{"},
		},
		{
			"rbrace",
			'}',
			Token{RBRACE, "}"},
		},
		{
			"comma",
			',',
			Token{COMMA, ","},
		},
		{
			"semicolon",
			';',
			Token{SEMICOLON, ";"},
		},
		{
			"EOF",
			0,
			Token{EOF, ""},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, FromChar(tt.char))
		})
	}
}
