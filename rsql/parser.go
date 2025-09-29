package rsql

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Pattern struct {
	Type    TokenType
	Pattern string
}

type Parser struct {
	operators []string
}

func operatorPattern(operators []string) string {
	return fmt.Sprintf("^(%s)", strings.Join(operators, "|"))
}

func New() *Parser {
	return &Parser{
		operators: []string{"==", "!=", ">=", "<=", ">", "<", "=lt=", "=le=", "=gt=", "=ge=", "=in=", "=out=", "=like="},
	}
}

func (p *Parser) RegisterOperator(op string) {
	p.operators = append(p.operators, regexp.QuoteMeta(op))
	sort.Slice(p.operators, func(i, j int) bool {
		return len(p.operators[i]) > len(p.operators[j])
	})
}

func (p *Parser) Parse(input string) (Node, error) {
	patterns := []Pattern{
		{TokenAnd, `^;`},
		{TokenOr, `^,`},
		{TokenLParen, `^\(`},
		{TokenRParen, `^\)`},
		{TokenOperator, operatorPattern(p.operators)},

		// Quoted strings (double and single quotes)
		{TokenField, `^"(([^"\\]|\\.)*)"`},
		{TokenField, `^'(([^'\\]|\\.)*)'`},

		{TokenField, `^[^;,()\s=><!]+`},
	}

	tokens, err := Tokenize(input, patterns)
	if err != nil {
		return nil, err
	}

	ast, err := BuildAST(tokens)
	if err != nil {
		return nil, err
	}

	return ast, nil
}
