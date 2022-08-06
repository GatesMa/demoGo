package main

import (
	"fmt"
	"runtime"
)

func main() {
	//var wg sync.WaitGroup

	fmt.Println(runtime.NumCPU())

}
