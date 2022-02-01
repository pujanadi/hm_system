package main

import (
	_ "hotel_management_system/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("/static", "/Users/Rima/workspace/Go/hotel_management_system/static")

	beego.Run()
}
