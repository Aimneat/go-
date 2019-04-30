package main

import "fmt"

type humaner interface {
	//方法，只有声明没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type student struct {
	name string
	id int
}
//student实现了此方法
func (tmp *student) sayhi(){
	fmt.Printf("student[%s, %d ]sayhi\n",tmp.name,tmp.id)
}
type teacher struct {
	addr string
	group string
}
//teacher实现了此方法
func (tmp *teacher) sayhi(){
	fmt.Printf("teacher[%s, %s ]sayhi\n",tmp.addr,tmp.group)
}

type mystr string
//mystr实现了此方法
func (tmp *mystr)sayhi(){
	fmt.Printf("mystr[%s] sayhi\n",*tmp)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，多态
func whosayhi(i humaner){
	i.sayhi()
}
func main(){
	s:=&student{"mike",666}
	t:=&teacher{"bj","go"}
	var str mystr="hello mike"
	//调用同一个函数，不同表现，多态，多种形态
	whosayhi(s)
	whosayhi(t)
	whosayhi(&str)

	//创建一个切片
	x:=make([]humaner,3)
	x[0]=s
	x[1]=t
	x[2]=&str
	//第一个返回下标，第二个返回下标所对应的值
	for _,i:=range x{
		i.sayhi()
	}
}