/*
has_many关联模式：
1. 表示一对多关联
2. 在has_one模式的基础上，主体嵌入的是实体的数组
*/
package main

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model

	Name    string
	Student []Student
}

type Student struct {
	gorm.Model

	Name    string
	ClassID uint
}
