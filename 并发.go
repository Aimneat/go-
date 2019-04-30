package main

import (
	"fmt"
	"runtime"
)

func say1(s string){
	for i:=0;i<5;i++{
		runtime.Gosched() //表示让CPU把时间片让给别人，下次某个时间继续恢复执行该goroutine
		fmt.Println(s)
	}
}
func main(){
	go say1("world") //开一个新的Goroutines执行
	say1("hello") //当前Goroutines执行
}

//上面的多个goroutine运行在同一个进程里面，共享内存数据。
// 不要通过共享来通讯， 而要通过通讯来共享。

/*
Goexit

退出当前执行的goroutine，但是defer函数还会继续调用

Gosched

让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

NumCPU

返回 CPU 核数量

NumGoroutine

返回正在执行和排队的任务总数

GOMAXPROCS

用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
*/