package main

/**
* 打家劫舍：自己的方法想的偏复杂了
* ex: [2,7,9,3,1]
* 要点：dp[i] 表示选择 i 可获取的最大值，dp[i] = max(dp[i-2]+nums[i], dp[i-3]+nums[i])
* 选择 i 时，i-1 不能选择，i-2 可以选择，i-3 可以选择，不用考虑 i-4 及以前的选择，因为 i-4 及以前的选择可以合并到 i-2 或 i-3 中
 */
func rob2(nums []int) int {
	dp := make([]int, len(nums)) // dp[i] 表示选择 i 可获取的最大值
	for i := range nums {
		if i == 0 {
			dp[0] = nums[0]
		} else if i == 1 {
			dp[1] = nums[1]
		} else if i == 2 {
			dp[2] = dp[0] + nums[2]
		} else {
			dp[i] = max(dp[i-2]+nums[i], dp[i-3]+nums[i])
		}
	}
	if len(dp) == 0 {
		return 0
	} else if len(dp) == 1 {
		return dp[0]
	}
	return max(dp[len(dp)-1], dp[len(dp)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 其他的解法：dp[i]=max(dp[i−2]+nums[i],dp[i−1])
func rob(nums []int) int {
	dp := make([]int, len(nums)) // dp[i] 表示前 i 个房子可获取的最大值
	for i := range nums {
		if i == 0 {
			dp[0] = nums[0]
		} else if i == 1 {
			dp[1] = max(nums[0], nums[1])
		} else {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
	}
	if len(dp) == 0 {
		return 0
	}
	return dp[len(dp)-1]
}

func main() {
	println(rob([]int{2, 7, 9, 3, 1}))
	println(rob([]int{1, 2, 3, 1}))
}
