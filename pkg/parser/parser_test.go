package parser

import (
	"reflect"
	"testing"
	"tiny_compiler/pkg/tokenizer"
)

func TestParser(t *testing.T) {

	t.Run("basic function", func(t *testing.T) {
		var parser Parser
		got := parser.Parse([]tokenizer.Token{
			{
				Type:  "string",
				Value: "add",
			},
			{
				Type:  "paren",
				Value: "(",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "number",
				Value: "4",
			},
			{
				Type:  "paren",
				Value: ")",
			},
		})

		want := Program{
			Name: "main",
			Body: []interface{}{
				CallExpression{
					Name: "add",
					Args: []interface{}{
						NumberLiteral{Value: "2"},
						NumberLiteral{Value: "4"},
					},
				},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Error(got, "Is not equal to ", want)
		}

	})
}
