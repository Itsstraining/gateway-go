package models

type SubmissionResult struct {
	Id              string           `json:"id"`
	Runtime         uint             `json:"runtime"`
	TestcaseResults []TestcaseResult `json:"testcaseResults"`
	Timestamp       uint             `json:"timestamp"`
}
