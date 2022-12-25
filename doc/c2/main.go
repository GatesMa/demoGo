package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	fmt.Println("================================================")

	var _ [3]string

	arr1 := [...]int{1, 2, 3}
	fmt.Println(arr1[0], arr1[len(arr1)-1])

	for i, v := range arr1 {
		fmt.Printf("%d %d\n", i, v)
	}

	var arr2 = [3]string{"a", "b", "c"}
	for i, v := range arr2 {
		fmt.Printf("%d %#v\n", i, v)
	}

	fmt.Println("================================================")

	var strList = []string{"123"}
	fmt.Printf("%v %T\n", strList, strList)
	strList = append(strList, "2432")
	fmt.Printf("%v %T\n", strList, strList)

	fmt.Println("================================================")

	var nums []int
	for i := 1; i < 10; i++ {
		fmt.Printf("%d %d %v\n", len(nums), cap(nums), nums)
		nums = append(nums, i)
	}
	nums = append([]int{12345}, nums...)
	fmt.Println(nums[0])

	fmt.Println("================================================")

	var slice1 = []int{1, 2, 3}
	var slice2 = []int{5, 6, 7, 8, 9}
	copy(slice1, slice2)
	fmt.Printf("%v %v\n", slice1, slice2)
	var slice3 = []int{}
	copy(slice3, slice1)
	fmt.Printf("%v %v\n", slice1, slice2)
	slice1[1] = 999
	fmt.Printf("%v %v\n", slice1, slice2)

	var a = []int{1, 2, 3}
	var tmp = copy(a, a[1:]) // 删除开头1个元素
	fmt.Printf("%v %v\n", a, tmp)

	fmt.Println("================================================")
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %p ElemAddr: %p\n", value, &value, &slice[index])
	}

	var xxx = [][]int{{}}
	fmt.Printf("%T \n", xxx)

	fmt.Println("================================================")

	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	var strMap map[string]string
	fmt.Printf("%T %v %v\n", strMap, strMap, strMap == nil)
	strMap = make(map[string]string)
	fmt.Printf("%T %v %v\n", strMap, strMap, strMap == nil)
	strMap["name"] = "gatesma"
	fmt.Printf("%T %v\n", strMap, strMap)

	var sliceMap = make(map[string]*[]int, 10)
	sliceMap["666"] = &[]int{1, 2, 3}
	fmt.Printf("666 in sliceMap = %v\n", *sliceMap["666"])
	(*sliceMap["666"])[0] = 1234
	fmt.Printf("666 in sliceMap = %v\n", *sliceMap["666"])

	fmt.Println("================================================")
	// 遍历map
	rangeStrMap := make(map[string]string)
	rangeStrMap["name"] = "gatesma"
	rangeStrMap["hp"] = "100"
	rangeStrMap["height"] = "199"
	for k, v := range rangeStrMap {
		fmt.Printf("%s %s\n", k, v)
	}

	var rangeStrSlice = []string{}
	for k, _ := range rangeStrMap {
		rangeStrSlice = append(rangeStrSlice, k)
	}

	sort.Strings(rangeStrSlice)

	for k, v := range rangeStrSlice {
		fmt.Printf("%d %s\n", k, v)
	}

	fmt.Println("=========================sync.Map=======================")

	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", "200")
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Printf("iterate: %T %v %T %v\n", k, k, v, v)
		return true
	})

}
