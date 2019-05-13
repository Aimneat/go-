package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//并发聊天服务器有广播上线，广播消息，查询在线用户，修改用户名，用户主动退出，超时处理功能

type client struct {
	c chan string  //用户发送数据的管道
	name string  //用户名
	addr string  //网络地址
}

//保存在线用户  cliAddr====>client
var  onlineMap map[string]client
var message = make(chan string)

//新开一个协程，转发消息，只要有消息来了，遍历map，给map每个成员都发送此消息
func manager(){
	//给map分配空间
	onlineMap=make(map[string]client)

	for {
		msg:=<-message  //没有消息前，这里阻塞

		//遍历map，给map每个成员都发送此消息
		for _,cli:=range onlineMap{
			cli.c<-msg
		}
	}
}

func writeMsgToClient(cli client,conn net.Conn){
	for msg:=range cli.c{
		//给当前客户端发送信息
		conn.Write([]byte(msg+"\n"))
	}
}

func makeMsg(cli client,msg string)(buf string){
	buf="["+cli.addr+"]"+cli.name+": "+msg
	return
}

func handleconn(conn net.Conn){
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr:=conn.RemoteAddr().String()

	//创建一个结构体，默认用户名和网络地址一样
	cli:=client{make(chan string),cliAddr,cliAddr}

	//把结构体添加到map
	onlineMap[cliAddr]=cli

	//新开一个协程，专门给当前客户端发送信息
	go writeMsgToClient(cli,conn)
	//广播某个在线
	message <-makeMsg(cli,"login")
	//提示我是谁
	cli.c<-makeMsg(cli,"I am here")

	isQuit:=make(chan bool)  //对方是否主动退出
	hasData:=make(chan bool)  //对方是否有数据发送

	//新建一个协程，接收用户发送过来的数据
	go func(){
		buf:=make([]byte,2048)
		for {
			n,err:=conn.Read(buf)
			if n==0{  //对方断开，或者出问题
				isQuit<-true
				fmt.Println("conn.Read err=",err)
				return
			}
			msg:=string(buf[:n-1])  //通过windows nc测试，多了一个换行
			if len(msg)==3&&msg=="who"{
				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _,tmp:=range onlineMap{
					msg=tmp.addr+":"+tmp.name+"\n"
					conn.Write([]byte(msg))
				}
			}else if  len(msg)>=8&&msg[:6]=="rename"{
				//raname|mike
				name:=strings.Split(msg,"|")[1]
				cli.name=name
				onlineMap[cliAddr]=cli
				conn.Write([]byte("rename success\n"))
			}else{
				//转发此内容
				message<-makeMsg(cli,msg)
			}

			hasData<-true  //代表有数据
		}
	}()

	for{
		//通过select检测channel的流动
		select{
			case <-isQuit:
				delete(onlineMap,cliAddr)  //当前用户退出
				message<-makeMsg(cli,"login out\n")  //广播谁下线了
				return

			case <-hasData:

			case <-time.After(15*time.Second):  //15s后
				delete(onlineMap,cliAddr)    //当前用户从map移除
				message<-makeMsg(cli,"time out leave out")  //广播谁下线了
				return
		}
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

	//新开一个协程，转发消息，只要有消息来了，遍历map，给map每个成员都发送此消息
	go manager()

	//主协程，循环阻塞等待用户连接
	for{
		conn,err:=listenner.Accept()
		if err!= nil {
			fmt.Println("listenner.Accept err=", err)
			continue
		}

		go handleconn(conn) //处理用户链接
	}
}
