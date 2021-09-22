package main

import (
	"gorm.io/gorm"
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
