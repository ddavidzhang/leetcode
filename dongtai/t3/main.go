package main

import "fmt"

/*
* 要点：动态规划的思想，l 及 map 记录前一个字符符合条件的起始位置及字符
 */
func lengthOfLongestSubstring(s string) int {
	l, ans := 0, 0
	m := make(map[byte]bool)
	for r := range s {
		v, ok := m[s[r]]
		for ok && v {
			m[s[l]] = false
			l++
			v, ok = m[s[r]]
		}
		m[s[r]] = true
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("abcabcbb"))
}

// 比上面效率低，没有复用之前结果的 map，但更好理解一些
func lengthOfLongestSubstring2(s string) int {
	l, ans := 0, 0
	for r := range s {
		m := make(map[byte]bool)
		i := r
		for ; i >= l; i-- {
			if m[s[i]] {
				break
			}
			m[s[i]] = true
		}
		l = i + 1
		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}
