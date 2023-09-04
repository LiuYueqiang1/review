package main

import "fmt"

// 单向通道 向通道中放1-10中2的倍数 获取后相加
//<- chan int // 只接收通道，只能接收不能发送
//chan <- int // 只发送通道，只能发送不能接收

func Producer() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 通道关闭后不能存放，只能获取
	}()
	return ch
}

// 从通道中拿值
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res)
}
