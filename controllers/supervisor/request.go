package supervisor

type QuestionList struct {
	SupervisorId int64 `joson:"supervisorId"`
}
type UserInfo struct {
	SupervisorId int64 `joson:"supervisorId"`
}
type ScoreDistribution struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `joson:"questionId"`
}
type TeachersByQuestion struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `joson:"questionId"`
}
type SelfScore struct {
	SupervisorId int64 `joson:"supervisorId"`
	ExaminerId   int64 `joson:"examinerId"`
}
type AverageScore struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `joson:"questionId"`
}
type ProblemTest struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `joson:"questionId"`
}
type ArbitramentTest struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `joson:"questionId"`
}
type TeacherMonitoring struct {
	SupervisorId int64 `joson:"supervisorId"`

	QuestionId int64 `joson:"questionId"`
}
type SupervisorPoint struct {
	SupervisorId  int64  `joson:"supervisorId"`
	TestId        int64  `joson:"testId"`
	TestDetailIds string `joson:"testDetailIds"`
	Scores        string `joson:"scores"`
}

type ScoreProgress struct {
	SupervisorId int64  `joson:"supervisorId"`
	Subject      string `json:"subject"`
}
type ArbitramentUnmarkList struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `json:"questionId"`
}
type SelfMarkList struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `json:"questionId"`
}
type ProblemUnmarkList struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `json:"questionId"`
}
type SelfUnmarkList struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `json:"questionId"`
}
type ScoreDeviation struct {
	SupervisorId int64 `joson:"supervisorId"`
	QuestionId   int64 `joson:"questionId"`
}
