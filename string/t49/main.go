package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		var a [26]int
		for _, b := range s {
			a[b-'a']++
		}
		m[a] = append(m[a], s)
	}
	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(-101 % 10)
}
