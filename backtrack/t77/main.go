package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n; i++ {
			dfs(i+1, append(path, i))
		}
	}
	dfs(1, []int{})
	return res
}

func main() {
	fmt.Println(combine(4, 2))
}
