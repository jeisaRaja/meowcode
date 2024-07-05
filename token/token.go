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

	IDENT = "IDENT" 
	INT   = "INT"
  
	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

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

func LookUpIdentifier(literal string) TokenType {
	if tokenType, ok := keywords[literal]; ok {
		return tokenType
	}
	return IDENT
}
