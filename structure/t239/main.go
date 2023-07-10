package t239

/*
*
要点：slice 模拟队列，单调队列, 保证队列中的元素是单调递减的，每次滑动窗口，队列中的第一个元素就是最大值
go的slice可方便的调整长度，不需要像java那样维护一个队列的头部和尾部，使思路更清晰
时间复杂度：O(n)，每个元素最多进队列一次，出队列一次
*/
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var q []int
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			res = append(res, nums[q[0]])
			if q[0] == i-k+1 {
				q = q[1:]
			}
		}
	}
	return res
}
