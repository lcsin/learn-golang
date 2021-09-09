/*
type关键字的基本使用：
1. 类型定义
2. 类型别名：，

类型定义：
1. 语法：type name Type
2. 类型定义和go内置的数据类型不是同一个类型，无法进行比较和赋值
3. 类似华氏度和摄氏度等可能被大量使用并且具有一定特殊含义的数据类型，可以通过类型定义来定义新的数据类型
4. 函数作为特殊的数据类型也可以用于类型定义，使用自定义的函数数据类型，可以简化代码和提高代码的阅读性

类型别名：
1. 语法：type name = Type
2. 类型别名和go内置的数据类型是同一个类型，可以进行比较和赋值
3. 如byte、rune是uint8和int32的别名
*/
package main

import "fmt"

// 自定义新的类型
type myInt int
type myFunc func(int, int) int

// 使用自定义的新的 func(int, int) int 函数类型不但可以简化代码，还能提高代码的阅读性
func fun1() myFunc {
	return func(i int, i2 int) int {
		return i + i2
	}
}

// RMB 自定义类型别名RMB作为int的别名
type RMB = int

func main() {
	// 自定义的myInt类型和golang内置的基本数据类型(int)不是一个类型，无法进行比较
	var n1 myInt = 1
	var n2 int = 1
	fmt.Printf("n1 type: %T\n", n1)
	fmt.Printf("n2 type: %T\n", n2)

	f := fun1()
	sum := f(1, 2)
	fmt.Println(sum)

	// RMB是int类型的别名，可以进行比较
	var r1 RMB = 10086
	var r2 int = 10086
	fmt.Printf("r1 type: %T\n", r1)
	fmt.Printf("r2 type: %T\n", r2)
	fmt.Println(r1 == r2)
}
