package main

import (
	"fmt"
)

func main(){
//创建一个无缓冲的channel
ch:=make(chan int)

go func(){
	for i:=0;i<5;i++{
		ch<-i
}
	close(ch)  //不需要在写数据时，关闭channel

	//ch<-666 //关闭channel后无法再发送数据
}()


//for{
//	//如果ok为true，说明管道没有关闭
//	if num,ok:=<-ch;ok==true{
//		fmt.Println("num=",num)
//	}else{
//		break
//	}
//
//	}

	for num:=range ch{
		fmt.Println("num=",num)
	}
}
