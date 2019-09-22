package main

import (
	"github.com/astaxie/beego"
	_ "heimaDay2/models"
	_ "heimaDay2/routers"
	"strconv"
)

func main() {
//	beego.SetStaticPath("/down1", "img")
	beego.AddFuncMap("ShowPrePage",HandlePrePage)
	beego.AddFuncMap("ShowNextPage",HandleNextPage)
	beego.Run()
}

func HandlePrePage(data int)(string){
	pageIndex:=data-1
	pageIndex1:=strconv.Itoa(pageIndex)
	return pageIndex1
}

func HandleNextPage(data int)(int){
	pageIndex:=data+1
	//pageIndex1:=strconv.Itoa(pageIndex)
	return pageIndex
}

