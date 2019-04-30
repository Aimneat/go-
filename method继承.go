package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human   //只有类型没有名字，匿名字段，继承了Human的成员
	school string
}
type Employee struct {
	Human
	company string
}
func (h *Human)SayHi(){
	fmt.Printf("hi,i am %s you can call me on %s\n",h.name,h.phone)
}
//employee的方法和human的方法同名，这叫做重写
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}
func main(){
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()  //继承了human的方法
	sam.SayHi()
}