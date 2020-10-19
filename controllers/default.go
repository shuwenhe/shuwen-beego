package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.xstiku.com"
	c.Data["Email"] = "1201220707@pku.edu.cn"
	c.TplName = "index.html"
}
