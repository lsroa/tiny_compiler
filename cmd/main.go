package main

import "tiny_compiler/pkg/tokenizer"

func main() {
	var input = "add(2 4)"
	tokenizer.Tokenizer(input)
}
