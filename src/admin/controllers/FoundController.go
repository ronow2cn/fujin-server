package controllers

import (
	"comm/dbmgr"
	"comm/logger"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
)

var log = logger.DefaultLogger

type FoundController struct {
	beego.Controller
}

func (c *FoundController) Found() {
	page, _ := c.GetInt64("page")

	articles := dbmgr.GetArticlesByLimit(int(page), 20)

	c.Data["List"] = articles

	c.TplName = "found/found.html"

}
