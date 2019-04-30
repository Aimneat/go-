package main

import "fmt"

type student struct {
	name string
	id int
}

func main(){
	i:=make([]interface{},3)  //空接口万能类型，保存任意类型的值
	i[0]=1
	i[1]="hello go"
	i[2]=student{"mike",666}

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值，data分别是i[0],i[1],i[2]
	for index,data:=range i{
		//第一个返回的是值，第二个返回判断结果的真假
		if value ,ok:=data.(int);ok==true{
			fmt.Printf("x[%d]类型为int,内容为%d\n",index,value)
		}else if value ,ok:=data.(string);ok==true{
			fmt.Printf("x[%d]类型为string,内容为%s\n",index,value)
		}else if value ,ok:=data.(student);ok==true{
			fmt.Printf("x[%d]类型为student,内容为name=%s,id=%d\n",index,value.name,value.id)
		}
	}

	//类型查询，类型断言
	for index,data:=range i{
		switch value:=data.(type){
		case int:
			fmt.Printf("x[%d]类型为int,内容为%d\n",index,value)
		case string:
			fmt.Printf("x[%d]类型为string,内容为%s\n",index,value)
		case student:
			fmt.Printf("x[%d]类型为student,内容为name=%s,id=%d\n",index,value.name,value.id)
		}
	}
}