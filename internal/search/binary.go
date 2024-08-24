package search

// NOTE: sc: words slice must contain strings sorted in ascending order.
func Binary(words []string, target string) (int, bool) {
	low, high := 0, len(words)-1

	for low <= high {
		mid := (low + high) / 2
		word := words[mid]

		if target < word {
			high = mid - 1
		} else if target > word {
			low = mid + 1
		} else {
			return mid, true
		}
	}

	return -1, false
}
