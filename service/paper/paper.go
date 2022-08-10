package paper

import (
	"errors"
	"log"

	"github.com/open-ct/openscore/model"
)

func FindUnDistributeTest(questionId int64) ([]model.TestPaper, error) {
	// 是否需要二次阅卷
	var topic model.Topic
	topic.QuestionId = questionId
	if err := topic.GetTopic(questionId); err != nil {
		log.Println("试卷分配异常,无法获取试卷批改次数 ", err)
		return nil, err
	}

	if topic.ScoreType == 1 {
		return model.FindUnDistributeTestLimit1(questionId)
	}

	if topic.ScoreType == 2 {
		return model.FindUnDistributeTestLimit2(questionId)
	}
	return nil, errors.New("wrong ScoreType")
}
