package tokenizer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {

	t.Run("basic function", func(t *testing.T) {
		tokens := Tokenizer("add(2 4)")
		want := []string{"add", "(", "2", "4", ")"}
		for i, token := range tokens {
			if token.value != want[i] {
				t.Error(token.value, "is not equal to ", want[i])
			}
		}
	})
}
