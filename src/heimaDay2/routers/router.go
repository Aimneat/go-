package routers

import (
	"github.com/astaxie/beego/context"
	"heimaDay2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})

	beego.InsertFilter("/Article/*",beego.BeforeRouter,FilterFunc)
	beego.Router("/register", &controllers.MainController{})
	//注意：当实现了自定义的get请求方法，请求将不会访问默认方法
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/Article/index", &controllers.MainController{},"get:ShowIndex;post:HandleSelect")

	beego.Router("/Article/addArticle", &controllers.MainController{},"get:ShowAdd;post:HandleAdd")
	beego.Router("/Article/content", &controllers.MainController{},"get:ShowContent")
	beego.Router("/Article/update", &controllers.MainController{},"get:ShowUpdate;post:HandleUpdate")
	beego.Router("/Article/delete", &controllers.MainController{},"get:HandleDelete")

    //添加类型
	beego.Router("/Article/AddArticleType", &controllers.MainController{},"get:ShowAddType;post:HandleAddType")

    //退出登录
    beego.Router("/Article/Logout", &controllers.MainController{},"get:Logout")
}

var FilterFunc=func (ctx *context.Context){
	userName:=ctx.Input.Session("userName")
	if userName==nil{
		ctx.Redirect(302,"/login")
	}
}
