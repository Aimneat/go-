package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConn(conn net.Conn){
	defer conn.Close()
	//获取客户端的网络地址信息
	addr:=conn.RemoteAddr().String()
	fmt.Println("addr connect sucessful")

	buf:=make([]byte,2048)  //2048大小的缓冲区

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Printf("[%s]:%s\n",addr,string(buf[:n]))

		if "exit"==string(buf[:(n-1)]){
			fmt.Println(addr,"exit")
			return
		}

		//把数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err!= nil {
			fmt.Println("err=", err)
			return
		}

		//处理用户请求，新建衣蛾协程
		go  handleConn(conn)
	}
}