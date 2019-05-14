package main

import "fmt"

func testa(){
	fmt.Println("aaaa")
}
func testb(x int){
	//设置recover
	defer func (){
		//recover //可以打印panic的错误信息
		//fmt.Println(recover())
		if err:=recover();err !=nil{
			fmt.Println(err)
		}
	}() //别忘了(),调用此匿名函数
	var a [10]int
	a[x]=111  //当x>9时，导致数组越界，产生panic，导致程序崩溃
}
func testc(){
	fmt.Println("cccc")
}

func main(){
	testa()
	testb(1)  //什么也不打印
	testb(10)
	testc()
}