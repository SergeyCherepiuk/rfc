package sort

import "cmp"

type Direction int

const (
	Unknown Direction = iota
	Ascending
	Descending
)

func IsSorted[T cmp.Ordered](s []T) (Direction, bool) {
	prev := Unknown

	for i := 0; i < len(s)-1; i++ {
		switch prev {
		case Ascending:
			if s[i] > s[i+1] {
				return Unknown, false
			}

		case Descending:
			if s[i] < s[i+1] {
				return Unknown, false
			}

		case Unknown:
			if s[i] < s[i+1] {
				prev = Ascending
			} else if s[i] > s[i+1] {
				prev = Descending
			}
		}
	}

	return prev, true
}
