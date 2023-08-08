package main

import "fmt"

func rotate(nums []int, k int) {
	l, r, n := 0, k, 0
	t := nums[l]
	for n <= len(nums) {
		tt := nums[r]
		nums[r] = t
		n++
		t = tt
		if r == l && n < len(nums) {
			l = (l + 1) % len(nums)
			r = (l + k) % len(nums)
			t = nums[l]
		} else {
			r = (r + k) % len(nums)
		}
	}
}

func main() {
	a := []int{-1, -100, 3, 99}
	rotate(a, 2)
	fmt.Println(a)

}
