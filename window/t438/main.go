package main

import "fmt"

/*
* 要点：滑动窗口遍历所有可能的子串
 */
func findAnagrams(s string, p string) []int {
	var rs []int
	var pcs [26]int
	for _, c := range p {
		pcs[c-'a']++
	}
	var wcs [26]int
	for l, r := 0, len(p)-1; r < len(s); l, r = l+1, r+1 {
		if l == 0 {
			for i := l; i <= r; i++ {
				wcs[s[i]-'a']++
			}
		} else {
			wcs[s[l-1]-'a']--
			wcs[s[r]-'a']++
		}
		if wcs == pcs {
			rs = append(rs, l)
		}
	}
	return rs
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
