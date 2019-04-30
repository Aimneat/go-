package main

import "fmt"

//函数的返回值是一个匿名函数，返回一个函数类型
func test02()func ()int {
	var x int
	return func()int{
		x++
		return x*x
	}
}
func main(){
	f:=test02()
	/*
闭包就是一个函数“捕获”了和它在同一作用域的其它常量和变量。
它不关心这些捕获了的变量或常量是否已经草出了作用域，
所以只有闭包还在使用它，这些变量就还会存在
 */
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}