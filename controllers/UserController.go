package controllers

import (
	"fmt"
	"html/template"
	"myphotos/models"
	"os"

	"github.com/astaxie/beego"
)

// UserController ...
// Struct of UserController
type UserController struct {
	beego.Controller
}

// Register ...
// Get login page
func (uc *UserController) Register() {
	user := models.User{}
	if uc.Ctx.Input.Method() == "POST" {
		if err := uc.ParseForm(&user); err != nil {
			beego.Error("Couldn't parse the form. Reason: ", err)
			uc.Data["msg"] = err
		} else {
			uc.Data["Users"] = user
			isValid, valid := user.Validate()
			if !isValid {
				uc.Data["Errors"] = valid.ErrorsMap
			} else {
				uc.Data["msg"], err = user.InsertNew()
				fmt.Println(uc.Data["msg"])
				os.Exit(1)
				if err == nil {
					uc.Redirect("/login", 302)
				}
			}
		}
	}
	uc.setLayout("register.tpl", true)
}

func (uc *UserController) setLayout(tplName string, addXsrf bool) {
	uc.Data["Website"] = "MyPhotos.com"
	uc.Layout = "master-layout.tpl"
	uc.LayoutSections = make(map[string]string)
	uc.LayoutSections["Header"] = "header.tpl"
	uc.LayoutSections["Footer"] = "footer.tpl"
	uc.TplName = tplName
	if addXsrf == true {
		uc.Data["xsrfdata"] = template.HTML(uc.XSRFFormHTML())
	}
}

func (uc *UserController) Login() {
	uc.setLayout("login.tpl", true)
}
