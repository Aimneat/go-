package main

import (
	"fmt"
	"time"
)

//主协程退出了，其它子协程也要跟着退出，（有的程序可能导致子协程还没来得及调用）
func main(){
	go func(){
		i:=0
		for{
			i++
			fmt.Println("子协程 i= ",i)
			time.Sleep(time.Second)
		}
	}()

	i:=0
	for{
		i++
		fmt.Println("main i= ",i)
		time.Sleep(time.Second)

		if i==2{
			break
		}
	}
}