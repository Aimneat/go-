package main

import "fmt"

func producer(out chan<-int){  //此通道只能写不能读
	for i:=0;i<10;i++{
		out<-i*i
	}
	close(out)
}
func consumer(in <-chan int){
	for num:=range in{
		fmt.Println("num=",num)
	}
}

func main(){
	ch:=make(chan int)

	//生产者，生产数字，写入channel
	go producer(ch)  //  channel传参，引用传递

	consumer(ch)
}