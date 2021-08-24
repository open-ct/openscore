package requests
type QuestionList struct {
	SupervisorId string   `joson:"supervisorId"`
}
type UserInfo struct {
	SupervisorId string   `joson:"supervisorId"`
}
type ScoreDistribution struct {
	SupervisorId string   `joson:"supervisorId"`
	QuestionId int64  `joson:"questionId"`
}
type TeachersByQuestion struct {
	SupervisorId string   `joson:"supervisorId"`
	QuestionId int64  `joson:"questionId"`
}
type SelfScore struct {
	SupervisorId string   `joson:"supervisorId"`
	ExaminerId string  `joson:"examinerId"`
}
type AverageScore struct {
	SupervisorId string   `joson:"supervisorId"`
	QuestionId int64  `joson:"questionId"`
}
type ProblemTest struct {
	SupervisorId string   `joson:"supervisorId"`
	QuestionId int64  `joson:"questionId"`
}
type ArbitramentTest struct {
	SupervisorId string   `joson:"supervisorId"`
	QuestionId int64  `joson:"questionId"`
}
type TeacherMonitoring struct {
	SupervisorId string   `joson:"supervisorId"`

	QuestionId int64  `joson:"questionId"`
}
type SupervisorPoint struct {
	SupervisorId string   `joson:"supervisorId"`
	TestId int64   `joson:"testId"`
	TestDetailIds string  `joson:"testDetailIds"`
	Scores string  `joson:"scores"`

}

type ScoreProgress struct {
	SupervisorId string   `joson:"supervisorId"`
}