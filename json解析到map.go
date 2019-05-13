package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonbuf := `{
	"company":"itcast",
	"subjects":["go","c++","python"],
	"isok":true,
    "price":666.666

	
	}`

	//创建一个map
	m:=make(map[string]interface{},4)

	err :=json.Unmarshal([]byte(jsonbuf),&m)  //第二个参数要地址传递
	if err!=nil{
		fmt.Println("err= ",err)
		return
	}
	fmt.Printf("tmp=%+v\n",m)

	var str string
	//str=string(m["company"])  //error，无法转换
	//fmt.Println("str=",str)


	//类型断言
	for key,value:=range m{
		switch data:=value.(type) {
		case string:
			str =data
			fmt.Printf("map[%s]的值类型为string，value=%s\n",key,str)
		case bool :
			fmt.Printf("map[%s]的值类型为bool，value=%v\n",key,data)
		case float64 :
			fmt.Printf("map[%s]的值类型为float64，value=%f\n",key,data)
		case []string :
			fmt.Printf("map[%s]的值类型为[]string，value=%v\n",key,data)
		case []interface{} :
			fmt.Printf("map[%s]的值类型为[]interface，value=%v\n",key,data)
		}
		fmt.Printf("")
	}
}