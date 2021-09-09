package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello,", msg)
}

func (p Person) Info() {
	fmt.Printf("姓名：%s, 年龄：%d, 性别：%s\n", p.Name, p.Age, p.Sex)
}

func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is:", getType.Name())
	fmt.Println("get Kind is:", getType.Kind())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("字段名称：%s, 字段类型：%s, 字段数值：%v\n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s, 方法类型：%v\n", method.Name, method.Type)
	}
}

func main() {
	// 已知原有类型
	var f float64 = 3.14
	value := reflect.ValueOf(f)
	convertValue := value.Interface().(float64)
	fmt.Printf("convertValue Type:%T, value:%v\n", convertValue, convertValue)

	// 未知原有类型
	p1 := Person{"张三", 28, "男"}
	GetMessage(p1)
}
