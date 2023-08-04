package main

/**
 * 要点：N 皇后，回溯法，递归的写法及 path 的定义可参考
 */
func solveNQueens(n int) [][]string {
	var res [][]string
	var dfs func(row int, path []int)
	dfs = func(row int, path []int) {
		if row == n {
			res = append(res, convert(path))
			return
		}
		for col := 0; col < n; col++ {
			if isValid(row, col, path) {
				dfs(row+1, append(path, col))
			}
		}
	}
	dfs(0, []int{})
	return res
}

func isValid(row, col int, path []int) bool {
	for i, c := range path {
		if c == col || abs(row-i) == abs(col-c) {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func convert(path []int) []string {
	var res []string
	for _, c := range path {
		var s string
		for i := 0; i < len(path); i++ {
			if i == c {
				s += "Q"
			} else {
				s += "."
			}
		}
		res = append(res, s)
	}
	return res
}
