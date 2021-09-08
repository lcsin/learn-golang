package main

import "fmt"

func main() {
	// 声明一个int类型的切片slice1,没有初始化表示还没有分配内存是一个空切片，值为nil即[]，
	var slice1 []int
	fmt.Printf("slice1: %#v\n", slice1)

	// 声明并初始化一个int类型的切片slice2，值为[1 2 3 4]
	slice2 := []int{1, 2, 3, 4}
	fmt.Printf("slice2: %#v\n", slice2)

	// 使用make()函数创建一个int类型的切片，值默认为数据类型的零值
	// make([]type, len, cap)，不指定cap时，cap默认大小为len
	slice3 := make([]int, 4, 8)
	slice4 := make([]int, 4)
	fmt.Printf("slice3: %#v\n", slice3)
	fmt.Printf("slice4: %#v len: %d cap: %d\n", slice4, len(slice4), cap(slice4))

	// 通过数组创建切片，该数组将作为切片的底层数组
	array := [...]int{1, 2, 3, 4, 5}
	slice5 := array[0:3]
	fmt.Printf("array: %#v\n", array)
	fmt.Printf("slice5: %#v\n", slice5)
}
