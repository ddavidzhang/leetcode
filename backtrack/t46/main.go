package main

import "fmt"

/*
* 要点：可以利用 slice 来避免重复元素的问题
 */
func permute(nums []int) [][]int {
	var res [][]int
	var dfs func(nums []int, path []int)
	dfs = func(nums []int, path []int) {
		if len(nums) == 0 {
			res = append(res, path)
			return
		}
		for i := 0; i < len(nums); i++ {
			dfs(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(path, nums[i]))
		}
	}
	dfs(nums, []int{})
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
