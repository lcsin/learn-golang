package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type User struct {
	gorm.Model

	Name       string `gorm:"default:guest"`
	Age        int
	Birthday   time.Time
	CompanyId  uint
	CreditCard CreditCard
}

type CreditCard struct {
	ID       uint
	Number   string
	UserId   uint
	Username string
}

// QueryOne 检索单个对象
func QueryOne(db *gorm.DB) {
	var user User
	// 根据主键升序获取第一条记录
	// select * from users order by id limit 1;
	result := db.First(&user)
	fmt.Printf("u1 age:%s, id:%d\n", user.Name, user.ID)

	// 获取一条记录
	// select * from users limit 1;
	db.Take(&user)
	fmt.Printf("u2 name:%s, id:%d\n", user.Name, user.ID)

	// 根据主键降序获取最后一条记录
	// select * from users order by id desc limit 1;
	db.Last(&user)
	fmt.Printf("u3 id:%s, id:%d\n", user.Name, user.ID)

	// 检查 ErrRecordNotFound 错误
	fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
}

// FirstAndLastQuery
// 查询数据库时添加了 LIMIT 1 条件，且没有找到记录时，会返回 ErrRecordNotFound 错误
// First 和 Last 只有在目标 struct 是指针或者通过 db.Model() 指定 model 时，该方法才有效。
func FirstAndLastQuery(db *gorm.DB) {
	var user User

	// 有效，因为目标 struct 是指针
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	db.First(&user)

	// 有效，因为通过 `db.Model()` 指定了 model
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result)

	// 无效
	result = map[string]interface{}{}
	db.Table("users").First(&result)

	// 配合 Take 有效
	result = map[string]interface{}{}
	db.Table("users").Take(&result)

	// 未指定主键，会根据第一个字段排序(即：`Code`)
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
	type Language struct {
		Code string
		Name string
	}
	db.First(&Language{})
}

// QueryByID 用主键检索
func QueryByID(db *gorm.DB) {
	var user User
	var users []User

	// SELECT * FROM users WHERE id = 27;
	db.First(&user, 27)
	fmt.Println("user name:", user.Name)

	// SELECT * FROM users WHERE id = 10;
	db.First(&user, "27")
	fmt.Println("user name:", user.Name)

	// SELECT * FROM users WHERE id IN (1,2,3);
	db.Find(&users, []int{22, 23, 24})
	for _, u := range users {
		fmt.Println("user name:", u.Name)
	}

	// 如果主键是字符串（例如像 uuid），查询将被写成这样：
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
	//db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
}

// QueryAll 检索全部对象
func QueryAll(db *gorm.DB) {
	var users []User

	// 获取全部记录
	// SELECT * FROM users;
	result := db.Find(&users)

	fmt.Println("count:", result.RowsAffected)
	fmt.Println("error:", result.Error)
}

// OptionQueryWithString String条件查询
func OptionQueryWithString(db *gorm.DB) {
	var user User
	var users []User

	// 获取第一条匹配的记录
	// SELECT * FROM users WHERE name = '小坂菜绪' ORDER BY id LIMIT 1;
	db.Where("name = ?", "小坂菜绪").First(&user)
	fmt.Println("username:", user.Name)

	// 获取全部匹配的记录
	// SELECT * FROM users WHERE name <> '小坂菜绪';
	db.Where("name <> ?", "小坂菜绪").Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}

	// IN
	// SELECT * FROM users WHERE name IN ('小坂菜绪','金村美玖');
	db.Where("name IN ?", []string{"小坂菜绪", "金村美玖"}).Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}

	// LIKE
	// SELECT * FROM users WHERE name LIKE '%美%';
	db.Where("name LIKE ?", "%美%").Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}

	// AND
	// SELECT * FROM users WHERE name = '松田好花' AND age >= 18;
	db.Where("name = ? AND age >= ?", "松田好花", "18").Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}

	// Time
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	db.Where("updated_at > ?", "2000-01-01 00:00:00").Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}

	// BETWEEN
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	db.Where("created_at BETWEEN ? AND ?", "2000-01-01 00:00:00", "2022-01-01 00:00:00").Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}
}

// OptionQueryWithStructAndMap 结构条件查询
// 当使用结构作为条件查询时，GORM 只会查询非零值字段
// 如果想要包含零值查询条件，你可以使用 map，其会包含所有 key-value 的查询条件
func OptionQueryWithStructAndMap(db *gorm.DB) {
	var user User
	var users []User

	// Struct
	// SELECT * FROM users WHERE name = "松田好花" AND age = 20 ORDER BY id LIMIT 1;
	db.Where(&User{Name: "松田好花", Age: 20}).First(&user)
	fmt.Println("username:", user.Name)

	// Map
	// SELECT * FROM users WHERE name = "松田好花" AND age = 20;
	db.Where(map[string]interface{}{"name": "松田好花", "age": 20}).Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}

	// 主键切片条件
	// SELECT * FROM users WHERE id IN (20, 21, 22);
	db.Where([]int64{20, 21, 22}).Find(&users)
	for _, u := range users {
		fmt.Println("username:", u.Name)
	}
}

// InternalRelationQuery 内联条件查询
// 查询条件也可以被内联到 First 和 Find 之类的方法中，其用法类似于 Where。
func InternalRelationQuery(db *gorm.DB) {
	var user User
	var users []User

	// 根据主键获取记录，如果是非整型主键
	// SELECT * FROM users WHERE id = '22';
	db.First(&user, "id = ?", "22")
	fmt.Println("1.username:", user.Name)

	// Plain SQL
	// SELECT * FROM users WHERE name = "金村美玖";
	db.Find(&user, "name = ?", "金村美玖")
	fmt.Println("2.username:", user.Name)

	// SELECT * FROM users WHERE name <> "松田好花" AND age > 20;
	db.Find(&users, "name <> ? AND age > ?", "松田好花", 20)
	for _, u := range users {
		fmt.Println("3.username:", u.Name)
	}

	// Struct
	// SELECT * FROM users WHERE age = 20;
	db.Find(&users, User{Age: 20})
	for _, u := range users {
		fmt.Println("4.username:", u.Name)
	}

	// Map
	// SELECT * FROM users WHERE age = 20;
	db.Find(&users, map[string]interface{}{"age": 20})
	for _, u := range users {
		fmt.Println("5.username:", u.Name)
	}
}

// NotOptionQuery not 查询
func NotOptionQuery(db *gorm.DB) {
	var user User
	var users []User

	// SELECT * FROM users WHERE NOT name = "松田好花" ORDER BY id LIMIT 1;
	db.Not("name = ?", "松田好花").First(&user)
	fmt.Println("1.username:", user.Name)

	// Not In
	// SELECT * FROM users WHERE name NOT IN ("松田好花", "金村美玖");
	db.Not(map[string]interface{}{"name": []string{"松田好花", "金村美玖"}}).Find(&users)
	for _, u := range users {
		fmt.Println("2.username:", u.Name)
	}

	// Struct
	// SELECT * FROM users WHERE name <> "松田好花" AND age <> 18 ORDER BY id LIMIT 1;
	db.Not(User{Name: "松田好花", Age: 18}).First(&user)
	fmt.Println("3.username:", user.Name)

	// 不在主键切片中的记录
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
	db.Not([]int64{1, 2, 3}).First(&user)
	fmt.Println("4.username:", user.Name)
}

// OrOptionQuery Or条件查询
func OrOptionQuery(db *gorm.DB) {
	var users []User

	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
	db.Where("name = ?", "松田好花").Or("name = ?", "小坂菜绪").Find(&users)
	for _, u := range users {
		fmt.Println("1.username:", u.Name)
	}

	// Struct
	// SELECT * FROM users WHERE name = '松田好花' OR (name = '小坂菜绪' AND age = 18);
	db.Where("name = '松田好花'").Or(User{Name: "小坂菜绪", Age: 18}).Find(&users)
	for _, u := range users {
		fmt.Println("2.username:", u.Name)
	}

	// Map
	// SELECT * FROM users WHERE name = '松田好花' OR (name = '金村美玖' AND age = 18);
	db.Where("name = '松田好花'").Or(map[string]interface{}{"name": "金村美玖", "age": 18}).Find(&users)
	for _, u := range users {
		fmt.Println("3.username:", u.Name)
	}
}

// SelectQuery 选择特定字段查询
func SelectQuery(db *gorm.DB) {
	var users []User

	// SELECT name, age FROM users;
	db.Select("name", "age").Find(&users)
	for _, user := range users {
		fmt.Printf("1.username:%s, age:%d\n", user.Name, user.Age)
	}

	// SELECT name, age FROM users;
	db.Select([]string{"name", "age"}).Find(&users)
	for _, user := range users {
		fmt.Printf("2.username:%s, age:%d\n", user.Name, user.Age)
	}

	// SELECT COALESCE(age,'42') FROM users;
	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	for _, user := range users {
		fmt.Printf("3.username:%s, age:%d\n", user.Name, user.Age)
	}
}

// OrderQuery 排序查询
func OrderQuery(db *gorm.DB) {
	var users []User

	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("age desc, name").Find(&users)
	for _, user := range users {
		fmt.Printf("1.username:%s, age:%d\n", user.Name, user.Age)
	}

	// 多个 order
	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("age desc").Order("name").Find(&users)
	for _, user := range users {
		fmt.Printf("2.username:%s, age:%d\n", user.Name, user.Age)
	}

	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&users)
	for _, user := range users {
		fmt.Printf("3.username:%s, age:%d\n", user.Name, user.Age)
	}
}

// LimitAndOffsetQuery
// Limit 指定获取记录的最大数量 Offset 指定在开始返回记录之前要跳过的记录数量
func LimitAndOffsetQuery(db *gorm.DB) {
	var users []User

	// SELECT * FROM users LIMIT 3;
	db.Limit(5).Find(&users)
	for _, user := range users {
		fmt.Println("1.username:", user.Name)
	}

	// 通过 -1 消除 Limit 条件
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
	db.Limit(3).Find(&users).Limit(-1).Find(&users)
	for _, user := range users {
		fmt.Println("2.username:", user.Name)
	}

	// SELECT * FROM users OFFSET 3;
	db.Limit(10).Offset(3).Find(&users)
	for _, user := range users {
		fmt.Println("3.username:", user.Name)
	}

	// SELECT * FROM users OFFSET 5 LIMIT 10;
	db.Limit(10).Offset(5).Find(&users)
	for _, user := range users {
		fmt.Println("4.username:", user.Name)
	}

	// 通过 -1 消除 Offset 条件
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)
	db.Limit(10).Offset(2).Find(&users).Offset(-1).Find(&users)
	for _, user := range users {
		fmt.Println("5.username:", user.Name)
	}
}

type Result struct {
	Age   uint
	Total int
	Name  string
}

// GroupAndHavingQuery Group和Having查询
func GroupAndHavingQuery(db *gorm.DB) {
	var results []Result
	var result Result

	// SELECT age, sum(age) as total FROM `users` GROUP BY `age`
	db.Model(&User{}).Select("age, sum(age) as total").Group("age").Find(&results)
	for _, result := range results {
		fmt.Println("age:", result.Age, "total:", result.Total)
	}

	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "guest"
	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "guest").Find(&result)
	fmt.Println("name:", result.Name, "total:", result.Total)
}

// DistinctQuery 去重查询
func DistinctQuery(db *gorm.DB) {
	var users []User
	db.Distinct("name", "age").Order("name, age desc").Find(&users)
	for _, user := range users {
		fmt.Println("username:", user.Name, "age:", user.Age)
	}
}

type CardResult struct {
	Name   string
	Number string
}
type Company struct {
	ID   uint
	Name string
}

// JoinQuery 联表查询
func JoinQuery(db *gorm.DB) {
	var results []CardResult
	var users []User

	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	db.Model(&User{}).Select("users.name, credit_cards.number").Joins("left join credit_cards on credit_cards.user_id = users.id").Scan(&results)
	for _, result := range results {
		fmt.Println("1.username:", result.Name, "card number:", result.Number)
	}

	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
	db.Joins("companies").Find(&users)
	for _, user := range users {
		fmt.Println("2.username:", user.Name, "companyId:", user.CompanyId)
	}
}

func main() {
	// 连接 MySQL 数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//QueryOne(db)
	//QueryByID(db)
	//QueryAll(db)
	//OptionQueryWithString(db)
	//OptionQueryWithStructAndMap(db)
	//InternalRelationQuery(db)
	//NotOptionQuery(db)
	//OrOptionQuery(db)
	//SelectQuery(db)
	//OrderQuery(db)
	//LimitAndOffsetQuery(db)
	//GroupAndHavingQuery(db)
	//DistinctQuery(db)
	//JoinQuery(db)
	var results []CreditCard
	db.Model(&CreditCard{}).Select("credit_cards.id, number, users.name username, user_id").Joins("left join users on user_id = users.id").Find(&results)
	for _, r := range results {
		fmt.Println("id:", r.ID, "number:", r.Number, "username:", r.Username, "userId:", r.UserId)
	}
}
