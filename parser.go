package go_kit

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	internal "github.com/lubosgarancovsky/go-kit/internal/rsql"
)

type Parser struct {
	operators []string
}

func operatorPattern(operators []string) string {
	return fmt.Sprintf("^(%s)", strings.Join(operators, "|"))
}

func NewRSQLParser() *Parser {
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

func (p *Parser) Parse(input string) (internal.Node, error) {
	patterns := []internal.Pattern{
		{internal.TokenAnd, `^;`},
		{internal.TokenOr, `^,`},
		{internal.TokenLParen, `^\(`},
		{internal.TokenRParen, `^\)`},
		{internal.TokenOperator, operatorPattern(p.operators)},

		// Quoted strings (double and single quotes)
		{internal.TokenField, `^"(([^"\\]|\\.)*)"`},
		{internal.TokenField, `^'(([^'\\]|\\.)*)'`},

		{internal.TokenField, `^[^;,()\s=><!]+`},
	}

	tokens, err := internal.Tokenize(input, patterns)
	if err != nil {
		return nil, err
	}

	ast, err := internal.BuildAST(tokens)
	if err != nil {
		return nil, err
	}

	return ast, nil
}
