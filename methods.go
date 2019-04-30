package main

import (
	"fmt"
	"math"
)

type Vertexr struct{
	X,Y float64
}

type MyFloat float64
func (f MyFloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}
/*
方法接收者 出现在 func 关键字和方法名之间的参数中
 */
func (v *Vertexr) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
func main(){
	v:=&Vertexr{3,4}
	fmt.Println(v.Abs())

	f:=MyFloat(-math.Sqrt(2))
	fmt.Println(f.Abs())
}