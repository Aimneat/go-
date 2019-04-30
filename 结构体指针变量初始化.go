package main

import (
	"fmt"
)

type Student struct{
	id int
	name string
	sex byte
	age int
	addr string
}
func test(p *Student){
	p.id=666
}
func main(){
	p2:=&Student{name:"mike",addr:"bj"}
	p2.age=18

	var s Student
	p1:=&s
	//通过指针操作成员，p1.name和（*p1).name完全等价，只能用.运算符
	p1.name="bob"
	fmt.Println("p1 = ",p1)
	test(&s) //地址传递（引用传递），形参可以改变实参
	fmt.Println("p1 = ",s)

	fmt.Printf("p2 type is %T\n",p2)
	fmt.Println("p2 = ",p2)

	//通过new申请一个结构体
	p3:=new(Student)
	p3.name="Tim"
	fmt.Println("p3 = ",p3)

}