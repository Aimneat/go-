package main

import "fmt"

type person struct {
	name string
	sex byte
	age int
}
func (p *person)setInfoPointer(){  //接收者为指针变量，引用传递
	fmt.Printf("sestinfopointer:%p,%v\n",p,p)
}
func main(){
	p:=person{"mike",'m',18}
	fmt.Printf("main:%p,%v\n",&p,p)

	//保存方式入口地址
	//f:=p.setInfoPointer  //这就是方法值，调用函数时，无需再传递接受者，隐藏了接收者
	//f()

	//方法表达式
	f:=(*person).setInfoPointer
	f(&p)  //显式把接收者传递过去==》p.setinfopointer
}