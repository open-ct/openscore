package controllers

type WriteUserRequest struct {
	SubjectId        int64               `json:"subject_id"`
	SupervisorNumber int                 `json:"supervisor_number"`
	List             []QuestionAndNumber `json:"list"`
}

type QuestionAndNumber struct {
	Id  int64 `json:"id"`
	Num int   `json:"num"`
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AddTopic struct {
	TopicName     string           `json:"topicName"`
	ScoreType     int64            `json:"scoreType"`
	Score         int64            `json:"score"`
	Error         int64            `json:"error"`
	SubjectName   string           `json:"subjectName"`
	TopicDetails  []AddTopicDetail `json:"topicDetails"`
	SelfScoreRate int64            `json:"self_score_rate"`
}

type AddTopicDetail struct {
	TopicDetailName  string `json:"topicDetailName"`
	DetailScoreTypes string `json:"DetailScoreTypes"`
	DetailScore      int64  `json:"detailScore"`
}

type QuestionBySubList struct {
	SubjectName string `json:"subjectName"`
}

type DistributionInfo struct {
	QuestionId int64 `json:"questionId"`
}

type DeleteTest struct {
	QuestionId int64 `json:"questionId"`
}

type Distribution struct {
	QuestionId int64 `json:"questionId"`
	TestNumber int   `json:"testNumber"`
	UserNumber int   `json:"userNumber"`
}

type ReadFile struct {
	PicName string `json:"picName"`
}

type DistributionRecord struct {
	SubjectName string `json:"subjectName"`
}

type TestRequest struct {
	TestId int64 `json:"testId"`
}

type TestPoint struct {
	Scores       string `json:"scores"`
	TestId       int64  `json:"testId"`
	TestDetailId string `json:"testDetailId"`
}

type TestProblem struct {
	ProblemType    int64  `json:"problemType"`
	TestId         int64  `json:"testId"`
	ProblemMessage string `json:"problemMessage"`
}

type ExampleDetail struct {
	ExampleTestId int64 `json:"exampleTestId"`
}

type Question struct {
	QuestionId int64 `json:"questionId"`
}

type SelfScore struct {
	ExaminerId int64 `json:"examinerId"`
}

type SupervisorPoint struct {
	TestId        int64  `json:"testId"`
	TestDetailIds string `json:"testDetailIds"`
	Scores        string `json:"scores"`
}

type ScoreProgress struct {
	Subject string `json:"subject"`
}
