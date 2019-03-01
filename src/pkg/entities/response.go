package entities

type AnalyzeTextResponse struct {
	Clout     float64 `json:"clout"`
	Tone      float64 `json:"tone"`
	Analytic  float64 `json:"analytic"`
	Authentic float64 `json:"authentic"`
	NumWords  float64 `json:"numWords"`
	NumChars  float64 `json:"numChars"`
	AvgLength float64 `json:"avgWordLength"`
}

type AnalyzeTextsMetadata struct {
	AvgLength float64 `json:"avgLength"`
	NumWords  float64 `json:"numWords"`
	NumChars  float64 `json:"numChars"`
}

type AnalyzeSummaryResponse struct {
	AnalyzeTexts []AnalyzeTextResponse `json:"analyzeTexts"`
	Metadata     AnalyzeTextsMetadata  `json:"metadata"`
}
