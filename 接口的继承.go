package main

import (
	"fmt"
)

type humaner interface { //子集
	//方法，只有声明没有实现，由别的类型（自定义类型）实现
	sayhi()
}
type personer interface { //超集
	humaner  //匿名字段，继承了sayhi（）
	sing (lrc string)
}
type student struct {
	name string
	id int
}
//student实现了sayhi方法
func (tmp *student) sayhi(){
	fmt.Printf("student[%s, %d ]sayhi\n",tmp.name,tmp.id)
}
//
func (tmp *student)sing(lrc string){
	fmt.Println("students在唱歌",lrc)
}

func main(){
	//定义一个接口类型的变量
	var i personer  // 超集
	s:=&student{"mike",666}
	i=s
	i.sayhi() //继承过来的方法
	i.sing("好汉歌")

	var x humaner
	x=s  //超集可以转换为子集，子集不可以转换为超集
	x.sayhi()
}
