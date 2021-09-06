package responses
type AddTopicVO struct {
	QuestionId int64
    QuestionDetailIds []AddTopicDetailVO
}
type AddTopicDetailVO struct {
	QuestionDetailId int64
}

type SubjectListVO struct {
	SubjectId  int64
	SubjectName string
}
type QuestionBySubListVO struct {
	QuestionId  int64
	QuestionName string
}
type DistributionInfoVO struct {
	ImportTestNumber  int64
	OnlineNumber int64
}
