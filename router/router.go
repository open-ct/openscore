package router

import (
	// "github.com/astaxie/beego"
	beego "github.com/beego/beego/v2/server/web"
	"openscore/controller"
)

func init() {
	beego.Router("/", &controller.TestPaperApiController{})
	beego.Router("/api/login", &controller.ApiController{}, "post:Login")
	beego.Router("/api/logout", &controller.ApiController{}, "post:Logout")
	beego.Router("/api/get-account", &controller.ApiController{}, "get:GetAccount")
	beego.Router("/openct/marking/score/test/display", &controller.TestPaperApiController{}, "post:Display")
	beego.Router("/openct/marking/score/test/list", &controller.TestPaperApiController{}, "post:List")
	beego.Router("/openct/marking/score/self/list", &controller.TestPaperApiController{}, "post:SelfScoreList")
	// beego.Router("/openct/marking/score/self/point", &controllers.TestPaperApiController{}, "post:SelfMarkPoint")
	beego.Router("/openct/marking/score/test/point", &controller.TestPaperApiController{}, "post:Point")
	beego.Router("/openct/marking/score/test/problem", &controller.TestPaperApiController{}, "post:Problem")
	beego.Router("/openct/marking/score/test/answer", &controller.TestPaperApiController{}, "post:Answer")
	beego.Router("/openct/marking/score/test/example/detail", &controller.TestPaperApiController{}, "post:ExampleDetail")
	beego.Router("/openct/marking/score/test/example/list", &controller.TestPaperApiController{}, "post:ExampleList")
	beego.Router("/openct/marking/score/test/review", &controller.TestPaperApiController{}, "post:Review")
	beego.Router("/openct/marking/score/test/review/point", &controller.TestPaperApiController{}, "post:ReviewPoint")
	// beego.Router("/api/get-users", &controllers.ApiController{}, "GET:GetUsers")

	/**
	  chen :阅卷组长端
	*/
	// beego.Router("/openct/marking/supervisor/question/list", &controllers.SupervisorApiController{}, "post:QuestionList")
	beego.Router("/openct/marking/supervisor/user/info", &controller.SupervisorApiController{}, "post:UserInfo")
	beego.Router("/openct/marking/supervisor/teacher/monitoring", &controller.SupervisorApiController{}, "post:TeacherMonitoring")
	beego.Router("/openct/marking/supervisor/score/distribution", &controller.SupervisorApiController{}, "post:ScoreDistribution")
	beego.Router("/openct/marking/supervisor/question/teacher/list", &controller.SupervisorApiController{}, "post:TeachersByQuestion")
	beego.Router("/openct/marking/supervisor/self/score", &controller.SupervisorApiController{}, "post:SelfScore")
	beego.Router("/openct/marking/supervisor/average/score", &controller.SupervisorApiController{}, "post:AverageScore")
	beego.Router("/openct/marking/supervisor/problem/list", &controller.SupervisorApiController{}, "post:ProblemTest")
	beego.Router("/openct/marking/supervisor/arbitrament/list", &controller.SupervisorApiController{}, "post:ArbitramentTest")
	beego.Router("/openct/marking/supervisor/score/progress", &controller.SupervisorApiController{}, "post:ScoreProgress")
	beego.Router("/openct/marking/supervisor/point", &controller.SupervisorApiController{}, "post:SupervisorPoint")
	beego.Router("/openct/marking/supervisor/arbitrament/unmark/list", &controller.SupervisorApiController{}, "post:ArbitramentUnmarkList")
	beego.Router("/openct/marking/supervisor/selfMark/list", &controller.SupervisorApiController{}, "post:SelfMarkList")
	beego.Router("/openct/marking/supervisor/selfMark/unmark/list", &controller.SupervisorApiController{}, "post:SelfUnmarkList")
	beego.Router("/openct/marking/supervisor/problem/unmark/list", &controller.SupervisorApiController{}, "post:ProblemUnmarkList")
	beego.Router("/openct/marking/supervisor/score/deviation", &controller.SupervisorApiController{}, "post:ScoreDeviation")

	/**
	  chen :管理员端
	*/
	// beego.Router("/openct/marking/admin/uploadPic",&controllers.AdminApiController{},"post:UploadPic")
	beego.Router("/openct/marking/admin/readExcel", &controller.AdminApiController{}, "post:ReadExcel")
	beego.Router("/openct/marking/admin/readExcel", &controller.AdminApiController{}, "OPTIONS:ReadExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &controller.AdminApiController{}, "post:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &controller.AdminApiController{}, "OPTIONS:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &controller.AdminApiController{}, "post:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &controller.AdminApiController{}, "OPTIONS:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/distribution", &controller.AdminApiController{}, "post:Distribution")
	beego.Router("/openct/marking/admin/distribution/info", &controller.AdminApiController{}, "post:DistributionInfo")
	beego.Router("/openct/marking/admin/questionBySubList", &controller.AdminApiController{}, "post:QuestionBySubList")
	beego.Router("/openct/marking/admin/insertTopic", &controller.AdminApiController{}, "post:InsertTopic")
	beego.Router("/openct/marking/admin/subjectList", &controller.AdminApiController{}, "post:SubjectList")
	beego.Router("/openct/marking/admin/topicList", &controller.AdminApiController{}, "post:TopicList")
	beego.Router("/openct/marking/admin/DistributionRecord", &controller.AdminApiController{}, "post:DistributionRecord")

	beego.Router("/openct/marking/admin/img", &controller.AdminApiController{}, "post:Pic")

}
