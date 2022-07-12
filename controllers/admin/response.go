package admin

import (
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
