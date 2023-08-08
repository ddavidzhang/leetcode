package main

import "fmt"

/**
* 要点：双指针，每次移动较小的那个指针
* 双指针问题一般还是需要达到 O(n) 的时间复杂度，根据实际问题特性来调整指针移动方向
 */
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ma := 0
	for l < r {
		ma = max(ma, area(l, r, height[l], height[r]))
		hl, hr := height[l], height[r]
		if hl < hr {
			for height[l] <= hl && l < r {
				l++
			}
		} else {
			for height[r] <= hr && l < r {
				r--
			}
		}
	}
	return ma
}

func area(i, j, hi, hj int) int {
	return min(hi, hj) * (j - i)
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
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
