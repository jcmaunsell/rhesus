package lexer

import (
	"github.com/jcmaunsell/rhesus/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			"symbols",
			`=+(){},;`,
			[]token.Token{
				token.Assign,
				token.Plus,
				token.LeftParen,
				token.RightParen,
				token.LeftBrace,
				token.RightBrace,
				token.Comma,
				token.Semicolon,
				token.EndOfFile,
			},
		},
		{
			"real code",
			`let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);`,
			[]token.Token{
				token.Let,
				token.Identifier("five"),
				token.Assign,
				token.Integer("5"),
				token.Semicolon,

				token.Let,
				token.Identifier("ten"),
				token.Assign,
				token.Integer("10"),
				token.Semicolon,

				// let add = fn(x, y)
				token.Let,
				token.Identifier("add"),
				token.Assign,
				token.Function,
				token.LeftParen,
				token.Identifier("x"),
				token.Comma,
				token.Identifier("y"),
				token.RightParen,

				// { x + y; };
				token.LeftBrace,
				token.Identifier("x"),
				token.Plus,
				token.Identifier("y"),
				token.Semicolon,
				token.RightBrace,
				token.Semicolon,

				// let result = add(five, ten);
				token.Let,
				token.Identifier("result"),
				token.Assign,
				token.Identifier("add"),
				token.LeftParen,
				token.Identifier("five"),
				token.Comma,
				token.Identifier("ten"),
				token.RightParen,
				token.Semicolon,

				token.EndOfFile,
			},
		},
		{
			"extended token set",
			`!-/*5;
	5 < 10 > 5;
`,
			[]token.Token{
				token.Bang,
				token.Minus,
				token.Slash,
				token.Asterisk,
				token.Integer("5"),
				token.Semicolon,
				token.Integer("5"),
				token.LessThan,
				token.Integer("10"),
				token.GreaterThan,
				token.Integer("5"),
				token.Semicolon,
			},
		},
		{
			"booleans",
			`if (5 < 10) {
	return true;
} else {
	return false;
}`,
			[]token.Token{
				token.If,
				token.LeftParen,
				token.Integer("5"),
				token.LessThan,
				token.Integer("10"),
				token.RightParen,
				token.LeftBrace,
				token.Return,
				token.True,
				token.Semicolon,
				token.RightBrace,
				token.Else,
				token.LeftBrace,
				token.Return,
				token.False,
				token.Semicolon,
				token.RightBrace,
			},
		},
		{
			"equality",
			`10 == 10;
10 != 9;`,
			[]token.Token{
				token.Integer("10"),
				token.EqualTo,
				token.Integer("10"),
				token.Semicolon,
				token.Integer("10"),
				token.NotEqualTo,
				token.Integer("9"),
				token.Semicolon,
			},
		},
	} {
		l := New(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			for _, expected := range tt.expected {
				assert.Equal(t, expected, l.NextToken())
			}
		})
	}
}
