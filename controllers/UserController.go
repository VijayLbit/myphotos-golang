package controllers

import (
	"html/template"
	"myphotos/models"

	"github.com/astaxie/beego"
)

// UserController ...
// Struct of UserController
type UserController struct {
	ExtendedController
}

// Register ...
// Get login page
func (uc *UserController) Register() {
	user := models.User{}
	if uc.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		if err := uc.ParseForm(&user); err != nil {
			beego.Error("Couldn't parse the form. Reason: ", err)
			flash.Error("Technical error: Failed parsing form. Try again")
			uc.Data["msg"] = err
		} else {
			uc.Data["Users"] = user
			isValid, valid := user.Validate()
			if !isValid {
				uc.Data["Errors"] = valid.ErrorsMap
				flash.Error("Error: Some inputs failed validation!")
				flash.Store(&uc.Controller)
			} else {
				uc.Data["msg"], err = user.InsertNew()
				flash.Success("User inserted successfully; you may login now")
				flash.Store(&uc.Controller)
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
	beego.ReadFromRequest(&uc.Controller)
	userId := uc.Session.Get("userId")
	flash := beego.NewFlash()
	if userId != nil {
		flash.Success("Already logged in")
		flash.Store(&uc.Controller)
		uc.Redirect("/user-home", 302)
	}
	if uc.Ctx.Input.Method() == "POST" {
		ua := models.UserAuth{}
		if err := uc.ParseForm(&ua); err != nil {
			beego.Error("Couldn't parse the form. Reason: ", err)
			flash.Error("Technical error: Failed parsing form. Try again")
			uc.Data["msg"] = err
		} else {
			user, valid := ua.VerifyLoginRequest()
			if user.Id == 0 {
				uc.Data["Errors"] = valid.ErrorsMap
				flash.Error("Error: Login failed!")
				flash.Store(&uc.Controller)
			} else {
				uc.Session.Set("userId", user.Id)
				uc.Session.Set("userName", user.Name)
				uc.Redirect("/user-home", 302)
			}
		}
	}
	uc.setLayout("login.tpl", true)
}

func (uc *UserController) UserHome() {
	beego.ReadFromRequest(&uc.Controller)
	flash := beego.NewFlash()
	userId := uc.Session.Get("userId")
	uc.Data["loginId"] = 0
	if userId != nil {
		uc.Data["loginId"] = userId
		uc.Data["name"] = uc.Session.Get("userName")
	} else {
		flash.Error("Login to access your account!")
		flash.Store(&uc.Controller)
		uc.Redirect("/login", 302)
	}
	uc.setLayout("user-home.tpl", false)
}

func (uc *UserController) Logout() {
	flash := beego.NewFlash()
	if err := uc.Session.Flush(); err == nil {
		flash.Success("Successfully logged out")
	} else {
		flash.Error("Error logging out; try again")
	}
	flash.Store(&uc.Controller)
	uc.Redirect("/login", 302)
}
