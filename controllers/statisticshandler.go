package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shuwenhe/shuwen-beego/dao"
	"github.com/shuwenhe/shuwen-beego/utils"
)

type StatisticsController struct {
	beego.Controller
}

func (ctx *StatisticsController) StatisticsAll() {
	statistics, err := dao.StatisticsAll()
	if err != nil {
		ctx.Data["json"] = utils.NewResult(utils.Fail, "No get the statistics msg!", err.Error())
		ctx.ServeJSON()
		return
	}
	ctx.Data["json"] = utils.NewResult(utils.Fail, "Get the statistics msg!", statistics)
	ctx.ServeJSON()
	return
}
