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
	words := textToWords(req.Text, true, true)
	total := float64(len(words))
	clout := float64(analysis.GetClout(words)) / total
	tone := float64(analysis.GetTone(words)) / total
	analytic := float64(analysis.GetAnalytic(words)) / total
	res := entities.AnalyzeTextResponse{
		Clout:    clout,
		Tone:     tone,
		Analytic: analytic,
	}
	resJson, err := encodeBody(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

func decodeBody(r *http.Request, entity interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&entity)
	return
}

func encodeBody(payload interface{}) ([]byte, error) {
	return json.Marshal(payload)
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
