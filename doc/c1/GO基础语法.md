# 

## Go语言变量的声明（使用var关键字）

Go语言是静态类型语言，因此变量（variable）是有明确类型的，编译器也会检查变量类型的正确性。在数学概念中，变量表示没有固定值且可改变的数。但从计算机系统实现角度来看，变量是一段或多段用来存储数据的内存。

声明变量的一般形式是使用 var 关键字：

```
var name type
```

其中，var 是声明变量的关键字，name 是变量名，type 是变量的类型。

需要注意的是，Go语言和许多编程语言不同，它在声明变量时将变量的类型放在变量的名称之后。这样做的好处就是可以避免像C语言中那样含糊不清的声明形式，例如：`int* a, b`; 。其中只有 a 是指针而 b 不是。如果你想要这两个变量都是指针，则需要将它们分开书写。而在 Go 中，则可以和轻松地将它们都声明为指针类型：

```
var a, b *int
```

**Go语言的基本类型有：**

- bool
- string
- int、int8、int16、int32、int64
- uint、uint8、uint16、uint32、uint64、uintptr
- byte // uint8 的别名
- rune // int32 的别名 代表一个 Unicode 码
- float32、float64
- complex64、complex128

**批量格式**

觉得每行都用 var 声明变量比较烦琐？没关系，还有一种为懒人提供的定义变量的方法：

```
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
```

**简短格式**

`名字 := 表达式`

## 匿名变量

匿名变量的特点是一个下画线“\_”，“\_”本身就是一个特殊的标识符，被称为空白标识符。
它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，
因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，
只需要在变量声明的地方使用下画线替换即可。例如：

## Go语言变量的作用域

一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为作用域。

了解变量的作用域对我们学习Go语言来说是比较重要的，因为Go语言会在编译时检查每个变量是否使用过，一旦出现未使用的变量，就会报编译错误。如果不能理解变量的作用域，就有可能会带来一些不明所以的编译错误。

根据变量定义位置的不同，可以分为以下三个类型：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

**局部变量**

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，函数的参数和返回值变量都属于局部变量。
局部变量不是一直存在的，它只在定义它的函数被调用后存在，函数调用结束后这个局部变量就会被销毁。

**全局变量**

在函数体外声明的变量称之为全局变量，全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，
当然，不包含这个全局变量的源文件需要使用“import”关键字引入全局变量所在的源文件之后才能使用这个全局变量。

全局变量声明必须以 var 关键字开头，****如果想要在外部包中使用全局变量的首字母必须大写****。

Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。

**形式参数**

在定义函数时函数名后面括号中的变量叫做形式参数（简称形参）。形式参数只在函数调用时才会生效，函数调用结束后就会被销毁，在函数未被调用时，函数的形参并不占用实际的存储单元，也没有实际值。

形式参数会作为函数的局部变量来使用。

## Go语言整型（整数类型）

Go语言的数值类型分为以下几种：整数、浮点数、复数，其中每一种都包含了不同大小的数值类型，例如有符号整数包含 int8、int16、int32、int64 等，每种数值类型都决定了对应的大小范围和是否支持正负符号。本节我们主要介绍一下整数类型。

用来表示 Unicode 字符的 rune 类型和 int32 类型是等价的，通常用于表示一个 Unicode 码点。这两个名称可以互换使用。同样，byte 和 uint8 也是等价类型，byte 类型一般用于强调数值是一个原始的数据而不是一个小的整数。

最后，还有一种无符号的整数类型 uintptr，它没有指定具体的 bit 大小但是足以容纳指针。uintptr 类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

尽管在某些特定的运行环境下 int、uint 和 uintptr 的大小可能相等，但是它们依然是不同的类型，比如 int 和 int32，虽然 int 类型的大小也可能是 32 bit，但是在需要把 int 类型当做 int32 类型使用的时候必须显示的对类型进行转换，反之亦然。

Go语言中有符号整数采用 2 的补码形式表示，也就是最高 bit 位用来表示符号位，一个 n-bit 的有符号数的取值范围是从 -2(n-1) 到 2(n-1)-1。无符号整数的所有 bit 位都用于表示非负数，取值范围是 0 到 2n-1。例如，int8 类型整数的取值范围是从 -128 到 127，而 uint8 类型整数的取值范围是从 0 到 255。

**哪些情况下使用 int 和 uint**

程序逻辑对整型范围没有特殊需求。例如，对象的长度使用内建 len() 函数返回，这个长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用 int 来表示。

反之，在二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用 int 和 uint。

## Go语言浮点类型（小数类型）

Go语言提供了两种精度的浮点数 float32 和 float64，它们的算术规范由 IEEE754 浮点数国际标准定义，该浮点数规范被所有现代的 CPU 支持。

这些浮点数类型的取值范围可以从很微小到很巨大。浮点数取值范围的极限值可以在 math 包中找到：

- 常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38；
- 常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308；
- float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。

一个 float32 类型的浮点数可以提供大约 6 个十进制数的精度，而 float64 则可以提供约 15 个十进制数的精度，通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大。

## Go语言bool类型（布尔类型）

一个布尔类型的值只有两种：`true 或 false`。`if 和 for` 语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。

一元操作符!对应逻辑非操作，因此!true的值为 false，更复杂一些的写法是(!true==false) ==true，实际开发中我们应尽量采用比较简洁的布尔表达式，就像用 x 来表示x==true。

```
var aVar = 10
aVar == 5  // false
aVar == 10 // true
aVar != 5  // true
aVar != 10 // false
```

Go语言对于值之间的比较有非常严格的限制，只有两个相同类型的值才可以进行比较，如果值的类型是接口（interface），那么它们也必须都实现了相同的接口。如果其中一个值是常量，那么另外一个值可以不是常量，但是类型必须和该常量类型相同。如果以上条件都不满足，则必须将其中一个值的类型转换为和另外一个值的类型相同之后才可以进行比较。

布尔值可以和 &&（AND）和 ||（OR）操作符结合，并且有短路行为，如果运算符左边的值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值

Go语言中不允许将整型强制转换为布尔型，代码如下：

```
var n bool
fmt.Println(int(n) * 2)

编译错误，输出如下：
cannot convert n (type bool) to type int
布尔型无法参与数值运算，也无法与其他类型进行转换。

数字到布尔型的逆转换非常简单，不过为了保持对称，我们也可以封装一个函数：
// itob报告是否为非零。
func itob(i int) bool { return i != 0 }
```

## Go语言字符串

一个字符串是一个不可改变的字节序列，字符串可以包含任意的数据，但是通常是用来包含可读的文本，字符串是 UTF-8 字符的一个序列（当字符为 ASCII 码表上的字符时则占用 1 个字节，其它字符根据需要占用 2-4 个字节）。

UTF-8 是一种被广泛使用的编码格式，是文本文件的标准编码，其中包括 XML 和 JSON 在内也都使用该编码。由于该编码对占用字节长度的不定性，在Go语言中字符串也可能根据需要占用 1 至 4 个字节，这与其它编程语言如 C++、Java 或者 Python 不同（Java 始终使用 2 个字节）。Go语言这样做不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码。

字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容，更深入地讲，字符串是字节的定长数组。

**定义字符串**

可以使用双引号

```
""
```

来定义字符串，字符串中可以使用转义字符来实现换行、缩进等效果，常用的转义字符包括：

- \n：换行符
- \r：回车符
- \t：tab 键
- \u 或 \U：Unicode 字符
- \\：反斜杠自身

一般的比较运算符（==、!=、<、<=、>=、>）是通过在内存中按字节比较来实现字符串比较的，因此比较的结果是字符串自然编码的顺序。字符串所占的字节长度可以通过函数 len() 来获取，例如 len(str)。

字符串的内容（纯字节）可以通过标准索引法来获取，在方括号

```
[ ]
```

内写入索引，索引从 0 开始计数：

- 字符串 str 的第 1 个字节：str[0]
- 第 i 个字节：str[i - 1]
- 最后 1 个字节：str[len(str)-1]

需要注意的是，这种转换方案只对纯 ASCII 码的字符串有效。

> 注意：获取字符串中某个字节的地址属于非法行为，例如 &str[i]。

**字符串拼接符“+”**

两个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起。将 s2 追加到 s1 尾部并生成一个新的字符串 s。

**字符串实现基于 UTF-8 编码**

Go语言中字符串的内部实现使用 UTF-8 编码，通过 rune 类型，可以方便地对每个 UTF-8 字符进行访问。当然，Go语言也支持按照传统的 ASCII 码方式逐字符进行访问。

**定义多行字符串**

在Go语言中，使用双引号书写字符串的方式是字符串常见表达方式之一，被称为字符串字面量（string literal），这种双引号字面量不能跨行，如果想要在源码中嵌入一个多行字符串时，就必须使用`反引号，代码如下：

```
const str = `第一行
第二行
第三行
\r\n
`
fmt.Println(str)
```

### 1. Go语言计算字符串长度——len()和RuneCountInString()

Go 语言的内建函数 len()，可以用来获取切片、字符串、通道（channel）等的长度。下面的代码可以用 len() 来获取字符串的长度。

```
tip1 := "genji is a ninja"
fmt.Println(len(tip1))
tip2 := "忍者"
fmt.Println(len(tip2))

程序输出如下：
16
6
```

下面的代码展示如何计算UTF-8的字符个数。

```
fmt.Println(utf8.RuneCountInString("忍者"))
fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))
```

## 2. Go语言遍历字符串——获取每一个字符串元素

**遍历每一个ASCII字符**

遍历 ASCII 字符使用 for 的数值循环进行遍历，直接取每个字符串的下标获取 ASCII 字符，如下面的例子所示。

```
theme := "狙击 start"
for i := 0; i < len(theme); i++ {
    fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
}
```

这种模式下取到的汉字“惨不忍睹”。由于没有使用 Unicode，汉字被显示为乱码。

**按Unicode字符遍历字符串**

```
theme := "狙击 start"
for _, s := range theme {
    fmt.Printf("Unicode: %c  %d\n", s, s)
}
```

### 总结

- ASCII 字符串遍历直接使用下标。
- Unicode 字符串遍历用 for range。

### 3. Go语言字符串截取（获取字符串的某一段字符）

```
tracer := "死神来了, 死神bye bye"
comma := strings.Index(tracer, ", ")
pos := strings.Index(tracer[comma:], "死神")
fmt.Println(comma, pos, tracer[comma+pos:])

12 2 死神bye bye
```

字符串索引比较常用的有如下几种方法：
strings.Index：正向搜索子字符串。
strings.LastIndex：反向搜索子字符串。
搜索的起始位置可以通过切片偏移制作。

### 4. Go语言字符串拼接（连接）

连接字符串这么简单，还需要学吗？确实，Go 语言和大多数其他语言一样，使用+对字符串进行连接操作，非常直观。

但问题来了，好的事物并非完美，简单的东西未必高效。除了加号连接字符串，Go 语言中也有类似于 StringBuilder 的机制来进行高效的字符串连接，例如：

```
hammer := "吃我一锤"
sickle := "死吧"
// 声明字节缓冲
var stringBuilder bytes.Buffer
// 把字符串写入缓冲
stringBuilder.WriteString(hammer)
stringBuilder.WriteString(sickle)
// 将缓冲以字符串形式输出
fmt.Println(stringBuilder.String())
```

bytes.Buffer 是可以缓冲并可以往里面写入各种字节数组的。字符串也是一种字节数组，使用 WriteString() 方法进行写入。

将需要连接的字符串，通过调用 WriteString() 方法，写入 stringBuilder 中，然后再通过 stringBuilder.String() 方法将缓冲转换为字符串。

### 5. Go语言fmt.Sprintf（格式化输出）

格式化在逻辑中非常常用。使用格式化函数，要注意写法：`fmt.Sprintf(格式化样式, 参数列表…)`

```
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
fmt.Printf("使用'%%+v' %+v\n", profile)
fmt.Printf("使用'%%#v' %#v\n", profile)
fmt.Printf("使用'%%T' %T\n", profile)
```

```
已采集2个药草, 还需要8个完成任务
月球基地 3.14159 true
使用'%+v' &{Name:rat HP:150}
使用'%#v' &struct { Name string; HP int }{Name:"rat", HP:150}
使用'%T' *struct { Name string; HP int }

```


| -   | -                                         |
| --- | ----------------------------------------- |
| %v  | 按值的本来值输出                          |
| %+v | 在 %v 基础上，对结构体字段名和值进行展开 |
| %#v | 输出 Go 语言语法格式的值                  |
| %T  | 输出 Go 语言语法格式的类型和值            |
| %%  | 输出 % 本体                              |
| %b  | 整型以二进制方式显示                      |
| %o  | 整型以八进制方式显示                      |
| %d  | 整型以十进制方式显示                      |
| %x  | 整型以十六进制方式显示                    |
| %X  | 整型以十六进制、字母大写方式显示          |
| %U  | Unicode 字符                             |
| %f  | 浮点数                                    |
| %p  | 指针，十六进制方式显示                    |

## Go语言字符类型（byte和rune）

字符串中的每一个元素叫做“字符”，在遍历或者单个获取字符串元素时可以获得字符。

Go语言的字符有以下两种：

- 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
- 另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。

在 ASCII 码表中，A 的值是 65，使用 16 进制表示则为 41，所以下面的写法是等效的：

`var ch byte = 65 或 var ch byte = '\x41'      //（\x 总是紧跟着长度为 2 的 16 进制数）`

```
Unicode 与 ASCII 类似，都是一种字符集。

字符集为每个字符分配一个唯一的 ID，我们使用到的所有字符在 Unicode 字符集中都有一个唯一的 ID，例如上面例子中的 a 在 Unicode 与 ASCII 中的编码都是 97。汉字“你”在 Unicode 中的编码为 20320，在不同国家的字符集中，字符所对应的 ID 也会不同。而无论任何情况下，Unicode 中的字符的 ID 都是不会变化的。

UTF-8 是编码规则，将 Unicode 中字符的 ID 以某种方式进行编码，UTF-8 的是一种变长编码规则，从 1 到 4 个字节不等。编码规则如下：
0xxxxxx 表示文字符号 0～127，兼容 ASCII 字符集。
从 128 到 0x10ffff 表示其他字符。

根据这个规则，拉丁文语系的字符编码一般情况下每个字符占用一个字节，而中文每个字符占用 3 个字节。

广义的 Unicode 指的是一个标准，它定义了字符集及编码规则，即 Unicode 字符集和 UTF-8、UTF-16 编码等。

```

## Go语言数据类型转换

在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明：

`valueOfTypeB = typeB(valueOfTypeA)`

类型转换只能在定义正确的情况下转换成功，例如从一个取值范围较小的类型转换到一个取值范围较大的类型（将 int16 转换为 int32）。当从一个取值范围较大的类型转换到取值范围较小的类型时（将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况。

只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型），不同底层类型的变量相互转换时会引发编译错误（如将 bool 类型转换为 int 类型）：

```
package main
import (
        "fmt"
        "math"
)
func main() {
        // 输出各数值范围
        fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
        fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
        fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
        fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
        // 初始化一个32位整型值
        var a int32 = 1047483647
        // 输出变量的十六进制形式和十进制值
        fmt.Printf("int32: 0x%x %d\n", a, a)
        // 将a变量数值转换为十六进制, 发生数值截断
        b := int16(a)
        // 输出变量的十六进制形式和十进制值
        fmt.Printf("int16: 0x%x %d\n", b, b)
        // 将常量保存为float32类型
        var c float32 = math.Pi
        // 转换为int类型, 浮点发生精度丢失
        fmt.Println(int(c))
}

int8 range: -128 127
int16 range: -32768 32767
int32 range: -2147483648 2147483647
int64 range: -9223372036854775808 9223372036854775807
int32: 0x3e6f54ff 1047483647
int16: 0x54ff 21759
3

```

## Go语言指针详解

与 Java 和 .NET 等编程语言不同，Go语言为程序员提供了控制数据结构指针的能力，但是，并不能进行指针运算。Go语言允许你控制特定集合的数据结构、分配的数量以及内存访问模式，这对于构建运行良好的系统是非常重要的。指针对于性能的影响不言而喻，如果你想要做系统编程、操作系统或者网络应用，指针更是不可或缺的一部分。

指针（pointer）在Go语言中可以被拆分为两个核心概念：

- 类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
- 切片，由指向起始元素的原始指针、元素数量和容量组成。

受益于这样的约束和拆分，Go语言的指针类型变量即拥有指针高效访问的特点，又不会发生指针偏移，从而避免了非法修改关键性数据的问题。同时，垃圾回收也比较容易对不会发生偏移的指针进行检索和回收。

切片比原始指针具备更强大的特性，而且更为安全。切片在发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。

**认识指针地址和指针类型**

一个指针变量可以指向任何一个值的内存地址，它所指向的值的内存地址在 32 和 64 位机器上分别占用 4 或 8 个字节，占用字节的大小与所指向的值的大小无关。当一个指针被定义后没有分配到任何变量时，它的默认值为 nil。指针变量通常缩写为 ptr。

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作），格式如下：

`ptr := &v    // v 的类型为 T`

其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做 T 的指针类型，*代表指针。

**从指针获取指针指向的值**

当使用&操作符对普通变量进行取地址操作并得到变量的指针后，可以对指针使用*操作符，也就是指针取值，代码如下:

```
package main
import (
    "fmt"
)
func main() {
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
}


ptr type: *string
address: 0xc000010310
value type: string
value: Malibu Point 10880, 90265
```

**使用指针修改值**

通过指针不仅可以取值，也可以修改值。

```
// 交换函数
func swap(a, b *int) {
    // 取a指针的值, 赋给临时变量t
    t := *a
    // 取b指针的值, 赋给a指针指向的变量
    *a = *b
    // 将a指针的值赋给b指针指向的变量
    *b = t
}
func main() {
// 准备两个变量, 赋值1和2
    x, y := 1, 2
    // 交换变量值
    swap(&x, &y)
    // 输出变量值
    fmt.Println(x, y)
}
```

**示例：使用指针变量获取命令行的输入信息**

Go语言内置的 flag 包实现了对命令行参数的解析，flag 包使得开发命令行工具更为简单。

下面的代码通过提前定义一些命令行指令和对应的变量，并在运行时输入对应的参数，经过 flag 包的解析后即可获取命令行的数据。

```
package main
// 导入系统包
import (
    "flag"
    "fmt"
)
// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
func main() {
    // 解析命令行参数
    flag.Parse()
    // 输出命令行参数
    fmt.Println(*mode)
}
```

将这段代码命名为 main.go，然后使用如下命令行运行：
`go run main.go --mode=fast`

命令行输出结果如下：
`fast`

代码说明如下：

```
第 10 行，通过 flag.String，定义一个 mode 变量，这个变量的类型是 *string。后面 3 个参数分别如下：
参数名称：在命令行输入参数时，使用这个名称。
参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。
参数说明：使用 -help 时，会出现在说明中。
第 15 行，解析命令行参数，并将结果写入到变量 mode 中。
第 18 行，打印 mode 指针所指向的变量。
```

**创建指针的另一种方法——new() 函数**

Go语言还提供了另外一种方法来创建指针变量，格式如下：`new(类型)`
一般这样写：

```
str := new(string)
*str = "Go语言教程"
fmt.Println(*str)

new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
```

## Go语言变量的生命周期

变量的生命周期指的是在程序运行期间变量有效存在的时间间隔。

变量的生命周期与 变量的作用域 有着不可分割的联系：

- 全局变量：它的生命周期和整个程序的运行周期是一致的；
- 局部变量：它的生命周期则是动态的，从创建这个变量的声明语句开始，到这个变量不再被引用为止；
- 形式参数和函数返回值：它们都属于局部变量，在函数被调用的时候创建，函数调用结束后被销毁。

栈和堆的区别在于：

- 堆（heap）：堆是用于存放进程执行中被动态分配的内存段。它的大小并不固定，可动态扩张或缩减。当进程调用 malloc 等函数分配内存时，新分配的内存就被动态加入到堆上（堆被扩张）。当利用 free 等函数释放内存时，被释放的内存从堆中被剔除（堆被缩减）；
- 栈(stack)：栈又称堆栈， 用来存放程序暂时创建的局部变量，也就是我们函数的大括号`{ }`中定义的局部变量。

在程序的编译阶段，编译器会根据实际情况自动选择在栈或者堆上分配局部变量的存储空间，不论使用 var 还是 new 关键字声明变量都不会影响编译器的选择。

```
var global *int
func f() {
    var x int
    x = 1
    global = &x
}
func g() {
    y := new(int)
    *y = 1
}
```

上述代码中，函数 f 里的变量 x 必须在堆上分配，因为它在函数退出后依然可以通过包一级的 global 变量找到，虽然它是在函数内部定义的。用Go语言的术语说，这个局部变量 x 从函数 f 中逃逸了。

相反，当函数 g 返回时，变量 *y 不再被使用，也就是说可以马上被回收的。因此，*y 并没有从函数 g 中逃逸，编译器可以选择在栈上分配 *y 的存储空间，也可以选择在堆上分配，然后由Go语言的 GC（垃圾回收机制）回收这个变量的内存空间。

在实际的开发中，并不需要刻意的实现变量的逃逸行为，因为逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。

虽然Go语言能够帮助我们完成对内存的分配和释放，但是为了能够开发出高性能的应用我们任然需要了解变量的声明周期。例如，如果将局部变量赋值给全局变量，将会阻止 GC 对这个局部变量的回收，导致不必要的内存占用，从而影响程序的性能。

## Go语言常量和const关键字

Go语言中的常量使用关键字 const 定义，用于存储不会改变的数据，常量是在编译时被创建的，即使定义在函数内部也是如此，并且只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。由于编译时的限制，定义常量的表达式必须为能被编译器求值的常量表达式。

`const pi = 3.14159 // 相当于 math.Pi 的近似值`

所有常量的运算都可以在编译期完成，这样不仅可以减少运行时的工作，也方便其他代码的编译优化，当操作数是常量时，一些运行时的错误也可以在编译时被发现，例如整数除零、字符串索引越界、任何导致无效浮点数的操作等。

常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是返回常量结果：len、cap、real、imag、complex 和 unsafe.Sizeof。

因为它们的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类型的长度：`var p [IPv4Len]byte`

**iota 常量生成器**

常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一。

```
type Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

**无类型常量**

Go语言的常量有个不同寻常之处。虽然一个常量可以有任意一个确定的基础类型，例如 int 或 float64，或者是类似 time.Duration 这样的基础类型，但是许多常量并没有一个明确的基础类型。

编译器为这些没有明确的基础类型的数字常量提供比基础类型更高精度的算术运算，可以认为至少有 256bit 的运算精度。这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。

通过延迟明确常量的具体类型，不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。

对于常量面值，不同的写法可能会对应不同的类型。例如 0、0.0、0i 和 \u0000 虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型。同样，true 和 false 也是无类型的布尔类型，字符串面值常量是无类型的字符串类型。

## Go语言type关键字（类型别名）

类型别名是 Go 1.9 版本添加的新功能，主要用于解决代码升级、迁移中存在的类型兼容性问题。在 C/C++ 语言中，代码重构升级可以使用宏快速定义一段新的代码，Go语言中没有选择加入宏，而是解决了重构中最麻烦的类型名变更问题。

```
这种是创建了新类型
type byte uint8
type rune int32

类型别名
type byte = uint8
type rune = int32
```

## Go语言关键字与标识符简述

Go语言的词法元素包括 5 种，分别是标识符（identifier）、关键字（keyword）、操作符（operator）、分隔符（delimiter）、字面量（literal），它们是组成Go语言代码和程序的最基本单位。

**关键字**
关键字即是被Go语言赋予了特殊含义的单词，也可以称为保留字。

Go语言中的关键字一共有 25 个：


| -        | -           | -      | -         | -      |
| -------- | ----------- | ------ | --------- | ------ |
| break    | default     | func   | interface | select |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

**标识符**


| -      | -       | -       | -       | -      | -       | -         | -          | -       |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |
