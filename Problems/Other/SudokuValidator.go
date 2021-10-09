package Problems

/*

Sudoku Background
Sudoku is a game played on a 9x9 grid. The goal of the game is to fill
all cells of the grid with digits from 1 to 9, so that each column, each row,
and each of the nine 3x3 sub-grids (also known as blocks)
contain all of the digits from 1 to 9.

validSolution([
  [5, 3, 4, 6, 7, 8, 9, 1, 2],
  [6, 7, 2, 1, 9, 5, 3, 4, 8],
  [1, 9, 8, 3, 4, 2, 5, 6, 7],
  [8, 5, 9, 7, 6, 1, 4, 2, 3],
  [4, 2, 6, 8, 5, 3, 7, 9, 1],
  [7, 1, 3, 9, 2, 4, 8, 5, 6],
  [9, 6, 1, 5, 3, 7, 2, 8, 4],
  [2, 8, 7, 4, 1, 9, 6, 3, 5],
  [3, 4, 5, 2, 8, 6, 1, 7, 9]
]); // => true
*/

func SudokuValidator(m [][]int) bool {
	if len(m) == 0 || len(m) != len(m[0]) {
		return false
	}

	// store hash map of all rows, cols
	row_map := make(map[int]map[int]int)
	col_map := make(map[int]map[int]int)

	for cols := 0; cols < len(m); cols++ {
		for rows := 0; rows < len(m); rows++ {
			if row_map[cols] == nil {
				row_map[cols] = make(map[int]int)
			}
			if col_map[rows] == nil {
				col_map[rows] = make(map[int]int)
			}

			row_map[cols][m[cols][rows]]++
			col_map[rows][m[cols][rows]]++

			if col_map[rows][m[cols][rows]] > 1 || row_map[cols][m[cols][rows]] > 1 {
				return false
			}
		}
	}

	return true
}
