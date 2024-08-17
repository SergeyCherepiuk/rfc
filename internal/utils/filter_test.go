package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterRegularWords(t *testing.T) {
	type TestCase struct {
		name     string
		words    []string
		expected []string
	}

	tests := []TestCase{
		{
			name:     "EmptySentence",
			words:    []string{},
			expected: []string{},
		},
		{
			name:     "NoWordsToFilter",
			words:    []string{"Hello", "world"},
			expected: []string{"Hello", "world"},
		},
		{
			name:     "FilterZeroOrOneLetterWords",
			words:    []string{"Hello", "a", "b", "world", ""},
			expected: []string{"Hello", "world"},
		},
		{
			name:     "FilterUppercaseWords",
			words:    []string{"GO", "Hello", "TCP", "UDP", "world", "END"},
			expected: []string{"Hello", "world"},
		},
		{
			name:     "FilterWordsWithNumbers",
			words:    []string{"Hello", "test123", "world", "456"},
			expected: []string{"Hello", "world"},
		},
		{
			name:     "AllWordsFiltered",
			words:    []string{"a", "", "123", "XYZ"},
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := FilterRegularWords(test.words)
			assert.Equal(t, test.expected, actual)
		})
	}
}
