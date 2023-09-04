package main

import "fmt"

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// 接收完成最后一个值后返回false
	close(ch)
	// 从一个关闭的通道c接收到最后一个值之后，任何从c接收的值都会成功而不被阻塞，为通道元素返回零值。
	//形式x, ok:= <-c也会将ok设置为false
	f2(ch)
}
