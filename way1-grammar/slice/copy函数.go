package main

import "fmt"

func main() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 使用array作为底层数组创建切片s1，s1的指针指向array[8]的值
	s1 := array[8:]
	// 使用array作为底层数组创建切片s2，s1的指针指向array[5]的值
	s2 := array[5:]
	fmt.Println(array)
	fmt.Println(s1, s2)

	// 使用copy()函数把s2的值copy到s1，copy()函数也会改变切片的底层数组
	// copy函数在把s2的5、6复制到s1时，实质是通过指针将底层数组8、9修改为5、6，因此s1和s2原本8和9的值都变为了5和6
	copy(s1, s2)
	fmt.Println(s1, s2)
	fmt.Println(array)
}
