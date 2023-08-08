package main

import "fmt"

func decodeString(s string) string {
	var stack []rune
	for _, c := range s {
		if c == ']' {
			var cs []rune
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for t != '[' {
				cs = append([]rune{t}, cs...)
				t = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			t = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n, i := 0, 1
			for t >= '0' && t <= '9' {
				n = int(t-'0')*i + n
				i *= 10
				if len(stack) > 0 {
					t = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			if len(stack) > 0 { // 补回不是数字的那个字符
				stack = append(stack, t)
			}
			for i := 0; i < n; i++ {
				stack = append(stack, cs...)
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}
