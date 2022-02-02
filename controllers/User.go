package controllers

import (
	"fmt"

	"github.com/astaxie/beego"

	uuid "github.com/google/uuid"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Post() {
	uid := uuid.New()

	fmt.Print(uid)

	// newUser := Auth{Username: "admin", Password: "password"}

	// newUserDetail := User{Id: uid, DisplayName: "John Doe", UserLevel: "0", IsLogin: "Y"}

	userService := AuthService{}

}
