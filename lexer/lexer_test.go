package lexer

import (
	"jeisaraja/meowcode/token"
	"log"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
    cat five = 5;
    cat ten = 10;
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "cat"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.VAR, "cat"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		log.Println(i+23, tt, tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
