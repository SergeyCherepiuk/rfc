package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditDistance(t *testing.T) {
	type TestCase struct {
		name     string
		word1    string
		word2    string
		expected int
	}

	testCases := []TestCase{
		{"BothStringsEmpty", "", "", 0},
		{"FirstStringEmpty", "", "a", 1},
		{"SecondStringEmpty", "a", "", 1},
		{"IdenticalStrings", "abc", "abc", 0},
		{"InsertionRequired", "abc", "abcd", 1},
		{"DeletionRequired", "abcd", "abc", 1},
		{"SubstitutionRequired", "abc", "azc", 1},
		{"SubstitutionAndInsertion", "kitten", "sitting", 3},
		{"EmptyVsNonEmptyString", "", "abcdef", 6},
		{"SpecialCharacters", "ab!@#$", "ab#$", 2},
		{"RepeatedCharsNeedInsertion", "aa", "aaaa", 2},
		{"RepeatedCharsNeedDeletion", "aaaa", "aa", 2},
		{"RepeatedCharsNeedSubstitution", "aaaa", "bbbb", 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := EditDistance(tc.word1, tc.word2)
			actualReversed := EditDistance(tc.word2, tc.word1)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, actual, actualReversed)
		})
	}
}
