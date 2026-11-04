// Name matcher: match person names across different systems using a
// multi-strategy approach combining edit distance and phonetic encoding.
//
// Demonstrates: combining Metric interface with Encoder for a real-world
// matching pipeline, custom scoring function, all four phonetic algorithms.
package main

import (
	"fmt"
	"strings"

	"github.com/jcoruiz/strsim"
)

// Record from system A.
type record struct {
	name  string
	email string
}

var systemA = []record{
	{"Robert Johnson", "rjohnson@corp.com"},
	{"Catherine O'Brien", "cobrien@corp.com"},
	{"Stephen García", "sgarcia@corp.com"},
	{"Michael Smith", "msmith@corp.com"},
	{"Elisabeth Brown", "ebrown@corp.com"},
}

// Names from system B that we need to match to system A.
var systemB = []string{
	"Bob Johnson",       // Nickname
	"Katherine O'Brian", // Spelling variation + typo
	"Steven Garcia",     // Anglicized + accent removed
	"Micheal Smyth",     // Typo + phonetic equivalent
	"Elizabeth Browne",  // Alternative spelling
	"Sarah Williams",    // No match expected
}

func main() {
	fmt.Print("=== Cross-System Name Matching ===\n\n")

	systemANames := make([]string, len(systemA))
	for i, r := range systemA {
		systemANames[i] = r.name
	}

	for _, nameB := range systemB {
		best, score := matchName(nameB, systemANames)

		if score >= 0.60 {
			rec := systemA[best.Index]
			fmt.Printf("  %-25s → %-25s (score: %.2f, email: %s)\n",
				nameB, best.Value, score, rec.email)
		} else {
			fmt.Printf("  %-25s → NO MATCH (best: %q, score: %.2f)\n",
				nameB, best.Value, score)
		}
	}

	// Show the phonetic codes for interesting pairs.
	fmt.Print("\n=== Phonetic Analysis ===\n\n")

	pairs := [][2]string{
		{"Smith", "Smyth"},
		{"Stephen", "Steven"},
		{"Catherine", "Katherine"},
		{"García", "Garcia"},
	}

	fmt.Printf("  %-12s %-12s %-8s %-8s %-10s %-10s %-8s %-8s\n",
		"Name A", "Name B", "Sdx A", "Sdx B", "Meta A", "Meta B", "NYSIIS A", "NYSIIS B")
	fmt.Println("  " + strings.Repeat("-", 86))

	for _, p := range pairs {
		fmt.Printf("  %-12s %-12s %-8s %-8s %-10s %-10s %-8s %-8s\n",
			p[0], p[1],
			strsim.Soundex(p[0]), strsim.Soundex(p[1]),
			strsim.Metaphone(p[0]), strsim.Metaphone(p[1]),
			strsim.NYSIIS(p[0]), strsim.NYSIIS(p[1]),
		)
	}
}

// matchName scores a name against candidates using a weighted combination of
// edit distance similarity and phonetic matching.
func matchName(query string, candidates []string) (strsim.Match, float64) {
	jw := strsim.NewJaroWinkler()
	best := strsim.FindBestMatch(query, candidates, jw)

	// Boost score if phonetic codes match.
	phoneticBonus := 0.0

	if strsim.SoundexMatch(query, best.Value) {
		phoneticBonus += 0.05
	}

	if strsim.MetaphoneMatch(query, best.Value) {
		phoneticBonus += 0.05
	}

	if strsim.DoubleMetaphoneMatch(query, best.Value) {
		phoneticBonus += 0.05
	}

	finalScore := best.Similarity + phoneticBonus
	if finalScore > 1.0 {
		finalScore = 1.0
	}

	return best, finalScore
}
