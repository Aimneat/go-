package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human  //匿名字段Human
	school string
	loan float32
}
type Employee struct {
	Human
	company string
	money float32
}
//Human对象实现Sayhi方法
func (h *Human)SayHi(){
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}
func (h *Human)Sing(lyrics string){
	fmt.Println("la la la la la ",lyrics)
}
func (h *Human)Guzzle(beerStein string){
	fmt.Println("guzzle guzzle guzzle...",beerStein)
}
//Employee 实现Sayhi方法
func (e *Employee)SayHi(){
	fmt.Printf("Hi,I am %s ,I work at %s. Call me on %s\n",e.name,
		e.company,e.phone)//此句可以分成多行
}

//Student 实现Sayhi方法
func (s *Student)SayHi(){
	fmt.Printf("Hi,I am %s ,I study at %s. Call me on %s\n",s.name,
		s.school,s.phone)//此句可以分成多行
}
//Student实现BorrowMoney方法
func (s *Student)BorrowMoney(amount float32){
	s.loan+=amount
}
//Employee 实现SpendSalary方法
func (e *Employee)SpendSalary(amount float32){
	e.money-=amount
}
//定义interface
type Men interface {
	SayHi()
	Sing(lyrice string)
	Guzzle(beerStein string)
}
type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}
type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main(){
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
//	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var m Men

	m=&mike
	fmt.Println("This is mike,a Student:")
	m.SayHi()
	m.Guzzle("mu mu mu ")
	m.Sing("233")

	var y YoungChap
	y=&paul
	fmt.Println("This is paul,a Student:")
	y.SayHi()
	y.Sing("yi yi yi")
	y.BorrowMoney(1000)
	fmt.Println(paul.loan)

	var em ElderlyGent
	em=&sam
	fmt.Println("This is sam,a Employee")
	em.SayHi()
	em.Sing("e e e ")
	em.SpendSalary(10000)
	fmt.Println(sam.money)
}