package tokenizer

import (
	"tiny_compiler/pkg/utils"
)

type Token struct {
	t_type string
	value  string
}

func Tokenizer(input string) []Token {
	len := len(input)
	tokens := []Token{}
	parens := [...]string{"[", "]", "(", ")"}
	// closeOpen := map[string]string{
	// 	"[": "]",
	// 	"(": ")",
	// }
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
			tokens = append(tokens, Token{value: letters, t_type: "string"})
			continue
		}
		if utils.IsNumber(char) {
			numbers := ""
			for utils.IsNumber(char) {
				numbers += char
				i += 1
				char = string(input[i])
			}
			tokens = append(tokens, Token{value: numbers, t_type: "number"})
			continue
		}
		for _, paren := range parens {
			if char == paren {
				tokens = append(tokens, Token{value: char, t_type: "paren"})
			}
		}
		i += 1
	}

	return tokens
}
