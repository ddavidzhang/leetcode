package main

import "fmt"

/**
双指针的解法：a、b两个数组分别记录从左到右和从右到左的最大值
*/

func trap(height []int) int {
	l := len(height)
	a := make([]int, l)
	b := make([]int, l)
	a[0] = 0
	for i := 1; i < l; i++ {
		if height[i-1] > a[i-1] {
			a[i] = height[i-1]
		} else {
			a[i] = a[i-1]
		}
	}
	b[l-1] = 0
	for i := l - 2; i >= 0; i-- {
		if height[i+1] > b[i+1] {
			b[i] = height[i+1]
		} else {
			b[i] = b[i+1]
		}
	}
	sum := 0
	for i := 0; i < l; i++ {
		s := 0
		if a[i] > height[i] && b[i] > height[i] {
			if a[i] < b[i] {
				s = a[i] - height[i]
			} else {
				s = b[i] - height[i]
			}
		}
		if s > 0 {
			sum += s
		}
	}
	return sum
}

// 单调栈的解法
func trap2(height []int) int {
	var stack []int
	sum := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := i - stack[len(stack)-1] - 1
			bounded_height := min(height[i], height[stack[len(stack)-1]]) - height[top]
			sum += distance * bounded_height
		}
		stack = append(stack, i)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
