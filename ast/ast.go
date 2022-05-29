package ast

import "github.com/jcmaunsell/rhesus/token"

type Node interface {
	TokenLiteral() token.Literal
}

type Statement Node

type Expr Node

type Program []Statement

type Let struct {
	Name  *Identifier
	Value Expr
}

func (l Let) TokenLiteral() token.Literal {
	return token.LET
}

type Identifier token.Literal

func (i Identifier) TokenLiteral() token.Literal {
	return token.IDENT
}
