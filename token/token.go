package token

type Type string

type Token struct {
	Type    Type
	Literal Literal
}

type Literal string

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	// Identifiers and literals

	IDENT Type = "IDENT"
	INT   Type = "INT"

	// Operators

	ASSIGN Type = "="
	PLUS   Type = "+"

	// Delimiters

	COMMA     Type = ","
	SEMICOLON Type = ";"
	LPAREN    Type = "("
	RPAREN    Type = ")"
	LBRACE    Type = "{"
	RBRACE    Type = "}"

	FUNCTION Type = "function"
	DEFINE   Type = "def"

	// Keywords

	FN  Literal = "fn"
	LET Literal = "let"
)

func Let() Token {
	return Identifier(LET)
}

func Identifier(literal Literal) Token {
	return Token{TypeOf(literal), literal}
}

func Integer(literal Literal) Token {
	return Token{INT, literal}
}

func Assign() Token {
	return Token{ASSIGN, "="}
}

func Plus() Token {
	return Token{PLUS, "+"}
}

func Semicolon() Token {
	return Token{SEMICOLON, ";"}
}

func Function() Token {
	return Identifier(FN)
}

func LeftParen() Token {
	return Token{LPAREN, Literal('(')}
}

func RightParen() Token {
	return Token{RPAREN, Literal(')')}
}

func LeftBrace() Token {
	return Token{LBRACE, Literal('{')}
}

func RightBrace() Token {
	return Token{RBRACE, Literal('}')}
}

func Comma() Token {
	return Token{COMMA, Literal(',')}
}

func EndOfFile() Token {
	return Token{EOF, ""}
}

func IllegalCharacter(char byte) Token {
	return Token{ILLEGAL, Literal(char)}
}

func FromChar(char byte) Token {
	switch char {
	case '=':
		return Assign()
	case '+':
		return Plus()
	case ',':
		return Comma()
	case ';':
		return Semicolon()
	case '(':
		return LeftParen()
	case ')':
		return RightParen()
	case '{':
		return LeftBrace()
	case '}':
		return RightBrace()
	case 0:
		return EndOfFile()
	default:
		return IllegalCharacter(char)
	}
}

func TypeOf(identifier Literal) Type {
	switch identifier {
	case FN:
		return FUNCTION
	case LET:
		return DEFINE
	default:
		return IDENT
	}
}
