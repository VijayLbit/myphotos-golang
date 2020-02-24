package routers

import (
	"myphotos/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/register", &controllers.UserController{}, "get,post:Register")
}
