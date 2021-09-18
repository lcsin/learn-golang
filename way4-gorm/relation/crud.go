/*
关联模式的CRUD：
1. 对所有关联模式都适用
2. 需要以主体作为model
*/
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RelationAutoMigrate(db *gorm.DB) {
	// belongs_to关联模式下，会创建Dog表和Girl表
	db.AutoMigrate(&Dog{})
	db.AutoMigrate(&Boy{}, &Cat{})
}

// RelationCreate
// 在创建、更新记录时，GORM 会通过 Upsert 自动保存关联及其引用记录。
func RelationCreate(db *gorm.DB) {
	g := Girl{
		Name: "斋藤飞鸟",
	}
	d := Dog{
		Name:   "李四",
		GirlID: 2,
		Girl:   g,
	}
	// 在创建d数据时，也会创建g数据
	db.Create(&d)

	c := Cat{
		Name: "日向雏田",
	}
	b := Boy{
		Name: "漩涡鸣人",
		Cat:  c,
	}
	db.Create(&b)

}

// FindRelation 关联查找
func FindRelation(db *gorm.DB) {
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
		GirlID: 1,
	}
	var g Girl

	// belongs_to模式的关联查找
	db.Model(&d).Association("Girl").Find(&g)
	fmt.Println("girl name:", g.Name)

	b := Boy{
		Model: gorm.Model{
			ID: 1,
		},
	}
	var c Cat

	// has_one模式的关联查找
	db.Model(&b).Association("Cat").Find(&c)
	fmt.Println("cat name:", c.Name)

	cls := Class{
		Model: gorm.Model{
			ID: 1,
		},
	}

	var s []Student
	db.Model(&cls).Association("Student").Find(&s)
	for _, student := range s {
		fmt.Println("student name:", student.Name)
	}
}

// AppendRelation 添加关联
func AppendRelation(db *gorm.DB) {
	s := []Student{
		{Name: "李四", ClassID: 1},
		{Name: "王五", ClassID: 1},
	}
	c := Class{
		Model: gorm.Model{
			ID: 1,
		},
	}

	// 添加关联
	db.Model(&c).Association("Student").Append(s)
}

// ReplaceRelation 替换关联
func ReplaceRelation(db *gorm.DB) {
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
	}
	g := Girl{
		Model: gorm.Model{
			ID: 2,
		},
	}
	db.Model(&d).Association("Girl").Replace(&g)

	b := Boy{
		Model: gorm.Model{
			ID: 1,
		},
	}
	c1 := Cat{
		Model: gorm.Model{
			ID: 1,
		},
	}
	c2 := Cat{
		Model: gorm.Model{
			ID: 2,
		},
	}
	db.Model(&b).Association("Cat").Replace(&c1, &c2)
}

// PreloadRelationQuery 预加载查询
func PreloadRelationQuery(db *gorm.DB) {
	var cls Class
	db.Preload("Student").Find(&cls)

	for _, student := range cls.Student {
		fmt.Println("1.id:", student.ID, "name:", student.Name)
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
	var class Class

	// 内联条件
	db.Preload("Student", "name <> ?", "赵四").Find(&class)
	for _, student := range class.Student {
		fmt.Println("student name:", student.Name)
	}

	// 自定义预加载
	db.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Where("name Like ?", "%四%")
	}).Find(&class)
	for _, student := range class.Student {
		fmt.Println("student name:", student.Name)
	}
}

func main() {
	// 连接 MySQL 数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//RelationAutoMigrate(db)
	//RelationCreate(db)
	//FindRelation(db)
	//AppendRelation(db)
	//ReplaceRelation(db)
	//PreloadRelationQuery(db)
	OptionPreloadQuery(db)
}
