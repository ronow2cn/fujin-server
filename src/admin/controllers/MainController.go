package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Home() {
	c.Ctx.Redirect(302, "/admin/index")
}

func (c *MainController) Index() {
	c.TplName = "index.html"
}
