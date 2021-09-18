/*
has_one关联模式：
1. 属于一对一模式
2. A has_one B 需要具备两个条件：
	2-1. A实体作为主体需要内嵌B实体
	2-2. B实体需要A主体的主键ID作为外键
*/
package main

import "gorm.io/gorm"

type Boy struct {
	gorm.Model

	Name string
	Cat  Cat
}

type Cat struct {
	gorm.Model

	Name  string
	BoyID uint
}
