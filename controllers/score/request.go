package score

type TestDisplay struct {
	UserId string `json:"userId"`
	TestId int64  `json:"testId"`
}

type TestList struct {
	UserId string `json:"userId"`
}

type TeacherSelfMarkList struct {
	UserId string `json:"userId"`
}

type TestPoint struct {
	UserId       string `json:"userId"`
	Scores       string `json:"scores"`
	TestId       int64  `json:"testId"`
	TestDetailId string `json:"testDetailId"`
}

type TestProblem struct {
	UserId         string `json:"userId"`
	ProblemType    int64  `json:"problemType"`
	TestId         int64  `json:"testId"`
	ProblemMessage string `json:"problemMessage"`
}

type TestAnswer struct {
	UserId string `json:"userId"`
	TestId int64  `json:"testId"`
}

type ExampleDetail struct {
	UserId        string `json:"userId"`
	ExampleTestId int64  `json:"exampleTestId"`
}

type ExampleList struct {
	UserId string `json:"userId"`
	TestId int64  `json:"TestId"`
}

type TestReview struct {
	UserId string `json:"userId"`
}
