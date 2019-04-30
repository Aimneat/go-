package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func puanduandaxiao(x *int,y *int){
//	if *x>*y{
//		*x,*y=*y,*x
//	}
//}
//func main(){
//	a:=[]int{3,5,6,8,4,1,2,5,8,3,1,4}
//	for i:=0;i<len(a)-1;i++{
//		for j:=0;j<len(a)-1-i;j++{
//			puanduandaxiao(&a[j],&a[j+1])
//		}
//	}
//	fmt.Println("a=",a)
//}
func main(){
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n:=len(a)

	for i:=0;i<n;i++{
		a[i]=rand.Intn(100)
		fmt.Printf("%d,",a[i])
	}
	fmt.Printf("\n")

	for i:=0;i<n-1;i++{
		for j:=0;j<n-1-i;j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	for i:=0;i<n;i++{
		fmt.Printf("%d,",a[i])
	}
}