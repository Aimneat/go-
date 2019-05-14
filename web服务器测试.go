package main

import (
	"fmt"
	"net/http"
)

//w,给客户端回复数据
//r，读取客户端发送的数据
//服务端编写的业务逻辑处理程序
func myHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("r.method=",r.Method)
	w.Write([]byte("hello go ")) //给客户端回复数据
	fmt.Fprintln(w,"hello world")
}

func main(){
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/go",myHandler)

	//在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe(":8000",nil)
}