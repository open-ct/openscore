package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"openscore/controllers"
	"openscore/controllers/admin"
	"openscore/controllers/score"
	"openscore/controllers/supervisor"
)

func init() {
	beego.Router("/", &score.TestPaperApiController{})

	api := new(controllers.ApiController)
	beego.Router("/api/login", api, "post:SignIn")
	beego.Router("/api/logout", api, "post:SignOut")
	beego.Router("/api/get-account", api, "get:GetAccount")

	testPaperApi := new(score.TestPaperApiController)
	testNs := beego.NewNamespace("/openct/marking/score",
		beego.NSNamespace("/test",
			beego.NSRouter("/display", testPaperApi, "post:Display"),
			beego.NSRouter("/list", testPaperApi, "post:List"),
			// beego.NSRouter("/point", testPaperApi, "post:SelfMarkPoint"),
			beego.NSRouter("/point", testPaperApi, "post:Point"),
			beego.NSRouter("/problem", testPaperApi, "post:Problem"),
			beego.NSRouter("/answer", testPaperApi, "post:Answer"),
			beego.NSRouter("/example/detail", testPaperApi, "post:ExampleDetail"),
			beego.NSRouter("/example/list", testPaperApi, "post:ExampleList"),
			beego.NSRouter("/review", testPaperApi, "post:Review"),
			beego.NSRouter("/review/point", testPaperApi, "post:ReviewPoint"),
		),
		beego.NSRouter("/self/list", testPaperApi, "post:SelfScoreList"),
	)
	beego.AddNamespace(testNs)

	/**
	  chen :阅卷组长端
	*/
	superApi := new(supervisor.SupervisorApiController)
	superNs := beego.NewNamespace("/openct/marking/supervisor",
		beego.NSRouter("/user/info", superApi, "post:UserInfo"),
		beego.NSRouter("/teacher/monitoring", superApi, "post:TeacherMonitoring"),
		// beego.NSRouter("/question/list", superApi, "post:QuestionList"),
		beego.NSRouter("/score/distribution", superApi, "post:ScoreDistribution"),
		beego.NSRouter("/question/teacher/list", superApi, "post:TeachersByQuestion"),
		beego.NSRouter("/self/score", superApi, "post:SelfScore"),
		beego.NSRouter("/average/score", superApi, "post:AverageScore"),
		beego.NSRouter("/problem/list", superApi, "post:ProblemTest"),
		beego.NSRouter("/arbitrament/list", superApi, "post:ArbitramentTest"),
		beego.NSRouter("/score/progress", superApi, "post:ScoreProgress"),
		beego.NSRouter("/point", superApi, "post:SupervisorPoint"),
		beego.NSRouter("/arbitrament/unmark/list", superApi, "post:ArbitramentUnmarkList"),
		beego.NSRouter("/selfMark/list", superApi, "post:SelfMarkList"),
		beego.NSRouter("/selfMark/unmark/list", superApi, "post:SelfUnmarkList"),
		beego.NSRouter("/problem/unmark/list", superApi, "post:ProblemUnmarkList"),
		beego.NSRouter("/score/deviation", superApi, "post:ScoreDeviation"),
		beego.NSRouter("/writeScoreExcel", superApi, "post:WriteScoreExcel"),
	)
	beego.AddNamespace(superNs)

	/**
	  chen :管理员端
	*/
	adminApi := new(admin.AdminApiController)
	adminNs := beego.NewNamespace("/openct/marking/admin",
		beego.NSRouter("/readExcel", adminApi, "post:ReadExcel"),
		beego.NSRouter("/readExcel", adminApi, "OPTIONS:ReadExcel"),
		// beego.NSRouter("/uploadPic", adminApi, "post:UploadPic"),
		// beego.NSRouter("/readUserExcel", adminApi, "post:ReadUserExcel"),
		beego.NSRouter("/readExampleExcel", adminApi, "post:ReadExampleExcel"),
		beego.NSRouter("/readExampleExcel", adminApi, "OPTIONS:ReadExampleExcel"),
		beego.NSRouter("/readAnswerExcel", adminApi, "post:ReadAnswerExcel"),
		beego.NSRouter("/readAnswerExcel", adminApi, "OPTIONS:ReadAnswerExcel"),
		beego.NSRouter("/distribution", adminApi, "post:Distribution"),
		beego.NSRouter("/distribution/info", adminApi, "post:DistributionInfo"),
		beego.NSRouter("/questionBySubList", adminApi, "post:QuestionBySubList"),
		beego.NSRouter("/insertTopic", adminApi, "post:InsertTopic"),
		beego.NSRouter("/subjectList", adminApi, "post:SubjectList"),
		beego.NSRouter("/topicList", adminApi, "post:TopicList"),
		beego.NSRouter("/DistributionRecord", adminApi, "post:DistributionRecord"),
		beego.NSRouter("/img", adminApi, "post:Pic"),
	)
	beego.AddNamespace(adminNs)
}
