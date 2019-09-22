package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"heimaDay2/models"
	"math"
	"path"
	"strconv"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
/*简单插入过程

	//1.有orm对象
	o:=orm.NewOrm()
	//2.有一个要插入数据的结构体对象
	user:=models.User{}
	//3.对结构体赋值
	user.Name="google"
	user.Pwd="888"
	//4.插入
	_,err:=o.Insert(&user)
	if err!=nil{
		fmt.Println("插入失败",err)
		return
	}

*/

/*
  简单查询有关的代码
 */
 /*
 	//1.有orm对象
	o:=orm.NewOrm()
 	//2.查询的对象
	user:=models.User{}
 	//3.指定查询对象字段值
 	   //user.Id=1
 	   user.Name="google"
 	//4.查询
	   //err:=o.Read(&user)
	   err:=o.Read(&user,"name")
	if err!=nil{
		fmt.Println("查询失败",err)
		return
	}

	fmt.Println("查询成功",user)
*/

/*
  更新相关的代码
 */
 /*
	//1.有orm对象
	o:=orm.NewOrm()
	//2.需要更新的结构体对象
	user:=models.User{}
	//3.查到需要更新的数据
	user.Id=1
	err:=o.Read(&user)
	//4.给数据重新赋值
  	if err ==nil{
  		user.Name="micsoft"
  		user.Pwd="666"
		//5.更新
		_,err1:=o.Update(&user)
		if err1!=nil{
			fmt.Println("更新失败",err1)
			return
		}
	}
*/


/*
	  删除相关的代码
*/
/*
	   //1.有orm对象
	   o:=orm.NewOrm()
	   //2.需要删除的结构体对象
	   user:=models.User{}
	   //3.指定删除哪一条数据
	   user.Id=1
	   //4.删除
	_,err:=o.Delete(&user)
	if err!=nil{
		fmt.Println("删除失败",err)
		return
	}
*/


/*
	   c.Data["data"] = "今天中午吃饺子"
	   c.TplName = "test.html"
*/

	c.TplName = "register.html"
}

   func (c *MainController) Post() {
/*	   c.Data["data"] = "今天中午吃面条"

	   c.TplName = "test.html"
*/

    /*
    注册业务代码
     */
     //1.拿到数据
	userName:=c.GetString("userName")
	pwd:=c.GetString("pwd")
	   //fmt.Println(userName,pwd)
	//2.对数据进行校验
	if userName==""||pwd==""{
		fmt.Println("数据不能空")
		c.Redirect("/register",302)
		return
	}
	//3.插入数据库
	o:=orm.NewOrm()
	user:=models.User{}
	user.Name=userName
	user.Pwd=pwd
	_,err:=o.Insert(&user)
	if err!=nil{
		 fmt.Println("插入失败",err)
		 c.Redirect("/register",302)
		 return
	}
	//4.返回登录界面
	  //c.Ctx.WriteString("注册成功")
	  c.TplName="login.html"  //同时还可以给视图传输数据
	  //c.Redirect("/login",302)
   }

func (c *MainController)ShowLogin(){
	name:=c.Ctx.GetCookie("userName")

	if name!=""{
		c.Data["userName"]=name
		c.Data["check"]="checked"
	}

	c.TplName="login.html"
}
//登录业务处理
func (c *MainController)HandleLogin(){
	//c.Ctx.WriteString("这是登陆的POST请求")
	//1.拿到数据
	name:=c.GetString("userName")
	pwd:=c.GetString("pwd")
	//2.判断数据是否合法
	if name==""||pwd==""{
		fmt.Println("输入数据不合法")
		c.TplName="login.html"
		return
	}
	//3.查询账号密码是否正确
	o:=orm.NewOrm()
	user:=models.User{}

	user.Name=name
	err:=o.Read(&user,"Name")
	if err!=nil{
		fmt.Println("查询失败")
		c.TplName="login.html"
		return
	}

	//记住用户名
	check:=c.GetString("remember")
	if check=="on"{
		c.Ctx.SetCookie("userName",name,time.Second*300)
	}else{
		c.Ctx.SetCookie("userName","deleteCookies",-1)
	}

	c.SetSession("userName",name)

	//4.跳转
	//c.Ctx.WriteString("欢迎你，登录成功")

	c.Redirect("/Article/index",302)
}

//处理下拉框改变法的请求
func (this *MainController)HandleSelect()  {
	//接收数据
	typeName:=this.GetString("select")
	//fmt.Println(typeName)
	//2.处理数据
	if typeName==""{
		fmt.Println("下拉框传递数据失败")
		return
	}
	//3.查询数据
	o:=orm.NewOrm()
	var articles[]models.Article
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)
	fmt.Println(articles)
}

//显示首页内容
func (c *MainController)ShowIndex(){

	/*
	userName:=c.GetSession("userName")
	if userName==nil{
		c.Redirect("/login",302)    //如果没有session就回到登录界面
		return
	}
	*/

	o:=orm.NewOrm()
	qs:=o.QueryTable("Article")
	var articles []models.Article
	/*_,err:=qs.All(&articles)
	if err!=nil{
		fmt.Println("查询所有文章信息出错")
		return
	}*/

	//pageIndex1:=1
	pageIndex:=c.GetString("pageIndex")
	pageIndex1,err:=strconv.Atoi(pageIndex)
	if err!=nil{
		pageIndex1=1
		/*
		fmt.Println("获取页数失败")
		return
		*/
	}

	typeName:=c.GetString("select")
	var count int64
	if typeName==""{
		var err1 error
		count,err1=qs.Count() //返回数据条目数
		if err1!=nil{
			fmt.Println("查询文章信息个数出错")
			return
		}
	}else{
		var err1 error
		count,err1=qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count() //返回数据条目数
		if err1!=nil{
			fmt.Println("查询文章信息个数出错")
			return
	}
	}

	//获取总页数
	pageSize:=2

	//pageIndex:=1
	start:=pageSize*(pageIndex1-1)
	qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles) //1. pagesize 一页显示多少 2.start 起始位置

	pageCount:=float64(count)/float64(pageSize)
	pageCount1:=math.Ceil(pageCount)   //天花板函数获取正确总页码

	FirstPage:=false
	if pageIndex1==1{
		FirstPage=true
	}
	LastPage:=false
	if pageIndex1==int(pageCount1){
		LastPage=true
	}

	//获取类型数据
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	c.Data["types"]=types

	//根据类型获取数据
	//接收数据
	//typeName:=c.GetString("select")
	fmt.Println(typeName)
	//2.处理数据
	var articleswithtype []models.Article
	if typeName==""{
		fmt.Println("下拉框传递数据失败")
		qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articleswithtype)
	}else{
		qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articleswithtype)
	}
	//3.查询数据



	c.Data["typeName"]=typeName
	c.Data["count"]=count

	c.Data["FirstPage"]=FirstPage
	c.Data["LastPage"]=LastPage
	c.Data["pageCount"]=pageCount1
	c.Data["pageIndex"]=pageIndex1
	//fmt.Println("count= ",count)
	//fmt.Println(articles)
	c.Data["articles"]=articleswithtype

	c.Layout="layout.html"
	c.TplName="index.html"

	//bug:index界面无法选择其它类型文章
	//待优化：选择文章类型后，不能快速查看全部文章

}

//显示添加文章界面
func (c *MainController)ShowAdd(){
	//查询类型数据然后传递到视图中
	//
	o:=orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	c.Data["types"]=types
	c.TplName="add.html"
}

//处理添加文章界面数据
func (c *MainController)HandleAdd(){
	//1.拿到数据
	artiName:= c.GetString("articleName")
	artiContent:= c.GetString("content")

	file,h,err2:=c.GetFile("uploadname")
	defer file.Close()

	//1.要限定格式
	fileext:=path.Ext(h.Filename)
	if fileext!=".jpg"&&fileext!=".png"{
		fmt.Println("上传文件格式错误")
		return
	}
	//2.限制大小
	if h.Size>50000000{
		fmt.Println("上传文件过大")
		return
	}
	//3.需要对文件重命名，防止文件名重复
	//filename:=time.Now().Format("2006-01-02 15:04:05")+fileext  //时间规定死了，否则出现的时间不对 6-1-2 3:4:5,Windows好像无法识别这个时间格式
	filename:=time.Now().Format("2006-01-02-15-04-05")+fileext
	if err2!=nil{
		fmt.Println("上传文件失败")
		return
	}else{
		c.SaveToFile("uploadname","./static/img/"+filename)
	}
	  fmt.Println(artiName,artiContent)
	//2.判断数据是否合法
	if artiContent==""||artiName==""{
		fmt.Println("添加文章数据错误")
		return
	}
	//3.插入数据
	o:=orm.NewOrm()
	arti:=models.Article{}
	arti.ArtiName=artiName
	arti.Acontent=artiContent
	arti.Aimg="/static/img/"+filename

	//4.返回文章界面

	//给article对象赋值
	//获取到下拉框传递过来的类型数据
	 typeName:=c.GetString("select")
	//类型判断
	if typeName==""{
		fmt.Println("下拉框数据读取错误")
		return
	}
	//获取type对象
	var artiType models.ArticleType
	 artiType.TypeName=typeName
	 err3:=o.Read(&artiType,"TypeName")
	if err3!=nil{
		fmt.Println("获取类型错误")
		return
	}
	 arti.ArticleType=&artiType

	_,err1:=o.Insert(&arti)
	if err1!=nil{
		fmt.Println("插入数据库错误")
		return
	}

	c.Redirect("/Article/index",302)
}

//显示内容详情界面
func (c *MainController)ShowContent(){
	//1.获取文章ID
	id,err:=c.GetInt("id")
	//fmt.Println("id is ",id)
	if err!=nil{
		fmt.Println("获取文章ID错误",err)
		return
	}
	//2.查询数据库获取数据
	o:=orm.NewOrm()
	arti:=models.Article{Id:id}
	err=o.Read(&arti)
	if err!=nil{
		fmt.Println("查询错误",err)
		return
	}
	arti.Acount+=1

	//多对多插入读者
	//1获取操作对象
	  //article:=models.Article{Id:id}
	//2获取多对多操作对象
	m2m:=o.QueryM2M(&arti,"User")
	//3获取插入对象
	userName:=c.GetSession("userName")
	user:=models.User{}
	user.Name=userName.(string)
	o.Read(&user,"Name")
	//4.多对多插入
	_,err=m2m.Add(&user)
	if err!=nil{
		fmt.Println("插入失败",err)
		return
	}

	o.Update(&arti)  //没有指定更新哪一列，他会自己查


	 o.LoadRelated(&arti,"User")
	  //o.QueryTable("Article").RelatedSel("User").Filter("User_User_Name",userName.(string)).Distinct().Filter("Id",id).One(&arti)
	  /*.Filter("字段名_表名_查询对应的字段)
	  						上面的方法2暂时有问题																												*/

	fmt.Println(arti)
	//3.传递数据给视图
	c.Data["article"]=arti
	c.Layout="layout.html"
	c.LayoutSections=make(map[string]string)
	c.LayoutSections["contentHead"]="head.html"
	c.TplName="content.html"
}

//显示编辑界面
func (c *MainController)ShowUpdate(){
	//1.获取文章ID
	id,err:=c.GetInt("id")
	//fmt.Println("id is ",id)
	if err!=nil{
		fmt.Println("获取文章ID错误",err)
		return
	}
	//2.查询数据库获取数据
	o:=orm.NewOrm()
	arti:=models.Article{Id:id}
	err=o.Read(&arti)
	if err!=nil{
		fmt.Println("查询错误",err)
		return
	}


	//3.传递数据给视图
	c.Data["article"]=arti

	c.TplName="update.html"
}

//处理更新业务数据
func (c *MainController)HandleUpdate(){
	//1.拿到数据
	id,_:=c.GetInt("id")
	artiName:= c.GetString("articleName")
	content:= c.GetString("content")
	file,h,err2:=c.GetFile("uploadname")

	var filename string
	if err2!=nil{
		fmt.Println("上传文件失败")
		return
	}else{
		defer file.Close()
		//1.要限定格式
		fileext:=path.Ext(h.Filename)
		if fileext!=".jpg"&&fileext!=".png"{
			fmt.Println("上传文件格式错误")
			return
		}
		//2.限制大小
		if h.Size>50000000{
			fmt.Println("上传文件过大")
			return
		}
		//3.需要对文件重命名，防止文件名重复
		//filename:=time.Now().Format("2006-01-02 15:04:05")+fileext  //时间规定死了，否则出现的时间不对 6-1-2 3:4:5 ,Windows好像无法识别这个时间格式
		filename:=time.Now().Format("2006-01-02-15-04-05")+fileext  //时间规定死了，否则出现的时间不对 6-1-2 3:4:5
		c.SaveToFile("uploadname","./static/img/"+filename)
	}



	//2.对数据进行一个处理
	if content==""||artiName==""{
		fmt.Println("更新获取数据失败")
		return
	}
	//3.更新操作
	o:=orm.NewOrm()
	arti:=models.Article{Id:id}
	err1:=o.Read(&arti)
	if err1!=nil{
		fmt.Println("查询数据错误",err1)
		return
	}
	arti.ArtiName=artiName
	arti.Acontent=content
	arti.Aimg="/static/img/"+filename
	_,err:=o.Update(&arti,"ArtiName","Acontent","Aimg")
	if err!=nil{
		fmt.Println("更新数据显示错误",err)
		return
	}
	//3.返回列表页面

	c.Redirect("/Article/index",302)
}

//删除操作
func (c *MainController)HandleDelete(){
	//1.拿到数据
	id,err:=c.GetInt("id")
	if err!=nil{
		fmt.Println("获取id数据错误")
		return
	}

	//2.执行删除操作
	o:=orm.NewOrm()
	arti:=models.Article{Id:id}
	err=o.Read(&arti)
	if err!=nil{
		fmt.Println("查询错误",err)
		return
	}
	o.Delete(&arti)
	//3.返回列表页面

	c.Redirect("/Article/index",302)
}

func (this *MainController)ShowAddType(){
	//1.读取类型表，显示数据
	o:=orm.NewOrm()
	var artiTypes[]models.ArticleType
	//查询
	_,err:=o.QueryTable("ArticleType").All(&artiTypes)
	if err!=nil{
		fmt.Println("查询类型错误或为空")
	}
	this.Data["types"]=artiTypes
	this.TplName="addType.html"
}

//处理添加类型业务
func (this *MainController)HandleAddType(){
	//1.获取数据
	typename:=this.GetString("typeName")

	//2.判断数据
	if typename==""{
		fmt.Println("添加类型数据为空")
		return
	}
	//3.执行插入操作
	o:=orm.NewOrm()
	var artiType models.ArticleType
	artiType.TypeName=typename
	_,err:=o.Insert(&artiType)
	if err!=nil{
		fmt.Println("插入失败")
	}
	//4.展示视图
	this.Redirect("/AddArticleType",302)
}

//退出登录
func (this *MainController)Logout(){
	//1.删除登录状态
	this.DelSession("userName")

	//2.跳转登录界面
	this.Redirect("/login",302)
}

