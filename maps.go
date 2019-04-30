package main

import "fmt"

type Vertex1 struct{
	Lat,Long float64
}
//map作为函数参数实际上是作为引用传递的
var m=map[string]Vertex1{
	"Bell labs":{
		40.68433,-74.39967,
	},
	"google":{
		37.42202,-122.08408,
	},

}

func main(){
/*	m=make(map[string]Vertex1)
	m["Bell Labs"]=Vertex1{
		40.68433,-74.39967,
	}
*/
//m4:=map[int]string{1:"mike",2:"go",3:"C++“}
	fmt.Println(m)

	mm:=make(map[string]int)
	mm["answer"]=42
	fmt.Printf("%d ",mm["answer"])
	mm["answer"]=48
	fmt.Printf("%d ",mm["answer"])
	delete(mm,"answer")
	fmt.Printf("%d ",mm["answer"])
/*
通过双赋值检测某个键存在：

elem, ok = m[key]
如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值
第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为true
 */
	v,ok:=mm["answer"]
	fmt.Println("the value:",v,"present?",ok)

}
