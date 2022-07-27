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
	"log"

	"github.com/open-ct/openscore/auth"
	"github.com/open-ct/openscore/util"

	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

type TestPaperApiController struct {
	beego.Controller
}

type SupervisorApiController struct {
	beego.Controller
}

type AdminApiController struct {
	beego.Controller
}

func (c *ApiController) GetSessionUser() *auth.Claims {
	s := c.GetSession("user")
	if s == nil {
		return nil
	}
	log.Println(s)

	claims := &auth.Claims{}
	err := util.JsonToStruct(s.(string), claims)
	if err != nil {
		panic(err)
	}

	return claims
}

func (c *ApiController) SetSessionUser(claims *auth.Claims) {
	if claims == nil {
		c.DelSession("user")
		return
	}

	s := util.StructToJson(claims)
	c.SetSession("user", s)
}

func (c *ApiController) GetSessionUsername() string {
	claims := c.GetSessionUser()
	if claims == nil {
		return ""
	}
	return claims.Username
}
