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
	beego.Router("/", &controllers.TestPaperApiController{})
	beego.Router("/api/login", &controllers.ApiController{}, "post:Login")
	beego.Router("/api/logout", &controllers.ApiController{}, "post:Logout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "get:GetAccount")
	beego.Router("/openct/marking/score/test/display", &controllers.TestPaperApiController{}, "post:Display")
	beego.Router("/openct/marking/score/test/list", &controllers.TestPaperApiController{}, "post:List")
	beego.Router("/openct/marking/score/self/list", &controllers.TestPaperApiController{}, "post:SelfScoreList")
	//beego.Router("/openct/marking/score/self/point", &controllers.TestPaperApiController{}, "post:SelfMarkPoint")
	beego.Router("/openct/marking/score/test/point", &controllers.TestPaperApiController{}, "post:Point")
	beego.Router("/openct/marking/score/test/problem", &controllers.TestPaperApiController{}, "post:Problem")
	beego.Router("/openct/marking/score/test/answer", &controllers.TestPaperApiController{}, "post:Answer")
	beego.Router("/openct/marking/score/test/example/detail", &controllers.TestPaperApiController{}, "post:ExampleDetail")
	beego.Router("/openct/marking/score/test/example/list", &controllers.TestPaperApiController{}, "post:ExampleList")
	beego.Router("/openct/marking/score/test/review", &controllers.TestPaperApiController{}, "post:Review")
	beego.Router("/openct/marking/score/test/review/point", &controllers.TestPaperApiController{}, "post:ReviewPoint")
	// beego.Router("/api/get-users", &controllers.ApiController{}, "GET:GetUsers")

	/**
	  chen :阅卷组长端
	*/
	//beego.Router("/openct/marking/supervisor/question/list", &controllers.SupervisorApiController{}, "post:QuestionList")
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
	beego.Router("/openct/marking/supervisor/selfMark/list", &controllers.SupervisorApiController{}, "post:SelfMarkList")
	beego.Router("/openct/marking/supervisor/selfMark/unmark/list", &controllers.SupervisorApiController{}, "post:SelfUnmarkList")
	beego.Router("/openct/marking/supervisor/problem/unmark/list", &controllers.SupervisorApiController{}, "post:ProblemUnmarkList")
	beego.Router("/openct/marking/supervisor/score/deviation", &controllers.SupervisorApiController{}, "post:ScoreDeviation")

	/**
	  chen :管理员端
	*/
	//beego.Router("/openct/marking/admin/uploadPic",&controllers.AdminApiController{},"post:UploadPic")
	beego.Router("/openct/marking/admin/readExcel", &controllers.AdminApiController{}, "post:ReadExcel")
	beego.Router("/openct/marking/admin/readExcel", &controllers.AdminApiController{}, "OPTIONS:ReadExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &controllers.AdminApiController{}, "post:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &controllers.AdminApiController{}, "OPTIONS:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &controllers.AdminApiController{}, "post:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &controllers.AdminApiController{}, "OPTIONS:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/distribution", &controllers.AdminApiController{}, "post:Distribution")
	beego.Router("/openct/marking/admin/distribution/info", &controllers.AdminApiController{}, "post:DistributionInfo")
	beego.Router("/openct/marking/admin/questionBySubList", &controllers.AdminApiController{}, "post:QuestionBySubList")
	beego.Router("/openct/marking/admin/insertTopic", &controllers.AdminApiController{}, "post:InsertTopic")
	beego.Router("/openct/marking/admin/subjectList", &controllers.AdminApiController{}, "post:SubjectList")
	beego.Router("/openct/marking/admin/topicList", &controllers.AdminApiController{}, "post:TopicList")
	beego.Router("/openct/marking/admin/DistributionRecord", &controllers.AdminApiController{}, "post:DistributionRecord")

	beego.Router("/openct/marking/admin/img", &controllers.AdminApiController{}, "post:Pic")

}
