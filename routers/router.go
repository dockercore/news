package routers

import (
	"news/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/reg", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{},"get:Showlogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{},"get:ShowIndex")



}
