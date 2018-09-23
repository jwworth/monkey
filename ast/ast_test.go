package ast

import (
	"monkey/token"
	"testing"
)

// Construct the AST and test it by hand
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "hundredDays"},
					Value: "hundredDays",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "codeChallenge"},
					Value: "codeChallenge",
				},
			},
		},
	}

	if program.String() != "let hundredDays = codeChallenge;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
