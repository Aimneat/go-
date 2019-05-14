package main

import (
	"fmt"
	"strings"
)

func main(){
	//“hellogo”中是否包含“hello”，包含返回true，反之为false
	fmt.Println(strings.Contains("hellogo","hello"))
	fmt.Println(strings.Contains("hello","go"))

	//Joins 组合
	s:=[]string{"abc","hello","mike"}
	buf:=strings.Join(s,"x")
	fmt.Println("buf=",buf)

	//Index 查找子串的位置
	fmt.Println(strings.Index("abchellogo","hello"))
	fmt.Println(strings.Contains("abchello","go")) //不包含子串返回-1(后改为false）

	buf=strings.Repeat("go",3)
	fmt.Println("buf=",buf)  //"gogogo"

	//Split 以指定的分隔符拆分
	buf="hello@abc@go"
	s2:=strings.Split(buf,"@")
	fmt.Println("s2=",s2)

	//Trim 去掉两头的字符
	buf =strings.Trim("   are u ok?   "," ")
	fmt.Printf("buf=#%s#\n",buf)

	//去掉空格，把元素放入切片中
	s3:=strings.Fields("   are u ok?   ")
	for i,data:=range s3{
		fmt.Println(i,", ",data)
	}
}