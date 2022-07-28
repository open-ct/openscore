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
	"encoding/gob"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type ApiController struct {
	beego.Controller
}

func init() {
	gob.Register(auth.Claims{})
}

func GetUserName(user *auth.User) string {
	if user == nil {
		return ""
	}

	return user.Name
}

func (c *ApiController) GetSessionClaims() *auth.Claims {
	s := c.Controller.GetSession("user")
	if s == nil {
		return nil
	}

	claims := s.(auth.Claims)
	return &claims
}

func (c *ApiController) SetSessionClaims(claims *auth.Claims) {
	if claims == nil {
		c.DelSession("user")
		return
	}

	c.SetSession("user", *claims)
}

func (c *ApiController) GetSessionUser() *auth.User {
	claims := c.GetSessionClaims()

	if claims == nil {
		return nil
	}

	return &claims.User
}

func (c *ApiController) SetSessionUser(user *auth.User) {
	if user == nil {
		c.DelSession("user")
		return
	}

	claims := c.GetSessionClaims()
	if claims != nil {
		claims.User = *user
		c.SetSessionClaims(claims)
	}
}

func (c *ApiController) GetSessionUsername() string {
	user := c.GetSessionUser()
	if user == nil {
		return ""
	}

	return GetUserName(user)
}

func (c *ApiController) GetSessionUserId() (int64, error) {
	id := c.Controller.GetSession("userId").(int64)
	fmt.Println("id: ", id)

	return id, nil
}
