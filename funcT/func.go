package main

import "fmt"



func foo1(a string, b int) int {
	return 100
}


// 返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	return 666, 777
}

// 返回多个返回值，有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	// r1 r2 属于foo3的形参，初始化默认的值是0
	// r1 r2 作用域空间 是foo3 整个函数体的{}空间
	fmt.Println("r1 = ", r1) // 0
	fmt.Println("r2 = ", r2) // 0

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	fmt.Println(foo1("1", 2))
	fmt.Println(foo2("1", 2))
	fmt.Println(foo3("1", 2))
	fmt.Println(foo4("1", 2))

}


