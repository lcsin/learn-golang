package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Create 用数据的指针创建记录
func Create(db *gorm.DB) {
	user := User{
		Name:     "张三",
		Age:      18,
		Birthday: time.Now(),
	}
	// 通过数据的指针来创建记录
	result := db.Create(&user)
	// 返回插入数据的主键
	fmt.Println("user id:", user.ID)
	// 返回error
	fmt.Println("error:", result.Error)
	// 返回插入记录的条数
	fmt.Println("count", result.RowsAffected)
}

// SelectCreate 用指定的字段创建记录
func SelectCreate(db *gorm.DB) {
	user := User{
		Name:     "张三",
		Age:      18,
		Birthday: time.Now(),
	}

	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	result := db.Select("Name", "Age").Create(&user)

	fmt.Println("user id:", user.ID)
	fmt.Println("error:", result.Error)
	fmt.Println("count", result.RowsAffected)
}

// OmitCreate 忽略指定的字段创建记录
func OmitCreate(db *gorm.DB) {
	user := User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "李四",
		Age:      20,
		Birthday: time.Now(),
	}

	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
	result := db.Omit("Name", "Age", "CreatedAt").Create(&user)

	fmt.Println("user id:", user.ID)
	fmt.Println("error:", result.Error)
	fmt.Println("count:", result.RowsAffected)
}

// BatchCreate 批量插入
func BatchCreate(db *gorm.DB) {
	users := []User{
		{Name: "user1", Birthday: time.Now()},
		{Name: "user2", Birthday: time.Now()},
		{Name: "user3", Birthday: time.Now()},
	}
	result := db.Create(&users)

	fmt.Println("error:", result.Error)
	fmt.Println("count:", result.RowsAffected)
	for _, user := range users {
		fmt.Println("user id:", user.ID)
	}
}

// CreateInBatches 批量创建
func CreateInBatches(db *gorm.DB, num int) {
	users := []User{
		{Name: "user14", Birthday: time.Now()},
		{Name: "user15", Birthday: time.Now()},
		{Name: "user16", Birthday: time.Now()},
		{Name: "user17", Birthday: time.Now()},
		{Name: "user18", Birthday: time.Now()},
		{Name: "user19", Birthday: time.Now()},
	}

	results := db.CreateInBatches(users, num)
	fmt.Println("error:", results.Error)
	fmt.Println("count:", results.RowsAffected)
	for _, user := range users {
		fmt.Println("user id:", user.ID)
	}
}

// BeforeCreate 创建钩子
func (u User) BeforeCreate(db *gorm.DB) (err error) {
	fmt.Println("Create Hooks...")
	return
}

// SkipHooks 使用SkipHooks跳过钩子方法
func SkipHooks(db *gorm.DB) {
	user := User{
		Name:     "user001",
		Age:      10086,
		Birthday: time.Now(),
	}
	result := db.Session(&gorm.Session{SkipHooks: true}).Create(&user)

	fmt.Println("user id:", user.ID)
	fmt.Println("error:", result.Error)
	fmt.Println("count:", result.RowsAffected)
}

// MapCreate 根据Map创建
func MapCreate(db *gorm.DB) {
	users := []map[string]interface{}{
		{"Name": "小坂菜绪", "Age": 19},
		{"Name": "金村美玖", "Age": 19},
		{"Name": "滨岸hiyori", "Age": 19},
	}
	result := db.Model(&User{}).Create(users)

	fmt.Println("error:", result.Error)
	fmt.Println("count:", result.RowsAffected)
}

// RelationCreate 关联创建，会同时创建两张表数据
func RelationCreate(db *gorm.DB) {
	user := User{
		Name:     "松田好花",
		Age:      20,
		Birthday: time.Now(),
		CreditCard: CreditCard{
			Number: "10000018612",
		},
	}
	db.Create(&user)
}

// DefaultCreate 根据default标签默认值创建
func DefaultCreate(db *gorm.DB) {
	user := User{
		Age:      22,
		Birthday: time.Now(),
	}
	db.Create(&user)
}

func main() {
	// 连接 MySQL 数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Create(db)
	//SelectCreate(db)
	//OmitCreate(db)
	//BatchCreate(db)
	//CreateInBatches(db, 3)
	//SkipHooks(db)
	//MapCreate(db)
	//RelationCreate(db)
	DefaultCreate(db)
}
