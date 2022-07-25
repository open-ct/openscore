package controllers

import (
	"encoding/gob"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"log"
)

type ApiController struct {
	beego.Controller
}

func init() {
	gob.Register(auth.Claims{})
}

// 用户登录

func (c *ApiController) UserLogin() {
	defer c.ServeJSON()
	var requestBody QuestionBySubList
	var resp Response
	var err error

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	if err != nil {
		log.Println(err)
		resp = Response{"10001", "cannot unmarshal", err}
		c.Data["json"] = resp
		return
	}

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
		// c.DelSession("user")
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
	// user := c.GetSessionUser()
	// if user == nil {
	// 	return 0, errors.New("cant find session info")
	// }

	// u, err := model.GetUserByCasdoorName(user.Name)
	//
	// return u.UserId, err
	return 0, nil
}
