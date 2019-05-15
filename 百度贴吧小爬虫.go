package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//爬取网页内容
func httpGet(url string)(result string,err error){
	resp,err1:=http.Get(url)
	if err1 != nil {
		err=err1
		return
	}
	defer resp.Body.Close()

	//读取网页body内容
	buf:=make([]byte,1024*4)
	for{
		n,err:=resp.Body.Read(buf)
		if n==0{    //读取结束或者出问题
			fmt.Println("resp.Body.Read err = ",err)
			break
		}
		result+=string(buf[:n])
	}
	return
}

func doWork(start,end int){
	fmt.Printf("正在爬取%d到%d的页面\n",start,end)

	//明确目标（要知道你准备在哪个范围或者网站去搜索）
	for i:=start;i<=end;i++{
		url:="http://tieba.baidu.com/f?kw=%E5%AD%A6%E4%B9%A0&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
		fmt.Println("url=",url)

		//2)爬（将所有的网站内容全部爬下来）
		result,err:=httpGet(url)
		if err != nil {
			fmt.Println("httpGet err=", err)
			continue
		}

		//把内容写入文件
		fileName:="学习吧第"+strconv.Itoa(i)+"页.html"
		f,err1:=os.Create(fileName)
		if err1!= nil {
			fmt.Println("os.Create err1=", err1)
			continue
		}
		f.WriteString(result)

		f.Close()  //关闭文件
	}
}

func main(){
	var start,end int
	fmt.Printf("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=1）：")
	fmt.Scan(&end)

	doWork(start,end)
}