package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"log"
	"openscore/controllers"
	"openscore/pkg/token"
)

func init() {

	api := new(controllers.ApiController)
	beego.Router("/", api)
	beego.Router("/api/login", api, "post:SignIn")
	beego.Router("/api/logout", api, "post:SignOut")
	beego.Router("/api/get-account", api, "get:GetAccount")

	beego.Router("/openct/login", api, "post:Login")

	FilterUser := func(ctx *context.Context) {
		authorization := ctx.Input.Header("Authorization")
		if len(authorization) == 0 {
			api.ResponseError("cant get Authorization")
			ctx.Redirect(302, "/login")
		}
		res, err := token.ResolveToken(authorization)
		if err != nil {
			log.Println(err)
			api.ResponseError("cant resolve Authorization", err)
		}
		ctx.Input.SetData("userId", res.Id)
		ctx.Input.SetData("typeId", res.TypeId)
	}

	beego.InsertFilter("/openct/marking", beego.BeforeRouter, FilterUser)

	testNs := beego.NewNamespace("/openct/marking/score",
		beego.NSNamespace("/test",
			beego.NSRouter("/display", api, "post:Display"),
			beego.NSRouter("/list", api, "post:List"),
			// beego.NSRouter("/point", api, "post:SelfMarkPoint"),
			beego.NSRouter("/point", api, "post:Point"),
			beego.NSRouter("/problem", api, "post:Problem"),
			beego.NSRouter("/answer", api, "post:Answer"),
			beego.NSRouter("/example/detail", api, "post:ExampleDetail"),
			beego.NSRouter("/example/list", api, "post:ExampleList"),
			beego.NSRouter("/review", api, "post:Review"),
			beego.NSRouter("/review/point", api, "post:ReviewPoint"),
		),
		beego.NSRouter("/self/list", api, "post:SelfScoreList"),
	)
	beego.AddNamespace(testNs)

	/**
	  chen :阅卷组长端
	*/
	superNs := beego.NewNamespace("/openct/marking/supervisor",
		beego.NSRouter("/user/info", api, "post:UserInfo"),
		beego.NSRouter("/teacher/monitoring", api, "post:TeacherMonitoring"),
		// beego.NSRouter("/question/list", api, "post:QuestionList"),
		beego.NSRouter("/score/distribution", api, "post:ScoreDistribution"),
		beego.NSRouter("/question/teacher/list", api, "post:TeachersByQuestion"),
		beego.NSRouter("/self/score", api, "post:SelfScore"),
		beego.NSRouter("/average/score", api, "post:AverageScore"),
		beego.NSRouter("/problem/list", api, "post:ProblemTest"),
		beego.NSRouter("/arbitrament/list", api, "post:ArbitramentTest"),
		beego.NSRouter("/score/progress", api, "post:ScoreProgress"),
		beego.NSRouter("/point", api, "post:SupervisorPoint"),
		beego.NSRouter("/arbitrament/unmark/list", api, "post:ArbitramentUnmarkList"),
		beego.NSRouter("/selfMark/list", api, "post:SelfMarkList"),
		beego.NSRouter("/selfMark/unmark/list", api, "post:SelfUnmarkList"),
		beego.NSRouter("/problem/unmark/list", api, "post:ProblemUnmarkList"),
		beego.NSRouter("/score/deviation", api, "post:ScoreDeviation"),
		beego.NSRouter("/writeScoreExcel", api, "post:WriteScoreExcel"),
	)
	beego.AddNamespace(superNs)

	/**
	  chen :管理员端
	*/
	adminNs := beego.NewNamespace("/openct/marking/admin",
		beego.NSRouter("/readExcel", api, "post:ReadExcel"),
		beego.NSRouter("/readExcel", api, "OPTIONS:ReadExcel"),
		// beego.NSRouter("/uploadPic", api, "post:UploadPic"),
		beego.NSRouter("/readExampleExcel", api, "post:ReadExampleExcel"),
		beego.NSRouter("/readExampleExcel", api, "OPTIONS:ReadExampleExcel"),
		beego.NSRouter("/readAnswerExcel", api, "post:ReadAnswerExcel"),
		beego.NSRouter("/readAnswerExcel", api, "OPTIONS:ReadAnswerExcel"),
		beego.NSRouter("/distribution", api, "post:Distribution"),
		beego.NSRouter("/distribution/info", api, "post:DistributionInfo"),
		beego.NSRouter("/questionBySubList", api, "post:QuestionBySubList"),
		beego.NSRouter("/insertTopic", api, "post:InsertTopic"),
		beego.NSRouter("/subjectList", api, "post:SubjectList"),
		beego.NSRouter("/topicList", api, "post:TopicList"),
		beego.NSRouter("/DistributionRecord", api, "post:DistributionRecord"),
		beego.NSRouter("/img", api, "post:Pic"),
	)
	beego.AddNamespace(adminNs)
}
