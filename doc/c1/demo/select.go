package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
		// this will not work, cause the main func is done
		time.Sleep(10000000000)
		fmt.Println("sleep end")
	}()

	fibo(quit, c)
	fmt.Println("QUIT!")
}

/**
 *	fibo 斐波那契数列
 */
func fibo(quit chan int, c chan int) {

	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
