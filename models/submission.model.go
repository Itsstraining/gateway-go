package models

type Submission struct {
	Id               string           `json:"id"`
	UserID           string           `json:"userId"`
	ChallengeId      string           `json:"challengeId"`
	Timestamp        uint64           `json:"timestamp"`
	SubmissionResult SubmissionResult `json:"submissionResult,omitempty"`
}
