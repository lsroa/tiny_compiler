package parser

import (
	"errors"
	"tiny_compiler/pkg/tokenizer"
)

type NumberLiteral struct {
	Value string
}

type CallExpression struct {
	Name string
	Args []interface{}
}

type Program struct {
	Name string
	Body interface{}
}

type Parser struct {
	tokens []tokenizer.Token
	cursor int
	ast    Program
}

func (p *Parser) Parse(tokens []tokenizer.Token) Program {
	p.tokens = tokens
	p.cursor = 0

	var body []interface{}

	for p.cursor < len(p.tokens) {
		if p.tokens[p.cursor].Type != "string" {
			body = append(body, p.walk())
		}
		p.cursor += 1
	}

	return Program{
		Name: "main",
		Body: body,
	}
}

func (p *Parser) walk() interface{} {
	token := p.tokens[p.cursor]

	if token.Type == "number" {
		p.cursor += 1
		return NumberLiteral{
			Value: token.Value,
		}
	}

	if token.Value == "(" {
		previous_token := p.tokens[p.cursor-1]
		p.cursor += 1
		token = p.tokens[p.cursor]

		var args []interface{}

		for token.Value != ")" {
			args = append(args, p.walk())
			token = p.tokens[p.cursor]
		}

		p.cursor += 1

		return CallExpression{
			Name: previous_token.Value,
			Args: args,
		}
	}

	return errors.New("Invalid token")
}
