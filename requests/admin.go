package requests

type AddTopic struct {
	AdminId string   `json:"adminId"`
	TopicName string  `json:"topicName"`
	ScoreType int64  `json:"scoreType"`
	Score int64  `json:"score"`
	Error int64  `json:"error"`
	SubjectName string  `json:"subjectName"`
	TopicDetails []AddTopicDetail `json:"topicDetails"`
}
type  AddTopicDetail struct {
	TopicDetailName string  `json:"topicDetailName"`
	DetailScoreTypes string  `json:"DetailScoreTypes"`
	DetailScore int64  `json:"detailScore"`
}

type SubjectList struct {
	AdminId string   `json:"adminId"`
}

type QuestionBySubList struct {
	AdminId string   `json:"adminId"`
	SubjectName string  `json:"subjectName"`
}
type DistributionInfo struct {
	AdminId string   `json:"adminId"`
	QuestionId int64  `json:"questionId"`
}

type Distribution struct {
	AdminId string   `json:"adminId"`
	QuestionId int64  `json:"questionId"`
	TestNumber int `json:"testNumber"`
	UserNumber int `json:"userNumber"`

}
//type ReadExcel struct {
//	AdminId string   `json:"adminId"`
////	FilePath string  `json:"filePath"`
//	Excel []byte    `json:"excel"`
//}
//
//type ReadExcelBytes struct {
//	AdminId string   ` form:"adminId"  binding:"required"`
//	Excel []byte    `form:"excel"   binding:"required"`
//}
type ReadFile struct {
	PicName string 	`json:"picName"`
}

type TopicList struct {
	AdminId string   `json:"adminId"`
}

type DistributionRecord struct {
	AdminId string   `json:"adminId"`
	SubjectName string  `json:"subjectName"`
}