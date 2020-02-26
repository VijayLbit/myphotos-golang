package controllers

import (
	"html/template"

	"github.com/astaxie/beego"
)

// MainController ...
// Struct of MainController
type MainController struct {
	beego.Controller
}

// Get ...
// Default get page
func (c *MainController) Get() {
	c.Data["Website"] = "MyPhotos.com"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "master-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = "index.tpl"
	c.Data["requestUri"] = c.Ctx.Request.RequestURI
}
