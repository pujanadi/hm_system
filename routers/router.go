package routers

import (
	"hotel_management_system/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/auth", &controllers.AuthController{})
}
