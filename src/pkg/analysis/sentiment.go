package analysis

import (
	"strings"

	"pkg/summary"
)

// GetClout loosely calculates the clout of a
// slice of words (+1 for we/you and -1 for i)
func GetClout(words []string) (clout int) {
	for _, w := range words {
		wLower := strings.ToLower(w)
		if inList(wLower, summary.IWords) {
			clout -= 1
			continue
		}
		if inList(wLower, summary.YouWords) || inList(wLower, summary.WeWords) {
			clout += 1
		}
	}
	return
}

// GetTone loosely calculates the tone of a slice
// of words (+1 for pos emo and -1 for neg emo)
func GetTone(words []string) (tone int) {
	for _, w := range words {
		wLower := strings.ToLower(w)
		if inList(wLower, summary.PositiveEmotions) {
			tone += 1
			continue
		}
		if inList(wLower, summary.NegativeEmotions) {
			tone -= 1
		}
	}
	return
}

// GetAnalytic loosely calculates the analytical thinking
// of a slice of words (+1 for article, prep and -1 for
// ppron, ipron, auxverb, conj, adverb, negation)
func GetAnalytic(words []string) (analytic int) {
	analytic = 30
	for _, w := range words {
		w = strings.ToLower(w)
		if inList(w, summary.Articles) || inList(w, summary.Prepositions) {
			analytic += 1
			continue
		}
		if inList(w, summary.PersonalPronouns) || inList(w, summary.ImpersonalPronouns) || inList(w, summary.AuxiliaryVerbs) || inList(w, summary.Conjunctions) || inList(w, summary.Adverbs) || inList(w, summary.Negations) {
			analytic -= 1
			continue
		}
	}
	return
}

func GetAuthentic(words []string) (authentic int) {
	for _, w := range words {
		w = strings.ToLower(w)
		if inList(w, summary.IWords) || inList(w, summary.SheHe) || inList(w, summary.TheyWords) || inList(w, summary.Exclusive) {
			authentic += 1
			continue
		}
		if inList(w, summary.NegativeEmotions) || inList(w, summary.Motion) {
			authentic -= 1
			continue
		}
	}
	return
}

// inList determines whether a word
// is contained in a slice of words
func inList(word string, words []string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}
