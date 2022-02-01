package controllers

import (
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	fmt.Println("username : %v", username)
	fmt.Println("password : %v", password)

	// this.ViewPath = "http://localhost:3000"
	// c.Redirect("localhost:3000", 200)
	// c.Ctx.Redirect(200, "index.tpl")
	// HeaderTemp := map[string][]string{
	// 	"Host": {"http://localhost:3000"},
	// 	"Origin": {"http://localhost:3000"},
	// 	"Referer": {"http://localhost:3000"},
	// }
	c.Ctx.Request.Header.Set("Host", "aaaaa")
	c.Ctx.Redirect(200, "http://localhost:3000")
	c.Ctx.Request.Response.Header.Set("Referer", "http://localhost:3000")
	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, "http://localhost:3000/", 200)
}

func (c *AuthController) Get() {
	c.Data["json"] = "OK"

	c.ServeJSON()
}
