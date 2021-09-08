package main

import "fmt"

/*
1. 当slice容量不足时，会重新分配内存，并进行扩容，容量大小扩容为之前的两倍
2. 重新分配内存表示slice所指向的底层数组发生变化，因此slice指针的地址会发生改变，但slice本身的地址不变
*/
func main() {
	// 初始化一个长度为3，容量为3的切片
	slice := []int{1, 2, 3}
	fmt.Println(slice, len(slice), cap(slice))

	// 通过append()函数，每次往切片中追加一个元素
	for i := 0; i < 15; i++ {
		slice = append(slice, i+100)
		fmt.Printf("%p ", slice)
		fmt.Printf("%p ", &slice)
		fmt.Println(slice, len(slice), cap(slice))
	}
}
