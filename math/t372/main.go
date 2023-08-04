package main

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := pow(a, last)
	part2 := pow(superPow(a, b), 10)
	return (part1 * part2) % 1337
}

func pow(a, b int) int {
	a %= 1337
	res := 1
	for i := 0; i < b; i++ {
		res *= a
		res %= 1337
	}
	return res
}
