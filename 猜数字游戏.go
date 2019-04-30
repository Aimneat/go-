package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int){
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num=rand.Intn(10000)
		if num>=1000{
			break
		}
	}
	*p= num
}
func GetNum(s []int,num int){
	for i:=3;i>=0;i--{
		s[i]=num%10
		num/=10
	}
}
func OnGame(randSlice []int){
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Printf("请输入一个4位数：")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合要求")
		}

		GetNum(keySlice, num)
		fmt.Println("keySlice= ", keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}
		}
		if n==4{
			fmt.Println("全都猜对了 ")
			break
		}
	}
}
func main(){
	var randNum int
	CreatNum(&randNum)
	fmt.Println("randNum: ",randNum)

	randSlice:=make([]int,4)
	GetNum(randSlice,randNum)
	fmt.Println("randSlice= ",randSlice)

	OnGame(randSlice)
}
