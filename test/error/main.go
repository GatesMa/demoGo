package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("start")
	err := rtErr(10000)
	fmt.Println(err)
	fmt.Println("end")

}

func rtErr(intput int) error {
	fmt.Println("func start")
	if intput == 10000 {
		panic("panic panic panic")
	}
	return errors.New("error happen")
}
