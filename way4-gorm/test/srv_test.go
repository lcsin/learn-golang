package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func InitDB(dsn string) *gorm.DB {
	// 连接 MySQL 数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type CmtSystemMsg struct {
	ID              uint      `json:"id" gorm:"primaryKey;type:bigint(20);comment:用户id"`
	Title           string    `json:"title" gorm:"type:varchar(255);comment:消息标题"`                          //
	Content         string    `json:"content" gorm:"type:text;comment:消息内容"`                                //
	Image           string    `json:"image" gorm:"type:varchar(255);comment:消息图片"`                          //
	UserId          string    `json:"userId" gorm:"type:bigint;comment:用户id 0-全部 ，其他指特定用户"`                 //
	MsgType         string    `json:"msgType" gorm:"type:tinyint;comment:消息类型 0-普通消息 1-banner消息"`           //
	TargetType      string    `json:"targetType" gorm:"type:tinyint;comment:跳转类型，0 不跳转，1 图片，2-H5，3-app内页面"` //
	Redirect        string    `json:"redirect" gorm:"type:varchar(255);comment:跳转url或者app shema"`           //
	IntentUri       string    `json:"intentUri" gorm:"type:varchar(255);comment:APP内部跳转地址"`                 //
	ShowImages      string    `json:"showImages" gorm:"type:text;comment:活动图片，json数组[]"`                    //
	ShowStatus      string    `json:"showStatus" gorm:"type:tinyint;comment:是否显示 0-不显示 1-显示"`               //
	Push            string    `json:"push" gorm:"type:tinyint;comment:是否推送 0-不推送 1-推送"`                     //
	PushStatus      string    `json:"pushStatus" gorm:"type:tinyint;comment:推送状态 0-未推送 1-已推送 2-已撤回"`        //
	PushImmediately string    `json:"pushImmediately" gorm:"type:tinyint;comment:是否立即推送 0-按时推送 1-立即推送"`     //
	PushReserveTime time.Time `json:"pushReserveTime" gorm:"type:timestamp;comment:推送预约时间,非立即推送，此字段才有效"`    //
	PushTime        time.Time `json:"pushTime" gorm:"type:timestamp;comment:发送推送的实际时间"`                     //
	PushRetractTime time.Time `json:"pushRetractTime" gorm:"type:timestamp;comment:推送撤销时间"`                 //
	PushTaskid      string    `json:"pushTaskid" gorm:"type:varchar(50);comment:个推返回的taskid"`               //
	ValidTime       time.Time `json:"validTime" gorm:"type:timestamp;comment:消息有效时间 ，为null表示一直展示"`          //
	Deleted         string    `json:"deleted" gorm:"type:tinyint;comment:是否删除 0 有效 1 无效"`                   //
	CreateTime      time.Time `json:"createTime" gorm:"type:timestamp;comment:创建时间"`                        //
	ModifyTime      time.Time `json:"modifyTime" gorm:"type:timestamp;comment:修改时间"`                        //
	Creator         string    `json:"creator" gorm:"type:varchar(64);comment:创建人"`                          //
	Modifier        string    `json:"modifier" gorm:"type:varchar(64);comment:修改人"`                         //

	Rewards []CmtMsgReward `json:"rewards" gorm:"foreignKey:msg_id"`
}

func (CmtSystemMsg) TableName() string {
	return "cmt_system_msg"
}

type CmtMsgReward struct {
	ID         uint      `json:"id" gorm:"primaryKey;type:bigint(20);comment:主键"`
	MsgID      uint      `json:"msg_id" gorm:"type:bigint(20);comment:消息id"`
	RewardType string    `json:"reward_type" gorm:"type:tinyint;comment:奖励类型 1-活跃度 2-经验  3-金币 4-钻石 5-装扮"`
	RewardNum  uint      `json:"reward_num" gorm:"type:int;comment:奖励数量"`
	RewardIcon string    `json:"reward_icon" gorm:"type:varchar;comment:奖励图标"`
	RelationId uint      `json:"relation_id" gorm:"type:bigint(20);comment:关联奖励物品id"`
	ValidTime  int       `json:"valid_time" gorm:"type:int;comment:有效时间,单位天"`
	CreateTime time.Time `json:"create_time" gorm:"type:timestamp;comment:创建时间"`
	Deleted    uint      `json:"deleted"`
	ModifyTime time.Time `json:"modify_time"`
	Creator    string    `json:"creator"`
	Modifier   string    `json:"modifier"`
}

func (CmtMsgReward) TableName() string {
	return "cmt_msg_reward"
}

func TestSysMessageReward(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/zsxyadmin?charset=utf8mb4&parseTime=True&loc=Local"
	var message []CmtSystemMsg
	InitDB(dsn).Preload("Rewards").Find(&message)
	for _, msg := range message {
		fmt.Println(msg)
	}
}

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

func TestClassStudent(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/learn_db?charset=utf8mb4&parseTime=True&loc=Local"
	var class []Class
	InitDB(dsn).Preload("Student").Find(&class)
	for _, c := range class {
		for _, student := range c.Student {
			fmt.Println("class name:", c.Name, "student name:", student.Name)
		}
	}
}
