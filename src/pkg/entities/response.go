package entities

type AnalyzeTextResponse struct {
	Clout    float64 `json:"clout"`
	Tone     float64 `json:"tone"`
	Analytic float64 `json:"analytic"`
}
