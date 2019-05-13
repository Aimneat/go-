package main

import (
	"fmt"
	"runtime"
)

//主协程退出了，其它子协程也要跟着退出，（有的程序可能导致子协程还没来得及调用）
func main(){
	go func(){

		for i:=0;i<5;i++{

			fmt.Println("子协程 i= ",i)
			//time.Sleep(time.Second)
		}
	}()


	for i:=0;i<2;i++{
		fmt.Println("main i= ",i)
		//让出时间片，先让别的协程执行，那个协程的执行完时间片后，再回来执行此协程
		runtime.Gosched()
		//time.Sleep(time.Second)

	}
}