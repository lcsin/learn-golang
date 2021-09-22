package main

import (
	"fmt"
	"gorm.io/gorm"
)

// PreloadRelationQuery 预加载查询
func PreloadRelationQuery(db *gorm.DB) {
	var cls []Class
	db.Preload("Student").Find(&cls)

	for _, c := range cls {
		for _, student := range c.Student {
			fmt.Println("1.class name:", c.Name, "student name:", student.Name)
		}
	}

	var d []Dog
	db.Preload("Girl").Find(&d)
	for _, dog := range d {
		fmt.Println("2.dog name:", dog.Name, "girl name:", dog.Girl.Name)
	}

	var b []Boy
	db.Preload("Cat").Find(&b)
	for _, boy := range b {
		fmt.Println("3.boy name:", boy.Name, "cat name:", boy.Cat.Name)
	}
}

// OptionPreloadQuery 带条件的预加载
func OptionPreloadQuery(db *gorm.DB) {
	var class []Class

	// 内联条件
	db.Preload("Student", "name <> ?", "赵四").Find(&class)
	for _, c := range class {
		for _, stu := range c.Student {
			fmt.Println("1. student name:", stu.Name, "class name:", c.Name)
		}
	}

	// 自定义预加载
	db.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Where("name Like ?", "%四%")
	}).Find(&class)
	for _, c := range class {
		for _, stu := range c.Student {
			fmt.Println("2. student name:", stu.Name, "class name:", c.Name)
		}
	}
}
