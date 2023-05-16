package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("=======================4.1======================")
	var x = 11
	if x > 10 {
		println("x > 10")
	} else {
		println("x <= 10")
	}

	fmt.Println("=======================4.2======================")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("1+2+3+...+9=", sum)

	var i int64
	fmt.Println("i = ", i)
	fmt.Println("=======================4.3======================")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println("=======================4.4======================")
	str := "gates-马龙"
	for k, v := range str {
		fmt.Printf("%d:%c ", k, v)
	}

	var intArr = [3]int{1, 2, 3}
	fmt.Printf("%T %T\n", intArr[0:2], intArr)

	fmt.Println("=======================4.5======================")

	c := make(chan int, 1)

	c <- 1

	fmt.Println("end put 1 to channel\n", <-c)
	time.Sleep(1000)
	c <- 2
	fmt.Println("end put 2 to channel\n", <-c)
	time.Sleep(1000)
	c <- 3
	fmt.Println("end put 3 to channel\n", <-c)
	time.Sleep(1000)
	close(c)

	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
		fmt.Println(3)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

}
