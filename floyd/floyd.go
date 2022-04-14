package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	result := make([][]int, rows)

	current := 0
	for i := 0; i < rows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			current++
			result[i][j] = current
		}
	}

	return result
}
