package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string
	Pwd string
	//Age string

	Article []*Article `orm:"rel(m2m)"`
}

type Article struct {
	Id int `orm:"pk;auto"` //主键 并且自动增长
	ArtiName string `orm:"size(20)"` //ArtiName长度为20
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string
	Aimg string

	User []*User `orm:"reverse(many)"`
}



func init()  {
	//设置数据库的基本信息
	orm.RegisterDataBase("default","mysql","root:123@tcp(127.0.0.1:3306)/test5?charset=utf8")
	//映射modle数据
	orm.RegisterModel(new(User),new(Article))
	//生成表
	orm.RunSyncdb("default",false,true)
}