package routers

import (
	"admin/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

//增加拦截器。
var filterAdmin = func(ctx *context.Context) {
	url := ctx.Input.URL()
	logs.Info("##### filter url : %s, %v", url, ctx.Input.Session("userinfo"))
	//TODO 如果判断用户未登录。
	_, ok := ctx.Input.Session("userinfo").(string)
	if !ok && url != "/login" {
		logs.Info("##### Redirect url : %s", url)
		ctx.Redirect(302, "/login")
		return
	}

}

func init() {

	LoginController := &controllers.LoginController{}
	beego.Router("/login", LoginController, "*:Login")
	beego.Router("/loginreq", LoginController, "*:LoginReq")
	beego.Router("/logoutreq", LoginController, "*:LogoutReq")

	MainController := &controllers.MainController{}
	beego.Router("/", MainController, "*:Home")
	beego.Router("/admin/index", MainController, "*:Index")

	userInfoController := &controllers.UserInfoController{}
	beego.Router("/admin/userInfo/edit", userInfoController, "get:Edit")
	beego.Router("/admin/userInfo/delete", userInfoController, "post:Delete")
	beego.Router("/admin/userInfo/save", userInfoController, "post:Save")
	beego.Router("/admin/userInfo/list", userInfoController, "get:List")

	foundController := &controllers.FoundController{}
	beego.Router("/admin/found/found", foundController, "*:Found")

	beego.InsertFilter("/admin/*", beego.BeforeRouter, filterAdmin)
}
