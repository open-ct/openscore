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
	beego.Router("/", &controllers.TestPaperApiController{})
	beego.Router("/api/login", &controllers.ApiController{}, "post:Login")
	beego.Router("/api/logout", &controllers.ApiController{}, "post:Logout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "get:GetAccount")
	beego.Router("/openct/marking/score/test/display", &controllers.TestPaperApiController{}, "post:Display")
	beego.Router("/openct/marking/score/test/list", &controllers.TestPaperApiController{}, "post:List")
	beego.Router("/openct/marking/score/test/point", &controllers.TestPaperApiController{}, "post:Point")
	beego.Router("/openct/marking/score/test/problem", &controllers.TestPaperApiController{}, "post:Problem")
	beego.Router("/openct/marking/score/test/answer", &controllers.TestPaperApiController{}, "post:Answer")
	beego.Router("/openct/marking/score/test/example/detail", &controllers.TestPaperApiController{}, "post:ExampleDeatil")
	beego.Router("/openct/marking/score/test/example/list", &controllers.TestPaperApiController{}, "post:ExampleList")
	beego.Router("/openct/marking/score/test/review", &controllers.TestPaperApiController{}, "post:Review")

	// beego.Router("/api/get-users", &controllers.ApiController{}, "GET:GetUsers")
}
