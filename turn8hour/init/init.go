package main

import (
	. "demoGo/turn8hour/init/lib1"
	mylib2 "demoGo/turn8hour/init/lib2"
	"fmt"
)

func init() {
	fmt.Println("init")
}

func main() {
	//lib1.Lib1Test()
	Lib1Test()
	//lib2.Lib2Test()
	mylib2.Lib2Test()
}
