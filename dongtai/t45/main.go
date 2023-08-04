package main

import "fmt"

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < l {
				if dp[i+j] == 0 || dp[i+j] > dp[i]+1 {
					dp[i+j] = dp[i] + 1
				}
			}
		}
	}
	return dp[l-1]
}
func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
