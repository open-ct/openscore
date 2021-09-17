package responses

import (
	"openscore/models"
	"time"
)

type SubTopicPlus struct {
	models.SubTopic
	Test_detail_id int64 `json:"test_detail_id"`
}
type TestPaperInfoPlus struct {
	models.TestPaperInfo
	PicCode  string  `json:"picCode"`
}

type TestDisplay struct {
	QuestionId   int64                  `json:"questionId"`
	QuestionName string                 `json:"questionName"`
	TestId       int64                  `json:"testId"`
	SubTopics    []SubTopicPlus         `json:"subTopic"`
	TestInfos    []TestPaperInfoPlus `json:"testInfos"`

}

type TestList struct {
	TestId []int64 `json:"TestIds"`
}

type TestAnswer struct {
	Pics []string `json:"Pics"`

}

type ExampleList struct {
	TestPapers []models.TestPaper `json:"exampleTestPapers"`
}

type TestReview struct {
	TestId    []int64     `json:"testId"`
	Score     []int64     `json:"score"`
	ScoreTime []time.Time `json:"score_time"`
}

type ExampleDetail struct {
	QuestionName string                   `json:"questionName"`
	Test         [][]models.TestPaperInfo `json:"test"`
}
