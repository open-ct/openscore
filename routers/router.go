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

	/**
	  chen :阅卷组长端
	*/
	beego.Router("/openct/marking/supervisor/question/list", &controllers.SupervisorApiController{}, "post:QuestionList")
	beego.Router("/openct/marking/supervisor/user/info", &controllers.SupervisorApiController{}, "post:UserInfo")
	beego.Router("/openct/marking/supervisor/teacher/monitoring", &controllers.SupervisorApiController{}, "post:TeacherMonitoring")
	beego.Router("/openct/marking/supervisor/score/distribution", &controllers.SupervisorApiController{}, "post:ScoreDistribution")
	beego.Router("/openct/marking/supervisor/question/teacher/list", &controllers.SupervisorApiController{}, "post:TeachersByQuestion")
	beego.Router("/openct/marking/supervisor/self/score", &controllers.SupervisorApiController{}, "post:SelfScore")
	beego.Router("/openct/marking/supervisor/average/score", &controllers.SupervisorApiController{}, "post:AverageScore")
	beego.Router("/openct/marking/supervisor/problem/list", &controllers.SupervisorApiController{}, "post:ProblemTest")
	beego.Router("/openct/marking/supervisor/arbitrament/list", &controllers.SupervisorApiController{}, "post:ArbitramentTest")
	beego.Router("/openct/marking/supervisor/score/progress", &controllers.SupervisorApiController{}, "post:ScoreProgress")
	beego.Router("/openct/marking/supervisor/point", &controllers.SupervisorApiController{}, "post:SupervisorPoint")
	beego.Router("/openct/marking/supervisor/arbitrament/unmark/list", &controllers.SupervisorApiController{}, "post:ArbitramentUnmarkList")
	beego.Router("/openct/marking/supervisor/problem/unmark/list", &controllers.SupervisorApiController{}, "post:ProblemUnmarkList")
	beego.Router("/openct/marking/supervisor/score/deviation", &controllers.SupervisorApiController{}, "post:ScoreDeviation")

	/**
	  chen :管理员端
	*/
	//beego.Router("/openct/marking/admin/uploadPic",&controllers.AdminApiController{},"post:UploadPic")
	beego.Router("/openct/marking/admin/readExcel",&controllers.AdminApiController{},"post:ReadExcel")
	beego.Router("/openct/marking/admin/distribution",&controllers.AdminApiController{},"post:Distribution")
	beego.Router("/openct/marking/admin/distribution/info",&controllers.AdminApiController{},"post:DistributionInfo")
	beego.Router("/openct/marking/admin/questionBySubList",&controllers.AdminApiController{},"post:QuestionBySubList")
	beego.Router("/openct/marking/admin/insertTopic",&controllers.AdminApiController{},"post:InsertTopic")
	beego.Router("/openct/marking/admin/subjectList",&controllers.AdminApiController{},"post:SubjectList")

	beego.Router("/openct/marking/admin/img",&controllers.AdminApiController{},"post:Pic")

}
