package score

import (
	"openscore/model"
	"time"
)

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
