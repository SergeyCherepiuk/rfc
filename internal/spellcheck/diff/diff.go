package diff

func EditDistance(word1, word2 string) int {
	rows, cols := len(word2)+1, len(word1)+1

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < cols; i++ {
		matrix[0][i] = i
	}

	for i := 0; i < rows; i++ {
		matrix[i][0] = i
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			var (
				deletion     = matrix[i-1][j] + 1
				insertion    = matrix[i][j-1] + 1
				substitution = matrix[i-1][j-1] + boolToInt(word1[j-1] != word2[i-1])
			)
			matrix[i][j] = min(deletion, insertion, substitution)
		}
	}

	return matrix[rows-1][cols-1]
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
