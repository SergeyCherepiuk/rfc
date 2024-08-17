package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitWithWhiteList(t *testing.T) {
	type TestCase struct {
		name      string
		text      string
		whitelist []rune
		expected  []string
	}

	tests := []TestCase{
		{
			name:      "LowercaseLettersWhitelist",
			text:      "ABCdef!ghi123",
			whitelist: []rune("abcdefghijklmnopqrstuvwxyz"),
			expected:  []string{"def", "ghi"},
		},
		{
			name:      "EmptyText",
			text:      "",
			whitelist: []rune("abc"),
			expected:  []string{},
		},
		{
			name:      "EmptyWhitelist",
			text:      "Hello World!",
			whitelist: []rune(""),
			expected:  []string{},
		},
		{
			name:      "OnlyWhitelistCharacters",
			text:      "abcabc",
			whitelist: []rune("abc"),
			expected:  []string{"abcabc"},
		},
		{
			name:      "NoWhitelistCharacters",
			text:      "!!!",
			whitelist: []rune("abc"),
			expected:  []string{},
		},
		{
			name:      "LeadingNonWhitelistCharacters",
			text:      "!!!hello",
			whitelist: []rune("abcdefghijklmnopqrstuvwxyz"),
			expected:  []string{"hello"},
		},
		{
			name:      "TrailingNonWhitelistCharacters",
			text:      "hello!!!",
			whitelist: []rune("abcdefghijklmnopqrstuvwxyz"),
			expected:  []string{"hello"},
		},
		{
			name:      "AdjacentNonWhitelistCharacters",
			text:      "hello,,world",
			whitelist: []rune("abcdefghijklmnopqrstuvwxyz"),
			expected:  []string{"hello", "world"},
		},
		{
			name:      "SingleCharacterElements",
			text:      "a,b,c",
			whitelist: []rune("abc"),
			expected:  []string{"a", "b", "c"},
		},
		{
			name:      "UnicodeCharacters",
			text:      "こんにちは世界",
			whitelist: []rune("こにちは世界"),
			expected:  []string{"こ", "にちは世界"},
		},
		{
			name:      "MixedASCIIAndUnicode",
			text:      "Hello世界123",
			whitelist: []rune("elo世界123"),
			expected:  []string{"ello世界123"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := SplitWithWhiteList(tc.text, tc.whitelist)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
