package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(word string) {
	for {
		fmt.Println(word)
		time.Sleep(time.Second)
	}
}

func main() {
	go say("1")
	println("runtime.NumCPU", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	println("runtime.NumCPU", runtime.NumCPU())
	say("2")

}
