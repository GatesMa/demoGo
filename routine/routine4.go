package main

import (
	"fmt"
	"time"
)

func main() {
	// 带有缓冲的channel
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		fmt.Println("func")
		time.Sleep(2 * time.Second)
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送的元素 =", i, "len(c) = ", len(c), " cap(c) = ", cap((c)))
		}
	}()

	fmt.Println("main")

	time.Sleep(1 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println("try get i, ", i)
		num := <-c // 从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}
	fmt.Println("main 结束")
}
