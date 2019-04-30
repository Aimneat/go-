package main

import "fmt"

//实现2数相加
//面向过程
func add01(a,b int)int{
	return a+b
}
//面向对象，方法：给某个类型绑定一个函数
type long int
func (tmp long)add02(other long)long{
	return tmp+other
}

type person struct {
	name string
	sex byte
	age int
}
//带有接收者的函数叫方法,person不能为指针或接口
func (tmp2 person) printInfo(){  //接收者为普通变量，非指针，值传递，一份拷贝
	fmt.Println("tmp= ",tmp2)
}
//通过一个函数，给成员赋值
func (p *person)setInfo(n string,s byte,a int){  //接收者为指针变量，引用传递
	p.name=n
	p.sex=s
	p.age=a
}
func main(){
	result:=add01(1,1) //普通函数调用方法
	fmt.Println("result= ",result)

	//定义一个变量
	var a long =2
	//调用方法格式：变量名.函数（所需参数）
	r:=a.add02(3)
	fmt.Println("result= ",r)

	 //面向对象只是换了一种表现形式

	 p:=person{"mike",'m',18}
	 p.printInfo()
	 p.setInfo("gg",'s',19)  //对象调用方法集中的方法，指针变量和普通变量可自动转换
	 p.printInfo()

	 var p2 person
	(&p2).setInfo("yoyo",'f',22)
	 p2.printInfo()
}