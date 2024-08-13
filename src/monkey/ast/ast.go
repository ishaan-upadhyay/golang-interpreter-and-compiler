package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	// Used for debugging and testing.
	TokenLiteral() string
	String() string
}

/**
 * For the purposes of Monkey, a statement is an instruction
 * which does not produce a value.
 */
type Statement interface {
	Node
	// Dummy marker method to determine whether we've used
	// an expression where a statement should be used.
	statementNode()
}

/**
 * For the purposes of Monkey, an expression is an instruction
 * which does produce a value.
 */
type Expression interface {
	Node
	expressionNode()
}

/**
 * Root node of all ASTs produced by the parser.
 * All valid Monkey programs are a series of statements, contained in Program.Statements
 */
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/**
 * Matches the let <identifier> = <expression> rule.
 */
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Matches the return <expression> rule.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Matches the <IDENTIFIER> non-terminal.
//
//	Why is an identifier an expression?
//	For the sake of simplicity - when used elsewhere (not at bind time), identifiers
//	do produce values.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// An expression statement - included to mimic scripting languages
// which allow single lines consisting of expressions.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
