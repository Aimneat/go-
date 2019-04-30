package main

import (
	  "fmt"
)

const (
	x=iota
	y=iota
	z=iota
	w
)

const v=iota
const(
	h,i,j =iota,iota,iota
)

const(
	a =iota
	b="B"
	c=iota
	d,e,f=iota,iota,iota
	g=iota
)

func main(){
	fmt.Println(a,b,c,d,e,f,g,h,i,j,x,y,z,w,v)
}