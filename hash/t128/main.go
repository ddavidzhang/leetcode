package main

/*
*
要点：要求 O(n) 的时间复杂度，每个元素进 map 一次，出一次
*/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		m[n] = true
	}
	ans, t := 0, 0
	for _, n := range nums {
		if m[n] {
			t = 1
			for i := n + 1; m[i]; i++ {
				t++
				m[i] = false
			}
			for i := n - 1; m[i]; i-- {
				t++
				m[i] = false
			}
			if t > ans {
				ans = t
			}
		}
	}
	return ans
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	println(longestConsecutive(nums))
}
