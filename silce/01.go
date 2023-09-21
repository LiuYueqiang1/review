package main

import "fmt"

func main() {
	a := make([]int, 2)
	a = append(a, 1, 2)
	fmt.Println("a", a) //0 0 1 2
	b := []int{3, 4, 5}
	copy(a, b)
	fmt.Println("a", a) //a [3 4 5 2]
	fmt.Println("b", b) //b [3 4 5]

	var c = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		c = append(c, fmt.Sprintf("%v", i))
	}
	fmt.Println("c", c, "len", len(c), "cap", cap(c))
	//c [     0 1 2 3 4 5 6 7 8 9] len 15 cap 20
	//cap 两倍扩容=20
}
