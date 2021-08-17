package utils

import "regexp"

func IsNumber(input string) bool {
	matched, _ := regexp.MatchString("[0-9]", input)
	return matched
}

func IsLetter(input string) bool {
	matched, _ := regexp.MatchString("[a-z]", input)
	return matched
}
