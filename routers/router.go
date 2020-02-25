package routers

import (
	"myphotos/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/logout", &controllers.UserController{}, "get,post:Logout")
	beego.Router("/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/user-home", &controllers.UserController{}, "get:UserHome")
}
