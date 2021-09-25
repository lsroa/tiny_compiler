package tokenizer

import (
	"tiny_compiler/pkg/utils"
)

type Token struct {
	Type  string
	Value string
}

func Tokenizer(input string) []Token {
	len := len(input)
	tokens := []Token{}
	parens := [...]string{"[", "]", "(", ")"}
	i := 0

	for i < len {
		char := string(input[i])
		if char == " " {
			i += 1
			continue
		}

		if utils.IsLetter(char) {
			letters := ""
			for utils.IsLetter(char) {
				letters += char
				i += 1
				char = string(input[i])
			}
			tokens = append(tokens, Token{Value: letters, Type: "string"})
			continue
		}

		if utils.IsNumber(char) {
			numbers := ""
			for utils.IsNumber(char) {
				numbers += char
				i += 1
				char = string(input[i])
			}
			tokens = append(tokens, Token{Value: numbers, Type: "number"})
			continue
		}

		for _, paren := range parens {
			if char == paren {
				tokens = append(tokens, Token{Value: char, Type: "paren"})
			}
		}
		i += 1
	}

	return tokens
}
