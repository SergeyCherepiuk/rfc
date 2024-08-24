package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowercaseTransformer(t *testing.T) {
	type TestCase struct {
		name     string
		words    []string
		expected []string
	}

	tests := []TestCase{
		{
			name:     "BasicCase",
			words:    []string{"Hello", "DEAR", "world"},
			expected: []string{"hello", "dear", "world"},
		},
		{
			name:     "EmptySentence",
			words:    []string{},
			expected: []string{},
		},
	}

	transformer := NewLowercaseTransformer()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := transformer.Transform(test.words)
			assert.Equal(t, test.expected, actual)
		})
	}
}
