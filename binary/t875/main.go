package main

import "fmt"

/*
*
+ 要点：搜索具有方向性，可以利用二分查找来加速搜索
*/
func minEatingSpeed(piles []int, h int) int {
	min, max := 1, piles[0]
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	l, r := min, max
	for l < r {
		m := l + (r-l)/2
		s := 0
		for _, pile := range piles {
			if pile%m == 0 {
				s += pile / m
			} else {
				s += pile/m + 1
			}
		}
		if s <= h {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
