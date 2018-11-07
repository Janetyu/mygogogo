package routers

import (
	"class/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})

	beego.Router("/register", &controllers.MainController{})
	//注意：当实现了自定义的请求方法，请求将不会访问默认方法
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/index",&controllers.MainController{},"get:ShowIndex")

	beego.Router("/addArticle",&controllers.MainController{},"get:ShowAdd;post:HandleAdd")

	beego.Router("/content",&controllers.MainController{},"get:ShowContent")

	beego.Router("/update",&controllers.MainController{},"get:ShowUpdate;post:HandleUpdate")

	beego.Router("/delete",&controllers.MainController{},"get:HandleDelete")

	beego.Router("/addType",&controllers.MainController{},"get:ShowAddType;post:HandleAddType")
}
