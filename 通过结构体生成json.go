package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写
type it struct{
	Company string     `json:"-"`//此字段不会输出到屏幕
	Subjects []string  `json:"subjects"` //二次编码
	Isok bool			`json:",string"`
	Price float64		`json:",string"`
}
func main(){
	//定义一个结构体变量，同时初始化
	s:=it{"itcast",[]string{"go","c++","python"},true,666.666}
	//编码根据内容生成json文本
	buf,err:=json.MarshalIndent(s,""," ")//格式化编码
	if err!=nil{
		fmt.Println("err= ",err)
		return
	}
	fmt.Println("buf = ",string(buf))

	//创建一个map
	m:=make(map[string]interface{},4)
	m["company"]="itcccc"
	m["subjects"]=[]string{"go","c++","java"}
	m["isok"]=false

	//编码成json
	result,err1:=json.MarshalIndent(s,""," ")//格式化编码
	if err!=nil{
		fmt.Println("err= ",err1)
		return
	}
	fmt.Println("result = ",string(result))
}


