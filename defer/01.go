package main

import "fmt"

func main() {
	const (
		a = 1 << iota
		b
		c
	)
	const d = iota
	fmt.Println(a, b, c, d)
}

// 输出 1 2 4 0
