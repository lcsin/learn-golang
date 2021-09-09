/*
strings：字符串操作
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello,world!"
	// strings.Contains: 判断是否包含一个子串
	fmt.Println("string 'hello,world!' contains 'o': ", strings.Contains(s1, "o"))
	// strings.ContainsAny: 判断是否包含chars中的任意一个
	fmt.Println("string 'hello,world!' containsAny 'abh': ", strings.ContainsAny(s1, "abh"))
	// strings.Count: 统计一个子串出现的次数
	fmt.Println("string 'hello,world!' count 'l': ", strings.Count(s1, "l"))
	// strings.Index: 查找子串在字符串中首次出现的下标
	fmt.Println("substring 'llo' first index in 'hello,world!': ", strings.Index(s1, "llo"))
	// strings.IndexAny: 查找子串中任意字符在字符串中首次出现的位置
	fmt.Println("substring 'xyl' first IndexAny in 'hello,world!': ", strings.IndexAny(s1, "xyl"))
	// strings.LastIndex: 查找子串在字符串中最后一次出现的位置
	fmt.Println("substring 'l' last index in 'hello,world!': ", strings.LastIndex(s1, "l"))
	// strings.Split: 字符串切割
	fmt.Println("split string 'hello,world!' with ',': ", strings.Split(s1, ","))
	// strings.Repeat: 字符串重复
	fmt.Println("repeat twice string 'hello,world!': ", strings.Repeat(s1, 2))
	// strings.Replace: 字符串替换(指定替换次数)
	fmt.Println("replace twice 'l' to '*' with string 'hello,world!': ", strings.Replace(s1, "l", "*", 2))
	// strings.ReplaceAll: 字符串替换(全部)
	fmt.Println("replace all 'l' to '*' with string 'hello,world!': ", strings.ReplaceAll(s1, "l", "*"))

	// 11. strings.ToUpper: 字符串中所有字母转换为大写
	fmt.Println("string 'hello,world!' ToUpper: ", strings.ToUpper("hello,world!"))
	// 12. strings.ToLower: 字符串中所有字母转换为小写
	fmt.Println("string 'Hello,World!' ToLower: ", strings.ToLower("Hello,World!"))

	s2 := "2021.8.31.txt"
	// strings.HasPrefix: 判断一个字符串开头
	fmt.Println("2021.8.31.txt has prefix '2021': ", strings.HasPrefix(s2, "2021"))
	// strings.HasSuffix: 判断字符串结尾
	fmt.Println("2021.8.31.txt has suffix '.txt': ", strings.HasSuffix(s2, ".txt"))

	// strings.join: 字符串拼接
	strSlice := []string{"abc", "def", "hyz"}
	ch := "*"
	fmt.Println("join strings 'abc, def, hyz' with '*': ", strings.Join(strSlice, ch))

	// substring
	s3 := s1[0:5]
	fmt.Println("slice implement substring s1[0:5]: ", s3)
}
