package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var(
	v1=Vertex{3,4}
	v2=Vertex{X:1}
	v3=Vertex{}
	pv=&Vertex{4,5}
)

func  main(){
	i,j:=42,2701

	p:=&i
	fmt.Println(*p)
	*p=21
	fmt.Println(i)

	p=&j
	*p=*p/37
	fmt.Println(j)

	v:=Vertex{1,2}
	fmt.Println(v)
	v.X=4
	fmt.Println(v.X)

	z := &v
	z.X=1e9
	fmt.Println(v)

	fmt.Println(v1,*pv,v2,v3)
}