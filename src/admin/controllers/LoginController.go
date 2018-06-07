package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	c.TplName = "login/login.html"
}

func (c *LoginController) LoginReq() {
	username := c.GetString("username")
	passwd := c.GetString("password")

	if username == "admin" && passwd == "admin" {
		c.SetSession("userinfo", username)

		c.Data["json"] = &map[string]interface{}{"status": true, "info": "登录成功"}
		c.ServeJSON()
		return
	} else {

		c.Data["json"] = &map[string]interface{}{"status": false, "info": "用户名或密码错误"}
		c.ServeJSON()
		return
	}

}

func (c *LoginController) LogoutReq() {
	c.DelSession("userinfo")
	c.Ctx.Redirect(302, "/login")
}
