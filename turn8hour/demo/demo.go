package main

import "fmt"

func main() {
	a := []int{}
	println(a == nil)
	println(len(a), " ", cap(a))
	a = make([]int, 3, 6)
	fmt.Printf("%v\n", a)
	println(len(a), " ", cap(a))
	println(a)
	fmt.Println(a)

	for i, v := range a {
		println(i, " ", v)
	}

	fmt.Println("=========")
	var b []int
	println(b == nil)
}
