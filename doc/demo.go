package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {

	fmt.Println(runtime.NumCPU())

	x := "1313414141"
	typeOf := reflect.TypeOf(&x)
	fmt.Println(x, typeOf, typeOf.Kind())

}
