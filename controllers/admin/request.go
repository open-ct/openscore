package admin

type AddTopic struct {
	AdminId       string           `json:"adminId"`
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

type SubjectList struct {
	AdminId string `json:"adminId"`
}

type QuestionBySubList struct {
	AdminId     string `json:"adminId"`
	SubjectName string `json:"subjectName"`
}

type DistributionInfo struct {
	AdminId    string `json:"adminId"`
	QuestionId int64  `json:"questionId"`
}

type DeleteTest struct {
	AdminId    string `json:"adminId"`
	QuestionId int64  `json:"questionId"`
}

type Distribution struct {
	AdminId    string `json:"adminId"`
	QuestionId int64  `json:"questionId"`
	TestNumber int    `json:"testNumber"`
	UserNumber int    `json:"userNumber"`
}

type ReadFile struct {
	PicName string `json:"picName"`
}

type TopicList struct {
	AdminId string `json:"adminId"`
}

type DistributionRecord struct {
	AdminId     string `json:"adminId"`
	SubjectName string `json:"subjectName"`
}
