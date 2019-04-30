package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {

}
func (p *MyMux)ServeHTTP(w http.ResponseWriter,r *http.Request){
	if r.URL.Path=="/"{
		sayhelloName(w,r)
		return
	}
	http.NotFound(w,r)
	return
}
func sayhelloName(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello myroute!")
}
func main(){
	mux:=&MyMux{}
	http.ListenAndServe(":9090",mux)
}
/*
一.首先调用Http.HandleFunc
	按顺序做了几件事
	1调用了DefaultServeMux的HandleFunc
	2调用了DefaultServeMux的Handle
	3往DefaultServeMux的map[string]muxEntr中增加对应的handler的路由法则
二.其次调用http.ListenAndServe(":9090",nil)
	按顺序做了
	1实例化Server
	2调用Server的ListenAndServe()
	3调用netListen('tcp',addr)监听端口
	4启动一个for循环，在循环体中Accep请求
	5对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
	6.读取每个请求的内容w，err:=c.readRequest()
	7判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），
	handler就设置为DefaultServeMux
	8调用handler的ServeHttp
	9在这个例子中，下面就进入到DefaultServeMux.ServeHttp
	10根据request选择handler，并且进入到这个handler的ServeHTTP
		mux.handler(r).ServeHTTP(w,r)
	11选择handler
	A判断是否有路由能满足这个request（循环遍历ServeMux的muxEntry）
	B如果有路由满足，调用这个路由handler的ServeHTTP
	C如果没有路由满足，调用NotFoundHandler的ServeHTTP。
*/