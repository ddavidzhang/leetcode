package main

// 不用递归，用循环扩展的思想
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	ans := m[digits[0]]
	for i := 1; i < len(digits); i++ {
		var t []string
		for _, v := range ans {
			for _, v2 := range m[digits[i]] {
				t = append(t, v+v2)
			}
		}
		ans = t
	}
	return ans
}
