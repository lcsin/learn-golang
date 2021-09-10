package main

import (
	"fmt"
	"time"
)

func main() {
	// go关键字：用于启动一个goroutine协程
	// main函数也是一个goroutine，叫做main goroutine
	// 当main函数返回后，所有的goroutine都会被中断执行，程序退出
	go f1()
	fmt.Println("main...")
	time.Sleep(1000 * time.Millisecond)
}

func f1() {
	fmt.Println("hello,world!")
}
