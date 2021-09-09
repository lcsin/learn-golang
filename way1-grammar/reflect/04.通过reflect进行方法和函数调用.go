package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

func (p Student) Say(msg string) {
	fmt.Println("hello,", msg)
}

func (p Student) Info() {
	fmt.Printf("姓名：%s, 年龄：%d, 性别：%s\n", p.Name, p.Age, p.Sex)
}

func func1() {
	fmt.Println("hello,world!")
}

func func2(msg1, msg2 string) {
	fmt.Println(msg1, msg2)
}

func func3() int {
	return 10086
}

/*
通过reflect调用方法method的主要步骤：
1. 通过reflect.ValueOf()获取结构体reflect对象
2. 通过MethodByName()获取要调用的方法
3. 使用Call()函数进行调用

通过reflect调用函数func的主要步骤：
1. 通过reflect.ValueOf()获取函数的reflect对象
2. 通过Call()函数进行调用

Call()函数：
1. 参数和返回值均为[]reflect.Value切片
2. 参数为空时，可以传递nil
3. Call()方法最终调用真实的方法，所以参数需要保持一致
*/
func main() {
	// 调用method
	stu := Student{"张三", 18, "男"}
	value := reflect.ValueOf(stu)

	method1 := value.MethodByName("Say")
	fmt.Printf("method1 kind: %s, type: %s\n", method1.Kind(), method1.Type())
	args1 := []reflect.Value{reflect.ValueOf("world!")}
	method1.Call(args1)

	method2 := value.MethodByName("Info")
	fmt.Printf("method2 kind: %s, type: %s\n", method2.Kind(), method2.Type())
	method2.Call(nil)

	// 调用函数
	f1 := reflect.ValueOf(func1)

	fmt.Printf("f1 kind: %s, type: %s\n", f1.Kind(), f1.Type())
	f1.Call(nil)

	f2 := reflect.ValueOf(func2)
	fmt.Printf("f2 kind: %s, type: %s\n", f2.Kind(), f2.Type())
	f2.Call([]reflect.Value{reflect.ValueOf("hello,"), reflect.ValueOf("world!")})

	f3 := reflect.ValueOf(func3)
	fmt.Printf("f3 kind: %s, type: %s\n", f3.Kind(), f3.Type())
	result := f3.Call(nil)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
