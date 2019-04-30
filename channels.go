package main

import "fmt"

func sum(a []int,c chan int){
	sum:=0
	for _,v:=range a{
		sum+=v
	}
	c<-sum
}

func main(){
	a:=[]int{7,2, 8, -9, 4, 0}
	c:=make(chan int)
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x, y := <-c, <-c // 从 c 中获取
	fmt.Println(x,y,x+y)

	d:=make(chan int,2)
	d<-1
	d<-2
	fmt.Println(<-d)
	d<-3
	fmt.Println(<-d)
	fmt.Println(<-d)
}
