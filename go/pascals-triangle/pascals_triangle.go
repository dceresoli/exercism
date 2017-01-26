package pascal

// recursive version
func pascal(row, col int) int {
	if row == 1 || col == 1 || col == row {
		return 1
	}
	return pascal(row-1, col) + pascal(row-1, col-1)
}

func TriangleRow(n int) []int {
	res := make([]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = pascal(n, i)
	}
	return res
}

func Triangle(n int) [][]int {
	res := make([][]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = TriangleRow(i)
	}
	return res
}
