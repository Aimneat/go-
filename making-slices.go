package main

import "fmt"

func main(){
	a:=make([]int,5)
	//make用于内建类型（map、slice 和channel）的内存分配。
	// new用于各种类型的内存分配。
	printSlice("a",a)
	b:=make([]int,0,5)
	printSlice("b",b)
	c:=b[:2]
	printSlice("c",c)
	d:=c[2:5]
	printSlice("d",d)

	var a1 []int
	printSlice("a1",a1)

	a1=append(a1,0)
	printSlice("a1",a1)

	a1=append(a1,1)
	printSlice("a1",a1)
	//如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组
	a1=append(a1,2,3,4)
	printSlice("a1",a1)
}

func printSlice(s string,x[]int){
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}