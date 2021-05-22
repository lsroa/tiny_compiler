package main

import (
	"bytes"
	"fmt"
)

func main() {
	var input = "Hello, world"
	tokenizer(input)
}

type token struct {
	ttype string
	value string
}

func tokenizer(input string) {
	len = len(input)
	var tokens token[] = token[len]
	// fmt.Println(len)
	// fmt.Println(input)
	// fmt.Println(&input)
	// fmt.Println(*input)

	for i := 0; i < len; i++ {
		fmt.Printf("%c \n", input[i])
		if (Match("\(|\)",input[i])){
			println("Parenthesis %c",input[i])
		}
	}
}
