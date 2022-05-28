package token

type Token struct {
	Type    Type
	Literal Literal
}

type Type string

type Literal string

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	// Identifiers and literals

	IDENT Type = "IDENT"
	INT   Type = "INT"

	// Operators

	ASSIGN   Type = "="
	PLUS     Type = "+"
	MINUS    Type = "-"
	BANG     Type = "!"
	ASTERISK Type = "*"
	SLASH    Type = "/"
	LT       Type = "<"
	GT       Type = ">"

	// Delimiters

	COMMA     Type = ","
	SEMICOLON Type = ";"
	LPAREN    Type = "("
	RPAREN    Type = ")"
	LBRACE    Type = "{"
	RBRACE    Type = "}"

	// Keywords

	FN     Type = "fn"
	LET    Type = "let"
	TRUE   Type = "true"
	FALSE  Type = "false"
	IF     Type = "if"
	ELSE   Type = "else"
	RETURN Type = "return"
)

func Symbol(typ Type) Token {
	return Token{typ, Literal(typ)}
}

func Keyword(typ Type) Token {
	return Token{typ, Literal(typ)}
}

func Identifier(literal Literal) Token {
	return Token{typeOf(literal), literal}
}

func Integer(literal Literal) Token {
	return Token{INT, literal}
}

var (
	Assign   = Symbol(ASSIGN)
	Plus     = Symbol(PLUS)
	Minus    = Symbol(MINUS)
	Bang     = Symbol(BANG)
	Asterisk = Symbol(ASTERISK)
	Slash    = Symbol(SLASH)

	LessThan    = Symbol(LT)
	GreaterThan = Symbol(GT)

	Semicolon  = Symbol(SEMICOLON)
	LeftParen  = Symbol(LPAREN)
	RightParen = Symbol(RPAREN)
	LeftBrace  = Symbol(LBRACE)
	RightBrace = Symbol(RBRACE)
	Comma      = Symbol(COMMA)

	EndOfFile = Token{EOF, ""}

	Function = Keyword(FN)
	Let      = Keyword(LET)
	True     = Keyword(TRUE)
	False    = Keyword(FALSE)
	If       = Keyword(IF)
	Else     = Keyword(ELSE)
	Return   = Keyword(RETURN)
)

func IllegalCharacter(char byte) Token {
	return Token{ILLEGAL, Literal(char)}
}

func FromChar(char byte) Token {
	switch char {
	case '=':
		return Assign
	case '+':
		return Plus
	case '-':
		return Minus
	case '!':
		return Bang
	case '*':
		return Asterisk
	case '/':
		return Slash
	case '<':
		return LessThan
	case '>':
		return GreaterThan
	case ',':
		return Comma
	case ';':
		return Semicolon
	case '(':
		return LeftParen
	case ')':
		return RightParen
	case '{':
		return LeftBrace
	case '}':
		return RightBrace
	case 0:
		return EndOfFile
	default:
		return IllegalCharacter(char)
	}
}

func typeOf(identifier Literal) Type {
	switch Type(identifier) {
	case FN:
		fallthrough
	case LET:
		fallthrough
	case TRUE:
		fallthrough
	case FALSE:
		fallthrough
	case IF:
		fallthrough
	case ELSE:
		fallthrough
	case RETURN:
		return Type(identifier)
	default:
		return IDENT
	}
}
