package main

import (
	"fmt"
	"regexp"
)

func main(){
	buf:=`	<div>哈哈</div>
	<div>你过来啊
	不敢吧
	 哎哎哎</div>
	<div>爱你</div>
	<div>大国如果</div>
`

	reg:=regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg==nil{
		fmt.Println("MustCompile err")
		return
	}
	//2）根据规则提取关键信息
	result2:=reg.FindAllStringSubmatch(buf,-1) //-1表示提取所有匹配的个数
	//fmt.Println("result1=",result2)

	//过滤<></>
	for _,text:=range result2{
		//fmt.Println("text[0]= ",text[0])  //带<></>
		fmt.Println("text[1]= ",text[1])//不带<></>
	}
}