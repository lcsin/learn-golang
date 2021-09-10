package main

import (
	"fmt"
	"sync"
)

/*
sync.WaitGroup：等待所有的并发任务完成
1. Add(n)：WaitGroup的值+1，表示总共有n个任务要等待
2. Done()：WaitGroup的值-1，表示有一个goroutine完成了任务
3. Wait()：阻塞等待所有的任务完成，直到WaitGroup的值为0
*/
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
