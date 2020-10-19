package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shuwenhe/shuwen-beego/dao"
	"github.com/shuwenhe/shuwen-beego/utils"
)

type ArticleController struct {
	beego.Controller
}

func (ctx *ArticleController) Page() {
	count, _ := dao.ArticleCount()
	if count == 0 {
		ctx.Data["json"] = utils.NewResult(utils.Fail, "Not get the count!")
		ctx.ServeJSON()
		return
	}
	pi, _ := ctx.GetInt("pi")
	ps, _ := ctx.GetInt("ps")
	articles, err := dao.ArticlePage(pi, ps)
	if err != nil {
		ctx.Data["json"] = utils.NewResult(utils.Fail, err.Error())
		ctx.ServeJSON()
		return
	}
	ctx.Data["json"] = utils.NewPage(utils.Success, "Page data", articles, count)
	ctx.ServeJSON()
	return
}

func (ctx *ArticleController) Get() {
	ctx.TplName = "page.html"
}
