package main

import "fmt"

func removeElement(nums []int, val int) int {
	for i,v:=range nums{
		if v==val{
			nums=append(nums[:i],nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func main(){
	nums:=[]int{0,1,2,2,3,0,4,2}
	lens:=removeElement(nums,2)
	fmt.Println("len=",lens)
}