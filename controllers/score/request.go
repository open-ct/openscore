package score

type TestDisplay struct {
	UserId int64 `json:"userId"`
	TestId int64 `json:"testId"`
}

type TestList struct {
	UserId int64 `json:"userId"`
}

type TeacherSelfMarkList struct {
	UserId int64 `json:"userId"`
}

type TestPoint struct {
	UserId       int64  `json:"userId"`
	Scores       string `json:"scores"`
	TestId       int64  `json:"testId"`
	TestDetailId string `json:"testDetailId"`
}

type TestProblem struct {
	UserId         int64  `json:"userId"`
	ProblemType    int64  `json:"problemType"`
	TestId         int64  `json:"testId"`
	ProblemMessage string `json:"problemMessage"`
}

type TestAnswer struct {
	UserId int64 `json:"userId"`
	TestId int64 `json:"testId"`
}

type ExampleDetail struct {
	UserId        int64 `json:"userId"`
	ExampleTestId int64 `json:"exampleTestId"`
}

type ExampleList struct {
	UserId int64 `json:"userId"`
	TestId int64 `json:"TestId"`
}

type TestReview struct {
	UserId int64 `json:"userId"`
}
