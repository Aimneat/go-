package main

import (
	"fmt"
	"math"
)

func adder() func(int) int{
	sum:=0
	return func(x int)int{
		sum+=x
		return sum
	}
}

//...type 表示不定参数类型
//注意：不定参数，一定（只能）放在形参中的最后一个参数
//func MyFunc02(args ...int){  //传递的实参可以是0个或多个
//	fmt.Println("len(args)=",len(args))
//	//返回2个值，第一个是下标，第二个是下标对应的数
//	for  i,_:=range args{
//		fmt.Printf("args[%d]=%d\n",i,args[i])
//	}
//
//}

func main(){
	hypot:=func(x,y float64)float64{
		return math.Sqrt(x*x+y*y)
	}
	fmt.Println(hypot(3,4))

	/*
	Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上
	 */
	pos,neg:=adder(),adder()
	for i:=0;i<10;i++{
		fmt.Println(pos(i),neg(-2*i),)
	}
}