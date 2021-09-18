package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//type User struct {
//	ID   uint
//	Name string
//	Age  uint
//}

type APIUser struct {
	ID   uint
	Name string
}

// IntelligentSelectQuery 智能选择字段 。
func IntelligentSelectQuery(db *gorm.DB) {
	var apiUsers []APIUser

	// 查询时会自动选择 `id`, `name` 字段
	// SELECT `id`, `name` FROM `users` LIMIT 10
	db.Model(&User{}).Limit(10).Find(&apiUsers)
	for _, user := range apiUsers {
		fmt.Println("id:", user.ID, "name:", user.Name)
	}
}

// ChildQuery 子查询
func ChildQuery(db *gorm.DB) {
	var users []User

	// select * from users where age > (SELECT AVG(age) FROM "users");
	db.Where("age > (?)", db.Table("users").Select("AVG(age)")).Find(&users)
	for _, user := range users {
		fmt.Println("1.username:", user.Name, "age:", user.Age)
	}
	type Result struct {
		Name   string
		Avgage float64
	}

	var results []Result
	subQuery := db.Select("AVG(age)").Where("name LIKE ?", "%g%").Table("users")
	db.Table("users").Select("AVG(age) as avgage, name").Group("name").Having("100 > (?)", subQuery).Find(&results)
	for _, result := range results {
		fmt.Println("2.username:", result.Name, "age:", result.Avgage)
	}
}

func main() {
	// 连接 MySQL 数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//IntelligentSelectQuery(db)
	ChildQuery(db)
}
