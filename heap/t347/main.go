package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

type fre struct {
	num int
	f   int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*fre)
	for _, n := range nums {
		if f, ok := m[n]; ok {
			f.f++
		} else {
			m[n] = &fre{n, 1}
		}
	}
	var fs []*fre
	for _, v := range m {
		fs = append(fs, v)
	}
	ki := findKthLargest(fs, k)
	var ans []int
	for i := ki; i < len(fs); i++ {
		ans = append(ans, fs[i].num)
	}
	return ans
}

// 快速选择算法, 时间复杂度 O(n)，空间复杂度 O(1)，效率最高
func findKthLargest(nums []*fre, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []*fre, l, r, index int) int {
	q := randomPartition(nums, l, r)
	if q == index {
		return q
	} else if q < index {
		return quickSelect(nums, q+1, r, index)
	}
	return quickSelect(nums, l, q-1, index)
}

func randomPartition(nums []*fre, l, r int) int {
	i := rand.Intn(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []*fre, l, r int) int {
	x := nums[r].f
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j].f <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
