package main

// 暴力解法，时间复杂度 O(n^2)，空间复杂度 O(1)
func subarraySum(nums []int, k int) int {
	sum := 0
	rs := 0
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				rs++
			}
		}
	}
	return rs
}

func main() {
	println(subarraySum([]int{1, 1, 1}, 2))
	println(subarraySum([]int{1, 2, 3}, 3))
	println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
	println(subarraySum([]int{1}, 0))
	println(subarraySum([]int{-1, -1, 1}, 0))
}

// 前缀和解法，时间复杂度 O(n)，空间复杂度 O(n)
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

/**
输入：nums = [1,2,3], k = 3
输出：2
*/
