流程控制是每种编程语言控制逻辑走向和执行次序的重要部分，流程控制可以说是一门语言的“经脉”。

Go 语言的常用流程控制有 if 和 for，而 switch 和 goto 主要是为了简化代码、降低重复代码而生的结构，属于扩展类的流程控制。

本章主要介绍了 Go 语言中的基本流程控制语句，包括分支语句（if 和 switch）、循环（for）和跳转（goto）语句。另外，还有循环控制语句（break 和 continue），前者的功能是中断循环或者跳出 switch 判断，后者的功能是继续 for 的下一个循环。

## Go语言if else（分支结构）

```go
var ten int = 11
if ten > 10 {
    fmt.Println(">10")
} else {
    fmt.Println("<=10")
}
```

**特殊写法**

if 还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断，代码如下：

```
if err := Connect(); err != nil {
    fmt.Println(err)
    return
}
```

Connect 是一个带有返回值的函数，err:=Connect() 是一个语句，执行 Connect 后，将错误保存到 err 变量中。

err != nil 才是 if 的判断表达式，当 err 不为空时，打印错误并返回。

这种写法可以将返回值与判断放在一行进行处理，而且返回值的作用范围被限制在 if、else 语句组合中。

## Go语言for（循环结构）

与多数语言不同的是，Go语言中的循环语句只支持 for 关键字，而不支持 while 和 do-while 结构，关键字 for 的基本使用方法与C语言和 C++ 中非常接近

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

可以看到比较大的一个不同在于 for 后面的条件表达式不需要用圆括号`()`括起来，Go语言还进一步考虑到无限循环的场景，让开发者不用写无聊的 `for(;;){}`和`do{} while(1);`，而直接简化为如下的写法：

```go
sum := 0
for {
    sum++
    if sum > 100 {
        break
    }
}
```

使用循环语句时，需要注意的有以下几点：

* 左花括号`{`必须与 for 处于同一行。
* Go语言中的 for 循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别是，Go语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量。
* Go语言的 for 循环同样支持 continue 和 break 来控制循环，但是它提供了一个更高级的 break，可以选择中断哪一个循环，如下例：

## Go语言for range（键值循环）

for range 结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用，for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句，一般形式为：

```go
for key, val := range coll {
    ...
}
```

需要要注意的是，val 始终为集合中**对应索引的值拷贝**，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。一个字符串是 Unicode 编码的字符（或称之为 rune ）集合，因此也可以用它来迭代字符串：

### 遍历数组、切片——获得索引和值

在遍历代码中，key 和 value 分别代表切片的下标及下标对应的值，下面的代码展示如何遍历切片，数组也是类似的遍历方法：

```go
for key, value := range []int{1, 2, 3, 4} {
    fmt.Printf("key:%d  value:%d\n", key, value)
}
```

### 遍历通道（channel）——接收通道数据

for range 可以遍历通道（channel），但是通道在遍历时，只输出一个值，即管道内的类型对应的数据。

下面代码为我们展示了通道的遍历：

```go
c := make(chan int)
go func() {
    c <- 1
    c <- 2
    c <- 3
    close(c)
}()
for v := range c {
    fmt.Println(v)
}
```

## Go语言switch case语句

Go语言的 switch 要比C语言的更加通用，表达式不需要为常量，甚至不需要为整数，case 按照从上到下的顺序进行求值，直到找到匹配的项，如果 switch 没有表达式，则对 true 进行匹配，因此，可以将 if else-if else 改写成一个 switch。

### 基本写法

Go语言改进了 switch 的语法设计，case 与 case 之间是独立的代码块，不需要通过 break 语句跳出当前 case 代码块以避免执行到下一行，示例代码如下：
