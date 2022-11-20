package main

import "fmt"

func main() {

	fmt.Println("================================================")

	var _ [3]string

	arr1 := [...]int{1, 2, 3}
	fmt.Println(arr1[0], arr1[len(arr1)-1])

	for i, v := range arr1 {
		fmt.Printf("%d %d\n", i, v)
	}

	var arr2 = [3]string{"a", "b", "c"}
	for i, v := range arr2 {
		fmt.Printf("%d %#v\n", i, v)
	}

}
