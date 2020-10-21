package routers

import (
	"github.com/shuwenhe/shuwen-beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/class/all", &controllers.ClassController{}, "get,post:Test")
	beego.Router("/page", &controllers.ArticleController{})
	beego.Router("/api/article/page", &controllers.ArticleController{}, "*:Page")
	beego.Router("/api/statisticsAll", &controllers.StatisticsController{}, "*:StatisticsAll")
}
