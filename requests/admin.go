package requests
type AddTopic struct {
	AdminId string   `joson:"adminId"`
	TopicName string  `joson:"topicName"`
	ScoreType int64  `joson:"scoreType"`
	Score int64  `joson:"score"`
	Error int64  `joson:"error"`
	SubjectName string  `joson:"subjectName"`
	TopicDetails []AddTopicDetail `joson:"topicDetails"`
}
type  AddTopicDetail struct {
	TopicDetailName string  `joson:"topicDetailName"`
	DetailScoreTypes string  `joson:"DetailScoreTypes"`
	DetailScore int64  `joson:"detailScore"`
}

type SubjectList struct {
	AdminId string   `joson:"adminId"`
}

type QuestionBySubList struct {
	AdminId string   `joson:"adminId"`
	SubjectName string  `json:"subjectName"`
}
type DistributionInfo struct {
	AdminId string   `joson:"adminId"`
	QuestionId int64  `json:"questionId"`
}

type Distribution struct {
	AdminId string   `joson:"adminId"`
	QuestionId int64  `json:"questionId"`
	TestNumber int `json:"testNumber"`
	UserNumber int `json:"userNumber"`

}
type ReadExcel struct {
	AdminId string   `joson:"adminId"`
	FilePath string  `json:"filePath"`

}

type ReadFile struct {
	PicName string 	`json:"picName"`
}
