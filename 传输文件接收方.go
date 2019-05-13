package main

import (

	"fmt"
	"io"
	"net"
	"os"
)

func recvfile(fileName string,conn net.Conn) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}

	buf := make([]byte, 1024*4)

	//读文件内容，读多少发多少
	for {
		n, err := conn.Read(buf) //通过互联网接收对方发过来的文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err=", err)
			}
			return
		}
		f.Write(buf[:n]) //往文件写入内容
	}
}
func main(){

		//监听
		listenner,err:=net.Listen("tcp",":8000")
		if err != nil {
			fmt.Println("net.Listen err=", err)
			return
		}
		defer listenner.Close()

		//阻塞等待用户连接
		conn,err1:=listenner.Accept()
		if err1 != nil {
			fmt.Println("listenner.Accept err=", err1)
			return
		}

		defer conn.Close()

		buf:=make([]byte,1024)
		var n int
		n,err=conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fileName:=string(buf[:n])

		//回复“ok”
		conn.Write([]byte("ok"))

		//接收文件内容
		recvfile(fileName,conn)
}
