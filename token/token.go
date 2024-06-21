package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t Token) String() string {
	switch t.Type {
	case EOF:
		return "EOF"

	default:
		if len(t.Literal) > 10 {
			return fmt.Sprintf("Type: %s, Literal:%.10q...", t.Type, t.Literal)
		}
		return fmt.Sprintf("Type: %s, Literal:%q", t.Type, t.Literal)
	}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MODULO   = "%"

	// Conditionals
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
	WHILE    = "WHILE"
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	DEFAULT  = "DEFAULT"
	CONTINUE = "CONTINUE"
	BREAK    = "BREAK"
)

var keywords = map[string]TokenType{
	"func":     FUNCTION,
	"let":      LET,
	"const":    CONST,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"for":      FOR,
	"while":    WHILE,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"continue": CONTINUE,
	"break":    BREAK,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}
