package models

type Challenge struct {
	Id          string     `json:"id"`
	AuthorId    string     `json:"authorId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Languages   []string   `json:"languages"`
	Testcases   []Testcase `json:"testcases"`
}
