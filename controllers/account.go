package controllers

import (
	_ "embed"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"log"
	"openscore/model"
	"openscore/util"
)

//go:embed token_jwt_key.pem
var JwtPublicKey string

func init() {
	InitAuthConfig()
}

func InitAuthConfig() {
	var casdoorEndpoint, _ = beego.AppConfig.String("casdoorEndpoint")
	var clientId, _ = beego.AppConfig.String("clientId")
	var clientSecret, _ = beego.AppConfig.String("clientSecret")
	var casdoorOrganization, _ = beego.AppConfig.String("casdoorOrganization")
	var casdoorApplication, _ = beego.AppConfig.String("casdoorApplication")

	auth.InitConfig(casdoorEndpoint, clientId, clientSecret, JwtPublicKey, casdoorOrganization, casdoorApplication)
}

/*
// @Title Signin
// @Description sign in as a member
// @Param   code     QueryString    string  true        "The code to sign in"
// @Param   state     QueryString    string  true        "The state"
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /signin [post]
// @Tag Account API*/
func (c *ApiController) SignIn() {
	input, _ := c.Input()
	code := input.Get("code")
	state := input.Get("state")

	token, err := auth.GetOAuthToken(code, state)

	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	claims, err := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	user, err := model.GetUserByCasdoorName(claims.User.Name)
	if err != nil {
		c.ResponseError(err.Error())
		return
	}

	// 首次登录
	if user == nil {
		tag := &struct {
			Subject  string `json:"subject"`
			UserType int64  `json:"user_type"`
		}{}
		if err := json.Unmarshal([]byte(claims.Tag), tag); err != nil {
			log.Println(err)
			c.Data["json"] = Response{Status: "30001", Msg: "用户首次登录: 解析tag错误", Data: err}
			return
		}
		u := model.User{
			UserType:    tag.UserType,
			SubjectName: tag.Subject,
			CasdoorName: claims.User.Name,
		}
		fmt.Println("insert: ", "insert")

		if err := u.Insert(); err != nil {
			log.Println(err)
			c.Data["json"] = Response{Status: "30001", Msg: "用户首次登录错误", Data: err}
			return
		}
		user = &u
	}

	if err := model.UpdateMemberOnlineStatus(user.UserId, true, util.GetCurrentTime()); err != nil {
		c.ResponseError(err.Error())
		return
	}

	claims.AccessToken = token.AccessToken
	c.SetSessionClaims(claims)

	c.ResponseOk(claims)
}

/*
// @Title Signout
// @Description sign out the current member
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /signout [post]
// @Tag Account API*/
func (c *ApiController) SignOut() {
	claims := c.GetSessionClaims()
	if claims != nil {
		user, err := model.GetUserByCasdoorName(claims.User.Name)
		if err != nil {
			c.ResponseError(err.Error())
			return
		}

		if err := model.UpdateMemberOnlineStatus(user.UserId, false, util.GetCurrentTime()); err != nil {
			c.ResponseError(err.Error())
			return
		}
	}

	c.SetSessionClaims(nil)

	c.ResponseOk()
}

/*
// @Title GetAccount
// @Description Get current account
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /get-account [get]
// @Tag Account API*/
func (c *ApiController) GetAccount() {
	if c.RequireSignedIn() {
		return
	}

	claims := c.GetSessionClaims()

	c.ResponseOk(claims)
}

func (c *ApiController) UpdateAccountBalance(amount int) {
	user := c.GetSessionUser()
	user.Score += amount
	c.SetSessionUser(user)
}

func (c *ApiController) UpdateAccountConsumptionSum(amount int) {
	user := c.GetSessionUser()
	user.Karma += amount
	c.SetSessionUser(user)
}
