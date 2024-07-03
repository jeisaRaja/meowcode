package token

var keywords = map[string]TokenType{
	"litter": FUNCTION,
	"cat":    VAR,
	"hiss":   IF,
	"purr":   ELSE,
	"pounce": RETURN,
	"meow":   TRUE,
	"mrow":   FALSE,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"

	RETURN = "RETURN"
	IF     = "IF"
	ELSE   = "ELSE"

	TRUE  = "TRUE"
	FALSE = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
