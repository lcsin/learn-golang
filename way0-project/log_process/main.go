package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Reader 读取日志接口
type Reader interface {
	Reader(rc chan []byte)
}

// Writer 写入日志接口
type Writer interface {
	Writer(wc chan string)
}

type ReadFromFile struct {
	path string
}

// Reader 从文件中读取
func (r *ReadFromFile) Reader(rc chan []byte) {
	// 打开文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	// 从文件末尾逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

type WriteToInfluxDB struct {
	dsn string
}

// Writer 写入解析后的日志信息
func (w *WriteToInfluxDB) Writer(wc chan string) {
	for v := range wc {
		fmt.Println(v)
	}
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Reader
	write Writer
}

// Process 解析读取的日志信息
func (l *LogProcess) Process() {
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

func main() {
	r := &ReadFromFile{path: "./access.log"}
	w := &WriteToInfluxDB{dsn: "username&password.."}
	lg := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go lg.read.Reader(lg.rc)
	go lg.Process()
	go lg.write.Writer(lg.wc)
	time.Sleep(time.Second * 30)
}
