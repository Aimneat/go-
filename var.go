package main

import (
	"fmt"
	"math"
	"math/cmplx"
)
var i,j int=1,2
var c, python, java bool  //没有明确初始值的变量声明会被赋予它们的零值（ 0 false  空字符串）
var (
	ToBe bool =false
	MaxInt uint64 =1<<64-1
	z complex128=cmplx.Sqrt(-5+12i)
)
const Pi=3.14
func main(){
	k:=3  //函数外的每个语句必须以关键字（var，func 等等），因此：=结构不能在函数外使用

	var x,y int =3,4
	var f float64=math.Sqrt(float64(x*x+y*y))
	var zz uint =uint(f)  //go 在不同类型的项之间赋值时需要显示转换
	fmt.Println(x,y,zz,Pi)
	fmt.Println(i,j,k,c,python,java)
	fmt.Println("type:%T Value :%v \n",ToBe,ToBe)
	fmt.Println("type:%T Value :%v \n",MaxInt,MaxInt)
	fmt.Println("type:%T Value :%v \n",z,z)

	a,b:=10,30
	a,b=b,a
	fmt.Printf("a=%d,b=%d",a,b)
}
