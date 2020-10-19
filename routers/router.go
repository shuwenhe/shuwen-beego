package routers

import (
	"beego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/class/all", &controllers.ClassController{}, "get,post:Test")
}
