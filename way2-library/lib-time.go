package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("当前时间:", now)

	// 获取指定时间
	date := time.Date(2008, 7, 15, 16, 30, 28, 0, time.Local)
	fmt.Println("指定时间:", date)

	// 时间格式化:time->string,格式化模板为：2006/1/2 3/4/5,该时间为go诞生的时间
	nowFmt := now.Format("2006-1-2 15:04:05")
	fmt.Println("格式化时间:", nowFmt)

	// string->time:需要解析模板和字符串格式一致
	str := "2021-8-24 18:30:37"
	strTime, err := time.Parse("2006-1-2 15:04:05", str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("字符串转时间:", strTime)

	// 获取当前时间的年,月,日
	year, month, day := now.Date()
	fmt.Println("年/月/日", year, month, day)
	// 获取当前时间的时,分,秒
	clock, min, sec := now.Clock()
	fmt.Println("时/分/秒", clock, min, sec)

	// 获取当前时间的时间单元
	now.Year()    // 年
	now.YearDay() // 当年过了多少天
	now.Month()   // 月
	now.Weekday() // 周
	now.Day()     // 日
	now.Hour()    // 小时
	now.Minute()  // 分钟
	now.Second()  // 秒

	// 时间戳
	unix := now.Unix()     // 秒
	nano := now.UnixNano() // 纳秒
	fmt.Println("时间戳(秒): ", unix)
	fmt.Println("时间戳(纳秒): ", nano)

	// 时间计算
	addNow := now.Add(time.Hour * 24)
	fmt.Println("增加时间: ", addNow)
	addDate := now.AddDate(1, 1, 0)
	fmt.Println("增加日期: ", addDate)
	subTime := now.Sub(addNow)
	fmt.Println("时间的差值: ", subTime)

	// 时间睡眠
	time.Sleep(time.Second * 3)
	fmt.Println("程序睡眠了: ", time.Second*3)
}
