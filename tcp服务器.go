package main

import (
	"fmt"
	"net"
)

func main(){
	//监听
	listener,err:=net.Listen("tcp",":8000")
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	defer listener.Close()

	//阻塞等待用户链接
	conn,err2:=listener.Accept()
	if err2!=nil{
		fmt.Println("err=",err2)
		return
	}

	//接收用户的请求
	buf:=make([]byte,1024)  //1024大小的缓冲区
	n,err3:=conn.Read(buf)
	if err3!=nil{
		fmt.Println("err=",err3)
		return
	}
	fmt.Println("buf=",string(buf[:n]))

	defer conn.Close()  //关闭当前用户链接
}
