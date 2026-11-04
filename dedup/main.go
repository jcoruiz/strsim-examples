// Dedup: detect and group duplicate entries in a dataset using multiple strategies.
//
// Demonstrates: FindAboveThreshold, phonetic matching, ASCIIOnly fast path.
package main

import (
	"fmt"

	"github.com/jcoruiz/strsim"
)

var contacts = []string{
	"John Smith",
	"Jon Smith",
	"John Smyth",
	"Jonathan Smith",
	"Jane Smith",
	"Michael Johnson",
	"Micheal Johnson",
	"Mike Johnson",
	"Robert Williams",
	"Bob Williams",
	"Roberto Williams",
}

func main() {
	fmt.Print("=== Duplicate Detection ===\n\n")

	// Strategy 1: Edit distance similarity (catches typos).
	fmt.Println("Strategy 1: Jaro-Winkler (typo detection)")
	jw := &strsim.JaroWinklerMetric{
		BoostThreshold: 0.7,
		PrefixSize:     4,
		ASCIIOnly:      true, // All ASCII names: use the fast path.
	}
	findDuplicates(contacts, 0.85, jw)

	// Strategy 2: Phonetic matching (catches different spellings of same name).
	fmt.Println("\nStrategy 2: Soundex (phonetic matching)")
	for i := 0; i < len(contacts); i++ {
		for j := i + 1; j < len(contacts); j++ {
			if strsim.SoundexMatch(contacts[i], contacts[j]) {
				fmt.Printf("  Soundex match: %q ↔ %q  [%s]\n",
					contacts[i], contacts[j],
					strsim.Soundex(contacts[i]))
			}
		}
	}

	// Strategy 3: Double Metaphone (handles more linguistic variations).
	fmt.Println("\nStrategy 3: Double Metaphone (linguistic variations)")
	for i := 0; i < len(contacts); i++ {
		for j := i + 1; j < len(contacts); j++ {
			if strsim.DoubleMetaphoneMatch(contacts[i], contacts[j]) {
				p1, a1 := strsim.DoubleMetaphone(contacts[i])
				fmt.Printf("  DM match: %q ↔ %q  [%s/%s]\n",
					contacts[i], contacts[j], p1, a1)
			}
		}
	}
}

func findDuplicates(items []string, threshold float64, m strsim.Metric) {
	seen := make(map[int]bool)

	for i, item := range items {
		if seen[i] {
			continue
		}

		remaining := items[i+1:]
		matches := strsim.FindAboveThreshold(item, remaining, threshold, m)

		if len(matches) > 0 {
			fmt.Printf("  Group: %q\n", item)
			for _, match := range matches {
				globalIdx := match.Index + i + 1
				seen[globalIdx] = true
				fmt.Printf("    ↳ %q (%.2f)\n", match.Value, match.Similarity)
			}
		}
	}
}
