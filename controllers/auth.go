package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct {
	beego.Controller
}

type LoginDetails struct {
	Username string
	Password string
}

// var hmacSampleSecret []byte

func (c *AuthController) Post() {
	// signinKey := []byte("aadfaofhaoijkdalknflakdfijajdslkjflakdjfladksjl")
	var details LoginDetails

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &details)

	// err := json.NewDecoder(c.Ctx.Request.Body).Decode(&requestParam)

	if err == nil {
		fmt.Println("usernameJSONReq : ", details.Username)
	} else {
		fmt.Println("error : ", err)
	}

	// data := map[string]string{
	// 	"username": "Fauzi Pujanadi",
	// 	"status":   "Logged in",
	// }

	// Create the Claims
	// claims := &jwt.RegisteredClaims{
	// 	ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
	// 	Issuer:    "test",
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// ss, err := token.SignedString(signinKey)
	// fmt.Printf("%v %v", ss, err)

	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Username string `json:"username"`
		jwt.RegisteredClaims
	}

	// Create the claims
	claims := MyCustomClaims{
		details.Username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "HM System",
			Subject:   "user",
			ID:        "1",
			Audience:  []string{"internal_user"},
		},
	}

	// Create claims while leaving out some of the optional fields
	// claims = MyCustomClaims{
	// 	"fauzi",
	// 	jwt.RegisteredClaims{
	// 		// Also fixed dates can be used for the NumericDate
	// 		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
	// 		Issuer:    "HM System",
	// 	},
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	// fmt.Printf("%v %v", ss, err)

	if err != nil {
		fmt.Printf("%v", err)
	}

	res := map[string]string{"token": ss}

	// jsonString, _ := json.Marshal(res)

	fmt.Println(res)

	c.Data["json"] = res
	c.ServeJSON()

}

func (c *AuthController) Get() {
	c.Data["json"] = "OK"

	c.ServeJSON()
}
