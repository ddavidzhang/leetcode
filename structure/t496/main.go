package t496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	m := make(map[int]int)
	for _, v := range nums2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	var res []int
	for _, v := range nums1 {
		if v, ok := m[v]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return res
}
