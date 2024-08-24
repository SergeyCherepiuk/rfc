package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSorted(t *testing.T) {
	type TestCase struct {
		name              string
		slice             []int
		expectedDirection Direction
		expectedResult    bool
	}

	testCases := []TestCase{
		{
			name:              "EmptySlice",
			slice:             []int{},
			expectedDirection: Unknown,
			expectedResult:    true,
		},
		{
			name:              "SortedAscending",
			slice:             []int{1, 4, 8, 9},
			expectedDirection: Ascending,
			expectedResult:    true,
		},
		{
			name:              "SortedDescending",
			slice:             []int{5, 3, -4, -7},
			expectedDirection: Descending,
			expectedResult:    true,
		},
		{
			name:              "SortedUnknown",
			slice:             []int{1, 1, 1, 1, 1, 1},
			expectedDirection: Unknown,
			expectedResult:    true,
		},
		{
			name:              "NotSorted",
			slice:             []int{6, 3, 10},
			expectedDirection: Unknown,
			expectedResult:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualDirection, actualResult := IsSorted(tc.slice)
			assert.Equal(t, tc.expectedDirection, actualDirection)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}
