package controllers

import (
	"log"
	"openscore/auth"
	"openscore/util"

	beego "github.com/beego/beego/v2/server/web"
	// "github.com/astaxie/beego"
)

type ApiController struct {
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
