package main

import (
	_ "news/routers"
	"github.com/astaxie/beego"
)

func main() {

	//beego.AddFuncMap("hi",hello)
	beego.AddFuncMap("showprepage",prepage)
	beego.AddFuncMap("shownextpage",shownextpage)
	beego.Run()
}


//上一页函数
func prepage(pageindex int)(preIndex int){
	preIndex= pageindex-1
	return
}

//下一页函数
func shownextpage(pageindex int)(nextIndex int){
	nextIndex = pageindex +1
	return
}



/*
//html里面调用
func hello(in string)(out string){
	out = in + "world"
	return
}
*/

