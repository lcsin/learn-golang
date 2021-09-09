package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 3.14
	getValue := reflect.ValueOf(f)
	getType := reflect.TypeOf(f)
	fmt.Println("f value:", getValue)
	fmt.Println("f type:", getType)

	var person struct {
		Name string
		Age  int
		Sex  string
	}
	personValue := reflect.ValueOf(person)
	personType := reflect.TypeOf(person)
	fmt.Println("person value:", personValue)
	fmt.Println("person type:", personType)
}
