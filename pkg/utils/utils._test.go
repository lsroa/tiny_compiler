package utils

import "testing"

func TestUtils(t *testing.T) {
	t.Run("IsNumber", func(t *testing.T) {
		input := "23"
		got := IsNumber(input)
		if got != true {
			t.Error(input, "is a number")
		}
	})

	t.Run("IsNumber", func(t *testing.T) {
		input := "23.5"
		got := IsNumber(input)
		if got != true {
			t.Error(input, "is a number")
		}
	})

	t.Run("IsNumber", func(t *testing.T) {
		input := "23.5.6"
		got := IsNumber(input)
		if got != false {
			t.Error(input, "is not a number")
		}
	})
}
