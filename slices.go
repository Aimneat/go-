package main

import "fmt"

func main(){
	//切片作为函数参数是引用传递
	p:=[]int{2,3,5,7,11,13}
	fmt.Println(("p=="),p)
	for i:=0;i<len(p);i++{
		fmt.Printf("p[%d]==%d ",i,p[i])
	}
	fmt.Printf("\n")
	fmt.Println("p cap ",cap(p))
	fmt.Println("p len ",len(p))
	fmt.Println("p[1:4]==",p[1:4])
	fmt.Println("p[1:4] cap ",cap(p[1:4]))  //容量为cap（p）-1
	fmt.Println("p[1:4] len ",len(p[1:4]))  //长度为4-1=3
	fmt.Println("p[4:]==",p[4:])
	fmt.Println("p[4:] cap is ",cap(p[4:]))//容量为cap（p）-1
	fmt.Println("p[4:] len is",len(p[4:]))  //长度为len（p）-4=2
}