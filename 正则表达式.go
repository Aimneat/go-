package main

import (
	"fmt"
	"regexp"
)

func main(){
	buf:="abc azc a7c aac 888 a9c tac"
	//1)解释规则，它会解析正则表达式，如果成功 返回解释器
	//reg1:=regexp.MustCompile("a.c")
	// reg1:=regexp.MustCompile("a[0-9]c")
	reg1 := regexp.MustCompile(`a\dc`)  //注意不是单引号
	if reg1==nil{
		fmt.Println("regexp err")
		return
	}
	//2）根据规则提取关键信息
	result1:=reg1.FindAllStringSubmatch(buf,-1) //-1表示提取所有匹配的个数
	fmt.Println("result1=",result1)


	buf1:="43.12 567 adad  1.23 7. 8.9 adfeweweq 6.66 7.8    "
	//1)解释规则，它会解析正则表达式，+匹配前一个字符的1次或多次
	reg2:=regexp.MustCompile(`\d+\.\d+`)
	if reg2==nil{
		fmt.Println("MustCompile err")
		return
	}
	//2）根据规则提取关键信息
	result2:=reg2.FindAllStringSubmatch(buf1,-1) //-1表示提取所有匹配的个数
	fmt.Println("result1=",result2)
}
