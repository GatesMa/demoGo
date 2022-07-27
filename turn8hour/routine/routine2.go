package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 退出当前goroutine
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	// 防止程序退出
	for {
		time.Sleep(1 * time.Second)
	}
}
