# Go语言函数（Go语言func）

函数是组织好的、可重复使用的、用来实现单一或相关联功能的代码段，其可以提高应用的模块性和代码的重复利用率。

Go 语言支持普通函数、匿名函数和闭包，从设计上对函数进行了优化和改进，让函数使用起来更加方便。

Go 语言的函数属于“一等公民”（first-class），也就是说：

* 函数本身可以作为值进行传递。
* 支持匿名函数和闭包（closure）。
* 函数可以满足接口。

## 1. Go语言函数声明（函数定义）

Go语言里面拥三种类型的函数：

* 普通的带有名字的函数
* 匿名函数或者 lambda 函数
* 方法

```go
func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3,4)) // "5"
```

如果一个函数在声明时，包含返回值列表，那么该函数必须以 return 语句结尾，除非函数明显无法运行到结尾处，例如函数在结尾时调用了 panic 异常或函数中存在无限循环。

函数的类型被称为函数的标识符，如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和标识符，形参和返回值的变量名不影响函数标识符也不影响它们是否可以以省略参数类型的形式表示。

每一次函数在调用时都必须按照声明顺序为所有参数提供实参（参数值），在函数调用时，Go语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

在函数中，实参通过值传递的方式进行传递，因此函数的形参是实参的拷贝，对形参进行修改不会影响实参，但是，如果实参包括引用类型，如指针、slice(切片)、map、function、channel 等类型，实参可能会由于函数的间接引用被修改。

### 函数的返回值

Go语言支持多返回值，多返回值能方便地获得函数执行后的多个返回参数，Go语言经常使用多返回值中的最后一个返回参数返回函数执行中可能发生的错误，示例代码如下：

`conn, err := connectToNetwork()`

## 2. Go语言函数变量——把函数作为值保存到变量中

在Go语言中，函数也是一种类型，可以和其他类型一样保存在变量中，下面的代码定义了一个函数变量 f，并将一个函数名为 fire() 的函数赋给函数变量 f，这样调用函数变量 f 时，实际调用的就是 fire() 函数，代码如下：

```go
package main
import (
    "fmt"
)
func fire() {
    fmt.Println("fire")
}
func main() {
    var f func()
    f = fire
    f()
}
```

## 3. Go语言匿名函数——没有函数名字的函数

#### 1) 在定义时调用匿名函数

```go
func(data int) {
    fmt.Println("hello", data)
}(100)
```

#### 2) 将匿名函数赋值给变量

```go
// 将匿名函数体保存到f()中
f := func(data int) {
    fmt.Println("hello", data)
}
// 使用f()调用
f(100)
```

### 匿名函数用作回调函数

下面的代码实现对切片的遍历操作，遍历中访问每个元素的操作使用匿名函数来实现，用户传入不同的匿名函数体可以实现对元素不同的遍历操作，代码如下：

```go
package main
import (
    "fmt"
)
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
    for _, v := range list {
        f(v)
    }
}
func main() {
    // 使用匿名函数打印切片内容
    visit([]int{1, 2, 3, 4}, func(v int) {
        fmt.Println(v)
    })
}
```

### 使用匿名函数实现操作封装

```go
package main
import (
    "flag"
    "fmt"
)
var skillParam = flag.String("skill", "", "skill to perform")
func main() {
    flag.Parse()
    var skill = map[string]func(){
        "fire": func() {
            fmt.Println("chicken fire")
        },
        "run": func() {
            fmt.Println("soldier run")
        },
        "fly": func() {
            fmt.Println("angel fly")
        },
    }
    if f, ok := skill[*skillParam]; ok {
        f()
    } else {
        fmt.Println("skill not found")
    }
}
```

```
PS D:\code> go run main.go --skill=fly
angel fly
PS D:\code> go run main.go --skill=run
soldier run 
```

## 4. Go语言函数类型实现接口——把函数作为接口来调用

函数和其他类型一样都属于“一等公民”，其他类型能够实现接口，函数也可以，本节将对结构体与函数实现接口的过程进行对比。

首先给出本节完整的代码：

```go
package main

import (
    "fmt"
)

// 调用器接口
type Invoker interface {
    // 需要实现一个Call方法
    Call(interface{})
}

// 结构体类型
type Struct struct {
}

// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
    fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {

    // 调用f函数本体
    f(p)
}

func main() {

    // 声明接口变量
    var invoker Invoker

    // 实例化结构体
    s := new(Struct)

    // 将实例化的结构体赋值到接口
    invoker = s

    // 使用接口调用实例化结构体的方法Struct.Call
    invoker.Call("hello")

    // 将匿名函数转为FuncCaller类型，再赋值给接口
    invoker = FuncCaller(func(v interface{}) {
        fmt.Println("from function", v)
    })

    // 使用接口调用FuncCaller.Call，内部会调用函数本体
    invoker.Call("hello")
}
```

## 5. Go语言闭包（Closure）——引用了外部变量的匿名函数

Go语言中闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或者删除，在闭包中可以继续使用这个自由变量，因此，简单的说：

函数 + 引用环境 = 闭包

#### 其它编程语言中的闭包

闭包（Closure）在某些编程语言中也被称为 Lambda 表达式。

闭包对环境中变量的引用过程也可以被称为“捕获”，在 [C++](http://c.biancheng.net/cplus/)11 标准中，捕获有两种类型，分别是引用和复制，可以改变引用的原值叫做“引用捕获”，捕获的过程值被复制到闭包中使用叫做“复制捕获”。

在 Lua 语言中，将被捕获的变量起了一个名字叫做 Upvalue，因为捕获过程总是对闭包上方定义过的自由变量进行引用。

闭包在各种语言中的实现也是不尽相同的，在 Lua 语言中，无论闭包还是函数都属于 Prototype 概念，被捕获的变量以 Upvalue 的形式引用到闭包中。

C++ 与 [C#](http://c.biancheng.net/csharp/) 中为闭包创建了一个类，而被捕获的变量在编译时放到类中的成员中，闭包在访问被捕获的变量时，实际上访问的是闭包隐藏类的成员。

#### 在闭包内部修改引用的变量

闭包对它作用域上部的变量可以进行修改，修改引用的变量会对变量进行实际修改，通过下面的例子来理解：

```go
// 准备一个字符串
str := "hello world"
// 创建一个匿名函数
foo := func() {
   
    // 匿名函数中访问str
    str = "hello dude"
}
// 调用匿名函数
foo()
```

## 6. Go语言可变参数（变参函数）

在C语言时代大家一般都用过 printf() 函数，从那个时候开始其实已经在感受可变参数的魅力和价值，如同C语言中的 printf() 函数，Go语言标准库中的 fmt.Println() 等函数的实现也依赖于语言的可变参数功能。

可变参数是指函数传入的参数个数是可变的，为了做到这点，首先需要将函数定义为可以接受可变参数的类型

```go
func myfunc(args ...int) {
    for _, arg := range args {
        fmt.Println(arg)
    }
}
```

## 7. Go语言defer（延迟执行语句）

Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

## 8. Go语言处理运行时错误

Go语言的错误处理思想及设计包含以下特征：

* 一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误。
* 在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。


## 9. Go语言宕机（panic）——程序终止运行



Go语言可以在程序中手动触发宕机，让程序崩溃，这样开发者可以及时地发现错误，同时减少可能的损失。

Go语言程序在宕机时，会将堆栈和 goroutine 信息输出到控制台，所以宕机也可以方便地知晓发生错误的位置，那么我们要如何触发宕机呢，示例代码如下所示：

```go
package main
func main() {
    panic("crash")
}
```


## 10. Go语言宕机恢复（recover）——防止程序崩溃

Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。

**panic 和 recover 的关系**

panic 和 recover 的组合有如下特性：* 有 panic 没 recover，程序宕机。

* 有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。

**提示**

虽然 panic/recover 能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性。

在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛，直到程序整体崩溃。

如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置。


## 11. Go语言Test功能测试函数详解

Go语言中提供了 testing 包来实现单元测试功能

要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时文件名必须以`_test.go`结尾，单元测试源码文件可以由多个测试用例（可以理解为函数）组成，每个测试用例的名称需要以 Test 为前缀，例如：

```go
func TestXxx( t *testing.T ){
    //......
}
```
