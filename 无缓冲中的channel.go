package main

import (
	"fmt"
	"time"
)

func main(){
	//创建一个无缓冲的channel
	ch:=make(chan int,0)
	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n",len(ch),cap(ch))

	go func(){
		for i:=0;i<3;i++{

			fmt.Println("子协程 i= ",i)
			ch<-i//往chan写内容
			//fmt.Printf("len(ch)=%d,cap(ch)=%d\n",len(ch),cap(ch))
			//time.Sleep(time.Second)
		}
	}()

	time.Sleep(2*time.Second)

	for i:=0;i<3;i++{
		num:=<-ch    //读取管道中内容，没有内容前阻塞
		fmt.Println("main i= ",num)

	}
}