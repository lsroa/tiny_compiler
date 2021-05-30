package main

import (
	"fmt"
	"regexp"
)

func main() {
	var input = "add(2 4)"
	tokenizer(input)
}

type token struct {
	t_type string
	value  string
}

func isLetter(input string) bool {
	matched, _ := regexp.MatchString("[a-z]", input)
	return matched
}

func isNumber(input string) bool {
	matched, _ := regexp.MatchString("[0-9]", input)
	return matched
}

func tokenizer(input string) {
	len := len(input)
	tokens := []token{}
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
		if isLetter(char) {
			letters := ""
			for isLetter(char) {
				letters += char
				i += 1
				char = string(input[i])
			}
			tokens = append(tokens, token{value: letters, t_type: "string"})
			continue
		}
		if isNumber(char) {
			numbers := ""
			for isNumber(char) {
				numbers += char
				i += 1
				char = string(input[i])
			}
			tokens = append(tokens, token{value: numbers, t_type: "number"})
			continue
		}
		for _, paren := range parens {
			if char == paren {
				tokens = append(tokens, token{value: char, t_type: "paren"})
			}
		}
		i += 1
	}
	fmt.Println("tokens: ", tokens)
}
