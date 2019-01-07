package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.html"
}


func (c *MainController) Post() {
//1.拿到数据
userName := c.GetString("name")
    pwd  := c.GetString("pwd")

//2.校验数据
if userName  == ""|| pwd == ""  {
	beego.Info("数据不能为空")
	c.Redirect("/reg", 302)
	//return
}
//3.插入数据库
     o := orm.NewOrm()
     user := models.User{}
     user.Name = userName
     user.Pwd = pwd

     _,err := o.Insert(&user)
     if err != nil {
		 beego.Info("插入数据失败")
		 c.Redirect("/reg", 302)
		 return
	 }
     c.Ctx.WriteString("插入成功")

}

/*
登录方法
 */
 //登录get 方法
func (c *MainController) Showlogin() {
c.TplName = "login.html"
}
//登录post方法
func (c *MainController) HandleLogin() {
	//1.拿到数据
	userName := c.GetString("username")
	pwd   := c.GetString("pwd")
	//2.判断数据
	if userName == "" || pwd == "" {
		beego.Info("输入数据不合法，请从新输入")
		c.TplName = "login.html"
		return
	}
	//3.插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd


	err := o.Read(&user,"Name","pwd")
	if err != nil {
		beego.Info("查询失败")
		c.Redirect("/reg", 302)
		return
	}

	c.Ctx.WriteString("登录成功")


}
