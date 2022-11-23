package main

import "fmt"

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

}
