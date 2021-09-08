package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{8, 9, 10}

	// append()函数支持追加一个或多个元素以及切片
	s3 := append(s1, 4)
	s4 := append(s1, 5, 6, 7)
	s5 := append(s1, s2...)

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)
	fmt.Println("s5: ", s5)
}
