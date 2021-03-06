package responses

type QuestionListVO struct {
	QuestionId int64
	QuestionName string
}

type UserInfoVO struct {
		UserName string
		SubjectName string
}

type TeacherMonitoringVO struct {
	UserId string
	UserName string
	TestDistributionNumber int64
	TestSuccessNumber int64
	TestRemainingNumber int64
	TestProblemNumber int64
	MarkingSpeed  float64
	AverageScore float64
	Validity float64
	StandardDeviation float64
	EvaluationIndex float64
	OnlineTime float64
}

type ScoreDistributionVO struct {
	Score int64
	Rate float64
}

type TeacherVO struct {
	UserId string
	UserName string
}

type SelfScoreRecordVO struct {
	TestId int64
	Score int64
	SelfScore int64

}
type ScoreAverageVO struct {
	UserId string
	UserName string
	Average float64

}
type ProblemUnderCorrectedPaperVO struct {

	TestId int64
	ExaminerId string
	ExaminerName string
	ProblemType  int64
	ProblemMes string
}

type  ArbitramentTestVO struct {
	TestId int64
	ExaminerFirstId  string
	ExaminerFirstName  string
	ExaminerFirstScore int64
	ExaminerSecondId string
	ExaminerSecondName string
	ExaminerSecondScore int64
	PracticeError int64
	StandardError int64

}

type ScoreProgressVO struct {
	QuestionId int64
	QuestionName string
	SubjectName string
    ImportNumber int64
	AverageScore float64

	FinishNumber int64
	FinishRate float64
	UnfinishedNumber float64
	UnfinishedRate float64
	IsAllFinished int64
	//AverageSpeed float64
	FirstFinishedNumber int64
	FirstFinishedRate float64
	FirstUnfinishedNumber float64
	FirstUnfinishedRate float64
	IsFirstFinished int64

	SecondFinishedNumber int64
	SecondFinishedRate float64
	SecondUnfinishedNumber float64
	SecondUnfinishedRate float64
	IsSecondFinished int64

	ThirdFinishedNumber int64
	ThirdFinishedRate float64
	ThirdUnfinishedNumber float64
	ThirdUnfinishedRate float64
	IsThirdFinished int64

	ArbitramentNumber int64
	ArbitramentRate float64
	ArbitramentFinishedNumber int64
	ArbitramentFinishedRate float64
	ArbitramentUnfinishedNumber int64
	ArbitramentUnfinishedRate float64
	IsArbitramentFinished int64


	ProblemNumber int64
	ProblemRate float64
	ProblemFinishedNumber int64
	ProblemFinishedRate float64
	ProblemUnfinishedNumber int64
	ProblemUnfinishedRate float64
	IsProblemFinished int64
}

