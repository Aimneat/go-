package main

import (
	"fmt"
	"math"
)

type Abser interface {
	//方法只有声明没有实现，由别的类型（自定义类型）实现
	Abs() float64
}
func main(){
	//定义接口类型变量
	var a Abser

	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给a赋值
	f:=Myfloat(-math.Sqrt2)
	v:=Vertex{3,4}

	a=f  // a MyFloat 实现了 Abser
	a=&v  // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
}

type Myfloat float64
//Myfloat实现了Abs方法
func (f Myfloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X,Y float64
}
//Vertex实现了Abs方法
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}