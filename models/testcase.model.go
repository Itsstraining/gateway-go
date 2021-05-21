package models

type Testcase struct {
	Id             string  `json:"id"`
	Title          string  `json:"title,omitempty"`
	Input          string  `json:"input"`
	ExpectedOutput string  `json:"expectedOuput"`
	Timeout        uint    `json:"timeout"`
	Visible        bool    `json:"visible"`
	Score          float32 `json:"score"`
}
