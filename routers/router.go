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

package routers

import (
	"github.com/astaxie/beego"
	"github.com/open-ct/openscore/controllers"
)

func init() {

	api := new(controllers.ApiController)
	beego.Router("/", api)
	beego.Router("/api/signin", api, "post:SignIn")
	beego.Router("/api/signout", api, "post:SignOut")
	beego.Router("/api/get-account", api, "get:GetAccount")

	beego.Router("/openct/login", api, "post:UserLogin")

	// beego.InsertFilter("/openct/marking/score/*", beego.BeforeRouter, filter.AuthScore)
	// beego.InsertFilter("/openct/marking/supervisor/*", beego.BeforeRouter, filter.AuthSupervisor)
	// TODO admin Filter

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
		beego.NSRouter("/list", api, "post:SuperVisorList"),
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
		beego.NSRouter("/writeUserExcel", api, "post:WriteUserExcel"),
		beego.NSRouter("/readExampleExcel", api, "post:ReadExampleExcel"),
		beego.NSRouter("/readExampleExcel", api, "OPTIONS:ReadExampleExcel"),
		beego.NSRouter("/readAnswerExcel", api, "post:ReadAnswerExcel"),
		beego.NSRouter("/readAnswerExcel", api, "OPTIONS:ReadAnswerExcel"),
		// beego.NSRouter("/distribution", api, "post:Distribution"), // Deprecated
		// beego.NSRouter("/distribution/info", api, "post:DistributionInfo"), // Deprecated
		beego.NSRouter("/questionBySubList", api, "post:QuestionBySubList"),
		beego.NSRouter("/insertTopic", api, "post:InsertTopic"),
		beego.NSRouter("/subjectList", api, "post:SubjectList"),
		beego.NSRouter("/topicList", api, "post:TopicList"),
		// beego.NSRouter("/distributionRecord", api, "post:DistributionRecord"), // Deprecated
		beego.NSRouter("/createUser", api, "post:CreateUser"),
		beego.NSRouter("/deleteUser", api, "post:DeleteUser"),
		beego.NSRouter("/updateUser", api, "post:UpdateUser"),
		beego.NSRouter("/listUsers", api, "post:ListUsers"),
		beego.NSRouter("/createSmallQuestion", api, "post:CreateSmallQuestion"),
		beego.NSRouter("/deleteSmallQuestion", api, "post:DeleteSmallQuestion"),
		beego.NSRouter("/updateSmallQuestion", api, "post:UpdateSmallQuestion"),
		beego.NSRouter("/deleteQuestion", api, "post:DeleteQuestion"),
		beego.NSRouter("/updateQuestion", api, "post:UpdateQuestion"),
		beego.NSRouter("/teachingPaperGrouping", api, "post:TeachingPaperGrouping"),
		beego.NSRouter("/listPaperGroups", api, "post:ListPaperGroups"),
		beego.NSRouter("/listTestPapersByQuestionId", api, "post:ListTestPapersByQuestionId"),
		beego.NSRouter("/updateUserQualified", api, "post:UpdateUserQualified"),
		beego.NSRouter("/listGroupGrades", api, "post:ListGroupGrades"),
		beego.NSRouter("/deletePaperFromGroup", api, "post:DeletePaperFromGroup"),
		beego.NSRouter("/listSchools", api, "post:ListSchools"),
		beego.NSRouter("/listTestPaperInfo", api, "post:ListTestPaperInfo"),
	)
	beego.AddNamespace(adminNs)
}
