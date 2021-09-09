/**
strconv: 字符串和基本类型之间的转换
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string转换为基本数据类型: strconv.ParseXxx
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseInt("10086", 10, 64))
	fmt.Println(strconv.ParseUint("10086", 10, 64))
	fmt.Println(strconv.ParseFloat("10086", 64))
	fmt.Println(strconv.ParseComplex("10086", 64))

	// 基本数据类型转换为string: strconv.FormatXxx
	fmt.Println(strconv.FormatInt(10086, 2))
	fmt.Println(strconv.FormatUint(10086, 10))
	fmt.Println(strconv.FormatFloat(1.0086, 'b', -1, 64))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatComplex(10086, 'b', -1, 64))

	// 最常用的int和string间的转换: strconv.Atoi()、strconv.Itoa()
	fmt.Println(strconv.Atoi("10086"))
	fmt.Println(strconv.Itoa(10086))
}
