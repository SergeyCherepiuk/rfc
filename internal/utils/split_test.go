package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitByAny(t *testing.T) {
	type TestCase struct {
		name       string
		text       string
		separators []rune
		expected   []string
	}

	tests := []TestCase{
		{
			name:       "BasicSplit",
			text:       "Hello World!",
			separators: []rune{' ', '!'},
			expected:   []string{"Hello", "World"},
		},
		{
			name:       "EmptyString",
			text:       "",
			separators: []rune{' ', ','},
			expected:   []string{},
		},
		{
			name:       "NoSeparators",
			text:       "HelloWorld",
			separators: []rune{},
			expected:   []string{"HelloWorld"},
		},
		{
			name:       "AllSeparators",
			text:       "!!!",
			separators: []rune{'!'},
			expected:   []string{},
		},
		{
			name:       "LeadingSeparators",
			text:       "!!Hello",
			separators: []rune{'!'},
			expected:   []string{"Hello"},
		},
		{
			name:       "TrailingSeparators",
			text:       "Hello!!",
			separators: []rune{'!'},
			expected:   []string{"Hello"},
		},
		{
			name:       "AdjacentSeparators",
			text:       "Hello,,World",
			separators: []rune{','},
			expected:   []string{"Hello", "World"},
		},
		{
			name:       "MixedSeparators",
			text:       "Hello, World! How are you?",
			separators: []rune{',', ' ', '!', '?'},
			expected:   []string{"Hello", "World", "How", "are", "you"},
		},
		{
			name:       "SingleCharacterElements",
			text:       "a,b,c",
			separators: []rune{','},
			expected:   []string{"a", "b", "c"},
		},
		{
			name:       "UnicodeCharacters",
			text:       "こんにちは世界",
			separators: []rune("ん世"),
			expected:   []string{"こ", "にちは", "界"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := SplitByAny(tc.text, tc.separators)
			assert.Equal(t, tc.expected, actual, "Failed for case: %s", tc.name)
		})
	}
}
