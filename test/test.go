package main

import "context"

type Data struct {
	name string
}

type TestInterface interface {
	Method() int
}

func (data *Data) Method() (a int) {
	return
}

func main() {
	//var wg sync.WaitGroup

	//fmt.Println(runtime.NumCPU())
	//var data Data
	//
	//fmt.Printf("%v\n", data)
	//data.name = "dadada"
	//
	//fmt.Printf("%v\n", data)

	todo := context.TODO()

	cancel, cancelFunc := context.WithCancel(todo)

	cancel.Value("")
	cancelFunc()

}
