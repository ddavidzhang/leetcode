package main

import (
	"container/heap"
	"math/rand"
)

/**
最大堆的解法, 时间复杂度 O(nlogn)，空间复杂度 O(n)
*/

type IntHeap []int

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func findKthLargest3(nums []int, k int) int {
	h := &IntHeap{}
	for _, v := range nums {
		h.Push(v)
	}
	heap.Init(h)
	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

func main() {
	println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

type IntHeap2 []int

func (h *IntHeap2) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap2) Len() int {
	return len(*h)
}

func (h IntHeap2) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func findKthLargest2(nums []int, k int) int {
	h := &IntHeap2{}
	for i := 0; i < k; i++ {
		h.Push(nums[i])
	}
	heap.Init(h)
	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return heap.Pop(h).(int)
}

// 快速选择算法, 时间复杂度 O(n)，空间复杂度 O(1)，效率最高
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, index int) int {
	q := randomPartition(nums, l, r)
	if q == index {
		return nums[q]
	} else if q < index {
		return quickSelect(nums, q+1, r, index)
	}
	return quickSelect(nums, l, q-1, index)
}

func randomPartition(nums []int, l, r int) int {
	i := rand.Intn(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partition(nums, l, r)
}

func partition(nums []int, l, r int) int {
	x := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
