package token

/*
 * Defined here as a string for easier debugging.
 * More typically, an int or byte might be used,
 * which would support >= 255 tokens with better performance.
 */
type TokenType string

type Token struct {
	Type   TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	/*
	 * Identifiers and literals.
	 * IDENT covers symbols like add, foobar, x, y, ...
	 * INT covers integers
	 */
	IDENT = "IDENT" 
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET		 = "LET"
)