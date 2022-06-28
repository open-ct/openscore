/*
 * @Author: Junlang
 * @Date: 2021-07-18 00:42:25
 * @LastEditTime: 2021-07-24 20:32:02
 * @LastEditors: Junlang
 * @FilePath: /openscore/routers/router.go
 */
package routers

import (
	"openscore/controllers"

	// "github.com/astaxie/beego"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.ApiController{})
	beego.Router("/api/login", &controllers.ApiController{}, "POST:Login")
	beego.Router("/api/logout", &controllers.ApiController{}, "POST:Logout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "GET:GetAccount")
	// beego.Router("/api/get-users", &controllers.ApiController{}, "GET:GetUsers")
}
