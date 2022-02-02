package routers

import (
	"hotel_management_system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/auth", &controllers.AuthController{})
}
