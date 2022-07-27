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
	beego.Router("/", &controllers.ApiController{})
	beego.Router("/api/login", &controllers.ApiController{}, "post:Login")
	beego.Router("/api/logout", &controllers.ApiController{}, "post:Logout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "get:GetAccount")
	beego.Router("/openct/marking/score/test/display", &controllers.ApiController{}, "post:Display")
	beego.Router("/openct/marking/score/test/list", &controllers.ApiController{}, "post:List")
	beego.Router("/openct/marking/score/self/list", &controllers.ApiController{}, "post:SelfScoreList")
	//beego.Router("/openct/marking/score/self/point", &controllers.ApiController{}, "post:SelfMarkPoint")
	beego.Router("/openct/marking/score/test/point", &controllers.ApiController{}, "post:Point")
	beego.Router("/openct/marking/score/test/problem", &controllers.ApiController{}, "post:Problem")
	beego.Router("/openct/marking/score/test/answer", &controllers.ApiController{}, "post:Answer")
	beego.Router("/openct/marking/score/test/example/detail", &controllers.ApiController{}, "post:ExampleDetail")
	beego.Router("/openct/marking/score/test/example/list", &controllers.ApiController{}, "post:ExampleList")
	beego.Router("/openct/marking/score/test/review", &controllers.ApiController{}, "post:Review")
	beego.Router("/openct/marking/score/test/review/point", &controllers.ApiController{}, "post:ReviewPoint")
	// beego.Router("/api/get-users", &controllers.ApiController{}, "GET:GetUsers")

	/**
	  chen :阅卷组长端
	*/
	//beego.Router("/openct/marking/supervisor/question/list", &controllers.ApiController{}, "post:QuestionList")
	beego.Router("/openct/marking/supervisor/user/info", &controllers.ApiController{}, "post:UserInfo")
	beego.Router("/openct/marking/supervisor/teacher/monitoring", &controllers.ApiController{}, "post:TeacherMonitoring")
	beego.Router("/openct/marking/supervisor/score/distribution", &controllers.ApiController{}, "post:ScoreDistribution")
	beego.Router("/openct/marking/supervisor/question/teacher/list", &controllers.ApiController{}, "post:TeachersByQuestion")
	beego.Router("/openct/marking/supervisor/self/score", &controllers.ApiController{}, "post:SelfScore")
	beego.Router("/openct/marking/supervisor/average/score", &controllers.ApiController{}, "post:AverageScore")
	beego.Router("/openct/marking/supervisor/problem/list", &controllers.ApiController{}, "post:ProblemTest")
	beego.Router("/openct/marking/supervisor/arbitrament/list", &controllers.ApiController{}, "post:ArbitramentTest")
	beego.Router("/openct/marking/supervisor/score/progress", &controllers.ApiController{}, "post:ScoreProgress")
	beego.Router("/openct/marking/supervisor/point", &controllers.ApiController{}, "post:SupervisorPoint")
	beego.Router("/openct/marking/supervisor/arbitrament/unmark/list", &controllers.ApiController{}, "post:ArbitramentUnmarkList")
	beego.Router("/openct/marking/supervisor/selfMark/list", &controllers.ApiController{}, "post:SelfMarkList")
	beego.Router("/openct/marking/supervisor/selfMark/unmark/list", &controllers.ApiController{}, "post:SelfUnmarkList")
	beego.Router("/openct/marking/supervisor/problem/unmark/list", &controllers.ApiController{}, "post:ProblemUnmarkList")
	beego.Router("/openct/marking/supervisor/score/deviation", &controllers.ApiController{}, "post:ScoreDeviation")

	/**
	  chen :管理员端
	*/
	//beego.Router("/openct/marking/admin/uploadPic",&controllers.ApiController{},"post:UploadPic")
	beego.Router("/openct/marking/admin/readExcel", &controllers.ApiController{}, "post:ReadExcel")
	beego.Router("/openct/marking/admin/readExcel", &controllers.ApiController{}, "OPTIONS:ReadExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &controllers.ApiController{}, "post:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readExampleExcel", &controllers.ApiController{}, "OPTIONS:ReadExampleExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &controllers.ApiController{}, "post:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/readAnswerExcel", &controllers.ApiController{}, "OPTIONS:ReadAnswerExcel")
	beego.Router("/openct/marking/admin/distribution", &controllers.ApiController{}, "post:Distribution")
	beego.Router("/openct/marking/admin/distribution/info", &controllers.ApiController{}, "post:DistributionInfo")
	beego.Router("/openct/marking/admin/questionBySubList", &controllers.ApiController{}, "post:QuestionBySubList")
	beego.Router("/openct/marking/admin/insertTopic", &controllers.ApiController{}, "post:InsertTopic")
	beego.Router("/openct/marking/admin/subjectList", &controllers.ApiController{}, "post:SubjectList")
	beego.Router("/openct/marking/admin/topicList", &controllers.ApiController{}, "post:TopicList")
	beego.Router("/openct/marking/admin/DistributionRecord", &controllers.ApiController{}, "post:DistributionRecord")

	beego.Router("/openct/marking/admin/img", &controllers.ApiController{}, "post:Pic")

}
