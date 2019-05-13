package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个channel
var ch=make(chan int)

//定义一个打印机，参数为字符串，按每个字符打印
func printer(str string){
	for _,data:=range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//使person1先执行完
func person1(){
	printer("he")  //打印完成之前是不会给管道写数据的
	ch<-666  //给管道写数据，发送
}
func person2(){
	<-ch  //从管道取数据，接收，如果通道没有数据它就会阻塞
	printer("wo")

}

func main(){
	go person1()
	go person2()

	i:=0
	for i<20000000000{
		//不让主协程先结束的循环
		i++
	}
}