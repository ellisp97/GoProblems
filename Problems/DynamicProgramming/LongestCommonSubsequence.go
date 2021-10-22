package Problems

func LongestCommonSubsequence(text1, text2 string) int {
	grid := make([][]int, len(text2)+1)
	for i := range grid {
		grid[i] = make([]int, len(text1)+1)
	}

	for col := len(text2) - 1; col >= 0; col-- {
		for row := len(text1) - 1; row >= 0; row-- {
			// 2 Cases
			// 1. Either the letters are the same so take result of diag bottom right + 1
			// 2. letters do not match so take the maximum result from bottom or right
			if text1[row] == text2[col] {
				grid[row][col] = 1 + grid[row+1][col+1]
			} else {
				grid[row][col] = max(grid[row+1][col], grid[row][col+1])
			}
		}
	}
	return grid[0][0]
}

func LongestCommonSubsequenceLessSpace(text1, text2 string) int {

	if len(text2) < len(text1) {
		text1, text2 = text2, text1
	}

	previous := make([]int, len(text2)+1)

	for col := len(text2) - 1; col >= 0; col-- {
		current := make([]int, len(text1)+1)
		for row := len(text1) - 1; row >= 0; row-- {
			// 2 Cases
			// 1. Either the letters are the same so take result of diag bottom right + 1
			// 2. letters do not match so take the maximum result from bottom or right
			if text1[row] == text2[col] {
				current[row] = 1 + previous[row+1]
			} else {
				current[row] = max(previous[row], current[row+1])
			}
		}
		previous = current
	}
	return previous[0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
