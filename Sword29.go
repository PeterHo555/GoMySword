package main

func spiralOrder(matrix [][]int) []int {
	//if len(matrix) == 0 || len(matrix[0]) == 0 {
	//	return []int{}
	//}
	//d := [][]int{{0, 1},{1, 0},{0, -1},{-1, 0}}
	//rows, cols := len(matrix), len(matrix[0])
	//res := make([]int, rows * cols)
	//vis := make([][]bool, rows)
	//for i := 0; i < rows; i++ {
	//	vis[i] = make([]bool, cols)
	//}
	//r, c, dir := 0, 0, 0
	//for i := 0; i < rows * cols; i++ {
	//	res[i] = matrix[r][c]
	//	vis[r][c] = true
	//	nextR, nextC := r + d[dir][0], c + d[dir][1]
	//	if nextR < 0 || nextR >= rows || nextC < 0 || nextC >= cols || vis[nextR][nextC] {
	//		dir = (dir + 1) % 4
	//	}
	//	r, c = r + d[dir][0], c + d[dir][1]
	//}
	//return res

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	vis := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		vis[i] = make([]bool, columns)
	}
	var (
		total = rows * columns
		res = make([]int, total)
		r, c = 0, 0
		directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		direIndex = 0
	)

	for i := 0; i < total; i++ {
		res[i] = matrix[r][c]
		vis[r][c] = true
		nextRow, nextColumn := r + directions[direIndex][0], c + directions[direIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || vis[nextRow][nextColumn] {
			direIndex = (direIndex + 1) % 4
		}
		r += directions[direIndex][0]
		c += directions[direIndex][1]
	}
	return res
}
