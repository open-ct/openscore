package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"openscore/controllers"
	"openscore/controllers/admin"
	"openscore/controllers/score"
	"openscore/controllers/supervisor"
)

func init() {
	// var scor score.TestPaperApiController
	// beego.Get("/", scor.Display)
	beego.Router("/", &score.TestPaperApiController{})
	beego.Router("/api/login", &controllers.ApiController{}, "post:Login")
	beego.Router("/api/logout", &controllers.ApiController{}, "post:Logout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "get:GetAccount")
	beego.Router("/openct/marking/score/test/display", &score.TestPaperApiController{}, "post:Display")
	beego.Router("/openct/marking/score/test/list", &score.TestPaperApiController{}, "post:List")
	beego.Router("/openct/marking/score/self/list", &score.TestPaperApiController{}, "post:SelfScoreList")
	// beego.Router("/openct/marking/score/self/point", &score.TestPaperApiController{}, "post:SelfMarkPoint")
	beego.Router("/openct/marking/score/test/point", &score.TestPaperApiController{}, "post:Point")
	beego.Router("/openct/marking/score/test/problem", &score.TestPaperApiController{}, "post:Problem")
	beego.Router("/openct/marking/score/test/answer", &score.TestPaperApiController{}, "post:Answer")
	beego.Router("/openct/marking/score/test/example/detail", &score.TestPaperApiController{}, "post:ExampleDetail")
	beego.Router("/openct/marking/score/test/example/list", &score.TestPaperApiController{}, "post:ExampleList")
	beego.Router("/openct/marking/score/test/review", &score.TestPaperApiController{}, "post:Review")
	beego.Router("/openct/marking/score/test/review/point", &score.TestPaperApiController{}, "post:ReviewPoint")
	// beego.Router("/api/get-users", &score.ApiController{}, "GET:GetUsers")

	/**
	  chen :阅卷组长端
	*/
	// beego.Router("/openct/marking/supervisor/question/list", &supervisor.SupervisorApiController{}, "post:QuestionList")
	beego.Router("/openct/marking/supervisor/user/info", &supervisor.SupervisorApiController{}, "post:UserInfo")
	beego.Router("/openct/marking/supervisor/teacher/monitoring", &supervisor.SupervisorApiController{}, "post:TeacherMonitoring")
	beego.Router("/openct/marking/supervisor/score/distribution", &supervisor.SupervisorApiController{}, "post:ScoreDistribution")
	beego.Router("/openct/marking/supervisor/question/teacher/list", &supervisor.SupervisorApiController{}, "post:TeachersByQuestion")
	beego.Router("/openct/marking/supervisor/self/score", &supervisor.SupervisorApiController{}, "post:SelfScore")
	beego.Router("/openct/marking/supervisor/average/score", &supervisor.SupervisorApiController{}, "post:AverageScore")
	beego.Router("/openct/marking/supervisor/problem/list", &supervisor.SupervisorApiController{}, "post:ProblemTest")
	beego.Router("/openct/marking/supervisor/arbitrament/list", &supervisor.SupervisorApiController{}, "post:ArbitramentTest")
	beego.Router("/openct/marking/supervisor/score/progress", &supervisor.SupervisorApiController{}, "post:ScoreProgress")
	beego.Router("/openct/marking/supervisor/point", &supervisor.SupervisorApiController{}, "post:SupervisorPoint")
	beego.Router("/openct/marking/supervisor/arbitrament/unmark/list", &supervisor.SupervisorApiController{}, "post:ArbitramentUnmarkList")
	beego.Router("/openct/marking/supervisor/selfMark/list", &supervisor.SupervisorApiController{}, "post:SelfMarkList")
	beego.Router("/openct/marking/supervisor/selfMark/unmark/list", &supervisor.SupervisorApiController{}, "post:SelfUnmarkList")
	beego.Router("/openct/marking/supervisor/problem/unmark/list", &supervisor.SupervisorApiController{}, "post:ProblemUnmarkList")
	beego.Router("/openct/marking/supervisor/score/deviation", &supervisor.SupervisorApiController{}, "post:ScoreDeviation")

	/**
	  chen :管理员端
	*/
	// beego.Router("/openct/marking/admin/uploadPic",&admin.AdminApiController{},"post:UploadPic")
	beego.Router("/openct/marking/admin/readExcel", &admin.AdminApiController{}, "post:ReadExcel")
	beego.Router("/openct/marking/admin/readExcel", &admin.AdminApiController{}, "OPTIONS:ReadExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &admin.AdminApiController{}, "post:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &admin.AdminApiController{}, "OPTIONS:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &admin.AdminApiController{}, "post:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &admin.AdminApiController{}, "OPTIONS:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/distribution", &admin.AdminApiController{}, "post:Distribution")
	beego.Router("/openct/marking/admin/distribution/info", &admin.AdminApiController{}, "post:DistributionInfo")
	beego.Router("/openct/marking/admin/questionBySubList", &admin.AdminApiController{}, "post:QuestionBySubList")
	beego.Router("/openct/marking/admin/insertTopic", &admin.AdminApiController{}, "post:InsertTopic")
	beego.Router("/openct/marking/admin/subjectList", &admin.AdminApiController{}, "post:SubjectList")
	beego.Router("/openct/marking/admin/topicList", &admin.AdminApiController{}, "post:TopicList")
	beego.Router("/openct/marking/admin/DistributionRecord", &admin.AdminApiController{}, "post:DistributionRecord")

	beego.Router("/openct/marking/admin/img", &admin.AdminApiController{}, "post:Pic")

}
