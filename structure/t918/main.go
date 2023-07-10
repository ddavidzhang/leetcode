package main

import "fmt"

/*
*
要点：环形可用总和-最小来解决
*/
func maxSubarraySumCircular(nums []int) int {
	var maxSum = nums[0]
	var minSum = nums[0]
	var sum = nums[0]
	var maxSub = nums[0]
	var minSub = nums[0]
	for i := 1; i < len(nums); i++ {
		maxSum = max(maxSum+nums[i], nums[i])
		maxSub = max(maxSub, maxSum)
		minSum = min(minSum+nums[i], nums[i])
		minSub = min(minSub, minSum)
		sum += nums[i]
	}
	if sum == minSub {
		return maxSub
	}
	return max(maxSub, sum-minSub)
}

func max(k int, i int) int {
	if k > i {
		return k
	}
	return i
}

func min(k int, i int) int {
	if k < i {
		return k
	}
	return i
}

func main() {
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
}
