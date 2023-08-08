package main

func minWindow(s string, t string) string {
	var m = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	var sm = make(map[byte]int)
	l, r := 0, 0
	var ml = len(s) + 1
	var res string
	for r < len(s) {
		for r < len(s) && !cover(sm, m) { // 判断 m 进行优化计算量
			sm[s[r]]++
			r++
		}
		for l < r {
			if m[s[l]] <= 0 { // 判断 m 进行优化计算量
				sm[s[l]]--
				l++
			} else if cover(sm, m) {
				if r-l < ml {
					ml = r - l
					res = s[l:r]
				}
				sm[s[l]]--
				l++
			} else {
				break
			}
		}
	}
	return res
}

func cover(sm, m map[byte]int) bool {
	for k, v := range m {
		if sm[k] < v {
			return false
		}
	}
	return true
}

func main() {
	println(minWindow("acbbaca", "aba"))
}
