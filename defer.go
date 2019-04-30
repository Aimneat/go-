package main

import "fmt"

//在defer后指定的函数会在函数退出前调用,如果有很多调用defer，那么defer是采用后进先出模式。
//哪怕这个函数或某个延时调用发生错误，这些调用依旧会被执行

func main(){

	//fmt.Printf("hello ")
	//
	//for i:=0;i<10;i++{
	//	defer fmt.Printf("%d ",i)
	//}
	//
	//defer fmt.Printf("world \n")

	a:=10
	b:=20
	defer func(a,b int){
		fmt.Println("a=",a,"b=",b)
	}(a,b)
	a=111
	b=222
	fmt.Println("a=",a,"b=",b)
}