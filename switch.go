package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println("When's Saturday?")
	today:=time.Now().Weekday()
	switch time.Saturday {
	case today+0:
		fmt.Println("Today") //默认包含break，
		//fallthrough  //不跳出switch语句，后面的无条件执行
	case today+1:
		fmt.Println("Tomorrow")
	case today+2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}

	t:=time.Now()
	switch  {
	case t.Hour()<12:
		fmt.Println("Good morning")
	case t.Hour()<17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
