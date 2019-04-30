package main

import "fmt"

/*
在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型

type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
 */

type testInt func(int) bool // 声明了一个函数类型
func isOdd(integer int)bool{
	if integer%2==0{
		return false
	}
	return true
}
func isEven(integer int)bool{
	if integer%2==0{
		return true
	}
	return false
}
//声明的函数类型在这个地方当做一个参数
func filter(slice []int, f testInt)[]int{
	var result []int
	for _,value:=range slice{
		if f(value){
			result=append(result,value)
		}
	}
	return result
}
func main(){
	slice :=[]int {1,2,3,4,5,7}
	fmt.Println("slice =",slice)
	odd:=filter(slice,isOdd)
	fmt.Println("Odd elements of slice are:",odd)
	even:=filter(slice,isEven)
	fmt.Println("Even elements of slice are :",even)
}