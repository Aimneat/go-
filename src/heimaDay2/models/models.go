package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
import   _ "github.com/go-sql-driver/mysql"

//models beego一般放表的设计
type User struct {
	Id int
	Name string `orm:"unique"`
	Pwd string
	Articles []*Article `orm:"rel(m2m)"`
}
//文章结构体  文章表和文章类型表是一对多
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null "`
	Acontent string
	Aimg string
	ArticleType *ArticleType `orm:"rel(fk)"`
	User []*User`orm:"reverse(many)"`
}
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}


func init()  {
	// set default database
	orm.RegisterDataBase("default" , "mysql", "root:root@tcp(localhost:3306)/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User),new(Article),new(ArticleType))

	// create table
	orm.RunSyncdb("default" , false, true)
}
