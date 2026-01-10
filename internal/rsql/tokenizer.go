package rsql

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenType string

type Pattern struct {
	Type    TokenType
	Pattern string
}

const (
	TokenField    TokenType = "FIELD"
	TokenOperator TokenType = "OPERATOR"
	TokenAnd      TokenType = "AND"
	TokenOr       TokenType = "OR"
	TokenLParen   TokenType = "LPAREN"
	TokenRParen   TokenType = "RPAREN"
	TokenEOF      TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s", t.Type, t.Value)
}

func Tokenize(input string, patterns []Pattern) ([]Token, error) {
	var tokens []Token
	input = strings.TrimSpace(input)

	for len(input) > 0 {
		matched := false
		for _, p := range patterns {
			re := regexp.MustCompile(p.Pattern)
			loc := re.FindStringIndex(input)

			if loc != nil && loc[0] == 0 {
				match := input[loc[0]:loc[1]]

				tokens = append(tokens, Token{Type: p.Type, Value: match})

				input = input[loc[1]:]
				input = strings.TrimSpace(input)

				matched = true
				break
			}
		}

		if !matched {
			return nil, fmt.Errorf("unexpected token near: %s", input)
		}
	}

	return tokens, nil
}
