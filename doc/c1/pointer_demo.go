package main

import "fmt"

type animal struct {
	name string
}

func (a animal) setName_v(name string) animal {
	a.name = name
	return a
}

func (a animal) setName_wrong(name *string) {
	a.name = *name //作用域只在函数内部， 垃圾回收了所有内部参数
}

func (a *animal) setName(name *string) {
	a.name = *name //作用域是外部的地址，垃圾并不会回收地址
}
func setName(p *animal, name string) {
	p.name = name //自带的数据机构 可以传 地址可以传value，看具体情况节约内存
}
func main() {
	an := animal{name: "jerry"}
	fmt.Println(an)
	setName(&an, "tommy")
	fmt.Println(an)
	strname := "zhangsan" //数据结构很大建议用地址，这样不用走栈内存。
	an.setName(&strname)
	fmt.Println(an)
	an = an.setName_v("lili")
	fmt.Println(an)
}
