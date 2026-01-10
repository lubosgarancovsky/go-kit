package rsql

import "fmt"

type ASTBuilder struct {
	tokens []Token
	pos    int
}

func BuildAST(tokens []Token) (Node, error) {
	builder := ASTBuilder{tokens: tokens, pos: 0}
	return builder.parseExpression()
}

func (b *ASTBuilder) parseExpression() (Node, error) {
	var nodes []Node

	for {
		node, err := b.parseTerm()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)

		if b.current().Type == TokenOr {
			b.advance()
		} else {
			break
		}
	}

	if len(nodes) == 1 {
		return nodes[0], nil
	}

	return &LogicalNode{
		Operator: TokenOr,
		Children: nodes,
	}, nil
}

func (b *ASTBuilder) parseTerm() (Node, error) {
	var nodes []Node

	for {
		node, err := b.parseFactor()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)

		if b.current().Type == TokenAnd {
			b.advance()
		} else {
			break
		}
	}

	if len(nodes) == 1 {
		return nodes[0], nil
	}

	return &LogicalNode{
		Operator: TokenAnd,
		Children: nodes,
	}, nil
}

func (b *ASTBuilder) parseFactor() (Node, error) {
	token := b.current()

	if token.Type == TokenLParen {
		b.advance()
		node, err := b.parseExpression()
		if err != nil {
			return nil, err
		}
		if b.current().Type != TokenRParen {
			return nil, fmt.Errorf("Unexpected token. Expected \")\"")
		}
		b.advance()
		return node, nil
	}
	return b.parseComparison()
}

func (b *ASTBuilder) parseComparison() (Node, error) {
	field := b.current()
	if field.Type != TokenField {
		return nil, fmt.Errorf("Unexpected token. Expected field, got %v", field)
	}
	b.advance()

	op := b.current()
	if op.Type != TokenOperator {
		return nil, fmt.Errorf("Unexpected token. Expected operator, got %v", op)
	}
	b.advance()

	val := b.current()
	comparisonValue := []string{stripQuotes(val.Value)}

	if val.Type != TokenField {
		if val.Type == TokenLParen {
			valueArr, err := b.parseArray()
			if err != nil {
				return nil, err
			} else {
				comparisonValue = valueArr
			}
		} else {
			return nil, fmt.Errorf("Unexpected token. Expected value, got %v", val)
		}
	}
	b.advance()

	return &ComparisonNode{
		Field:    field.Value,
		Operator: op.Value,
		Value:    comparisonValue,
	}, nil
}

func (b *ASTBuilder) parseArray() ([]string, error) {
	result := []string{}

	for {
		prevToken := b.current()
		b.advance()
		token := b.current()

		// If array is closed and previos token is VALUE, parsing is done
		if token.Type == TokenRParen {
			if prevToken.Type != TokenField {
				return result, fmt.Errorf("Unexpected token before %v %v", token.Type, token.Value)
			}
			return result, nil
		}

		// If token is VALUE, prevToken must be COMMA or LPAREN
		if token.Type == TokenField {
			if prevToken.Type != TokenOr && prevToken.Type != TokenLParen {
				return result, fmt.Errorf("Unexpected token before %v %v", token.Type, token.Value)
			}

			result = append(result, stripQuotes(token.Value))
			continue
		}

		// If token is COMMA, skip to next token
		if token.Type == TokenOr {
			continue
		}

		// If EOF, array was not closed properly, break the loop
		if token.Type == TokenEOF {
			break
		}

	}
	return result, fmt.Errorf("Syntax error: Array is not closed")
}

func (b *ASTBuilder) current() Token {
	if b.pos >= len(b.tokens) {
		return Token{Type: "EOF", Value: ""}
	}
	return b.tokens[b.pos]
}

func (b *ASTBuilder) advance() {
	b.pos++
}

func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
