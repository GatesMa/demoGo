package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1, ch2, c3 chan int
	fmt.Println("c3 ", c3)
	c3 = make(chan int)
	fmt.Println("c3 ", c3)
	var i1, i2 int

	go func() {
		c3 <- 1
	}()

	time.Sleep(time.Second * 2)

	select {
	case i1 = <-ch1:
		fmt.Println("receive ", i1, " from ch1")
	case ch2 <- i2:
		fmt.Println("send ", i2, " to ch2")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}
