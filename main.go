package main

import (
	_ "hotel_management_system/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.SetStaticPath("/static", "/Users/Rima/workspace/Go/hotel_management_system/static")

	// CORS for https://foo.* origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// // - Credentials share
	// beego.InsertFilter("/auth", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowOrigins:     []string{"https://*.foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	// allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowOrigins:     []string{"https://*.foo.com"},
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type"},
		// ExposeHeaders:    []string{"Content-Length"},
		// AllowCredentials: true,
	}))

	beego.Run()
}
