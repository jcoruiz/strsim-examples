// Spell checker: suggest corrections for misspelled words using multiple algorithms.
//
// Demonstrates: DistanceMetric interface, comparing edit distances,
// combining edit distance with phonetic encoding for better suggestions.
package main

import (
	"fmt"

	"github.com/jcoruiz/strsim"
)

var words = []string{
	"accept", "access", "account", "achieve", "across",
	"action", "address", "advance", "advice", "affect",
	"already", "always", "analyze", "another", "answer",
	"appear", "approach", "arrange", "attempt", "balance",
	"because", "before", "believe", "benefit", "between",
	"business", "calendar", "category", "certain", "change",
	"collect", "column", "comment", "commit", "complete",
	"concern", "connect", "consider", "contain", "continue",
	"control", "correct", "create", "current", "customer",
	"database", "default", "define", "deliver", "describe",
	"develop", "different", "discuss", "display", "document",
	"environment", "example", "execute", "exercise", "experience",
	"function", "general", "generate", "important", "improve",
	"include", "increase", "indicate", "industry", "information",
	"language", "library", "measure", "message", "necessary",
	"network", "observe", "operate", "organize", "original",
	"package", "perform", "possible", "practice", "prepare",
	"process", "produce", "program", "project", "provide",
	"question", "receive", "recognize", "recommend", "reference",
	"release", "remember", "require", "resource", "response",
	"schedule", "separate", "similar", "specific", "standard",
	"strategy", "strength", "structure", "succeed", "suggest",
	"surprise", "technical", "together", "transfer", "understand",
}

func main() {
	misspelled := []string{
		"recieve",    // common: ie/ei swap
		"definately", // common: ite vs ately
		"enviroment", // missing letter
		"langauge",   // transposition
		"seperate",   // common: e vs a
		"neccessary", // double letter error
		"reccomend",  // double letter error
		"strucutre",  // transposition
	}

	fmt.Print("=== Spell Checker ===\n\n")

	// Use multiple distance metrics to compare suggestions.
	metrics := []struct {
		name string
		dm   strsim.DistanceMetric
	}{
		{"Levenshtein", strsim.NewLevenshtein()},
		{"OSA", strsim.NewOSA()},
		{"Damerau-Lev", strsim.NewDamerauLevenshtein()},
	}

	for _, word := range misspelled {
		fmt.Printf("  %q:\n", word)

		for _, m := range metrics {
			best := strsim.FindBestMatch(word, words, m.dm)
			dist := m.dm.Distance(word, best.Value)
			fmt.Printf("    %-12s → %q (distance: %d, similarity: %.2f)\n",
				m.name, best.Value, dist, best.Similarity)
		}
		fmt.Println()
	}
}
