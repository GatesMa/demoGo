package main

import "fmt"

type Data struct {
	name string
}

func main() {
	//var wg sync.WaitGroup

	//fmt.Println(runtime.NumCPU())
	var data Data

	fmt.Printf("%v\n", data)
	data.name = "dadada"

	fmt.Printf("%v\n", data)

}
