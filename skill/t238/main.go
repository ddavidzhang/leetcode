package main

import "fmt"

func productExceptSelf(nums []int) []int {
	a := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			a[i] = nums[i]
		} else {
			a[i] = a[i+1] * nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[i] = nums[i]
		} else {
			nums[i] = nums[i-1] * nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		a[i] = 1
		if i-1 >= 0 {
			a[i] *= nums[i-1]
		}
		if i+1 < len(nums) {
			a[i] *= a[i+1]
		}
	}
	return a
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
