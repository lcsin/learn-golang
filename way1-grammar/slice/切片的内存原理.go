package main

import "fmt"

/*
1. 切片的本质是一个指针指向一个底层数组
2. 多个不同的切片可能指向同一个底层数组，此时它们的指针地址相同
*/
func main() {
	// 以array作为切片的底层数组创建切片，切片的指针指向了array[0]的位置
	array := [...]int{1, 2, 3, 4, 5}
	slice1 := array[0:3]
	fmt.Printf("array: %#v\n", array)
	fmt.Printf("slice: %#v\n", slice1)
	fmt.Println("=====================")

	// 通过指针修改slice1的值，也会修改底层数组array的值
	var p = &slice1[1]
	*p = 199
	fmt.Println("array:", array)
	fmt.Println("slice:", slice1)
	fmt.Println("=====================")

	// 修改切片的内容，实质就是修改该切片指向的底层数组的内容
	slice1[2] = 10086
	fmt.Println("array:", array)
	fmt.Println("slice:", slice1)
	fmt.Println("=====================")

	// 通过slice1创建的slice2共同指向了一个底层数组即array
	// 所以对slice2进行修改也会修改slice1和array的值
	slice2 := slice1[0:2]
	slice2[0] = 110
	slice2[1] = 120
	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("=====================")

	// array地址与slice1,slice2的指针的值相同，因为slice1和slice2都指向了array作为底层数组
	fmt.Printf("%p\n", &array)
	fmt.Printf("%p\n", slice1)
	fmt.Printf("%p\n", slice2)
	fmt.Println("=====================")

	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(array)
	fmt.Println(slice1)
}
