package main

import "fmt"

func main() {
	// 通过切片截取字符串
	str := "hello,world!"
	s1 := str[0:5]
	fmt.Println("s1:", s1)
	s2 := str[6:]
	fmt.Println("s2:", s2)

	/*
		由于字符串是不可变的，所以str[0] = '1'无法通过编译
		因为slice的本质是指针指向一个底层数组，所以s1[0] = '1'也无法通过编译
		str[0] = '1'
		s1[0] = '1'
	*/

	// 通过切片修改字符串，存在中文的情况下，需要将[]byte修改为[]rune
	// 具体步骤：先将字符串转换为[]byte或[]rune切片，通过操作字符进行修改，然后再转换为字符串
	s := []byte(str)
	s = s[0:6]
	s = append(s, 'g', 'o', 'l', 'a', 'n', 'g')
	str = string(s)
	fmt.Println(str)

}
