/*
belongs_to关联模式：
1. 属于一对一模式
2. A belongs_to B 需要具备两个条件：
	2-1. A实体作为主体需要内嵌B实体
	2-2. A主体需要B实体主键ID作为外键
*/
package main

import "gorm.io/gorm"

// Dog 属于 Girl
type Dog struct {
	gorm.Model

	Name   string
	GirlID uint
	Girl   Girl
}

type Girl struct {
	gorm.Model

	Name string
}
