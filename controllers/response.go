package controllers

import (
	"github.com/open-ct/openscore/model"
	"time"
)

type AddTopicVO struct {
	QuestionId        int64
	QuestionDetailIds []AddTopicDetailVO
}

type AddTopicDetailVO struct {
	QuestionDetailId int64
}

type SubjectListVO struct {
	SubjectId   int64
	SubjectName string
}

type QuestionBySubListVO struct {
	QuestionId   int64
	QuestionName string
}

type DistributionInfoVO struct {
	ImportTestNumber int64
	LeftTestNumber   int
	OnlineNumber     int64
	ScoreType        int64
}

type TopicVO struct {
	TopicId        int64
	SubjectName    string
	TopicName      string
	Score          int64
	StandardError  int64
	ScoreType      int64
	ImportTime     time.Time
	SubTopicVOList []SubTopicVO
}

type SubTopicVO struct {
	SubTopicId        int64
	SubTopicName      string
	Score             int64
	ScoreDistribution string
}

type DistributionRecordVO struct {
	TopicId                int64
	TopicName              string
	ImportNumber           int64
	DistributionTestNumber int64
	DistributionUserNumber int64
}

type SubTopicPlus struct {
	model.SubTopic
	TestDetailId int64 `json:"test_detail_id"`
}
type TestPaperInfoPlus struct {
	model.TestPaperInfo
	PicCode string `json:"picCode"`
}

type TestDisplayResponse struct {
	QuestionId   int64               `json:"questionId"`
	QuestionName string              `json:"questionName"`
	TestId       int64               `json:"testId"`
	SubTopics    []SubTopicPlus      `json:"subTopic"`
	TestInfos    []TestPaperInfoPlus `json:"testInfos"`
}

type TestListResponse struct {
	TestId []int64 `json:"TestIds"`
}

type TestAnswerResponse struct {
	Pics []string `json:"Pics"`
}

type ExampleListResponse struct {
	TestPapers []model.TestPaper `json:"exampleTestPapers"`
}

type TestReviewResponse struct {
	TestId    []int64     `json:"testId"`
	Score     []int64     `json:"score"`
	ScoreTime []time.Time `json:"score_time"`
}

type ExampleDetailResponse struct {
	QuestionName string                  `json:"questionName"`
	Test         [][]model.TestPaperInfo `json:"test"`
}

type QuestionListVO struct {
	QuestionId   int64
	QuestionName string
}

type UserInfoVO struct {
	UserName    string
	SubjectName string
}

type TeacherMonitoringVO struct {
	UserId                 int64
	UserName               string
	TestDistributionNumber int64
	TestSuccessNumber      float64
	TestRemainingNumber    int64
	TestProblemNumber      int64
	MarkingSpeed           float64
	PredictTime            float64
	AverageScore           float64
	Validity               float64
	StandardDeviation      float64
	IsOnline               bool
}

type ScoreDistributionVO struct {
	Score int64
	Rate  float64
}

type TeacherVO struct {
	UserId   int64
	UserName string
}

type SelfScoreRecordVO struct {
	TestId        int64
	Score         int64
	SelfScore     int64
	IsQualified   int64
	Error         float64
	StandardError int64
}
type ScoreAverageVO struct {
	UserId   int64
	UserName string
	Average  float64
}
type ProblemUnderCorrectedPaperVO struct {
	TestId       int64
	ExaminerId   int64
	ExaminerName string
	ProblemType  int64
	ProblemMes   string
}
type ProblemUnmarkListVO struct {
	TestId int64
}
type SelfUnmarkListVO struct {
	TestId int64
}

type ArbitramentTestVO struct {
	TestId              int64
	ExaminerFirstId     int64
	ExaminerFirstName   string
	ExaminerFirstScore  int64
	ExaminerSecondId    int64
	ExaminerSecondName  string
	ExaminerSecondScore int64
	ExaminerThirdId     int64
	ExaminerThirdName   string
	ExaminerThirdScore  int64
	PracticeError       int64
	StandardError       int64
}

type ArbitramentUnmarkListVO struct {
	TestId int64
}
type SelfMarkListVO struct {
	TestId        int64
	Score         int64
	SelfScore     int64
	Error         float64
	StandardError int64
	Userid        int64
	Name          string
}

type ScoreProgressVO struct {
	// 问题id 问题名 导入试卷数
	QuestionId   int64
	QuestionName string
	ImportNumber int64

	// 在线人数 ，分配人数 平均分 平均速度  在线预计时间 ，预计时间 ,自评指数
	OnlineUserNumber       int64
	DistributionUserNumber int64

	OnlineAverageScore float64
	AverageScore       float64
	OnlineAverageSpeed float64
	AverageSpeed       float64
	OnlinePredictTime  float64
	PredictTime        float64

	SelfScoreRate int64

	// 完成阅卷数 完成率 未完成数 未完成率  是否全部完成
	FinishNumber     int64
	FinishRate       float64
	UnfinishedNumber float64
	UnfinishedRate   float64
	IsAllFinished    string

	FirstFinishedNumber   int64
	FirstFinishedRate     float64
	FirstUnfinishedNumber float64
	FirstUnfinishedRate   float64
	IsFirstFinished       string

	SecondFinishedNumber   int64
	SecondFinishedRate     float64
	SecondUnfinishedNumber float64
	SecondUnfinishedRate   float64
	IsSecondFinished       string

	ThirdFinishedNumber   int64
	ThirdFinishedRate     float64
	ThirdUnfinishedNumber float64
	ThirdUnfinishedRate   float64
	IsThirdFinished       string

	// 仲裁卷生产数量 生产率   完成数，完成率 未完成数 未完成率  是否全部完成
	ArbitramentNumber           int64
	ArbitramentRate             float64
	ArbitramentFinishedNumber   int64
	ArbitramentFinishedRate     float64
	ArbitramentUnfinishedNumber int64
	ArbitramentUnfinishedRate   float64
	IsArbitramentFinished       string

	ProblemNumber           int64
	ProblemRate             float64
	ProblemFinishedNumber   int64
	ProblemFinishedRate     float64
	ProblemUnfinishedNumber int64
	ProblemUnfinishedRate   float64
	IsProblemFinished       string
}

type ScoreDeviationVO struct {
	UserId         int64
	UserName       string
	DeviationScore float64
}
