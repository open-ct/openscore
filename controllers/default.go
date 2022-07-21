/*
 * @Author: Junlang
 * @Date: 2021-07-18 00:45:51
 * @LastEditTime: 2021-07-24 20:25:03
 * @LastEditors: Junlang
 * @FilePath: /openscore/controllers/default.go
 */
package controllers

import (
	"fmt"
	"log"
	"openscore/auth"
	"openscore/models"

	beego "github.com/beego/beego/v2/server/web"
)

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

var CasdoorEndpoint, _ = beego.AppConfig.String("casdoorEndpoint")
var ClientId, _ = beego.AppConfig.String("clientId")
var ClientSecret, _ = beego.AppConfig.String("clientSecret")
var JwtSecret, _ = beego.AppConfig.String("jwtSecret")
var CasdoorOrganization, _ = beego.AppConfig.String("casdoorOrganization")

func init() {
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtSecret, CasdoorOrganization)
}

func (c *ApiController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	a := new(models.Topic)
	fmt.Println(a)
	c.Ctx.WriteString("hello OpenCT")
	c.Data["json"] = "hello OpenCT"
	c.ServeJSON()
}

func (c *ApiController) Login() {
	input, _ := c.Input()
	code := input.Get("code")
	state := input.Get("state")

	token, err := auth.GetOAuthToken(code, state)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	claims, err := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		panic(err)
	}
	log.Println(claims)

	claims.AccessToken = token.AccessToken
	c.SetSessionUser(claims)

	resp := &Response{Status: "ok", Msg: "", Data: claims}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *ApiController) Logout() {
	var resp Response

	c.SetSessionUser(nil)

	resp = Response{Status: "ok", Msg: ""}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *ApiController) GetAccount() {

	var resp Response

	if c.GetSessionUser() == nil {
		resp = Response{Status: "error", Msg: "please sign in first", Data: c.GetSessionUser()}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	claims := c.GetSessionUser()
	userObj := claims
	resp = Response{Status: "ok", Msg: "", Data: userObj}

	c.Data["json"] = resp

	c.ServeJSON()
}
