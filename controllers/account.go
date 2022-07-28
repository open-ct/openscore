// Copyright 2022 The OpenCT Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	_ "embed"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/open-ct/openscore/service/user"
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

// UserLogin 用户登录
func (c *ApiController) UserLogin() {
	defer c.ServeJSON()
	var req LoginRequest

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.ResponseError("cannot unmarshal", err.Error())
		return
	}

	token, err := user.Login(req.Account, req.Password)
	if err != nil {
		c.ResponseError("cannot login", err.Error())
		return
	}

	resp := struct {
		Token string `json:"token"`
	}{token}

	c.ResponseOk(resp)
}

// func (c *ApiController) Login() {
// 	input, _ := c.Input()
// 	account := input.Get("account")
// 	password := input.Get("password")
//
// 	fmt.Println("account: ", account)
//
// 	token, err := user.Login(account, password)
// 	if err != nil {
// 		c.ResponseError(err.Error())
// 		return
// 	}
//
//
//
// 	c.ResponseOk(resp)
// }

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

	// user, err := model.GetUserByCasdoorName(claims.User.Name)
	// if err != nil {
	// 	c.ResponseError(err.Error())
	// 	return
	// }
	//
	// if err := model.UpdateMemberOnlineStatus(user.UserId, true, util.GetCurrentTime()); err != nil {
	// 	c.ResponseError(err.Error())
	// 	return
	// }

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
	// claims := c.GetSessionClaims()
	// if claims != nil {
	// 	user, err := model.GetUserByCasdoorName(claims.User.Name)
	// 	if err != nil {
	// 		c.ResponseError(err.Error())
	// 		return
	// 	}
	//
	// 	if err := model.UpdateMemberOnlineStatus(user.UserId, false, util.GetCurrentTime()); err != nil {
	// 		c.ResponseError(err.Error())
	// 		return
	// 	}
	// }

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
	u := c.GetSessionUser()
	u.Score += amount
	c.SetSessionUser(u)
}

func (c *ApiController) UpdateAccountConsumptionSum(amount int) {
	u := c.GetSessionUser()
	u.Karma += amount
	c.SetSessionUser(u)
}
