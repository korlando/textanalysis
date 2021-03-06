package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"pkg/analysis"
	"pkg/entities"
	"pkg/summary"
)

func AnalyzeText(w http.ResponseWriter, r *http.Request) {
	var req entities.AnalyzeTextRequest
	err := decodeBody(r, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	texts := []string{}
	if req.Text != "" {
		texts = append(texts, req.Text)
	}
	if len(req.Texts) > 0 {
		texts = append(texts, req.Texts...)
	}
	analyzeTexts := []entities.AnalyzeTextResponse{}
	totalNumWords := float64(0)
	totalNumChars := float64(0)
	for _, text := range texts {
		words := textToWords(text, true, true)
		numWords := float64(len(words))
		totalNumWords += numWords
		clout := float64(analysis.GetClout(words)) / numWords
		tone := float64(analysis.GetTone(words)) / numWords
		analytic := float64(analysis.GetAnalytic(words)) / numWords
		authentic := float64(analysis.GetAuthentic(words)) / numWords
		numChars := float64(0)
		for _, w := range words {
			numChars += float64(len(w))
		}
		totalNumChars += numChars
		avgLength := numChars / numWords
		analyzeTexts = append(analyzeTexts, entities.AnalyzeTextResponse{
			Clout:     clout,
			Tone:      tone,
			Analytic:  analytic,
			Authentic: authentic,
			NumWords:  numWords,
			NumChars:  numChars,
			AvgLength: avgLength,
		})
	}
	res := entities.AnalyzeSummaryResponse{
		AnalyzeTexts: analyzeTexts,
		Metadata: entities.AnalyzeTextsMetadata{
			AvgLength: totalNumChars / totalNumWords,
			NumWords:  totalNumWords,
			NumChars:  totalNumChars,
		},
	}
	jsonResponse(w, res)
}

func decodeBody(r *http.Request, entity interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&entity)
	return
}

func encodeBody(payload interface{}) ([]byte, error) {
	return json.Marshal(payload)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	resJson, err := encodeBody(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

// textToWords takes in text and returns a slice of words (strings);
// optionally trims symbols from individual words and converts to lowercase
func textToWords(text string, trimSymbols, lower bool) (words []string) {
	words = strings.Fields(text)
	if !trimSymbols && !lower {
		return
	}
	for i, w := range words {
		result := w
		if trimSymbols {
			start := 0
			end := len(w)
			for j, b := range w {
				if !isSymbol(b) {
					start = j
					break
				}
			}
			for j := len(w) - 1; j >= start; j-- {
				if !isSymbol(rune(w[j])) {
					end = j + 1
					break
				}
			}
			result = string(w[start:end])
		}
		if lower {
			words[i] = strings.ToLower(result)
		} else {
			words[i] = result
		}
	}
	return
}

// isSymbol decides whether a
// rune is in the symbol list
func isSymbol(r rune) bool {
	for _, s := range summary.Symbols {
		if s == r {
			return true
		}
	}
	return false
}
