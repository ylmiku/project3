package routers

import (
	"github.com/astaxie/beego"
	"project3/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
