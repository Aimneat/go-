package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"regexp"
	"strconv"
)

//bug：发现会少爬取一页

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
		n,_:=resp.Body.Read(buf)
		if n==0{    //读取结束或者出问题
			//fmt.Println("resp.Body.Read err = ",err)
			break
		}
		result+=string(buf[:n])
	}
	return
}

//开始爬取每一个笑话，每一个段子title,conntent,err:=spiderOneJoy(data)
func spiderOneJoy(url string)(title string,content string,err error) {
	//开始爬取页面内容
	result,err1:=httpGet(url)
	if err1 != nil {
		//fmt.Println("httpGet err=", err1)
		err=err1
		return
	}
	//取关键信息
	//取标题 <h1> 标题</h1> 只取一个
	re1:=regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1== nil {
		//fmt.Println("regexp.MustCompile err=", err)
		err=fmt.Errorf("%s","regexp.MustCompile err1 ")
		return
	}
	//取标题内容
	tmpTitle:=re1.FindAllStringSubmatch(result,1) //最后一个参数为1，只过滤第一个
	for _,data:=range tmpTitle{
		title=data[1]
		//title=strings.Replace(title,"\t","\n",-1)
		break
	}
	//取段子 <div class="content-txt pt10"> 段子内容<a id="prev" href="
	re2:=regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2== nil {
		//fmt.Println("regexp.MustCompile err=", err)
		err=fmt.Errorf("%s","regexp.MustCompile err2 ")
		return
	}
	//取内容
	tmpContent:=re2.FindAllStringSubmatch(result,-1) //最后一个参数为1，只过滤第一个
	for _,data:=range tmpContent{
		content=data[1]
		content=strings.Replace(content,"\t","",-1)  //去掉tab字符
		break
	}
	return
}

//把内容写入文件
func storeJoyToFile(i int,fileTitle,fileContent []string){
	fileName:="段子第"+strconv.Itoa(i)+"页.txt"
	f,err1:=os.Create(fileName)
	if err1!= nil {
		fmt.Println("os.Create err1=", err1)
		return
	}

	//写内容
	n:=len(fileTitle)
	for i:=0;i<n;i++{
		//写标题
		f.WriteString(fileTitle[i])
		//写内容
		f.WriteString(fileContent[i])

		f.WriteString("\n=============================================\n")
	}


	defer f.Close()  //关闭文件
}

func spiderPage(i int,page chan int){

	//1)明确爬取的url
	//https://www.pengfu.com/index_1.html
	url:="https://www.pengfue.com/xiaohua_"+strconv.Itoa(i)+".html"
	//fmt.Printf("正在爬取第%d页网页：%s\n",i,url)

	//2)爬（将所有的网站内容全部爬下来）
	result,err:=httpGet(url)
	if err != nil {
		fmt.Println("httpGet err=", err)
		return
	}
	//fmt.Println("r=",result)

	//3）取，<h1 class="dp-b"><a href="  一个段子url链接  "
	  //a解释表达式
	  re:=regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re== nil {
		fmt.Println("regexp.MustCompile err=", err)
		return
	}
	  //b取关键信息
		joyUrls:=re.FindAllStringSubmatch(result,-1)
		//fmt.Println("joyUrls=", joyUrls)

		fileTitle:=make([]string,0)
		fileContent:=make([]string,0)

		//c)取网址
		//第一个返回下标，第二个返回内容
		for _,data:=range joyUrls{
			//fmt.Println("url=", data[1])
			//开始爬取每一个段子
			title,content,err:=spiderOneJoy(data[1])
			if err != nil {
				fmt.Println("spiderOneJoy err=", err)
				continue
			}
			//fmt.Printf("title =#%v#", title)
			//fmt.Printf("content=#%v#", content)
			//fmt.Println("title =", title)
			//fmt.Println("content=", content)

			fileTitle=append(fileTitle,title)  //追加内容
			fileContent=append(fileContent,content)
		}

		//fmt.Println("fileTitle= ",fileTitle)
		//fmt.Println("fileContent=",fileContent)

	////把内容写入文件
	storeJoyToFile(i,fileTitle,fileContent)


	page<-i
}

func doWork(start,end int){
	fmt.Printf("准备爬取%d到%d的页面\n",start,end)

	page:=make(chan int)

	//明确目标（要知道你准备在哪个范围或者网站去搜索）
	for i:=start;i<=end;i++{

		go spiderPage(i,page)


	}

	for i:=start;i<=end;i++{

		go spiderPage(i,page)
		fmt.Printf("第%d页网页爬取完成\n",<-page)

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
