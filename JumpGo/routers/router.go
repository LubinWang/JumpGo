package routers

import (
	"github.com/LubinWang/JumpGo/JumpGo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
