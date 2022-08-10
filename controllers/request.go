package controllers

type UpdateSmallQuestionRequest struct {
	QuestionDetailId    int64  `json:"question_detail_id"`
	QuestionDetailName  string `json:"question_detail_name"`
	QuestionDetailScore int64  `json:"question_detail_score"`
	ScoreType           string `json:"score_type"`
}

type CreateSmallQuestionRequest struct {
	QuestionDetailName  string `json:"question_detail_name"`
	QuestionId          int64  `json:"question_id"`
	QuestionDetailScore int64  `json:"question_detail_score"`
	ScoreType           string `json:"score_type"`
}

type DeleteSmallQuestionRequest struct {
	QuestionDetailId int64 `json:"question_detail_id"`
}

type DeleteQuestionRequest struct {
	QuestionId int64 `json:"question_id"`
}

type UpdateQuestionRequest struct {
	QuestionId    int64  `json:"question_id"`
	QuestionName  string `json:"question_name"`
	StandardError int64  `json:"standard_error"`
	QuestionScore int64  `json:"question_score"`
	ScoreType     int64  `json:"score_type"`
}

type UpdateUserRequest struct {
	Account     string `json:"account"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	SubjectName string `json:"subject_name"`
	UserType    int64  `json:"user_type"`
}

type DeleteUserRequest struct {
	Account string `json:"account"`
}

type CreateUserRequest struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	SubjectName string `json:"subject_name"`
	QuestionId  int64  `json:"question_id"`
	UserType    int64  `json:"user_type"`
}

type WriteUserRequest struct {
	SubjectName      string              `json:"subject_name"`
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
