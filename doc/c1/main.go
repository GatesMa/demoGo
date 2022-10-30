package main

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	var name = "gatesma"
	print(name)

	//Go语言的基本类型有：
	//bool
	//string
	//int、int8、int16、int32、int64
	//uint、uint8、uint16、uint32、uint64、uintptr
	//byte // uint8 的别名
	//rune // int32 的别名 代表一个 Unicode 码
	//float32、float64
	//complex64、complex128

	f := 0.0
	println("\n" + reflect.TypeOf(f).String())
	i := 0
	println("\n" + reflect.TypeOf(i).String())

	var b bool
	println("\n"+reflect.TypeOf(b).String(), b)

	print("\n=====================\n")

	//var (
	//	a int
	//	b string
	//	c []float32
	//	d func() bool
	//	e struct {
	//		x int
	//	}
	//)

	//除 var 关键字外，还可使用更加简短的变量定义和初始化语法。
	//名字 := 表达式
	//
	//需要注意的是，简短模式（short variable declaration）有以下限制：
	//定义变量，同时显式初始化。
	//不能提供数据类型。
	//只能用在函数内部。

	//整型和浮点型变量的默认值为 0 和 0.0。
	//字符串变量的默认值为空字符串。
	//布尔型变量默认为 false。
	//切片、函数、指针变量的默认为 nil。
	//float 默认是 float64，为了不减少精度

	print("\n=====================\n")
	//匿名变量
	/**
	匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符。
	它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，
	因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，
	只需要在变量声明的地方使用下画线替换即可。例如：
	*/

	//局部变量
	/**
	在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，函数的参数和返回值变量都属于局部变量。
	局部变量不是一直存在的，它只在定义它的函数被调用后存在，函数调用结束后这个局部变量就会被销毁。
	*/

	//全局变量
	/**
	在函数体外声明的变量称之为全局变量，全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，
	当然，不包含这个全局变量的源文件需要使用“import”关键字引入全局变量所在的源文件之后才能使用这个全局变量。
	全局变量声明必须以 var 关键字开头，****如果想要在外部包中使用全局变量的首字母必须大写****。
	*/

	var x = 3.4e10
	println("\n", x, "\n")
	fmt.Printf("%.2f\n", x)

	str := "Beginning of the string " +
		"second part of the string"
	fmt.Printf("%s\n", str)

	const str1 = `第一行
第二行
第三行
\r\n
`
	fmt.Println(str1)

	// 遍历字符串
	str2 := "爆龙6634737战士！！!"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("ascii:%c %d\n", str2[i], str2[i])
	}
	//按Unicode字符遍历字符串
	for _, s := range str2 {
		fmt.Printf("unicode:%c %d\n", s, s)
	}

	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ")
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])

	hammer := "吃我一锤"
	sickle := "死吧"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())
	fmt.Sprintf("%d", 'c')

	fmt.Println("\n=========")
	var progress = 2
	var target = 8
	// 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)
	fmt.Println(title)
	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)
	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v' %v\n", profile)
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var d int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", d, d)
	// 将a变量数值转换为十六进制, 发生数值截断
	bb := int16(d)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", bb, bb)
	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))

	fmt.Println("\n================指针==================")

	var car = 1
	var str3 = "banana"
	fmt.Printf("%p %p\n", &car, &str3)

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	sa := 1
	sb := 2
	swap(&sa, &sb)
	fmt.Printf("\n%d, %d\n", sa, sb)

	str4 := new(string)
	*str4 = "Go语言教程"
	fmt.Println(*str4)

	var x1 float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	fmt.Println(x1, y, z)

}

func swap(a *int, b *int) {
	//// 取a指针的值, 赋给临时变量t
	//t := *a
	//// 取b指针的值, 赋给a指针指向的变量
	//*a = *b
	//// 将a指针的值赋给b指针指向的变量
	//*b = t

	*a, *b = *b, *a
}
