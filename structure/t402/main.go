package main

import "fmt"

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	// stack = stack[:len(stack)-k]
	var res []byte
	var i int
	for i < len(stack) && stack[i] == '0' {
		i++
	}
	for i < len(stack) {
		res = append(res, stack[i])
		i++
	}
	if len(res) == 0 {
		return "0"
	}
	return string(res)
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))

}
