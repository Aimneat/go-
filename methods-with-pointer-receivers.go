package main

import (
	"fmt"
	"math"
)

type Vertexq struct{
	X,Y float64
}
func (v *Vertexq) Scale(f float64){
	v.X=v.X*f
	v.Y=v.Y*f
}

func (v *Vertexq) Abs() float64{
	return math.Sqrt(v.X*v.Y+v.Y*v.Y)
}

func main(){
	v:=&Vertexq{3,4}
	v.Scale(5)
	fmt.Println(v,v.Abs())
}
