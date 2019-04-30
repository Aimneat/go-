package main

import "fmt"

func swapa(p1,p2 *int){
	*p1,*p2=*p2,*p1
}
func main(){
	a,b:=10,20

	swapa(&a,&b)  //地址传递
	fmt.Printf("main:a=%d,b=%d\n",a,b)
}