package main
//反射就是能检查程序在运行时的状态。

//要去反射是一个类型的值（这些值都实现了空interfa）
//1.首先需要把它转化成reflect对象
/*t:=reflect.TypeOf(i)
v:=reflect.ValueOf(i)
2.将reflect对象转换为相应的值
tag:=t.Elem().Field(0).Tag//获取定义在struct里面的标签
name:=v.Elem().Field(0).String()//获取存储在第一个字段里面的值
3.获取反射值能返回相应的类型的数值
var x float64=3.4
v:=reflect.ValueOf(x)
fmt.Println("type:",v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
*/