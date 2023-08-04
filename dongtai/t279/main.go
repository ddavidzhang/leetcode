package main

/**
* 完全平方数, 本质跟凑硬币是一样的
 */
func numSquares(n int) int {
	s := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		if i*i > n {
			break
		}
		s = append(s, i*i)
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		r := i
		for k := 0; k < len(s); k++ {
			if i-s[k] < 0 {
				break
			}
			if dp[i-s[k]]+1 < r {
				r = dp[i-s[k]] + 1
			}
		}
		dp[i] = r
	}
	return dp[n]
}

func main() {
	println(numSquares(12))
	println(numSquares(13))
}
