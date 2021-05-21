package models

type TestcaseResult struct {
	Id           string   `json:"id"`
	Testcase     Testcase `json:"testcase"`
	ActualOutput string   `json:"actualOutput"`
	StdErr       string   `json:"stdErr"`
	ExecutedTime uint     `json:"excutedTime"`
}
