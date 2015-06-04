package routers

import (
	"github.com/astaxie/beego"
	"luckperson/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/push-person", &controllers.MainController{}, "get,post:Push")

}
