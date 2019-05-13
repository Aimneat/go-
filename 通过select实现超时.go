package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int)  //数字通信
	quit:=make(chan bool)//程序是否结束

	//消费者，从channel中读取内容
	go func (){
		for {
			select {
				case num:=<-ch:
					fmt.Println("num=",num)
				case <-time.After(2*time.Second):
					fmt.Println("超时")
					quit<-true
			}
		}
	}()
	for i:=0;i<3;i++{
		ch<-i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}