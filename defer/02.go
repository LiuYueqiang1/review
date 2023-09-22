package main

import "fmt"

// 1、返回值赋值
// 2、defer语句
// 3、真正的return语句
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 首先x=5，return x=5，x++   返回值为5， x=6

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// x=5， 返回值为x ，x++ =6

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// x=5 ,return y=x=5, x++=6 ,y=5

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// x=5,  return 5 ,在匿名函数中，对x的副本进行递增操作, 而不是地址, return &x=5

func f5() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}

// 对x的地址进行操作，return x=6

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
