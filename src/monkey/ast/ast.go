package ast

import "monkey/token"

type Node interface {
	// Used for debugging and testing.
	TokenLiteral() string
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

// Matches the return <expression> rule.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

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
