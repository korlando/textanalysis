package entities

type AnalyzeTextRequest struct {
	Text  string   `json:"text"`
	Texts []string `json:"texts"`
}
