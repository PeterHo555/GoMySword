package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || matrix == nil {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	curRow := 0
	curCol := col - 1
	for curRow < row && curCol >= 0 {
		if matrix[curRow][curCol] == target {
			return true
		} else if matrix[curRow][curCol] < target {
			curRow++
		} else {
			curCol--
		}
	}
	return false
}