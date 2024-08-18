package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLowercaseWords(t *testing.T) {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ToLowercaseWords(test.words)
			assert.Equal(t, test.expected, actual)
		})
	}
}
