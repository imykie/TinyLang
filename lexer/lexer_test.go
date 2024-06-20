package lexer

import (
	"testing"

	"github.com/imykie/tinylang/token"
)

type TestToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {

	tests := map[string]struct {
		input  string
		result []TestToken
	}{
		"test1": {
			input: `=+-(){},;`,
			result: []TestToken{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.MINUS, "-"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		"test2": {
			input: `let five = 5;
					let ten = 10;
					let add = func(x, y) {
						x + y;
					};
					let result = add(five, ten);
			`,
			result: []TestToken{
				{token.LET, "let"},
				{token.IDENTIFIER, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENTIFIER, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENTIFIER, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "func"},
				{token.LPAREN, "("},
				{token.IDENTIFIER, "x"},
				{token.COMMA, ","},
				{token.IDENTIFIER, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENTIFIER, "x"},
				{token.PLUS, "+"},
				{token.IDENTIFIER, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENTIFIER, "result"},
				{token.ASSIGN, "="},
				{token.IDENTIFIER, "add"},
				{token.LPAREN, "("},
				{token.IDENTIFIER, "five"},
				{token.COMMA, ","},
				{token.IDENTIFIER, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		"test3": {
			input: `!-/*5;
   					5 < 10 > 5;
			`,
			result: []TestToken{
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		"test4": {
			input: `if (5 < 10) {
						return true;
					} else {
						return false;
					}
			`,
			result: []TestToken{
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		"test5": {
			input: `10 == 10;
					10 != 9;
					10 <= 9;
					10 >= 9;
					`,
			result: []TestToken{
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.INT, "10"},
				{token.NOT_EQ, "!="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.INT, "10"},
				{token.LT_EQ, "<="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.INT, "10"},
				{token.GT_EQ, ">="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l := New(test.input)
			for i, tt := range test.result {
				tok := l.NextToken()
				if tok.Type != tt.expectedType {
					t.Fatalf("token:[%d] - tokentype wrong. expected=%q, got=%q",
						i, tt.expectedType, tok.Type)
				}
				if tok.Literal != tt.expectedLiteral {
					t.Fatalf("token:[%d] - literal wrong. expected=%q, got=%q",
						i, tt.expectedLiteral, tok.Literal)
				}
			}
		})
	}
}
