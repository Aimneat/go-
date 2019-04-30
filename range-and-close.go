package main

import "fmt"

func fibonacci(n int,c chan int){
	x,y:=0,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
		}
	close(c)
}
func main(){
	c:=make(chan int,10)
	go fibonacci(cap(c),c)
	//channel不需要经常去关闭。应该在生产者的地方关闭channel，而不是在消费的地方去关闭它，这样容易引起panic。

	for i:=range c{
		fmt.Println(i)
	}
}
