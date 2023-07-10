package main

import "fmt"

func removeDuplicateLetters(s string) string {
	var stack []byte
	var m = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	var visited = make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if visited[s[i]] {
			m[s[i]]--
			continue
		}
		for len(stack) > 0 && s[i] < stack[len(stack)-1] && m[stack[len(stack)-1]] > 0 {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		visited[s[i]] = true
		m[s[i]]--
	}
	return string(stack)
}

func main() {
	fmt.Println('a' - 'a')
}
