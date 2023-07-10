package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	for i, _ := range res {
		res[i] = -1
	}
	var stack []int
	for i := 0; i < l*2; i++ {
		v := nums[i%l]
		for len(stack) > 0 && v > nums[stack[len(stack)-1]%l] {
			res[stack[len(stack)-1]%l] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
