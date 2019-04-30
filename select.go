package main

import "fmt"

//select 默认是是阻塞的，只有当监听的channel中有发送或接收可以进行时才会进行，
// 当多个channel都准备好的时候，select是随机的选择一个执行的
func fibonacci(c,quit chan int){
	x,y:=0,1
	for{
		select {
		case c<-x:
			x,y=y,x+y
			case <-quit:
				fmt.Println("quit")
				return
		}
	}
}
func main(){
	c:=make (chan int)
	quit:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-c)
		}
		quit<-0
	}()
	fibonacci(c,quit)
}
