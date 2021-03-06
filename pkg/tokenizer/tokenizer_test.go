package tokenizer

import (
	"testing"
)

func TestTokenizer(t *testing.T) {

	t.Run("basic function", func(t *testing.T) {
		tokens := Tokenizer("add(2 4)")
		want := []string{"add", "(", "2", "4", ")"}
		for i, token := range tokens {
			if token.Value != want[i] {
				t.Error(token.Value, "is not equal to ", want[i])
			}
		}
	})
}
