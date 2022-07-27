package main

import "fmt"

const (
	// 可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING  = iota // iota = 0
	SHANGHAI        // iota = 1
	SHENZHEN        // iota = 2
)

func main() {
	const a int = 10
	println(a)
	fmt.Printf("a:%T, %d\n", a, a)
	fmt.Println(SHANGHAI)

}
