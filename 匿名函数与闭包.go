package main

import (
	"fmt"
)

/*
闭包就是一个函数“捕获”了和它在同一作用域的其它常量和变量。
它不关心这些捕获了的变量或常量是否已经草出了作用域，
所以只有闭包还在使用它，这些变量就还会存在
 */

func main(){
	a:=10
	str:="mike"
	//匿名函数没有函数名字
	f1:=func (){
		//闭包是以引用方式捕获外部变量
		a=6
		fmt.Println("a =",a)
		fmt.Println("str =",str)
	}
	f1() //调用

	//定义匿名函数，同时调用
	func(){

		fmt.Println("a =",a)
		fmt.Println("str =",str)
	}() //后面的（）代表调用此匿名函数

	//匿名函数，有参有返回值
	x,y:=func(i,j int)(max,min int){
		if i>j{
			max=i
			min=j
		}else{
			max=j
			min=i
		}
		return
	}(10,20)
	fmt.Printf("x=%d,y=%d\n",x,y)
}