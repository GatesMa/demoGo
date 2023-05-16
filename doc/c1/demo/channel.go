package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	//fmt.Println(x, y, x+y)

	c := make(chan int)
	go func() {

		for v := range c {
			fmt.Println("v: ", v)
		}
		close(c)
	}()

	for i := 0; i < 100; i++ {
		c <- i
	}

	//fmt.Println("v: ", <-c)
	//fmt.Println("v: ", <-c)
	//fmt.Println("v: ", <-c)
}
