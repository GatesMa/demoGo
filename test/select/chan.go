package main

import "time"

func main() {
	var quit = make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		quit <- 1
	}()

	println("start waiting")
	<-quit
	println("end")

}
