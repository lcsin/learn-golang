package main

import (
	"fmt"
	"reflect"
)

/*
1. 通过reflect对象修改实际变量的值需要通过指针去修改，所以reflect.ValueOf()需要传递指针参数
2. 如果传递的不是指针参数，Elem()将会报错，且CanSet()为false
*/
func main() {
	// 修改基本数据类型的值
	var f float64 = 3.14
	value := reflect.ValueOf(&f)
	if value.Kind() == reflect.Ptr {
		elem := value.Elem()
		if elem.CanSet() && elem.Kind() == reflect.Float64 {
			elem.SetFloat(3.1415967)
			fmt.Println("f new value:", elem)
			fmt.Println("f value:", f)
		}
	}

	// 修改结构体类型的值
	type user struct {
		Name string
		Age  int
		Sex  string
	}
	u := user{"张三", 19, "男"}
	value = reflect.ValueOf(&u)
	if value.Kind() == reflect.Ptr {
		elem := value.Elem()
		if elem.CanSet() {
			name := elem.FieldByName("Name")
			if name.Kind() == reflect.String {
				name.SetString("李四")
			}
			age := elem.FieldByName("Age")
			if age.Kind() == reflect.Int {
				age.SetInt(29)
			}
			sex := elem.FieldByName("Sex")
			if sex.Kind() == reflect.String {
				sex.SetString("女")
			}
			fmt.Println("u new value:", elem)
			fmt.Println("u value:", u)
		}
	}
}
